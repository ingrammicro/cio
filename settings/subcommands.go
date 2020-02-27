package settings

import (
	"github.com/ingrammicro/cio/settings/cloud_accounts"
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
	}
}
