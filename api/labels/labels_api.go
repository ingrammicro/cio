// Copyright (c) 2017-2021 Ingram Micro Inc.

package labels

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathLabels = "/labels"
const APIPathLabelResources = "/labels/%s/resources"
const APIPathLabelResource = "/labels/%s/resources/%s/%s"

// LabelService manages label operations
type LabelService struct {
	concertoService utils.ConcertoService
}

// NewLabelService returns a Concerto labels service
func NewLabelService(concertoService utils.ConcertoService) (*LabelService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &LabelService{
		concertoService: concertoService,
	}, nil
}

// ListLabels returns the list of labels as an array of Label
func (ls *LabelService) ListLabels() (labels []*types.Label, err error) {
	log.Debug("ListLabels")

	data, status, err := ls.concertoService.Get(APIPathLabels)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &labels); err != nil {
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
func (ls *LabelService) CreateLabel(labelParams *map[string]interface{}) (label *types.Label, err error) {
	log.Debug("CreateLabel")

	data, status, err := ls.concertoService.Post(APIPathLabels, labelParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &label); err != nil {
		return nil, err
	}

	return label, nil
}

// AddLabel assigns a single label from a single labelable resource
func (ls *LabelService) AddLabel(
	labelID string,
	labelParams *map[string]interface{},
) (labeledResources []*types.LabeledResource, err error) {
	log.Debug("AddLabel")

	data, status, err := ls.concertoService.Post(fmt.Sprintf(APIPathLabelResources, labelID), labelParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &labeledResources); err != nil {
		return nil, err
	}

	return labeledResources, nil
}

// RemoveLabel de-assigns a single label from a single labelable resource
func (ls *LabelService) RemoveLabel(labelID string, resourceType string, resourceID string) error {
	log.Debug("RemoveLabel")

	data, status, err := ls.concertoService.Delete(
		fmt.Sprintf(APIPathLabelResource, labelID, resourceType, resourceID),
	)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
