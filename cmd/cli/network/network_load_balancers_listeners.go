// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

import (
	"github.com/ingrammicro/cio/cmd/cli"
	"strings"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/utils"
	"github.com/spf13/viper"
)

func init() {
	fLoadBalancerId := cmd.FlagContext{Type: cmd.String, Name: cmd.LoadBalancerId, Required: true,
		Usage: "Identifier of the load balancer"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Listener Id"}

	fProtocol := cmd.FlagContext{Type: cmd.String, Name: cmd.Protocol, Required: true,
		Usage: "The protocol of the listener"}

	fPort := cmd.FlagContext{Type: cmd.Int, Name: cmd.Port, Required: true, Usage: "Port of the listener"}

	fTargetGroupId := cmd.FlagContext{Type: cmd.String, Name: cmd.TargetGroupId, Required: true,
		Usage: "Identifier of the target group of the listener"}

	fCertificateId := cmd.FlagContext{Type: cmd.String, Name: cmd.CertificateId,
		Usage: "Identifier of the certificate"}

	fListenerId := cmd.FlagContext{Type: cmd.String, Name: cmd.ListenerId, Required: true,
		Usage: "Identifier of the listener"}

	fField := cmd.FlagContext{Type: cmd.String, Name: cmd.Field, Required: true,
		Usage: "Field of the rule. It supports the rule fields available in load balancer plan"}

	fValues := cmd.FlagContext{Type: cmd.String, Name: cmd.Values, Required: true, Usage: "Values of the rule"}

	fIdRule := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Rule Id"}

	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all listeners of a load balancer",
		RunMethod:    ListenerList,
		FlagContexts: []cmd.FlagContext{fLoadBalancerId}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the listener identified by the given id",
		RunMethod:    ListenerShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new listener in a load balancer",
		RunMethod:    ListenerCreate,
		FlagContexts: []cmd.FlagContext{fLoadBalancerId, fProtocol, fPort, fTargetGroupId, fCertificateId}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing listener identified by the given id",
		RunMethod:    ListenerUpdate,
		FlagContexts: []cmd.FlagContext{fId, fTargetGroupId, fCertificateId}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a listener",
		RunMethod:    ListenerDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "retry",
		Short:        "Retries the application of a listener of a load balancer",
		RunMethod:    ListenerRetry,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "list-rules",
		Short:        "Lists all rules of a listener",
		RunMethod:    ListenerListRules,
		FlagContexts: []cmd.FlagContext{fListenerId}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "create-rule",
		Short:        "Creates a rule in a listener",
		RunMethod:    ListenerCreateRule,
		FlagContexts: []cmd.FlagContext{fListenerId, fField, fValues, fTargetGroupId}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "update-rule",
		Short:        "Updates a rule in a target listener",
		RunMethod:    ListenerUpdateRule,
		FlagContexts: []cmd.FlagContext{fIdRule, fListenerId, fField, fValues}},
	)
	cmd.NewCommand(listenersCmd, &cmd.CommandContext{
		Use:          "delete-rule",
		Short:        "Destroys a rule in a listener",
		RunMethod:    ListenerDeleteRule,
		FlagContexts: []cmd.FlagContext{fIdRule, fListenerId}},
	)
}

// ListenerList subcommand function
func ListenerList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	listeners, err := svc.ListListeners(cmd.GetContext(), viper.GetString(cmd.LoadBalancerId))
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer listeners data", err)
		return err
	}
	if err = formatter.PrintList(listeners); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerShow subcommand function
func ListenerShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	listener, err := svc.GetListener(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive load balancer listener data", err)
		return err
	}
	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerCreate subcommand function
func ListenerCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	listenerIn := map[string]interface{}{
		"protocol":                viper.GetString(cmd.Protocol),
		"port":                    viper.GetInt(cmd.Port),
		"default_target_group_id": viper.GetString(cmd.TargetGroupId),
	}
	cmd.SetParamString("certificate_id", cmd.CertificateId, listenerIn)

	listener, err := svc.CreateListener(cmd.GetContext(), viper.GetString(cmd.LoadBalancerId), &listenerIn)
	if err != nil {
		formatter.PrintError("Couldn't create load balancer listener", err)
		return err
	}

	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerUpdate subcommand function
func ListenerUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	listenerIn := map[string]interface{}{
		"default_target_group_id": viper.GetString(cmd.TargetGroupId),
	}
	cmd.SetParamString("certificate_id", cmd.CertificateId, listenerIn)

	listener, err := svc.UpdateListener(cmd.GetContext(), viper.GetString(cmd.Id), &listenerIn)
	if err != nil {
		formatter.PrintError("Couldn't update load balancer listener", err)
		return err
	}

	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerDelete subcommand function
func ListenerDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	listener, err := svc.DeleteListener(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete load balancer listener", err)
		return err
	}

	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerRetry subcommand function
func ListenerRetry() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	listenerIn := map[string]interface{}{}
	listener, err := svc.RetryListener(cmd.GetContext(), viper.GetString(cmd.Id), &listenerIn)
	if err != nil {
		formatter.PrintError("Couldn't retry load balancer listener", err)
		return err
	}

	if err = formatter.PrintItem(*listener); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerListRules subcommand function
func ListenerListRules() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	rules, err := svc.ListRules(cmd.GetContext(), viper.GetString(cmd.ListenerId))
	if err != nil {
		formatter.PrintError("Couldn't receive listener rules data", err)
		return err
	}
	if err = formatter.PrintList(rules); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerCreateRule subcommand function
func ListenerCreateRule() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ruleIn := map[string]interface{}{
		"field":           viper.GetString(cmd.Field),
		"values":          utils.RemoveDuplicates(strings.Split(viper.GetString(cmd.Values), ",")),
		"target_group_id": viper.GetString(cmd.TargetGroupId),
	}

	rule, err := svc.CreateRule(cmd.GetContext(), viper.GetString(cmd.ListenerId), &ruleIn)
	if err != nil {
		formatter.PrintError("Couldn't create listener rule", err)
		return err
	}

	if err = formatter.PrintItem(*rule); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerUpdateRule subcommand function
func ListenerUpdateRule() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ruleIn := map[string]interface{}{
		"field":  viper.GetString(cmd.Field),
		"values": utils.RemoveDuplicates(strings.Split(viper.GetString(cmd.Values), ",")),
	}

	rule, err := svc.UpdateRule(cmd.GetContext(), viper.GetString(cmd.ListenerId), viper.GetString(cmd.Id), &ruleIn)
	if err != nil {
		formatter.PrintError("Couldn't update listener rule", err)
		return err
	}

	if err = formatter.PrintItem(*rule); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ListenerDeleteRule subcommand function
func ListenerDeleteRule() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteRule(cmd.GetContext(), viper.GetString(cmd.ListenerId), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete listener rule", err)
		return err
	}
	return nil
}
