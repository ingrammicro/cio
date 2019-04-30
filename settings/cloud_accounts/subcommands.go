package cloud_accounts

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns cloud accounts commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the cloud accounts of the account group.",
			Action: cmd.CloudAccountList,
		},
	}
}
