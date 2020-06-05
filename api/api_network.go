// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListCertificates returns the list of certificates in a load balancer by its ID, as an array of Certificate
func (imco *IMCOClient) ListCertificates(loadBalancerID string) (certificates []*types.Certificate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkLoadBalancerCertificates, loadBalancerID), true, &certificates)
	if err != nil {
		return nil, err
	}
	return certificates, nil
}

// GetCertificate returns a certificate by its ID
func (imco *IMCOClient) GetCertificate(loadBalancerID string, certificateID string,
) (certificate *types.Certificate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
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
func (imco *IMCOClient) CreateCertificate(loadBalancerID string, certificateParams *map[string]interface{},
) (certificate *types.Certificate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
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
func (imco *IMCOClient) UpdateCertificate(loadBalancerID string, certificateID string,
	certificateParams *map[string]interface{},
) (certificate *types.Certificate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) DeleteCertificate(loadBalancerID string, certificateID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(
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
func (imco *IMCOClient) ListFirewallProfiles() (firewallProfiles []*types.FirewallProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathNetworkFirewallProfiles, true, &firewallProfiles)
	if err != nil {
		return nil, err
	}
	return firewallProfiles, nil
}

// GetFirewallProfile returns a firewallProfile by its ID
func (imco *IMCOClient) GetFirewallProfile(firewallProfileID string,
) (firewallProfile *types.FirewallProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkFirewallProfile, firewallProfileID), true, &firewallProfile)
	if err != nil {
		return nil, err
	}
	return firewallProfile, nil
}

// CreateFirewallProfile creates a firewallProfile
func (imco *IMCOClient) CreateFirewallProfile(firewallProfileParams *map[string]interface{},
) (firewallProfile *types.FirewallProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathNetworkFirewallProfiles, firewallProfileParams, true, &firewallProfile)
	if err != nil {
		return nil, err
	}
	return firewallProfile, nil
}

// UpdateFirewallProfile updates a firewallProfile by its ID
func (imco *IMCOClient) UpdateFirewallProfile(firewallProfileID string, firewallProfileParams *map[string]interface{},
) (firewallProfile *types.FirewallProfile, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) DeleteFirewallProfile(firewallProfileID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkFirewallProfile, firewallProfileID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetFloatingIP returns a FloatingIP by its ID
func (imco *IMCOClient) GetFloatingIP(floatingIPID string) (floatingIP *types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkFloatingIp, floatingIPID), true, &floatingIP)
	if err != nil {
		return nil, err
	}
	return floatingIP, nil
}

// CreateFloatingIP creates a FloatingIP
func (imco *IMCOClient) CreateFloatingIP(floatingIPParams *map[string]interface{},
) (floatingIP *types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathNetworkFloatingIps, floatingIPParams, true, &floatingIP)
	if err != nil {
		return nil, err
	}
	return floatingIP, nil
}

// UpdateFloatingIP updates a FloatingIP by its ID
func (imco *IMCOClient) UpdateFloatingIP(floatingIPID string, floatingIPParams *map[string]interface{},
) (floatingIP *types.FloatingIP, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkFloatingIp, floatingIPID), floatingIPParams, true, &floatingIP)
	if err != nil {
		return nil, err
	}
	return floatingIP, nil
}

