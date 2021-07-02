// Copyright (c) 2017-2021 Ingram Micro Inc.

package storage_plans

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns storage plans commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "show",
			Usage:  "Shows information about the storage plan identified by the given id",
			Action: cmd.StoragePlanShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Storage Plan Id",
				},
			},
		},
	}
}
