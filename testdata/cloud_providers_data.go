// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/api/types"

// GetCloudProviderData loads test data
func GetCloudProviderData() []*types.CloudProvider {

	return []*types.CloudProvider{
		{
			ID:   "fakeID0",
			Name: "fakeName0",
		},
		{
			ID:   "fakeID1",
			Name: "fakeName1",
		},
	}
}
