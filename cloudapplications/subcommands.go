package cloudapplications

import (
	"github.com/ingrammicro/cio/cloudapplications/deployments"
	"github.com/ingrammicro/cio/cloudapplications/templates"
	"github.com/urfave/cli"
)

// SubCommands returns cloud applications commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "templates",
			Usage:       "Provides information about CAT",
			Subcommands: append(templates.SubCommands()),
		},
		{
			Name:        "deployments",
			Usage:       "Provides information about CAT deployments",
			Subcommands: append(deployments.SubCommands()),
		},
	}
}
