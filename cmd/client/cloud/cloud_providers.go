// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fCloudProviderId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudProviderId, Required: true,
		Usage: "Cloud provider id"}

	cloudProvidersCmd := cmd.NewCommand(CloudCmd, &cmd.CommandContext{
		Use:   "cloud-providers",
		Short: "Provides information on cloud providers"},
	)
	cmd.NewCommand(cloudProvidersCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists all available cloud providers",
		RunMethod: CloudProviderList},
	)
	cmd.NewCommand(cloudProvidersCmd, &cmd.CommandContext{
		Use:          "list-storage-plans",
		Short:        "This action lists the storage plans offered by the cloud provider identified by the given id",
		RunMethod:    CloudProviderStoragePlansList,
		FlagContexts: []cmd.FlagContext{fCloudProviderId}},
	)
	cmd.NewCommand(cloudProvidersCmd, &cmd.CommandContext{
		Use:          "list-load-balancer-plans",
		Short:        "This action lists the load balancer plans offered by the cloud provider identified by the given id",
		RunMethod:    CloudProviderLoadBalancerPlansList,
		FlagContexts: []cmd.FlagContext{fCloudProviderId}},
	)
	cmd.NewCommand(cloudProvidersCmd, &cmd.CommandContext{
		Use:          "list-cluster-plans",
		Short:        "This action lists the cluster plans offered by the cloud provider identified by the given id",
		RunMethod:    CloudProviderClusterPlansList,
		FlagContexts: []cmd.FlagContext{fCloudProviderId}},
	)
}

// CloudProviderList subcommand function
func CloudProviderList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	cloudProviders, err := svc.ListCloudProviders()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	if err = formatter.PrintList(cloudProviders); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudProviderStoragePlansList subcommand function
func CloudProviderStoragePlansList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	storagePlans, err := svc.ListServerStoragePlans(viper.GetString(cmd.CloudProviderId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive storage plans data", err)
	}

	cloudProvidersMap := cmd.LoadCloudProvidersMapping()
	locationsMap := cmd.LoadLocationsMapping()

	for id, sp := range storagePlans {
		storagePlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
		storagePlans[id].LocationName = locationsMap[sp.LocationID]
	}

	if err = formatter.PrintList(storagePlans); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudProviderLoadBalancerPlansList subcommand function
func CloudProviderLoadBalancerPlansList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	loadBalancerPlans, err := svc.ListLoadBalancerPlans(viper.GetString(cmd.CloudProviderId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer plans data", err)
	}

	cloudProvidersMap := cmd.LoadCloudProvidersMapping()
	for id, sp := range loadBalancerPlans {
		loadBalancerPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
	}

	if err = formatter.PrintList(loadBalancerPlans); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudProviderClusterPlansList subcommand function
func CloudProviderClusterPlansList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	loadBalancerPlans, err := svc.ListClusterPlans(viper.GetString(cmd.CloudProviderId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cluster plans data", err)
	}

	cloudProvidersMap := cmd.LoadCloudProvidersMapping()
	for id, sp := range loadBalancerPlans {
		loadBalancerPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
	}

	if err = formatter.PrintList(loadBalancerPlans); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
