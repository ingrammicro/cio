// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// CreateTemporaryArchiveMocked test mocked function
func CreateTemporaryArchiveMocked(t *testing.T, temporaryArchiveIn *types.TemporaryArchive) *types.TemporaryArchive {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveIn)
	assert.Nil(err, "TemporaryArchive test data corrupted")

	// to json
	dOut, err := json.Marshal(temporaryArchiveIn)
	assert.Nil(err, "TemporaryArchive test data corrupted")

	// call service
	cs.On("Post", APIPathPluginsToscaTemporaryArchives, mapIn).Return(dOut, 200, nil)
	temporaryArchiveOut, err := ds.CreateTemporaryArchive(mapIn)

	assert.Nil(err, "Error creating temporary archive")
	assert.Equal(temporaryArchiveIn, temporaryArchiveOut, "CreateTemporaryArchive returned different temporary archive")

	return temporaryArchiveOut
}

// CreateTemporaryArchiveFailErrMocked test mocked function
func CreateTemporaryArchiveFailErrMocked(
	t *testing.T,
	temporaryArchiveIn *types.TemporaryArchive,
) *types.TemporaryArchive {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveIn)
	assert.Nil(err, "TemporaryArchive test data corrupted")

	// to json
	dOut, err := json.Marshal(temporaryArchiveIn)
	assert.Nil(err, "TemporaryArchive test data corrupted")

	// call service
	cs.On("Post", APIPathPluginsToscaTemporaryArchives, mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	temporaryArchiveOut, err := ds.CreateTemporaryArchive(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(temporaryArchiveOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return temporaryArchiveOut
}

// CreateTemporaryArchiveFailStatusMocked test mocked function
func CreateTemporaryArchiveFailStatusMocked(
	t *testing.T,
	temporaryArchiveIn *types.TemporaryArchive,
) *types.TemporaryArchive {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveIn)
	assert.Nil(err, "TemporaryArchive test data corrupted")

	// to json
	dOut, err := json.Marshal(temporaryArchiveIn)
	assert.Nil(err, "TemporaryArchive test data corrupted")

	// call service
	cs.On("Post", APIPathPluginsToscaTemporaryArchives, mapIn).Return(dOut, 499, nil)
	temporaryArchiveOut, err := ds.CreateTemporaryArchive(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(temporaryArchiveOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return temporaryArchiveOut
}

// CreateTemporaryArchiveFailJSONMocked test mocked function
func CreateTemporaryArchiveFailJSONMocked(
	t *testing.T,
	temporaryArchiveIn *types.TemporaryArchive,
) *types.TemporaryArchive {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveIn)
	assert.Nil(err, "TemporaryArchive test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", APIPathPluginsToscaTemporaryArchives, mapIn).Return(dIn, 200, nil)
	temporaryArchiveOut, err := ds.CreateTemporaryArchive(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(temporaryArchiveOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return temporaryArchiveOut
}

// UploadTemporaryArchiveMocked test mocked function
func UploadTemporaryArchiveMocked(t *testing.T, temporaryArchiveIn *types.TemporaryArchive) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := temporaryArchiveIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 200, nil)
	err = ds.UploadTemporaryArchive(sourceFilePath, targetURL)
	assert.Nil(err, "Error uploading temporary archive file")
}

// UploadTemporaryArchiveFailStatusMocked test mocked function
func UploadTemporaryArchiveFailStatusMocked(t *testing.T, temporaryArchiveIn *types.TemporaryArchive) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := temporaryArchiveIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 403, nil)
	err = ds.UploadTemporaryArchive(sourceFilePath, targetURL)
	assert.NotNil(err, "We are expecting an error")
}

// UploadTemporaryArchiveFailErrMocked test mocked function
func UploadTemporaryArchiveFailErrMocked(t *testing.T, temporaryArchiveIn *types.TemporaryArchive) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := temporaryArchiveIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 403, fmt.Errorf("mocked error"))
	err = ds.UploadTemporaryArchive(sourceFilePath, targetURL)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// CreateTemporaryArchiveImportMocked test mocked function
