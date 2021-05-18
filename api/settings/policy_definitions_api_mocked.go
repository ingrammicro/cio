// Copyright (c) 2017-2021 Ingram Micro Inc.

package settings

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// ListDefinitionsMocked test mocked function
func ListDefinitionsMocked(t *testing.T, policyDefinitionsIn []*types.PolicyDefinition) []*types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionsIn)
	assert.Nil(err, "PolicyDefinitions test data corrupted")

	// call service
	cs.On("Get", APIPathPolicyDefinitions).Return(dIn, 200, nil)
	policyDefinitionsOut, err := ds.ListDefinitions()

	assert.Nil(err, "Error getting policy definitions")
	assert.Equal(policyDefinitionsIn, policyDefinitionsOut, "ListDefinitions returned different policy definitions")

	return policyDefinitionsOut
}

// ListDefinitionsFailErrMocked test mocked function
func ListDefinitionsFailErrMocked(
	t *testing.T,
	policyDefinitionsIn []*types.PolicyDefinition,
) []*types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionsIn)
	assert.Nil(err, "PolicyDefinitions test data corrupted")

	// call service
	cs.On("Get", APIPathPolicyDefinitions).Return(dIn, 200, fmt.Errorf("mocked error"))
	policyDefinitionsOut, err := ds.ListDefinitions()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyDefinitionsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyDefinitionsOut
}

// ListDefinitionsFailStatusMocked test mocked function
func ListDefinitionsFailStatusMocked(
	t *testing.T,
	policyDefinitionsIn []*types.PolicyDefinition,
) []*types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionsIn)
	assert.Nil(err, "PolicyDefinitions test data corrupted")

	// call service
	cs.On("Get", APIPathPolicyDefinitions).Return(dIn, 499, nil)
	policyDefinitionsOut, err := ds.ListDefinitions()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyDefinitionsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyDefinitionsOut
}

// ListDefinitionsFailJSONMocked test mocked function
func ListDefinitionsFailJSONMocked(
	t *testing.T,
	policyDefinitionsIn []*types.PolicyDefinition,
) []*types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathPolicyDefinitions).Return(dIn, 200, nil)
	policyDefinitionsOut, err := ds.ListDefinitions()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyDefinitionsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyDefinitionsOut
}

// GetDefinitionMocked test mocked function
func GetDefinitionMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID)).Return(dIn, 200, nil)
	policyDefinitionOut, err := ds.GetDefinition(policyDefinitionIn.ID)

	assert.Nil(err, "Error getting policy definition")
	assert.Equal(*policyDefinitionIn, *policyDefinitionOut, "GetDefinition returned different policy definition")

	return policyDefinitionOut
}

// GetDefinitionFailErrMocked test mocked function
func GetDefinitionFailErrMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	policyDefinitionOut, err := ds.GetDefinition(policyDefinitionIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyDefinitionOut
}

// GetDefinitionFailStatusMocked test mocked function
func GetDefinitionFailStatusMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID)).Return(dIn, 499, nil)
	policyDefinitionOut, err := ds.GetDefinition(policyDefinitionIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyDefinitionOut
}

// GetDefinitionFailJSONMocked test mocked function
func GetDefinitionFailJSONMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID)).Return(dIn, 200, nil)
	policyDefinitionOut, err := ds.GetDefinition(policyDefinitionIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyDefinitionOut
}

// CreateDefinitionMocked test mocked function
func CreateDefinitionMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// to json
	dOut, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Post", APIPathPolicyDefinitions, mapIn).Return(dOut, 200, nil)
	policyDefinitionOut, err := ds.CreateDefinition(mapIn)

	assert.Nil(err, "Error creating policy definition")
	assert.Equal(policyDefinitionIn, policyDefinitionOut, "CreateDefinition returned different policy definition")

	return policyDefinitionOut
}

// CreateDefinitionFailErrMocked test mocked function
func CreateDefinitionFailErrMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// to json
	dOut, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Post", APIPathPolicyDefinitions, mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	policyDefinitionOut, err := ds.CreateDefinition(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyDefinitionOut
}

// CreateDefinitionFailStatusMocked test mocked function
func CreateDefinitionFailStatusMocked(
	t *testing.T,
	policyDefinitionIn *types.PolicyDefinition,
) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// to json
	dOut, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Post", APIPathPolicyDefinitions, mapIn).Return(dOut, 499, nil)
	policyDefinitionOut, err := ds.CreateDefinition(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyDefinitionOut
}

