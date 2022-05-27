// Copyright (c) 2017-2022 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/types"

// GetServerPlanData loads test data
func GetServerPlanData() []*types.ServerPlan {

	return []*types.ServerPlan{
		{
			ID:                  "fakeID0",
			Name:                "fakeName0",
			Memory:              512,
			CPUs:                2,
			Storage:             2048,
			LocationID:          "fakeLocationID0",
			LocationName:        "fakeLocationName0",
			RealmID:             "fakeRealmID0",
			RealmProviderName:   "fakeRealmProviderNameID0",
			FlavourProviderName: "fakeFlavourProviderNameID0",
			CloudProviderID:     "fakeCloudProviderID0",
			CloudProviderName:   "fakeCloudProviderName0",
		},
		{
			ID:                  "fakeID1",
			Name:                "fakeName1",
			Memory:              256,
			CPUs:                3,
			Storage:             1024,
			LocationID:          "fakeLocationID1",
			LocationName:        "fakeLocationName1",
			RealmID:             "fakeRealmID1",
			RealmProviderName:   "fakeRealmProviderNameID1",
			FlavourProviderName: "fakeFlavourProviderNameID1",
			CloudProviderID:     "fakeCloudProviderID1",
			CloudProviderName:   "fakeCloudProviderName1",
		},
	}
}
