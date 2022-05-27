// Copyright (c) 2017-2022 Ingram Micro Inc.

// +build windows

package firewall

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/agent"

	"github.com/ingrammicro/cio/cmd/agent/firewall/discovery"
	"github.com/ingrammicro/cio/types"
)

func driverName() string {
	return "windows"
}

func Apply(policy types.Policy) error {
	err := flush()
	if err != nil {
		return err
	}
	for i, rule := range policy.Rules {
		cidr := rule.Cidr
		if rule.Cidr == "0.0.0.0/0" {
			cidr = "any"
		}
		ruleCmd := fmt.Sprintf(
			"netsh advfirewall firewall add rule "+
				"name=\"Concerto firewall %d\" "+
				"dir=in action=allow "+
				"remoteip=\"%s\" "+
				"protocol=\"%s\" "+
				"localport=\"%d-%d\"",
			i, cidr, rule.Protocol, rule.MinPort, rule.MaxPort)
		agent.RunCmd(ruleCmd)
	}

	agent.RunCmd("netsh advfirewall set allprofiles state on")
	return nil
}

func flush() error {
	fc, err := discovery.CurrentFirewallRules()
	if err != nil {
		return err
	}
	agent.RunCmd("netsh advfirewall set allprofiles state off")
	agent.RunCmd("netsh advfirewall set allprofiles firewallpolicy allowinbound,allowoutbound")
	//utils.RunCmd("netsh advfirewall firewall delete rule name=all")
	for _, r := range fc[0].Rules {
		agent.RunCmd(fmt.Sprintf("netsh advfirewall firewall delete rule name=%q", r.Name))
	}
	return nil
}
