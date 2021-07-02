// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudspecificextension

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCseDeployments = "/cse/deployments"
const APIPathCseDeployment = "/cse/deployments/%s"
const APIPathCseTemplateDeployments = "/cse/templates/%s/deployments"

// CloudSpecificExtensionDeploymentService manages cloud specific extension deployment operations
type CloudSpecificExtensionDeploymentService struct {
	concertoService utils.ConcertoService
}

// NewCloudSpecificExtensionDeploymentService returns a Concerto cloud specific extension deployment service
func NewCloudSpecificExtensionDeploymentService(
	concertoService utils.ConcertoService,
) (*CloudSpecificExtensionDeploymentService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudSpecificExtensionDeploymentService{
		concertoService: concertoService,
	}, nil
}

// ListDeployments returns the list of cloud specific extension deployments as an array of
// CloudSpecificExtensionDeployment
func (cseds *CloudSpecificExtensionDeploymentService) ListDeployments() (
	deployments []*types.CloudSpecificExtensionDeployment, err error,
) {
	log.Debug("ListDeployments")

	data, status, err := cseds.concertoService.Get(APIPathCseDeployments)
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
func (cseds *CloudSpecificExtensionDeploymentService) GetDeployment(
	deploymentID string,
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("GetDeployment")

	data, status, err := cseds.concertoService.Get(fmt.Sprintf(APIPathCseDeployment, deploymentID))
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
func (cseds *CloudSpecificExtensionDeploymentService) CreateDeployment(
	templateID string,
	deploymentParams *map[string]interface{},
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("CreateDeployment")

	data, status, err := cseds.concertoService.Post(
		fmt.Sprintf(APIPathCseTemplateDeployments, templateID),
		deploymentParams,
	)
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
func (cseds *CloudSpecificExtensionDeploymentService) UpdateDeployment(
	deploymentID string,
	deploymentParams *map[string]interface{},
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("UpdateDeployment")

	data, status, err := cseds.concertoService.Put(fmt.Sprintf(APIPathCseDeployment, deploymentID), deploymentParams)
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
func (cseds *CloudSpecificExtensionDeploymentService) DeleteDeployment(
	deploymentID string,
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("DeleteDeployment")

	data, status, err := cseds.concertoService.Delete(fmt.Sprintf(APIPathCseDeployment, deploymentID))
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
