// Copyright (c) 2017-2022 Ingram Micro Inc.

package cli

import (
	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/utils/format"
	"golang.org/x/net/context"
)

// WireUpAPIClient prepares common resources to send request to Orchestrator API
func WireUpAPIClient() (svc *api.ClientAPI, config *configuration.Config, f format.Formatter) {
	ds, config, f := cmd.WireUpAPI()
	svc = new(api.ClientAPI)
	svc.HTTPClient = *ds
	return svc, config, f
}

// LoadCloudProvidersMapping retrieves Cloud Providers and create a map between ID and Name
func LoadCloudProvidersMapping(ctx context.Context) map[string]string {
	logger.DebugFuncInfo()
	svc, _, formatter := WireUpAPIClient()

	cloudProviders, err := svc.ListCloudProviders(ctx)
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	cloudProvidersMap := make(map[string]string)
	for _, cloudProvider := range cloudProviders {
		cloudProvidersMap[cloudProvider.ID] = cloudProvider.Name
	}
	return cloudProvidersMap
}

// LoadLocationsMapping retrieves Locations and create a map between ID and Name
func LoadLocationsMapping(ctx context.Context) map[string]string {
	logger.DebugFuncInfo()
	svc, _, formatter := WireUpAPIClient()

	locations, err := svc.ListLocations(ctx)
	if err != nil {
		formatter.PrintFatal("Couldn't receive location data", err)
	}
	locationsMap := make(map[string]string)
	for _, location := range locations {
		locationsMap[location.ID] = location.Name
	}
	return locationsMap
}
