// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"golang.org/x/net/context"
)

// ListCertificates returns the list of certificates in a load balancer by its ID, as an array of Certificate
func (imco *ClientAPI) ListCertificates(ctx context.Context, loadBalancerID string,
) (certificates []*types.Certificate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerCertificates, loadBalancerID),
		true,
		&certificates,
	)
	if err != nil {
		return nil, err
	}
	return certificates, nil
}

// GetCertificate returns a certificate by its ID
func (imco *ClientAPI) GetCertificate(ctx context.Context, loadBalancerID string, certificateID string,
) (certificate *types.Certificate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerCertificate, loadBalancerID, certificateID),
		true,
		&certificate,
	)
	if err != nil {
		return nil, err
	}
	return certificate, nil
}

// CreateCertificate creates a certificate in a load balancer by its ID
func (imco *ClientAPI) CreateCertificate(ctx context.Context, loadBalancerID string,
	certificateParams *map[string]interface{},
) (certificate *types.Certificate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerCertificates, loadBalancerID),
		certificateParams,
		true,
		&certificate,
	)
	if err != nil {
		return nil, err
	}
	return certificate, nil
}

// UpdateCertificate updates a certificate by its ID
func (imco *ClientAPI) UpdateCertificate(ctx context.Context, loadBalancerID string, certificateID string,
	certificateParams *map[string]interface{},
) (certificate *types.Certificate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerCertificate, loadBalancerID, certificateID),
		certificateParams,
		true,
		&certificate,
	)
	if err != nil {
		return nil, err
	}
	return certificate, nil
}

// DeleteCertificate deletes a certificate by its ID
func (imco *ClientAPI) DeleteCertificate(ctx context.Context, loadBalancerID string, certificateID string,
) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerCertificate, loadBalancerID, certificateID),
		true,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

// ListFirewallProfiles returns the list of firewallProfiles as an array of FirewallProfile
func (imco *ClientAPI) ListFirewallProfiles(ctx context.Context,
) (firewallProfiles []*types.FirewallProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathNetworkFirewallProfiles, true, &firewallProfiles)
	if err != nil {
		return nil, err
	}
	return firewallProfiles, nil
}

// GetFirewallProfile returns a firewallProfile by its ID
func (imco *ClientAPI) GetFirewallProfile(ctx context.Context, firewallProfileID string,
) (firewallProfile *types.FirewallProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkFirewallProfile, firewallProfileID), true, &firewallProfile)
	if err != nil {
		return nil, err
	}
	return firewallProfile, nil
}

// CreateFirewallProfile creates a firewallProfile
func (imco *ClientAPI) CreateFirewallProfile(ctx context.Context, firewallProfileParams *map[string]interface{},
) (firewallProfile *types.FirewallProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathNetworkFirewallProfiles, firewallProfileParams, true, &firewallProfile)
	if err != nil {
		return nil, err
	}
	return firewallProfile, nil
}

// UpdateFirewallProfile updates a firewallProfile by its ID
func (imco *ClientAPI) UpdateFirewallProfile(ctx context.Context, firewallProfileID string,
	firewallProfileParams *map[string]interface{},
) (firewallProfile *types.FirewallProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathNetworkFirewallProfile, firewallProfileID),
		firewallProfileParams,
		true,
		&firewallProfile,
	)
	if err != nil {
		return nil, err
	}
	return firewallProfile, nil
}

// DeleteFirewallProfile deletes a firewallProfile by its ID
func (imco *ClientAPI) DeleteFirewallProfile(ctx context.Context, firewallProfileID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkFirewallProfile, firewallProfileID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetFloatingIP returns a FloatingIP by its ID
func (imco *ClientAPI) GetFloatingIP(ctx context.Context, floatingIPID string,
) (floatingIP *types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkFloatingIp, floatingIPID), true, &floatingIP)
	if err != nil {
		return nil, err
	}
	return floatingIP, nil
}

