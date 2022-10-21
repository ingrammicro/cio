// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListLabels returns the list of labels as an array of Label
func (imco *ClientAPI) ListLabels(ctx context.Context) (labels []*types.Label, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathLabels, true, &labels)
	if err != nil {
		return nil, err
	}
	// exclude internal labels (with a Namespace defined)
	var filteredLabels []*types.Label
	for _, label := range labels {
		if label.Namespace == "" {
			filteredLabels = append(filteredLabels, label)
		}
	}
	return filteredLabels, nil
}

// CreateLabel creates a label
func (imco *ClientAPI) CreateLabel(ctx context.Context, labelParams *map[string]interface{},
) (label *types.Label, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, PathLabels, labelParams, true, &label)
	if err != nil {
		return nil, err
	}
	return label, nil
}

// AddLabel assigns a single label from a single labelable resource
func (imco *ClientAPI) AddLabel(ctx context.Context, labelID string, labelParams *map[string]interface{},
) (labeledResources []*types.LabeledResource, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, fmt.Sprintf(PathLabelResources, labelID), labelParams, true, &labeledResources)
	if err != nil {
		return nil, err
	}
	return labeledResources, nil
}

// RemoveLabel de-assigns a single label from a single labelable resource
func (imco *ClientAPI) RemoveLabel(ctx context.Context, labelID string, resourceType string, resourceID string,
) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathLabelResource, labelID, resourceType, resourceID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListCloudApplicationDeployments returns the list of cloud application deployments as an array of
// CloudApplicationDeployment
func (imco *ClientAPI) ListCloudApplicationDeployments(ctx context.Context,
) (deployments []*types.CloudApplicationDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathLabels, true, &deployments)
	if err != nil {
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
