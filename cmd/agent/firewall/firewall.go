// Copyright (c) 2017-2022 Ingram Micro Inc.

package firewall

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/agent"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	fCidr := cmd.FlagContext{Type: cmd.String, Name: cmd.Cidr, Required: true, Usage: "CIDR"}

	fMinPort := cmd.FlagContext{Type: cmd.Int, Name: cmd.MinPort, Required: true, Usage: "Minimum Port"}

	fMaxPort := cmd.FlagContext{Type: cmd.Int, Name: cmd.MaxPort, Required: true, Usage: "Maximum Port"}

	fIpProtocol := cmd.FlagContext{Type: cmd.String, Name: cmd.IpProtocol, Required: true,
		Usage: "Ip protocol udp or tcp"}

	fRules := cmd.FlagContext{Type: cmd.String, Name: cmd.Rules, Required: true,
		Usage: "JSON array in the form " +
			"'[{\"ip_protocol\":\"...\", \"min_port\":..., \"max_port\":..., \"cidr_ip\":\"...\"}, ... ]'"}

	firewallCmd := cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:   "firewall",
		Short: "Manages Firewall Policies within a Host"},
	)
	cmd.NewCommand(firewallCmd, &cmd.CommandContext{
		Use:          "add",
		Short:        "Adds a single firewall rule to host",
		RunMethod:    cmdAdd,
		FlagContexts: []cmd.FlagContext{fCidr, fMinPort, fMaxPort, fIpProtocol}},
	)
	cmd.NewCommand(firewallCmd, &cmd.CommandContext{
		Use:       "apply",
		Short:     "Applies selected firewall rules in host",
		RunMethod: cmdApply},
	)
	cmd.NewCommand(firewallCmd, &cmd.CommandContext{
		Use:          "check",
		Short:        "Checks if a firewall rule exists in host",
		RunMethod:    cmdCheck,
		FlagContexts: []cmd.FlagContext{fCidr, fMinPort, fMaxPort, fIpProtocol}},
	)
	cmd.NewCommand(firewallCmd, &cmd.CommandContext{
		Use:       "flush",
		Short:     "Flushes all firewall rules from host",
		RunMethod: cmdFlush},
	)
	cmd.NewCommand(firewallCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists all firewall rules associated to host",
		RunMethod: cmdList},
	)
	cmd.NewCommand(firewallCmd, &cmd.CommandContext{
		Use:          "remove",
		Short:        "Removes a firewall rule to host",
		RunMethod:    cmdRemove,
		FlagContexts: []cmd.FlagContext{fCidr, fMinPort, fMaxPort, fIpProtocol}},
	)
	cmd.NewCommand(firewallCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates all firewall rules",
		RunMethod:    cmdUpdate,
		FlagContexts: []cmd.FlagContext{fRules}},
	)
}

const CurrentFirewallDriverDebugTrace = "Current firewall driver %s"

func cmdList() error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return RuleList()
}

func cmdApply() error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	policy := PolicyGet()
	// Only apply firewall if we get a non-empty set of rules
	if len(policy.Rules) > 0 {
		return Apply(*policy)
	}
	return flush()
}

func cmdFlush() error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return flush()
}

func cmdCheck() error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	_, _, exists := RuleCheck()
	fmt.Printf("%t\n", exists)
	return nil
}

func cmdAdd() error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return RuleAdd()
}

func cmdUpdate() error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return RulesUpdate()
}

func cmdRemove() error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return RuleRemove()
}

// PolicyGet subcommand function
func PolicyGet() *types.Policy {
	logger.DebugFuncInfo()
	svc, _, formatter := agent.WireUpAPIServer()

	policy, err := svc.GetPolicy(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive firewall policy data", err)
	}
	return policy
}

// RuleList subcommand function
func RuleList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := agent.WireUpAPIServer()

	policy, err := svc.GetPolicy(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive firewall policy data", err)
	}
	if err = formatter.PrintList(policy.ActualRules); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// RuleCheck subcommand function
func RuleCheck() (*types.Policy, *types.PolicyRule, bool) {
	logger.DebugFuncInfo()

	// API accepts only 1 rule
	rule := &types.PolicyRule{
		Cidr:     viper.GetString(cmd.Cidr),
		MinPort:  viper.GetInt(cmd.MinPort),
		MaxPort:  viper.GetInt(cmd.MaxPort),
		Protocol: viper.GetString(cmd.IpProtocol),
	}
	policy := PolicyGet()
	return policy, rule, policy.CheckPolicyRule(*rule)
}

// RuleAdd subcommand function
func RuleAdd() error {
	logger.DebugFuncInfo()
	svc, _, formatter := agent.WireUpAPIServer()

	_, newRule, exists := RuleCheck()
	if exists == false {
		nRule := map[string]interface{}{
			"rule": *newRule,
		}
		_, err := svc.AddPolicyRule(cmd.GetContext(), &nRule)
		if err != nil {
			formatter.PrintFatal("Couldn't add firewall policy rule data", err)
		}
	}
	return nil
}

// RulesUpdate subcommand function
func RulesUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := agent.WireUpAPIServer()

	params, err := cmd.FlagConvertParamsJSON(cmd.Rules)
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}
	in := make(map[string]interface{})
	in["firewall_profile"] = params

	_, err = svc.UpdatePolicy(cmd.GetContext(), &in)
	if err != nil {
		formatter.PrintFatal("Couldn't update firewall policy data", err)
	}
	return nil
}

// RuleRemove subcommand function
func RuleRemove() error {
	logger.DebugFuncInfo()
	svc, _, formatter := agent.WireUpAPIServer()

	policy, existingRule, exists := RuleCheck()
	if exists == true {
		for i, rule := range policy.Rules {
			if rule == *existingRule {
				policy.Rules = append(policy.Rules[:i], policy.Rules[1+i:]...)
				break
			}
		}

		in := make(map[string]interface{})
		in["firewall_profile"] = policy

		_, err := svc.UpdatePolicy(cmd.GetContext(), &in)
		if err != nil {
			formatter.PrintFatal("Couldn't update firewall policy data", err)
		}
	}
	return nil
}
