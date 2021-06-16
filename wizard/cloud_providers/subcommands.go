// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud_providers

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns cloud providers commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Cloud Providers",
			Action: cmd.WizCloudProviderList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "app-id",
					Usage: "Identifier of the App",
				},
				cli.StringFlag{
					Name:  "location-id",
					Usage: "Identifier of the Location",
				},
			},
		},
	}
}
