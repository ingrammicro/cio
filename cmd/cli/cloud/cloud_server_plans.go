// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloud

import (
	"context"
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
		Short:        "This action shows information about the server plan identified by the given id",
		RunMethod:    ServerPlanShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// ServerPlanList subcommand function
func ServerPlanList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	serverPlans, err := svc.ListServerPlans(
		ctx,
		viper.GetString(cmd.CloudProviderId),
		viper.GetString(cmd.RealmId))
	if err != nil {
		formatter.PrintError("Couldn't receive server plans data", err)
		return err
	}
	if err = FormatServerPlansResponse(ctx, serverPlans, formatter); err != nil {
		return err
	}
	return nil
}

// FormatServerPlansResponse processes and prints received server plans
func FormatServerPlansResponse(ctx context.Context, serverPlans []*types.ServerPlan, formatter format.Formatter) error {
	logger.DebugFuncInfo()

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(ctx)
	if err != nil {
		return err
	}
	locationsMap, err := cli.LoadLocationsMapping(ctx)
	if err != nil {
		return err
	}

	for id, sp := range serverPlans {
		serverPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
		serverPlans[id].LocationName = locationsMap[sp.LocationID]
	}

	if err := formatter.PrintList(serverPlans); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerPlanShow subcommand function
func ServerPlanShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	serverPlan, err := svc.GetServerPlan(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive server plan data", err)
		return err
	}

	locationsMap, err := cli.LoadLocationsMapping(ctx)
	if err != nil {
		return err
	}
	serverPlan.LocationName = locationsMap[serverPlan.LocationID]

	if err = formatter.PrintItem(*serverPlan); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
