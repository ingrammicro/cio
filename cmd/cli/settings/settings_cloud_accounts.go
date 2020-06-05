// Copyright (c) 2017-2022 Ingram Micro Inc.

package settings

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Cloud Account Id"}

	cloudAccountsCmd := cmd.NewCommand(settingsCmd, &cmd.CommandContext{
		Use:   "cloud-accounts",
		Short: "Provides information about cloud accounts"},
	)
	cmd.NewCommand(cloudAccountsCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists the cloud accounts of the account group",
		RunMethod: CloudAccountList},
	)
	cmd.NewCommand(cloudAccountsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about a specific cloud account",
		RunMethod:    CloudAccountShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// CloudAccountList subcommand function
func CloudAccountList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cloudAccounts, err := svc.ListCloudAccounts(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudAccount data", err)
	}

	cloudProvidersMap := cli.LoadCloudProvidersMapping(cmd.GetContext())

	for id, ca := range cloudAccounts {
		cloudAccounts[id].CloudProviderName = cloudProvidersMap[ca.CloudProviderID]
	}

	if err = formatter.PrintList(cloudAccounts); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudAccountShow subcommand function
func CloudAccountShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cloudAccount, err := svc.GetCloudAccount(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudAccount data", err)
	}

	cloudProvidersMap := cli.LoadCloudProvidersMapping(cmd.GetContext())

	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
