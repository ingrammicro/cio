// Copyright (c) 2017-2021 Ingram Micro Inc.

package blueprint

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathBlueprintScripts = "/blueprint/scripts"
const APIPathBlueprintScript = "/blueprint/scripts/%s"
const APIPathBlueprintScriptAttachments = "/blueprint/scripts/%s/attachments"
const APIPathBlueprintScriptAttachmentUploaded = "/blueprint/attachments/%s/uploaded"

// ScriptService manages script operations
type ScriptService struct {
	concertoService utils.ConcertoService
}

// NewScriptService returns a Concerto script service
func NewScriptService(concertoService utils.ConcertoService) (*ScriptService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ScriptService{
		concertoService: concertoService,
	}, nil
}

// ListScripts returns the list of scripts as an array of Scripts
func (sc *ScriptService) ListScripts() (scripts []*types.Script, err error) {
	log.Debug("ListScripts")

	data, status, err := sc.concertoService.Get(APIPathBlueprintScripts)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

// GetScript returns a script by its ID
func (sc *ScriptService) GetScript(scriptID string) (script *types.Script, err error) {
	log.Debug("GetScript")

	data, status, err := sc.concertoService.Get(fmt.Sprintf(APIPathBlueprintScript, scriptID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// CreateScript creates a script
func (sc *ScriptService) CreateScript(scriptParams *map[string]interface{}) (script *types.Script, err error) {
	log.Debug("CreateScript")

	data, status, err := sc.concertoService.Post(APIPathBlueprintScripts, scriptParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// UpdateScript updates a script by its ID
func (sc *ScriptService) UpdateScript(
	scriptID string,
	scriptParams *map[string]interface{},
) (script *types.Script, err error) {
	log.Debug("UpdateScript")

	data, status, err := sc.concertoService.Put(fmt.Sprintf(APIPathBlueprintScript, scriptID), scriptParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// DeleteScript deletes a script by its ID
func (sc *ScriptService) DeleteScript(scriptID string) (err error) {
	log.Debug("DeleteScript")

	data, status, err := sc.concertoService.Delete(fmt.Sprintf(APIPathBlueprintScript, scriptID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// AddScriptAttachment adds an attachment to script by its ID
func (sc *ScriptService) AddScriptAttachment(
	scriptID string,
	attachmentIn *map[string]interface{},
) (script *types.Attachment, err error) {
	log.Debug("AddScriptAttachment")

	data, status, err := sc.concertoService.Post(fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID), attachmentIn)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// UploadScriptAttachment uploads an attachment file
func (sc *ScriptService) UploadScriptAttachment(sourceFilePath string, targetURL string) error {
	log.Debug("UploadScriptAttachment")

	data, status, err := sc.concertoService.PutFile(sourceFilePath, targetURL)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// UploadedScriptAttachment sets "uploaded" status to the attachment by its ID
func (sc *ScriptService) UploadedScriptAttachment(
	attachmentID string,
	attachmentParams *map[string]interface{},
) (attachment *types.Attachment, err error) {
	log.Debug("UploadedScriptAttachment")

	data, status, err := sc.concertoService.Put(
		fmt.Sprintf(APIPathBlueprintScriptAttachmentUploaded, attachmentID),
		attachmentParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &attachment); err != nil {
		return nil, err
	}

	return attachment, nil
}

// ListScriptAttachments returns the list of Attachments for a given script ID
func (sc *ScriptService) ListScriptAttachments(scriptID string) (attachments []*types.Attachment, err error) {
	log.Debug("ListScriptAttachments")

	data, status, err := sc.concertoService.Get(fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &attachments); err != nil {
		return nil, err
	}

	return attachments, nil
}
