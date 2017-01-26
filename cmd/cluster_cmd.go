package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/cluster"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpCluster prepares common resources to send request to Concerto API
func WireUpCluster(c *cli.Context) (cs *cluster.ClusterService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	cs, err = cluster.NewClusterService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up cluster service", err)
	}

	return cs, f
}

// ClusterList subcommand function
func ClusterList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	clusters, err := clusterSvc.GetClusterList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cluster data", err)
	}
	if err = formatter.PrintList(clusters); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ClusterCreate subcommand function
func ClusterCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"name"}, formatter)
	cluster, err := clusterSvc.CreateCluster(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create cluster", err)
	}
	if err = formatter.PrintItem(*cluster); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ClusterDelete subcommand function
func ClusterDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := clusterSvc.DeleteCluster(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete cluster", err)
	}
	return nil
}

// ClusterStart subcommand function
func ClusterStart(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := clusterSvc.StartCluster(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't start cluster", err)
	}
	return nil
}

// ClusterStop subcommand function
func ClusterStop(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := clusterSvc.StopCluster(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't stop cluster", err)
	}
	return nil
}

// ClusterEmpty subcommand function
func ClusterEmpty(c *cli.Context) error {
	debugCmdFuncInfo(c)
	clusterSvc, formatter := WireUpCluster(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := clusterSvc.EmptyCluster(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't empty cluster", err)
	}
	return nil
}
