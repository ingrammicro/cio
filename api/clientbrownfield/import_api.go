// Copyright (c) 2017-2021 Ingram Micro Inc.

package clientbrownfield

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathBlueprintCloudAccountImportServers = "/brownfield/cloud_accounts/%s/import_servers"
const APIPathBlueprintCloudAccountImportVpcs = "/brownfield/cloud_accounts/%s/import_vpcs"
const APIPathBlueprintCloudAccountImportFloatingIPs = "/brownfield/cloud_accounts/%s/import_floating_ips"
const APIPathBlueprintCloudAccountImportVolumes = "/brownfield/cloud_accounts/%s/import_volumes"
const APIPathBlueprintCloudAccountImportKubernetesClusters = "/brownfield/cloud_accounts/%s/import_kubernetes_clusters"
const APIPathBlueprintCloudAccountImportPolicies = "/brownfield/cloud_accounts/%s/import_policies"

// ImportService manages brownfield import operations
type ImportService struct {
	concertoService utils.ConcertoService
}

// NewImportService returns a Concerto Import service
func NewImportService(concertoService utils.ConcertoService) (*ImportService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ImportService{
		concertoService: concertoService,
	}, nil
}

// ImportServers imports brownfield servers
func (is *ImportService) ImportServers(
	cloudAccountID string,
	params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportServers")

	data, status, err := is.concertoService.Put(
		fmt.Sprintf(APIPathBlueprintCloudAccountImportServers, cloudAccountID),
		params,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

// ImportVPCs imports brownfield vpcs
func (is *ImportService) ImportVPCs(
	cloudAccountID string,
	params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportVPCs")

	data, status, err := is.concertoService.Put(
		fmt.Sprintf(APIPathBlueprintCloudAccountImportVpcs, cloudAccountID),
		params,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

// ImportFloatingIPs imports brownfield floating ips
func (is *ImportService) ImportFloatingIPs(
	cloudAccountID string,
	params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportFloatingIPs")

	data, status, err := is.concertoService.Put(
		fmt.Sprintf(APIPathBlueprintCloudAccountImportFloatingIPs, cloudAccountID),
		params,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

// ImportVolumes imports brownfield volumes
func (is *ImportService) ImportVolumes(
	cloudAccountID string,
	params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportVolumes")

	data, status, err := is.concertoService.Put(
		fmt.Sprintf(APIPathBlueprintCloudAccountImportVolumes, cloudAccountID),
		params,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

// ImportKubernetesClusters imports brownfield kubernetes clusters
func (is *ImportService) ImportKubernetesClusters(
	cloudAccountID string,
	params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportKubernetesClusters")

	data, status, err := is.concertoService.Put(
		fmt.Sprintf(APIPathBlueprintCloudAccountImportKubernetesClusters, cloudAccountID),
		params,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

// ImportPolicies imports brownfield policies
func (is *ImportService) ImportPolicies(
	cloudAccountID string,
	params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportPolicies")

	data, status, err := is.concertoService.Put(
		fmt.Sprintf(APIPathBlueprintCloudAccountImportPolicies, cloudAccountID),
		params,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}
