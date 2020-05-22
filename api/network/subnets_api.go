package network

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// SubnetService manages subnet operations
type SubnetService struct {
	concertoService utils.ConcertoService
}

// NewSubnetService returns a Concerto Subnet service
func NewSubnetService(concertoService utils.ConcertoService) (*SubnetService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &SubnetService{
		concertoService: concertoService,
	}, nil
}

// ListSubnets returns the list of Subnets of a VPC as an array of Subnet
func (ss *SubnetService) ListSubnets(vpcID string) (subnets []*types.Subnet, err error) {
	log.Debug("ListSubnets")

	data, status, err := ss.concertoService.Get(fmt.Sprintf("/network/vpcs/%s/subnets", vpcID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &subnets); err != nil {
		return nil, err
	}

	return subnets, nil
}

// GetSubnet returns a Subnet by its ID
func (ss *SubnetService) GetSubnet(subnetID string) (subnet *types.Subnet, err error) {
	log.Debug("GetSubnet")

	data, status, err := ss.concertoService.Get(fmt.Sprintf("/network/subnets/%s", subnetID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &subnet); err != nil {
		return nil, err
	}

	return subnet, nil
}

// CreateSubnet creates a Subnet
func (ss *SubnetService) CreateSubnet(subnetParams *map[string]interface{}, vpcID string) (subnet *types.Subnet, err error) {
	log.Debug("CreateSubnet")

	data, status, err := ss.concertoService.Post(fmt.Sprintf("/network/vpcs/%s/subnets", vpcID), subnetParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &subnet); err != nil {
		return nil, err
	}

	return subnet, nil
}

// UpdateSubnet updates a Subnet by its ID
func (ss *SubnetService) UpdateSubnet(subnetParams *map[string]interface{}, ID string) (subnet *types.Subnet, err error) {
	log.Debug("UpdateSubnet")

	data, status, err := ss.concertoService.Put(fmt.Sprintf("/network/subnets/%s", ID), subnetParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &subnet); err != nil {
		return nil, err
	}

	return subnet, nil
}

// DeleteSubnet deletes a Subnet by its ID
func (ss *SubnetService) DeleteSubnet(subnetID string) (err error) {
	log.Debug("DeleteSubnet")

	data, status, err := ss.concertoService.Delete(fmt.Sprintf("/network/subnets/%s", subnetID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ListSubnetServers returns the list of Servers of a Subnet as an array of Server
func (ss *SubnetService) ListSubnetServers(subnetID string) (servers []*types.Server, err error) {
	log.Debug("ListSubnetServers")

	data, status, err := ss.concertoService.Get(fmt.Sprintf("/network/subnets/%s/servers", subnetID))

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

// ListSubnetServerArrays returns the list of Server arrays of a Subnet as an array of ServerArray
func (ss *SubnetService) ListSubnetServerArrays(subnetID string) (serverArrays []*types.ServerArray, err error) {
	log.Debug("ListSubnetServerArrays")

	data, status, err := ss.concertoService.Get(fmt.Sprintf("/network/subnets/%s/server_arrays", subnetID))

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
