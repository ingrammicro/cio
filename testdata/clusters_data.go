// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/api/types"

// GetClusterData loads test data
func GetClusterData() []*types.Cluster {

	return []*types.Cluster{
		{
			ID:                      "fakeID0",
			Name:                    "fakeName0",
			State:                   "fakeState0",
			RemoteID:                "fakeRemoteID0",
			CloudAccountID:          "fakeCloudAccountID0",
			RealmID:                 "fakeRealmID0",
			VpcID:                   "fakeVpcID0",
			Brownfield:              false,
			Version:                 "1.2",
			Endpoint:                "fakeEndpoint0",
			ClusterPlanID:           "fakeClusterPlanID0",
			PublicAccessIPAddresses: []string{},
			ErrorEventID:            "fakeErrorEventID0",
		},
	}
}

// GetClusterPlanData loads test data
func GetClusterPlanData() []*types.ClusterPlan {

	return []*types.ClusterPlan{
		{
			ID:                  "fakeID0",
			Name:                "fakeName0",
			AvailableVersions:   []string{"1.0", "1.1", "1.2"},
			DefaultVersion:      "1.2",
			MaxPodsPerNode:      50,
			MaxNodesPerNodePool: 50,
			CloudProviderID:     "fakeCloudProviderID0",
			CloudProviderName:   "fakeCloudProviderName0",
			RealmID:             "fakeRealmID0",
			RealmProviderName:   "fakeRealmProviderName0",
			FlavourProviderName: "fakeFlavourProviderName0",
			Deprecated:          false,
		},
	}
}
