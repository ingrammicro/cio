// Copyright (c) 2017-2022 Ingram Micro Inc.

// +build linux

package firewall

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/agent"
	"os"

	"github.com/ingrammicro/cio/types"

	log "github.com/sirupsen/logrus"
)

func driverName() string {
	return "iptables"
}

func Apply(policy types.Policy) error {
	var exitCode int
	agent.RunCmd("/sbin/iptables -w -N CONCERTO")
	agent.RunCmd("/sbin/iptables -w -F CONCERTO")
	agent.RunCmd("/sbin/iptables -w -P INPUT DROP")

	_, exitCode, _, _ = agent.RunCmd("/sbin/iptables -w -C INPUT -i lo -j ACCEPT")
	if exitCode != 0 {
		agent.RunCmd("/sbin/iptables -w -A INPUT -i lo -j ACCEPT")
	}

	_, exitCode, _, _ = agent.RunCmd("/sbin/iptables -w -C INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")
	if exitCode != 0 {
		agent.RunCmd("/sbin/iptables -w -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT")
	}

	for _, rule := range policy.Rules {
		agent.RunCmd(
			fmt.Sprintf(
				"/sbin/iptables -w -A CONCERTO -s %s -p %s --dport %d:%d -j ACCEPT",
				rule.Cidr,
				rule.Protocol,
				rule.MinPort,
				rule.MaxPort,
			),
		)
	}

	_, exitCode, _, _ = agent.RunCmd("/sbin/iptables -w -C INPUT -j CONCERTO")
	if exitCode != 0 {
		log.Debugln("Concerto Chain is not existent adding it to INPUT")
		agent.RunCmd("/sbin/iptables -w -A INPUT -j CONCERTO")
	}

	return nil
}

func flush() error {
	if _, err := os.Stat("/etc/redhat-release"); err == nil {
		agent.RunCmd("firewall-cmd --set-default-zone=trusted")
	}
	agent.RunCmd("/sbin/iptables -w -P INPUT ACCEPT")
	agent.RunCmd("/sbin/iptables -w -F CONCERTO")
	agent.RunCmd("/sbin/iptables -w -D INPUT -j CONCERTO")
	agent.RunCmd("/sbin/iptables -w -X CONCERTO")
	return nil
}
