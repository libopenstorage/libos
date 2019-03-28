// Package cluster implements a cluster state machine.  It relies on a cluster
// wide keyvalue store for coordinating the state of the cluster.
// It also stores the state of the cluster in this keyvalue store.
package manager

import (
	"container/list"
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/libopenstorage/gossip"
	"github.com/libopenstorage/gossip/types"
	"github.com/libopenstorage/openstorage/api"
	"github.com/libopenstorage/openstorage/cluster"
	"github.com/libopenstorage/openstorage/config"
	"github.com/libopenstorage/openstorage/objectstore"
	"github.com/libopenstorage/openstorage/osdconfig"
	"github.com/libopenstorage/openstorage/pkg/auth"
	"github.com/libopenstorage/openstorage/pkg/clusterdomain"
	sched "github.com/libopenstorage/openstorage/schedpolicy"
	"github.com/libopenstorage/openstorage/secrets"
	"github.com/libopenstorage/systemutils"
	"github.com/portworx/kvdb"
	"github.com/sirupsen/logrus"
)

const (
	heartbeatKey       = "heartbeat"
	clusterLockKey     = "/cluster/lock"
	gossipVersionKey   = "Gossip Version"
	decommissionErrMsg = "Node %s must be offline or in maintenance " +
		"mode to be decommissioned."
)

var (
	stopHeartbeat            = make(chan bool)
	errClusterInitialized    = errors.New("openstorage.cluster: already initialized")
	errClusterNotInitialized = errors.New("openstorage.cluster: not initialized")

	inst *ClusterManager
	// Inst returns an instance of an already instantiated cluster manager.
	// This function can be overridden for testing purposes
	Inst = func() (cluster.Cluster, error) {
		return clusterInst()
	}
)

// ClusterManager implements the cluster interface
type ClusterManager struct {
	size                 int
	listeners            *list.List
	config               config.ClusterConfig
	kv                   kvdb.Kvdb
	status               api.Status
	nodeCache            map[string]api.Node // Cached info on the nodes in the cluster.
	nodeCacheLock        sync.Mutex
	nodeStatuses         map[string]api.Status // Set of nodes currently marked down.
	gossip               gossip.Gossiper
	gossipVersion        string
	gossipPort           string
	gEnabled             bool
	selfNode             api.Node
	selfNodeLock         sync.Mutex // Lock that guards data and label of selfNode
	system               systemutils.System
	configManager        osdconfig.ConfigCaller
	schedManager         sched.SchedulePolicyProvider
	objstoreManager      objectstore.ObjectStore
	secretsManager       secrets.Secrets
	systemTokenManager   auth.TokenGenerator
	clusterDomainManager clusterdomain.ClusterDomainProvider
	snapshotPrefixes     []string
	selfClusterDomain    string
}

// Init instantiates a new cluster manager.
func Init(cfg config.ClusterConfig) error {
	if inst != nil {
		return errClusterInitialized
	}

	kv := kvdb.Instance()
	if kv == nil {
		return errors.New("KVDB is not yet initialized.  " +
			"A valid KVDB instance required for the cluster to start.")
	}

	inst = &ClusterManager{
		listeners:          list.New(),
		config:             cfg,
		kv:                 kv,
		nodeCache:          make(map[string]api.Node),
		nodeStatuses:       make(map[string]api.Status),
		systemTokenManager: auth.SystemTokenManagerInst(),
	}

	return nil
}

func clusterInst() (cluster.Cluster, error) {
	if inst == nil {
		return nil, errClusterNotInitialized
	}
	return inst, nil
}

type checkFunc func(cluster.ClusterInfo) error

func ifaceToIp(iface *net.Interface) (string, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip == nil || ip.IsLoopback() {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}
		if ip.String() == "" {
			continue // address is empty string
		}
		return ip.String(), nil
	}

	return "", errors.New("Node not connected to the network.")
}

func ifaceNameToIp(ifaceName string) (string, error) {
	stdout, err := exec.Command("/usr/sbin/ip", "a", "show", ifaceName, "label", ifaceName).Output()
	if err != nil {
		return "", err
	}
	ipOp := string(stdout)
	// Parse the output of command /usr/bin/ip a show eth0 label eth0:0
	ipOpParts := strings.Fields(ipOp)
	for i, tokens := range ipOpParts {
		// Only check for ipv4 addresses
		if tokens == "inet" {
			ip := ipOpParts[i+1]
			// Remove the mask
			ipAddr := strings.Split(ip, "/")
			if strings.Contains(ipAddr[0], "127") {
				// Loopback address
				continue
			}
			if ipAddr[0] == "" {
				// Address is empty string
				continue
			}
			return ipAddr[0], nil
		}
	}
	return "", fmt.Errorf("Unable to find Ip address for given interface")
}

// ExternalIp returns the mgmt and data ip based on the config
func ExternalIp(config *config.ClusterConfig) (string, string, error) {
	mgmtIp := ""
	dataIp := ""

	var err error
	if config.MgmtIp == "" && config.MgtIface != "" {
		mgmtIp, err = ifaceNameToIp(config.MgtIface)
		if err != nil {
			return "", "", errors.New("Invalid data network interface " +
				"specified.")
		}
	} else if config.MgmtIp != "" {
		mgmtIp = config.MgmtIp
	}

	if config.DataIp == "" && config.DataIface != "" {
		dataIp, err = ifaceNameToIp(config.DataIface)
		if err != nil {
			return "", "", errors.New("Invalid data network interface " +
				"specified.")
		}
	} else if config.DataIp != "" {
		dataIp = config.DataIp
	}

	if mgmtIp != "" && dataIp != "" {
		return mgmtIp, dataIp, nil
	} else if mgmtIp != "" { // dataIp is empty
		return mgmtIp, mgmtIp, nil
	} else if dataIp != "" { // mgmtIp is empty
		return dataIp, dataIp, nil
	} // both are empty, try to pick first available interface for both

	// No network interface specified, pick first default.
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		mgmtIp, err = ifaceToIp(&iface)
		if err != nil {
			logrus.Printf("Skipping interface without IP: %v: %v",
				iface, err)
			continue
		}
		return mgmtIp, mgmtIp, err
	}

	return "", "", errors.New("Node not connected to the network.")
}

// getNodeEntry is internal helper method, shared between Inspect() and enumerateNodesFromCache()
// Parameter 'clustDBRef' may be a pointer to "empty" struct, in which case it'll be populated, but it must not be NULL.
// Also, it's caller's responsibility to lock the access to the NodeCache.
func (c *ClusterManager) getNodeEntry(nodeID string, clustDBRef *cluster.ClusterInfo) (api.Node, error) {
	var n api.Node
	var ok bool

	nodeID, _ = c.nodeIdFromIp(nodeID)

	if nodeID == c.selfNode.Id {
		n = *c.getCurrentState()
	} else if n, ok = c.nodeCache[nodeID]; !ok {
		return api.Node{}, errors.New("Unable to locate node with provided UUID.")
	} else if n.Status == api.Status_STATUS_OFFLINE &&
		(n.DataIp == "" || n.MgmtIp == "") {
		// cached info unstable, read from DB
		if clustDBRef.Id == "" {
			// We've been passed "empty" struct, lazy-init before use
			clusterDB, _, _ := readClusterInfo()
			*clustDBRef = clusterDB
		}
		// Gossip does not have essential information of
		// an offline node. Provide the essential data
		// that we have in the cluster db
		if v, ok := clustDBRef.NodeEntries[n.Id]; ok {
			n.SchedulerNodeName = v.SchedulerNodeName
			n.MgmtIp = v.MgmtIp
			n.DataIp = v.DataIp
			n.Hostname = v.Hostname
			n.NodeLabels = v.NodeLabels
		} else {
			logrus.Warnf("Could not query NodeID %v", nodeID)
			// Node entry won't be refreshed form DB, will use the "offline" original
		}
	}
	return n, nil
}

// Inspect inspects given node and returns the state
func (c *ClusterManager) Inspect(nodeID string) (api.Node, error) {
	c.nodeCacheLock.Lock()
	defer c.nodeCacheLock.Unlock()
	return c.getNodeEntry(nodeID, &cluster.ClusterInfo{})
}

// AddEventListener adds a new listener
func (c *ClusterManager) AddEventListener(listener cluster.ClusterListener) error {
	logrus.Printf("Adding cluster event listener: %s", listener.String())
	c.listeners.PushBack(listener)
	return nil
}

