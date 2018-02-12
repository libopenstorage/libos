// osdconfig is a package to work with distributed config parameters
package osdconfig

import "net/http"

// A config manager interface allows management of osdconfig parameters
// It defines setters, getters and callback management functions
type ConfigManager interface {
	// GetClusterConf fetches cluster configuration data from a backend such as kvdb
	GetClusterConf() (*ClusterConfig, error)

	// Fetch node configuration data using node id
	GetNodeConf(nodeID string) (*NodeConfig, error)

	// SetClusterConf pushes cluster configuration data to the backend
	// It is assumed that the backend will notify the implementor of this interface
	// when a change is triggered
	SetClusterConf(config *ClusterConfig) error

	// SetNodeConf pushes node configuration data to the backend
	// It is assumed that the backend will notify the implementor of this interface
	// when a change is triggered
	SetNodeConf(config *NodeConfig) error

	// TuneCluster registers a user defined function as callback watching for changes
	// in the cluster configuration
	WatchCluster(name string, cb func(config *ClusterConfig) error) error

	// TuneNode registers a user defined function as callback watching for changes
	// in the node configuration
	WatchNode(name string, cb func(config *NodeConfig) error) error

	// GetHTTPFunc generates an http handler
	GetHTTPFunc(state interface{}, fn interface{}) (func(w http.ResponseWriter, r *http.Request), error)

	// Close performs internal cleanup
	Close()
}
