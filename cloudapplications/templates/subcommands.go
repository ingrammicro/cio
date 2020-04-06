package templates

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns cloud application templates commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists CATs",
			Action: cmd.CloudApplicationTemplateList,
		},
		{
			Name:   "show",
			Usage:  "Shows CAT",
			Action: cmd.CloudApplicationTemplateShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CAT Id",
				},
			},
		},
		{
			Name:   "upload",
			Usage:  "Uploads a CAT",
			Action: cmd.CloudApplicationTemplateUpload,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the CAT",
				},
				cli.StringFlag{
					Name:  "filepath",
					Usage: "path to CAT file",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a CAT",
			Action: cmd.CloudApplicationTemplateDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "CAT Id",
				},
			},
		},
	}
}
