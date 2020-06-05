// Copyright (c) 2017-2022 Ingram Micro Inc.

package wizard

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/cmd/cli/cloud"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fAppId := cmd.FlagContext{Type: cmd.String, Name: cmd.AppId, Required: true, Usage: "Identifier of the App"}

	fLocationId := cmd.FlagContext{Type: cmd.String, Name: cmd.LocationId, Required: true,
		Usage: "Identifier of the Location"}

	fCloudProviderId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudProviderId, Required: true,
		Usage: "Identifier of the Cloud Provider"}

	serverPlansCmd := cmd.NewCommand(wizardCmd, &cmd.CommandContext{
		Use:   "server-plans",
		Short: "Provides information about server plans"},
	)
	cmd.NewCommand(serverPlansCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists the available server Plans",
		RunMethod:    WizServerPlanList,
		FlagContexts: []cmd.FlagContext{fAppId, fLocationId, fCloudProviderId}},
	)
}

// WizServerPlanList subcommand function
func WizServerPlanList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverPlans, err := svc.ListWizardServerPlans(cmd.GetContext(),
		viper.GetString(cmd.AppId),
		viper.GetString(cmd.LocationId),
		viper.GetString(cmd.CloudProviderId),
	)
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}
	if err = cloud.FormatServerPlansResponse(serverPlans, formatter); err != nil {
		return err
	}
	return nil
}
