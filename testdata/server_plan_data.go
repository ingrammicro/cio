// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/api/types"

// GetServerPlanData loads test data
func GetServerPlanData() []*types.ServerPlan {

	return []*types.ServerPlan{
		{
			ID:                "fakeID0",
			Name:              "fakeName0",
			Memory:            512,
			CPUs:              2,
			Storage:           2048,
			LocationID:        "fakeLocationID0",
			LocationName:      "fakeLocationName0",
			CloudProviderID:   "fakeCloudProviderID0",
			CloudProviderName: "fakeCloudProviderName0",
		},
		{
			ID:                "fakeID1",
			Name:              "fakeName1",
			Memory:            256,
			CPUs:              3,
			Storage:           1024,
			LocationID:        "fakeLocationID1",
			LocationName:      "fakeLocationName1",
			CloudProviderID:   "fakeCloudProviderID1",
			CloudProviderName: "fakeCloudProviderName1",
		},
	}
}
