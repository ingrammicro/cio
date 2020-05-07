package cmd

import (
	"github.com/ingrammicro/cio/api/network"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpTargetGroup prepares common resources to send request to Concerto API
func WireUpTargetGroup(c *cli.Context) (ds *network.TargetGroupService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewTargetGroupService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up TargetGroup service", err)
	}

	return ds, f
}

// TargetGroupList subcommand function
func TargetGroupList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"load-balancer-id"}, formatter)
	targetGroups, err := svc.ListTargetGroups(c.String("load-balancer-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer target groups data", err)
	}
	if err = formatter.PrintList(targetGroups); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TargetGroupShow subcommand function
func TargetGroupShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	lb, err := svc.GetTargetGroup(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer target group data", err)
	}
	if err = formatter.PrintItem(*lb); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TargetGroupCreate subcommand function
func TargetGroupCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"load-balancer-id", "name", "protocol", "port", "health-check-protocol", "health-check-port", "health-check-interval", "health-check-threshold-count", "health-check-path"}, formatter)
	targetGroupIn := map[string]interface{}{
		"name":                         c.String("name"),
		"protocol":                     c.String("protocol"),
		"port":                         c.Int("port"),
		"health_check_protocol":        c.String("health-check-protocol"),
		"health_check_port":            c.Int("health-check-port"),
		"health_check_interval":        c.Int("health-check-interval"),
		"health_check_threshold_count": c.Int("health-check-threshold-count"),
		"health_check_path":            c.String("health-check-path"),
	}
	if c.IsSet("stickiness") {
		targetGroupIn["stickiness"] = c.Bool("stickiness")
	}

	targetGroup, err := svc.CreateTargetGroup(c.String("load-balancer-id"), &targetGroupIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create load balancer target group", err)
	}

	if err = formatter.PrintItem(*targetGroup); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TargetGroupUpdate subcommand function
func TargetGroupUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"id", "name", "protocol", "port", "health-check-protocol", "health-check-port", "health-check-interval", "health-check-threshold-count", "health-check-path"}, formatter)
	targetGroupIn := map[string]interface{}{
		"name":                         c.String("name"),
		"protocol":                     c.String("protocol"),
		"port":                         c.Int("port"),
		"health_check_protocol":        c.String("health-check-protocol"),
		"health_check_port":            c.Int("health-check-port"),
		"health_check_interval":        c.Int("health-check-interval"),
		"health_check_threshold_count": c.Int("health-check-threshold-count"),
		"health_check_path":            c.String("health-check-path"),
	}
	if c.IsSet("stickiness") {
		targetGroupIn["stickiness"] = c.Bool("stickiness")
	}

	targetGroup, err := svc.UpdateTargetGroup(c.String("id"), &targetGroupIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update load balancer target group", err)
	}

	if err = formatter.PrintItem(*targetGroup); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TargetGroupDelete subcommand function
func TargetGroupDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	targetGroup, err := svc.DeleteTargetGroup(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete load balancer target group", err)
	}

	if err = formatter.PrintItem(*targetGroup); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	// @TODO wait while decommissioning?
	return nil
}

// TargetGroupRetry subcommand function
func TargetGroupRetry(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	targetGroupIn := map[string]interface{}{}

	targetGroup, err := svc.RetryTargetGroup(c.String("id"), &targetGroupIn)
	if err != nil {
		formatter.PrintFatal("Couldn't retry load balancer target group", err)
	}

	if err = formatter.PrintItem(*targetGroup); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TargetGroupListTargets subcommand function
func TargetGroupListTargets(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"target-group-id"}, formatter)
	targets, err := svc.ListTargets(c.String("target-group-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive targets data", err)
	}
	if err = formatter.PrintList(targets); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TargetGroupCreateTarget subcommand function
func TargetGroupCreateTarget(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"target-group-id", "resource-type", "resource-id"}, formatter)
	targetIn := map[string]interface{}{
		"id":            c.String("resource-id"),
		"resource_type": c.String("resource-type"),
	}

	target, err := svc.CreateTarget(c.String("target-group-id"), &targetIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create target", err)
	}

	if err = formatter.PrintItem(*target); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TargetGroupDeleteTarget subcommand function
func TargetGroupDeleteTarget(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpTargetGroup(c)

	checkRequiredFlags(c, []string{"target-group-id", "resource-type", "resource-id"}, formatter)

	err := svc.DeleteTarget(c.String("target-group-id"), c.String("resource-type"), c.String("resource-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete target", err)
	}
	return nil
}
