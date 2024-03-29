// Copyright (c) 2017-2021 Ingram Micro Inc.

package server_plan

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns server plan commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "This action lists the server plans offered by the cloud provider identified by the given id.",
			Action: cmd.ServerPlanList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud-provider-id",
					Usage: "Cloud provider id",
				},
				cli.StringFlag{
					Name:  "realm-id",
					Usage: "Identifier of the realm to which the server plan belongs",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "This action shows information about the Server Plan identified by the given id.",
			Action: cmd.ServerPlanShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server plan id",
				},
			},
		},
	}
}
