// Copyright (c) 2017-2022 Ingram Micro Inc.

// +build linux darwin

package brownfield

import (
	"github.com/ingrammicro/cio/cmd/agent"
	"github.com/ingrammicro/cio/cmd/agent/firewall"
	"github.com/ingrammicro/cio/types"
)

func Apply(p *types.Policy) error {
	agent.RunCmd("/sbin/iptables -w -F INPUT")

	if len(p.Rules) > 0 {
		return firewall.Apply(*p)
	}
	return nil
}
