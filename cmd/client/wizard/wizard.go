// Copyright (c) 2017-2021 Ingram Micro Inc.

package wizard

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var WizardCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	WizardCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "wizard",
		Short:   "Manages wizard related commands for apps, locations, cloud providers, server plans",
		Aliases: []string{"wiz"}},
	)
}
