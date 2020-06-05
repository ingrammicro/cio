// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// GetAttachment returns a attachment by its ID
func (imco *IMCOClient) GetAttachment(attachmentID string) (attachment *types.Attachment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathBlueprintAttachment, attachmentID), true, &attachment)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

// DeleteAttachment deletes a attachment by its ID
func (imco *IMCOClient) DeleteAttachment(attachmentID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathBlueprintAttachment, attachmentID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListCookbookVersions returns the list of cookbook versions as an array of CookbookVersion
func (imco *IMCOClient) ListCookbookVersions() (cookbookVersions []*types.CookbookVersion, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathBlueprintCookbookVersions, true, &cookbookVersions)
	if err != nil {
		return nil, err
	}
	return cookbookVersions, nil
}

// GetCookbookVersion returns a cookbook version by its ID
func (imco *IMCOClient) GetCookbookVersion(cookbookVersionID string,
) (cookbookVersion *types.CookbookVersion, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathBlueprintCookbookVersion, cookbookVersionID), true, &cookbookVersion)
	if err != nil {
		return nil, err
	}
	return cookbookVersion, nil
}

// CreateCookbookVersion creates a new cookbook version
func (imco *IMCOClient) CreateCookbookVersion(cookbookVersionParams *map[string]interface{},
) (cookbookVersion *types.CookbookVersion, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathBlueprintCookbookVersions, cookbookVersionParams, true, &cookbookVersion)
	if err != nil {
		return nil, err
	}
	return cookbookVersion, nil
}

// ProcessCookbookVersion process a cookbook version by its ID
func (imco *IMCOClient) ProcessCookbookVersion(cookbookVersionID string, cookbookVersionParams *map[string]interface{},
) (cookbookVersion *types.CookbookVersion, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
		fmt.Sprintf(pathBlueprintCookbookVersionProcess, cookbookVersionID),
		cookbookVersionParams,
		true,
		&cookbookVersion,
	)
	if err != nil {
		return nil, err
	}
	return cookbookVersion, nil
}

// DeleteCookbookVersion deletes a cookbook version by its ID
func (imco *IMCOClient) DeleteCookbookVersion(cookbookVersionID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathBlueprintCookbookVersion, cookbookVersionID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListScripts returns the list of scripts as an array of Scripts
func (imco *IMCOClient) ListScripts() (scripts []*types.Script, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathBlueprintScripts, true, &scripts)
	if err != nil {
		return nil, err
	}
	return scripts, nil
}

// GetScript returns a script by its ID
func (imco *IMCOClient) GetScript(scriptID string) (script *types.Script, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathBlueprintScript, scriptID), true, &script)
	if err != nil {
		return nil, err
	}
	return script, nil
}

// CreateScript creates a script
func (imco *IMCOClient) CreateScript(scriptParams *map[string]interface{}) (script *types.Script, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathBlueprintScripts, scriptParams, true, &script)
	if err != nil {
		return nil, err
	}
	return script, nil
}

// UpdateScript updates a script by its ID
func (imco *IMCOClient) UpdateScript(scriptID string, scriptParams *map[string]interface{},
) (script *types.Script, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathBlueprintScript, scriptID), scriptParams, true, &script)
	if err != nil {
		return nil, err
	}
	return script, nil
}

// DeleteScript deletes a script by its ID
func (imco *IMCOClient) DeleteScript(scriptID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathBlueprintScript, scriptID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// AddScriptAttachment adds an attachment to script by its ID
func (imco *IMCOClient) AddScriptAttachment(scriptID string, attachmentIn *map[string]interface{},
) (script *types.Attachment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathBlueprintScriptAttachments, scriptID), attachmentIn, true, &script)
	if err != nil {
		return nil, err
	}
	return script, nil
}

