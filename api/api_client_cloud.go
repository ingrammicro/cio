// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"golang.org/x/net/context"
)

// ListStorageVolumes returns the list of Volumes as an array of Volume
func (imco *ClientAPI) ListStorageVolumes(ctx context.Context, serverID string) (volumes []*types.Volume, err error) {
	logger.DebugFuncInfo()

	path := pathStorageVolumes
	if serverID != "" {
		path = fmt.Sprintf(pathCloudServerVolumes, serverID)
	}
	_, err = imco.GetAndCheck(ctx, path, true, &volumes)
	if err != nil {
		return nil, err
	}
	return volumes, nil
}

// ListCloudProviders returns the list of cloudProviders as an array of CloudProvider
func (imco *ClientAPI) ListCloudProviders(ctx context.Context) (cloudProviders []*types.CloudProvider, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathCloudCloudProviders, true, &cloudProviders)
	if err != nil {
		return nil, err
	}
	return cloudProviders, nil
}

// ListServerStoragePlans returns the list of storage plans as an array of StoragePlan
func (imco *ClientAPI) ListServerStoragePlans(ctx context.Context, providerID string,
) (storagePlans []*types.StoragePlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudCloudProviderStoragePlans, providerID), true, &storagePlans)
	if err != nil {
		return nil, err
	}
	return storagePlans, nil
}

// ListLoadBalancerPlans returns the list of load balancer plans as an array of LoadBalancerPlan
func (imco *ClientAPI) ListLoadBalancerPlans(ctx context.Context, providerID string,
) (loadBalancerPlans []*types.LoadBalancerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
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
func (imco *ClientAPI) ListClusterPlans(ctx context.Context, providerID string,
) (clusterPlans []*types.ClusterPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudCloudProviderClusterPlans, providerID), true, &clusterPlans)
	if err != nil {
		return nil, err
	}
	return clusterPlans, nil
}

// ListGenericImages returns the list of generic images as an array of GenericImage
func (imco *ClientAPI) ListGenericImages(ctx context.Context) (genericImages []*types.GenericImage, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathCloudGenericImages, true, &genericImages)
	if err != nil {
		return nil, err
	}
	return genericImages, nil
}

// ListServerArrays returns the list of server arrays as an array of ServerArray
func (imco *ClientAPI) ListServerArrays(ctx context.Context) (serverArrays []*types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathCloudServerArrays, true, &serverArrays)
	if err != nil {
		return nil, err
	}
	return serverArrays, nil
}

