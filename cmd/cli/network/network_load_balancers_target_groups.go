// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fLoadBalancerId := cmd.FlagContext{Type: cmd.String, Name: cmd.LoadBalancerId, Required: true,
		Usage: "Identifier of the load balancer of the target group"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Target group Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the target group"}

	fProtocol := cmd.FlagContext{Type: cmd.String, Name: cmd.Protocol, Required: true,
		Usage: "The protocol of the target group"}

	fPort := cmd.FlagContext{Type: cmd.Int, Name: cmd.Port, Required: true, Usage: "Port of the target group"}

	fHealthCheckProtocol := cmd.FlagContext{Type: cmd.String, Name: cmd.HealthCheckProtocol, Required: true,
		Usage: "The protocol of the health check of the target group"}

	fHealthCheckPort := cmd.FlagContext{Type: cmd.Int, Name: cmd.HealthCheckPort, Required: true,
		Usage: "Port of the health check of the target group"}

	fHealthCheckInterval := cmd.FlagContext{Type: cmd.Int, Name: cmd.HealthCheckInterval, Required: true,
		Usage: "Interval of the health check of the target group"}

	fHealthCheckThresholdCount := cmd.FlagContext{Type: cmd.Int, Name: cmd.HealthCheckThresholdCount, Required: true,
		Usage: "Threshold count of the health check of the target group"}

	fHealthCheckPath := cmd.FlagContext{Type: cmd.String, Name: cmd.HealthCheckPath, Required: true,
		Usage: "Path of the health check of the target group"}

	fStickiness := cmd.FlagContext{Type: cmd.Bool, Name: cmd.Stickiness,
		Usage: "Flag to indicate whether requests from the same origin must be redirected to the same target"}

	fTargetGroupId := cmd.FlagContext{Type: cmd.String, Name: cmd.TargetGroupId, Required: true,
		Usage: "Target group Id"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, Required: true,
		Usage: "The identifier for the type of resource, specifically \"server\" or \"server_array\""}

	fResourceId := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceId, Required: true,
		Usage: "The identifier for the target resource"}

	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all target groups of a load balancer",
		RunMethod:    TargetGroupList,
		FlagContexts: []cmd.FlagContext{fLoadBalancerId}},
	)
	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the target group identified by the given id",
		RunMethod:    TargetGroupShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:       "create",
		Short:     "Creates a new target group in a load balancer",
		RunMethod: TargetGroupCreate,
		FlagContexts: []cmd.FlagContext{fLoadBalancerId, fName, fProtocol, fPort, fHealthCheckProtocol,
			fHealthCheckPort, fHealthCheckInterval, fHealthCheckThresholdCount, fHealthCheckPath, fStickiness}},
	)
	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:       "update",
		Short:     "Updates an existing target group identified by the given id",
		RunMethod: TargetGroupUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName, fProtocol, fPort, fHealthCheckProtocol, fHealthCheckPort,
			fHealthCheckInterval, fHealthCheckThresholdCount, fHealthCheckPath, fStickiness}},
	)
	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a target group",
		RunMethod:    TargetGroupDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:          "retry",
		Short:        "Retries the application of a target group of a load balancer",
		RunMethod:    TargetGroupRetry,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:          "list-targets",
		Short:        "Lists all targets in a target group",
		RunMethod:    TargetGroupListTargets,
		FlagContexts: []cmd.FlagContext{fTargetGroupId}},
	)
	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:          "create-target",
		Short:        "Creates a target in a target group",
		RunMethod:    TargetGroupCreateTarget,
		FlagContexts: []cmd.FlagContext{fTargetGroupId, fResourceType, fResourceId}},
	)
	cmd.NewCommand(targetGroupsCmd, &cmd.CommandContext{
		Use:          "delete-target",
		Short:        "Destroys a target in a target group",
		RunMethod:    TargetGroupDeleteTarget,
		FlagContexts: []cmd.FlagContext{fTargetGroupId, fResourceType, fResourceId}},
	)
}

