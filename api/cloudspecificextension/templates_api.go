// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudspecificextension

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCseTemplates = "/cse/templates"
const APIPathCseTemplate = "/cse/templates/%s"

// CloudSpecificExtensionTemplateService manages cloud specific extension template operations
type CloudSpecificExtensionTemplateService struct {
	concertoService utils.ConcertoService
}

// NewCloudSpecificExtensionTemplateService returns a Concerto cloud specific extension template service
func NewCloudSpecificExtensionTemplateService(
	concertoService utils.ConcertoService,
) (*CloudSpecificExtensionTemplateService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudSpecificExtensionTemplateService{
		concertoService: concertoService,
	}, nil
}

// ListTemplates returns the list of cloud specific extension templates as an array of CloudSpecificExtensionTemplate
func (csets *CloudSpecificExtensionTemplateService) ListTemplates() (
	templates []*types.CloudSpecificExtensionTemplate, err error,
) {
	log.Debug("ListTemplates")

	data, status, err := csets.concertoService.Get(APIPathCseTemplates)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templates); err != nil {
		return nil, err
	}

	return templates, nil
}

// GetTemplate returns a cloud specific extension template by its ID
func (csets *CloudSpecificExtensionTemplateService) GetTemplate(
	templateID string,
) (template *types.CloudSpecificExtensionTemplate, err error) {
	log.Debug("GetTemplate")

	data, status, err := csets.concertoService.Get(fmt.Sprintf(APIPathCseTemplate, templateID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &template); err != nil {
		return nil, err
	}

	return template, nil
}

// CreateTemplate creates a cloud specific extension template
func (csets *CloudSpecificExtensionTemplateService) CreateTemplate(
	templateParams *map[string]interface{},
) (template *types.CloudSpecificExtensionTemplate, err error) {
	log.Debug("CreateTemplate")

	data, status, err := csets.concertoService.Post(APIPathCseTemplates, templateParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &template); err != nil {
		return nil, err
	}

	return template, nil
}

// UpdateTemplate updates a cloud specific extension template by its ID
func (csets *CloudSpecificExtensionTemplateService) UpdateTemplate(
	templateID string,
	templateParams *map[string]interface{},
) (template *types.CloudSpecificExtensionTemplate, err error) {
	log.Debug("UpdateTemplate")

	data, status, err := csets.concertoService.Put(fmt.Sprintf(APIPathCseTemplate, templateID), templateParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &template); err != nil {
		return nil, err
	}

	return template, nil
}

// ListDeployments returns the list of cloud specific extension deployments for a CSE template as an array of
// CloudSpecificExtensionDeployment
func (csets *CloudSpecificExtensionTemplateService) ListDeployments(
	templateID string,
) (deployments []*types.CloudSpecificExtensionDeployment, err error) {
	log.Debug("ListDeployments")

	data, status, err := csets.concertoService.Get(fmt.Sprintf(APIPathCseTemplateDeployments, templateID))
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

// DeleteTemplate deletes a cloud specific extension template by its ID
func (csets *CloudSpecificExtensionTemplateService) DeleteTemplate(templateID string) (err error) {
	log.Debug("DeleteTemplate")

	data, status, err := csets.concertoService.Delete(fmt.Sprintf(APIPathCseTemplate, templateID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
