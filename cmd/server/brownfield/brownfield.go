// Copyright (c) 2017-2021 Ingram Micro Inc.

package brownfield

import (
	"runtime"

	"github.com/ingrammicro/cio/cmd"
)

func init() {
	var flagContexts []cmd.FlagContext
	if runtime.GOOS == "windows" {
		flagContexts = append(flagContexts, cmd.FlagContext{Type: cmd.String, Name: cmd.AdminPassword,
			Usage: "The password for your current (administrator) user"},
		)
	}

	brownfieldCmd := cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:   "brownfield",
		Short: "Manages registration and configuration within an imported brownfield Host"},
	)
	cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use:       "register",
		Short:     "Register orchestrator agent within an imported brownfield Host",
		RunMethod: Register},
	)
	cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use:          "configure",
		Short:        "Configures an imported brownfield Host's FQDN, SSH access, agent services and firewall",
		RunMethod:    Configure,
		FlagContexts: flagContexts},
	)
}
