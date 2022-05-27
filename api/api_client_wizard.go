// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"golang.org/x/net/context"
)

// ListApps returns the list of apps as an array of App
func (imco *ClientAPI) ListApps(ctx context.Context) (apps []*types.WizardApp, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathWizardApps, true, &apps)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

// DeployApp deploys a app
func (imco *ClientAPI) DeployApp(ctx context.Context, appID string, appParams *map[string]interface{},
) (app *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, fmt.Sprintf(pathWizardAppDeploy, appID), appParams, true, &app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// ListLocations returns the list of locations as an array of Location
func (imco *ClientAPI) ListLocations(ctx context.Context) (locations []*types.Location, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathWizardLocations, true, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// ListWizardCloudProviders returns the list of cloud providers as an array of CloudProvider
func (imco *ClientAPI) ListWizardCloudProviders(ctx context.Context, appID string, locationID string,
) (cloudProviders []*types.CloudProvider, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathWizardCloudProviders, appID, locationID), true, &cloudProviders)
	if err != nil {
		return nil, err
	}
	return cloudProviders, nil
}

// ListWizardServerPlans returns the list of server plans as an array of ServerPlan
func (imco *ClientAPI) ListWizardServerPlans(ctx context.Context, appID string, locationID string,
	cloudProviderID string,
) (serverPlans []*types.ServerPlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(pathWizardServerPlans, appID, locationID, cloudProviderID),
		true,
		&serverPlans,
	)
	if err != nil {
		return nil, err
	}
	return serverPlans, nil
}
