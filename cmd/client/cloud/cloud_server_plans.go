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

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Server plan id"}

	serverPlansCmd := cmd.NewCommand(CloudCmd, &cmd.CommandContext{
		Use:   "server-plans",
		Short: "Provides information on server plans"},
	)
	cmd.NewCommand(serverPlansCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "This action lists the server plans offered by the cloud provider identified by the given id",
		RunMethod:    ServerPlanList,
		FlagContexts: []cmd.FlagContext{fCloudProviderId}},
	)
	cmd.NewCommand(serverPlansCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "This action shows information about the server Plan identified by the given id",
		RunMethod:    ServerPlanShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// ServerPlanList subcommand function
func ServerPlanList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	serverPlans, err := svc.ListServerPlans(viper.GetString(cmd.CloudProviderId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	cloudProvidersMap := cmd.LoadCloudProvidersMapping()
	locationsMap := cmd.LoadLocationsMapping()

	for id, sp := range serverPlans {
		serverPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
		serverPlans[id].LocationName = locationsMap[sp.LocationID]
	}

	if err = formatter.PrintList(serverPlans); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// ServerPlanShow subcommand function
func ServerPlanShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	serverPlan, err := svc.GetServerPlan(viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	locationsMap := cmd.LoadLocationsMapping()
	serverPlan.LocationName = locationsMap[serverPlan.LocationID]

	if err = formatter.PrintItem(*serverPlan); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
