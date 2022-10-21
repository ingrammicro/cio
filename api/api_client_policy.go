// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListPolicyDefinitions returns the list of policy definitions as an array of PolicyDefinition
func (imco *ClientAPI) ListPolicyDefinitions(ctx context.Context) (definitions []*types.PolicyDefinition, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathPolicyDefinitions, true, &definitions)
	if err != nil {
		return nil, err
	}
	return definitions, nil
}

// GetPolicyDefinition returns a policy definition by its ID
func (imco *ClientAPI) GetPolicyDefinition(ctx context.Context, definitionID string,
) (definition *types.PolicyDefinition, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathPolicyDefinition, definitionID), true, &definition)
	if err != nil {
		return nil, err
	}
	return definition, nil
}

// CreatePolicyDefinition creates a policy definition
func (imco *ClientAPI) CreatePolicyDefinition(ctx context.Context, definitionParams *map[string]interface{},
) (definition *types.PolicyDefinition, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, PathPolicyDefinitions, definitionParams, true, &definition)
	if err != nil {
		return nil, err
	}
	return definition, nil
}

// UpdatePolicyDefinition updates a policy definition by its ID
func (imco *ClientAPI) UpdatePolicyDefinition(ctx context.Context, definitionID string,
	definitionParams *map[string]interface{},
) (definition *types.PolicyDefinition, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathPolicyDefinition, definitionID), definitionParams, true, &definition)
	if err != nil {
		return nil, err
	}
	return definition, nil
}

// DeletePolicyDefinition deletes a policy definition by its ID
func (imco *ClientAPI) DeletePolicyDefinition(ctx context.Context, definitionID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathPolicyDefinition, definitionID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListPolicyDefinitionAssignments returns the list of policy assignments for a policy definition as an array of
// PolicyAssignment
func (imco *ClientAPI) ListPolicyDefinitionAssignments(ctx context.Context, definitionID string,
) (assignments []*types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathPolicyDefinitionAssignments, definitionID), true, &assignments)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

// GetPolicyAssignment returns an assignment by its ID
func (imco *ClientAPI) GetPolicyAssignment(ctx context.Context, assignmentID string,
) (assignment *types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathPolicyAssignment, assignmentID), true, &assignment)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

// CreatePolicyAssignment creates an assignment
func (imco *ClientAPI) CreatePolicyAssignment(ctx context.Context, definitionID string,
	assignmentParams *map[string]interface{},
) (assignment *types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(PathPolicyDefinitionAssignments, definitionID),
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
func (imco *ClientAPI) UpdatePolicyAssignment(ctx context.Context, assignmentID string,
	assignmentParams *map[string]interface{},
) (assignment *types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathPolicyAssignment, assignmentID), assignmentParams, true, &assignment)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}

// DeletePolicyAssignment deletes an assignment by its ID
func (imco *ClientAPI) DeletePolicyAssignment(ctx context.Context, assignmentID string,
) (assignment *types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathPolicyAssignment, assignmentID), true, &assignment)
	if err != nil {
		return nil, err
	}
	return assignment, nil
}
