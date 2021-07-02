// Copyright (c) 2017-2021 Ingram Micro Inc.

// +build linux darwin

package brownfield

import (
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/firewall"
	"github.com/ingrammicro/cio/utils"
)

func Apply(p *types.Policy) error {
	utils.RunCmd("/sbin/iptables -w -F INPUT")

	if len(p.Rules) > 0 {
		return firewall.Apply(*p)
	}
	return nil
}
