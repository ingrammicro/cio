package blueprint

import (
	"github.com/ingrammicro/cio/blueprint/attachments"
	"github.com/ingrammicro/cio/blueprint/cookbook_versions"
	"github.com/ingrammicro/cio/blueprint/scripts"
	"github.com/ingrammicro/cio/blueprint/templates"
	"github.com/urfave/cli"
)

// SubCommands returns blueprint commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "cookbook-versions",
			Usage:       "Provides information on chef cookbook versions",
			Subcommands: append(cookbook_versions.SubCommands()),
		},
		{
			Name:        "scripts",
			Usage:       "Allow the user to manage the scripts they want to run on the servers",
			Subcommands: append(scripts.SubCommands()),
		},
		{
			Name:        "attachments",
			Usage:       "Allow the user to manage the attachments they want to store on the servers",
			Subcommands: append(attachments.SubCommands()),
		},
		{
			Name:        "templates",
			Usage:       "Provides information on templates",
			Subcommands: append(templates.SubCommands()),
		},
	}
}