// UploadedScriptAttachment sets "uploaded" status to the attachment by its ID
func (imco *IMCOClient) UploadedScriptAttachment(attachmentID string, attachmentParams *map[string]interface{},
) (attachment *types.Attachment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
		fmt.Sprintf(pathBlueprintAttachmentUploaded, attachmentID),
		attachmentParams,
		true,
		&attachment,
	)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

// ListScriptAttachments returns the list of Attachments for a given script ID
func (imco *IMCOClient) ListScriptAttachments(scriptID string) (attachments []*types.Attachment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathBlueprintScriptAttachments, scriptID), true, &attachments)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

// ListTemplates returns the list of templates as an array of Template
func (imco *IMCOClient) ListTemplates() (templates []*types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathBlueprintTemplates, true, &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

// GetTemplate returns a template by its ID
func (imco *IMCOClient) GetTemplate(templateID string) (template *types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathBlueprintTemplate, templateID), true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CreateTemplate creates a template
func (imco *IMCOClient) CreateTemplate(templateParams *map[string]interface{}) (template *types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathBlueprintTemplates, templateParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// UpdateTemplate updates a template by its ID
func (imco *IMCOClient) UpdateTemplate(templateID string, templateParams *map[string]interface{},
) (template *types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathBlueprintTemplate, templateID), templateParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CompileTemplate requests compile for a given template by its ID
func (imco *IMCOClient) CompileTemplate(templateID string, payload *map[string]interface{},
) (template *types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathBlueprintTemplateCompile, templateID), payload, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// DeleteTemplate deletes a template by its ID
func (imco *IMCOClient) DeleteTemplate(templateID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathBlueprintTemplate, templateID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListTemplateScripts returns a list of templateScript by template ID
func (imco *IMCOClient) ListTemplateScripts(templateID string, scriptType string,
) (templateScript []*types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathBlueprintTemplateScriptsType, templateID, scriptType),
		true,
		&templateScript,
	)
	if err != nil {
		return nil, err
	}
	return templateScript, nil
}

// GetTemplateScript returns a templateScript
func (imco *IMCOClient) GetTemplateScript(templateID string, templateScriptID string,
) (templateScript *types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathBlueprintTemplateScript, templateID, templateScriptID),
		true,
		&templateScript,
	)
	if err != nil {
		return nil, err
	}
	return templateScript, nil
}

// CreateTemplateScript creates a templateScript
func (imco *IMCOClient) CreateTemplateScript(templateID string, templateScriptParams *map[string]interface{},
) (templateScript *types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
		fmt.Sprintf(pathBlueprintTemplateScripts, templateID),
		templateScriptParams,
		true,
		&templateScript,
	)
	if err != nil {
		return nil, err
	}
	return templateScript, nil
}

// UpdateTemplateScript updates a templateScript
func (imco *IMCOClient) UpdateTemplateScript(templateID string, templateScriptID string,
	templateScriptParams *map[string]interface{},
) (templateScript *types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
		fmt.Sprintf(pathBlueprintTemplateScript, templateID, templateScriptID),
		templateScriptParams,
		true,
		&templateScript,
	)
	if err != nil {
		return nil, err
	}
	return templateScript, nil
}

// DeleteTemplateScript deletes a template record
func (imco *IMCOClient) DeleteTemplateScript(templateID string, templateScriptID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathBlueprintTemplateScript, templateID, templateScriptID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ReorderTemplateScript returns a list of templateScript
func (imco *IMCOClient) ReorderTemplateScript(templateID string, templateScriptParams *map[string]interface{},
) (templateScript []*types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
		fmt.Sprintf(pathBlueprintTemplateScriptsReorder, templateID),
		templateScriptParams,
		true,
		&templateScript,
	)
	if err != nil {
		return nil, err
	}
	return templateScript, nil
}

// ListTemplateServers returns a list of templateServers by template ID
func (imco *IMCOClient) ListTemplateServers(templateID string) (templateServer []*types.TemplateServer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathBlueprintTemplateServers, templateID), true, &templateID)
	if err != nil {
		return nil, err
	}
	return templateServer, nil
}
