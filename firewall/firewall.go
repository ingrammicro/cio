// Copyright (c) 2017-2021 Ingram Micro Inc.

package firewall

import (
	"fmt"

	"github.com/ingrammicro/cio/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const CurrentFirewallDriverDebugTrace = "Current firewall driver %s"

func cmdList(c *cli.Context) error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return cmd.FirewallRuleList(c)
}

func cmdApply(c *cli.Context) error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	policy := cmd.FirewallPolicyGet(c)
	// Only apply firewall if we get a non-empty set of rules
	if len(policy.Rules) > 0 {
		return Apply(*policy)
	}
	return flush()
}

func cmdFlush(c *cli.Context) error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return flush()
}

func cmdCheck(c *cli.Context) error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	_, _, exists := cmd.FirewallRuleCheck(c)
	fmt.Printf("%t\n", exists)
	return nil
}

func cmdAdd(c *cli.Context) error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return cmd.FirewallRuleAdd(c)
}

func cmdUpdate(c *cli.Context) error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return cmd.FirewallRulesUpdate(c)
}

func cmdRemove(c *cli.Context) error {
	log.Debugf(CurrentFirewallDriverDebugTrace, driverName())
	return cmd.FirewallRuleRemove(c)
}
