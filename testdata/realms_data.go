// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/api/types"

// GetRealmData loads test data
func GetRealmData() []*types.Realm {

	return []*types.Realm{
		{
			ID:              "fakeID0",
			Name:            "fakeName0",
			LocationID:      "fakeLocationID0",
			CloudProviderID: "fakeCloudProviderID0",
			ProviderName:    "fakeProviderName0",
			Deprecated:      false,
		},
	}
}
