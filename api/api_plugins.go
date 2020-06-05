// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// GetCloudApplicationDeployment returns a cloud application deployment by its ID
func (imco *IMCOClient) GetCloudApplicationDeployment(deploymentID string,
) (deployment *types.CloudApplicationDeployment, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.getAndCheck(fmt.Sprintf(pathPluginsToscaDeployment, deploymentID), true, &deployment)
	if err != nil {
		return nil, status, err
	}
	return deployment, status, nil
}

// DeleteCloudApplicationDeployment deletes a cloud application deployment by its ID
func (imco *IMCOClient) DeleteCloudApplicationDeployment(deploymentID string,
) (deployment *types.CloudApplicationDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathPluginsToscaDeployment, deploymentID), true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// CreateCloudApplicationDeploymentTask creates a cloud application deployment task by a given CAT ID
func (imco *IMCOClient) CreateCloudApplicationDeploymentTask(catID string, deploymentParams *map[string]interface{},
) (deploymentTask *types.CloudApplicationDeploymentTask, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
		fmt.Sprintf(pathPluginsToscaCatDeploymentTasks, catID),
		deploymentParams,
		true,
		&deploymentTask,
	)
	if err != nil {
		return nil, err
	}
	return deploymentTask, nil
}

// GetCloudApplicationDeploymentTask gets a cloud application deployment task by its ID and given CAT ID
func (imco *IMCOClient) GetCloudApplicationDeploymentTask(catID string, deploymentTaskID string,
) (deploymentTask *types.CloudApplicationDeploymentTask, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathPluginsToscaCatDeploymentTask, catID, deploymentTaskID),
		true,
		&deploymentTask,
	)
	if err != nil {
		return nil, err
	}
	return deploymentTask, nil
}

// ListCloudApplicationTemplates returns the list of cloud application templates as an array of CloudApplicationTemplate
func (imco *IMCOClient) ListCloudApplicationTemplates() (templates []*types.CloudApplicationTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathPluginsToscaCats, true, &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

// GetCloudApplicationTemplate returns a cloud application template by its ID
func (imco *IMCOClient) GetCloudApplicationTemplate(templateID string,
) (template *types.CloudApplicationTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathPluginsToscaCat, templateID), true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CreateCloudApplicationTemplate creates a cloud application template
func (imco *IMCOClient) CreateCloudApplicationTemplate(catParams *map[string]interface{},
) (template *types.CloudApplicationTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathPluginsToscaCats, catParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// ParseMetadataCloudApplicationTemplate process cloud application template metadata
func (imco *IMCOClient) ParseMetadataCloudApplicationTemplate(templateID string,
) (template *types.CloudApplicationTemplate, err error) {
	logger.DebugFuncInfo()

	catIn := map[string]interface{}{}
	_, err = imco.putAndCheck(fmt.Sprintf(pathPluginsToscaCatParseMetadata, templateID), &catIn, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// DeleteCloudApplicationTemplate deletes a cloud application template by its ID
func (imco *IMCOClient) DeleteCloudApplicationTemplate(templateID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathPluginsToscaCat, templateID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// CreateTemporaryArchive creates a temporary archive
func (imco *IMCOClient) CreateTemporaryArchive(temporaryArchiveParams *map[string]interface{},
) (temporaryArchive *types.TemporaryArchive, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathPluginsToscaTemporaryArchives, temporaryArchiveParams, true, &temporaryArchive)
	if err != nil {
		return nil, err
	}
	return temporaryArchive, nil
}

// CreateTemporaryArchiveImport creates a temporary archive import
func (imco *IMCOClient) CreateTemporaryArchiveImport(temporaryArchiveID string,
	temporaryArchiveImportParams *map[string]interface{},
) (temporaryArchiveImport *types.TemporaryArchiveImport, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
		fmt.Sprintf(pathPluginsToscaTemporaryArchiveImport, temporaryArchiveID),
		temporaryArchiveImportParams,
		true,
		&temporaryArchiveImport,
	)
	if err != nil {
		return nil, err
	}
	return temporaryArchiveImport, nil
}

// GetTemporaryArchiveImport returns a temporary archive import by its ID
func (imco *IMCOClient) GetTemporaryArchiveImport(temporaryArchiveImportID string,
) (temporaryArchiveImport *types.TemporaryArchiveImport, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathPluginsToscaTemporaryArchiveImport, temporaryArchiveImportID),
		true,
		&temporaryArchiveImport,
	)
	if err != nil {
		return nil, err
	}
	return temporaryArchiveImport, nil
}

// CreateTemporaryArchiveExport creates a temporary archive export
func (imco *IMCOClient) CreateTemporaryArchiveExport(temporaryArchiveExportParams *map[string]interface{},
) (temporaryArchiveExport *types.TemporaryArchiveExport, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(
		pathPluginsToscaTemporaryArchivesExport,
		temporaryArchiveExportParams,
		true,
		&temporaryArchiveExport,
	)
	if err != nil {
		return nil, err
	}
	return temporaryArchiveExport, nil
}

// GetTemporaryArchiveExportTask returns a temporary archive export task by its ID
func (imco *IMCOClient) GetTemporaryArchiveExportTask(temporaryArchiveID string,
) (temporaryArchiveExportTask *types.TemporaryArchiveExportTask, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathPluginsToscaTemporaryArchiveExport, temporaryArchiveID),
		true,
		&temporaryArchiveExportTask,
	)
	if err != nil {
		return nil, err
	}
	return temporaryArchiveExportTask, nil
}
