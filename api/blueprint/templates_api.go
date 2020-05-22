package blueprint

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// TemplateService manages template operations
type TemplateService struct {
	concertoService utils.ConcertoService
}

// NewTemplateService returns a Concerto template service
func NewTemplateService(concertoService utils.ConcertoService) (*TemplateService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &TemplateService{
		concertoService: concertoService,
	}, nil
}

// ListTemplates returns the list of templates as an array of Template
func (tp *TemplateService) ListTemplates() (templates []*types.Template, err error) {
	log.Debug("ListTemplates")

	data, status, err := tp.concertoService.Get("/blueprint/templates")
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

// GetTemplate returns a template by its ID
func (tp *TemplateService) GetTemplate(templateID string) (template *types.Template, err error) {
	log.Debug("GetTemplate")

	data, status, err := tp.concertoService.Get(fmt.Sprintf("/blueprint/templates/%s", templateID))
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

// CreateTemplate creates a template
func (tp *TemplateService) CreateTemplate(templateParams *map[string]interface{}) (template *types.Template, err error) {
	log.Debug("CreateTemplate")

	data, status, err := tp.concertoService.Post("/blueprint/templates/", templateParams)
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

// UpdateTemplate updates a template by its ID
func (tp *TemplateService) UpdateTemplate(templateParams *map[string]interface{}, ID string) (template *types.Template, err error) {
	log.Debug("UpdateTemplate")

	data, status, err := tp.concertoService.Put(fmt.Sprintf("/blueprint/templates/%s", ID), templateParams)

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

// CompileTemplate requests compile for a given template by its ID
func (tp *TemplateService) CompileTemplate(payload *map[string]interface{}, templateID string) (template *types.Template, err error) {
	log.Debug("CompileTemplate")

	data, status, err := tp.concertoService.Put(fmt.Sprintf("/blueprint/templates/%s/compile", templateID), payload)
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

// DeleteTemplate deletes a template by its ID
func (tp *TemplateService) DeleteTemplate(templateID string) (err error) {
	log.Debug("DeleteTemplate")

	data, status, err := tp.concertoService.Delete(fmt.Sprintf("/blueprint/templates/%s", templateID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ================ Template Script =================

// ListTemplateScripts returns a list of templateScript by template ID
func (tp *TemplateService) ListTemplateScripts(templateID string, scriptType string) (templateScript []*types.TemplateScript, err error) {
	log.Debug("ListTemplateScripts")

	data, status, err := tp.concertoService.Get(fmt.Sprintf("/blueprint/templates/%s/scripts?type=%s", templateID, scriptType))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// GetTemplateScript returns a templateScript
func (tp *TemplateService) GetTemplateScript(templateID string, templateScriptID string) (templateScript *types.TemplateScript, err error) {
	log.Debug("GetTemplateScript")

	data, status, err := tp.concertoService.Get(fmt.Sprintf("/blueprint/templates/%s/scripts/%s", templateID, templateScriptID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// CreateTemplateScript creates a templateScript
func (tp *TemplateService) CreateTemplateScript(templateScriptParams *map[string]interface{}, templateID string) (templateScript *types.TemplateScript, err error) {
	log.Debug("CreateTemplateScript")

	data, status, err := tp.concertoService.Post(fmt.Sprintf("/blueprint/templates/%s/scripts", templateID), templateScriptParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// UpdateTemplateScript updates a templateScript
func (tp *TemplateService) UpdateTemplateScript(templateScriptParams *map[string]interface{}, templateID string, ID string) (templateScript *types.TemplateScript, err error) {
	log.Debug("UpdateTemplateScript")

	data, status, err := tp.concertoService.Put(fmt.Sprintf("/blueprint/templates/%s/scripts/%s", templateID, ID), templateScriptParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// DeleteTemplateScript deletes a template record
func (tp *TemplateService) DeleteTemplateScript(templateID string, templateScriptID string) (err error) {
	log.Debug("DeleteTemplateScript")

	data, status, err := tp.concertoService.Delete(fmt.Sprintf("/blueprint/templates/%s/scripts/%s", templateID, templateScriptID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ReorderTemplateScript returns a list of templateScript
func (tp *TemplateService) ReorderTemplateScript(templateScriptParams *map[string]interface{}, templateID string) (templateScript []*types.TemplateScript, err error) {
	log.Debug("ReorderTemplateScript")

	data, status, err := tp.concertoService.Put(fmt.Sprintf("/blueprint/templates/%s/scripts/reorder", templateID), templateScriptParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// ================ Template Servers =================

// ListTemplateServers returns a list of templateServers by template ID
func (tp *TemplateService) ListTemplateServers(templateID string) (templateServer []*types.TemplateServer, err error) {
	log.Debug("ListTemplateServers")

	data, status, err := tp.concertoService.Get(fmt.Sprintf("/blueprint/templates/%s/servers", templateID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateServer); err != nil {
		return nil, err
	}

	return templateServer, nil
}
