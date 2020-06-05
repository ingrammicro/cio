// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var CloudCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	CloudCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use: "cloud",
		Short: "Manages cloud related commands for server arrays, servers, generic images, ssh profiles, " +
			"cloud providers, realms, server plans and infrastructure archives",
		Aliases: []string{"clo"}},
	)
}
