package node_pools

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns node pool commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing node pools in a cluster",
			Action: cmd.NodePoolList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the node pool identified by the given id",
			Action: cmd.NodePoolShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node pool Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new node pool",
			Action: cmd.NodePoolCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the node pool",
				},
				cli.StringFlag{
					Name:  "subnet-id",
					Usage: "Identifier of the subnet where this node pool is deployed",
				},
				cli.StringFlag{
					Name:  "node-pool-plan-id",
					Usage: "Identifier of the node pool plan that this node pool is based",
				},
				cli.StringFlag{
					Name:  "cpu-type",
					Usage: "Type of CPU each node of the node pools will have. " +
						"Can be nil only if the node pool plan does not have any cpu types",
				},
				cli.IntFlag{
					Name:  "disk-size",
					Usage: "Size of the disk each node of the node pool will have, expressed in Gigabytes (GB)",
				},
				cli.IntFlag{
					Name:  "min-nodes",
					Usage: "Minimum number of nodes the node pool will have",
				},
				cli.IntFlag{
					Name:  "max-nodes",
					Usage: "Maximum number of nodes the node pool will have",
				},
				cli.IntFlag{
					Name:  "desired-nodes",
					Usage: "Amount of nodes the node pool will tend to have if the node pool does not have autoscaling",
				},
				cli.IntFlag{
					Name:  "pods-per-node",
					Usage: "Amount of pods each node of the node pool will have if the node pool plan supports it",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing node pool identified by the given id",
			Action: cmd.NodePoolUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node pool Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the node pool",
				},
				cli.IntFlag{
					Name:  "min-nodes",
					Usage: "Minimum number of nodes the node pool will have",
				},
				cli.IntFlag{
					Name:  "max-nodes",
					Usage: "Maximum number of nodes the node pool will have",
				},
				cli.IntFlag{
					Name:  "desired-nodes",
					Usage: "Amount of nodes the node pool will tend to have if the node pool does not have autoscaling",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a node pool",
			Action: cmd.NodePoolDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node pool Id",
				},
			},
		},
		{
			Name:   "retry",
			Usage:  "Retries the application of node pool identified by the given id",
			Action: cmd.NodePoolRetry,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node pool Id",
				},
			},
		},
		{
			Name:   "show-plan",
			Usage:  "Shows information about a specific node pool plan identified by the given id",
			Action: cmd.NodePoolPlanShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node pool plan Id",
				},
			},
		},
	}
}
