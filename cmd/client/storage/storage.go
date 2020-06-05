// Copyright (c) 2017-2021 Ingram Micro Inc.

package storage

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var StorageCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	StorageCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "storage",
		Short:   "Manages storage commands for plans and volumes",
		Aliases: []string{"st"}},
	)
}
