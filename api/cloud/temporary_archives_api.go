// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	log "github.com/sirupsen/logrus"

	"github.com/ingrammicro/cio/utils"
)

const APIPathPluginsToscaTemporaryArchives = "/plugins/tosca/temporary_archives"
const APIPathPluginsToscaTemporaryArchiveImport = "/plugins/tosca/temporary_archives/%s/import"
const APIPathPluginsToscaTemporaryArchivesExport = "/plugins/tosca/temporary_archives/export"
const APIPathPluginsToscaTemporaryArchiveExport = "/plugins/tosca/temporary_archives/%s/export"

// TemporaryArchiveService manages server operations
type TemporaryArchiveService struct {
	concertoService utils.ConcertoService
}

// NewTemporaryArchiveService returns a Concerto server service
func NewTemporaryArchiveService(concertoService utils.ConcertoService) (*TemporaryArchiveService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &TemporaryArchiveService{
		concertoService: concertoService,
	}, nil
}

// CreateTemporaryArchive creates a temporary archive
func (tas *TemporaryArchiveService) CreateTemporaryArchive(
	temporaryArchiveParams *map[string]interface{},
) (temporaryArchive *types.TemporaryArchive, err error) {
	log.Debug("CreateTemporaryArchive")

	data, status, err := tas.concertoService.Post(APIPathPluginsToscaTemporaryArchives, temporaryArchiveParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &temporaryArchive); err != nil {
		return nil, err
	}

	return temporaryArchive, nil
}

// UploadTemporaryArchive uploads a temporary archive file
func (tas *TemporaryArchiveService) UploadTemporaryArchive(sourceFilePath string, targetURL string) error {
	log.Debug("UploadTemporaryArchive")

	data, status, err := tas.concertoService.PutFile(sourceFilePath, targetURL)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// CreateTemporaryArchiveImport creates a temporary archive import
func (tas *TemporaryArchiveService) CreateTemporaryArchiveImport(
	temporaryArchiveID string,
	temporaryArchiveImportParams *map[string]interface{},
) (temporaryArchiveImport *types.TemporaryArchiveImport, err error) {
	log.Debug("CreateTemporaryArchiveImport")

	data, status, err := tas.concertoService.Post(
		fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveID),
		temporaryArchiveImportParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &temporaryArchiveImport); err != nil {
		return nil, err
	}

	return temporaryArchiveImport, nil
}

// GetTemporaryArchiveImport returns a temporary archive import by its ID
func (tas *TemporaryArchiveService) GetTemporaryArchiveImport(
	temporaryArchiveImportID string,
) (temporaryArchiveImport *types.TemporaryArchiveImport, err error) {
	log.Debug("GetTemporaryArchiveImport")

	data, status, err := tas.concertoService.Get(
		fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveImport, temporaryArchiveImportID),
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &temporaryArchiveImport); err != nil {
		return nil, err
	}

	return temporaryArchiveImport, nil
}

// CreateTemporaryArchiveExport creates a temporary archive export
func (tas *TemporaryArchiveService) CreateTemporaryArchiveExport(
	temporaryArchiveExportParams *map[string]interface{},
) (temporaryArchiveExport *types.TemporaryArchiveExport, err error) {
	log.Debug("CreateTemporaryArchiveExport")

	data, status, err := tas.concertoService.Post(
		APIPathPluginsToscaTemporaryArchivesExport,
		temporaryArchiveExportParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &temporaryArchiveExport); err != nil {
		return nil, err
	}

	return temporaryArchiveExport, nil
}

// GetTemporaryArchiveExportTask returns a temporary archive export task by its ID
func (tas *TemporaryArchiveService) GetTemporaryArchiveExportTask(
	temporaryArchiveID string,
) (temporaryArchiveExportTask *types.TemporaryArchiveExportTask, err error) {
	log.Debug("GetTemporaryArchiveExportTask")

	data, status, err := tas.concertoService.Get(
		fmt.Sprintf(APIPathPluginsToscaTemporaryArchiveExport, temporaryArchiveID),
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &temporaryArchiveExportTask); err != nil {
		return nil, err
	}

	return temporaryArchiveExportTask, nil
}

// DownloadTemporaryArchiveExport gets a file from given url saving file into given file path
func (tas *TemporaryArchiveService) DownloadTemporaryArchiveExport(
	url string,
	filepath string,
) (realFileName string, status int, err error) {
	log.Debug("DownloadTemporaryArchiveExport")

	realFileName, status, err = tas.concertoService.GetFile(url, filepath, false)
	if err != nil {
		return realFileName, status, err
	}

	return realFileName, status, nil
}
