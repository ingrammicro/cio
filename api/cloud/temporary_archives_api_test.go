// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewTemporaryArchiveServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewTemporaryArchiveService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestCreateTemporaryArchive(t *testing.T) {
	temporaryArchivesIn := testdata.GetTemporaryArchiveData()
	for _, temporaryArchiveIn := range temporaryArchivesIn {
		CreateTemporaryArchiveMocked(t, temporaryArchiveIn)
		CreateTemporaryArchiveFailErrMocked(t, temporaryArchiveIn)
		CreateTemporaryArchiveFailStatusMocked(t, temporaryArchiveIn)
		CreateTemporaryArchiveFailJSONMocked(t, temporaryArchiveIn)
	}
}

func TestUploadTemporaryArchive(t *testing.T) {
	temporaryArchivesIn := testdata.GetTemporaryArchiveData()
	for _, temporaryArchiveIn := range temporaryArchivesIn {
		UploadTemporaryArchiveMocked(t, temporaryArchiveIn)
		UploadTemporaryArchiveFailStatusMocked(t, temporaryArchiveIn)
		UploadTemporaryArchiveFailErrMocked(t, temporaryArchiveIn)
	}
}

func TestCreateTemporaryArchiveImport(t *testing.T) {
	temporaryArchivesIn := testdata.GetTemporaryArchiveData()
	temporaryArchiveImportsIn := testdata.GetTemporaryArchiveImportData()
	for _, temporaryArchiveImportIn := range temporaryArchiveImportsIn {
		CreateTemporaryArchiveImportMocked(t, temporaryArchivesIn[0].ID, temporaryArchiveImportIn)
		CreateTemporaryArchiveImportFailErrMocked(t, temporaryArchivesIn[0].ID, temporaryArchiveImportIn)
		CreateTemporaryArchiveImportFailStatusMocked(t, temporaryArchivesIn[0].ID, temporaryArchiveImportIn)
		CreateTemporaryArchiveImportFailJSONMocked(t, temporaryArchivesIn[0].ID, temporaryArchiveImportIn)
	}
}

func TestGetTemporaryArchiveImport(t *testing.T) {
	temporaryArchiveImportsIn := testdata.GetTemporaryArchiveImportData()
	for _, temporaryArchiveImportsIn := range temporaryArchiveImportsIn {
		GetTemporaryArchiveImportMocked(t, temporaryArchiveImportsIn)
		GetTemporaryArchiveImportFailErrMocked(t, temporaryArchiveImportsIn)
		GetTemporaryArchiveImportFailStatusMocked(t, temporaryArchiveImportsIn)
		GetTemporaryArchiveImportFailJSONMocked(t, temporaryArchiveImportsIn)
	}
}

func TestCreateTemporaryArchiveExport(t *testing.T) {
	temporaryArchiveExportsIn := testdata.GetTemporaryArchiveExportData()
	for _, temporaryArchiveExportIn := range temporaryArchiveExportsIn {
		CreateTemporaryArchiveExportMocked(t, temporaryArchiveExportIn)
		CreateTemporaryArchiveExportFailErrMocked(t, temporaryArchiveExportIn)
		CreateTemporaryArchiveExportFailStatusMocked(t, temporaryArchiveExportIn)
		CreateTemporaryArchiveExportFailJSONMocked(t, temporaryArchiveExportIn)
	}
}

func TestGetTemporaryArchiveExportTask(t *testing.T) {
	temporaryArchiveExportTasksIn := testdata.GetTemporaryArchiveExportTaskData()
	for _, temporaryArchiveExportTasksIn := range temporaryArchiveExportTasksIn {
		GetTemporaryArchiveExportTaskMocked(t, temporaryArchiveExportTasksIn)
		GetTemporaryArchiveExportTaskFailErrMocked(t, temporaryArchiveExportTasksIn)
		GetTemporaryArchiveExportTaskFailStatusMocked(t, temporaryArchiveExportTasksIn)
		GetTemporaryArchiveExportTaskFailJSONMocked(t, temporaryArchiveExportTasksIn)
	}
}

func TestDownloadTemporaryArchiveExport(t *testing.T) {
	downloadTemporaryArchiveExportDataIn := testdata.GetDownloadTemporaryArchiveExportData()
	DownloadTemporaryArchiveExportMocked(t, downloadTemporaryArchiveExportDataIn)
	DownloadTemporaryArchiveExportFailErrMocked(t, downloadTemporaryArchiveExportDataIn)
}