// UpdateData updates self node data
func (c *ClusterManager) UpdateData(nodeData map[string]interface{}) error {
	c.selfNodeLock.Lock()
	defer c.selfNodeLock.Unlock()
	for dataKey, dataValue := range nodeData {
		c.selfNode.NodeData[dataKey] = dataValue
	}
	return nil
}

func (c *ClusterManager) UpdateLabels(nodeLabels map[string]string) error {
	c.selfNodeLock.Lock()
	defer c.selfNodeLock.Unlock()
	if c.selfNode.NodeLabels == nil {
		c.selfNode.NodeLabels = make(map[string]string)
	}
	for labelKey, labelValue := range nodeLabels {
		c.selfNode.NodeLabels[labelKey] = labelValue
	}
	return nil
}

func (c *ClusterManager) UpdateSchedulerNodeName(schedulerNodeName string) error {
	c.selfNodeLock.Lock()
	defer c.selfNodeLock.Unlock()
	c.selfNode.SchedulerNodeName = schedulerNodeName

	updateCallbackFn := func(db *cluster.ClusterInfo) (bool, error) {
		nodeEntry, ok := db.NodeEntries[c.selfNode.Id]
		if !ok {
			return false, fmt.Errorf("Node not found in cluster database")
		}
		nodeEntry.SchedulerNodeName = schedulerNodeName
		db.NodeEntries[c.selfNode.Id] = nodeEntry
		return true, nil
	}

	return updateLockedDB("update-scheduler-name", c.selfNode.Id, updateCallbackFn)
}

// GetData returns self node's data
func (c *ClusterManager) GetData() (map[string]*api.Node, error) {
	nodes := make(map[string]*api.Node)
	c.nodeCacheLock.Lock()
	defer c.nodeCacheLock.Unlock()
	for _, value := range c.nodeCache {
		copyValue := value.Copy()
		nodes[value.Id] = copyValue
	}
	return nodes, nil
}

func (c *ClusterManager) nodeIdFromIp(idIp string) (string, error) {
	// Caller's responsibility to lock the access to the NodeCache.
	for _, n := range c.nodeCache {
		if n.DataIp == idIp || n.MgmtIp == idIp {
			return n.Id, nil // return Id
		}
	}

	return idIp, errors.New("Failed to locate IP in this cluster.") // return input value
}

// GetNodeIdFromIp returns a Node Id given an IP.
func (c *ClusterManager) GetNodeIdFromIp(idIp string) (string, error) {
	addr := net.ParseIP(idIp)
	if addr != nil { // Is an IP, lookup Id
		c.nodeCacheLock.Lock()
		defer c.nodeCacheLock.Unlock()
		return c.nodeIdFromIp(idIp)
	}

	return idIp, nil // return input, assume its an Id
}

// getCurrentState always returns the copy of selfNode that
// cluster manager maintains. It also updates the selfNode
// with latest data.
func (c *ClusterManager) getCurrentState() *api.Node {
	c.selfNodeLock.Lock()
	defer c.selfNodeLock.Unlock()
	c.selfNode.Timestamp = time.Now()

	c.selfNode.Cpu, _, _ = c.system.CpuUsage()
	c.selfNode.MemTotal, c.selfNode.MemUsed, c.selfNode.MemFree = c.system.MemUsage()

	c.selfNode.Timestamp = time.Now()

	for e := c.listeners.Front(); e != nil; e = e.Next() {
		listenerDataMap := e.Value.(cluster.ClusterListener).ListenerData()
		if listenerDataMap == nil {
			continue
		}
		for key, val := range listenerDataMap {
			c.selfNode.NodeData[key] = val
		}
	}

	nodeCopy := (&c.selfNode).Copy()
	return nodeCopy
}

func (c *ClusterManager) getNonDecommisionedPeers(
	db cluster.ClusterInfo,
) map[types.NodeId]types.NodeUpdate {
	peers := make(map[types.NodeId]types.NodeUpdate)
	for _, nodeEntry := range db.NodeEntries {
		if nodeEntry.Status == api.Status_STATUS_DECOMMISSION {
			continue
		}
		peers[types.NodeId(nodeEntry.Id)] = types.NodeUpdate{
			Addr:          nodeEntry.DataIp + ":" + c.gossipPort,
			QuorumMember:  !nodeEntry.NonQuorumMember,
			ClusterDomain: nodeEntry.ClusterDomain,
		}
	}
	return peers
}

// Get the latest config.
func (c *ClusterManager) watchDB(key string, opaque interface{},
	kvp *kvdb.KVPair, watchErr error) error {

	db, kvdbVersion, err := readClusterInfo()

	if err != nil {
		logrus.Warnln("Failed to read database after update ", err)
		// Exit since an update may be missed here.
		os.Exit(1)
	}

	// Update all the listeners with the new db
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		err := e.Value.(cluster.ClusterListener).UpdateCluster(&c.selfNode, &db)
		if err != nil {
			logrus.Warnln("Failed to notify ", e.Value.(cluster.ClusterListener).String())
		}
	}

	for _, nodeEntry := range db.NodeEntries {
		if nodeEntry.Status == api.Status_STATUS_DECOMMISSION {
			logrus.Infof("ClusterManager watchDB, node ID "+
				"%s state is Decommission.",
				nodeEntry.Id)

			n, found := c.getNodeCacheEntry(nodeEntry.Id)
			if !found {
				logrus.Errorf("ClusterManager watchDB, "+
					"node ID %s not in node cache",
					nodeEntry.Id)
				continue
			}

			if n.Status == api.Status_STATUS_DECOMMISSION {
				logrus.Infof("ClusterManager watchDB, "+
					"node ID %s is already decommission "+
					"on this node",
					nodeEntry.Id)
				continue
			}

			logrus.Infof("ClusterManager watchDB, "+
				"decommsission node ID %s on this node",
				nodeEntry.Id)

			n.Status = api.Status_STATUS_DECOMMISSION
			c.putNodeCacheEntry(nodeEntry.Id, n)
			// We are getting decommissioned!!
			if nodeEntry.Id == c.selfNode.Id {
				// We are getting decommissioned.
				// Stop the heartbeat and stop the watch
				stopHeartbeat <- true
				c.gossip.Stop(time.Duration(10 * time.Second))
				return fmt.Errorf("stop watch")
			}
		}
	}

	c.size = db.Size

	c.gossip.UpdateCluster(c.getNonDecommisionedPeers(db))

	// update the nodeCache and remove any nodes not present in cluster database
	c.nodeCacheLock.Lock()
	defer c.nodeCacheLock.Unlock()
	for _, n := range c.nodeCache {
		_, found := db.NodeEntries[n.Id]
		if !found {
			delete(c.nodeCache, n.Id)
		}
	}

	if watchErr != nil && c.selfNode.Status != api.Status_STATUS_DECOMMISSION {
		logrus.Errorf("ClusterManager watch stopped, restarting (err: %v)",
			watchErr)
		c.startClusterDBWatch(kvdbVersion, kvdb.Instance())
	}
	return watchErr
}

func (c *ClusterManager) getLatestNodeConfig(nodeId string) *cluster.NodeEntry {
	db, _, err := readClusterInfo()
	if err != nil {
		logrus.Warnln("Failed to read the database for updating config")
		return nil
	}

	ne, exists := db.NodeEntries[nodeId]
	if !exists {
		logrus.Warnln("Could not find info for node with id ", nodeId)
		return nil
	}

	return &ne
}

func (c *ClusterManager) initNode(db *cluster.ClusterInfo) (*api.Node, bool) {
	_, exists := db.NodeEntries[c.selfNode.Id]

	// Add us into the database.
	labels := make(map[string]string)
	labels[gossipVersionKey] = c.gossipVersion
	nodeEntry := cluster.NodeEntry{
		Id:                c.selfNode.Id,
		SchedulerNodeName: c.selfNode.SchedulerNodeName,
		MgmtIp:            c.selfNode.MgmtIp,
		DataIp:            c.selfNode.DataIp,
		GenNumber:         c.selfNode.GenNumber,
		StartTime:         c.selfNode.StartTime,
		MemTotal:          c.selfNode.MemTotal,
		Hostname:          c.selfNode.Hostname,
		NodeLabels:        labels,
		GossipPort:        c.selfNode.GossipPort,
		ClusterDomain:     c.selfClusterDomain,
	}

	db.NodeEntries[c.config.NodeId] = nodeEntry

	logrus.Infof("Node %s joining cluster...", c.config.NodeId)
	logrus.Infof("Cluster ID: %s", c.config.ClusterId)
	logrus.Infof("Node Mgmt IP: %s", c.selfNode.MgmtIp)
	logrus.Infof("Node Data IP: %s", c.selfNode.DataIp)
	if len(c.selfClusterDomain) > 0 {
		logrus.Infof("Node's Cluster Domain: %s", c.selfClusterDomain)
	}

	return &c.selfNode, exists
}

