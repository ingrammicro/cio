// Copyright (c) 2017-2022 Ingram Micro Inc.

// +build linux darwin

package discovery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func CurrentFirewallRules() ([]*FirewallChain, error) {
	output, err := exec.Command("/sbin/iptables", "-L", "-n", "-v").Output()
	if err != nil {
		return nil, fmt.Errorf("running iptables list command to obtain current firewall rules: %v", err)
	}
	chains, err := parseIptablesOutput(string(output))
	if err != nil {
		return nil, fmt.Errorf("parsing iptables list command output to obtain current firewall rules: %v", err)
	}
	return chains, nil
}

func parseIptablesOutput(output string) ([]*FirewallChain, error) {
	var chains []*FirewallChain
	cs := strings.Split(output, "\n\n")
	fmt.Printf("Found %d chains\n", len(cs))
	for _, c := range cs {
		chain, err := parseIptablesChain(c)
		if err == nil {
			chains = append(chains, chain)
		} else {
			fmt.Printf("Warning: error occurred while parsing iptables chain: %v\n", err)
		}
	}
	return chains, nil
}

var iptablesChainHeaderRegexp = regexp.MustCompile(
	`\AChain (?P<name>[a-zA-Z0-9-]+) \((policy (?P<policy>[a-zA-Z0-9-]+) )?`)

func buildIptablesFirewallChain(header string, rules []string) (*FirewallChain, error) {
	chain := &FirewallChain{}
	match := iptablesChainHeaderRegexp.FindStringSubmatch(header)
	for i, name := range iptablesChainHeaderRegexp.SubexpNames() {
		switch name {
		case "name":
			chain.Name = match[i]
		case "policy":
			chain.Policy = match[i]
		default:
		}
	}
	if chain.Name == "" {
		return nil, fmt.Errorf("found no name for chain in '%s'", header)
	}
	for _, r := range rules {
		if r != "" {
			rule, err := parseIptablesRule(r)
			if err != nil {
				fmt.Printf("Warning: cannot parse rule for chain %s : %v\n", chain.Name, err)
			} else {
				if rule != nil {
					chain.Rules = append(chain.Rules, rule)
				}
			}
		}
	}
	return chain, nil
}

func parseIptablesChain(c string) (*FirewallChain, error) {
	lines := strings.Split(c, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("chain output has no header")
	}
	if len(lines) < 2 {
		return nil, fmt.Errorf("chain output has no rules header")
	}
	header := lines[0]
	matched, err := regexp.Match(iptablesChainHeaderRegexp.String(), []byte(header))
	if err != nil {
		return nil, fmt.Errorf("cannot parse chain header '%s': %v", header, err)
	}
	if !matched {
		return nil, fmt.Errorf("cannot parse chain header '%s'", header)
	}
	rules := lines[2:]
	return buildIptablesFirewallChain(header, rules)
}

var iptablesRuleFieldSeparator = regexp.MustCompile("[[:blank:]]+")
var iptablesRuleDPortRegexp = regexp.MustCompile(`(tcp|udp) dpts?:(?P<minPort>\d+)(:(?P<maxPort>\d+))?`)
var iptablesRuleStateRegexp = regexp.MustCompile(`state [[:alpha:]]+(,[[:alpha:]]+)*`)
var iptablesRuleStatsInfoRegexp = regexp.MustCompile(`^ ?\d+[A-Z]? \d+[A-Z]? `)

func parseRuleDPorts(r string) ([2]int, error) {
	match := iptablesRuleDPortRegexp.FindStringSubmatch(r)
	dports := [2]int{0, 0}
	for i, name := range iptablesRuleDPortRegexp.SubexpNames() {
		var err error
		switch name {
		case "minPort":
			dports[0], err = strconv.Atoi(match[i])
			if err != nil {
				return dports, fmt.Errorf("rule '%s' has invalid destination %s specification '%s'", r, name, match[i])
			}
		case "maxPort":
			dports[1], _ = strconv.Atoi(match[i])
			if err != nil {
				dports[1] = 0
			}
		default:
		}
	}
	if dports[1] == 0 {
		dports[1] = dports[0]
	}
	return dports, nil
}

func parseRule(r, target, protocol, source string) (*FirewallRule, error) {
	var dports [2]int
	var err error
	matchString := iptablesRuleDPortRegexp.FindString(r)
	if len(matchString) == 0 {
		dports = [2]int{1, 65535}
	} else {
		dports, err = parseRuleDPorts(r)
		if err != nil {
			return nil, err
		}
	}
	rule := &FirewallRule{
		Target:   target,
		Protocol: protocol,
		Source:   source,
		Dports:   dports,
	}
	return rule, nil
}

func parseIptablesRule(r string) (*FirewallRule, error) {
	r = iptablesRuleFieldSeparator.ReplaceAllLiteralString(r, " ")
	r = iptablesRuleStatsInfoRegexp.ReplaceAllLiteralString(r, "")
	fields := iptablesRuleFieldSeparator.Split(r, 8)
	if len(fields) < 8 {
		return nil, fmt.Errorf("rule '%s' has too few fields", r)
	}
	matchString := iptablesRuleStateRegexp.FindString(r)
	if len(matchString) > 0 { //state condition, we ignore it
		return nil, nil
	}
	if fields[3] == "lo" { // incoming interface is localhost
		return nil, nil
	}
	return parseRule(r, fields[0], fields[1], fields[5])
}
