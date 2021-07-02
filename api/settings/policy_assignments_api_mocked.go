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

// ListAssignmentsMocked test mocked function
func ListAssignmentsMocked(
	t *testing.T,
	cloudAccountID string,
	policyAssignmentsIn []*types.PolicyAssignment,
) []*types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentsIn)
	assert.Nil(err, "PolicyAssignments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathSettingsCloudAccountPolicyAssignments, cloudAccountID)).Return(dIn, 200, nil)
	policyAssignmentsOut, err := ds.ListAssignments(cloudAccountID)

	assert.Nil(err, "Error getting policy assignments")
	assert.Equal(policyAssignmentsIn, policyAssignmentsOut, "ListAssignments returned different policy assignments")

	return policyAssignmentsOut
}

// ListAssignmentsFailErrMocked test mocked function
func ListAssignmentsFailErrMocked(
	t *testing.T,
	cloudAccountID string,
	policyAssignmentsIn []*types.PolicyAssignment,
) []*types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentsIn)
	assert.Nil(err, "PolicyAssignments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathSettingsCloudAccountPolicyAssignments, cloudAccountID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	policyAssignmentsOut, err := ds.ListAssignments(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyAssignmentsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyAssignmentsOut
}

// ListAssignmentsFailStatusMocked test mocked function
func ListAssignmentsFailStatusMocked(
	t *testing.T,
	cloudAccountID string,
	policyAssignmentsIn []*types.PolicyAssignment,
) []*types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentsIn)
	assert.Nil(err, "PolicyAssignments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathSettingsCloudAccountPolicyAssignments, cloudAccountID)).Return(dIn, 499, nil)
	policyAssignmentsOut, err := ds.ListAssignments(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyAssignmentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyAssignmentsOut
}

// ListAssignmentsFailJSONMocked test mocked function
func ListAssignmentsFailJSONMocked(
	t *testing.T,
	cloudAccountID string,
	policyAssignmentsIn []*types.PolicyAssignment,
) []*types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathSettingsCloudAccountPolicyAssignments, cloudAccountID)).Return(dIn, 200, nil)
	policyAssignmentsOut, err := ds.ListAssignments(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyAssignmentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyAssignmentsOut
}

// GetAssignmentMocked test mocked function
func GetAssignmentMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID)).Return(dIn, 200, nil)
	policyAssignmentOut, err := ds.GetAssignment(policyAssignmentIn.ID)

	assert.Nil(err, "Error getting policy assignment")
	assert.Equal(*policyAssignmentIn, *policyAssignmentOut, "GetAssignment returned different policy assignment")

	return policyAssignmentOut
}

// GetAssignmentFailErrMocked test mocked function
func GetAssignmentFailErrMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	policyAssignmentOut, err := ds.GetAssignment(policyAssignmentIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyAssignmentOut
}

// GetAssignmentFailStatusMocked test mocked function
func GetAssignmentFailStatusMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID)).Return(dIn, 499, nil)
	policyAssignmentOut, err := ds.GetAssignment(policyAssignmentIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyAssignmentOut
}

// GetAssignmentFailJSONMocked test mocked function
func GetAssignmentFailJSONMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID)).Return(dIn, 200, nil)
	policyAssignmentOut, err := ds.GetAssignment(policyAssignmentIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyAssignmentOut
}

// CreateAssignmentMocked test mocked function
func CreateAssignmentMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// to json
	dOut, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathPolicyDefinitionAssignments, policyAssignmentIn.DefinitionID), mapIn).
		Return(dOut, 200, nil)
	policyAssignmentOut, err := ds.CreateAssignment(policyAssignmentIn.DefinitionID, mapIn)

	assert.Nil(err, "Error creating policy assignment")
	assert.Equal(policyAssignmentIn, policyAssignmentOut, "CreateAssignment returned different policy assignment")

	return policyAssignmentOut
}

