// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"

	"github.com/ingrammicro/cio/api/network"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpLoadBalancer prepares common resources to send request to Concerto API
func WireUpLoadBalancer(c *cli.Context) (ds *network.LoadBalancerService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewLoadBalancerService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up LoadBalancer service", err)
	}

	return ds, f
}

// LoadBalancerList subcommand function
func LoadBalancerList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpLoadBalancer(c)

	loadBalancers, err := svc.ListLoadBalancers()
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancers data", err)
	}

	labelables := make([]types.Labelable, len(loadBalancers))
	for i := 0; i < len(loadBalancers); i++ {
		labelables[i] = types.Labelable(loadBalancers[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	loadBalancers = make([]*types.LoadBalancer, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.LoadBalancer)
		if !ok {
			formatter.PrintFatal(LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.LoadBalancer, got a %T", labelable))
		}
		loadBalancers[i] = v
	}
	if err = formatter.PrintList(loadBalancers); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// LoadBalancerShow subcommand function
func LoadBalancerShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	loadBalancer, err := svc.GetLoadBalancer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// LoadBalancerCreate subcommand function
func LoadBalancerCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpLoadBalancer(c)
	checkRequiredFlags(c, []string{"name", "cloud-account-id", "vpc-id", "plan-id", "realm-id"}, formatter)
	loadBalancerIn := map[string]interface{}{
		"name":                  c.String("name"),
		"cloud_account_id":      c.String("cloud-account-id"),
		"vpc_id":                c.String("vpc-id"),
		"load_balancer_plan_id": c.String("plan-id"),
		"realm_id":              c.String("realm-id"),
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		loadBalancerIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	loadBalancer, err := svc.CreateLoadBalancer(&loadBalancerIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create load balancer", err)
	}

	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// LoadBalancerUpdate subcommand function
func LoadBalancerUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)
	loadBalancerIn := map[string]interface{}{
		"name": c.String("name"),
	}

	loadBalancer, err := svc.UpdateLoadBalancer(c.String("id"), &loadBalancerIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update load balancer", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// LoadBalancerDelete subcommand function
func LoadBalancerDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	loadBalancer, err := svc.DeleteLoadBalancer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete load balancer", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// LoadBalancerRetry subcommand function
func LoadBalancerRetry(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	loadBalancerIn := map[string]interface{}{}

	loadBalancer, err := svc.RetryLoadBalancer(c.String("id"), &loadBalancerIn)
	if err != nil {
		formatter.PrintFatal("Couldn't retry load balancer", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	loadBalancer.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// LoadBalancerPlanShow subcommand function
func LoadBalancerPlanShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	lbp, err := svc.GetLoadBalancerPlan(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer plan data", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	lbp.CloudProviderName = cloudProvidersMap[lbp.CloudProviderID]

	if err = formatter.PrintItem(*lbp); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
