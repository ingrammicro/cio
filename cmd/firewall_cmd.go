package cmd

import (
	"github.com/ingrammicro/cio/api/firewall"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpFirewall prepares common resources to send request to Concerto API
func WireUpFirewall(c *cli.Context) (ds *firewall.FirewallService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = firewall.NewFirewallService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up firewall service", err)
	}

	return ds, f
}

// FirewallPolicyGet subcommand function
func FirewallPolicyGet(c *cli.Context) *types.Policy {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpFirewall(c)

	policy, err := svc.GetPolicy()
	if err != nil {
		formatter.PrintFatal("Couldn't receive firewall policy data", err)
	}
	return policy
}

// FirewallRuleList subcommand function
func FirewallRuleList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpFirewall(c)

	policy, err := svc.GetPolicy()
	if err != nil {
		formatter.PrintFatal("Couldn't receive firewall policy data", err)
	}
	if err = formatter.PrintList(policy.ActualRules); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// FirewallRuleCheck subcommand function
func FirewallRuleCheck(c *cli.Context) (*types.Policy, *types.PolicyRule, bool) {
	debugCmdFuncInfo(c)
	_, formatter := WireUpFirewall(c)

	checkRequiredFlags(c, []string{"cidr", "min-port", "max-port", "ip-protocol"}, formatter)

	// API accepts only 1 rule
	rule := &types.PolicyRule{
		Cidr:     c.String("cidr"),
		MinPort:  c.Int("min-port"),
		MaxPort:  c.Int("max-port"),
		Protocol: c.String("ip-protocol"),
	}
	policy := FirewallPolicyGet(c)
	return policy, rule, policy.CheckPolicyRule(*rule)
}

// FirewallRuleAdd subcommand function
func FirewallRuleAdd(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpFirewall(c)

	_, newRule, exists := FirewallRuleCheck(c)
	if exists == false {
		nRule := map[string]interface{}{
			"rule": *newRule,
		}
		_, err := svc.AddPolicyRule(&nRule)
		if err != nil {
			formatter.PrintFatal("Couldn't add firewall policy rule data", err)
		}
	}
	return nil
}

// FirewallRulesUpdate subcommand function
func FirewallRulesUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpFirewall(c)

	checkRequiredFlags(c, []string{"rules"}, formatter)
	params, err := utils.FlagConvertParamsJSON(c, []string{"rules"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}
	in := make(map[string]interface{})
	in["firewall_profile"] = params

	_, err = svc.UpdatePolicy(&in)
	if err != nil {
		formatter.PrintFatal("Couldn't update firewall policy data", err)
	}
	return nil
}

// FirewallRuleRemove subcommand function
func FirewallRuleRemove(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpFirewall(c)

	policy, existingRule, exists := FirewallRuleCheck(c)
	if exists == true {
		for i, rule := range policy.Rules {
			if rule == *existingRule {
				policy.Rules = append(policy.Rules[:i], policy.Rules[1+i:]...)
				break
			}
		}

		in := make(map[string]interface{})
		in["firewall_profile"] = policy

		_, err := svc.UpdatePolicy(&in)
		if err != nil {
			formatter.PrintFatal("Couldn't update firewall policy data", err)
		}
	}
	return nil
}
