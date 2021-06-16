// Copyright (c) 2017-2021 Ingram Micro Inc.

package target_groups

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns target groups commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all target groups of a load balancer",
			Action: cmd.TargetGroupList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the target group identified by the given id",
			Action: cmd.TargetGroupShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Target group Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new target group in a load balancer",
			Action: cmd.TargetGroupCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "load-balancer-id",
					Usage: "Identifier of the load balancer of the target group",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the target group",
				},
				cli.StringFlag{
					Name:  "protocol",
					Usage: "The protocol of the target group",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "Port of the target group",
				},
				cli.StringFlag{
					Name:  "health-check-protocol",
					Usage: "The protocol of the health check of the target group",
				},
				cli.IntFlag{
					Name:  "health-check-port",
					Usage: "Port of the health check of the target group",
				},
				cli.IntFlag{
					Name:  "health-check-interval",
					Usage: "Interval of the health check of the target group",
				},
				cli.IntFlag{
					Name:  "health-check-threshold-count",
					Usage: "Threshold count of the health check of the target group",
				},
				cli.StringFlag{
					Name:  "health-check-path",
					Usage: "Path of the health check of the target group",
				},
				cli.BoolFlag{
					Name:  "stickiness",
					Usage: "Flag to indicate whether requests from the same origin must be redirected to the same target",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing target group identified by the given id",
			Action: cmd.TargetGroupUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Target group Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the target group",
				},
				cli.StringFlag{
					Name:  "protocol",
					Usage: "The protocol of the target group",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "Port of the target group",
				},
				cli.StringFlag{
					Name:  "health-check-protocol",
					Usage: "The protocol of the health check of the target group",
				},
				cli.IntFlag{
					Name:  "health-check-port",
					Usage: "Port of the health check of the target group",
				},
				cli.IntFlag{
					Name:  "health-check-interval",
					Usage: "Interval of the health check of the target group",
				},
				cli.IntFlag{
					Name:  "health-check-threshold-count",
					Usage: "Threshold count of the health check of the target group",
				},
				cli.StringFlag{
					Name:  "health-check-path",
					Usage: "Path of the health check of the target group",
				},
				cli.BoolFlag{
					Name:  "stickiness",
					Usage: "Flag to indicate whether requests from the same origin must be redirected to the same target",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a target group",
			Action: cmd.TargetGroupDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Target group Id",
				},
			},
		},
		{
			Name:   "retry",
			Usage:  "Retries the application of a target group of a load balancer",
			Action: cmd.TargetGroupRetry,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Target group Id",
				},
			},
		},
		{
			Name:   "list-targets",
			Usage:  "Lists all targets in a target group",
			Action: cmd.TargetGroupListTargets,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "target-group-id",
					Usage: "Identifier of the target group",
				},
			},
		},
		{
			Name:   "create-target",
			Usage:  "Creates a target in a target group",
			Action: cmd.TargetGroupCreateTarget,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "target-group-id",
					Usage: "Identifier of the target group",
				},
				cli.StringFlag{
					Name:  "resource-type",
					Usage: "The identifier for the type of resource, specifically \"server\" or \"server_array\"",
				},
				cli.StringFlag{
					Name:  "resource-id",
					Usage: "The identifier for the target resource",
				},
			},
		},
		{
			Name:   "delete-target",
			Usage:  "Destroys a target in a target group",
			Action: cmd.TargetGroupDeleteTarget,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "target-group-id",
					Usage: "Identifier of the target group",
				},
				cli.StringFlag{
					Name:  "resource-type",
					Usage: "The identifier for the type of resource, specifically \"server\" or \"server_array\"",
				},
				cli.StringFlag{
					Name:  "resource-id",
					Usage: "The identifier for the target resource",
				},
			},
		},
	}
}
