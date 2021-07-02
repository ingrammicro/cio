// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud_providers

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns cloud providers commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available cloud providers",
			Action: cmd.CloudProviderList,
		},
		{
			Name:   "list-storage-plans",
			Usage:  "This action lists the storage plans offered by the cloud provider identified by the given id",
			Action: cmd.CloudProviderStoragePlansList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud-provider-id",
					Usage: "Cloud provider id",
				},
			},
		},
		{
			Name:   "list-load-balancer-plans",
			Usage:  "This action lists the load balancer plans offered by the cloud provider identified by the given id",
			Action: cmd.CloudProviderLoadBalancerPlansList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud-provider-id",
					Usage: "Cloud provider id",
				},
			},
		},
		{
			Name:   "list-cluster-plans",
			Usage:  "This action lists the cluster plans offered by the cloud provider identified by the given id",
			Action: cmd.CloudProviderClusterPlansList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud-provider-id",
					Usage: "Cloud provider id",
				},
			},
		},
	}
}