// CreateDefinitionFailJSONMocked test mocked function
func CreateDefinitionFailJSONMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", APIPathPolicyDefinitions, mapIn).Return(dIn, 200, nil)
	policyDefinitionOut, err := ds.CreateDefinition(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyDefinitionOut
}

// UpdateDefinitionMocked test mocked function
func UpdateDefinitionMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID), mapIn).Return(dIn, 200, nil)
	policyDefinitionOut, err := ds.UpdateDefinition(policyDefinitionIn.ID, mapIn)

	assert.Nil(err, "Error parsing policy definition metadata")
	assert.Equal(*policyDefinitionIn, *policyDefinitionOut, "UpdateDefinition returned different policy definition")

	return policyDefinitionOut
}

// UpdateDefinitionFailErrMocked test mocked function
func UpdateDefinitionFailErrMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID), mapIn).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	policyDefinitionOut, err := ds.UpdateDefinition(policyDefinitionIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyDefinitionOut
}

// UpdateDefinitionFailStatusMocked test mocked function
func UpdateDefinitionFailStatusMocked(
	t *testing.T,
	policyDefinitionIn *types.PolicyDefinition,
) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID), mapIn).Return(dIn, 499, nil)
	policyDefinitionOut, err := ds.UpdateDefinition(policyDefinitionIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyDefinitionOut
}

// UpdateDefinitionFailJSONMocked test mocked function
func UpdateDefinitionFailJSONMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) *types.PolicyDefinition {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID), mapIn).Return(dIn, 200, nil)
	policyDefinitionOut, err := ds.UpdateDefinition(policyDefinitionIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyDefinitionOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyDefinitionOut
}

// DeleteDefinitionMocked test mocked function
func DeleteDefinitionMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteDefinition(policyDefinitionIn.ID)

	assert.Nil(err, "Error deleting policy definition")
}

// DeleteDefinitionFailErrMocked test mocked function
func DeleteDefinitionFailErrMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteDefinition(policyDefinitionIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteDefinitionFailStatusMocked test mocked function
func DeleteDefinitionFailStatusMocked(t *testing.T, policyDefinitionIn *types.PolicyDefinition) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyDefinitionIn)
	assert.Nil(err, "PolicyDefinition test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathPolicyDefinition, policyDefinitionIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteDefinition(policyDefinitionIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// ListDefinitionAssignmentsMocked test mocked function
func ListDefinitionAssignmentsMocked(
	t *testing.T,
	definitionID string,
	policyAssignmentsIn []*types.PolicyAssignment,
) []*types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentsIn)
	assert.Nil(err, "PolicyAssignments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyDefinitionAssignments, definitionID)).Return(dIn, 200, nil)
	policyAssignmentsOut, err := ds.ListAssignments(definitionID)

	assert.Nil(err, "Error getting policy definitions")
	assert.Equal(policyAssignmentsIn, policyAssignmentsOut, "ListAssignments returned different policy definitions")

	return policyAssignmentsOut
}

// ListDefinitionAssignmentsFailErrMocked test mocked function
func ListDefinitionAssignmentsFailErrMocked(
	t *testing.T,
	definitionID string,
	policyAssignmentsIn []*types.PolicyAssignment,
) []*types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentsIn)
	assert.Nil(err, "PolicyAssignments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyDefinitionAssignments, definitionID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	policyAssignmentsOut, err := ds.ListAssignments(definitionID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyAssignmentsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyAssignmentsOut
}

// ListDefinitionAssignmentsFailStatusMocked test mocked function
func ListDefinitionAssignmentsFailStatusMocked(
	t *testing.T,
	definitionID string,
	policyAssignmentsIn []*types.PolicyAssignment,
) []*types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentsIn)
	assert.Nil(err, "PolicyAssignments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyDefinitionAssignments, definitionID)).Return(dIn, 499, nil)
	policyAssignmentsOut, err := ds.ListAssignments(definitionID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyAssignmentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyAssignmentsOut
}

// ListDefinitionAssignmentsFailJSONMocked test mocked function
func ListDefinitionAssignmentsFailJSONMocked(
	t *testing.T,
	definitionID string,
	policyAssignmentsIn []*types.PolicyAssignment,
) []*types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyDefinitionService(cs)
	assert.Nil(err, "Couldn't load policyDefinition service")
	assert.NotNil(ds, "PolicyDefinition service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyDefinitionAssignments, definitionID)).Return(dIn, 200, nil)
	policyAssignmentsOut, err := ds.ListAssignments(definitionID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyAssignmentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyAssignmentsOut
}
