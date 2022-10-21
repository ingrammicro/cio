// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var targetGroupsCmd *cobra.Command
var listenersCmd *cobra.Command
var certificatesCmd *cobra.Command

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Load balancer Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the load balancer"}

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account in which is deployed"}

	fVpcId := cmd.FlagContext{Type: cmd.String, Name: cmd.VpcId, Required: true,
		Usage: "Identifier of the VPC in which the load balancer is"}

	fPlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.PlanId, Required: true,
		Usage: "Identifier of the load balancer plan"}

	fRealmId := cmd.FlagContext{Type: cmd.String, Name: cmd.RealmId, Required: true,
		Usage: "Identifier of the realm in which is deployed"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with load balancer"}

	fIdLoadBalancerPlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true,
		Usage: "Load balancer plan Id"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{
		Type:         cmd.String,
		Name:         cmd.ResourceType,
		DefaultValue: "load_balancer",
		Hidden:       true,
		Usage:        "Resource Type",
	}

	loadBalancersCmd := cmd.NewCommand(networkCmd, &cmd.CommandContext{
		Use:   "load-balancers",
		Short: "Provides information about load balancers"},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists all existing load balancers",
		RunMethod: LoadBalancerList},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the load balancer identified by the given id",
		RunMethod:    LoadBalancerShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new load balancer",
		RunMethod:    LoadBalancerCreate,
		FlagContexts: []cmd.FlagContext{fName, fCloudAccountId, fVpcId, fPlanId, fRealmId, fLabels}},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing load balancer identified by the given id",
		RunMethod:    LoadBalancerUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName}},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a load balancer",
		RunMethod:    LoadBalancerDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:          "retry",
		Short:        "Retries the application of a load balancer",
		RunMethod:    LoadBalancerRetry,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:          "show-plan",
		Short:        "Shows information about a specific load balancer plan",
		RunMethod:    LoadBalancerPlanShow,
		FlagContexts: []cmd.FlagContext{fIdLoadBalancerPlanId}},
	)

	targetGroupsCmd = cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:   "target-groups",
		Short: "Provides information about load balancer target groups"},
	)
	listenersCmd = cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:   "listeners",
		Short: "Provides information about load balancer listeners"},
	)
	certificatesCmd = cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:   "certificates",
		Short: "Provides information about load balancer certificates"},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(loadBalancersCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// LoadBalancerList subcommand function
func LoadBalancerList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	loadBalancers, err := svc.ListLoadBalancers(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive load balancers data", err)
		return err
	}

	labelables := make([]types.Labelable, len(loadBalancers))
	for i := 0; i < len(loadBalancers); i++ {
		labelables[i] = types.Labelable(loadBalancers[i])
	}
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	filteredLabelables, err := labels.LabelFiltering(labelables, labelIDsByName)
	if err != nil {
		return err
	}
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	loadBalancers = make([]*types.LoadBalancer, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.LoadBalancer)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.LoadBalancer, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		loadBalancers[i] = v
	}
	if err = formatter.PrintList(loadBalancers); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// LoadBalancerShow subcommand function
func LoadBalancerShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	loadBalancer, err := svc.GetLoadBalancer(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer data", err)
		return err
	}
	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// LoadBalancerCreate subcommand function
func LoadBalancerCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	loadBalancerIn := map[string]interface{}{
		"name":                  viper.GetString(cmd.Name),
		"cloud_account_id":      viper.GetString(cmd.CloudAccountId),
		"vpc_id":                viper.GetString(cmd.VpcId),
		"load_balancer_plan_id": viper.GetString(cmd.PlanId),
		"realm_id":              viper.GetString(cmd.RealmId),
	}

	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	if viper.IsSet(cmd.Labels) {
		loadBalancerIn["label_ids"], err = labels.LabelResolution(
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	loadBalancer, err := svc.CreateLoadBalancer(cmd.GetContext(), &loadBalancerIn)
	if err != nil {
		formatter.PrintError("Couldn't create load balancer", err)
		return err
	}

	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// LoadBalancerUpdate subcommand function
func LoadBalancerUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	loadBalancerIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	loadBalancer, err := svc.UpdateLoadBalancer(cmd.GetContext(), viper.GetString(cmd.Id), &loadBalancerIn)
	if err != nil {
		formatter.PrintError("Couldn't update load balancer", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// LoadBalancerDelete subcommand function
func LoadBalancerDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	loadBalancer, err := svc.DeleteLoadBalancer(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete load balancer", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// LoadBalancerRetry subcommand function
func LoadBalancerRetry() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	loadBalancerIn := map[string]interface{}{}

	loadBalancer, err := svc.RetryLoadBalancer(cmd.GetContext(), viper.GetString(cmd.Id), &loadBalancerIn)
	if err != nil {
		formatter.PrintError("Couldn't retry load balancer", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// LoadBalancerPlanShow subcommand function
func LoadBalancerPlanShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	lbp, err := svc.GetLoadBalancerPlan(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer plan data", err)
		return err
	}

	if err = formatter.PrintItem(*lbp); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
