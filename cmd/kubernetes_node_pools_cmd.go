// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"github.com/ingrammicro/cio/api/kubernetes"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpNodePool prepares common resources to send request to Concerto API
func WireUpNodePool(c *cli.Context) (ds *kubernetes.NodePoolService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = kubernetes.NewNodePoolService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up node pool service", err)
	}

	return ds, f
}

// NodePoolList subcommand function
func NodePoolList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	nodePoolSvc, formatter := WireUpNodePool(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	nodePools, err := nodePoolSvc.ListNodePools(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive node pool data", err)
	}

	if err = formatter.PrintList(nodePools); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// NodePoolShow subcommand function
func NodePoolShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	nodePoolSvc, formatter := WireUpNodePool(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	nodePool, err := nodePoolSvc.GetNodePool(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive node pool data", err)
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// NodePoolCreate subcommand function
func NodePoolCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	nodePoolSvc, formatter := WireUpNodePool(c)

	checkRequiredFlags(c, []string{"id", "name", "node-pool-plan-id"}, formatter)
	nodePoolIn := map[string]interface{}{
		"name":              c.String("name"),
		"node_pool_plan_id": c.String("node-pool-plan-id"),
	}

	if c.IsSet("subnet-id") {
		nodePoolIn["subnet_id"] = c.String("subnet-id")
	}
	if c.IsSet("cpu-type") {
		nodePoolIn["cpu_type"] = c.String("cpu-type")
	}
	if c.IsSet("disk-size") {
		nodePoolIn["disk_size"] = c.Int("disk-size")
	}
	if c.IsSet("min-nodes") {
		nodePoolIn["min_nodes"] = c.Int("min-nodes")
	}
	if c.IsSet("max-nodes") {
		nodePoolIn["max_nodes"] = c.Int("max-nodes")
	}
	if c.IsSet("desired-nodes") {
		nodePoolIn["desired_nodes"] = c.Int("desired-nodes")
	}
	if c.IsSet("pods-per-node") {
		nodePoolIn["pods_per_node"] = c.Int("pods-per-node")
	}

	nodePool, err := nodePoolSvc.CreateNodePool(c.String("id"), &nodePoolIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create node pool", err)
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// NodePoolUpdate subcommand function
func NodePoolUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	nodePoolSvc, formatter := WireUpNodePool(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	nodePoolIn := map[string]interface{}{}
	if c.IsSet("name") {
		nodePoolIn["name"] = c.String("name")
	}
	if c.IsSet("min-nodes") {
		nodePoolIn["min_nodes"] = c.String("min-nodes")
	}
	if c.IsSet("max-nodes") {
		nodePoolIn["max_nodes"] = c.String("max-nodes")
	}
	if c.IsSet("desired-nodes") {
		nodePoolIn["desired_nodes"] = c.String("desired-nodes")
	}

	nodePool, err := nodePoolSvc.UpdateNodePool(c.String("id"), &nodePoolIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update node pool", err)
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// NodePoolDelete subcommand function
func NodePoolDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	nodePoolSvc, formatter := WireUpNodePool(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	nodePool, err := nodePoolSvc.DeleteNodePool(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete node pool", err)
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// NodePoolRetry subcommand function
func NodePoolRetry(c *cli.Context) error {
	debugCmdFuncInfo(c)
	nodePoolSvc, formatter := WireUpNodePool(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	nodePool, err := nodePoolSvc.RetryNodePool(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't retry node pool", err)
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// NodePoolPlanShow subcommand function
func NodePoolPlanShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	nodePoolSvc, formatter := WireUpNodePool(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	nodePoolPlan, err := nodePoolSvc.GetNodePoolPlan(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't show node pool plan", err)
	}

	if err = formatter.PrintItem(*nodePoolPlan); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