func (c *ClusterManager) cleanupInit(db *cluster.ClusterInfo, self *api.Node) error {
	var resErr error
	var err error

	logrus.Infof("Cleanup Init services")

	for e := c.listeners.Front(); e != nil; e = e.Next() {
		logrus.Warnf("Cleanup Init for service %s.",
			e.Value.(cluster.ClusterListener).String())

		err = e.Value.(cluster.ClusterListener).CleanupInit(self, db)
		if err != nil {
			logrus.Warnf("Failed to Cleanup Init %s: %v",
				e.Value.(cluster.ClusterListener).String(), err)
			resErr = err
		}

	}

	return resErr
}

// Initialize node and alert listeners that we are initializing a node in the cluster.
func (c *ClusterManager) initNodeInCluster(
	clusterInfo *cluster.ClusterInfo,
	self *api.Node,
	exist bool,
	nodeInitialized bool,
) ([]cluster.FinalizeInitCb, error) {
	// If I am already in the cluster map, don't add me again.
	if exist {
		return nil, nil
	}

	if nodeInitialized {
		logrus.Errorf(cluster.ErrInitNodeNotFound.Error())
		return nil, cluster.ErrInitNodeNotFound
	}

	// Alert all listeners that we are a new node and we are initializing.
	finalizeCbs := make([]cluster.FinalizeInitCb, 0)
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		finalizeCb, err := e.Value.(cluster.ClusterListener).Init(self, clusterInfo)
		if err != nil {
			if self.Status != api.Status_STATUS_MAINTENANCE {
				self.Status = api.Status_STATUS_ERROR
			}
			logrus.Warnf("Failed to initialize Init %s: %v",
				e.Value.(cluster.ClusterListener).String(), err)
			c.cleanupInit(clusterInfo, self)
			return nil, err
		}
		if finalizeCb != nil {
			finalizeCbs = append(finalizeCbs, finalizeCb)
		}
	}

	return finalizeCbs, nil
}

// Alert all listeners that we are joining the cluster
func (c *ClusterManager) joinCluster(
	self *api.Node,
	exist bool,
) error {
	// Listeners may update initial state, so snap again.
	// The cluster db may have diverged since we waited for quorum
	// in between. Snapshot is created under cluster db lock to make
	// sure cluster db updates do not happen during snapshot, otherwise
	// there may be a mismatch between db updates from listeners and
	// cluster db state.
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.LockWithID(clusterLockKey, c.config.NodeId)
	if err != nil {
		logrus.Warnln("Unable to obtain cluster lock before creating snapshot: ",
			err)
		return err
	}
	initState, err := snapAndReadClusterInfo(c.snapshotPrefixes)
	kvdb.Unlock(kvlock)
	if err != nil {
		logrus.Panicf("Fatal, Unable to create snapshot: %v", err)
		return err
	}
	defer func() {
		if initState.Collector != nil {
			initState.Collector.Stop()
			initState.Collector = nil
		}
	}()

	// Alert all listeners that we are joining the cluster.
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		err := e.Value.(cluster.ClusterListener).Join(self, initState)
		if err != nil {
			if self.Status != api.Status_STATUS_MAINTENANCE {
				self.Status = api.Status_STATUS_ERROR
			}
			logrus.Warnf("Failed to initialize Join %s: %v",
				e.Value.(cluster.ClusterListener).String(), err)

			if exist == false {
				c.cleanupInit(initState.ClusterInfo, self)
			}
			logrus.Errorln("Failed to join cluster.", err)
			return err
		}
	}
	selfNodeEntry, ok := initState.ClusterInfo.NodeEntries[c.config.NodeId]
	if !ok {
		logrus.Panicln("Fatal, Unable to find self node entry in local cache")
	}

	_, _, err = c.updateNodeEntryDB(selfNodeEntry, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClusterManager) initClusterForListeners(
	self *api.Node,
) error {
	err := error(nil)

	// Alert all listeners that we are initializing a new cluster.
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		err = e.Value.(cluster.ClusterListener).ClusterInit(self)
		if err != nil {
			if self.Status != api.Status_STATUS_MAINTENANCE {
				self.Status = api.Status_STATUS_ERROR
			}
			logrus.Printf("Failed to initialize %s",
				e.Value.(cluster.ClusterListener).String())
			goto done
		}
	}
done:
	return err
}

func (c *ClusterManager) startClusterDBWatch(lastIndex uint64,
	kv kvdb.Kvdb) error {
	logrus.Infof("Cluster manager starting watch at version %d", lastIndex)
	go kv.WatchKey(ClusterDBKey, lastIndex, nil, c.watchDB)
	return nil
}

func (c *ClusterManager) startHeartBeat(
	clusterInfo *cluster.ClusterInfo,
	activeMap types.ClusterDomainsActiveMap,
) {
	gossipStoreKey := types.StoreKey(heartbeatKey + c.config.ClusterId)

	node := c.getCurrentState()
	c.putNodeCacheEntry(c.selfNode.Id, *node)
	c.gossip.UpdateSelf(gossipStoreKey, *node)
	var nodeIps []string

	gossipConfig := types.GossipStartConfiguration{
		ActiveMap: activeMap,
	}
	gossipConfig.Nodes = make(map[types.NodeId]types.GossipNodeConfiguration)

	for nodeId, nodeEntry := range clusterInfo.NodeEntries {
		if nodeId == node.Id {
			continue
		}
		labels := nodeEntry.NodeLabels
		version, ok := labels[gossipVersionKey]
		if !ok || version != c.gossipVersion {
			// Do not add nodes with mismatched version
			continue
		}
		nodeIp := nodeEntry.DataIp + ":" + c.gossipPort
		gossipConfig.Nodes[types.NodeId(nodeId)] = types.GossipNodeConfiguration{
			KnownUrl:      nodeIp,
			ClusterDomain: nodeEntry.ClusterDomain,
		}

		gossipPort := nodeEntry.GossipPort
		if gossipPort == "" {
			// The cluster DB does not have the gossip port value
			// The probability of this happening is close to 0
			// In an event if this happens, lets use our own gossip port
			// for this node. If that node has a different port, once that
			// node pings us, gossip protocol will automatically update the port
			gossipPort = c.gossipPort
		}
		nodeIps = append(nodeIps, nodeEntry.DataIp+":"+gossipPort)
	}
	if len(nodeIps) > 0 {
		logrus.Infof("Starting Gossip... Gossiping to these nodes : %v", nodeIps)
	} else {
		logrus.Infof("Starting Gossip...")
	}

	if len(activeMap) > 0 {
		gossipConfig.QuorumProviderType = types.QUORUM_PROVIDER_FAILURE_DOMAINS
	} else {
		gossipConfig.QuorumProviderType = types.QUORUM_PROVIDER_DEFAULT
	}

	c.gossip.Start(gossipConfig)
	c.gossip.UpdateCluster(c.getNonDecommisionedPeers(*clusterInfo))

	lastUpdateTs := time.Now()
	for {
		select {
		case <-stopHeartbeat:
			return
		default:
			node = c.getCurrentState()

			currTime := time.Now()
			diffTime := currTime.Sub(lastUpdateTs)
			if diffTime > 10*time.Second {
				logrus.Warnln("No gossip update for ", diffTime.Seconds(), "s")
			}
			c.gossip.UpdateSelf(gossipStoreKey, *node)
			lastUpdateTs = currTime
		}
		time.Sleep(2 * time.Second)
	}
}

