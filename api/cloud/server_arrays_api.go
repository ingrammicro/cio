// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCloudServerArrays = "/cloud/server_arrays"
const APIPathCloudServerArray = "/cloud/server_arrays/%s"
const APIPathCloudServerArrayBoot = "/cloud/server_arrays/%s/boot"
const APIPathCloudServerArrayShutdown = "/cloud/server_arrays/%s/shutdown"
const APIPathCloudServerArrayEmpty = "/cloud/server_arrays/%s/empty"
const APIPathCloudServerArrayServers = "/cloud/server_arrays/%s/servers"

// ServerArrayService manages server array operations
type ServerArrayService struct {
	concertoService utils.ConcertoService
}

// NewServerArrayService returns a Concerto server array service
func NewServerArrayService(concertoService utils.ConcertoService) (*ServerArrayService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ServerArrayService{
		concertoService: concertoService,
	}, nil
}

// ListServerArrays returns the list of server arrays as an array of ServerArray
func (sas *ServerArrayService) ListServerArrays() (serverArrays []*types.ServerArray, err error) {
	log.Debug("ListServerArrays")

	data, status, err := sas.concertoService.Get(APIPathCloudServerArrays)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArrays); err != nil {
		return nil, err
	}

	return serverArrays, nil
}

// GetServerArray returns a server array by its ID
func (sas *ServerArrayService) GetServerArray(serverArrayID string) (serverArray *types.ServerArray, err error) {
	log.Debug("GetServerArray")

	data, status, err := sas.concertoService.Get(fmt.Sprintf(APIPathCloudServerArray, serverArrayID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// CreateServerArray creates a server array
func (sas *ServerArrayService) CreateServerArray(
	serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	log.Debug("CreateServerArray")

	data, status, err := sas.concertoService.Post(APIPathCloudServerArrays, serverArrayParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// UpdateServerArray updates a server array by its ID
func (sas *ServerArrayService) UpdateServerArray(
	serverArrayID string,
	serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	log.Debug("UpdateServerArray")

	data, status, err := sas.concertoService.Put(
		fmt.Sprintf(APIPathCloudServerArray, serverArrayID),
		serverArrayParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// BootServerArray boots a server array by its ID
func (sas *ServerArrayService) BootServerArray(
	serverArrayID string,
	serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	log.Debug("BootServerArray")

	data, status, err := sas.concertoService.Put(
		fmt.Sprintf(APIPathCloudServerArrayBoot, serverArrayID),
		serverArrayParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// ShutdownServerArray shuts down a server array by its ID
func (sas *ServerArrayService) ShutdownServerArray(
	serverArrayID string,
	serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	log.Debug("ShutdownServerArray")

	data, status, err := sas.concertoService.Put(
		fmt.Sprintf(APIPathCloudServerArrayShutdown, serverArrayID),
		serverArrayParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// EmptyServerArray empties a server array by its ID
func (sas *ServerArrayService) EmptyServerArray(
	serverArrayID string,
	serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	log.Debug("EmptyServerArray")

	data, status, err := sas.concertoService.Put(
		fmt.Sprintf(APIPathCloudServerArrayEmpty, serverArrayID),
		serverArrayParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// EnlargeServerArray enlarges a server array by its ID
func (sas *ServerArrayService) EnlargeServerArray(
	serverArrayID string,
	serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	log.Debug("EnlargeServerArray")

	data, status, err := sas.concertoService.Post(
		fmt.Sprintf(APIPathCloudServerArrayServers, serverArrayID),
		serverArrayParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverArray); err != nil {
		return nil, err
	}

	return serverArray, nil
}

// ListServerArrayServers returns the list of servers in a server array as an array of Server
func (sas *ServerArrayService) ListServerArrayServers(serverArrayID string) (servers []*types.Server, err error) {
	log.Debug("ListServerArrayServers")

	data, status, err := sas.concertoService.Get(fmt.Sprintf(APIPathCloudServerArrayServers, serverArrayID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &servers); err != nil {
		return nil, err
	}

	return servers, nil
}

// DeleteServerArray deletes a server array by its ID
func (sas *ServerArrayService) DeleteServerArray(serverArrayID string) (err error) {
	log.Debug("DeleteServerArray")

	data, status, err := sas.concertoService.Delete(fmt.Sprintf(APIPathCloudServerArray, serverArrayID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
