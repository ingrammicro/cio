package cloudapplication

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// CloudApplicationTemplateService manages cloud application template operations
type CloudApplicationTemplateService struct {
	concertoService utils.ConcertoService
}

// NewCloudApplicationTemplateService returns a Concerto cloud application template service
func NewCloudApplicationTemplateService(concertoService utils.ConcertoService) (*CloudApplicationTemplateService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudApplicationTemplateService{
		concertoService: concertoService,
	}, nil
}

// ListTemplates returns the list of cloud application templates as an array of CloudApplicationTemplate
func (cats *CloudApplicationTemplateService) ListTemplates() (templates []*types.CloudApplicationTemplate, err error) {
	log.Debug("ListTemplates")

	data, status, err := cats.concertoService.Get("/plugins/tosca/cats")
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

// GetTemplate returns a cloud application template by its ID
func (cats *CloudApplicationTemplateService) GetTemplate(templateID string) (template *types.CloudApplicationTemplate, err error) {
	log.Debug("GetTemplate")

	data, status, err := cats.concertoService.Get(fmt.Sprintf("/plugins/tosca/cats/%s", templateID))
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

// CreateTemplate creates a cloud application template
func (cats *CloudApplicationTemplateService) CreateTemplate(catVector *map[string]interface{}) (template *types.CloudApplicationTemplate, err error) {
	log.Debug("CreateTemplate")

	data, status, err := cats.concertoService.Post("/plugins/tosca/cats", catVector)
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

// UploadTemplate uploads a cloud application template file
func (cats *CloudApplicationTemplateService) UploadTemplate(sourceFilePath string, targetURL string) error {
	log.Debug("UploadTemplate")

	data, status, err := cats.concertoService.PutFile(sourceFilePath, targetURL)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ParseMetadataTemplate process cloud application template metadata
func (cats *CloudApplicationTemplateService) ParseMetadataTemplate(templateID string) (template *types.CloudApplicationTemplate, err error) {
	log.Debug("ParseMetadataTemplate")

	catIn := map[string]interface{}{}
	data, status, err := cats.concertoService.Put(fmt.Sprintf("/plugins/tosca/cats/%s/parse_metadata", templateID), &catIn)
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

// DeleteTemplate deletes a cloud application template by its ID
func (cats *CloudApplicationTemplateService) DeleteTemplate(templateID string) (err error) {
	log.Debug("DeleteTemplate")

	data, status, err := cats.concertoService.Delete(fmt.Sprintf("/plugins/tosca/cats/%s", templateID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
