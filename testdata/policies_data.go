// Copyright (c) 2017-2022 Ingram Micro Inc.

package testdata

import (
	types2 "github.com/ingrammicro/cio/types"
)

// GetPolicyDefinitionData loads test data
func GetPolicyDefinitionData() []*types2.PolicyDefinition {
	return []*types2.PolicyDefinition{
		{
			ID:          "fakeID0",
			Name:        "fakeName0",
			Description: "fakeDescription0",
			Definition:  "fakeDefinition0",
			Parameters:  map[string]interface{}{"fakeParameter0": "fakeVal0", "fakeParameter1": "fakeVal1"},
			Builtin:     false,
		},
		{
			ID:          "fakeID1",
			Name:        "fakeName1",
			Description: "fakeDescription1",
			Definition:  "fakeDefinition1",
			Parameters:  map[string]interface{}{"fakeParameter0": "fakeVal0", "fakeParameter1": "fakeVal1"},
			Builtin:     true,
		},
	}

}

// GetPolicyAssignmentData loads test data
func GetPolicyAssignmentData() []*types2.PolicyAssignment {
	return []*types2.PolicyAssignment{
		{
			ID:              "fakeID0",
			Name:            "fakeName0",
			Description:     "fakeDescription0",
			State:           "fakeState0",
			RemoteID:        "fakeRemoteID0",
			CloudAccountID:  "fakeCloudAccountID0",
			ErrorEventID:    "fakeErrorEventID0",
			ParameterValues: map[string]interface{}{"fakeParameter0": "fakeVal0", "fakeParameter1": "fakeVal1"},
			DefinitionID:    "fakeDefinitionID0",
			ResellerApplied: false,
		},
		{
			ID:              "fakeID1",
			Name:            "fakeName1",
			Description:     "fakeDescription1",
			State:           "fakeState1",
			RemoteID:        "fakeRemoteID1",
			CloudAccountID:  "fakeCloudAccountID1",
			ErrorEventID:    "fakeErrorEventID1",
			ParameterValues: map[string]interface{}{"fakeParameter0": "fakeVal0", "fakeParameter1": "fakeVal1"},
			DefinitionID:    "fakeDefinitionID1",
			ResellerApplied: true,
		},
	}
}
