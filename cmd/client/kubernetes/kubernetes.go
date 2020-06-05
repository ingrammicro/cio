// Copyright (c) 2017-2021 Ingram Micro Inc.

package kubernetes

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var KubernetesCmd *cobra.Command

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	KubernetesCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "kubernetes",
		Short:   "Manages kubernetes commands for clusters and node pools",
		Aliases: []string{"k8s"}},
	)
}
