package clientbrownfield

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// BrownfieldCloudAccountService manages brownfield cloud account operations
type BrownfieldCloudAccountService struct {
	concertoService utils.ConcertoService
}

// NewBrownfieldCloudAccountService returns a Concerto BrownfieldCloudAccount service
func NewBrownfieldCloudAccountService(concertoService utils.ConcertoService) (*BrownfieldCloudAccountService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &BrownfieldCloudAccountService{
		concertoService: concertoService,
	}, nil
}

// ListBrownfieldCloudAccounts returns the list of Brownfield Cloud Accounts as an array of CloudAccount
func (bcas *BrownfieldCloudAccountService) ListBrownfieldCloudAccounts() (cloudAccounts []*types.CloudAccount, err error) {
	log.Debug("ListBrownfieldCloudAccounts")

	data, status, err := bcas.concertoService.Get("/brownfield/cloud_accounts")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccounts); err != nil {
		return nil, err
	}

	return cloudAccounts, nil
}

// GetBrownfieldCloudAccount returns a Brownfield Cloud Account by its ID
func (bcas *BrownfieldCloudAccountService) GetBrownfieldCloudAccount(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("GetBrownfieldCloudAccount")

	data, status, err := bcas.concertoService.Get(fmt.Sprintf("/brownfield/cloud_accounts/%s", cloudAccountID))
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

// DiscoverServers discovers brownfield servers
func (bcas *BrownfieldCloudAccountService) DiscoverServers(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("DiscoverServers")

	serversIn := new(map[string]interface{})
	data, status, err := bcas.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/discover", cloudAccountID), serversIn)

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

// ListServers lists brownfield servers import candidates
func (bcas *BrownfieldCloudAccountService) ListServers(cloudAccountID string) (serversImportCandidates []*types.ServerImportCandidate, err error) {
	log.Debug("ListServers")

	data, status, err := bcas.concertoService.Get(fmt.Sprintf("/brownfield/cloud_accounts/%s/import_candidates", cloudAccountID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serversImportCandidates); err != nil {
		return nil, err
	}

	return serversImportCandidates, nil
}

// DiscoverVPCs discovers brownfield VPCs
func (bcas *BrownfieldCloudAccountService) DiscoverVPCs(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("DiscoverVPCs")

	vpcsIn := new(map[string]interface{})
	data, status, err := bcas.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_vpcs", cloudAccountID), vpcsIn)

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

// ListVPCs lists brownfield VPCs import candidates
func (bcas *BrownfieldCloudAccountService) ListVPCs(cloudAccountID string) (vpcsImportCandidates []*types.VpcImportCandidate, err error) {
	log.Debug("ListVPCs")

	data, status, err := bcas.concertoService.Get(fmt.Sprintf("/brownfield/cloud_accounts/%s/vpc_import_candidates", cloudAccountID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &vpcsImportCandidates); err != nil {
		return nil, err
	}

	return vpcsImportCandidates, nil
}

// DiscoverFloatingIPs discovers brownfield floating IPs
func (bcas *BrownfieldCloudAccountService) DiscoverFloatingIPs(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("DiscoverFloatingIPs")

	floatingIPsIn := new(map[string]interface{})
	data, status, err := bcas.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_floating_ips", cloudAccountID), floatingIPsIn)

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

// ListFloatingIPs lists brownfield floating IPs import candidates
func (bcas *BrownfieldCloudAccountService) ListFloatingIPs(cloudAccountID string) (floatingIPsImportCandidates []*types.FloatingIPImportCandidate, err error) {
	log.Debug("ListFloatingIPs")

	data, status, err := bcas.concertoService.Get(fmt.Sprintf("/brownfield/cloud_accounts/%s/floating_ip_import_candidates", cloudAccountID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIPsImportCandidates); err != nil {
		return nil, err
	}

	return floatingIPsImportCandidates, nil
}

// DiscoverVolumes discovers brownfield Volumes
func (bcas *BrownfieldCloudAccountService) DiscoverVolumes(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("DiscoverVolumes")

	volumesIn := new(map[string]interface{})
	data, status, err := bcas.concertoService.Put(fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_volumes", cloudAccountID), volumesIn)

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

// ListVolumes lists brownfield volumes import candidates
func (bcas *BrownfieldCloudAccountService) ListVolumes(cloudAccountID string) (volumesImportCandidates []*types.VolumeImportCandidate, err error) {
	log.Debug("ListVolumes")

	data, status, err := bcas.concertoService.Get(fmt.Sprintf("/brownfield/cloud_accounts/%s/volume_import_candidates", cloudAccountID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volumesImportCandidates); err != nil {
		return nil, err
	}

	return volumesImportCandidates, nil
}