// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudspecificextensions

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var CloudSpecificExtensionsCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	CloudSpecificExtensionsCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "cloud-specific-extensions",
		Short:   "Manages cloud specific extensions -CSEs- templates and deployments",
		Aliases: []string{"cse"}},
	)
}
