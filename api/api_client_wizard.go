// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListApps returns the list of apps as an array of App
func (imco *ClientAPI) ListApps(ctx context.Context) (apps []*types.WizardApp, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathWizardApps, true, &apps)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

// DeployApp deploys a app
func (imco *ClientAPI) DeployApp(ctx context.Context, appID string, appParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, fmt.Sprintf(PathWizardAppDeploy, appID), appParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// ListLocations returns the list of locations as an array of Location
func (imco *ClientAPI) ListLocations(ctx context.Context) (locations []*types.Location, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathWizardLocations, true, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// ListWizardCloudProviders returns the list of cloud providers as an array of CloudProvider
func (imco *ClientAPI) ListWizardCloudProviders(ctx context.Context, appID string, locationID string,
) (cloudProviders []*types.CloudProvider, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathWizardCloudProviders, appID, locationID), true, &cloudProviders)
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
		fmt.Sprintf(PathWizardServerPlans, appID, locationID, cloudProviderID),
		true,
		&serverPlans,
	)
	if err != nil {
		return nil, err
	}
	return serverPlans, nil
}
