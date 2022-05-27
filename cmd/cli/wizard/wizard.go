// Copyright (c) 2017-2022 Ingram Micro Inc.

package wizard

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var wizardCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	wizardCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "wizard",
		Short:   "Manages wizard related commands for apps, locations, cloud providers, server plans",
		Aliases: []string{"wiz"}},
	)
}
