// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathNetworkVpcs = "/network/vpcs"
const APIPathNetworkVpc = "/network/vpcs/%s"
const APIPathNetworkVpcDiscard = "/network/vpcs/%s/discard"

// VPCService manages VPC operations
type VPCService struct {
	concertoService utils.ConcertoService
}

// NewVPCService returns a Concerto VPC service
func NewVPCService(concertoService utils.ConcertoService) (*VPCService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &VPCService{
		concertoService: concertoService,
	}, nil
}

// ListVPCs returns the list of VPCs as an array of VPC
func (vs *VPCService) ListVPCs() (vpcs []*types.Vpc, err error) {
	log.Debug("ListVPCs")

	data, status, err := vs.concertoService.Get(APIPathNetworkVpcs)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpcs); err != nil {
		return nil, err
	}

	return vpcs, nil
}

// GetVPC returns a VPC by its ID
func (vs *VPCService) GetVPC(vpcID string) (vpc *types.Vpc, err error) {
	log.Debug("GetVPC")

	data, status, err := vs.concertoService.Get(fmt.Sprintf(APIPathNetworkVpc, vpcID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpc); err != nil {
		return nil, err
	}

	return vpc, nil
}

// CreateVPC creates a VPC
func (vs *VPCService) CreateVPC(vpcParams *map[string]interface{}) (vpc *types.Vpc, err error) {
	log.Debug("CreateVPC")

	data, status, err := vs.concertoService.Post(APIPathNetworkVpcs, vpcParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpc); err != nil {
		return nil, err
	}

	return vpc, nil
}

// UpdateVPC updates a VPC by its ID
func (vs *VPCService) UpdateVPC(vpcID string, vpcParams *map[string]interface{}) (vpc *types.Vpc, err error) {
	log.Debug("UpdateVPC")

	data, status, err := vs.concertoService.Put(fmt.Sprintf(APIPathNetworkVpc, vpcID), vpcParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpc); err != nil {
		return nil, err
	}

	return vpc, nil
}

// DeleteVPC deletes a VPC by its ID
func (vs *VPCService) DeleteVPC(vpcID string) (err error) {
	log.Debug("DeleteVPC")

	data, status, err := vs.concertoService.Delete(fmt.Sprintf(APIPathNetworkVpc, vpcID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DiscardVPC discards a VPC by its ID
func (vs *VPCService) DiscardVPC(vpcID string) (err error) {
	log.Debug("DiscardVPC")

	data, status, err := vs.concertoService.Delete(fmt.Sprintf(APIPathNetworkVpcDiscard, vpcID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
