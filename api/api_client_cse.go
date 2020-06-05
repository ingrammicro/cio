// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"golang.org/x/net/context"
)

// ListCloudSpecificExtensionDeployments returns the list of cloud specific extension deployments as an array of
// CloudSpecificExtensionDeployment
func (imco *ClientAPI) ListCloudSpecificExtensionDeployments(ctx context.Context) (
	deployments []*types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathCseDeployments, true, &deployments)
	if err != nil {
		return nil, err
	}
	return deployments, nil
}

// GetCloudSpecificExtensionDeployment returns a cloud specific extension deployment by its ID
func (imco *ClientAPI) GetCloudSpecificExtensionDeployment(ctx context.Context, deploymentID string,
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCseDeployment, deploymentID), true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// CreateCloudSpecificExtensionDeployment creates a cloud specific extension deployment
func (imco *ClientAPI) CreateCloudSpecificExtensionDeployment(ctx context.Context, templateID string,
	deploymentParams *map[string]interface{},
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(
		ctx,
		fmt.Sprintf(pathCseTemplateDeployments, templateID),
		deploymentParams,
		true,
		&deployment,
	)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// UpdateCloudSpecificExtensionDeployment updates a cloud specific extension deployment by its ID
func (imco *ClientAPI) UpdateCloudSpecificExtensionDeployment(ctx context.Context, deploymentID string,
	deploymentParams *map[string]interface{},
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathCseDeployment, deploymentID), deploymentParams, true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// DeleteCloudSpecificExtensionDeployment deletes a cloud specific extension deployment by its ID
func (imco *ClientAPI) DeleteCloudSpecificExtensionDeployment(ctx context.Context, deploymentID string,
) (deployment *types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathCseDeployment, deploymentID), true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// ListCloudSpecificExtensionTemplates returns the list of cloud specific extension templates as an array of
// CloudSpecificExtensionTemplate
func (imco *ClientAPI) ListCloudSpecificExtensionTemplates(ctx context.Context) (
	templates []*types.CloudSpecificExtensionTemplate, err error,
) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathCseTemplates, true, &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

// GetCloudSpecificExtensionTemplate returns a cloud specific extension template by its ID
func (imco *ClientAPI) GetCloudSpecificExtensionTemplate(ctx context.Context, templateID string,
) (template *types.CloudSpecificExtensionTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCseTemplate, templateID), true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CreateCloudSpecificExtensionTemplate creates a cloud specific extension template
func (imco *ClientAPI) CreateCloudSpecificExtensionTemplate(ctx context.Context,
	templateParams *map[string]interface{},
) (template *types.CloudSpecificExtensionTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathCseTemplates, templateParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// UpdateCloudSpecificExtensionTemplate updates a cloud specific extension template by its ID
func (imco *ClientAPI) UpdateCloudSpecificExtensionTemplate(ctx context.Context, templateID string,
	templateParams *map[string]interface{},
) (template *types.CloudSpecificExtensionTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathCseTemplate, templateID), templateParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// ListCloudSpecificExtensionTemplateDeployments returns the list of cloud specific extension deployments for a CSE
// template as an array of CloudSpecificExtensionDeployment
func (imco *ClientAPI) ListCloudSpecificExtensionTemplateDeployments(ctx context.Context, templateID string,
) (deployments []*types.CloudSpecificExtensionDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathCseTemplateDeployments, templateID), true, &deployments)
	if err != nil {
		return nil, err
	}
	return deployments, nil
}

// DeleteCloudSpecificExtensionTemplate deletes a cloud specific extension template by its ID
func (imco *ClientAPI) DeleteCloudSpecificExtensionTemplate(ctx context.Context, templateID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathCseTemplate, templateID), true, nil)
	if err != nil {
		return err
	}
	return nil
}
