package definitions

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns policy definitions commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists policy definitions",
			Action: cmd.PolicyDefinitionList,
		},
		{
			Name:   "show",
			Usage:  "Shows policy definition",
			Action: cmd.PolicyDefinitionShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Policy Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates policy definition",
			Action: cmd.PolicyDefinitionCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the policy definition",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the policy definition",
				},
				cli.StringFlag{
					Name:  "definition",
					Usage: "The definition used to configure the policy, as a json formatted parameter. \n\ti.e: --definition '{\"parameters\": {\"prefix\": {\"type\": \"string\",\"metadata\": {\"description\": \"prefix data\"}},\"suffix\": {\"type\": \"string\",\"metadata\": {\"description\": \"suffix data\"}}},\"policyRule\": {\"if\": {\"not\": {\"field\": \"name\",\"like\": \"[concat(parameters('prefix'), '*', parameters('suffix'))]\"}},\"then\": {\"effect\": \"audit\"}}}'",
				},
				cli.StringFlag{
					Name:  "definition-from-file",
					Usage: "The definition used to configure the policy, from file or STDIN, as a json formatted parameter. \n\tFrom file: --definition-from-file def.json \n\tFrom STDIN: --definition-from-file -",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing policy definition identified by the given id",
			Action: cmd.PolicyDefinitionUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Policy Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the policy definition",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the policy definition",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a policy definition",
			Action: cmd.PolicyDefinitionDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Policy Id",
				},
			},
		},
		{
			Name:   "list-assignments",
			Usage:  "lists policy assignments",
			Action: cmd.PolicyDefinitionListAssignments,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Policy Id",
				},
			},
		},
	}
}