func (c *ClusterManager) updateClusterStatus() {
	gossipStoreKey := types.StoreKey(heartbeatKey + c.config.ClusterId)
	for {
		node := c.getCurrentState()
		c.putNodeCacheEntry(node.Id, *node)

		// Process heartbeats from other nodes...
		gossipValues := c.gossip.GetStoreKeyValue(gossipStoreKey)

		numNodes := 0
		for id, gossipNodeInfo := range gossipValues {
			numNodes = numNodes + 1

			// Check to make sure we are not exceeding the size of the cluster.
			if c.size > 0 && numNodes > c.size {
				logrus.Fatalf("Fatal, number of nodes in the cluster has"+
					"exceeded the cluster size: %d > %d", numNodes, c.size)
				os.Exit(1)
			}

			// Special handling for self node
			if id == types.NodeId(node.Id) {
				// TODO: Implement State Machine for node statuses similar to the one in gossip
				if c.selfNode.Status == api.Status_STATUS_OK &&
					gossipNodeInfo.Status == types.NODE_STATUS_SUSPECT_NOT_IN_QUORUM {
					// Current:
					// Cluster Manager Status: UP.
					// Gossip Status: Suspecting Not in Quorum (stays in this state for quorumTimeout)
					// New:
					// Cluster Manager: Not in Quorum
					// Cluster Manager does not have a Suspect in Quorum status
					logrus.Warnf("Can't reach quorum no. of nodes. Suspecting out of quorum...")
					c.selfNode.Status = api.Status_STATUS_NOT_IN_QUORUM
					c.status = api.Status_STATUS_NOT_IN_QUORUM
				} else if (c.selfNode.Status == api.Status_STATUS_NOT_IN_QUORUM ||
					c.selfNode.Status == api.Status_STATUS_OK) &&
					(gossipNodeInfo.Status == types.NODE_STATUS_NOT_IN_QUORUM ||
						gossipNodeInfo.Status == types.NODE_STATUS_DOWN) {
					// Current:
					// Cluster Manager Status: UP or Not in Quorum.
					// Gossip Status: Not in Quorum or DOWN
					// New:
					// Cluster Manager: DOWN
					// Gossip waited for quorumTimeout and indicates we are Not in Quorum and should go Down
					logrus.Warnf("Not in quorum. Gracefully shutting down...")
					c.gossip.UpdateSelfStatus(types.NODE_STATUS_DOWN)
					c.selfNode.Status = api.Status_STATUS_OFFLINE
					c.status = api.Status_STATUS_NOT_IN_QUORUM
					c.Shutdown()
					os.Exit(1)
				} else if c.selfNode.Status == api.Status_STATUS_NOT_IN_QUORUM &&
					gossipNodeInfo.Status == types.NODE_STATUS_UP {
					// Current:
					// Cluster Manager Status: Not in Quorum
					// Gossip Status: Up
					// New:
					// Cluster Manager : UP
					c.selfNode.Status = api.Status_STATUS_OK
					c.status = api.Status_STATUS_OK
				} else {
					// Ignore the update
				}
				continue
			}

			// Notify node status change if required.
			peerNodeInCache := api.Node{}
			peerNodeInCache.Id = string(id)
			peerNodeInCache.Status = api.Status_STATUS_OK

			// Initialize a no-op notify listeners function
			notifyListenerFn := func() {}
			var peerNodeCopy *api.Node

			switch {
			case gossipNodeInfo.Status == types.NODE_STATUS_DOWN:
				// Replace the status of this node in cache to offline
				peerNodeInCache.Status = api.Status_STATUS_OFFLINE
				lastStatus, ok := c.nodeStatuses[string(id)]
				if !ok {
					// This node was probably added recently into gossip node
					// map through cluster database and is yet to reach out to us.
					// Mark this node down.
					logrus.Warnln("Detected new node with ", id,
						" to be offline due to inactivity.")

				} else {
					if lastStatus == peerNodeInCache.Status {
						break
					}
					logrus.Warnln("Detected node ", id,
						" to be offline due to inactivity.")
				}

				c.nodeStatuses[string(id)] = peerNodeInCache.Status
				peerNodeCopy = peerNodeInCache.Copy()
				notifyListenerFn = func() {
					for e := c.listeners.Front(); e != nil && c.gEnabled; e = e.Next() {
						err := e.Value.(cluster.ClusterListener).Update(peerNodeCopy)
						if err != nil {
							logrus.Warnln("Failed to notify ",
								e.Value.(cluster.ClusterListener).String())
						}
					}
				}

			case gossipNodeInfo.Status == types.NODE_STATUS_UP:
				peerNodeInCache.Status = api.Status_STATUS_OK
				lastStatus, ok := c.nodeStatuses[string(id)]
				if ok && lastStatus == peerNodeInCache.Status {
					break
				}
				c.nodeStatuses[string(id)] = peerNodeInCache.Status

				// A node discovered in the cluster.
				logrus.Infoln("Detected node", peerNodeInCache.Id,
					" to be in the cluster.")

				peerNodeCopy = peerNodeInCache.Copy()
				notifyListenerFn = func() {
					for e := c.listeners.Front(); e != nil && c.gEnabled; e = e.Next() {
						err := e.Value.(cluster.ClusterListener).Add(peerNodeCopy)
						if err != nil {
							logrus.Warnln("Failed to notify ",
								e.Value.(cluster.ClusterListener).String())
						}
					}
				}
			}

			// Update cache with gossip data
			if gossipNodeInfo.Value != nil {
				peerNodeInGossip, ok := gossipNodeInfo.Value.(api.Node)
				if ok {
					if peerNodeInCache.Status == api.Status_STATUS_OFFLINE {
						// Overwrite the status of Node in Gossip data with Down
						peerNodeInGossip.Status = peerNodeInCache.Status
					} else {
						if peerNodeInGossip.Status == api.Status_STATUS_MAINTENANCE {
							// If the node sent its status as Maintenance
							// do not overwrite it with online
						} else {
							peerNodeInGossip.Status = peerNodeInCache.Status
						}
					}
					c.putNodeCacheEntry(peerNodeInGossip.Id, peerNodeInGossip)
				} else {
					logrus.Errorln("Unable to get node info from gossip")
					c.putNodeCacheEntry(peerNodeInCache.Id, peerNodeInCache)
				}
			} else {
				c.putNodeCacheEntry(peerNodeInCache.Id, peerNodeInCache)
			}

			// Notify the listeners
			notifyListenerFn()
		}
		time.Sleep(2 * time.Second)
	}
}

// DisableUpdates disables gossip updates
func (c *ClusterManager) DisableUpdates() error {
	logrus.Warnln("Disabling gossip updates")
	c.gEnabled = false

	return nil
}

// EnableUpdates enables gossip updates
func (c *ClusterManager) EnableUpdates() error {
	logrus.Warnln("Enabling gossip updates")
	c.gEnabled = true

	return nil
}

// GetGossipState returns current gossip state
func (c *ClusterManager) GetGossipState() *cluster.ClusterState {
	gossipStoreKey := types.StoreKey(heartbeatKey + c.config.ClusterId)
	nodeValue := c.gossip.GetStoreKeyValue(gossipStoreKey)
	nodes := make([]types.NodeValue, len(nodeValue), len(nodeValue))
	i := 0
	for _, value := range nodeValue {
		nodes[i] = value
		i++
	}

	return &cluster.ClusterState{NodeStatus: nodes}
}

func (c *ClusterManager) waitForQuorum(exist bool) error {
	// Max quorum retries allowed = 600
	// 600 * 2 seconds (gossip interval) = 20 minutes before it restarts
	quorumRetries := 0
	for {
		gossipSelfStatus := c.gossip.GetSelfStatus()
		if c.selfNode.Status == api.Status_STATUS_NOT_IN_QUORUM &&
			gossipSelfStatus == types.NODE_STATUS_UP {
			// Node not initialized yet
			// Achieved quorum in the cluster.
			// Lets start the node
			c.selfNode.Status = api.Status_STATUS_INIT
			err := c.joinCluster(&c.selfNode, exist)
			if err != nil {
				if c.selfNode.Status != api.Status_STATUS_MAINTENANCE {
					c.selfNode.Status = api.Status_STATUS_ERROR
				}
				return err
			}
			break
		} else {
			c.status = api.Status_STATUS_NOT_IN_QUORUM
			if quorumRetries == 600 {
				err := fmt.Errorf("Unable to achieve Quorum." +
					" Timeout 20 minutes exceeded.")
				logrus.Warnln("Failed to join cluster: ", err)
				c.status = api.Status_STATUS_NOT_IN_QUORUM
				c.selfNode.Status = api.Status_STATUS_OFFLINE
				c.gossip.UpdateSelfStatus(types.NODE_STATUS_DOWN)
				return err
			}
			if quorumRetries == 0 {
				logrus.Infof("Waiting for the cluster to reach quorum...")
			}
			time.Sleep(types.DEFAULT_GOSSIP_INTERVAL)
			quorumRetries++
		}
	}
	// Update the listeners that we have joined the cluster and
	// and our quorum status
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		err := e.Value.(cluster.ClusterListener).JoinComplete(&c.selfNode)
		if err != nil {
			logrus.Warnln("Failed to notify ", e.Value.(cluster.ClusterListener).String())
		}
	}

	// Update the status after the listeners are started to ensure all REST points are available.
	c.status = api.Status_STATUS_OK
	c.selfNode.Status = api.Status_STATUS_OK

	return nil
}