// CreateFloatingIP creates a FloatingIP
func (imco *ClientAPI) CreateFloatingIP(ctx context.Context, floatingIPParams *map[string]interface{},
) (floatingIP *types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathNetworkFloatingIps, floatingIPParams, true, &floatingIP)
	if err != nil {
		return nil, err
	}
	return floatingIP, nil
}

// UpdateFloatingIP updates a FloatingIP by its ID
func (imco *ClientAPI) UpdateFloatingIP(ctx context.Context, floatingIPID string,
	floatingIPParams *map[string]interface{},
) (floatingIP *types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathNetworkFloatingIp, floatingIPID),
		floatingIPParams,
		true,
		&floatingIP,
	)
	if err != nil {
		return nil, err
	}
	return floatingIP, nil
}

// AttachFloatingIP attaches a FloatingIP by its ID
func (imco *ClientAPI) AttachFloatingIP(ctx context.Context, floatingIPID string,
	floatingIPParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(pathNetworkFloatingIpAttachedServer, floatingIPID),
		floatingIPParams,
		true,
		&server,
	)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// DetachFloatingIP detaches a FloatingIP by its ID
func (imco *ClientAPI) DetachFloatingIP(ctx context.Context, floatingIPID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkFloatingIpAttachedServer, floatingIPID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFloatingIP deletes a FloatingIP by its ID
func (imco *ClientAPI) DeleteFloatingIP(ctx context.Context, floatingIPID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkFloatingIp, floatingIPID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DiscardFloatingIP discards a FloatingIP by its ID
func (imco *ClientAPI) DiscardFloatingIP(ctx context.Context, floatingIPID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkFloatingIpDiscard, floatingIPID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListListeners returns the list of listeners in a load balancer by its ID, as an array of Listener
func (imco *ClientAPI) ListListeners(ctx context.Context, loadBalancerID string,
) (listeners []*types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkLoadBalancerListeners, loadBalancerID), true, &listeners)
	if err != nil {
		return nil, err
	}
	return listeners, nil
}

// GetListener returns a listener by its ID
func (imco *ClientAPI) GetListener(ctx context.Context, listenerID string) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkListener, listenerID), true, &listener)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// CreateListener creates a listener in a load balancer by its ID
func (imco *ClientAPI) CreateListener(ctx context.Context, loadBalancerID string,
	listenerParams *map[string]interface{},
) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerListeners, loadBalancerID),
		listenerParams,
		true,
		&listener,
	)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// UpdateListener updates a listener by its ID
func (imco *ClientAPI) UpdateListener(ctx context.Context, listenerID string, listenerParams *map[string]interface{},
) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathNetworkListener, listenerID), listenerParams, true, &listener)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// DeleteListener deletes a listener by its ID
func (imco *ClientAPI) DeleteListener(ctx context.Context, listenerID string) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkListener, listenerID), true, &listener)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// RetryListener retries a listener by its ID
func (imco *ClientAPI) RetryListener(ctx context.Context, listenerID string, listenerParams *map[string]interface{},
) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathNetworkListenerRetry, listenerID), listenerParams, true, &listener)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// ListRules returns the list of rules in a listener by its ID, as an array of ListenerRule
func (imco *ClientAPI) ListRules(ctx context.Context, listenerID string,
) (listenerRules []*types.ListenerRule, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkListenerRules, listenerID), true, &listenerRules)
	if err != nil {
		return nil, err
	}
	return listenerRules, nil
}

// CreateRule creates a rule in a listener by its ID
func (imco *ClientAPI) CreateRule(ctx context.Context, listenerID string, listenerRuleParams *map[string]interface{},
) (listenerRule *types.ListenerRule, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(pathNetworkListenerRules, listenerID),
		listenerRuleParams,
		true,
		&listenerRule,
	)
	if err != nil {
		return nil, err
	}
	return listenerRule, nil
}

// UpdateRule updates a rule in a listener by its ID
func (imco *ClientAPI) UpdateRule(ctx context.Context, listenerID string, listenerRuleID string,
	listenerRuleParams *map[string]interface{},
) (listenerRule *types.ListenerRule, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathNetworkListenerRule, listenerID, listenerRuleID),
		listenerRuleParams,
		true,
		&listenerRule,
	)
	if err != nil {
		return nil, err
	}
	return listenerRule, nil
}