// TargetGroupList subcommand function
func TargetGroupList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	targetGroups, err := svc.ListTargetGroups(cmd.GetContext(), viper.GetString(cmd.LoadBalancerId))
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer target groups data", err)
		return err
	}
	if err = formatter.PrintList(targetGroups); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// TargetGroupShow subcommand function
func TargetGroupShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	lb, err := svc.GetTargetGroup(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer target group data", err)
		return err
	}
	if err = formatter.PrintItem(*lb); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

func getTargetGroupIn() map[string]interface{} {
	targetGroupIn := map[string]interface{}{
		"name":                         viper.GetString(cmd.Name),
		"protocol":                     viper.GetString(cmd.Protocol),
		"port":                         viper.GetInt(cmd.Port),
		"health_check_protocol":        viper.GetString(cmd.HealthCheckProtocol),
		"health_check_port":            viper.GetInt(cmd.HealthCheckPort),
		"health_check_interval":        viper.GetInt(cmd.HealthCheckInterval),
		"health_check_threshold_count": viper.GetInt(cmd.HealthCheckThresholdCount),
		"health_check_path":            viper.GetString(cmd.HealthCheckPath),
	}
	cmd.SetParamBool("stickiness", cmd.Stickiness, targetGroupIn)
	return targetGroupIn
}

// TargetGroupCreate subcommand function
func TargetGroupCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	targetGroupIn := getTargetGroupIn()
	targetGroup, err := svc.CreateTargetGroup(cmd.GetContext(), viper.GetString(cmd.LoadBalancerId), &targetGroupIn)
	if err != nil {
		formatter.PrintError("Couldn't create load balancer target group", err)
		return err
	}

	if err = formatter.PrintItem(*targetGroup); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// TargetGroupUpdate subcommand function
func TargetGroupUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	targetGroupIn := getTargetGroupIn()
	targetGroup, err := svc.UpdateTargetGroup(cmd.GetContext(), viper.GetString(cmd.Id), &targetGroupIn)
	if err != nil {
		formatter.PrintError("Couldn't update load balancer target group", err)
		return err
	}

	if err = formatter.PrintItem(*targetGroup); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// TargetGroupDelete subcommand function
func TargetGroupDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	targetGroup, err := svc.DeleteTargetGroup(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete load balancer target group", err)
		return err
	}

	if err = formatter.PrintItem(*targetGroup); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// TargetGroupRetry subcommand function
func TargetGroupRetry() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	targetGroupIn := map[string]interface{}{}
	targetGroup, err := svc.RetryTargetGroup(cmd.GetContext(), viper.GetString(cmd.Id), &targetGroupIn)
	if err != nil {
		formatter.PrintError("Couldn't retry load balancer target group", err)
		return err
	}

	if err = formatter.PrintItem(*targetGroup); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// TargetGroupListTargets subcommand function
func TargetGroupListTargets() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	targets, err := svc.ListTargets(cmd.GetContext(), viper.GetString(cmd.TargetGroupId))
	if err != nil {
		formatter.PrintError("Couldn't receive targets data", err)
		return err
	}
	if err = formatter.PrintList(targets); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// TargetGroupCreateTarget subcommand function
func TargetGroupCreateTarget() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	targetIn := map[string]interface{}{
		"id":            viper.GetString(cmd.ResourceId),
		"resource_type": viper.GetString(cmd.ResourceType),
	}

	target, err := svc.CreateTarget(cmd.GetContext(), viper.GetString(cmd.TargetGroupId), &targetIn)
	if err != nil {
		formatter.PrintError("Couldn't create target", err)
		return err
	}

	if err = formatter.PrintItem(*target); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// TargetGroupDeleteTarget subcommand function
func TargetGroupDeleteTarget() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteTarget(cmd.GetContext(),
		viper.GetString(cmd.TargetGroupId),
		viper.GetString(cmd.ResourceType),
		viper.GetString(cmd.ResourceId),
	)
	if err != nil {
		formatter.PrintError("Couldn't delete target", err)
		return err
	}
	return nil
}
