// Copyright (c) 2017-2021 Ingram Micro Inc.

package settings

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/cobra"
)

var PoliciesCmd *cobra.Command

func init() {
	PoliciesCmd = cmd.NewCommand(SettingsCmd, &cmd.CommandContext{
		Use:   "policies",
		Short: "Provides information about policies"},
	)
}
