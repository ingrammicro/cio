// Copyright (c) 2017-2021 Ingram Micro Inc.

package blueprint

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// ListScriptsMocked test mocked function
func ListScriptsMocked(t *testing.T, scriptsIn []*types.Script) []*types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", APIPathBlueprintScripts).Return(dIn, 200, nil)
	scriptsOut, err := ds.ListScripts()
	assert.Nil(err, "Error getting script list")
	assert.Equal(scriptsIn, scriptsOut, "ListScripts returned different scripts")

	return scriptsOut
}

// ListScriptsFailErrMocked test mocked function
func ListScriptsFailErrMocked(t *testing.T, scriptsIn []*types.Script) []*types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", APIPathBlueprintScripts).Return(dIn, 200, fmt.Errorf("mocked error"))
	scriptsOut, err := ds.ListScripts()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scriptsOut
}

// ListScriptsFailStatusMocked test mocked function
func ListScriptsFailStatusMocked(t *testing.T, scriptsIn []*types.Script) []*types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", APIPathBlueprintScripts).Return(dIn, 499, nil)
	scriptsOut, err := ds.ListScripts()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scriptsOut
}

// ListScriptsFailJSONMocked test mocked function
func ListScriptsFailJSONMocked(t *testing.T, scriptsIn []*types.Script) []*types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathBlueprintScripts).Return(dIn, 200, nil)
	scriptsOut, err := ds.ListScripts()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptsOut
}

// GetScriptMocked test mocked function
func GetScriptMocked(t *testing.T, script *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(script)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintScript, script.ID)).Return(dIn, 200, nil)
	scriptOut, err := ds.GetScript(script.ID)
	assert.Nil(err, "Error getting script")
	assert.Equal(*script, *scriptOut, "GetScript returned different scripts")

	return scriptOut
}

// GetScriptFailErrMocked test mocked function
func GetScriptFailErrMocked(t *testing.T, script *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(script)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintScript, script.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	scriptOut, err := ds.GetScript(script.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scriptOut
}

// GetScriptFailStatusMocked test mocked function
func GetScriptFailStatusMocked(t *testing.T, script *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(script)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintScript, script.ID)).Return(dIn, 499, nil)
	scriptOut, err := ds.GetScript(script.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scriptOut
}

// GetScriptFailJSONMocked test mocked function
func GetScriptFailJSONMocked(t *testing.T, script *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintScript, script.ID)).Return(dIn, 200, nil)
	scriptOut, err := ds.GetScript(script.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptOut
}

// CreateScriptMocked test mocked function
func CreateScriptMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", APIPathBlueprintScripts, mapIn).Return(dOut, 200, nil)
	scriptOut, err := ds.CreateScript(mapIn)
	assert.Nil(err, "Error creating script list")
	assert.Equal(*scriptIn, *scriptOut, "CreateScript returned different scripts")

	return scriptOut
}

// CreateScriptFailErrMocked test mocked function
func CreateScriptFailErrMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", APIPathBlueprintScripts, mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	scriptOut, err := ds.CreateScript(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scriptOut
}

// CreateScriptFailStatusMocked test mocked function
func CreateScriptFailStatusMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", APIPathBlueprintScripts, mapIn).Return(dOut, 499, nil)
	scriptOut, err := ds.CreateScript(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scriptOut
}

// CreateScriptFailJSONMocked test mocked function
func CreateScriptFailJSONMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", APIPathBlueprintScripts, mapIn).Return(dIn, 200, nil)
	scriptOut, err := ds.CreateScript(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptOut
}

// UpdateScriptMocked test mocked function
func UpdateScriptMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathBlueprintScript, scriptIn.ID), mapIn).Return(dOut, 200, nil)
	scriptOut, err := ds.UpdateScript(scriptIn.ID, mapIn)
	assert.Nil(err, "Error updating script list")
	assert.Equal(*scriptIn, *scriptOut, "UpdateScript returned different scripts")

	return scriptOut
}

// UpdateScriptFailErrMocked test mocked function
func UpdateScriptFailErrMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathBlueprintScript, scriptIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	scriptOut, err := ds.UpdateScript(scriptIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return scriptOut
}

// UpdateScriptFailStatusMocked test mocked function
func UpdateScriptFailStatusMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathBlueprintScript, scriptIn.ID), mapIn).Return(dOut, 499, nil)
	scriptOut, err := ds.UpdateScript(scriptIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return scriptOut
}

// UpdateScriptFailJSONMocked test mocked function
func UpdateScriptFailJSONMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathBlueprintScript, scriptIn.ID), mapIn).Return(dIn, 200, nil)
	scriptOut, err := ds.UpdateScript(scriptIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptOut
}

// DeleteScriptMocked test mocked function
func DeleteScriptMocked(t *testing.T, scriptIn *types.Script) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathBlueprintScript, scriptIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteScript(scriptIn.ID)
	assert.Nil(err, "Error deleting script")

}

// DeleteScriptFailErrMocked test mocked function
func DeleteScriptFailErrMocked(t *testing.T, scriptIn *types.Script) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathBlueprintScript, scriptIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteScript(scriptIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteScriptFailStatusMocked test mocked function
func DeleteScriptFailStatusMocked(t *testing.T, scriptIn *types.Script) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathBlueprintScript, scriptIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteScript(scriptIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// AddScriptAttachmentMocked test mocked function
func AddScriptAttachmentMocked(t *testing.T, attachmentIn *types.Attachment, scriptID string) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID), mapIn).Return(dIn, 200, nil)
	attachmentOut, err := ds.AddScriptAttachment(scriptID, mapIn)
	assert.Nil(err, "Error getting template list")
	assert.Equal(attachmentIn, attachmentOut, "AddScriptAttachment returned different attachments")

	return attachmentOut
}

