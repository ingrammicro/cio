// Copyright (c) 2017-2022 Ingram Micro Inc.

package storage

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var storageCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	storageCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "storage",
		Short:   "Manages storage commands for plans and volumes",
		Aliases: []string{"st"}},
	)
}
