// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloud

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fCloudProviderId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudProviderId, Required: true,
		Usage: "Cloud provider id"}

	providersCmd := cmd.NewCommand(cloudCmd, &cmd.CommandContext{
		Use:   "providers",
		Short: "Provides information on cloud providers"},
	)
	cmd.NewCommand(providersCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists all available cloud providers",
		RunMethod: CloudProviderList},
	)
	cmd.NewCommand(providersCmd, &cmd.CommandContext{
		Use:          "list-storage-plans",
		Short:        "This action lists the storage plans offered by the cloud provider identified by the given id",
		RunMethod:    CloudProviderStoragePlansList,
		FlagContexts: []cmd.FlagContext{fCloudProviderId}},
	)
	cmd.NewCommand(providersCmd, &cmd.CommandContext{
		Use:          "list-load-balancer-plans",
		Short:        "This action lists the load balancer plans offered by the cloud provider identified by the given id",
		RunMethod:    CloudProviderLoadBalancerPlansList,
		FlagContexts: []cmd.FlagContext{fCloudProviderId}},
	)
	cmd.NewCommand(providersCmd, &cmd.CommandContext{
		Use:          "list-cluster-plans",
		Short:        "This action lists the cluster plans offered by the cloud provider identified by the given id",
		RunMethod:    CloudProviderClusterPlansList,
		FlagContexts: []cmd.FlagContext{fCloudProviderId}},
	)
}

// CloudProviderList subcommand function
func CloudProviderList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cloudProviders, err := svc.ListCloudProviders(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive cloudProvider data", err)
		return err
	}
	if err = formatter.PrintList(cloudProviders); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CloudProviderStoragePlansList subcommand function
func CloudProviderStoragePlansList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	storagePlans, err := svc.ListServerStoragePlans(cmd.GetContext(), viper.GetString(cmd.CloudProviderId))
	if err != nil {
		formatter.PrintError("Couldn't receive storage plans data", err)
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(cmd.GetContext())
	if err != nil {
		return err
	}
	locationsMap, err := cli.LoadLocationsMapping(cmd.GetContext())
	if err != nil {
		return err
	}

	for id, sp := range storagePlans {
		storagePlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
		storagePlans[id].LocationName = locationsMap[sp.LocationID]
	}

	if err = formatter.PrintList(storagePlans); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CloudProviderLoadBalancerPlansList subcommand function
func CloudProviderLoadBalancerPlansList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	loadBalancerPlans, err := svc.ListLoadBalancerPlans(cmd.GetContext(), viper.GetString(cmd.CloudProviderId))
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer plans data", err)
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(cmd.GetContext())
	if err != nil {
		return err
	}
	for id, sp := range loadBalancerPlans {
		loadBalancerPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
	}

	if err = formatter.PrintList(loadBalancerPlans); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CloudProviderClusterPlansList subcommand function
func CloudProviderClusterPlansList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	loadBalancerPlans, err := svc.ListClusterPlans(cmd.GetContext(), viper.GetString(cmd.CloudProviderId))
	if err != nil {
		formatter.PrintError("Couldn't receive cluster plans data", err)
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(cmd.GetContext())
	if err != nil {
		return err
	}
	for id, sp := range loadBalancerPlans {
		loadBalancerPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
	}

	if err = formatter.PrintList(loadBalancerPlans); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