// AddScriptAttachmentFailErrMocked test mocked function
func AddScriptAttachmentFailErrMocked(t *testing.T, attachmentIn *types.Attachment, scriptID string) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID), mapIn).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	attachmentOut, err := ds.AddScriptAttachment(scriptID, mapIn)
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return attachmentOut
}

// AddScriptAttachmentFailStatusMocked test mocked function
func AddScriptAttachmentFailStatusMocked(
	t *testing.T,
	attachmentIn *types.Attachment,
	scriptID string,
) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID), mapIn).Return(dIn, 499, nil)
	attachmentOut, err := ds.AddScriptAttachment(scriptID, mapIn)
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return attachmentOut
}

// AddScriptAttachmentFailJSONMocked test mocked function
func AddScriptAttachmentFailJSONMocked(
	t *testing.T,
	attachmentIn *types.Attachment,
	scriptID string,
) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID), mapIn).Return(dIn, 200, nil)
	attachmentOut, err := ds.AddScriptAttachment(scriptID, mapIn)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return attachmentOut
}

// UploadScriptAttachmentMocked test mocked function
func UploadScriptAttachmentMocked(t *testing.T, attachmentIn *types.Attachment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := attachmentIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 200, nil)
	err = ds.UploadScriptAttachment(sourceFilePath, targetURL)
	assert.Nil(err, "Error uploading attachment file")
}

// UploadScriptAttachmentFailStatusMocked test mocked function
func UploadScriptAttachmentFailStatusMocked(t *testing.T, attachmentIn *types.Attachment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := attachmentIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 403, nil)
	err = ds.UploadScriptAttachment(sourceFilePath, targetURL)
	assert.NotNil(err, "We are expecting an error")
}

// UploadScriptAttachmentFailErrMocked test mocked function
func UploadScriptAttachmentFailErrMocked(t *testing.T, attachmentIn *types.Attachment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := attachmentIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 403, fmt.Errorf("mocked error"))
	err = ds.UploadScriptAttachment(sourceFilePath, targetURL)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// UploadedScriptAttachmentMocked test mocked function
func UploadedScriptAttachmentMocked(t *testing.T, attachmentIn *types.Attachment) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathBlueprintScriptAttachmentUploaded, attachmentIn.ID), mapIn).Return(dIn, 200, nil)
	attachmentOut, err := ds.UploadedScriptAttachment(attachmentIn.ID, mapIn)
	assert.Nil(err, "Error setting uploaded status to attachmentIn")
	assert.Equal(*attachmentIn, *attachmentOut, "UploadedScriptAttachment returned different attachments")

	return attachmentOut
}

// UploadedScriptAttachmentFailErrMocked test mocked function
func UploadedScriptAttachmentFailErrMocked(t *testing.T, attachmentIn *types.Attachment) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathBlueprintScriptAttachmentUploaded, attachmentIn.ID), mapIn).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	attachmentOut, err := ds.UploadedScriptAttachment(attachmentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return attachmentOut
}

// UploadedScriptAttachmentFailStatusMocked test mocked function
func UploadedScriptAttachmentFailStatusMocked(t *testing.T, attachmentIn *types.Attachment) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dIn, err := json.Marshal(attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathBlueprintScriptAttachmentUploaded, attachmentIn.ID), mapIn).Return(dIn, 499, nil)
	attachmentOut, err := ds.UploadedScriptAttachment(attachmentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return attachmentOut
}

// UploadedScriptAttachmentFailJSONMocked test mocked function
func UploadedScriptAttachmentFailJSONMocked(t *testing.T, attachmentIn *types.Attachment) *types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*attachmentIn)
	assert.Nil(err, "Script test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathBlueprintScriptAttachmentUploaded, attachmentIn.ID), mapIn).Return(dIn, 200, nil)
	attachmentOut, err := ds.UploadedScriptAttachment(attachmentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(attachmentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return attachmentOut
}

// ListScriptAttachmentsMocked test mocked function
func ListScriptAttachmentsMocked(t *testing.T, attachmentsIn []*types.Attachment, scriptID string) []*types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID)).Return(dIn, 200, nil)
	attachmentsOut, err := ds.ListScriptAttachments(scriptID)
	assert.Nil(err, "Error getting script attachments list")
	assert.Equal(attachmentsIn, attachmentsOut, "ListScriptAttachments returned different attachments")

	return attachmentsOut
}

// ListScriptAttachmentsFailErrMocked test mocked function
func ListScriptAttachmentsFailErrMocked(
	t *testing.T,
	attachmentsIn []*types.Attachment,
	scriptID string,
) []*types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	attachmentsOut, err := ds.ListScriptAttachments(scriptID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(attachmentsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return attachmentsOut
}

// ListScriptAttachmentsFailStatusMocked test mocked function
func ListScriptAttachmentsFailStatusMocked(
	t *testing.T,
	attachmentsIn []*types.Attachment,
	scriptID string,
) []*types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(attachmentsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID)).Return(dIn, 499, nil)
	attachmentsOut, err := ds.ListScriptAttachments(scriptID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(attachmentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return attachmentsOut
}

// ListScriptAttachmentsFailJSONMocked test mocked function
func ListScriptAttachmentsFailJSONMocked(
	t *testing.T,
	attachmentsIn []*types.Attachment,
	scriptID string,
) []*types.Attachment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintScriptAttachments, scriptID)).Return(dIn, 200, nil)
	attachmentsOut, err := ds.ListScriptAttachments(scriptID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(attachmentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return attachmentsOut
}
