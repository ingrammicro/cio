// +build solaris

package firewall

import (
	"fmt"

	"os"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

func driverName() string {
	return "iptables"
}

func Apply(policy types.Policy) error {

	// NO!
	f, err := os.Create("/etc/ipf/ipf.conf")

	if err != nil {
		return fmt.Errorf("Error opening /etc/ipf/ipf.conf : %s", err)
	}
	defer f.Close()

	f.WriteString("pass out on net0 from any to any keep state\n")
	f.WriteString("pass in quick on net0 proto icmp from any to any keep state\n")

	for _, rule := range policy.Rules {
		f.WriteString(fmt.Sprintf("pass in quick on net0 proto %s from %s to any %s\n", rule.Protocol, rule.Cidr, determinePort(rule.MinPort, rule.MaxPort)))
	}

	f.WriteString("block in on net0 from any to any\n")

	if output, exit, _, _ := utils.RunCmd("svcadm enable ipfilter; svcadm restart ipfilter; ipf -Fa -f /etc/ipf/ipf.conf"); exit != 0 {
		return fmt.Errorf("Error executing firewall enable: (%d) %s", exit, output)
	}

	return nil
}

func determinePort(min, max int) string {
	if min == max {
		return fmt.Sprintf("port = %d", min)
	}
	if min > 1 {
		if max < 65535 {
			return fmt.Sprintf("port %d >< %d", min-1, max+1)
		}
		return fmt.Sprintf("port => %d", min)
	}
	if max < 65535 {
		return fmt.Sprintf("port <= %d", max)
	}
	return ""
}

func flush() error {

	if output, exit, _, _ := utils.RunCmd("svcadm disable ipfilter"); exit != 0 {
		return fmt.Errorf("Error executing firewall flush: (%d) %s", exit, output)
	}
	return nil
}
