// Copyright (c) 2017-2021 Ingram Micro Inc.

package kubernetes

import (
	"github.com/ingrammicro/cio/kubernetes/clusters"
	"github.com/ingrammicro/cio/kubernetes/node_pools"
	"github.com/urfave/cli"
)

// SubCommands returns kubernetes commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "clusters",
			Usage:       "Provides information on kubernetes clusters",
			Subcommands: append(clusters.SubCommands()),
		},
		{
			Name:        "node-pools",
			Usage:       "Provides information on kubernetes node pools",
			Subcommands: append(node_pools.SubCommands()),
		},
	}
}
