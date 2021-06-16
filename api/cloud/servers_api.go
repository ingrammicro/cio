// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCloudServers = "/cloud/servers"
const APIPathCloudServer = "/cloud/servers/%s"
const APIPathCloudServerBoot = "/cloud/servers/%s/boot"
const APIPathCloudServerReboot = "/cloud/servers/%s/reboot"
const APIPathCloudServerShutdown = "/cloud/servers/%s/shutdown"
const APIPathCloudServerOverride = "/cloud/servers/%s/override"
const APIPathCloudServerFloatingIPs = "/cloud/servers/%s/floating_ips"
const APIPathCloudServerVolumes = "/cloud/servers/%s/volumes"
const APIPathCloudServerEvents = "/cloud/servers/%s/events"
const APIPathCloudServerOperationalScripts = "/cloud/servers/%s/operational_scripts"
const APIPathCloudServerOperationalScriptExecute = "/cloud/servers/%s/operational_scripts/%s/execute"

// ServerService manages server operations
type ServerService struct {
	concertoService utils.ConcertoService
}

// NewServerService returns a Concerto server service
func NewServerService(concertoService utils.ConcertoService) (*ServerService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ServerService{
		concertoService: concertoService,
	}, nil
}

// ListServers returns the list of servers as an array of Server
func (ss *ServerService) ListServers() (servers []*types.Server, err error) {
	log.Debug("ListServers")

	data, status, err := ss.concertoService.Get(APIPathCloudServers)
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

// GetServer returns a server by its ID
func (ss *ServerService) GetServer(serverID string) (server *types.Server, err error) {
	log.Debug("GetServer")

	data, status, err := ss.concertoService.Get(fmt.Sprintf(APIPathCloudServer, serverID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// CreateServer creates a server
func (ss *ServerService) CreateServer(serverParams *map[string]interface{}) (server *types.Server, err error) {
	log.Debug("CreateServer")

	data, status, err := ss.concertoService.Post(APIPathCloudServers, serverParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// UpdateServer updates a server by its ID
func (ss *ServerService) UpdateServer(
	serverID string,
	serverParams *map[string]interface{},
) (server *types.Server, err error) {
	log.Debug("UpdateServer")

	data, status, err := ss.concertoService.Put(fmt.Sprintf(APIPathCloudServer, serverID), serverParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// BootServer boots a server by its ID
func (ss *ServerService) BootServer(
	serverID string,
	serverParams *map[string]interface{},
) (server *types.Server, err error) {
	log.Debug("BootServer")

	data, status, err := ss.concertoService.Put(fmt.Sprintf(APIPathCloudServerBoot, serverID), serverParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// RebootServer reboots a server by its ID
func (ss *ServerService) RebootServer(
	serverID string,
	serverParams *map[string]interface{},
) (server *types.Server, err error) {
	log.Debug("RebootServer")

	data, status, err := ss.concertoService.Put(fmt.Sprintf(APIPathCloudServerReboot, serverID), serverParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// ShutdownServer shuts down a server by its ID
func (ss *ServerService) ShutdownServer(
	serverID string,
	serverParams *map[string]interface{},
) (server *types.Server, err error) {
	log.Debug("ShutdownServer")

	data, status, err := ss.concertoService.Put(fmt.Sprintf(APIPathCloudServerShutdown, serverID), serverParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// OverrideServer overrides a server by its ID
func (ss *ServerService) OverrideServer(
	serverID string,
	serverParams *map[string]interface{},
) (server *types.Server, err error) {
	log.Debug("OverrideServer")

	data, status, err := ss.concertoService.Put(fmt.Sprintf(APIPathCloudServerOverride, serverID), serverParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// DeleteServer deletes a server by its ID
func (ss *ServerService) DeleteServer(serverID string) (err error) {
	log.Debug("DeleteServer")

	data, status, err := ss.concertoService.Delete(fmt.Sprintf(APIPathCloudServer, serverID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ListServerFloatingIPs returns the list of floating IPs as an array of FloatingIP
func (ss *ServerService) ListServerFloatingIPs(serverID string) (floatingIPs []*types.FloatingIP, err error) {
	log.Debug("ListServerFloatingIPs")

	data, status, err := ss.concertoService.Get(fmt.Sprintf(APIPathCloudServerFloatingIPs, serverID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIPs); err != nil {
		return nil, err
	}

	return floatingIPs, nil
}

// ListServerVolumes returns the list of volumes as an array of Volume
func (ss *ServerService) ListServerVolumes(serverID string) (volumes []*types.Volume, err error) {
	log.Debug("ListServerVolumes")

	data, status, err := ss.concertoService.Get(fmt.Sprintf(APIPathCloudServerVolumes, serverID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volumes); err != nil {
		return nil, err
	}

	return volumes, nil
}

//======= Events ==========

// ListEvents returns a list of events by server ID
func (ss *ServerService) ListEvents(serverID string) (events []*types.Event, err error) {
	log.Debug("ListEvents")

	data, status, err := ss.concertoService.Get(fmt.Sprintf(APIPathCloudServerEvents, serverID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	return events, nil
}

//======= Operational Scripts ==========

// ListOperationalScripts returns a list of scripts by server ID
func (ss *ServerService) ListOperationalScripts(serverID string) (scripts []*types.ScriptChar, err error) {
	log.Debug("ListOperationalScripts")

	data, status, err := ss.concertoService.Get(fmt.Sprintf(APIPathCloudServerOperationalScripts, serverID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

// ExecuteOperationalScript executes an operational script by its server ID and the script id
func (ss *ServerService) ExecuteOperationalScript(
	serverID string,
	scriptID string,
	serverParams *map[string]interface{},
) (script *types.Event, err error) {
	log.Debug("ExecuteOperationalScript")

	data, status, err := ss.concertoService.Put(
		fmt.Sprintf(APIPathCloudServerOperationalScriptExecute, serverID, scriptID),
		serverParams,
	)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}
