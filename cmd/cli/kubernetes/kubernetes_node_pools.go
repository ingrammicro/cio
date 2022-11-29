// Copyright (c) 2017-2022 Ingram Micro Inc.

package kubernetes

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Node pool Id"}

	fClusterId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Cluster Id"}

	fNodePoolPlanIdReq := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Node pool plan Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Logical name of the node pool"}
	fNameReq := fName
	fNameReq.Required = true

	fSubnetId := cmd.FlagContext{Type: cmd.String, Name: cmd.SubnetId,
		Usage: "Identifier of the subnet where this node pool is deployed"}

	fNodePoolPlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.NodePoolPlanId, Required: true,
		Usage: "Identifier of the node pool plan that this node pool is based"}

	fCpuType := cmd.FlagContext{Type: cmd.String, Name: cmd.CpuType,
		Usage: "Type of CPU each node of the node pools will have. " +
			"Can be nil only if the node pool plan does not have any cpu types"}

	fDiskSize := cmd.FlagContext{Type: cmd.Int, Name: cmd.DiskSize,
		Usage: "Size of the disk each node of the node pool will have, expressed in Gigabytes (GB)"}

	fMinNodes := cmd.FlagContext{Type: cmd.Int, Name: cmd.MinNodes,
		Usage: "Minimum number of nodes the node pool will have"}

	fMaxNodes := cmd.FlagContext{Type: cmd.Int, Name: cmd.MaxNodes,
		Usage: "Maximum number of nodes the node pool will have"}

	fDesiredNodes := cmd.FlagContext{Type: cmd.Int, Name: cmd.DesiredNodes,
		Usage: "Amount of nodes the node pool will tend to have if the node pool does not have autoscaling"}

	fPodsPerNode := cmd.FlagContext{Type: cmd.Int, Name: cmd.PodsPerNode,
		Usage: "Amount of pods each node of the node pool will have if the node pool plan supports it"}

	nodePoolsCmd := cmd.NewCommand(kubernetesCmd, &cmd.CommandContext{
		Use:   "node-pools",
		Short: "Provides information on kubernetes node pools"},
	)
	cmd.NewCommand(nodePoolsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all existing node pools in a cluster",
		RunMethod:    NodePoolList,
		FlagContexts: []cmd.FlagContext{fClusterId}},
	)
	cmd.NewCommand(nodePoolsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the node pool identified by the given id",
		RunMethod:    NodePoolShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(nodePoolsCmd, &cmd.CommandContext{
		Use:       "create",
		Short:     "Creates a new node pool",
		RunMethod: NodePoolCreate,
		FlagContexts: []cmd.FlagContext{fClusterId, fNameReq, fSubnetId, fNodePoolPlanId, fCpuType, fDiskSize,
			fMinNodes, fMaxNodes, fDesiredNodes, fPodsPerNode}},
	)
	cmd.NewCommand(nodePoolsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing node pool identified by the given id",
		RunMethod:    NodePoolUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName, fMinNodes, fMaxNodes, fDesiredNodes}},
	)
	cmd.NewCommand(nodePoolsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a node pool",
		RunMethod:    NodePoolDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(nodePoolsCmd, &cmd.CommandContext{
		Use:          "retry",
		Short:        "Retries the application of node pool identified by the given id",
		RunMethod:    NodePoolRetry,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(nodePoolsCmd, &cmd.CommandContext{
		Use:          "show-plan",
		Short:        "Shows information about a specific node pool plan identified by the given id",
		RunMethod:    NodePoolPlanShow,
		FlagContexts: []cmd.FlagContext{fNodePoolPlanIdReq}},
	)
}

// NodePoolList subcommand function
func NodePoolList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	nodePools, err := svc.ListNodePools(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive node pools data", err)
		return err
	}

	if err = formatter.PrintList(nodePools); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// NodePoolShow subcommand function
func NodePoolShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	nodePool, err := svc.GetNodePool(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive node pool data", err)
		return err
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// NodePoolCreate subcommand function
func NodePoolCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	nodePoolIn := map[string]interface{}{
		"name":              viper.GetString(cmd.Name),
		"node_pool_plan_id": viper.GetString(cmd.NodePoolPlanId),
	}

	cmd.SetParamString("subnet_id", cmd.SubnetId, nodePoolIn)
	cmd.SetParamString("cpu_type", cmd.CpuType, nodePoolIn)
	cmd.SetParamInt("disk_size", cmd.DiskSize, nodePoolIn)
	cmd.SetParamInt("min_nodes", cmd.MinNodes, nodePoolIn)
	cmd.SetParamInt("max_nodes", cmd.MaxNodes, nodePoolIn)
	cmd.SetParamInt("desired_nodes", cmd.DesiredNodes, nodePoolIn)
	cmd.SetParamInt("pods_per_node", cmd.PodsPerNode, nodePoolIn)

	nodePool, err := svc.CreateNodePool(cmd.GetContext(), viper.GetString(cmd.Id), &nodePoolIn)
	if err != nil {
		formatter.PrintError("Couldn't create node pool", err)
		return err
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// NodePoolUpdate subcommand function
func NodePoolUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	nodePoolIn := map[string]interface{}{}
	cmd.SetParamString("name", cmd.Name, nodePoolIn)
	cmd.SetParamInt("min_nodes", cmd.MinNodes, nodePoolIn)
	cmd.SetParamInt("max_nodes", cmd.MaxNodes, nodePoolIn)
	cmd.SetParamInt("desired_nodes", cmd.DesiredNodes, nodePoolIn)

	nodePool, err := svc.UpdateNodePool(cmd.GetContext(), viper.GetString(cmd.Id), &nodePoolIn)
	if err != nil {
		formatter.PrintError("Couldn't update node pool", err)
		return err
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// NodePoolDelete subcommand function
func NodePoolDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	nodePool, err := svc.DeleteNodePool(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete node pool", err)
		return err
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// NodePoolRetry subcommand function
func NodePoolRetry() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	nodePool, err := svc.RetryNodePool(cmd.GetContext(), viper.GetString(cmd.Id), &map[string]interface{}{})
	if err != nil {
		formatter.PrintError("Couldn't retry node pool", err)
		return err
	}

	if err = formatter.PrintItem(*nodePool); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// NodePoolPlanShow subcommand function
func NodePoolPlanShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	nodePoolPlan, err := svc.GetNodePoolPlan(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive node pool plan", err)
		return err
	}

	if err = formatter.PrintItem(*nodePoolPlan); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
