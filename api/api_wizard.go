// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListApps returns the list of apps as an array of App
func (imco *IMCOClient) ListApps() (apps []*types.WizardApp, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathWizardApps, true, &apps)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

// DeployApp deploys a app
func (imco *IMCOClient) DeployApp(appID string, appParams *map[string]interface{}) (app *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathWizardAppDeploy, appID), appParams, true, &app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// ListLocations returns the list of locations as an array of Location
func (imco *IMCOClient) ListLocations() (locations []*types.Location, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathWizardLocations, true, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// ListWizardCloudProviders returns the list of cloud providers as an array of CloudProvider
func (imco *IMCOClient) ListWizardCloudProviders(appID string, locationID string,
) (cloudProviders []*types.CloudProvider, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathWizardCloudProviders, appID, locationID), true, &cloudProviders)
	if err != nil {
		return nil, err
	}
	return cloudProviders, nil
}

// ListWizardServerPlans returns the list of server plans as an array of ServerPlan
func (imco *IMCOClient) ListWizardServerPlans(appID string, locationID string, cloudProviderID string,
) (serverPlans []*types.ServerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathWizardServerPlans, appID, locationID, cloudProviderID),
		true,
		&serverPlans,
	)
	if err != nil {
		return nil, err
	}
	return serverPlans, nil
}
