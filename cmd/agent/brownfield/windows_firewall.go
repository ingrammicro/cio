// Copyright (c) 2017-2022 Ingram Micro Inc.

// +build windows

package brownfield

import (
	"github.com/ingrammicro/cio/cmd/agent/firewall"
	"github.com/ingrammicro/cio/types"
)

func Apply(p *types.Policy) error {
	if len(p.Rules) > 0 {
		return firewall.Apply(*p)
	}
	return nil
}
