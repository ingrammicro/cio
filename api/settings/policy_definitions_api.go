package settings

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// PolicyDefinitionService manages policy definition operations
type PolicyDefinitionService struct {
	concertoService utils.ConcertoService
}

// NewPolicyDefinitionService returns a Concerto policy definition service
func NewPolicyDefinitionService(concertoService utils.ConcertoService) (*PolicyDefinitionService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &PolicyDefinitionService{
		concertoService: concertoService,
	}, nil
}

// ListDefinitions returns the list of policy definitions as an array of PolicyDefinition
func (pds *PolicyDefinitionService) ListDefinitions() (definitions []*types.PolicyDefinition, err error) {
	log.Debug("ListDefinitions")

	data, status, err := pds.concertoService.Get("/policy/definitions")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &definitions); err != nil {
		return nil, err
	}

	return definitions, nil
}

// GetDefinition returns a policy definition by its ID
func (pds *PolicyDefinitionService) GetDefinition(definitionID string) (definition *types.PolicyDefinition, err error) {
	log.Debug("GetDefinition")

	data, status, err := pds.concertoService.Get(fmt.Sprintf("/policy/definitions/%s", definitionID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &definition); err != nil {
		return nil, err
	}

	return definition, nil
}

// CreateDefinition creates a policy definition
func (pds *PolicyDefinitionService) CreateDefinition(definitionParams *map[string]interface{}) (definition *types.PolicyDefinition, err error) {
	log.Debug("CreateDefinition")

	data, status, err := pds.concertoService.Post("/policy/definitions", definitionParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &definition); err != nil {
		return nil, err
	}

	return definition, nil
}

// UpdateDefinition updates a policy definition by its ID
func (pds *PolicyDefinitionService) UpdateDefinition(definitionID string, definitionParams *map[string]interface{}) (definition *types.PolicyDefinition, err error) {
	log.Debug("UpdateDefinition")

	data, status, err := pds.concertoService.Put(fmt.Sprintf("/policy/definitions/%s", definitionID), definitionParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &definition); err != nil {
		return nil, err
	}

	return definition, nil
}

// DeleteDefinition deletes a policy definition by its ID
func (pds *PolicyDefinitionService) DeleteDefinition(definitionID string) (err error) {
	log.Debug("DeleteDefinition")

	data, status, err := pds.concertoService.Delete(fmt.Sprintf("/policy/definitions/%s", definitionID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ListAssignments returns the list of policy assignments for a policy definition as an array of PolicyAssignment
func (pds *PolicyDefinitionService) ListAssignments(definitionID string) (assignments []*types.PolicyAssignment, err error) {
	log.Debug("ListAssignments")

	data, status, err := pds.concertoService.Get(fmt.Sprintf("/policy/definitions/%s/assignments", definitionID))
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
