// Copyright (c) 2017-2021 Ingram Micro Inc.

package listeners

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns listeners commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all listeners of a load balancer",
			Action: cmd.ListenerList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the listener identified by the given id",
			Action: cmd.ListenerShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Listener Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new listener in a load balancer",
			Action: cmd.ListenerCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer of the listener",
				},
				cli.StringFlag{
					Name:  "protocol",
					Usage: "The protocol of the listener",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "Port of the listener",
				},
				cli.StringFlag{
					Name:  "target-group-id",
					Usage: "Identifier of the target group of the listener",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing listener identified by the given id",
			Action: cmd.ListenerUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Listener Id",
				},
				cli.StringFlag{
					Name:  "target-group-id",
					Usage: "Identifier of the target group of the listener",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a listener",
			Action: cmd.ListenerDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Listener Id",
				},
			},
		},
		{
			Name:   "retry",
			Usage:  "Retries the application of a listener of a load balancer",
			Action: cmd.ListenerRetry,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Listener Id",
				},
			},
		},
		{
			Name:   "list-rules",
			Usage:  "Lists all rules of a listener",
			Action: cmd.ListenerListRules,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "listener-id",
					Usage: "Identifier of the listener",
				},
			},
		},
		{
			Name:   "create-rule",
			Usage:  "Creates a rule in a listener",
			Action: cmd.ListenerCreateRule,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "listener-id",
					Usage: "Identifier of the listener",
				},
				cli.StringFlag{
					Name:  "field",
					Usage: "Field of the rule. It supports the rule fields available in load balancer plan",
				},
				cli.StringFlag{
					Name:  "values",
					Usage: "Values of the rule",
				},
				cli.StringFlag{
					Name:  "target-group-id",
					Usage: "Identifier of the target group of the listener",
				},
			},
		},
		{
			Name:   "update-rule",
			Usage:  "Updates a rule in a target listener",
			Action: cmd.ListenerUpdateRule,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Rule Id",
				},
				cli.StringFlag{
					Name:  "listener-id",
					Usage: "Identifier of the listener",
				},
				cli.StringFlag{
					Name:  "field",
					Usage: "Field of the rule. It supports the rule fields available in load balancer plan",
				},
				cli.StringFlag{
					Name:  "values",
					Usage: "Values of the rule",
				},
			},
		},
		{
			Name:   "delete-rule",
			Usage:  "Destroys a rule in a listener",
			Action: cmd.ListenerDeleteRule,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Rule Id",
				},
				cli.StringFlag{
					Name:  "listener-id",
					Usage: "Identifier of the listener",
				},
			},
		},
	}
}