// AttachFloatingIP attaches a FloatingIP by its ID
func (imco *IMCOClient) AttachFloatingIP(floatingIPID string, floatingIPParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
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
func (imco *IMCOClient) DetachFloatingIP(floatingIPID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkFloatingIpAttachedServer, floatingIPID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFloatingIP deletes a FloatingIP by its ID
func (imco *IMCOClient) DeleteFloatingIP(floatingIPID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkFloatingIp, floatingIPID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DiscardFloatingIP discards a FloatingIP by its ID
func (imco *IMCOClient) DiscardFloatingIP(floatingIPID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkFloatingIpDiscard, floatingIPID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListListeners returns the list of listeners in a load balancer by its ID, as an array of Listener
func (imco *IMCOClient) ListListeners(loadBalancerID string) (listeners []*types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkLoadBalancerListeners, loadBalancerID), true, &listeners)
	if err != nil {
		return nil, err
	}
	return listeners, nil
}

// GetListener returns a listener by its ID
func (imco *IMCOClient) GetListener(listenerID string) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkListener, listenerID), true, &listener)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// CreateListener creates a listener in a load balancer by its ID
func (imco *IMCOClient) CreateListener(loadBalancerID string, listenerParams *map[string]interface{},
) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
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
func (imco *IMCOClient) UpdateListener(listenerID string, listenerParams *map[string]interface{},
) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkListener, listenerID), listenerParams, true, &listener)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// DeleteListener deletes a listener by its ID
func (imco *IMCOClient) DeleteListener(listenerID string) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkListener, listenerID), true, nil)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// RetryListener retries a listener by its ID
func (imco *IMCOClient) RetryListener(listenerID string, listenerParams *map[string]interface{},
) (listener *types.Listener, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkListenerRetry, listenerID), listenerParams, true, &listener)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

// ListRules returns the list of rules in a listener by its ID, as an array of ListenerRule
func (imco *IMCOClient) ListRules(listenerID string) (listenerRules []*types.ListenerRule, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkListenerRules, listenerID), true, &listenerRules)
	if err != nil {
		return nil, err
	}
	return listenerRules, nil
}

// CreateRule creates a rule in a listener by its ID
func (imco *IMCOClient) CreateRule(listenerID string, listenerRuleParams *map[string]interface{},
) (listenerRule *types.ListenerRule, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
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
func (imco *IMCOClient) UpdateRule(listenerID string, listenerRuleID string,
	listenerRuleParams *map[string]interface{},
) (listenerRule *types.ListenerRule, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) DeleteRule(listenerID string, listenerRuleID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkListenerRule, listenerID, listenerRuleID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListLoadBalancers returns the list of load balancers as an array of LoadBalancer
func (imco *IMCOClient) ListLoadBalancers() (loadBalancers []*types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathNetworkLoadBalancers, true, &loadBalancers)
	if err != nil {
		return nil, err
	}
	return loadBalancers, nil
}

// GetLoadBalancer returns a load balancer by its ID
func (imco *IMCOClient) GetLoadBalancer(loadBalancerID string) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkLoadBalancer, loadBalancerID), true, &loadBalancer)
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

// CreateLoadBalancer creates a load balancer
func (imco *IMCOClient) CreateLoadBalancer(loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathNetworkLoadBalancers, loadBalancerParams, true, &loadBalancer)
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

// UpdateLoadBalancer updates a load balancer by its ID
func (imco *IMCOClient) UpdateLoadBalancer(loadBalancerID string, loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) DeleteLoadBalancer(loadBalancerID string) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkLoadBalancer, loadBalancerID), true, &loadBalancer)
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}

// RetryLoadBalancer retries a load balancer by its ID
func (imco *IMCOClient) RetryLoadBalancer(loadBalancerID string, loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) GetLoadBalancerPlan(loadBalancerPlanID string,
) (loadBalancerPlan *types.LoadBalancerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkLoadBalancerPlan, loadBalancerPlanID), true, &loadBalancerPlan)
	if err != nil {
		return nil, err
	}
	return loadBalancerPlan, nil
}

// ListSubnets returns the list of Subnets of a VPC as an array of Subnet
func (imco *IMCOClient) ListSubnets(vpcID string) (subnets []*types.Subnet, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkVpcSubnets, vpcID), true, &subnets)
	if err != nil {
		return nil, err
	}
	return subnets, nil
}

// GetSubnet returns a Subnet by its ID
func (imco *IMCOClient) GetSubnet(subnetID string) (subnet *types.Subnet, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkSubnet, subnetID), true, &subnet)
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

