package testdata

import "github.com/ingrammicro/cio/api/types"

// GetTemporaryArchiveData loads test data
func GetTemporaryArchiveData() []*types.TemporaryArchive {
	return []*types.TemporaryArchive{
		{
			ID:        "fakeID0",
			UploadURL: "fakeUploadURL0",
		},
	}
}

// GetTemporaryArchiveImportData loads test data
func GetTemporaryArchiveImportData() []*types.TemporaryArchiveImport {
	return []*types.TemporaryArchiveImport{
		{
			ID:           "fakeID0",
			Type:         "fakeType0",
			LabelName:    "fakeLabelName0",
			LabelID:      "fakeLabelID0",
			State:        "fakeState0",
			Outputs:      "fakeOutputs0",
			UserID:       "fakeUserID0",
			ArchiveID:    "fakeArchiveID0",
			ErrorMessage: "fakeErrorMessage0",
		},
	}
}

// GetTemporaryArchiveExportData loads test data
func GetTemporaryArchiveExportData() []*types.TemporaryArchiveExport {
	return []*types.TemporaryArchiveExport{
		{
			ID:     "fakeID0",
			TaskID: "fakeTaskID0",
		},
	}
}

// GetTemporaryArchiveExportTaskData loads test data
func GetTemporaryArchiveExportTaskData() []*types.TemporaryArchiveExportTask {
	return []*types.TemporaryArchiveExportTask{
		{
			ID:           "fakeID0",
			Type:         "fakeType0",
			State:        "fakeState0",
			FileAccess:   "fakeFileAccess0",
			UserID:       "fakeUserID0",
			ArchiveID:    "fakeArchiveID0",
			ErrorMessage: "fakeErrorMessage0",
		},
	}
}

// GetDownloadTemporaryArchiveExportData loads test data
func GetDownloadTemporaryArchiveExportData() map[string]string {
	return map[string]string{
		"fakeURLToFile": "http://fakeURLToFile.xxx/filename.csar",
		"fakeFilePath":  "/temporary_archives/exports/filename.csar",
	}
}
