// Copyright (c) 2017-2021 Ingram Micro Inc.

package brownfield

import (
	"runtime"

	"github.com/urfave/cli"
)

// SubCommands returns brownfield commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "register",
			Usage:  "Register concerto agent within an imported brownfield Host",
			Action: cmdRegister,
		},
		{
			Name:   "configure",
			Usage:  "Configures an imported brownfield Host's FQDN, SSH access, agent services and firewall",
			Action: cmdConfigure,
			Flags:  configureFlags(),
		},
	}
}

func configureFlags() []cli.Flag {
	if runtime.GOOS != "windows" {
		return nil
	}
	return []cli.Flag{
		cli.StringFlag{
			Name:  "admin-password",
			Usage: "The password for your current (administrator) user",
		},
	}
}