func (c *ClusterManager) initializeCluster(db kvdb.Kvdb, selfClusterDomain string) (
	*cluster.ClusterInfo,
	error,
) {
	kvlock, err := db.LockWithID(clusterLockKey, c.config.NodeId)
	if err != nil {
		logrus.Panicln("Fatal, Unable to obtain cluster lock.", err)
	}
	defer db.Unlock(kvlock)

	clusterInfo, _, err := readClusterInfo()
	if err != nil {
		logrus.Panicln(err)
	}

	selfNodeEntry, ok := clusterInfo.NodeEntries[c.config.NodeId]
	if ok && selfNodeEntry.Status == api.Status_STATUS_DECOMMISSION {
		msg := fmt.Sprintf("Node is in decommision state, Node ID %s.",
			c.selfNode.Id)
		logrus.Errorln(msg)
		return nil, cluster.ErrNodeDecommissioned
	}
	// Set the clusterID in db
	clusterInfo.Id = c.config.ClusterId

	if clusterInfo.Status == api.Status_STATUS_INIT {
		logrus.Infoln("Initializing a new cluster.")
		// Initialize self node
		clusterInfo.Status = api.Status_STATUS_OK

		err = c.initClusterForListeners(&c.selfNode)
		if err != nil {
			logrus.Errorln("Failed to initialize the cluster.", err)
			return nil, err
		}
		// While we hold the lock write the cluster info
		// to kvdb.
		_, err := writeClusterInfo(&clusterInfo)
		if err != nil {
			logrus.Errorln("Failed to initialize the cluster.", err)
			return nil, err
		}
	} else if clusterInfo.Status&api.Status_STATUS_OK > 0 {
		logrus.Infoln("Cluster state is OK... Joining the cluster.")
	} else {
		return nil, errors.New("Fatal, Cluster is in an unexpected state.")
	}
	// Cluster database max size... 0 if unlimited.
	c.size = clusterInfo.Size
	c.status = api.Status_STATUS_OK
	return &clusterInfo, nil
}

func (c *ClusterManager) quorumMember() bool {
	if c.listeners.Len() == 0 {
		// If there are no listeners registered by the driver, assume
		// this node is a quorum member, so this becomes the default behavior
		// for drivers which do not implement the ClusterListener interface.
		return true
	}
	quorumRequired := false
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		quorumRequired = quorumRequired ||
			e.Value.(cluster.ClusterListener).QuorumMember(&c.selfNode)
	}
	return quorumRequired
}

func (c *ClusterManager) initListeners(
	db kvdb.Kvdb,
	clusterMaxSize int,
	nodeExists *bool,
	nodeInitialized bool,
	selfClusterDomain string,
) (uint64, *cluster.ClusterInfo, error) {
	// Initialize the cluster if required
	clusterInfo, err := c.initializeCluster(db, selfClusterDomain)
	if err != nil {
		return 0, nil, err
	}

	// Initialize the node in cluster
	self, exist := c.initNode(clusterInfo)
	*nodeExists = exist
	finalizeCbs, err := c.initNodeInCluster(
		clusterInfo,
		self,
		*nodeExists,
		nodeInitialized,
	)
	if err != nil {
		logrus.Errorln("Failed to initialize node in cluster.", err)
		return 0, nil, err
	}

	selfNodeEntry, ok := clusterInfo.NodeEntries[c.config.NodeId]
	if !ok {
		logrus.Panicln("Fatal, Unable to find self node entry in local cache")
	}

	// the inverse value is to handle upgrades.
	// This node does not participate in quorum decisions if it is
	// decommissioned or if none of the listeners require it.
	selfNodeEntry.NonQuorumMember =
		selfNodeEntry.Status == api.Status_STATUS_DECOMMISSION ||
			!c.quorumMember()
	if !selfNodeEntry.NonQuorumMember {
		logrus.Infof("This node participates in quorum decisions")
	} else {
		logrus.Infof("This node does not participates in quorum decisions")
	}

	initFunc := func(clusterInfo cluster.ClusterInfo) error {
		numNodes := 0
		for _, node := range clusterInfo.NodeEntries {
			if node.Status != api.Status_STATUS_DECOMMISSION {
				numNodes++
			}
		}
		if clusterMaxSize > 0 && numNodes > clusterMaxSize {
			return fmt.Errorf("Cluster is operating at maximum capacity "+
				"(%v nodes). Please remove a node before attempting to "+
				"add a new node.", clusterMaxSize)
		}

		// Finalize inits from subsystems under cluster db lock.
		for _, finalizeCb := range finalizeCbs {
			if err := finalizeCb(); err != nil {
				logrus.Errorf("Failed finalizing init: %s", err.Error())
				return err
			}
		}
		return nil
	}

	kvp, kvClusterInfo, err := c.updateNodeEntryDB(selfNodeEntry,
		initFunc)
	if err != nil {
		logrus.Errorln("Failed to save the database.", err)
		return 0, nil, err
	}
	if kvClusterInfo.Status == api.Status_STATUS_INIT {
		logrus.Panicln("Cluster in an unexpected state: ", kvClusterInfo.Status)
	}

	// update node cache with entries in the database at this point since
	// we are going to start watch at kvp.ModifiedIndex
	c.nodeCacheLock.Lock()
	defer c.nodeCacheLock.Unlock()
	for _, node := range c.nodes(kvClusterInfo) {
		c.nodeCache[node.Id] = node
	}
	return kvp.ModifiedIndex, kvClusterInfo, nil
}

func (c *ClusterManager) initializeAndStartHeartbeat(
	kvdb kvdb.Kvdb,
	clusterMaxSize int,
	exist *bool,
	nodeInitialized bool,
	selfClusterDomain string,
) (uint64, *cluster.ClusterInfo, error) {
	lastIndex, clusterInfo, err := c.initListeners(
		kvdb,
		clusterMaxSize,
		exist,
		nodeInitialized,
		selfClusterDomain,
	)
	if err != nil {
		return 0, nil, err
	}

	// Set the status to NOT_IN_QUORUM to start the node.
	// Once we achieve quorum then we actually join the cluster
	// and change the status to OK
	c.selfNode.Status = api.Status_STATUS_NOT_IN_QUORUM

	// Get the cluster domain info
	clusterDomainInfos, err := c.clusterDomainManager.EnumerateDomains()
	if err != nil && err != clusterdomain.ErrNotImplemented {
		return 0, nil, err
	}

	// Start heartbeating to other nodes.
	go c.startHeartBeat(
		clusterInfo,
		clusterdomain.GetActiveMapFromClusterDomainInfos(clusterDomainInfos),
	)
	return lastIndex, clusterInfo, nil
}

func (c *ClusterManager) setupManagers(config *cluster.ClusterServerConfiguration) {
	if config.ConfigSchedManager == nil {
		c.schedManager = sched.NewDefaultSchedulePolicy()
	} else {
		c.schedManager = config.ConfigSchedManager
	}

	if config.ConfigObjectStoreManager == nil {
		c.objstoreManager = objectstore.NewDefaultObjectStore()
	} else {
		c.objstoreManager = config.ConfigObjectStoreManager
	}

	if config.ConfigSecretManager == nil {
		c.secretsManager = secrets.NewDefaultSecrets()
	} else {
		c.secretsManager = config.ConfigSecretManager
	}

	if config.ConfigSystemTokenManager == nil {
		c.systemTokenManager = auth.NoAuth()
	} else {
		c.systemTokenManager = config.ConfigSystemTokenManager
	}

	if config.ConfigClusterDomainProvider == nil {
		c.clusterDomainManager = clusterdomain.NewDefaultClusterDomainPorvider()
	} else {
		c.clusterDomainManager = config.ConfigClusterDomainProvider
	}
}

// Start initiates the cluster manager and the cluster state machine
func (c *ClusterManager) Start(
	clusterMaxSize int,
	nodeInitialized bool,
	gossipPort string,
	selfClusterDomain string,
) error {
	return c.StartWithConfiguration(
		clusterMaxSize,
		nodeInitialized,
		gossipPort,
		[]string{ClusterDBKey},
		selfClusterDomain,
		&cluster.ClusterServerConfiguration{})
}

