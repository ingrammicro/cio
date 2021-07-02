// Copyright (c) 2017-2021 Ingram Micro Inc.

package settings

import (
	"github.com/ingrammicro/cio/settings/cloud_accounts"
	"github.com/ingrammicro/cio/settings/policies"
	"github.com/urfave/cli"
)

// SubCommands returns settings commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "cloud-accounts",
			Usage:       "Provides information about cloud accounts",
			Subcommands: append(cloud_accounts.SubCommands()),
		},
		{
			Name:        "policies",
			Usage:       "Provides information about policies",
			Subcommands: append(policies.SubCommands()),
		},
	}
}
