package import_candidates

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns import candidates commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list-servers",
			Usage:  "Lists the servers that are found on the cloud account identified by the given id, that can be imported into IMCO",
			Action: cmd.ImportCandidateServerList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "list-vpcs",
			Usage:  "Lists the VPCs that are found on the cloud account identified by the given id, that can be imported into IMCO",
			Action: cmd.ImportCandidateVPCList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "list-floating-ips",
			Usage:  "Lists the floating IPs that are found on the cloud account identified by the given id, that can be imported into IMCO",
			Action: cmd.ImportCandidateFloatingIPList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "list-volumes",
			Usage:  "Lists the volumes that are found on the cloud account identified by the given id, that can be imported into IMCO",
			Action: cmd.ImportCandidateVolumeList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "show-server",
			Usage:  "Show server import candidate, given its id",
			Action: cmd.ImportCandidateServerShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Import candidate server Id",
				},
			},
		},
		{
			Name:   "show-vpc",
			Usage:  "Show VPC import candidate, given its id",
			Action: cmd.ImportCandidateVPCShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Import candidate VPC Id",
				},
			},
		},
		{
			Name:   "show-floating-ip",
			Usage:  "Show floating IP import candidate, given its id",
			Action: cmd.ImportCandidateFloatingIPShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Import candidate floating IP Id",
				},
			},
		},
		{
			Name:   "show-volume",
			Usage:  "Show volume import candidate, given its id",
			Action: cmd.ImportCandidateVolumeShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Import candidate volume Id",
				},
			},
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
