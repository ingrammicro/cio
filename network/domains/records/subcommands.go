// Copyright (c) 2017-2021 Ingram Micro Inc.

package records

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns domains commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all DNS records of a domain",
			Action: cmd.DomainListRecords,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain-id",
					Usage: "Identifier of the DNS domain",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the DNS record identified by the given id",
			Action: cmd.DomainShowRecord,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Identifier of the DNS record",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a DNS record in a domain",
			Action: cmd.DomainCreateRecord,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain-id",
					Usage: "Identifier of the DNS domain",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the DNS record",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Type of the  DNS record, among 'a', 'aaaa', 'cname', 'mx', 'srv', 'txt'",
				},
				cli.StringFlag{
					Name:  "content",
					Usage: "Content of the DNS record",
				},
				cli.IntFlag{
					Name:  "ttl",
					Usage: "TTL of the DNS record. Defaults to 3600 if not provided",
				},
				cli.StringFlag{
					Name: "server-id",
					Usage: "Identifier of the Server that is wanted to be attached to the record. " +
						"Only valid for records of type 'a'",
				},
				cli.StringFlag{
					Name: "floating-ip-id",
					Usage: "Identifier of the floating IP that is wanted to be attached to the record. " +
						"Only valid for records of type 'a'",
				},
				cli.StringFlag{
					Name: "load-balancer-id",
					Usage: "Identifier of the load balancer that is wanted to be attached to the record. " +
						"Only valid for records of type 'cname'",
				},
				cli.IntFlag{
					Name: "priority",
					Usage: "Priority of the record. Only valid for 'mx' and 'srv' types. Defaults to 0. " +
						"Only valid for records of types 'mx' and 'srv'",
				},
				cli.IntFlag{
					Name: "weight",
					Usage: "Weight of the record. Only valid for 'srv' type. Defaults to 0. " +
						"Only valid for records of type 'srv'",
				},
				cli.IntFlag{
					Name: "port",
					Usage: "Port of the record. Only valid for 'srv' type. Defaults to 0. " +
						"Only valid for records of type 'srv'",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates a DNS record in a domain",
			Action: cmd.DomainUpdateRecord,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Identifier of the DNS record",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the DNS record",
				},
				cli.StringFlag{
					Name:  "content",
					Usage: "Content of the DNS record",
				},
				cli.IntFlag{
					Name:  "ttl",
					Usage: "TTL of the DNS record. Defaults to 3600 if not provided",
				},
				cli.IntFlag{
					Name: "priority",
					Usage: "Priority of the record. Only valid for 'mx' and 'srv' types. Defaults to 0. " +
						"Only valid for records of types 'mx' and 'srv'",
				},
				cli.IntFlag{
					Name: "weight",
					Usage: "Weight of the record. Only valid for 'srv' type. Defaults to 0. " +
						"Only valid for records of type 'srv'",
				},
				cli.IntFlag{
					Name: "port",
					Usage: "Port of the record. Only valid for 'srv' type. Defaults to 0. " +
						"Only valid for records of type 'srv'",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a DNS record in a domain",
			Action: cmd.DomainDeleteRecord,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Identifier of the DNS record",
				},
			},
		},
		{
			Name:   "retry",
			Usage:  "Retries the application of DNS record in a domain",
			Action: cmd.DomainRetryRecord,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Identifier of the DNS record",
				},
			},
		},
	}
}
