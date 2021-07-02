// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathNetworkVpcVpn = "/network/vpcs/%s/vpn"
const APIPathNetworkVpcVpnPlans = "/network/vpcs/%s/vpn_plans"

// VPNService manages VPN operations
type VPNService struct {
	concertoService utils.ConcertoService
}

// NewVPNService returns a Concerto VPN service
func NewVPNService(concertoService utils.ConcertoService) (*VPNService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &VPNService{
		concertoService: concertoService,
	}, nil
}

// GetVPN returns a VPN by VPC ID
func (vs *VPNService) GetVPN(vpcID string) (vpn *types.Vpn, err error) {
	log.Debug("GetVPN")

	data, status, err := vs.concertoService.Get(fmt.Sprintf(APIPathNetworkVpcVpn, vpcID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpn); err != nil {
		return nil, err
	}

	return vpn, nil
}

// CreateVPN creates a VPN
func (vs *VPNService) CreateVPN(vpcID string, vpnParams *map[string]interface{}) (vpn *types.Vpn, err error) {
	log.Debug("CreateVPN")

	data, status, err := vs.concertoService.Post(fmt.Sprintf(APIPathNetworkVpcVpn, vpcID), vpnParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpn); err != nil {
		return nil, err
	}

	return vpn, nil
}

// DeleteVPN deletes VPN by VPC ID
func (vs *VPNService) DeleteVPN(vpcID string) (err error) {
	log.Debug("DeleteVPN")

	data, status, err := vs.concertoService.Delete(fmt.Sprintf(APIPathNetworkVpcVpn, vpcID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ListVPNPlans returns the list of VPN plans for a given VPC ID
func (vs *VPNService) ListVPNPlans(vpcID string) (vpnPlans []*types.VpnPlan, err error) {
	log.Debug("ListVPNPlans")

	data, status, err := vs.concertoService.Get(fmt.Sprintf(APIPathNetworkVpcVpnPlans, vpcID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpnPlans); err != nil {
		return nil, err
	}

	return vpnPlans, nil
}
