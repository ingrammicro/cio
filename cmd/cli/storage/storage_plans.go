// Copyright (c) 2017-2022 Ingram Micro Inc.

package storage

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Storage Plan Id"}

	plansCmd := cmd.NewCommand(storageCmd, &cmd.CommandContext{
		Use:   "plans",
		Short: "Provides information on storage plans"},
	)
	cmd.NewCommand(plansCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the storage plan identified by the given id",
		RunMethod:    StoragePlanShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// StoragePlanShow subcommand function
func StoragePlanShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	storagePlans, err := svc.GetStoragePlan(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive storage plan data", err)
		return err
	}

	locationsMap, err := cli.LoadLocationsMapping(ctx)
	if err != nil {
		return err
	}
	storagePlans.LocationName = locationsMap[storagePlans.LocationID]

	if err = formatter.PrintItem(*storagePlans); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
