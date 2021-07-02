// Copyright (c) 2017-2021 Ingram Micro Inc.

package settings

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewPolicyAssignmentServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewPolicyAssignmentService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListAssignments(t *testing.T) {
	policyAssignmentsIn := testdata.GetPolicyAssignmentData()
	ListAssignmentsMocked(t, policyAssignmentsIn[0].CloudAccountID, policyAssignmentsIn)
	ListAssignmentsFailErrMocked(t, policyAssignmentsIn[0].CloudAccountID, policyAssignmentsIn)
	ListAssignmentsFailStatusMocked(t, policyAssignmentsIn[0].CloudAccountID, policyAssignmentsIn)
	ListAssignmentsFailJSONMocked(t, policyAssignmentsIn[0].CloudAccountID, policyAssignmentsIn)
}

func TestGetAssignment(t *testing.T) {
	policyAssignmentsIn := testdata.GetPolicyAssignmentData()
	for _, policyAssignmentIn := range policyAssignmentsIn {
		GetAssignmentMocked(t, policyAssignmentIn)
		GetAssignmentFailErrMocked(t, policyAssignmentIn)
		GetAssignmentFailStatusMocked(t, policyAssignmentIn)
		GetAssignmentFailJSONMocked(t, policyAssignmentIn)
	}
}

func TestCreateAssignment(t *testing.T) {
	policyAssignmentsIn := testdata.GetPolicyAssignmentData()
	for _, policyAssignmentIn := range policyAssignmentsIn {
		CreateAssignmentMocked(t, policyAssignmentIn)
		CreateAssignmentFailErrMocked(t, policyAssignmentIn)
		CreateAssignmentFailStatusMocked(t, policyAssignmentIn)
		CreateAssignmentFailJSONMocked(t, policyAssignmentIn)
	}
}

func TestUpdateAssignment(t *testing.T) {
	policyAssignmentsIn := testdata.GetPolicyAssignmentData()
	for _, policyAssignmentIn := range policyAssignmentsIn {
		UpdateAssignmentMocked(t, policyAssignmentIn)
		UpdateAssignmentFailErrMocked(t, policyAssignmentIn)
		UpdateAssignmentFailStatusMocked(t, policyAssignmentIn)
		UpdateAssignmentFailJSONMocked(t, policyAssignmentIn)
	}
}

func TestDeleteAssignment(t *testing.T) {
	policyAssignmentsIn := testdata.GetPolicyAssignmentData()
	for _, policyAssignmentIn := range policyAssignmentsIn {
		DeleteAssignmentMocked(t, policyAssignmentIn)
		DeleteAssignmentFailErrMocked(t, policyAssignmentIn)
		DeleteAssignmentFailStatusMocked(t, policyAssignmentIn)
		DeleteAssignmentFailJSONMocked(t, policyAssignmentIn)
	}
}
