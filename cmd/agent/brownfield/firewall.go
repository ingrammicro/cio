// Copyright (c) 2017-2022 Ingram Micro Inc.

package brownfield

import (
	"context"
	"fmt"

	"github.com/ingrammicro/cio/api"

	"strings"

	"github.com/ingrammicro/cio/cmd/agent/firewall/discovery"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils/format"
)

func configureFirewall(ctx context.Context, svc *api.ServerAPI, f format.Formatter) {
	chains, err := discovery.CurrentFirewallRules()
	if err != nil {
		f.PrintFatal("Cannot obtain current firewall rules", err)
	}
	flattenedInputChain, err := discovery.FlattenChain("INPUT", chains, nil)
	if err != nil {
		f.PrintFatal("Cannot flatten firewall INPUT chain", err)
	}
	fmt.Printf("After flattening chain: %d rules\n", len(flattenedInputChain.Rules))
	policy, err := startFirewallMapping(ctx, svc, flattenedInputChain.Rules)
	if err != nil {
		f.PrintFatal("Error starting the firewall mapping", err)
	}
	err = Apply(policy)
	if err != nil {
		f.PrintFatal("Applying firewall", err)
	}
}

func startFirewallMapping(ctx context.Context, svc *api.ServerAPI, rules []*discovery.FirewallRule) (
	p *types.Policy, err error) {
	payload := convertFirewallChainToPayload(rules)
	fmt.Printf("DEBUG: Sending following firewall profile: %+v\n", payload)

	response, status, err := svc.SetFirewallProfile(ctx, &payload)
	if err != nil {
		return
	}
	if status >= 300 {
		err = fmt.Errorf("server responded with %d code: %v", status, response)
		return
	}
	p = &(response.Profile)
	return
}

func convertFirewallChainToPayload(rules []*discovery.FirewallRule) map[string]interface{} {
	var fpRules []interface{}
	for _, r := range rules {
		newRules := convertRuleToPayload(r)
		if newRules != nil {
			fpRules = append(fpRules, newRules...)
		}
	}
	fp := map[string]interface{}{
		"firewall_profile": map[string]interface{}{
			"rules": fpRules,
		},
	}
	return fp
}

func convertRuleToPayload(rule *discovery.FirewallRule) []interface{} {
	var rules []interface{}
	protocol := strings.ToLower(rule.Protocol)
	if protocol != "all" && protocol != "tcp" && protocol != "udp" {
		return nil
	}
	if protocol == "all" || protocol == "tcp" {
		rules = append(rules,
			types.PolicyRule{
				Name:     rule.Name,
				Protocol: "tcp",
				Cidr:     rule.Source,
				MinPort:  rule.Dports[0],
				MaxPort:  rule.Dports[1],
			})
	}
	if protocol == "all" || protocol == "udp" {
		rules = append(rules,
			types.PolicyRule{
				Name:     rule.Name,
				Protocol: "udp",
				Cidr:     rule.Source,
				MinPort:  rule.Dports[0],
				MaxPort:  rule.Dports[1],
			})
	}
	return rules
}
