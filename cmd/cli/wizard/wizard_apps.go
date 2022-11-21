// Copyright (c) 2017-2022 Ingram Micro Inc.

package wizard

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true,
		Usage: "Identifier of the App which will be deployed"}

	fLocationId := cmd.FlagContext{Type: cmd.String, Name: cmd.LocationId, Required: true,
		Usage: "Identifier of the Location on which the App will be deployed"}

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the Cloud Account with which the App will be deployed"}

	fHostname := cmd.FlagContext{Type: cmd.String, Name: cmd.Hostname, Required: true,
		Usage: "A hostname for the cloud server to deploy"}

	fServerPlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerPlanId,
		Usage: "Identifier of the Server Plan on which the App will be deployed"}

	appsCmd := cmd.NewCommand(wizardCmd, &cmd.CommandContext{
		Use:   "apps",
		Short: "Provides information about apps"})
	cmd.NewCommand(appsCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists the available apps",
		RunMethod: AppList})
	cmd.NewCommand(appsCmd, &cmd.CommandContext{
		Use:          "deploy",
		Short:        "Deploys the app with the given id as a server on the cloud",
		RunMethod:    AppDeploy,
		FlagContexts: []cmd.FlagContext{fId, fLocationId, fCloudAccountId, fHostname, fServerPlanId}},
	)
}

// AppList subcommand function
func AppList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	apps, err := svc.ListApps(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive apps data", err)
		return err
	}
	if err = formatter.PrintList(apps); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// AppDeploy subcommand function
func AppDeploy() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	appIn := map[string]interface{}{
		"location_id":      viper.GetString(cmd.LocationId),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
		"hostname":         viper.GetString(cmd.Hostname),
	}
	if viper.IsSet(cmd.ServerPlanId) {
		appIn["server_plan_id"] = viper.GetString(cmd.ServerPlanId)
	}

	server, err := svc.DeployApp(cmd.GetContext(), viper.GetString(cmd.Id), &appIn)
	if err != nil {
		formatter.PrintError("Couldn't deploy app", err)
		return err
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
