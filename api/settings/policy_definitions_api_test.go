// Copyright (c) 2017-2021 Ingram Micro Inc.

package settings

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewPolicyDefinitionServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewPolicyDefinitionService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListDefinitions(t *testing.T) {
	policyDefinitionsIn := testdata.GetPolicyDefinitionData()
	ListDefinitionsMocked(t, policyDefinitionsIn)
	ListDefinitionsFailErrMocked(t, policyDefinitionsIn)
	ListDefinitionsFailStatusMocked(t, policyDefinitionsIn)
	ListDefinitionsFailJSONMocked(t, policyDefinitionsIn)
}

func TestGetDefinition(t *testing.T) {
	policyDefinitionsIn := testdata.GetPolicyDefinitionData()
	for _, policyDefinitionIn := range policyDefinitionsIn {
		GetDefinitionMocked(t, policyDefinitionIn)
		GetDefinitionFailErrMocked(t, policyDefinitionIn)
		GetDefinitionFailStatusMocked(t, policyDefinitionIn)
		GetDefinitionFailJSONMocked(t, policyDefinitionIn)
	}
}

func TestCreateDefinition(t *testing.T) {
	policyDefinitionsIn := testdata.GetPolicyDefinitionData()
	for _, policyDefinitionIn := range policyDefinitionsIn {
		CreateDefinitionMocked(t, policyDefinitionIn)
		CreateDefinitionFailErrMocked(t, policyDefinitionIn)
		CreateDefinitionFailStatusMocked(t, policyDefinitionIn)
		CreateDefinitionFailJSONMocked(t, policyDefinitionIn)
	}
}

func TestUpdateDefinition(t *testing.T) {
	policyDefinitionsIn := testdata.GetPolicyDefinitionData()
	for _, policyDefinitionIn := range policyDefinitionsIn {
		UpdateDefinitionMocked(t, policyDefinitionIn)
		UpdateDefinitionFailErrMocked(t, policyDefinitionIn)
		UpdateDefinitionFailStatusMocked(t, policyDefinitionIn)
		UpdateDefinitionFailJSONMocked(t, policyDefinitionIn)
	}
}

func TestDeleteDefinition(t *testing.T) {
	policyDefinitionsIn := testdata.GetPolicyDefinitionData()
	for _, policyDefinitionIn := range policyDefinitionsIn {
		DeleteDefinitionMocked(t, policyDefinitionIn)
		DeleteDefinitionFailErrMocked(t, policyDefinitionIn)
		DeleteDefinitionFailStatusMocked(t, policyDefinitionIn)
	}
}

func TestListDefinitionAssignments(t *testing.T) {
	policyDefinitionsIn := testdata.GetPolicyDefinitionData()
	policyAssignmentsIn := testdata.GetPolicyAssignmentData()
	for _, policyDefinitionIn := range policyDefinitionsIn {
		ListDefinitionAssignmentsMocked(t, policyDefinitionIn.ID, policyAssignmentsIn)
		ListDefinitionAssignmentsFailErrMocked(t, policyDefinitionIn.ID, policyAssignmentsIn)
		ListDefinitionAssignmentsFailStatusMocked(t, policyDefinitionIn.ID, policyAssignmentsIn)
		ListDefinitionAssignmentsFailJSONMocked(t, policyDefinitionIn.ID, policyAssignmentsIn)
	}
}
