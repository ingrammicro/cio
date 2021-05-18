// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudapplication

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathDeploymentLabels = "/labels"
const APIPathPluginsToscaDeployment = "/plugins/tosca/deployments/%s"
const APIPathPluginsToscaCatDeploymentTasks = "/plugins/tosca/cats/%s/deployment_tasks"
const APIPathPluginsToscaCatDeploymentTask = "/plugins/tosca/cats/%s/deployment_tasks/%s"

// CloudApplicationDeploymentService manages cloud application deployment operations
type CloudApplicationDeploymentService struct {
	concertoService utils.ConcertoService
}

// NewCloudApplicationDeploymentService returns a Concerto cloud application deployment service
func NewCloudApplicationDeploymentService(
	concertoService utils.ConcertoService,
) (*CloudApplicationDeploymentService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudApplicationDeploymentService{
		concertoService: concertoService,
	}, nil
}

// ListDeployments returns the list of cloud application deployments as an array of CloudApplicationDeployment
func (cads *CloudApplicationDeploymentService) ListDeployments() (
	deployments []*types.CloudApplicationDeployment, err error,
) {
	log.Debug("ListDeployments")

	data, status, err := cads.concertoService.Get(APIPathDeploymentLabels)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &deployments); err != nil {
		return nil, err
	}

	// Only takes internal labels (with a Namespace defined as cat:deployment)
	var filteredDeployments []*types.CloudApplicationDeployment
	for _, dep := range deployments {
		if dep.Namespace == "cat:deployment" {
			filteredDeployments = append(filteredDeployments, dep)
		}
	}

	return filteredDeployments, nil
}

// GetDeployment returns a cloud application deployment by its ID
func (cads *CloudApplicationDeploymentService) GetDeployment(
	deploymentID string,
) (deployment *types.CloudApplicationDeployment, status int, err error) {
	log.Debug("GetDeployment")

	data, status, err := cads.concertoService.Get(fmt.Sprintf(APIPathPluginsToscaDeployment, deploymentID))
	if err != nil {
		return nil, status, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &deployment); err != nil {
		return nil, status, err
	}

	return deployment, status, nil
}

// DeleteDeployment deletes a cloud application deployment by its ID
func (cads *CloudApplicationDeploymentService) DeleteDeployment(
	deploymentID string,
) (deployment *types.CloudApplicationDeployment, err error) {
	log.Debug("DeleteDeployment")

	data, status, err := cads.concertoService.Delete(fmt.Sprintf(APIPathPluginsToscaDeployment, deploymentID))
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

// CreateDeploymentTask creates a cloud application deployment task by a given CAT ID
func (cads *CloudApplicationDeploymentService) CreateDeploymentTask(
	catID string,
	deploymentParams *map[string]interface{},
) (deploymentTask *types.CloudApplicationDeploymentTask, err error) {
	log.Debug("CreateDeploymentTask")

	data, status, err := cads.concertoService.Post(
		fmt.Sprintf(APIPathPluginsToscaCatDeploymentTasks, catID),
		deploymentParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &deploymentTask); err != nil {
		return nil, err
	}

	return deploymentTask, nil
}

// GetDeploymentTask gets a cloud application deployment task by its ID and given CAT ID
func (cads *CloudApplicationDeploymentService) GetDeploymentTask(
	catID string,
	deploymentTaskID string,
) (deploymentTask *types.CloudApplicationDeploymentTask, err error) {
	log.Debug("GetDeploymentTask")

	data, status, err := cads.concertoService.Get(
		fmt.Sprintf(APIPathPluginsToscaCatDeploymentTask, catID, deploymentTaskID),
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &deploymentTask); err != nil {
		return nil, err
	}

	return deploymentTask, nil
}
