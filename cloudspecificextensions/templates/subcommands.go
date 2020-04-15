package templates

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns cloud specific extensions templates commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "List CSE templates",
			Action: cmd.CloudSpecificExtensionTemplateList,
		},
		{
			Name:   "show",
			Usage:  "Shows CSE template",
			Action: cmd.CloudSpecificExtensionTemplateShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CSE template Id",
				},
			},
		},
		{
			Name:   "import",
			Usage:  "Imports a CSE template",
			Action: cmd.CloudSpecificExtensionTemplateImport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the CSE template",
				},
				cli.StringFlag{
					Name:  "syntax",
					Usage: "Cloud provider syntax of the CSE template",
				},
				cli.StringFlag{
					Name:  "definition",
					Usage: "The definition used to configure the CSE template, as a json formatted parameter. \n\ti.e: --definition '{\"$schema\":\"https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#\",\"contentVersion\":\"1.0.0.0\",\"parameters\":{\"vmName\":{\"type\":\"string\",\"defaultValue\": \"simpleLinuxVM\",\"metadata\":{\"description\": \"The name of you Virtual Machine.\"}}}}'",
				},
				cli.StringFlag{
					Name:  "definition-from-file",
					Usage: "The definition used to configure the CSE template, from file or STDIN, as a json formatted parameter. \n\tFrom file: --definition-from-file def.json \n\tFrom STDIN: --definition-from-file -",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with CSE template",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing CSE template identified by the given id",
			Action: cmd.CloudSpecificExtensionTemplateUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CSE template Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the CSE template",
				},
			},
		},
		{
			Name:   "list-deployments",
			Usage:  "List CSE deployments of a CSE template",
			Action: cmd.CloudSpecificExtensionTemplateListDeployments,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CSE template Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a CSE template",
			Action: cmd.CloudSpecificExtensionTemplateDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CSE template Id",
				},
			},
		},
	}
}
