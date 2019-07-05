package wizard

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
)

// WizCloudProvidersService manages wizCloudProviders operations
type WizCloudProvidersService struct {
	concertoService utils.ConcertoService
}

// NewWizCloudProvidersService returns a Concerto wizCloudProviders service
func NewWizCloudProvidersService(concertoService utils.ConcertoService) (*WizCloudProvidersService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &WizCloudProvidersService{
		concertoService: concertoService,
	}, nil
}

// GetWizCloudProviderList returns the list of wizCloudProviders as an array of CloudProvider
func (dm *WizCloudProvidersService) GetWizCloudProviderList(AppID string, LocID string) (wizCloudProviders []*types.CloudProvider, err error) {
	log.Debug("GetWizCloudProviderList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/wizard/cloud_providers?app_id=%s&location_id=%s", AppID, LocID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &wizCloudProviders); err != nil {
		return nil, err
	}

	return wizCloudProviders, nil
}
