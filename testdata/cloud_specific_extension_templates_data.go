// Copyright (c) 2017-2022 Ingram Micro Inc.

package testdata

import "github.com/ingrammicro/cio/types"

// GetCloudSpecificExtensionTemplateData loads test data
func GetCloudSpecificExtensionTemplateData() []*types.CloudSpecificExtensionTemplate {
	return []*types.CloudSpecificExtensionTemplate{
		{
			ID:         "fakeID0",
			Name:       "fakeName0",
			Global:     false,
			Definition: "fakeDefinition0",
			Parameters: nil,
			Syntax:     "fakeSyntax0",
		},
		{
			ID:         "fakeID0",
			Name:       "fakeName0",
			Global:     false,
			Definition: "fakeDefinition0",
			Parameters: nil,
			Syntax:     "fakeSyntax0",
		},
	}
}

// GetCloudSpecificExtensionDeploymentData loads test data
func GetCloudSpecificExtensionDeploymentData() []*types.CloudSpecificExtensionDeployment {
	return []*types.CloudSpecificExtensionDeployment{
		{
			ID:              "fakeID0",
			Name:            "fakeName0",
			TemplateID:      "fakeTemplateID0",
			State:           "fakeState0",
			RemoteID:        "fakeRemoteID0",
			CloudAccountID:  "fakeCloudAccountID0",
			RealmID:         "fakeRealmID0",
			ErrorEventID:    "fakeErrorEventID0",
			ParameterValues: nil,
			Outputs:         nil,
		},
		{
			ID:              "fakeID1",
			Name:            "fakeName1",
			TemplateID:      "fakeTemplateID1",
			State:           "fakeState1",
			RemoteID:        "fakeRemoteID1",
			CloudAccountID:  "fakeCloudAccountID1",
			RealmID:         "fakeRealmID1",
			ErrorEventID:    "fakeErrorEventID1",
			ParameterValues: nil,
			Outputs:         nil,
		},
	}
}