// CreateAssignmentFailErrMocked test mocked function
func CreateAssignmentFailErrMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// to json
	dOut, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathPolicyDefinitionAssignments, policyAssignmentIn.DefinitionID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	policyAssignmentOut, err := ds.CreateAssignment(policyAssignmentIn.DefinitionID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyAssignmentOut
}

// CreateAssignmentFailStatusMocked test mocked function
func CreateAssignmentFailStatusMocked(
	t *testing.T,
	policyAssignmentIn *types.PolicyAssignment,
) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// to json
	dOut, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathPolicyDefinitionAssignments, policyAssignmentIn.DefinitionID), mapIn).
		Return(dOut, 499, nil)
	policyAssignmentOut, err := ds.CreateAssignment(policyAssignmentIn.DefinitionID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyAssignmentOut
}

// CreateAssignmentFailJSONMocked test mocked function
func CreateAssignmentFailJSONMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathPolicyDefinitionAssignments, policyAssignmentIn.DefinitionID), mapIn).
		Return(dIn, 200, nil)
	policyAssignmentOut, err := ds.CreateAssignment(policyAssignmentIn.DefinitionID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyAssignmentOut
}

// UpdateAssignmentMocked test mocked function
func UpdateAssignmentMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID), mapIn).Return(dIn, 200, nil)
	policyAssignmentOut, err := ds.UpdateAssignment(policyAssignmentIn.ID, mapIn)

	assert.Nil(err, "Error updating assignment")
	assert.Equal(*policyAssignmentIn, *policyAssignmentOut, "UpdateAssignment returned different policy assignment")

	return policyAssignmentOut
}

// UpdateAssignmentFailErrMocked test mocked function
func UpdateAssignmentFailErrMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID), mapIn).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	policyAssignmentOut, err := ds.UpdateAssignment(policyAssignmentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return policyAssignmentOut
}

// UpdateAssignmentFailStatusMocked test mocked function
func UpdateAssignmentFailStatusMocked(
	t *testing.T,
	policyAssignmentIn *types.PolicyAssignment,
) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID), mapIn).Return(dIn, 499, nil)
	policyAssignmentOut, err := ds.UpdateAssignment(policyAssignmentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return policyAssignmentOut
}

// UpdateAssignmentFailJSONMocked test mocked function
func UpdateAssignmentFailJSONMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID), mapIn).Return(dIn, 200, nil)
	policyAssignmentOut, err := ds.UpdateAssignment(policyAssignmentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyAssignmentOut
}

// DeleteAssignmentMocked test mocked function
func DeleteAssignmentMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID)).Return(dIn, 200, nil)
	policyAssignmentOut, err := ds.DeleteAssignment(policyAssignmentIn.ID)

	assert.Nil(err, "Error deleting policy assignment")
	assert.Equal(policyAssignmentIn, policyAssignmentOut, "DeleteAssignment returned different assignment")
}

// DeleteAssignmentFailErrMocked test mocked function
func DeleteAssignmentFailErrMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	policyAssignmentOut, err := ds.DeleteAssignment(policyAssignmentIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteAssignmentFailStatusMocked test mocked function
func DeleteAssignmentFailStatusMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// to json
	dIn, err := json.Marshal(policyAssignmentIn)
	assert.Nil(err, "PolicyAssignment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID)).Return(dIn, 499, nil)
	policyAssignmentOut, err := ds.DeleteAssignment(policyAssignmentIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteAssignmentFailJSONMocked test mocked function
func DeleteAssignmentFailJSONMocked(t *testing.T, policyAssignmentIn *types.PolicyAssignment) *types.PolicyAssignment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewPolicyAssignmentService(cs)
	assert.Nil(err, "Couldn't load policyAssignment service")
	assert.NotNil(ds, "PolicyAssignment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathPolicyAssignment, policyAssignmentIn.ID)).Return(dIn, 200, nil)
	policyAssignmentOut, err := ds.DeleteAssignment(policyAssignmentIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(policyAssignmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return policyAssignmentOut
}