func (c *ClusterManager) StartWithConfiguration(
	clusterMaxSize int,
	nodeInitialized bool,
	gossipPort string,
	snapshotPrefixes []string,
	selfClusterDomain string,
	config *cluster.ClusterServerConfiguration,
) error {
	var err error

	logrus.Infoln("Cluster manager starting...")

	snapshotPrefixes = append(snapshotPrefixes, ClusterDBKey)
	c.snapshotPrefixes = snapshotPrefixes

	kv := kvdb.Instance()

	// osdconfig manager should be instantiated as soon as kv is ready
	logrus.Info("initializing osdconfig manager")
	c.configManager, err = osdconfig.NewCaller(kv)
	if err != nil {
		return err
	}

	// Setup any default managers if none were provided
	c.setupManagers(config)

	c.gEnabled = true
	c.selfNode = api.Node{}
	c.selfNode.GenNumber = uint64(time.Now().UnixNano())
	c.selfNode.Id = c.config.NodeId
	c.selfNode.SchedulerNodeName = c.config.SchedulerNodeName
	c.selfNode.Status = api.Status_STATUS_INIT
	c.selfNode.MgmtIp, c.selfNode.DataIp, err = ExternalIp(&c.config)
	c.selfNode.StartTime = time.Now()
	c.selfNode.Hostname, _ = os.Hostname()
	c.gossipPort = gossipPort
	c.selfNode.GossipPort = gossipPort
	c.selfClusterDomain = selfClusterDomain
	if err != nil {
		logrus.Errorf("Failed to get external IP address for mgt/data interfaces: %s.",
			err)
		return err
	}

	c.selfNode.NodeData = make(map[string]interface{})
	c.system = systemutils.New()

	// Start the gossip protocol.
	// XXX Make the port configurable.
	gob.Register(api.Node{})
	gossipIntervals := types.GossipIntervals{
		GossipInterval:   types.DEFAULT_GOSSIP_INTERVAL,
		PushPullInterval: types.DEFAULT_PUSH_PULL_INTERVAL,
		ProbeInterval:    types.DEFAULT_PROBE_INTERVAL,
		ProbeTimeout:     types.DEFAULT_PROBE_TIMEOUT,
		QuorumTimeout:    types.DEFAULT_QUORUM_TIMEOUT,
	}
	c.gossip = gossip.New(
		c.selfNode.DataIp+":"+c.gossipPort,
		types.NodeId(c.config.NodeId),
		c.selfNode.GenNumber,
		gossipIntervals,
		types.GOSSIP_VERSION_2,
		c.config.ClusterId,
		selfClusterDomain,
	)
	c.gossipVersion = types.GOSSIP_VERSION_2

	var exist bool
	lastIndex, clusterInfo, err := c.initializeAndStartHeartbeat(
		kv,
		clusterMaxSize,
		&exist,
		nodeInitialized,
		selfClusterDomain,
	)
	if err != nil {
		return err
	}

	// Update all the listeners with the new db
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		err := e.Value.(cluster.ClusterListener).UpdateCluster(&c.selfNode, clusterInfo)
		if err != nil {
			logrus.Warnln("Failed to notify ", e.Value.(cluster.ClusterListener).String())
		}
	}

	_ = c.startClusterDBWatch(lastIndex, kv)

	err = c.waitForQuorum(exist)
	if err != nil {
		return err
	}

	go c.updateClusterStatus()
	go c.replayNodeDecommission()

	return nil
}

// NodeStatus returns the status of a node. It compares the status maintained by the
// cluster manager and the provided listener and returns the appropriate one
func (c *ClusterManager) NodeStatus() (api.Status, error) {
	clusterNodeStatus := c.selfNode.Status

	if clusterNodeStatus != api.Status_STATUS_OK {
		// Status of this node as seen by Cluster Manager is not OK
		// This takes highest precedence over other listener statuses.
		// Returning our status
		return clusterNodeStatus, nil
	}

	for e := c.listeners.Front(); e != nil; e = e.Next() {
		listenerStatus := e.Value.(cluster.ClusterListener).ListenerStatus()
		if listenerStatus == api.Status_STATUS_NONE {
			continue
		}
		if int(listenerStatus.StatusKind()) >= int(clusterNodeStatus.StatusKind()) {
			clusterNodeStatus = listenerStatus
		}
	}

	return clusterNodeStatus, nil
}

// PeerStatus returns the status of a peer node as seen by us
func (c *ClusterManager) PeerStatus(listenerName string) (map[string]api.Status, error) {
	statusMap := make(map[string]api.Status)
	var listenerStatusMap map[string]api.Status
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		if e.Value.(cluster.ClusterListener).String() == listenerName {
			listenerStatusMap = e.Value.(cluster.ClusterListener).ListenerPeerStatus()
			break
		}
	}
	c.nodeCacheLock.Lock()
	defer c.nodeCacheLock.Unlock()
	// Listener failed to provide peer status
	if listenerStatusMap == nil || len(listenerStatusMap) == 0 {
		for _, n := range c.nodeCache {
			if n.Id == c.selfNode.Id {
				// skip self
				continue
			}
			statusMap[n.Id] = n.Status
		}
		return statusMap, nil
	}
	// Compare listener's peer statuses and cluster provider's peer statuses
	for _, n := range c.nodeCache {
		if n.Id == c.selfNode.Id {
			// Skip self
			continue
		}
		clusterNodeStatus := n.Status
		listenerNodeStatus, ok := listenerStatusMap[n.Id]
		if !ok {
			// Could not find listener's peer status
			// Using cluster provider's peer status
			statusMap[n.Id] = clusterNodeStatus
		}
		if int(listenerNodeStatus.StatusKind()) >= int(clusterNodeStatus.StatusKind()) {
			// Use listener's peer status
			statusMap[n.Id] = listenerNodeStatus
		} else {
			// Use the cluster provider's peer status
			statusMap[n.Id] = clusterNodeStatus
		}
	}
	return statusMap, nil
}

