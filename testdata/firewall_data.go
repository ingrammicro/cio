// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import (
	"github.com/ingrammicro/cio/api/types"
)

// GetPolicyData loads test data
func GetPolicyData() *types.Policy {
	return &types.Policy{
		ActualRules: []types.PolicyRule{
			{
				Protocol: "fakeProtocol0",
				MinPort:  0,
				MaxPort:  1024,
				Cidr:     "fakeCidrIP0",
			},
		},
		Rules: []types.PolicyRule{
			{
				Protocol: "fakeProtocol1",
				MinPort:  0,
				MaxPort:  1024,
				Cidr:     "fakeCidrIP1",
			},
		},
	}
}