// DeleteRule deletes a rule in a listener by given IDs
func (imco *ClientAPI) DeleteRule(ctx context.Context, listenerID string, listenerRuleID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkListenerRule, listenerID, listenerRuleID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListLoadBalancers returns the list of load balancers as an array of LoadBalancer
func (imco *ClientAPI) ListLoadBalancers(ctx context.Context) (loadBalancers []*types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathNetworkLoadBalancers, true, &loadBalancers)
	if err != nil {
		return nil, err
	}
	return loadBalancers, nil
}

// GetLoadBalancer returns a load balancer by its ID
func (imco *ClientAPI) GetLoadBalancer(ctx context.Context, loadBalancerID string,
) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkLoadBalancer, loadBalancerID), true, &loadBalancer)
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

// CreateLoadBalancer creates a load balancer
func (imco *ClientAPI) CreateLoadBalancer(ctx context.Context, loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathNetworkLoadBalancers, loadBalancerParams, true, &loadBalancer)
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

// UpdateLoadBalancer updates a load balancer by its ID
func (imco *ClientAPI) UpdateLoadBalancer(ctx context.Context, loadBalancerID string,
	loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancer, loadBalancerID),
		loadBalancerParams,
		true,
		&loadBalancer,
	)
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

// DeleteLoadBalancer deletes a load balancer by its ID
func (imco *ClientAPI) DeleteLoadBalancer(ctx context.Context, loadBalancerID string,
) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkLoadBalancer, loadBalancerID), true, &loadBalancer)
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

// RetryLoadBalancer retries a load balancer by its ID
func (imco *ClientAPI) RetryLoadBalancer(ctx context.Context, loadBalancerID string,
	loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerRetry, loadBalancerID),
		loadBalancerParams,
		true,
		&loadBalancer,
	)
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

// GetLoadBalancerPlan returns a load balancer plan by its ID
func (imco *ClientAPI) GetLoadBalancerPlan(ctx context.Context, loadBalancerPlanID string,
) (loadBalancerPlan *types.LoadBalancerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerPlan, loadBalancerPlanID),
		true,
		&loadBalancerPlan,
	)
	if err != nil {
		return nil, err
	}
	return loadBalancerPlan, nil
}

// ListSubnets returns the list of Subnets of a VPC as an array of Subnet
func (imco *ClientAPI) ListSubnets(ctx context.Context, vpcID string) (subnets []*types.Subnet, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkVpcSubnets, vpcID), true, &subnets)
	if err != nil {
		return nil, err
	}
	return subnets, nil
}

// GetSubnet returns a Subnet by its ID
func (imco *ClientAPI) GetSubnet(ctx context.Context, subnetID string) (subnet *types.Subnet, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkSubnet, subnetID), true, &subnet)
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

// CreateSubnet creates a Subnet
func (imco *ClientAPI) CreateSubnet(ctx context.Context, vpcID string, subnetParams *map[string]interface{},
) (subnet *types.Subnet, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, fmt.Sprintf(pathNetworkVpcSubnets, vpcID), subnetParams, true, &subnet)
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

// UpdateSubnet updates a Subnet by its ID
func (imco *ClientAPI) UpdateSubnet(ctx context.Context, subnetID string, subnetParams *map[string]interface{},
) (subnet *types.Subnet, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathNetworkSubnet, subnetID), subnetParams, true, &subnet)
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

// DeleteSubnet deletes a Subnet by its ID
func (imco *ClientAPI) DeleteSubnet(ctx context.Context, subnetID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkSubnet, subnetID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListSubnetServers returns the list of Servers of a Subnet as an array of server
func (imco *ClientAPI) ListSubnetServers(ctx context.Context, subnetID string) (servers []*types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkSubnetServers, subnetID), true, &servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// ListSubnetServerArrays returns the list of server arrays of a Subnet as an array of ServerArray
func (imco *ClientAPI) ListSubnetServerArrays(ctx context.Context, subnetID string,
) (serverArrays []*types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkSubnetServerArrays, subnetID), true, &serverArrays)
	if err != nil {
		return nil, err
	}
	return serverArrays, nil
}

