// Copyright (c) 2017-2021 Ingram Micro Inc.

package wizard

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathWizardApps = "/wizard/apps"
const APIPathWizardAppDeploy = "/wizard/apps/%s/deploy"

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

// ListApps returns the list of apps as an array of App
func (as *AppService) ListApps() (apps []*types.WizardApp, err error) {
	log.Debug("ListApps")

	data, status, err := as.concertoService.Get(APIPathWizardApps)
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
func (as *AppService) DeployApp(appID string, appParams *map[string]interface{}) (app *types.Server, err error) {
	log.Debug("DeployApp")

	data, status, err := as.concertoService.Post(fmt.Sprintf(APIPathWizardAppDeploy, appID), appParams)

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
