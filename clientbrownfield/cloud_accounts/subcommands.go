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
			Usage:  "Lists the cloud accounts that support importing resources",
			Action: cmd.BrownfieldCloudAccountList,
		},
	}
}
