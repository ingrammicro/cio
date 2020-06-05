// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloudapplications

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var cloudApplicationsCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	cloudApplicationsCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "cloud-applications",
		Short:   "Manages cloud application templates -CATs- and deployments",
		Aliases: []string{"ca"}},
	)
}
