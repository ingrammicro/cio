package locations

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Locations.",
			Action: cmd.LocationList,
		},
	}
}
