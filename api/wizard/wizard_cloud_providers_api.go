package wizard

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// WizardCloudProviderService manages wizard cloud provider operations
type WizardCloudProviderService struct {
	concertoService utils.ConcertoService
}

// NewWizardCloudProviderService returns a Concerto WizardCloudProvider service
func NewWizardCloudProviderService(concertoService utils.ConcertoService) (*WizardCloudProviderService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &WizardCloudProviderService{
		concertoService: concertoService,
	}, nil
}

// ListWizardCloudProviders returns the list of cloud providers as an array of CloudProvider
func (wcps *WizardCloudProviderService) ListWizardCloudProviders(appID string, locationID string) (cloudProviders []*types.CloudProvider, err error) {
	log.Debug("ListWizardCloudProviders")

	data, status, err := wcps.concertoService.Get(fmt.Sprintf("/wizard/cloud_providers?app_id=%s&location_id=%s", appID, locationID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudProviders); err != nil {
		return nil, err
	}

	return cloudProviders, nil
}
