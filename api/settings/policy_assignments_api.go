// Copyright (c) 2017-2021 Ingram Micro Inc.

package settings

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathSettingsCloudAccountPolicyAssignments = "/settings/cloud_accounts/%s/policy_assignments"
const APIPathPolicyAssignment = "/policy/assignments/%s"

// PolicyAssignmentService manages policy assignment operations
type PolicyAssignmentService struct {
	concertoService utils.ConcertoService
}

// NewPolicyAssignmentService returns a Concerto policy assignment service
func NewPolicyAssignmentService(concertoService utils.ConcertoService) (*PolicyAssignmentService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &PolicyAssignmentService{
		concertoService: concertoService,
	}, nil
}

// ListAssignments returns the list of policy assignments as an array of PolicyAssignment
func (pas *PolicyAssignmentService) ListAssignments(
	cloudAccountID string,
) (assignments []*types.PolicyAssignment, err error) {
	log.Debug("ListAssignments")

	data, status, err := pas.concertoService.Get(
		fmt.Sprintf(APIPathSettingsCloudAccountPolicyAssignments, cloudAccountID),
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &assignments); err != nil {
		return nil, err
	}

	return assignments, nil
}

// GetAssignment returns an assignment by its ID
func (pas *PolicyAssignmentService) GetAssignment(assignmentID string) (assignment *types.PolicyAssignment, err error) {
	log.Debug("GetAssignment")

	data, status, err := pas.concertoService.Get(fmt.Sprintf(APIPathPolicyAssignment, assignmentID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &assignment); err != nil {
		return nil, err
	}

	return assignment, nil
}

// CreateAssignment creates an assignment
func (pas *PolicyAssignmentService) CreateAssignment(
	definitionID string,
	assignmentParams *map[string]interface{},
) (assignment *types.PolicyAssignment, err error) {
	log.Debug("CreateAssignment")

	data, status, err := pas.concertoService.Post(
		fmt.Sprintf(APIPathPolicyDefinitionAssignments, definitionID),
		assignmentParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &assignment); err != nil {
		return nil, err
	}

	return assignment, nil
}

// UpdateAssignment updates an assignment by its ID
func (pas *PolicyAssignmentService) UpdateAssignment(
	assignmentID string,
	assignmentParams *map[string]interface{},
) (assignment *types.PolicyAssignment, err error) {
	log.Debug("UpdateAssignment")

	data, status, err := pas.concertoService.Put(fmt.Sprintf(APIPathPolicyAssignment, assignmentID), assignmentParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &assignment); err != nil {
		return nil, err
	}

	return assignment, nil
}

// DeleteAssignment deletes an assignment by its ID
func (pas *PolicyAssignmentService) DeleteAssignment(
	assignmentID string,
) (assignment *types.PolicyAssignment, err error) {
	log.Debug("DeleteAssignment")

	data, status, err := pas.concertoService.Delete(fmt.Sprintf(APIPathPolicyAssignment, assignmentID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &assignment); err != nil {
		return nil, err
	}

	return assignment, nil
}
