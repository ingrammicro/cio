// Copyright (c) 2017-2021 Ingram Micro Inc.

package realms

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns realms commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all realms of a cloud provider",
			Action: cmd.RealmList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud provider id",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the realm identified by the given id",
			Action: cmd.RealmShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Realm id",
				},
			},
		},
		{
			Name:   "list-node-pool-plans",
			Usage:  "This action lists the node pool plans offered by the realm identified by the given id",
			Action: cmd.RealmNodePoolPlansList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Realm id",
				},
			},
		},
	}
}
