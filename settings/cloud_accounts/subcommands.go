// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud_accounts

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns cloud accounts commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the cloud accounts of the account group.",
			Action: cmd.CloudAccountList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific cloud account",
			Action: cmd.CloudAccountShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud Account Id",
				},
			},
		},
	}
}
