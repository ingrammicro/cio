// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/api/types"

// GetNodePoolData loads test data
func GetNodePoolData() []*types.NodePool {

	return []*types.NodePool{
		{
			ID:             "fakeID0",
			Name:           "fakeName0",
			State:          "fakeState0",
			RemoteID:       "fakeRemoteID0",
			ClusterID:      "fakeClusterID0",
			SubnetID:       "fakeSubnetID0",
			NodePoolPlanID: "fakeNodePoolPlanID0",
			DiskSize:       0,
			OSType:         "fakeOSType0",
			CpuType:        "fakeCpuType0",
			MinNodes:       0,
			MaxNodes:       0,
			DesiredNodes:   0,
			PodsPerNode:    0,
			Autoscaling:    false,
			Brownfield:     false,
			ErrorEventID:   "fakeErrorEventID0",
		},
	}
}

// GetNodePoolPlanData loads test data
func GetNodePoolPlanData() []*types.NodePoolPlan {

	return []*types.NodePoolPlan{
		{
			ID:                   "fakeID0",
			Name:                 "fakeName0",
			RemoteID:             "fakeRemoteID0",
			CPUTypes:             nil,
			CPUs:                 1,
			Memory:               2,
			RealmID:              "fakeRealmID0",
			ServerPlanID:         "fakeServerPlanID0",
			AutoscalingCapable:   false,
			PodsPerNodePresence:  false,
			FirstNodePoolCapable: false,
			Deprecated:           false,
		},
	}
}
