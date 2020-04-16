package clientbrownfield

import (
	"github.com/ingrammicro/cio/clientbrownfield/cloud_accounts"
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns client brownfield commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "cloud-accounts",
			Usage:       "Provides information about brownfield cloud accounts. Allows querying cloud accounts to discover (and list) candidate resources from IMCO.",
			Subcommands: append(cloud_accounts.SubCommands()),
		},
		{
			Name:   "import-server",
			Usage:  "Import server import candidate, given its id",
			Action: cmd.ImportCandidateServerImport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Import candidate server Id",
				},
				cli.StringFlag{
					Name:  "ssh-profile-id",
					Usage: "Identifier of the ssh profile which the server shall use",
				},
				cli.StringFlag{
					Name:  "ssh-profile-ids",
					Usage: "A list of comma separated ssh profiles ids",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "import-vpc",
			Usage:  "Import VPC import candidate, given its id",
			Action: cmd.ImportCandidateVPCImport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Import candidate VPC Id",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "import-floating-ip",
			Usage:  "Import Floating IP import candidate, given its id",
			Action: cmd.ImportCandidateFloatingIPImport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Import candidate floating IP Id",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "import-volume",
			Usage:  "Import volume import candidate, given its id",
			Action: cmd.ImportCandidateVolumeImport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Import candidate volume Id",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
	}
}
