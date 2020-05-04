package cloudspecificextension

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// CloudSpecificExtensionDeploymentService manages cloud specific extension deployment operations
type CloudSpecificExtensionDeploymentService struct {
	concertoService utils.ConcertoService
}

// NewCloudSpecificExtensionDeploymentService returns a Concerto cloud specific extension deployment service
func NewCloudSpecificExtensionDeploymentService(concertoService utils.ConcertoService) (*CloudSpecificExtensionDeploymentService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudSpecificExtensionDeploymentService{
		concertoService: concertoService,
	}, nil
}

// ListDeployments returns the list of cloud specific extension deployments as an array of CloudSpecificExtensionDeployment
func (csets *CloudSpecificExtensionDeploymentService) ListDeployments() (deployments []*types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("ListDeployments")

	data, status, err := csets.concertoService.Get("/cse/deployments")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &deployments); err != nil {
		return nil, err
	}

	return deployments, nil
}

// GetDeployment returns a cloud specific extension deployment by its ID
func (csets *CloudSpecificExtensionDeploymentService) GetDeployment(deploymentID string) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("GetDeployment")

	data, status, err := csets.concertoService.Get(fmt.Sprintf("/cse/deployments/%s", deploymentID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &deployment); err != nil {
		return nil, err
	}

	return deployment, nil
}

// CreateDeployment creates a cloud specific extension deployment
func (csets *CloudSpecificExtensionDeploymentService) CreateDeployment(templateID string, deploymentParams *map[string]interface{}) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("CreateDeployment")

	data, status, err := csets.concertoService.Post(fmt.Sprintf("/cse/templates/%s/deployments", templateID), deploymentParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &deployment); err != nil {
		return nil, err
	}

	return deployment, nil
}

// UpdateDeployment updates a cloud specific extension deployment by its ID
func (csets *CloudSpecificExtensionDeploymentService) UpdateDeployment(deploymentParams *map[string]interface{}, deploymentID string) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("UpdateDeployment")

	data, status, err := csets.concertoService.Put(fmt.Sprintf("/cse/deployments/%s", deploymentID), deploymentParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &deployment); err != nil {
		return nil, err
	}

	return deployment, nil
}

// DeleteDeployment deletes a cloud specific extension deployment by its ID
func (csets *CloudSpecificExtensionDeploymentService) DeleteDeployment(deploymentID string) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("DeleteDeployment")

	data, status, err := csets.concertoService.Delete(fmt.Sprintf("/cse/deployments/%s", deploymentID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &deployment); err != nil {
		return nil, err
	}

	return deployment, nil
}