func (c *ClusterManager) nodes(clusterDB *cluster.ClusterInfo) []api.Node {
	nodes := []api.Node{}
	for _, n := range clusterDB.NodeEntries {
		node := api.Node{}
		if n.Id == c.selfNode.Id {
			node = *c.getCurrentState()
		} else {
			node.Id = n.Id
			node.SchedulerNodeName = n.SchedulerNodeName
			node.Status = n.Status
			node.MgmtIp = n.MgmtIp
			node.DataIp = n.DataIp
			node.Hostname = n.Hostname
			node.NodeLabels = n.NodeLabels
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func (c *ClusterManager) enumerateFromClusterDB() []api.Node {
	clusterDB, _, err := readClusterInfo()
	if err != nil {
		logrus.Errorf("enumerateNodesFromClusterDB failed with error: %v", err)
		return make([]api.Node, 0)
	}
	return c.nodes(&clusterDB)
}

func (c *ClusterManager) enumerateFromCache() []api.Node {
	var clusterDB cluster.ClusterInfo
	c.nodeCacheLock.Lock()
	defer c.nodeCacheLock.Unlock()
	nodes := make([]api.Node, len(c.nodeCache))
	i := 0
	for _, n := range c.nodeCache {
		n, _ := c.getNodeEntry(n.Id, &clusterDB)
		nodes[i] = *n.Copy()
		i++
	}
	return nodes
}

// Enumerate lists all the nodes in the cluster.
func (c *ClusterManager) Enumerate() (api.Cluster, error) {
	clusterState := api.Cluster{
		Id:     c.config.ClusterId,
		Status: c.status,
		NodeId: c.selfNode.Id,
	}

	if c.selfNode.Status == api.Status_STATUS_NOT_IN_QUORUM ||
		c.selfNode.Status == api.Status_STATUS_MAINTENANCE {
		// If the node is not yet ready, query the cluster db
		// for node members since gossip is not ready yet.
		clusterState.Nodes = c.enumerateFromClusterDB()
	} else {
		clusterState.Nodes = c.enumerateFromCache()
	}

	// Allow listeners to add/modify data
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		if err := e.Value.(cluster.ClusterListener).Enumerate(clusterState); err != nil {
			logrus.Warnf("listener %s enumerate failed: %v",
				e.Value.(cluster.ClusterListener).String(), err)
			continue
		}
	}
	return clusterState, nil
}

func (c *ClusterManager) updateNodeEntryDB(
	nodeEntry cluster.NodeEntry,
	checkCbBeforeUpdate checkFunc,
) (*kvdb.KVPair, *cluster.ClusterInfo, error) {
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.LockWithID(clusterLockKey, c.config.NodeId)
	if err != nil {
		logrus.Warnln("Unable to obtain cluster lock for updating cluster DB.",
			err)
		return nil, nil, err
	}
	defer kvdb.Unlock(kvlock)

	currentState, _, err := readClusterInfo()
	if err != nil {
		return nil, nil, err
	}

	currentState.NodeEntries[nodeEntry.Id] = nodeEntry

	if checkCbBeforeUpdate != nil {
		err = checkCbBeforeUpdate(currentState)
		if err != nil {
			return nil, nil, err
		}
	}

	kvp, err := writeClusterInfo(&currentState)
	if err != nil {
		logrus.Errorln("Failed to save the database.", err)
	}
	return kvp, &currentState, err
}

// SetSize sets the maximum number of nodes in a cluster.
func (c *ClusterManager) SetSize(size int) error {
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.LockWithID(clusterLockKey, c.config.NodeId)
	if err != nil {
		logrus.Warnln("Unable to obtain cluster lock for updating config", err)
		return nil
	}
	defer kvdb.Unlock(kvlock)

	db, _, err := readClusterInfo()
	if err != nil {
		return err
	}

	db.Size = size

	_, err = writeClusterInfo(&db)

	return err
}

func (c *ClusterManager) getNodeInfoFromClusterDb(id string) (api.Node, error) {
	node := api.Node{Id: id}
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.LockWithID(clusterLockKey, c.config.NodeId)
	if err != nil {
		logrus.Warnln("Unable to obtain cluster lock for marking "+
			"node decommission", err)
		return node, err
	}
	defer kvdb.Unlock(kvlock)

	db, _, err := readClusterInfo()
	if err != nil {
		return node, err
	}

	nodeEntry, ok := db.NodeEntries[id]
	if !ok {
		msg := fmt.Sprintf("Node entry does not exist, Node ID %s", id)
		return node, errors.New(msg)
	}
	node.Status = nodeEntry.Status
	return node, nil
}

func (c *ClusterManager) markNodeDecommission(node api.Node) error {
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.LockWithID(clusterLockKey, c.config.NodeId)
	if err != nil {
		logrus.Warnln("Unable to obtain cluster lock for marking "+
			"node decommission",
			err)
		return err
	}
	defer kvdb.Unlock(kvlock)

	db, _, err := readClusterInfo()
	if err != nil {
		return err
	}

	nodeEntry, ok := db.NodeEntries[node.Id]
	if !ok {
		msg := fmt.Sprintf("Node entry does not exist, Node ID %s",
			node.Id)
		return errors.New(msg)
	}

	nodeEntry.Status = api.Status_STATUS_DECOMMISSION
	db.NodeEntries[node.Id] = nodeEntry

	if c.selfNode.Id == node.Id {
		c.selfNode.Status = api.Status_STATUS_DECOMMISSION
	}
	_, err = writeClusterInfo(&db)

	return err
}

func (c *ClusterManager) deleteNodeFromDB(nodeID string) error {
	// Delete node from cluster DB
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.LockWithID(clusterLockKey, c.config.NodeId)
	if err != nil {
		logrus.Panicln("fatal, unable to obtain cluster lock. ", err)
	}
	defer kvdb.Unlock(kvlock)

	currentState, _, err := readClusterInfo()
	if err != nil {
		logrus.Errorln("Failed to read cluster info. ", err)
		return err
	}

	delete(currentState.NodeEntries, nodeID)

	_, err = writeClusterInfo(&currentState)
	if err != nil {
		logrus.Errorln("Failed to save the database.", err)
	}
	return err
}

// Remove node(s) from the cluster permanently.
func (c *ClusterManager) Remove(nodes []api.Node, forceRemove bool) error {
	logrus.Infof("ClusterManager Remove node.")

	var resultErr error

	inQuorum := !(c.selfNode.Status == api.Status_STATUS_NOT_IN_QUORUM)

	for i, _ := range nodes {

		if id, cerr := c.GetNodeIdFromIp(nodes[i].Id); cerr == nil {
			if nodes[i].Id != id {
				nodes[i].Id = id
			}
		}

		node, exist := c.getNodeCacheEntry(nodes[i].Id)
		if !exist {
			node, resultErr = c.getNodeInfoFromClusterDb(nodes[i].Id)
			if resultErr != nil {
				logrus.Errorf("Error getting node info for id %s : %v", nodes[i].Id,
					resultErr)
				return fmt.Errorf("Node %s does not exist", nodes[i].Id)
			}
		}

		// If removing node is self and node is not in maintenance mode,
		// disallow node remove.
		if nodes[i].Id == c.selfNode.Id &&
			c.selfNode.Status != api.Status_STATUS_MAINTENANCE {
			msg := fmt.Sprintf(decommissionErrMsg, nodes[i].Id)
			logrus.Errorf(msg)
			return errors.New(msg)
		} else if nodes[i].Id != c.selfNode.Id && inQuorum {
			nodeCacheStatus := node.Status
			// If node is not down, do not remove it
			if nodeCacheStatus != api.Status_STATUS_OFFLINE &&
				nodeCacheStatus != api.Status_STATUS_MAINTENANCE &&
				nodeCacheStatus != api.Status_STATUS_DECOMMISSION {

				msg := fmt.Sprintf(decommissionErrMsg, nodes[i].Id)
				logrus.Errorf(msg+", node status: %s", nodeCacheStatus)
				return errors.New(msg)
			}
		}

		if forceRemove {
			// Mark the other node down so that it can be decommissioned.
			for e := c.listeners.Front(); e != nil; e = e.Next() {
				logrus.Infof("Remove node: ask cluster listener %s "+
					"to mark node %s down ",
					e.Value.(cluster.ClusterListener).String(), nodes[i].Id)
				err := e.Value.(cluster.ClusterListener).MarkNodeDown(&nodes[i])
				if err != nil {
					logrus.Warnf("Node mark down error: %v", err)
					return err
				}
			}
		}

		// Ask listeners, can we remove this node?
		for e := c.listeners.Front(); e != nil; e = e.Next() {
			logrus.Infof("Remove node: ask cluster listener: "+
				"can we remove node ID %s, %s",
				nodes[i].Id, e.Value.(cluster.ClusterListener).String())
			additionalMsg, err := e.Value.(cluster.ClusterListener).CanNodeRemove(&nodes[i])
			if err != nil && !(err == cluster.ErrRemoveCausesDataLoss && forceRemove) {
				msg := fmt.Sprintf("Cannot remove node ID %s: %s.", nodes[i].Id, err)
				if additionalMsg != "" {
					msg = msg + " " + additionalMsg
				}
				logrus.Warnf(msg)
				return errors.New(msg)
			}
		}

		err := c.markNodeDecommission(nodes[i])
		if err != nil {
			msg := fmt.Sprintf("Failed to mark node as "+
				"decommision, error %s",
				err)
			logrus.Errorf(msg)
			return errors.New(msg)
		}

		if !inQuorum {
			// If we are not in quorum, we only mark the node as decommissioned
			// since this node is not functional yet.
			continue
		}

		// Alert all listeners that we are removing this node.
		for e := c.listeners.Front(); e != nil; e = e.Next() {
			logrus.Infof("Remove node: notify cluster listener: %s",
				e.Value.(cluster.ClusterListener).String())
			err := e.Value.(cluster.ClusterListener).Remove(&nodes[i], forceRemove)
			if err != nil {
				if err != cluster.ErrNodeRemovePending {
					logrus.Warnf("Cluster listener failed to "+
						"remove node: %s: %s",
						e.Value.(cluster.ClusterListener).String(),
						err)
					return err
				} else {
					resultErr = err
				}
			}
		}
	}

	return resultErr
}

// NodeRemoveDone is called from the listeners when their job of Node removal is done.
func (c *ClusterManager) NodeRemoveDone(nodeID string, result error) {
	// XXX: only storage will make callback right now
	if result != nil {
		msg := fmt.Sprintf("Storage failed to decommission node %s, "+
			"error %s",
			nodeID,
			result)
		logrus.Errorf(msg)
		return
	}

	logrus.Infof("Cluster manager node remove done: node ID %s", nodeID)

	// Remove osdconfig data from etcd
	if err := c.configManager.DeleteNodeConf(nodeID); err != nil {
		logrus.Warn("error removing node from osdconfig:", err)
	}

	if err := c.deleteNodeFromDB(nodeID); err != nil {
		msg := fmt.Sprintf("Failed to delete node %s "+
			"from cluster database, error %s",
			nodeID, err)
		logrus.Errorf(msg)
	}
}

func (c *ClusterManager) replayNodeDecommission() {
	currentState, _, err := readClusterInfo()
	if err != nil {
		logrus.Infof("Failed to read cluster db for node decommissions: %v", err)
		return
	}

	for _, nodeEntry := range currentState.NodeEntries {
		if nodeEntry.Status == api.Status_STATUS_DECOMMISSION {
			logrus.Infof("Replay Node Remove for node ID %s", nodeEntry.Id)

			var n api.Node
			n.Id = nodeEntry.Id
			nodes := make([]api.Node, 0)
			nodes = append(nodes, n)
			err := c.Remove(nodes, false)
			if err != nil {
				logrus.Warnf("Failed to replay node remove: "+
					"node ID %s, error %s",
					nodeEntry.Id, err)
			}
		}
	}
}

// Shutdown can be called when THIS node is gracefully shutting down.
func (c *ClusterManager) Shutdown() error {
	db, _, err := readClusterInfo()
	if err != nil {
		logrus.Warnf("Could not read cluster database (%v).", err)
		return err
	}

	// Alert all listeners that we are shutting this node down.
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		logrus.Infof("Shutting down %s", e.Value.(cluster.ClusterListener).String())
		if err := e.Value.(cluster.ClusterListener).Halt(&c.selfNode, &db); err != nil {
			logrus.Warnf("Failed to shutdown %s",
				e.Value.(cluster.ClusterListener).String())
		}
	}
	return nil
}

func (c *ClusterManager) ClusterNotifyNodeDown(culpritNodeId string) (string, error) {
	killNodeId := c.gossip.ExternalNodeLeave(types.NodeId(culpritNodeId))
	return string(killNodeId), nil
}

func (c *ClusterManager) ClusterNotifyClusterDomainsUpdate(activeMap types.ClusterDomainsActiveMap) error {
	if c.gossip != nil {
		return c.gossip.UpdateClusterDomainsActiveMap(activeMap)
	}
	return nil
}

func (c *ClusterManager) EnumerateAlerts(ts, te time.Time, resource api.ResourceType) (*api.Alerts, error) {
	a := api.Alerts{}

	for e := c.listeners.Front(); e != nil; e = e.Next() {
		listenerAlerts, err := e.Value.(cluster.ClusterListener).EnumerateAlerts(ts, te, resource)
		if err != nil {
			continue
		}
		if listenerAlerts != nil {
			a.Alert = append(a.Alert, listenerAlerts.Alert...)
		}
	}
	return &a, nil
}

func (c *ClusterManager) EraseAlert(resource api.ResourceType, alertID int64) error {
	erased := false
	for e := c.listeners.Front(); e != nil; e = e.Next() {
		if err := e.Value.(cluster.ClusterListener).EraseAlert(resource, alertID); err != nil {
			continue
		}
		erased = true
	}
	if !erased {
		return fmt.Errorf("Unable to erase alert (%v)", alertID)
	}
	return nil
}

func (c *ClusterManager) getNodeCacheEntry(nodeId string) (api.Node, bool) {
	c.nodeCacheLock.Lock()
	defer c.nodeCacheLock.Unlock()
	n, ok := c.nodeCache[nodeId]
	return n, ok
}

func (c *ClusterManager) putNodeCacheEntry(nodeId string, node api.Node) {
	c.nodeCacheLock.Lock()
	defer c.nodeCacheLock.Unlock()
	c.nodeCache[nodeId] = node
}

// GetSelfDomain returns the cluster domain for this node
func (c *ClusterManager) GetSelfDomain() (*clusterdomain.ClusterDomainInfo, error) {
	return c.clusterDomainManager.GetSelfDomain()
}

// EnumerateDomains returns all the cluster domains in the cluster
func (c *ClusterManager) EnumerateDomains() ([]*clusterdomain.ClusterDomainInfo, error) {
	return c.clusterDomainManager.EnumerateDomains()
}

// InspectDomain returns the cluster domain info for the provided argument.
func (c *ClusterManager) InspectDomain(name string) (*clusterdomain.ClusterDomainInfo, error) {
	return c.clusterDomainManager.InspectDomain(name)
}

// DeleteDomain deletes a cluster domain entry
func (c *ClusterManager) DeleteDomain(name string) error {
	return c.clusterDomainManager.DeleteDomain(name)
}

// UpdateDomainState sets the cluster domain info object into kvdb
func (c *ClusterManager) UpdateDomainState(name string, state types.ClusterDomainState) error {
	return c.clusterDomainManager.UpdateDomainState(name, state)
}

// osdconfig.ConfigCaller compliance
func (c *ClusterManager) GetClusterConf() (*osdconfig.ClusterConfig, error) {
	return c.configManager.GetClusterConf()
}

func (c *ClusterManager) GetNodeConf(nodeID string) (*osdconfig.NodeConfig, error) {
	return c.configManager.GetNodeConf(nodeID)
}

func (c *ClusterManager) SetClusterConf(config *osdconfig.ClusterConfig) error {
	return c.configManager.SetClusterConf(config)
}

func (c *ClusterManager) SetNodeConf(config *osdconfig.NodeConfig) error {
	return c.configManager.SetNodeConf(config)
}

func (c *ClusterManager) DeleteNodeConf(nodeID string) error {
	return c.configManager.DeleteNodeConf(nodeID)
}

func (c *ClusterManager) EnumerateNodeConf() (*osdconfig.NodesConfig, error) {
	return c.configManager.EnumerateNodeConf()
}

// SchedPolicyCreate creates a policy with given name and schedule.
func (c *ClusterManager) SchedPolicyCreate(name, sched string) error {
	return c.schedManager.SchedPolicyCreate(name, sched)
}

// SchedPolicyUpdate updates a policy with given name and schedule.
func (c *ClusterManager) SchedPolicyUpdate(name, sched string) error {
	return c.schedManager.SchedPolicyUpdate(name, sched)
}

// SchedPolicyDelete deletes a policy with given name.
func (c *ClusterManager) SchedPolicyDelete(name string) error {
	return c.schedManager.SchedPolicyDelete(name)
}

// SchedPolicyEnumerate enumerates all configured policies or the ones specified.
func (c *ClusterManager) SchedPolicyEnumerate() ([]*sched.SchedPolicy, error) {
	return c.schedManager.SchedPolicyEnumerate()
}

// SchedPolicyGet returns schedule policy matching given name.
func (c *ClusterManager) SchedPolicyGet(name string) (*sched.SchedPolicy, error) {
	return c.schedManager.SchedPolicyGet(name)
}

// ObjectStoreInspect returns status of objectstore
func (c *ClusterManager) ObjectStoreInspect(objectstoreID string) (*api.ObjectstoreInfo, error) {
	return c.objstoreManager.ObjectStoreInspect(objectstoreID)
}

// ObjectStoreCreate objectstore on specified volume
func (c *ClusterManager) ObjectStoreCreate(volumeID string) (*api.ObjectstoreInfo, error) {
	return c.objstoreManager.ObjectStoreCreate(volumeID)
}

// ObjectStoreUpdate enable/disable objectstore
func (c *ClusterManager) ObjectStoreUpdate(objectstoreID string, enable bool) error {
	return c.objstoreManager.ObjectStoreUpdate(objectstoreID, enable)
}

// ObjectStoreDelete objectstore from cluster
func (c *ClusterManager) ObjectStoreDelete(objectstoreID string) error {
	return c.objstoreManager.ObjectStoreDelete(objectstoreID)
}

// SecretLogin create session with secret store
func (c *ClusterManager) SecretLogin(secretType string, secretConfig map[string]string) error {
	return c.secretsManager.SecretLogin(secretType, secretConfig)
}

// SecretSetDefaultSecretKey  sets the cluster wide secret key
func (c *ClusterManager) SecretSetDefaultSecretKey(secretKey string, override bool) error {
	return c.secretsManager.SecretSetDefaultSecretKey(secretKey, override)
}

// SecretGetDefaultSecretKey returns cluster wide secret key
func (c *ClusterManager) SecretGetDefaultSecretKey() (interface{}, error) {
	return c.secretsManager.SecretGetDefaultSecretKey()
}

// SecretCheckLogin validates session with secret store
func (c *ClusterManager) SecretCheckLogin() error {
	return c.secretsManager.SecretCheckLogin()
}

// SecretSet the given value/data against the key
func (c *ClusterManager) SecretSet(secretKey string, secretValue interface{}) error {
	return c.secretsManager.SecretSet(secretKey, secretValue)
}

// SecretGet retrieves the value/data for given key
func (c *ClusterManager) SecretGet(secretKey string) (interface{}, error) {
	return c.secretsManager.SecretGet(secretKey)
}

// Uuid returns the unique id of the cluster
func (c *ClusterManager) Uuid() string {
	if len(c.config.ClusterUuid) == 0 {
		return c.config.ClusterId
	}
	return c.config.ClusterUuid
}