// GetServerArray returns a server array by its ID
func (imco *ClientAPI) GetServerArray(ctx context.Context, serverArrayID string,
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudServerArray, serverArrayID), true, &serverArray)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// CreateServerArray creates a server array
func (imco *ClientAPI) CreateServerArray(ctx context.Context, serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathCloudServerArrays, serverArrayParams, true, &serverArray)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// UpdateServerArray updates a server array by its ID
func (imco *ClientAPI) UpdateServerArray(ctx context.Context, serverArrayID string,
	serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(
		ctx,
		fmt.Sprintf(pathCloudServerArray, serverArrayID),
		serverArrayParams,
		true,
		&serverArray,
	)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// BootServerArray boots a server array by its ID
func (imco *ClientAPI) BootServerArray(ctx context.Context, serverArrayID string,
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	serverArrayIn := map[string]interface{}{}
	_, err = imco.PutAndCheck(
		ctx,
		fmt.Sprintf(pathCloudServerArrayBoot, serverArrayID),
		&serverArrayIn,
		true,
		&serverArray,
	)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// ShutdownServerArray shuts down a server array by its ID
func (imco *ClientAPI) ShutdownServerArray(ctx context.Context, serverArrayID string,
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	serverArrayIn := map[string]interface{}{}
	_, err = imco.PutAndCheck(ctx,
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
func (imco *ClientAPI) EmptyServerArray(ctx context.Context, serverArrayID string,
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	serverArrayIn := map[string]interface{}{}
	_, err = imco.PutAndCheck(
		ctx,
		fmt.Sprintf(pathCloudServerArrayEmpty, serverArrayID),
		&serverArrayIn,
		true,
		&serverArray,
	)
	if err != nil {
		return nil, err
	}
	return serverArray, nil
}

// EnlargeServerArray enlarges a server array by its ID
func (imco *ClientAPI) EnlargeServerArray(ctx context.Context, serverArrayID string,
	serverArrayParams *map[string]interface{},
) (serverArray *types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
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
func (imco *ClientAPI) ListServerArrayServers(ctx context.Context, serverArrayID string,
) (servers []*types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudServerArrayServers, serverArrayID), true, &servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// DeleteServerArray deletes a server array by its ID
func (imco *ClientAPI) DeleteServerArray(ctx context.Context, serverArrayID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathCloudServerArray, serverArrayID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListServerPlans returns the list of serverPlans as an array of ServerPlan
func (imco *ClientAPI) ListServerPlans(ctx context.Context, providerID string, realmID string,
) (serverPlans []*types.ServerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(pathCloudCloudProviderServerPlansByRealm, providerID, realmID),
		true,
		&serverPlans)
	if err != nil {
		return nil, err
	}
	return serverPlans, nil
}

// GetServerPlan returns a serverPlan by its ID
func (imco *ClientAPI) GetServerPlan(ctx context.Context, planID string) (serverPlan *types.ServerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudServerPlan, planID), true, &serverPlan)
	if err != nil {
		return nil, err
	}
	return serverPlan, nil
}

// ListServers returns the list of servers as an array of server
func (imco *ClientAPI) ListServers(ctx context.Context) (servers []*types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathCloudServers, true, &servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// GetServer returns a server by its ID
func (imco *ClientAPI) GetServer(ctx context.Context, serverID string) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudServer, serverID), true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// CreateServer creates a server
func (imco *ClientAPI) CreateServer(ctx context.Context, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathCloudServers, serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// UpdateServer updates a server by its ID
func (imco *ClientAPI) UpdateServer(ctx context.Context, serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathCloudServer, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// BootServer boots a server by its ID
func (imco *ClientAPI) BootServer(ctx context.Context, serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathCloudServerBoot, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// RebootServer reboots a server by its ID
func (imco *ClientAPI) RebootServer(ctx context.Context, serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathCloudServerReboot, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// ShutdownServer shuts down a server by its ID
func (imco *ClientAPI) ShutdownServer(ctx context.Context, serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathCloudServerShutdown, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// OverrideServer overrides a server by its ID
func (imco *ClientAPI) OverrideServer(ctx context.Context, serverID string, serverParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathCloudServerOverride, serverID), serverParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// DeleteServer deletes a server by its ID
func (imco *ClientAPI) DeleteServer(ctx context.Context, serverID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathCloudServer, serverID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListServerFloatingIPs returns the list of floating IPs as an array of FloatingIP
func (imco *ClientAPI) ListServerFloatingIPs(ctx context.Context, serverID string,
) (floatingIPs []*types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudServerFloatingIps, serverID), true, &floatingIPs)
	if err != nil {
		return nil, err
	}
	return floatingIPs, nil
}

// ListServerVolumes returns the list of volumes as an array of Volume
func (imco *ClientAPI) ListServerVolumes(ctx context.Context, serverID string) (volumes []*types.Volume, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudServerVolumes, serverID), true, &volumes)
	if err != nil {
		return nil, err
	}
	return volumes, nil
}

// ListServerEvents returns a list of events by server ID
func (imco *ClientAPI) ListServerEvents(ctx context.Context, serverID string) (events []*types.Event, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudServerEvents, serverID), true, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// ListOperationalScripts returns a list of scripts by server ID
func (imco *ClientAPI) ListOperationalScripts(ctx context.Context, serverID string,
) (scripts []*types.ScriptChar, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudServerOperationalScripts, serverID), true, &scripts)
	if err != nil {
		return nil, err
	}
	return scripts, nil
}

// ExecuteOperationalScript executes an operational script by its server ID and the script id
func (imco *ClientAPI) ExecuteOperationalScript(ctx context.Context, serverID string, scriptID string,
	serverParams *map[string]interface{},
) (event *types.Event, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathCloudServerOperationalScriptExecute, serverID, scriptID),
		serverParams,
		true,
		&event,
	)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// ListSSHProfiles returns the list of sshProfiles as an array of SSHProfile
func (imco *ClientAPI) ListSSHProfiles(ctx context.Context) (sshProfiles []*types.SSHProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathCloudSshProfiles, true, &sshProfiles)
	if err != nil {
		return nil, err
	}
	return sshProfiles, nil
}

// GetSSHProfile returns a sshProfile by its ID
func (imco *ClientAPI) GetSSHProfile(ctx context.Context, sshProfileID string,
) (sshProfile *types.SSHProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudSshProfile, sshProfileID), true, &sshProfile)
	if err != nil {
		return nil, err
	}
	return sshProfile, nil
}

// CreateSSHProfile creates a sshProfile
func (imco *ClientAPI) CreateSSHProfile(ctx context.Context, sshProfileParams *map[string]interface{},
) (sshProfile *types.SSHProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathCloudSshProfiles, sshProfileParams, true, &sshProfile)
	if err != nil {
		return nil, err
	}
	return sshProfile, nil
}

// UpdateSSHProfile updates a sshProfile by its ID
func (imco *ClientAPI) UpdateSSHProfile(ctx context.Context, sshProfileID string,
	sshProfileParams *map[string]interface{},
) (sshProfile *types.SSHProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathCloudSshProfile, sshProfileID), sshProfileParams, true, &sshProfile)
	if err != nil {
		return nil, err
	}
	return sshProfile, nil
}

// DeleteSSHProfile deletes a sshProfile by its ID
func (imco *ClientAPI) DeleteSSHProfile(ctx context.Context, sshProfileID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathCloudSshProfile, sshProfileID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListFloatingIPs returns the list of FloatingIPs as an array of FloatingIP
func (imco *ClientAPI) ListFloatingIPs(ctx context.Context, serverID string,
) (floatingIPs []*types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	path := pathNetworkFloatingIps
	if serverID != "" {
		path = fmt.Sprintf(pathCloudServerFloatingIps, serverID)
	}
	_, err = imco.GetAndCheck(ctx, path, true, &floatingIPs)
	if err != nil {
		return nil, err
	}
	return floatingIPs, nil
}

// ListRealms returns the list of realms as an array of Realm
func (imco *ClientAPI) ListRealms(ctx context.Context, providerID string) (realms []*types.Realm, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudProviderRealms, providerID), true, &realms)
	if err != nil {
		return nil, err
	}
	return realms, nil
}

// GetRealm returns a realm by its ID
func (imco *ClientAPI) GetRealm(ctx context.Context, realmID string) (realm *types.Realm, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudRealm, realmID), true, &realm)
	if err != nil {
		return nil, err
	}
	return realm, nil
}

// ListRealmNodePoolPlans returns the list of node pool plans as an array of NodePoolPlan
func (imco *ClientAPI) ListRealmNodePoolPlans(ctx context.Context, realmID string,
) (nodePoolPlans []*types.NodePoolPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCloudRealmNodePoolPlans, realmID), true, &nodePoolPlans)
	if err != nil {
		return nil, err
	}
	return nodePoolPlans, nil
}
