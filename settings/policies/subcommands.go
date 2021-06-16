// Copyright (c) 2017-2021 Ingram Micro Inc.

package policies

import (
	"github.com/ingrammicro/cio/settings/policies/assignments"
	"github.com/ingrammicro/cio/settings/policies/definitions"
	"github.com/urfave/cli"
)

// SubCommands returns policy commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "definitions",
			Usage:       "Provides information about policy definitions",
			Subcommands: append(definitions.SubCommands()),
		},
		{
			Name:        "assignments",
			Usage:       "Provides information about policy assignments",
			Subcommands: append(assignments.SubCommands()),
		},
	}
}
