// Copyright (c) 2017-2021 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/api/types"

// GetCloudApplicationDeploymentData loads test data
func GetCloudApplicationDeploymentData() []*types.CloudApplicationDeployment {
	return []*types.CloudApplicationDeployment{
		{
			ID:        "fakeID0",
			Name:      "fakeName0",
			Namespace: "cat:deployment",
			Value:     "fakeValue0",
			CatID:     "fakeCatID0",
		},
		{
			ID:        "fakeID1",
			Name:      "fakeName1",
			Namespace: "cat:deployment",
			Value:     "fakeValue1",
			CatID:     "fakeCatID1",
		},
	}
}

// GetCloudApplicationDeploymentTaskData loads test data
func GetCloudApplicationDeploymentTaskData() []*types.CloudApplicationDeploymentTask {
	return []*types.CloudApplicationDeploymentTask{
		{
			ID:           "fakeID0",
			Type:         "fakeType0",
			LabelName:    "fakeLabelName0",
			LabelID:      "fakeLabelID0",
			State:        "fakeState0",
			ErrorMessage: "fakeErrorMessage0",
			Outputs:      "fakeOutputs0",
			UserID:       "fakeUserID0",
			ArchiveID:    "fakeArchiveID0",
			DeploymentID: "fakeDeploymentID0",
		},
		{
			ID:           "fakeID1",
			Type:         "fakeType1",
			LabelName:    "fakeLabelName1",
			LabelID:      "fakeLabelID1",
			State:        "fakeState1",
			ErrorMessage: "fakeErrorMessage1",
			Outputs:      "fakeOutputs1",
			UserID:       "fakeUserID1",
			ArchiveID:    "fakeArchiveID1",
			DeploymentID: "fakeDeploymentID1",
		},
	}
}
