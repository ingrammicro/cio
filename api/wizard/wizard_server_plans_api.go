package wizard

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// WizardServerPlanService manages wizard server plan operations
type WizardServerPlanService struct {
	concertoService utils.ConcertoService
}

// NewWizardServerPlanService returns a Concerto serverPlan service
func NewWizardServerPlanService(concertoService utils.ConcertoService) (*WizardServerPlanService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &WizardServerPlanService{
		concertoService: concertoService,
	}, nil
}

// ListWizardServerPlans returns the list of server plans as an array of ServerPlan
func (wsps *WizardServerPlanService) ListWizardServerPlans(appID string, locationID string, cloudProviderID string) (serverPlans []*types.ServerPlan, err error) {
	log.Debug("ListWizardServerPlans")

	data, status, err := wsps.concertoService.Get(fmt.Sprintf("/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s", appID, locationID, cloudProviderID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverPlans); err != nil {
		return nil, err
	}

	return serverPlans, nil
}