// CreateSubnet creates a Subnet
func (imco *IMCOClient) CreateSubnet(vpcID string, subnetParams *map[string]interface{},
) (subnet *types.Subnet, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathNetworkVpcSubnets, vpcID), subnetParams, true, &subnet)
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

// UpdateSubnet updates a Subnet by its ID
func (imco *IMCOClient) UpdateSubnet(subnetID string, subnetParams *map[string]interface{},
) (subnet *types.Subnet, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkSubnet, subnetID), subnetParams, true, &subnet)
	if err != nil {
		return nil, err
	}
	return subnet, nil
}

// DeleteSubnet deletes a Subnet by its ID
func (imco *IMCOClient) DeleteSubnet(subnetID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkSubnet, subnetID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListSubnetServers returns the list of Servers of a Subnet as an array of server
func (imco *IMCOClient) ListSubnetServers(subnetID string) (servers []*types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkSubnetServers, subnetID), true, &servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

// ListSubnetServerArrays returns the list of server arrays of a Subnet as an array of ServerArray
func (imco *IMCOClient) ListSubnetServerArrays(subnetID string) (serverArrays []*types.ServerArray, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkSubnetServerArrays, subnetID), true, &serverArrays)
	if err != nil {
		return nil, err
	}
	return serverArrays, nil
}

// ListTargetGroups returns the list of target groups in a load balancer by its ID, as an array of TargetGroup
func (imco *IMCOClient) ListTargetGroups(loadBalancerID string) (targetGroups []*types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkLoadBalancerTargetGroups, loadBalancerID), true, &targetGroups)
	if err != nil {
		return nil, err
	}
	return targetGroups, nil
}

// GetTargetGroup returns a target group by its ID
func (imco *IMCOClient) GetTargetGroup(targetGroupID string) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkTargetGroup, targetGroupID), true, &targetGroup)
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

// CreateTargetGroup creates a target group in a load balancer by its ID
func (imco *IMCOClient) CreateTargetGroup(loadBalancerID string, targetGroupParams *map[string]interface{},
) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
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
func (imco *IMCOClient) UpdateTargetGroup(targetGroupID string, targetGroupParams *map[string]interface{},
) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkTargetGroup, targetGroupID), targetGroupParams, true, &targetGroup)
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

// DeleteTargetGroup deletes a target group by its ID
func (imco *IMCOClient) DeleteTargetGroup(targetGroupID string) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkTargetGroup, targetGroupID), true, nil)
	if err != nil {
		return nil, err
	}
	return targetGroup, nil
}

// RetryTargetGroup retries a target group by its ID
func (imco *IMCOClient) RetryTargetGroup(targetGroupID string, targetGroupParams *map[string]interface{},
) (targetGroup *types.TargetGroup, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) ListTargets(targetGroupID string) (targets []*types.Target, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkTargetGroupTargets, targetGroupID), true, &targets)
	if err != nil {
		return nil, err
	}
	return targets, nil
}

// CreateTarget creates a target in a target group by its ID
func (imco *IMCOClient) CreateTarget(targetGroupID string, targetParams *map[string]interface{},
) (target *types.Target, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathNetworkTargetGroupTargets, targetGroupID), targetParams, true, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

// DeleteTarget deletes a target in a target group by given IDs and resource type
func (imco *IMCOClient) DeleteTarget(targetGroupID string, targetResourceType string, targetResourceID string,
) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(
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
func (imco *IMCOClient) ListVPCs() (vpcs []*types.Vpc, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathNetworkVpcs, true, &vpcs)
	if err != nil {
		return nil, err
	}
	return vpcs, nil
}

// GetVPC returns a VPC by its ID
func (imco *IMCOClient) GetVPC(vpcID string) (vpc *types.Vpc, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkVpc, vpcID), true, &vpc)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

// CreateVPC creates a VPC
func (imco *IMCOClient) CreateVPC(vpcParams *map[string]interface{}) (vpc *types.Vpc, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathNetworkVpcs, vpcParams, true, &vpc)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

