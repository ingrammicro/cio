// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudapplications

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var CloudApplicationsCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	CloudApplicationsCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "cloud-applications",
		Short:   "Manages cloud application templates -CATs- and deployments",
		Aliases: []string{"ca"}},
	)
}
