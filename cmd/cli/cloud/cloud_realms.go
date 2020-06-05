// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloud

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Cloud provider id"}

	fRealmId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Realm id"}

	realmsCmd := cmd.NewCommand(cloudCmd, &cmd.CommandContext{
		Use:   "realms",
		Short: "Provides information on realms"},
	)
	cmd.NewCommand(realmsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all realms of a cloud provider",
		RunMethod:    RealmList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(realmsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the realm identified by the given id",
		RunMethod:    RealmShow,
		FlagContexts: []cmd.FlagContext{fRealmId}},
	)
	cmd.NewCommand(realmsCmd, &cmd.CommandContext{
		Use:          "list-node-pool-plans",
		Short:        "This action lists the node pool plans offered by the realm identified by the given id",
		RunMethod:    RealmNodePoolPlansList,
		FlagContexts: []cmd.FlagContext{fRealmId}},
	)
}

// RealmList subcommand function
func RealmList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	realms, err := svc.ListRealms(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive realms data", err)
	}

	if err = formatter.PrintList(realms); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// RealmShow subcommand function
func RealmShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	realm, err := svc.GetRealm(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive realm data", err)
	}

	if err = formatter.PrintItem(*realm); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// RealmNodePoolPlansList subcommand function
func RealmNodePoolPlansList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	realms, err := svc.ListRealmNodePoolPlans(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive realms data", err)
	}

	if err = formatter.PrintList(realms); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
