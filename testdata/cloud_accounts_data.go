// Copyright (c) 2017-2022 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/types"

// GetCloudAccountData loads test data
func GetCloudAccountData() []*types.CloudAccount {

	return []*types.CloudAccount{
		{
			ID:                "fakeID0",
			Name:              "fakeName0",
			CloudProviderID:   "CloudProviderID0",
			CloudProviderName: "CloudProviderName0",
		},
		{
			ID:                "fakeID1",
			Name:              "fakeName1",
			CloudProviderID:   "CloudProviderID1",
			CloudProviderName: "CloudProviderName1",
		},
	}
}
