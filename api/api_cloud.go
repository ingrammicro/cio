// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListStorageVolumes returns the list of Volumes as an array of Volume
func (imco *IMCOClient) ListStorageVolumes(serverID string) (volumes []*types.Volume, err error) {
	logger.DebugFuncInfo()

	path := pathStorageVolumes
	if serverID != "" {
		path = fmt.Sprintf(pathCloudServerVolumes, serverID)
	}
	_, err = imco.getAndCheck(path, true, &volumes)
	if err != nil {
		return nil, err
	}
	return volumes, nil
}

// ListCloudProviders returns the list of cloudProviders as an array of CloudProvider
func (imco *IMCOClient) ListCloudProviders() (cloudProviders []*types.CloudProvider, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathCloudCloudProviders, true, &cloudProviders)
	if err != nil {
		return nil, err
	}
	return cloudProviders, nil
}

// ListServerStoragePlans returns the list of storage plans as an array of StoragePlan
func (imco *IMCOClient) ListServerStoragePlans(providerID string) (storagePlans []*types.StoragePlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudCloudProviderStoragePlans, providerID), true, &storagePlans)
	if err != nil {
		return nil, err
	}
	return storagePlans, nil
}

// ListLoadBalancerPlans returns the list of load balancer plans as an array of LoadBalancerPlan
func (imco *IMCOClient) ListLoadBalancerPlans(providerID string,
) (loadBalancerPlans []*types.LoadBalancerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathCloudCloudProviderLoadBalancerPlans, providerID),
		true,
		&loadBalancerPlans,
	)
	if err != nil {
		return nil, err
	}
	return loadBalancerPlans, nil
}

// ListClusterPlans returns the list of cluster plans as an array of ClusterPlan
func (imco *IMCOClient) ListClusterPlans(providerID string) (clusterPlans []*types.ClusterPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudCloudProviderClusterPlans, providerID), true, &clusterPlans)
	if err != nil {
		return nil, err
	}
	return clusterPlans, nil
}

// ListGenericImages returns the list of generic images as an array of GenericImage
func (imco *IMCOClient) ListGenericImages() (genericImages []*types.GenericImage, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathCloudGenericImages, true, &genericImages)
	if err != nil {
		return nil, err
	}
	return genericImages, nil
}

// ListServerArrays returns the list of server arrays as an array of ServerArray
func (imco *IMCOClient) ListServerArrays() (serverArrays []*types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathCloudServerArrays, true, &serverArrays)
	if err != nil {
		return nil, err
	}
	return serverArrays, nil
}

// GetServerArray returns a server array by its ID
func (imco *IMCOClient) GetServerArray(serverArrayID string) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudServerArray, serverArrayID), true, &serverArray)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// CreateServerArray creates a server array
func (imco *IMCOClient) CreateServerArray(serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathCloudServerArrays, serverArrayParams, true, &serverArray)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// UpdateServerArray updates a server array by its ID
func (imco *IMCOClient) UpdateServerArray(serverArrayID string, serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudServerArray, serverArrayID), serverArrayParams, true, &serverArray)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// BootServerArray boots a server array by its ID
func (imco *IMCOClient) BootServerArray(serverArrayID string) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	serverArrayIn := map[string]interface{}{}
	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudServerArrayBoot, serverArrayID), &serverArrayIn, true, &serverArray)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// ShutdownServerArray shuts down a server array by its ID
func (imco *IMCOClient) ShutdownServerArray(serverArrayID string) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	serverArrayIn := map[string]interface{}{}
	_, err = imco.putAndCheck(
		fmt.Sprintf(pathCloudServerArrayShutdown, serverArrayID),
		&serverArrayIn,
		true,
		&serverArray,
	)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// EmptyServerArray empties a server array by its ID
func (imco *IMCOClient) EmptyServerArray(serverArrayID string) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	serverArrayIn := map[string]interface{}{}
	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudServerArrayEmpty, serverArrayID), &serverArrayIn, true, &serverArray)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// EnlargeServerArray enlarges a server array by its ID
func (imco *IMCOClient) EnlargeServerArray(serverArrayID string, serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
		fmt.Sprintf(pathCloudServerArrayServers, serverArrayID),
		serverArrayParams,
		true,
		&serverArray,
	)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// ListServerArrayServers returns the list of servers in a server array as an array of server
func (imco *IMCOClient) ListServerArrayServers(serverArrayID string) (servers []*types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudServerArrayServers, serverArrayID), true, &servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// DeleteServerArray deletes a server array by its ID
func (imco *IMCOClient) DeleteServerArray(serverArrayID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathCloudServerArray, serverArrayID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListServerPlans returns the list of serverPlans as an array of ServerPlan
func (imco *IMCOClient) ListServerPlans(providerID string) (serverPlans []*types.ServerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudCloudProviderServerPlans, providerID), true, &serverPlans)
	if err != nil {
		return nil, err
	}
	return serverPlans, nil
}

// GetServerPlan returns a serverPlan by its ID
func (imco *IMCOClient) GetServerPlan(planID string) (serverPlan *types.ServerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudServerPlan, planID), true, &serverPlan)
	if err != nil {
		return nil, err
	}
	return serverPlan, nil
}

