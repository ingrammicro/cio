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
func (ts *TemplateService) ListTemplates() (templates []*types.Template, err error) {
	log.Debug("ListTemplates")

	data, status, err := ts.concertoService.Get("/blueprint/templates")
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
func (ts *TemplateService) GetTemplate(templateID string) (template *types.Template, err error) {
	log.Debug("GetTemplate")

	data, status, err := ts.concertoService.Get(fmt.Sprintf("/blueprint/templates/%s", templateID))
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
func (ts *TemplateService) CreateTemplate(templateParams *map[string]interface{}) (template *types.Template, err error) {
	log.Debug("CreateTemplate")

	data, status, err := ts.concertoService.Post("/blueprint/templates/", templateParams)
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
func (ts *TemplateService) UpdateTemplate(templateID string, templateParams *map[string]interface{}) (template *types.Template, err error) {
	log.Debug("UpdateTemplate")

	data, status, err := ts.concertoService.Put(fmt.Sprintf("/blueprint/templates/%s", templateID), templateParams)

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
func (ts *TemplateService) CompileTemplate(templateID string, payload *map[string]interface{}) (template *types.Template, err error) {
	log.Debug("CompileTemplate")

	data, status, err := ts.concertoService.Put(fmt.Sprintf("/blueprint/templates/%s/compile", templateID), payload)
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
func (ts *TemplateService) DeleteTemplate(templateID string) (err error) {
	log.Debug("DeleteTemplate")

	data, status, err := ts.concertoService.Delete(fmt.Sprintf("/blueprint/templates/%s", templateID))
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
func (ts *TemplateService) ListTemplateScripts(templateID string, scriptType string) (templateScript []*types.TemplateScript, err error) {
	log.Debug("ListTemplateScripts")

	data, status, err := ts.concertoService.Get(fmt.Sprintf("/blueprint/templates/%s/scripts?type=%s", templateID, scriptType))
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
func (ts *TemplateService) GetTemplateScript(templateID string, templateScriptID string) (templateScript *types.TemplateScript, err error) {
	log.Debug("GetTemplateScript")

	data, status, err := ts.concertoService.Get(fmt.Sprintf("/blueprint/templates/%s/scripts/%s", templateID, templateScriptID))
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
func (ts *TemplateService) CreateTemplateScript(templateID string, templateScriptParams *map[string]interface{}) (templateScript *types.TemplateScript, err error) {
	log.Debug("CreateTemplateScript")

	data, status, err := ts.concertoService.Post(fmt.Sprintf("/blueprint/templates/%s/scripts", templateID), templateScriptParams)
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
func (ts *TemplateService) UpdateTemplateScript(templateID string, templateScriptID string, templateScriptParams *map[string]interface{}) (templateScript *types.TemplateScript, err error) {
	log.Debug("UpdateTemplateScript")

	data, status, err := ts.concertoService.Put(fmt.Sprintf("/blueprint/templates/%s/scripts/%s", templateID, templateScriptID), templateScriptParams)

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
func (ts *TemplateService) DeleteTemplateScript(templateID string, templateScriptID string) (err error) {
	log.Debug("DeleteTemplateScript")

	data, status, err := ts.concertoService.Delete(fmt.Sprintf("/blueprint/templates/%s/scripts/%s", templateID, templateScriptID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ReorderTemplateScript returns a list of templateScript
func (ts *TemplateService) ReorderTemplateScript(templateID string, templateScriptParams *map[string]interface{}) (templateScript []*types.TemplateScript, err error) {
	log.Debug("ReorderTemplateScript")

	data, status, err := ts.concertoService.Put(fmt.Sprintf("/blueprint/templates/%s/scripts/reorder", templateID), templateScriptParams)
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
func (ts *TemplateService) ListTemplateServers(templateID string) (templateServer []*types.TemplateServer, err error) {
	log.Debug("ListTemplateServers")

	data, status, err := ts.concertoService.Get(fmt.Sprintf("/blueprint/templates/%s/servers", templateID))
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
