// Copyright (c) 2017-2022 Ingram Micro Inc.

package brownfield

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var brownfieldCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	brownfieldCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use: "brownfield",
		Short: "Manages brownfield resources, allowing users to discover and import servers, VPCs, floating IPs, " +
			"volumes and policies from different cloud accounts into the system",
		Aliases: []string{"bf"}},
	)
}
