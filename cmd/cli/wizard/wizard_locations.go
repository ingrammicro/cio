// Copyright (c) 2017-2022 Ingram Micro Inc.

package wizard

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
)

func init() {
	locationsCmd := cmd.NewCommand(wizardCmd, &cmd.CommandContext{
		Use:   "locations",
		Short: "Provides information about locations"},
	)
	cmd.NewCommand(locationsCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists the available Locations",
		RunMethod: LocationList},
	)
}

// LocationList subcommand function
func LocationList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	locations, err := svc.ListLocations(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive location data", err)
	}
	if err = formatter.PrintList(locations); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
