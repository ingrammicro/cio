// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloud

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/spf13/viper"
)

func init() {
	fCloudProviderId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudProviderId, Required: true,
		Usage: "Cloud provider id"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Server plan id"}
	fRealmId := cmd.FlagContext{Type: cmd.String, Name: cmd.RealmId, Required: true,
		Usage: "Identifier of the realm to which the server plan belongs"}

	serverPlansCmd := cmd.NewCommand(cloudCmd, &cmd.CommandContext{
		Use:   "server-plans",
		Short: "Provides information on server plans"},
	)
	cmd.NewCommand(serverPlansCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "This action lists the server plans offered by the cloud provider identified by the given id",
		RunMethod:    ServerPlanList,
		FlagContexts: []cmd.FlagContext{fCloudProviderId, fRealmId}},
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
	svc, _, formatter := cli.WireUpAPIClient()

	serverPlans, err := svc.ListServerPlans(
		cmd.GetContext(),
		viper.GetString(cmd.CloudProviderId),
		viper.GetString(cmd.RealmId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}
	if err = FormatServerPlansResponse(serverPlans, formatter); err != nil {
		return err
	}
	return nil
}

// FormatServerPlansResponse processes and prints received server plans
func FormatServerPlansResponse(serverPlans []*types.ServerPlan, formatter format.Formatter) error {
	logger.DebugFuncInfo()

	cloudProvidersMap := cli.LoadCloudProvidersMapping(cmd.GetContext())
	locationsMap := cli.LoadLocationsMapping(cmd.GetContext())

	for id, sp := range serverPlans {
		serverPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
		serverPlans[id].LocationName = locationsMap[sp.LocationID]
	}

	if err := formatter.PrintList(serverPlans); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// ServerPlanShow subcommand function
func ServerPlanShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverPlan, err := svc.GetServerPlan(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	locationsMap := cli.LoadLocationsMapping(cmd.GetContext())
	serverPlan.LocationName = locationsMap[serverPlan.LocationID]

	if err = formatter.PrintItem(*serverPlan); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
