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
			Usage:  "Lists the cloud accounts that support discovering and importing resources",
			Action: cmd.BrownfieldCloudAccountList,
		},
		{
			Name:   "discover-servers",
			Usage:  "Starts the process of discovering servers on the cloud account identified by the given id",
			Action: cmd.BrownfieldCloudAccountServersDiscover,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "discover-vpcs",
			Usage:  "Starts the process of discovering VPCs on the cloud account identified by the given id",
			Action: cmd.BrownfieldCloudAccountVPCsDiscover,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "discover-floating-ips",
			Usage:  "Starts the process of discovering floating IPs on the cloud account identified by the given id",
			Action: cmd.BrownfieldCloudAccountFloatingIPsDiscover,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "discover-volumes",
			Usage:  "Starts the process of discovering volumes on the cloud account identified by the given id",
			Action: cmd.BrownfieldCloudAccountVolumesDiscover,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
	}
}
