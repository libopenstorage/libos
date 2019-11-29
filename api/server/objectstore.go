package server

import (
	"encoding/json"
	"github.com/libopenstorage/openstorage/api"
	"net/http"
	"strconv"

	clustermanager "github.com/libopenstorage/openstorage/cluster/manager"
	"github.com/libopenstorage/openstorage/objectstore"
)

// swagger:operation GET /cluster/objectstore objectstore objectStoreInspect
//
// Lists Objectstore
//
// This will list current objectstores
//
// ---
// produces:
// - application/json
// parameters:
// - name: ID
//   in: query
//   description: ID of objectstore to inspect
//   type: string
// responses:
//   '200':
//     description: success
//     schema:
//      $ref: '#/definitions/ObjectstoreInfo'
func (c *clusterApi) objectStoreInspect(w http.ResponseWriter, r *http.Request) {
	method := "objectStoreInspect"
	var objstoreID string
	params := r.URL.Query()
	v := params[objectstore.ObjectStoreID]
	if v != nil {
		objstoreID = v[0]
	}

	ctx, err := c.annotateContext(r)
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	if conn, err := c.getConn(); err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		objectStoreClient := api.NewOpenStorageObjectstoreClient(conn)

		resp, err := objectStoreClient.Inspect(ctx, &api.SdkObjectstoreInspectRequest{
			ObjectstoreId: objstoreID,
		})

		if err != nil {
			c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(resp.ObjectstoreStatus); err != nil {
			c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// swagger:operation POST /cluster/objectstore objectstore objectStoreCreate
//
// Create an Object store
//
// This creates the volumes required to run the object store
//
// ---
// produces:
// - application/json
// parameters:
// - name: name
//   in: query
//   description: volume on which object store to run
//   required: true
//   type: string
// responses:
//   '200':
//     description: success
//     schema:
//      $ref: '#/definitions/ObjectstoreInfo'
func (c *clusterApi) objectStoreCreate(w http.ResponseWriter, r *http.Request) {
	method := "objectStoreCreate"
	params := r.URL.Query()
	volumeName := params[objectstore.VolumeName]

	if len(volumeName) == 0 || volumeName[0] == "" {
		c.sendError(c.name, method, w, "Missing volume name", http.StatusBadRequest)
		return
	}

	ctx, err := c.annotateContext(r)
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	if conn, err := c.getConn(); err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		objectStoreClient := api.NewOpenStorageObjectstoreClient(conn)

		resp, err := objectStoreClient.Create(ctx, &api.SdkObjectstoreCreateRequest{
			VolumeId: volumeName[0],
		})

		if err != nil {
			c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(resp.ObjectstoreStatus); err != nil {
			c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// swagger:operation PUT /cluster/objectstore objectstore objectStoreUpdate
//
// Updates object store
//
// This will enable/disable object store functionality.
//
// ---
// produces:
// - application/json
// parameters:
// - name: enable
//   in: query
//   description: enable/disable flag for object store
//   type: boolean
// - name: id
//   in: query
//   description: ID of objectstore to update
//   type: string
// responses:
//   '200':
//     description: success
func (c *clusterApi) objectStoreUpdate(w http.ResponseWriter, r *http.Request) {
	method := "objectStoreUpdate"
	var objstoreID string
	params := r.URL.Query()
	strEnable := params[objectstore.Enable]
	v := params[objectstore.ObjectStoreID]

	if v != nil {
		objstoreID = v[0]
	}

	if len(strEnable) == 0 && strEnable[0] == "" {
		c.sendError(c.name, method, w, "enable parameter not set", http.StatusInternalServerError)
		return
	}

	enable, err := strconv.ParseBool(strEnable[0])
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx, err := c.annotateContext(r)
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	if conn, err := c.getConn(); err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		objectStoreClient := api.NewOpenStorageObjectstoreClient(conn)

		_, err := objectStoreClient.Update(ctx, &api.SdkObjectstoreUpdateRequest{
			ObjectstoreId: objstoreID,
			Enable:        enable,
		})

		if err != nil {
			c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

// swagger:operation DELETE /cluster/objectstore objectstore objectStoreDelete
//
// Delete object store
//
// This will delete object store on node
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: query
//   description: ID of objectstore to delete
//   type: string
// responses:
//   '200':
//     description: success
func (c *clusterApi) objectStoreDelete(w http.ResponseWriter, r *http.Request) {
	method := "objectStoreDelete"
	var objstoreID string
	params := r.URL.Query()
	v := params[objectstore.ObjectStoreID]

	if v != nil {
		objstoreID = v[0]
	}

	inst, err := clustermanager.Inst()
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = inst.ObjectStoreDelete(objstoreID)
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
