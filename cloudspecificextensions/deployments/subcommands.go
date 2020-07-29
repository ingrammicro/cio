package deployments

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns cloud specific extensions deployments commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "List CSE deployments",
			Action: cmd.CloudSpecificExtensionDeploymentList,
		},
		{
			Name:   "show",
			Usage:  "Shows CSE deployment",
			Action: cmd.CloudSpecificExtensionDeploymentShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CSE deployment Id",
				},
			},
		},
		{
			Name:   "deploy",
			Usage:  "Deploys a new CSE deployment from CSE template",
			Action: cmd.CloudSpecificExtensionDeploymentCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CSE template Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the CSE deployment",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account in which is deployed",
				},
				cli.StringFlag{
					Name:  "realm-id",
					Usage: "Identifier of the realm in which is deployed",
				},
				cli.StringFlag{
					Name:  "parameters",
					Usage: "The parameters used to configure the CSE deployment, as a json formatted parameter. \n\ti.e: --parameters '{\"param1\":\"val1\",\"param2\":\"val2\",\"param3\":{\"id\":\"val3\"},\"param4\":true}'",
				},
				cli.StringFlag{
					Name:  "parameters-from-file",
					Usage: "The parameters used to configure the CSE deployment, from file or STDIN, as a json formatted parameter. \n\tFrom file: --parameters-from-file params.json \n\tFrom STDIN: --parameters-from-file -",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with CSE deployment",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing CSE deployment identified by the given id",
			Action: cmd.CloudSpecificExtensionDeploymentUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CSE deployment Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the CSE deployment",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a CSE deployment",
			Action: cmd.CloudSpecificExtensionDeploymentDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CSE deployment Id",
				},
			},
		},
	}
}
