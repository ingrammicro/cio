// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var NetworkCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	NetworkCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "network",
		Short:   "Manages network related commands",
		Aliases: []string{"net"}},
	)
}
