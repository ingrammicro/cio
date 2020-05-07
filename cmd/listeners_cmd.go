package cmd

import (
	"github.com/ingrammicro/cio/api/network"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
	"strings"
)

// WireUpListener prepares common resources to send request to Concerto API
func WireUpListener(c *cli.Context) (ds *network.ListenerService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewListenerService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up Listener service", err)
	}

	return ds, f
}

// ListenerList subcommand function
func ListenerList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"load-balancer-id"}, formatter)
	listeners, err := svc.ListListeners(c.String("load-balancer-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer listeners data", err)
	}
	if err = formatter.PrintList(listeners); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ListenerShow subcommand function
func ListenerShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	listener, err := svc.GetListener(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive load balancer listener data", err)
	}
	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ListenerCreate subcommand function
func ListenerCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"load-balancer-id", "protocol", "port", "target-group-id"}, formatter)
	listenerIn := map[string]interface{}{
		"protocol":                c.String("protocol"),
		"port":                    c.Int("port"),
		"default_target_group_id": c.String("target-group-id"),
	}

	listener, err := svc.CreateListener(c.String("load-balancer-id"), &listenerIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create load balancer listener", err)
	}

	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ListenerUpdate subcommand function
func ListenerUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"id", "target-group-id"}, formatter)
	listenerIn := map[string]interface{}{
		"default_target_group_id": c.String("target-group-id"),
	}

	listener, err := svc.UpdateListener(c.String("id"), &listenerIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update load balancer listener", err)
	}

	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ListenerDelete subcommand function
func ListenerDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := svc.DeleteListener(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete load balancer listener", err)
	}
	// @TODO wait while decommissioning?
	return nil
}

// ListenerRetry subcommand function
func ListenerRetry(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	listenerIn := map[string]interface{}{}

	listener, err := svc.RetryListener(c.String("id"), &listenerIn)
	if err != nil {
		formatter.PrintFatal("Couldn't retry load balancer listener", err)
	}

	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ListenerListRules subcommand function
func ListenerListRules(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"listener-id"}, formatter)
	rules, err := svc.ListRules(c.String("listener-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive listener rules data", err)
	}
	if err = formatter.PrintList(rules); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ListenerCreateRule subcommand function
func ListenerCreateRule(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"listener-id", "field", "values", "target-group-id"}, formatter)
	ruleIn := map[string]interface{}{
		"field":           c.String("field"),
		"values":          utils.RemoveDuplicates(strings.Split(c.String("values"), ",")),
		"target_group_id": c.String("target-group-id"),
	}

	rule, err := svc.CreateRule(c.String("listener-id"), &ruleIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create listener rule", err)
	}

	if err = formatter.PrintItem(*rule); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ListenerUpdateRule subcommand function
func ListenerUpdateRule(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"id", "listener-id", "field", "values"}, formatter)
	ruleIn := map[string]interface{}{
		"field":  c.String("field"),
		"values": utils.RemoveDuplicates(strings.Split(c.String("values"), ",")),
	}

	rule, err := svc.UpdateRule(c.String("listener-id"), c.String("id"), &ruleIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update listener rule", err)
	}

	if err = formatter.PrintItem(*rule); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ListenerDeleteRule subcommand function
func ListenerDeleteRule(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpListener(c)

	checkRequiredFlags(c, []string{"id", "listener-id"}, formatter)

	err := svc.DeleteRule(c.String("listener-id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete listener rule", err)
	}
	return nil
}
