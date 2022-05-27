// Copyright (c) 2017-2022 Ingram Micro Inc.

package settings

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var settingsCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	settingsCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "settings",
		Short:   "Provides settings for cloud accounts and policies",
		Aliases: []string{"set"}},
	)
}
