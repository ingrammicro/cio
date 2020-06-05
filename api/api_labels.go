// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListLabels returns the list of labels as an array of Label
func (imco *IMCOClient) ListLabels() (labels []*types.Label, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathLabels, true, &labels)
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
func (imco *IMCOClient) CreateLabel(labelParams *map[string]interface{}) (label *types.Label, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathLabels, labelParams, true, &label)
	if err != nil {
		return nil, err
	}
	return label, nil
}

// AddLabel assigns a single label from a single labelable resource
func (imco *IMCOClient) AddLabel(labelID string, labelParams *map[string]interface{},
) (labeledResources []*types.LabeledResource, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathLabelResources, labelID), labelParams, true, &labeledResources)
	if err != nil {
		return nil, err
	}
	return labeledResources, nil
}

// RemoveLabel de-assigns a single label from a single labelable resource
func (imco *IMCOClient) RemoveLabel(labelID string, resourceType string, resourceID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathLabelResource, labelID, resourceType, resourceID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListCloudApplicationDeployments returns the list of cloud application deployments as an array of
// CloudApplicationDeployment
func (imco *IMCOClient) ListCloudApplicationDeployments() (deployments []*types.CloudApplicationDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathLabels, true, &deployments)
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
