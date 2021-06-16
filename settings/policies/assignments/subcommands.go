// Copyright (c) 2017-2021 Ingram Micro Inc.

package assignments

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns policy assignments commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "lists policy assignments for a given cloud account",
			Action: cmd.PolicyAssignmentList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows policy assignments",
			Action: cmd.PolicyAssignmentShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Assignment Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates policy assignment",
			Action: cmd.PolicyAssignmentCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the policy assignment",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the policy assignment",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account where the assignment belongs",
				},
				cli.StringFlag{
					Name:  "definition-id",
					Usage: "Identifier of the policy definition to be assigned",
				},
				cli.StringFlag{
					Name: "parameters",
					Usage: "The parameters used to configure the policy assignment, as a json formatted parameter. \n\t" +
						"i.e: --parameters '{\"param1\":\"val1\",\"param2\":\"val2\",\"param3\":{\"id\":\"val3\"},\"param4\":true}'",
				},
				cli.StringFlag{
					Name: "parameters-from-file",
					Usage: "The parameters used to configure the policy assignment, from file or STDIN, " +
						"as a json formatted parameter. \n\t" +
						"From file: --parameters-from-file params.json \n\t" +
						"From STDIN: --parameters-from-file -",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing policy assignment identified by the given id",
			Action: cmd.PolicyAssignmentUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Assignment Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the policy assignment",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the policy assignment",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a policy assignment",
			Action: cmd.PolicyAssignmentDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Assignment Id",
				},
			},
		},
	}
}
