package clusters

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns clusters commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing clusters",
			Action: cmd.ClusterList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about the cluster identified by the given id",
			Action: cmd.ClusterShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new cluster",
			Action: cmd.ClusterCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the cluster",
				},
				cli.StringFlag{
					Name:  "version",
					Usage: "Kubernetes version of the cluster",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account the cluster will be deployed",
				},
				cli.StringFlag{
					Name:  "cluster-plan-id",
					Usage: "Identifier of the cluster plan that will use the cluster to be created",
				},
				cli.StringFlag{
					Name:  "public-access-ip-addresses",
					Usage: "A list of comma separated CIDR blocks the cluster will allow to receive requests",
				},
				cli.BoolFlag{
					Name:  "default-vpc-creation",
					Usage: "Flag indicating if the cluster must create a VPC first",
				},
				cli.StringFlag{
					Name:  "default-vpc-cidr",
					Usage: "CIDR block where the default VPC will have when created",
				},
				cli.StringFlag{
					Name:  "vpc-id",
					Usage: "Identifier of the VPC where the cluster will be deployed",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with cluster",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing cluster identified by the given id",
			Action: cmd.ClusterUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the cluster",
				},
				cli.StringFlag{
					Name:  "version",
					Usage: "Kubernetes version of the cluster",
				},
				cli.StringFlag{
					Name:  "public-access-ip-addresses",
					Usage: "A list of comma separated CIDR blocks the cluster will allow to receive requests",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a cluster",
			Action: cmd.ClusterDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "retry",
			Usage:  "Retries the application of cluster identified by the given id",
			Action: cmd.ClusterRetry,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "discard",
			Usage:  "Discards a cluster but does not delete it from the cloud provider",
			Action: cmd.ClusterDiscard,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "show-plan",
			Usage:  "Shows information about a specific cluster plan identified by the given id",
			Action: cmd.ClusterPlanShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster plan Id",
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
					Usage: "Cluster Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "cluster",
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
					Usage: "Cluster Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "cluster",
					Hidden: true,
				},
			},
		},
	}
}
