package clientbrownfield

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// ImportCandidateService manages brownfield import candidate operations
type ImportCandidateService struct {
	concertoService utils.ConcertoService
}

// NewImportCandidateService returns a Concerto ImportCandidate service
func NewImportCandidateService(concertoService utils.ConcertoService) (*ImportCandidateService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ImportCandidateService{
		concertoService: concertoService,
	}, nil
}

// ImportServers imports brownfield servers candidates
func (ics *ImportCandidateService) ImportServers(cloudAccountID string, params *map[string]interface{}) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportServers")

	data, status, err := ics.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/import_servers", cloudAccountID), params)
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

// ImportVPCs imports brownfield vpcs candidates
func (ics *ImportCandidateService) ImportVPCs(cloudAccountID string, params *map[string]interface{}) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportVPCs")

	data, status, err := ics.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/import_vpcs", cloudAccountID), params)
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

// ImportFloatingIPs imports brownfield floating ips candidates
func (ics *ImportCandidateService) ImportFloatingIPs(cloudAccountID string, params *map[string]interface{}) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportFloatingIPs")

	data, status, err := ics.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/import_floating_ips", cloudAccountID), params)
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

// ImportVolumes imports brownfield volumes candidates
func (ics *ImportCandidateService) ImportVolumes(cloudAccountID string, params *map[string]interface{}) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportVolumes")

	data, status, err := ics.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/import_volumes", cloudAccountID), params)
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

// ImportKubernetesClusters imports brownfield kubernetes clusters candidates
func (ics *ImportCandidateService) ImportKubernetesClusters(cloudAccountID string, params *map[string]interface{}) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportKubernetesClusters")

	data, status, err := ics.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/import_kubernetes_clusters", cloudAccountID), params)
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

// ImportPolicies imports brownfield policies candidates
func (ics *ImportCandidateService) ImportPolicies(cloudAccountID string, params *map[string]interface{}) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("ImportPolicies")

	data, status, err := ics.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/import_policies", cloudAccountID), params)
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