func CreateTemporaryArchiveImportMocked(
	t *testing.T,
	temporaryArchiveID string,
	temporaryArchiveImportIn *types.TemporaryArchiveImport,
) *types.TemporaryArchiveImport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveImportIn)

	// to json
	dOut, err := json.Marshal(temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveID), mapIn).
		Return(dOut, 200, nil)
	temporaryArchiveImportOut, err := ds.CreateTemporaryArchiveImport(temporaryArchiveID, mapIn)

	assert.Nil(err, "Error creating temporary archive import")
	assert.Equal(
		temporaryArchiveImportIn,
		temporaryArchiveImportOut,
		"CreateTemporaryArchiveImport returned different temporary archive import",
	)

	return temporaryArchiveImportOut
}

// CreateTemporaryArchiveImportFailErrMocked test mocked function
func CreateTemporaryArchiveImportFailErrMocked(
	t *testing.T,
	temporaryArchiveID string,
	temporaryArchiveImportIn *types.TemporaryArchiveImport,
) *types.TemporaryArchiveImport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// to json
	dOut, err := json.Marshal(temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	temporaryArchiveImportOut, err := ds.CreateTemporaryArchiveImport(temporaryArchiveID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(temporaryArchiveImportOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return temporaryArchiveImportOut
}

// CreateTemporaryArchiveImportFailStatusMocked test mocked function
func CreateTemporaryArchiveImportFailStatusMocked(
	t *testing.T,
	temporaryArchiveID string,
	temporaryArchiveImportIn *types.TemporaryArchiveImport,
) *types.TemporaryArchiveImport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// to json
	dOut, err := json.Marshal(temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveID), mapIn).
		Return(dOut, 499, nil)
	temporaryArchiveImportOut, err := ds.CreateTemporaryArchiveImport(temporaryArchiveID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(temporaryArchiveImportOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return temporaryArchiveImportOut
}

// CreateTemporaryArchiveImportFailJSONMocked test mocked function
func CreateTemporaryArchiveImportFailJSONMocked(
	t *testing.T,
	temporaryArchiveID string,
	temporaryArchiveImportIn *types.TemporaryArchiveImport,
) *types.TemporaryArchiveImport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveID), mapIn).
		Return(dIn, 200, nil)
	temporaryArchiveImportOut, err := ds.CreateTemporaryArchiveImport(temporaryArchiveID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(temporaryArchiveImportOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return temporaryArchiveImportOut
}

// GetTemporaryArchiveImportMocked test mocked function
func GetTemporaryArchiveImportMocked(
	t *testing.T,
	temporaryArchiveImportIn *types.TemporaryArchiveImport,
) *types.TemporaryArchiveImport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// to json
	dIn, err := json.Marshal(temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveImportIn.ID)).
		Return(dIn, 200, nil)
	temporaryArchiveImportOut, err := ds.GetTemporaryArchiveImport(temporaryArchiveImportIn.ID)

	assert.Nil(err, "Error getting temporary archive import")
	assert.Equal(
		*temporaryArchiveImportIn,
		*temporaryArchiveImportOut,
		"GetTemporaryArchiveImport returned different temporary archive import",
	)

	return temporaryArchiveImportOut
}