// UpdateVPC updates a VPC by its ID
func (imco *IMCOClient) UpdateVPC(vpcID string, vpcParams *map[string]interface{}) (vpc *types.Vpc, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkVpc, vpcID), vpcParams, true, &vpc)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}

// DeleteVPC deletes a VPC by its ID
func (imco *IMCOClient) DeleteVPC(vpcID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkVpc, vpcID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DiscardVPC discards a VPC by its ID
func (imco *IMCOClient) DiscardVPC(vpcID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkVpcDiscard, vpcID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetVPN returns a VPN by VPC ID
func (imco *IMCOClient) GetVPN(vpcID string) (vpn *types.Vpn, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkVpcVpn, vpcID), true, &vpn)
	if err != nil {
		return nil, err
	}
	return vpn, nil
}

// CreateVPN creates a VPN
func (imco *IMCOClient) CreateVPN(vpcID string, vpnParams *map[string]interface{}) (vpn *types.Vpn, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathNetworkVpcVpn, vpcID), vpnParams, true, &vpn)
	if err != nil {
		return nil, err
	}
	return vpn, nil
}

// DeleteVPN deletes VPN by VPC ID
func (imco *IMCOClient) DeleteVPN(vpcID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkVpcVpn, vpcID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListVPNPlans returns the list of VPN plans for a given VPC ID
func (imco *IMCOClient) ListVPNPlans(vpcID string) (vpnPlans []*types.VpnPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkVpcVpnPlans, vpcID), true, vpnPlans)
	if err != nil {
		return nil, err
	}
	return vpnPlans, nil
}

// ListDomains returns the list of domains as an array of Domain
func (imco *IMCOClient) ListDomains() (domains []*types.Domain, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathNetworkDnsDomains, true, &domains)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

// GetDomain returns a domain by its ID
func (imco *IMCOClient) GetDomain(domainID string) (domain *types.Domain, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkDnsDomain, domainID), true, &domain)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

// CreateDomain creates a domain
func (imco *IMCOClient) CreateDomain(domainParams *map[string]interface{}) (domain *types.Domain, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathNetworkDnsDomains, domainParams, true, &domain)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

// DeleteDomain deletes a domain by its ID
func (imco *IMCOClient) DeleteDomain(domainID string) (domain *types.Domain, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkDnsDomain, domainID), true, nil)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

// RetryDomain retries a domain by its ID
func (imco *IMCOClient) RetryDomain(domainID string) (domain *types.Domain, err error) {
	logger.DebugFuncInfo()

	domainParams := new(map[string]interface{})
	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkDnsDomainRetry, domainID), domainParams, true, &domain)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

// ListRecords returns the list of records as an array of Record for given domain
func (imco *IMCOClient) ListRecords(domainID string) (records []*types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkDnsDomainRecords, domainID), true, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetRecord returns a record by its ID
func (imco *IMCOClient) GetRecord(recordID string) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathNetworkDnsRecord, recordID), true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// CreateRecord creates a record in a domain
func (imco *IMCOClient) CreateRecord(domainID string, recordParams *map[string]interface{},
) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathNetworkDnsDomainRecords, domainID), recordParams, true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// UpdateRecord updates a record by its ID
func (imco *IMCOClient) UpdateRecord(recordID string, recordParams *map[string]interface{},
) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkDnsRecord, recordID), recordParams, true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// DeleteRecord deletes a record by its ID
func (imco *IMCOClient) DeleteRecord(recordID string) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathNetworkDnsRecord, recordID), true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// RetryRecord retries a record by its ID
func (imco *IMCOClient) RetryRecord(recordID string) (record *types.Record, err error) {
	logger.DebugFuncInfo()

	recordParams := new(map[string]interface{})
	_, err = imco.putAndCheck(fmt.Sprintf(pathNetworkDnsRecordRetry, recordID), recordParams, true, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}
