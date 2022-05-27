// Copyright (c) 2017-2022 Ingram Micro Inc.

package blueprint

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var blueprintCmd *cobra.Command

func init() {
	blueprintCmd = cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "blueprint",
		Short:   "Manages blueprint commands for scripts, cookbook versions and templates",
		Aliases: []string{"bl"}},
	)
}
