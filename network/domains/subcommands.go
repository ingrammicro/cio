package domains

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/network/domains/records"
	"github.com/urfave/cli"
)

// SubCommands returns domains commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all DNS domains",
			Action: cmd.DomainList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the DNS domain identified by the given id",
			Action: cmd.DomainShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new DNS domain",
			Action: cmd.DomainCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the DNS domain",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account in which the domain shall be registered",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with domain",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a DNS domain",
			Action: cmd.DomainDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "retry",
			Usage:  "Retries the application of DNS domain",
			Action: cmd.DomainRetry,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:        "records",
			Usage:       "Provides information about DNS records",
			Subcommands: append(records.SubCommands()),
		},
	}
}
