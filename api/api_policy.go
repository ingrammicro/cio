// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListPolicyDefinitions returns the list of policy definitions as an array of PolicyDefinition
func (imco *IMCOClient) ListPolicyDefinitions() (definitions []*types.PolicyDefinition, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathPolicyDefinitions, true, &definitions)
	if err != nil {
		return nil, err
	}
	return definitions, nil
}

// GetPolicyDefinition returns a policy definition by its ID
func (imco *IMCOClient) GetPolicyDefinition(definitionID string) (definition *types.PolicyDefinition, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathPolicyDefinition, definitionID), true, &definition)
	if err != nil {
		return nil, err
	}
	return definition, nil
}

// CreatePolicyDefinition creates a policy definition
func (imco *IMCOClient) CreatePolicyDefinition(definitionParams *map[string]interface{},
) (definition *types.PolicyDefinition, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathPolicyDefinitions, definitionParams, true, &definition)
	if err != nil {
		return nil, err
	}
	return definition, nil
}

// UpdatePolicyDefinition updates a policy definition by its ID
func (imco *IMCOClient) UpdatePolicyDefinition(definitionID string, definitionParams *map[string]interface{},
) (definition *types.PolicyDefinition, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathPolicyDefinition, definitionID), definitionParams, true, &definition)
	if err != nil {
		return nil, err
	}
	return definition, nil
}

// DeletePolicyDefinition deletes a policy definition by its ID
func (imco *IMCOClient) DeletePolicyDefinition(definitionID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathPolicyDefinition, definitionID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListPolicyDefinitionAssignments returns the list of policy assignments for a policy definition as an array of
// PolicyAssignment
func (imco *IMCOClient) ListPolicyDefinitionAssignments(definitionID string,
) (assignments []*types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathPolicyDefinitionAssignments, definitionID), true, &assignments)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

// GetPolicyAssignment returns an assignment by its ID
func (imco *IMCOClient) GetPolicyAssignment(assignmentID string) (assignment *types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathPolicyAssignment, assignmentID), true, &assignment)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

// CreatePolicyAssignment creates an assignment
func (imco *IMCOClient) CreatePolicyAssignment(definitionID string, assignmentParams *map[string]interface{},
) (assignment *types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
		fmt.Sprintf(pathPolicyDefinitionAssignments, definitionID),
		assignmentParams,
		true,
		&assignment,
	)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

// UpdatePolicyAssignment updates an assignment by its ID
func (imco *IMCOClient) UpdatePolicyAssignment(assignmentID string, assignmentParams *map[string]interface{},
) (assignment *types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathPolicyAssignment, assignmentID), assignmentParams, true, &assignment)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

// DeletePolicyAssignment deletes an assignment by its ID
func (imco *IMCOClient) DeletePolicyAssignment(assignmentID string) (assignment *types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathPolicyAssignment, assignmentID), true, &assignment)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}
