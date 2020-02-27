package ssh_profiles

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns ssh profiles commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available SSH profiles.",
			Action: cmd.SSHProfileList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the SSH profile identified by the given id.",
			Action: cmd.SSHProfileShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "SSH profile id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new SSH profile.",
			Action: cmd.SSHProfileCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the SSH profile",
				},
				cli.StringFlag{
					Name:  "public-key",
					Usage: "Public key of the SSH profile",
				},
				cli.StringFlag{
					Name:  "private-key",
					Usage: "Private key of the SSH profile",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with SSH profile",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing SSH profile",
			Action: cmd.SSHProfileUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "SSH profile id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the SSH profile",
				},
				cli.StringFlag{
					Name:  "public-key",
					Usage: "Public key of the SSH profile",
				},
				cli.StringFlag{
					Name:  "private-key",
					Usage: "Private key of the SSH profile",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a SSH profile",
			Action: cmd.SSHProfileDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "SSH profile id",
				},
			},
		},
		{
			Name:   "add-label",
			Usage:  "This action assigns a single label from a single labelable resource",
			Action: cmd.LabelAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "SSH profile id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "ssh_profile",
					Hidden: true,
				},
			},
		},
		{
			Name:   "remove-label",
			Usage:  "This action unassigns a single label from a single labelable resource",
			Action: cmd.LabelRemove,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "SSH profile id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "ssh_profile",
					Hidden: true,
				},
			},
		},
	}
}
