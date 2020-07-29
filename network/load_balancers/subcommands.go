package load_balancers

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/network/load_balancers/certificates"
	"github.com/ingrammicro/cio/network/load_balancers/listeners"
	"github.com/ingrammicro/cio/network/load_balancers/target_groups"
	"github.com/urfave/cli"
)

// SubCommands returns load balancers commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing load balancers",
			Action: cmd.LoadBalancerList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about the load balancer identified by the given id",
			Action: cmd.LoadBalancerShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new load balancer",
			Action: cmd.LoadBalancerCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the load balancer",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account in which is deployed",
				},
				cli.StringFlag{
					Name:  "vpc-id",
					Usage: "Identifier of the VPC in which the load balancer is",
				},
				cli.StringFlag{
					Name:  "plan-id",
					Usage: "Identifier of the load balancer plan",
				},
				cli.StringFlag{
					Name:  "realm-id",
					Usage: "Identifier of the realm in which is deployed",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with load balancer",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing load balancer identified by the given id",
			Action: cmd.LoadBalancerUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the load balancer",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a load balancer",
			Action: cmd.LoadBalancerDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer Id",
				},
			},
		},
		{
			Name:   "retry",
			Usage:  "Retries the application of a load balancer",
			Action: cmd.LoadBalancerRetry,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer Id",
				},
			},
		},
		{
			Name:   "show-plan",
			Usage:  "Shows information about a specific load balancer plan",
			Action: cmd.LoadBalancerPlanShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer plan Id",
				},
			},
		},
		{
			Name:        "target-groups",
			Usage:       "Provides information about load balancer target groups",
			Subcommands: append(target_groups.SubCommands()),
		},
		{
			Name:        "listeners",
			Usage:       "Provides information about load balancer listeners",
			Subcommands: append(listeners.SubCommands()),
		},
		{
			Name:        "certificates",
			Usage:       "Provides information about load balancer certificates",
			Subcommands: append(certificates.SubCommands()),
		},
		{
			Name:   "add-label",
			Usage:  "This action assigns a single label from a single labelable resource",
			Action: cmd.LabelAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "load_balancer",
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
					Usage: "Load balancer Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "load_balancer",
					Hidden: true,
				},
			},
		},
	}
}
