// Copyright (c) 2017-2022 Ingram Micro Inc.

package settings

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var policiesCmd *cobra.Command

func init() {
	policiesCmd = cmd.NewCommand(settingsCmd, &cmd.CommandContext{
		Use:   "policies",
		Short: "Provides information about policies"},
	)
}
