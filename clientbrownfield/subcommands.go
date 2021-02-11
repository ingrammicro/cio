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
			Usage:       "Provides information about brownfield cloud accounts. Allows querying cloud accounts to import candidate resources from IMCO.",
			Subcommands: append(cloud_accounts.SubCommands()),
		},
		{
			Name:   "import-servers",
			Usage:  "Import servers candidates, by given cloud account id",
			Action: cmd.ImportCandidateServers,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-vpcs",
			Usage:  "Import VPCs candidates, by given cloud account id",
			Action: cmd.ImportCandidateVPCs,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-floating-ips",
			Usage:  "Import Floating IPs candidates, by given cloud account id",
			Action: cmd.ImportCandidateFloatingIPs,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-volumes",
			Usage:  "Import volumes candidates, by given cloud account id",
			Action: cmd.ImportCandidateVolumes,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-kb-clusters",
			Usage:  "Import kubernetes clusters candidates, by given cloud account id",
			Action: cmd.ImportCandidateKubernetesClusters,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-policies",
			Usage:  "Import policies candidates, by given cloud account id",
			Action: cmd.ImportCandidatePolicies,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
	}
}
