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
			Usage:       "Provides information about brownfield cloud accounts. Allows querying cloud accounts to import resources from IMCO.",
			Subcommands: append(cloud_accounts.SubCommands()),
		},
		{
			Name:   "import-servers",
			Usage:  "Import servers for a given cloud account id",
			Action: cmd.ImportServers,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-vpcs",
			Usage:  "Import VPCs for a given cloud account id",
			Action: cmd.ImportVPCs,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-floating-ips",
			Usage:  "Import Floating IPs for a given cloud account id",
			Action: cmd.ImportFloatingIPs,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-volumes",
			Usage:  "Import volumes for a given cloud account id",
			Action: cmd.ImportVolumes,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-policies",
			Usage:  "Import policies for a given cloud account id",
			Action: cmd.ImportPolicies,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
		{
			Name:   "import-k8s-clusters",
			Usage:  "Import kubernetes clusters for a given cloud account id",
			Action: cmd.ImportKubernetesClusters,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cloud account Id",
				},
			},
		},
	}
}