// ListTargetGroups returns the list of target groups in a load balancer by its ID, as an array of TargetGroup
func (imco *ClientAPI) ListTargetGroups(ctx context.Context, loadBalancerID string,
) (targetGroups []*types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerTargetGroups, loadBalancerID),
		true,
		&targetGroups,
	)
	if err != nil {
		return nil, err
	}
	return targetGroups, nil
}

// GetTargetGroup returns a target group by its ID
func (imco *ClientAPI) GetTargetGroup(ctx context.Context, targetGroupID string,
) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkTargetGroup, targetGroupID), true, &targetGroup)
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

// CreateTargetGroup creates a target group in a load balancer by its ID
func (imco *ClientAPI) CreateTargetGroup(ctx context.Context, loadBalancerID string,
	targetGroupParams *map[string]interface{},
) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(pathNetworkLoadBalancerTargetGroups, loadBalancerID),
		targetGroupParams,
		true,
		&targetGroup,
	)
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

// UpdateTargetGroup updates a target group by its ID
func (imco *ClientAPI) UpdateTargetGroup(ctx context.Context, targetGroupID string,
	targetGroupParams *map[string]interface{},
) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathNetworkTargetGroup, targetGroupID),
		targetGroupParams,
		true,
		&targetGroup,
	)
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

// DeleteTargetGroup deletes a target group by its ID
func (imco *ClientAPI) DeleteTargetGroup(ctx context.Context, targetGroupID string,
) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkTargetGroup, targetGroupID), true, &targetGroup)
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

// RetryTargetGroup retries a target group by its ID
func (imco *ClientAPI) RetryTargetGroup(ctx context.Context, targetGroupID string,
	targetGroupParams *map[string]interface{},
) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathNetworkTargetGroupRetry, targetGroupID),
		targetGroupParams,
		true,
		&targetGroup,
	)
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

// ListTargets returns the list of targets in a target group by its ID, as an array of Target
func (imco *ClientAPI) ListTargets(ctx context.Context, targetGroupID string) (targets []*types.Target, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkTargetGroupTargets, targetGroupID), true, &targets)
	if err != nil {
		return nil, err
	}
	return targets, nil
}

// CreateTarget creates a target in a target group by its ID
func (imco *ClientAPI) CreateTarget(ctx context.Context, targetGroupID string, targetParams *map[string]interface{},
) (target *types.Target, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(pathNetworkTargetGroupTargets, targetGroupID),
		targetParams,
		true,
		&target,
	)
	if err != nil {
		return nil, err
	}
	return target, nil
}

// DeleteTarget deletes a target in a target group by given IDs and resource type
func (imco *ClientAPI) DeleteTarget(ctx context.Context, targetGroupID string, targetResourceType string,
	targetResourceID string,
) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx,
		fmt.Sprintf(pathNetworkTargetGroupTarget, targetGroupID, targetResourceType, targetResourceID),
		true,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

// ListVPCs returns the list of VPCs as an array of VPC
func (imco *ClientAPI) ListVPCs(ctx context.Context) (vpcs []*types.Vpc, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathNetworkVpcs, true, &vpcs)
	if err != nil {
		return nil, err
	}
	return vpcs, nil
}

// GetVPC returns a VPC by its ID
func (imco *ClientAPI) GetVPC(ctx context.Context, vpcID string) (vpc *types.Vpc, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkVpc, vpcID), true, &vpc)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

// CreateVPC creates a VPC
func (imco *ClientAPI) CreateVPC(ctx context.Context, vpcParams *map[string]interface{}) (vpc *types.Vpc, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathNetworkVpcs, vpcParams, true, &vpc)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

// UpdateVPC updates a VPC by its ID
func (imco *ClientAPI) UpdateVPC(ctx context.Context, vpcID string,
	vpcParams *map[string]interface{},
) (vpc *types.Vpc, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathNetworkVpc, vpcID), vpcParams, true, &vpc)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

