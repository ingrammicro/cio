// Copyright (c) 2017-2021 Ingram Micro Inc.

package deployments

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns cloud application deployments commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists deployments",
			Action: cmd.CloudApplicationDeploymentList,
		},
		{
			Name:   "show",
			Usage:  "Shows deployment",
			Action: cmd.CloudApplicationDeploymentShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Deployment Id",
				},
			},
		},
		{
			Name:   "deploy",
			Usage:  "Deploys a CAT",
			Action: cmd.CloudApplicationDeploymentDeploy,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cat Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the deployment",
				},
				cli.StringFlag{
					Name: "inputs",
					Usage: "The inputs used to configure the cloud application deployment, " +
						"as a json formatted parameter. \n\t" +
						"i.e: --inputs '{\"region\":{\"cloud_provider\":\"Azure\",\"name\":\"US\"}," +
						"\"server_plan\":\"Standard_D2_v3\",\"admin_user\":\"admin\",\"admin_password\":\"abc$1\"}'",
				},
				cli.StringFlag{
					Name: "inputs-from-file",
					Usage: "The inputs used to configure the cloud application deployment, " +
						"from file or STDIN, as a json formatted parameter. \n\t" +
						"From file: --inputs-from-file attrs.json \n\tFrom STDIN: --inputs-from-file -",
				},
				cli.IntFlag{
					Name:  "time, t",
					Usage: "Time lapse -seconds- for deployment status check",
					Value: cmd.DefaultTimeLapseDeploymentStatusCheck,
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a deployment",
			Action: cmd.CloudApplicationDeploymentDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Deployment Id",
				},
				cli.IntFlag{
					Name:  "time, t",
					Usage: "Time lapse -seconds- for deletion status check",
					Value: cmd.DefaultTimeLapseDeletionStatusCheck,
				},
			},
		},
	}
}
