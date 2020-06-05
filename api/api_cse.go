// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListCloudSpecificExtensionDeployments returns the list of cloud specific extension deployments as an array of
// CloudSpecificExtensionDeployment
func (imco *IMCOClient) ListCloudSpecificExtensionDeployments() (
	deployments []*types.CloudSpecificExtensionDeployment, err error,
) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathCseDeployments, true, &deployments)
	if err != nil {
		return nil, err
	}
	return deployments, nil
}

// GetCloudSpecificExtensionDeployment returns a cloud specific extension deployment by its ID
func (imco *IMCOClient) GetCloudSpecificExtensionDeployment(deploymentID string,
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCseDeployment, deploymentID), true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// CreateCloudSpecificExtensionDeployment creates a cloud specific extension deployment
func (imco *IMCOClient) CreateCloudSpecificExtensionDeployment(templateID string,
	deploymentParams *map[string]interface{},
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathCseTemplateDeployments, templateID), deploymentParams, true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// UpdateCloudSpecificExtensionDeployment updates a cloud specific extension deployment by its ID
func (imco *IMCOClient) UpdateCloudSpecificExtensionDeployment(deploymentID string,
	deploymentParams *map[string]interface{},
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCseDeployment, deploymentID), deploymentParams, true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// DeleteCloudSpecificExtensionDeployment deletes a cloud specific extension deployment by its ID
func (imco *IMCOClient) DeleteCloudSpecificExtensionDeployment(deploymentID string,
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathCseDeployment, deploymentID), true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// ListCloudSpecificExtensionTemplates returns the list of cloud specific extension templates as an array of
// CloudSpecificExtensionTemplate
func (imco *IMCOClient) ListCloudSpecificExtensionTemplates() (
	templates []*types.CloudSpecificExtensionTemplate, err error,
) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathCseTemplates, true, &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

// GetCloudSpecificExtensionTemplate returns a cloud specific extension template by its ID
func (imco *IMCOClient) GetCloudSpecificExtensionTemplate(templateID string,
) (template *types.CloudSpecificExtensionTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCseTemplate, templateID), true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CreateCloudSpecificExtensionTemplate creates a cloud specific extension template
func (imco *IMCOClient) CreateCloudSpecificExtensionTemplate(templateParams *map[string]interface{},
) (template *types.CloudSpecificExtensionTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathCseTemplates, templateParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// UpdateCloudSpecificExtensionTemplate updates a cloud specific extension template by its ID
func (imco *IMCOClient) UpdateCloudSpecificExtensionTemplate(templateID string,
	templateParams *map[string]interface{},
) (template *types.CloudSpecificExtensionTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathCseTemplate, templateID), templateParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// ListCloudSpecificExtensionTemplateDeployments returns the list of cloud specific extension deployments for a CSE
// template as an array of CloudSpecificExtensionDeployment
func (imco *IMCOClient) ListCloudSpecificExtensionTemplateDeployments(templateID string,
) (deployments []*types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathCseTemplateDeployments, templateID), true, &deployments)
	if err != nil {
		return nil, err
	}
	return deployments, nil
}

// DeleteCloudSpecificExtensionTemplate deletes a cloud specific extension template by its ID
func (imco *IMCOClient) DeleteCloudSpecificExtensionTemplate(templateID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathCseTemplate, templateID), true, nil)
	if err != nil {
		return err
	}
	return nil
}
