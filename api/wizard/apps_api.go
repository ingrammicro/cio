package wizard

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
)

// AppService manages app operations
type AppService struct {
	concertoService utils.ConcertoService
}

// NewAppService returns a Concerto app service
func NewAppService(concertoService utils.ConcertoService) (*AppService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &AppService{
		concertoService: concertoService,
	}, nil
}

// GetAppList returns the list of apps as an array of App
func (as *AppService) GetAppList() (apps []*types.WizardApp, err error) {
	log.Debug("GetAppList")

	data, status, err := as.concertoService.Get("/wizard/apps")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &apps); err != nil {
		return nil, err
	}

	return apps, nil
}

// DeployApp deploys a app
func (as *AppService) DeployApp(appVector *map[string]interface{}, ID string) (app *types.Server, err error) {
	log.Debug("DeployApp")

	data, status, err := as.concertoService.Post(fmt.Sprintf("/wizard/apps/%s/deploy", ID), appVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &app); err != nil {
		return nil, err
	}

	return app, nil
}
