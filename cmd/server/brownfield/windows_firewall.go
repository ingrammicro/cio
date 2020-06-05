// Copyright (c) 2017-2021 Ingram Micro Inc.

// +build windows

package brownfield

import (
	"github.com/ingrammicro/cio/cmd/server/firewall"
	"github.com/ingrammicro/cio/types"
)

func Apply(p *types.Policy) error {
	if len(p.Rules) > 0 {
		return firewall.Apply(*p)
	}
	return nil
}