// GetTemporaryArchiveImportFailErrMocked test mocked function
func GetTemporaryArchiveImportFailErrMocked(
	t *testing.T,
	temporaryArchiveImportIn *types.TemporaryArchiveImport,
) *types.TemporaryArchiveImport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// to json
	dIn, err := json.Marshal(temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveImportIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	temporaryArchiveImportOut, err := ds.GetTemporaryArchiveImport(temporaryArchiveImportIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(temporaryArchiveImportOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return temporaryArchiveImportOut
}

// GetTemporaryArchiveImportFailStatusMocked test mocked function
func GetTemporaryArchiveImportFailStatusMocked(
	t *testing.T,
	temporaryArchiveImportIn *types.TemporaryArchiveImport,
) *types.TemporaryArchiveImport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// to json
	dIn, err := json.Marshal(temporaryArchiveImportIn)
	assert.Nil(err, "TemporaryArchiveImport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveImportIn.ID)).
		Return(dIn, 499, nil)
	temporaryArchiveImportOut, err := ds.GetTemporaryArchiveImport(temporaryArchiveImportIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(temporaryArchiveImportOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return temporaryArchiveImportOut
}

// GetTemporaryArchiveImportFailJSONMocked test mocked function
func GetTemporaryArchiveImportFailJSONMocked(
	t *testing.T,
	temporaryArchiveImportIn *types.TemporaryArchiveImport,
) *types.TemporaryArchiveImport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveImportIn.ID)).
		Return(dIn, 200, nil)
	temporaryArchiveImportOut, err := ds.GetTemporaryArchiveImport(temporaryArchiveImportIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(temporaryArchiveImportOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return temporaryArchiveImportOut
}

// CreateTemporaryArchiveExportMocked test mocked function
func CreateTemporaryArchiveExportMocked(
	t *testing.T,
	temporaryArchiveExportIn *types.TemporaryArchiveExport,
) *types.TemporaryArchiveExport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveExportIn)
	assert.Nil(err, "TemporaryArchiveExport test data corrupted")

	// to json
	dOut, err := json.Marshal(temporaryArchiveExportIn)
	assert.Nil(err, "TemporaryArchiveExport test data corrupted")

	// call service
	cs.On("Post", APIPathPluginsToscaTemporaryArchivesExport, mapIn).Return(dOut, 200, nil)
	temporaryArchiveExportOut, err := ds.CreateTemporaryArchiveExport(mapIn)

	assert.Nil(err, "Error creating temporary archive export")
	assert.Equal(
		temporaryArchiveExportIn,
		temporaryArchiveExportOut,
		"CreateTemporaryArchiveExport returned different temporary archive export",
	)

	return temporaryArchiveExportOut
}

// CreateTemporaryArchiveExportFailErrMocked test mocked function
func CreateTemporaryArchiveExportFailErrMocked(
	t *testing.T,
	temporaryArchiveExportIn *types.TemporaryArchiveExport,
) *types.TemporaryArchiveExport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveExportIn)
	assert.Nil(err, "TemporaryArchiveExport test data corrupted")

	// to json
	dOut, err := json.Marshal(temporaryArchiveExportIn)
	assert.Nil(err, "TemporaryArchiveExport test data corrupted")

	// call service
	cs.On("Post", APIPathPluginsToscaTemporaryArchivesExport, mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	temporaryArchiveExportOut, err := ds.CreateTemporaryArchiveExport(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(temporaryArchiveExportOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return temporaryArchiveExportOut
}

// CreateTemporaryArchiveExportFailStatusMocked test mocked function
func CreateTemporaryArchiveExportFailStatusMocked(
	t *testing.T,
	temporaryArchiveExportIn *types.TemporaryArchiveExport,
) *types.TemporaryArchiveExport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveExportIn)
	assert.Nil(err, "TemporaryArchiveExport test data corrupted")

	// to json
	dOut, err := json.Marshal(temporaryArchiveExportIn)
	assert.Nil(err, "TemporaryArchiveExport test data corrupted")

	// call service
	cs.On("Post", APIPathPluginsToscaTemporaryArchivesExport, mapIn).Return(dOut, 499, nil)
	temporaryArchiveExportOut, err := ds.CreateTemporaryArchiveExport(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(temporaryArchiveExportOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return temporaryArchiveExportOut
}

// CreateTemporaryArchiveExportFailJSONMocked test mocked function
func CreateTemporaryArchiveExportFailJSONMocked(
	t *testing.T,
	temporaryArchiveExportIn *types.TemporaryArchiveExport,
) *types.TemporaryArchiveExport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*temporaryArchiveExportIn)
	assert.Nil(err, "TemporaryArchiveExport test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", APIPathPluginsToscaTemporaryArchivesExport, mapIn).Return(dIn, 200, nil)
	temporaryArchiveExportOut, err := ds.CreateTemporaryArchiveExport(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(temporaryArchiveExportOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return temporaryArchiveExportOut
}

// GetTemporaryArchiveExportTaskMocked test mocked function
func GetTemporaryArchiveExportTaskMocked(
	t *testing.T,
	temporaryArchiveExportTaskIn *types.TemporaryArchiveExportTask,
) *types.TemporaryArchiveExportTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// to json
	dIn, err := json.Marshal(temporaryArchiveExportTaskIn)
	assert.Nil(err, "TemporaryArchiveExportTask test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveExport, temporaryArchiveExportTaskIn.ID)).
		Return(dIn, 200, nil)
	temporaryArchiveExportTaskOut, err := ds.GetTemporaryArchiveExportTask(temporaryArchiveExportTaskIn.ID)

	assert.Nil(err, "Error getting temporary archive export task")
	assert.Equal(
		*temporaryArchiveExportTaskIn,
		*temporaryArchiveExportTaskOut,
		"GetTemporaryArchiveExportTask returned different temporary archive export task",
	)

	return temporaryArchiveExportTaskOut
}

// GetTemporaryArchiveExportTaskFailErrMocked test mocked function
func GetTemporaryArchiveExportTaskFailErrMocked(
	t *testing.T,
	temporaryArchiveExportTaskIn *types.TemporaryArchiveExportTask,
) *types.TemporaryArchiveExportTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// to json
	dIn, err := json.Marshal(temporaryArchiveExportTaskIn)
	assert.Nil(err, "TemporaryArchiveExportTask test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveExport, temporaryArchiveExportTaskIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	temporaryArchiveExportTaskOut, err := ds.GetTemporaryArchiveExportTask(temporaryArchiveExportTaskIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(temporaryArchiveExportTaskOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return temporaryArchiveExportTaskOut
}

// GetTemporaryArchiveExportTaskFailStatusMocked test mocked function
func GetTemporaryArchiveExportTaskFailStatusMocked(
	t *testing.T,
	temporaryArchiveExportTaskIn *types.TemporaryArchiveExportTask,
) *types.TemporaryArchiveExportTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// to json
	dIn, err := json.Marshal(temporaryArchiveExportTaskIn)
	assert.Nil(err, "TemporaryArchiveExportTask test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveExport, temporaryArchiveExportTaskIn.ID)).
		Return(dIn, 499, nil)
	temporaryArchiveExportTaskOut, err := ds.GetTemporaryArchiveExportTask(temporaryArchiveExportTaskIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(temporaryArchiveExportTaskOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return temporaryArchiveExportTaskOut
}

// GetTemporaryArchiveExportTaskFailJSONMocked test mocked function
func GetTemporaryArchiveExportTaskFailJSONMocked(
	t *testing.T,
	temporaryArchiveExportTaskIn *types.TemporaryArchiveExportTask,
) *types.TemporaryArchiveExportTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveExport, temporaryArchiveExportTaskIn.ID)).
		Return(dIn, 200, nil)
	temporaryArchiveExportTaskOut, err := ds.GetTemporaryArchiveExportTask(temporaryArchiveExportTaskIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(temporaryArchiveExportTaskOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return temporaryArchiveExportTaskOut
}

// DownloadTemporaryArchiveExportMocked test mocked function
func DownloadTemporaryArchiveExportMocked(t *testing.T, downloadTemporaryArchiveExportDataIn map[string]string) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	urlSource := downloadTemporaryArchiveExportDataIn["fakeURLToFile"]
	pathFile := downloadTemporaryArchiveExportDataIn["fakeFilePath"]

	// call service
	cs.On("GetFile", urlSource, pathFile).Return(pathFile, 200, nil)
	realFileName, status, err := ds.DownloadTemporaryArchiveExport(urlSource, pathFile)
	assert.Nil(err, "Error downloading temporary archive export file")
	assert.Equal(status, 200, "DownloadTemporaryArchiveExport returned invalid response")
	assert.Equal(realFileName, pathFile, "Invalid downloaded file path")
}

// DownloadTemporaryArchiveExportFailErrMocked test mocked function
func DownloadTemporaryArchiveExportFailErrMocked(t *testing.T, downloadTemporaryArchiveExportDataIn map[string]string) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemporaryArchiveService(cs)
	assert.Nil(err, "Couldn't load temporaryArchive service")
	assert.NotNil(ds, "TemporaryArchive service not instanced")

	urlSource := downloadTemporaryArchiveExportDataIn["fakeURLToFile"]
	pathFile := downloadTemporaryArchiveExportDataIn["fakeFilePath"]

	// call service
	cs.On("GetFile", urlSource, pathFile).Return("", 499, fmt.Errorf("mocked error"))
	_, status, err := ds.DownloadTemporaryArchiveExport(urlSource, pathFile)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(status, 499, "DownloadTemporaryArchiveExport returned an unexpected status code")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}
