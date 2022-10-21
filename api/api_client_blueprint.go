// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// GetAttachment returns a attachment by its ID
func (imco *ClientAPI) GetAttachment(ctx context.Context, attachmentID string,
) (attachment *types.Attachment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathBlueprintAttachment, attachmentID), true, &attachment)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

// DeleteAttachment deletes a attachment by its ID
func (imco *ClientAPI) DeleteAttachment(ctx context.Context, attachmentID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathBlueprintAttachment, attachmentID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListCookbookVersions returns the list of cookbook versions as an array of CookbookVersion
func (imco *ClientAPI) ListCookbookVersions(ctx context.Context,
) (cookbookVersions []*types.CookbookVersion, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathBlueprintCookbookVersions, true, &cookbookVersions)
	if err != nil {
		return nil, err
	}
	return cookbookVersions, nil
}

// GetCookbookVersion returns a cookbook version by its ID
func (imco *ClientAPI) GetCookbookVersion(ctx context.Context, cookbookVersionID string,
) (cookbookVersion *types.CookbookVersion, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathBlueprintCookbookVersion, cookbookVersionID), true, &cookbookVersion)
	if err != nil {
		return nil, err
	}
	return cookbookVersion, nil
}

// CreateCookbookVersion creates a new cookbook version
func (imco *ClientAPI) CreateCookbookVersion(ctx context.Context, cookbookVersionParams *map[string]interface{},
) (cookbookVersion *types.CookbookVersion, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, PathBlueprintCookbookVersions, cookbookVersionParams, true, &cookbookVersion)
	if err != nil {
		return nil, err
	}
	return cookbookVersion, nil
}

// ProcessCookbookVersion process a cookbook version by its ID
func (imco *ClientAPI) ProcessCookbookVersion(ctx context.Context, cookbookVersionID string,
	cookbookVersionParams *map[string]interface{},
) (cookbookVersion *types.CookbookVersion, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(PathBlueprintCookbookVersionProcess, cookbookVersionID),
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
func (imco *ClientAPI) DeleteCookbookVersion(ctx context.Context, cookbookVersionID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathBlueprintCookbookVersion, cookbookVersionID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListScripts returns the list of scripts as an array of Scripts
func (imco *ClientAPI) ListScripts(ctx context.Context) (scripts []*types.Script, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathBlueprintScripts, true, &scripts)
	if err != nil {
		return nil, err
	}
	return scripts, nil
}

// GetScript returns a script by its ID
func (imco *ClientAPI) GetScript(ctx context.Context, scriptID string) (script *types.Script, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathBlueprintScript, scriptID), true, &script)
	if err != nil {
		return nil, err
	}
	return script, nil
}

// CreateScript creates a script
func (imco *ClientAPI) CreateScript(ctx context.Context, scriptParams *map[string]interface{},
) (script *types.Script, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, PathBlueprintScripts, scriptParams, true, &script)
	if err != nil {
		return nil, err
	}
	return script, nil
}

// UpdateScript updates a script by its ID
func (imco *ClientAPI) UpdateScript(ctx context.Context, scriptID string, scriptParams *map[string]interface{},
) (script *types.Script, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathBlueprintScript, scriptID), scriptParams, true, &script)
	if err != nil {
		return nil, err
	}
	return script, nil
}

// DeleteScript deletes a script by its ID
func (imco *ClientAPI) DeleteScript(ctx context.Context, scriptID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathBlueprintScript, scriptID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// AddScriptAttachment adds an attachment to script by its ID
func (imco *ClientAPI) AddScriptAttachment(ctx context.Context, scriptID string, attachmentIn *map[string]interface{},
) (attachment *types.Attachment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, fmt.Sprintf(PathBlueprintScriptAttachments, scriptID), attachmentIn, true, &attachment)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

// UploadedScriptAttachment sets "uploaded" status to the attachment by its ID
func (imco *ClientAPI) UploadedScriptAttachment(ctx context.Context, attachmentID string,
	attachmentParams *map[string]interface{},
) (attachment *types.Attachment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBlueprintAttachmentUploaded, attachmentID),
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
func (imco *ClientAPI) ListScriptAttachments(ctx context.Context, scriptID string,
) (attachments []*types.Attachment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathBlueprintScriptAttachments, scriptID), true, &attachments)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

// ListTemplates returns the list of templates as an array of Template
func (imco *ClientAPI) ListTemplates(ctx context.Context) (templates []*types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathBlueprintTemplates, true, &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

// GetTemplate returns a template by its ID
func (imco *ClientAPI) GetTemplate(ctx context.Context, templateID string) (template *types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathBlueprintTemplate, templateID), true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CreateTemplate creates a template
func (imco *ClientAPI) CreateTemplate(ctx context.Context, templateParams *map[string]interface{},
) (template *types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, PathBlueprintTemplates, templateParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// UpdateTemplate updates a template by its ID
func (imco *ClientAPI) UpdateTemplate(ctx context.Context, templateID string, templateParams *map[string]interface{},
) (template *types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathBlueprintTemplate, templateID), templateParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CompileTemplate requests compile for a given template by its ID
func (imco *ClientAPI) CompileTemplate(ctx context.Context, templateID string, payload *map[string]interface{},
) (template *types.Template, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathBlueprintTemplateCompile, templateID), payload, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// DeleteTemplate deletes a template by its ID
func (imco *ClientAPI) DeleteTemplate(ctx context.Context, templateID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathBlueprintTemplate, templateID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListTemplateScripts returns a list of templateScript by template ID
func (imco *ClientAPI) ListTemplateScripts(ctx context.Context, templateID string, scriptType string,
) (templateScripts []*types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(PathBlueprintTemplateScriptsType, templateID, scriptType),
		true,
		&templateScripts,
	)
	if err != nil {
		return nil, err
	}
	return templateScripts, nil
}

// GetTemplateScript returns a templateScript
func (imco *ClientAPI) GetTemplateScript(ctx context.Context, templateID string, templateScriptID string,
) (templateScript *types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(PathBlueprintTemplateScript, templateID, templateScriptID),
		true,
		&templateScript,
	)
	if err != nil {
		return nil, err
	}
	return templateScript, nil
}

// CreateTemplateScript creates a templateScript
func (imco *ClientAPI) CreateTemplateScript(ctx context.Context, templateID string,
	templateScriptParams *map[string]interface{},
) (templateScript *types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(PathBlueprintTemplateScripts, templateID),
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
func (imco *ClientAPI) UpdateTemplateScript(ctx context.Context, templateID string, templateScriptID string,
	templateScriptParams *map[string]interface{},
) (templateScript *types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBlueprintTemplateScript, templateID, templateScriptID),
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
func (imco *ClientAPI) DeleteTemplateScript(ctx context.Context, templateID string, templateScriptID string,
) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathBlueprintTemplateScript, templateID, templateScriptID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// ReorderTemplateScript returns a list of templateScript
func (imco *ClientAPI) ReorderTemplateScript(ctx context.Context, templateID string,
	templateScriptParams *map[string]interface{},
) (templateScripts []*types.TemplateScript, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBlueprintTemplateScriptsReorder, templateID),
		templateScriptParams,
		true,
		&templateScripts,
	)
	if err != nil {
		return nil, err
	}
	return templateScripts, nil
}

// ListTemplateServers returns a list of templateServers by template ID
func (imco *ClientAPI) ListTemplateServers(ctx context.Context, templateID string,
) (templateServers []*types.TemplateServer, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathBlueprintTemplateServers, templateID), true, &templateServers)
	if err != nil {
		return nil, err
	}
	return templateServers, nil
}
