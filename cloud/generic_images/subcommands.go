// Copyright (c) 2017-2021 Ingram Micro Inc.

package generic_images

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns generic images commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "This action lists the available generic images.",
			Action: cmd.GenericImageList,
		},
	}
}