// ListServers returns the list of servers as an array of server
func (imco *IMCOClient) ListServers() (servers []*types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathCloudServers, true, &servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// GetServer returns a server by its ID
func (imco *IMCOClient) GetServer(serverID string) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudServer, serverID), true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// CreateServer creates a server
func (imco *IMCOClient) CreateServer(serverParams *map[string]interface{}) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathCloudServers, serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// UpdateServer updates a server by its ID
func (imco *IMCOClient) UpdateServer(serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudServer, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// BootServer boots a server by its ID
func (imco *IMCOClient) BootServer(serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudServerBoot, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// RebootServer reboots a server by its ID
func (imco *IMCOClient) RebootServer(serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudServerReboot, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// ShutdownServer shuts down a server by its ID
func (imco *IMCOClient) ShutdownServer(serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudServerShutdown, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// OverrideServer overrides a server by its ID
func (imco *IMCOClient) OverrideServer(serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudServerOverride, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// DeleteServer deletes a server by its ID
func (imco *IMCOClient) DeleteServer(serverID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathCloudServer, serverID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListServerFloatingIPs returns the list of floating IPs as an array of FloatingIP
func (imco *IMCOClient) ListServerFloatingIPs(serverID string) (floatingIPs []*types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudServerFloatingIps, serverID), true, &floatingIPs)
	if err != nil {
		return nil, err
	}
	return floatingIPs, nil
}

// ListServerVolumes returns the list of volumes as an array of Volume
func (imco *IMCOClient) ListServerVolumes(serverID string) (volumes []*types.Volume, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudServerVolumes, serverID), true, &volumes)
	if err != nil {
		return nil, err
	}
	return volumes, nil
}

// ListServerEvents returns a list of events by server ID
func (imco *IMCOClient) ListServerEvents(serverID string) (events []*types.Event, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudServerEvents, serverID), true, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// ListOperationalScripts returns a list of scripts by server ID
func (imco *IMCOClient) ListOperationalScripts(serverID string) (scripts []*types.ScriptChar, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudServerOperationalScripts, serverID), true, &scripts)
	if err != nil {
		return nil, err
	}
	return scripts, nil
}

// ExecuteOperationalScript executes an operational script by its server ID and the script id
func (imco *IMCOClient) ExecuteOperationalScript(serverID string, scriptID string,
	serverParams *map[string]interface{},
) (script *types.Event, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
		fmt.Sprintf(pathCloudServerOperationalScriptExecute, serverID, scriptID),
		serverParams,
		true,
		&serverParams,
	)
	if err != nil {
		return nil, err
	}
	return script, nil
}

// ListSSHProfiles returns the list of sshProfiles as an array of SSHProfile
func (imco *IMCOClient) ListSSHProfiles() (sshProfiles []*types.SSHProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathCloudSshProfiles, true, &sshProfiles)
	if err != nil {
		return nil, err
	}
	return sshProfiles, nil
}

// GetSSHProfile returns a sshProfile by its ID
func (imco *IMCOClient) GetSSHProfile(sshProfileID string) (sshProfile *types.SSHProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudSshProfile, sshProfileID), true, &sshProfile)
	if err != nil {
		return nil, err
	}
	return sshProfile, nil
}

// CreateSSHProfile creates a sshProfile
func (imco *IMCOClient) CreateSSHProfile(sshProfileParams *map[string]interface{},
) (sshProfile *types.SSHProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathCloudSshProfiles, sshProfileParams, true, &sshProfile)
	if err != nil {
		return nil, err
	}
	return sshProfile, nil
}

// UpdateSSHProfile updates a sshProfile by its ID
func (imco *IMCOClient) UpdateSSHProfile(sshProfileID string, sshProfileParams *map[string]interface{},
) (sshProfile *types.SSHProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCloudSshProfile, sshProfileID), sshProfileParams, true, &sshProfile)
	if err != nil {
		return nil, err
	}
	return sshProfile, nil
}

// DeleteSSHProfile deletes a sshProfile by its ID
func (imco *IMCOClient) DeleteSSHProfile(sshProfileID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathCloudSshProfile, sshProfileID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListFloatingIPs returns the list of FloatingIPs as an array of FloatingIP
func (imco *IMCOClient) ListFloatingIPs(serverID string) (floatingIPs []*types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	path := pathNetworkFloatingIps
	if serverID != "" {
		path = fmt.Sprintf(pathCloudServerFloatingIps, serverID)
	}
	_, err = imco.getAndCheck(path, true, &floatingIPs)
	if err != nil {
		return nil, err
	}
	return floatingIPs, nil
}

// ListRealms returns the list of realms as an array of Realm
func (imco *IMCOClient) ListRealms(providerID string) (realms []*types.Realm, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudProviderRealms, providerID), true, &realms)
	if err != nil {
		return nil, err
	}
	return realms, nil
}

// GetRealm returns a realm by its ID
func (imco *IMCOClient) GetRealm(realmID string) (realm *types.Realm, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudRealm, realmID), true, &realm)
	if err != nil {
		return nil, err
	}
	return realm, nil
}

// ListRealmNodePoolPlans returns the list of node pool plans as an array of NodePoolPlan
func (imco *IMCOClient) ListRealmNodePoolPlans(realmID string) (nodePoolPlans []*types.NodePoolPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCloudRealmNodePoolPlans, realmID), true, &nodePoolPlans)
	if err != nil {
		return nil, err
	}
	return nodePoolPlans, nil
}
