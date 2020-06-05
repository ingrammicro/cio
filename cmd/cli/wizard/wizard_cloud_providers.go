// Copyright (c) 2017-2022 Ingram Micro Inc.

package wizard

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fAppId := cmd.FlagContext{Type: cmd.String, Name: cmd.AppId, Required: true, Usage: "Identifier of the App"}
	fLocationId := cmd.FlagContext{Type: cmd.String, Name: cmd.LocationId, Required: true,
		Usage: "Identifier of the Location"}

	cloudProvidersCmd := cmd.NewCommand(wizardCmd, &cmd.CommandContext{
		Use:   "cloud-providers",
		Short: "Provides information about cloud providers"},
	)
	cmd.NewCommand(cloudProvidersCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists the available Cloud Providers",
		RunMethod:    WizCloudProviderList,
		FlagContexts: []cmd.FlagContext{fAppId, fLocationId}},
	)
}

// WizCloudProviderList subcommand function
func WizCloudProviderList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cloudProviders, err := svc.ListWizardCloudProviders(
		cmd.GetContext(),
		viper.GetString(cmd.AppId),
		viper.GetString(cmd.LocationId),
	)
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	if err = formatter.PrintList(cloudProviders); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
