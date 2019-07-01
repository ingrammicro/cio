package locations

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/cio/cmd"
)

// SubCommands returns locations commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Locations.",
			Action: cmd.LocationList,
		},
	}
}