// DeleteVPC deletes a VPC by its ID
func (imco *ClientAPI) DeleteVPC(ctx context.Context, vpcID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkVpc, vpcID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DiscardVPC discards a VPC by its ID
func (imco *ClientAPI) DiscardVPC(ctx context.Context, vpcID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkVpcDiscard, vpcID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetVPN returns a VPN by VPC ID
func (imco *ClientAPI) GetVPN(ctx context.Context, vpcID string) (vpn *types.Vpn, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkVpcVpn, vpcID), true, &vpn)
	if err != nil {
		return nil, err
	}
	return vpn, nil
}

// CreateVPN creates a VPN
func (imco *ClientAPI) CreateVPN(ctx context.Context, vpcID string,
	vpnParams *map[string]interface{},
) (vpn *types.Vpn, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, fmt.Sprintf(pathNetworkVpcVpn, vpcID), vpnParams, true, &vpn)
	if err != nil {
		return nil, err
	}
	return vpn, nil
}

// DeleteVPN deletes VPN by VPC ID
func (imco *ClientAPI) DeleteVPN(ctx context.Context, vpcID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkVpcVpn, vpcID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListVPNPlans returns the list of VPN plans for a given VPC ID
func (imco *ClientAPI) ListVPNPlans(ctx context.Context, vpcID string) (vpnPlans []*types.VpnPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkVpcVpnPlans, vpcID), true, &vpnPlans)
	if err != nil {
		return nil, err
	}
	return vpnPlans, nil
}

// ListDomains returns the list of domains as an array of Domain
func (imco *ClientAPI) ListDomains(ctx context.Context) (domains []*types.Domain, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathNetworkDnsDomains, true, &domains)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

// GetDomain returns a domain by its ID
func (imco *ClientAPI) GetDomain(ctx context.Context, domainID string) (domain *types.Domain, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkDnsDomain, domainID), true, &domain)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

// CreateDomain creates a domain
func (imco *ClientAPI) CreateDomain(ctx context.Context, domainParams *map[string]interface{},
) (domain *types.Domain, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathNetworkDnsDomains, domainParams, true, &domain)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

// DeleteDomain deletes a domain by its ID
func (imco *ClientAPI) DeleteDomain(ctx context.Context, domainID string) (domain *types.Domain, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkDnsDomain, domainID), true, &domain)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

// RetryDomain retries a domain by its ID
func (imco *ClientAPI) RetryDomain(ctx context.Context, domainID string) (domain *types.Domain, err error) {
	logger.DebugFuncInfo()

	domainParams := new(map[string]interface{})
	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathNetworkDnsDomainRetry, domainID), domainParams, true, &domain)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

// ListRecords returns the list of records as an array of Record for given domain
func (imco *ClientAPI) ListRecords(ctx context.Context, domainID string) (records []*types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkDnsDomainRecords, domainID), true, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetRecord returns a record by its ID
func (imco *ClientAPI) GetRecord(ctx context.Context, recordID string) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathNetworkDnsRecord, recordID), true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// CreateRecord creates a record in a domain
func (imco *ClientAPI) CreateRecord(ctx context.Context, domainID string, recordParams *map[string]interface{},
) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, fmt.Sprintf(pathNetworkDnsDomainRecords, domainID), recordParams, true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// UpdateRecord updates a record by its ID
func (imco *ClientAPI) UpdateRecord(ctx context.Context, recordID string, recordParams *map[string]interface{},
) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathNetworkDnsRecord, recordID), recordParams, true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// DeleteRecord deletes a record by its ID
func (imco *ClientAPI) DeleteRecord(ctx context.Context, recordID string) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathNetworkDnsRecord, recordID), true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// RetryRecord retries a record by its ID
func (imco *ClientAPI) RetryRecord(ctx context.Context, recordID string) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	recordParams := new(map[string]interface{})
	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathNetworkDnsRecordRetry, recordID), recordParams, true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}
