// Copyright (c) 2017-2021 Ingram Micro Inc.

package server_plans

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns server plans commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Server Plans.",
			Action: cmd.WizServerPlanList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "app-id",
					Usage: "Identifier of the App",
				},
				cli.StringFlag{
					Name:  "location-id",
					Usage: "Identifier of the Location",
				},
				cli.StringFlag{
					Name:  "cloud-provider-id",
					Usage: "Identifier of the Cloud Provider",
				},
			},
		},
	}
}
