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

// ImportServer imports a brownfield server import candidate
func (ics *ImportCandidateService) ImportServer(serverID string, serverIn *map[string]interface{}) (server *types.Server, err error) {
	log.Debug("ImportServer")

	data, status, err := ics.concertoService.Post(fmt.Sprintf("/brownfield/import_candidates/%s/import", serverID), serverIn)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// ImportVPC imports a brownfield vpc import candidate
func (ics *ImportCandidateService) ImportVPC(vpcID string, vpcIn *map[string]interface{}) (vpc *types.Vpc, err error) {
	log.Debug("ImportVPC")

	data, status, err := ics.concertoService.Post(fmt.Sprintf("/brownfield/vpc_import_candidates/%s/import", vpcID), vpcIn)
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

// ImportFloatingIP imports a brownfield floating ip import candidate
func (ics *ImportCandidateService) ImportFloatingIP(floatingIPID string, floatingIPIn *map[string]interface{}) (floatingIP *types.FloatingIP, err error) {
	log.Debug("ImportFloatingIP")

	data, status, err := ics.concertoService.Post(fmt.Sprintf("/brownfield/floating_ip_import_candidates/%s/import", floatingIPID), floatingIPIn)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &floatingIP); err != nil {
		return nil, err
	}

	return floatingIP, nil
}

// ImportVolume imports a brownfield volume import candidate
func (ics *ImportCandidateService) ImportVolume(volumeID string, volumeIn *map[string]interface{}) (volume *types.Volume, err error) {
	log.Debug("ImportVolume")

	data, status, err := ics.concertoService.Post(fmt.Sprintf("/brownfield/volume_import_candidates/%s/import", volumeID), volumeIn)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volume); err != nil {
		return nil, err
	}

	return volume, nil
}
