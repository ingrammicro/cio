package cloudspecificextensions

import (
	"github.com/ingrammicro/cio/cloudspecificextensions/deployments"
	"github.com/ingrammicro/cio/cloudspecificextensions/templates"
	"github.com/urfave/cli"
)

// SubCommands returns cloud specific extensions commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "templates",
			Usage:       "Provides information about CSE templates",
			Subcommands: append(templates.SubCommands()),
		},
		{
			Name:        "deployments",
			Usage:       "Provides information about CSE deployments",
			Subcommands: append(deployments.SubCommands()),
		},
	}
}
