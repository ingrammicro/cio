// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// GetCloudApplicationDeployment returns a cloud application deployment by its ID
func (imco *ClientAPI) GetCloudApplicationDeployment(ctx context.Context, deploymentID string,
) (deployment *types.CloudApplicationDeployment, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathPluginsToscaDeployment, deploymentID), true, &deployment)
	if err != nil {
		return nil, status, err
	}
	return deployment, status, nil
}

// DeleteCloudApplicationDeployment deletes a cloud application deployment by its ID
func (imco *ClientAPI) DeleteCloudApplicationDeployment(ctx context.Context, deploymentID string,
) (deployment *types.CloudApplicationDeployment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathPluginsToscaDeployment, deploymentID), true, &deployment)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// CreateCloudApplicationDeploymentTask creates a cloud application deployment task by a given CAT ID
func (imco *ClientAPI) CreateCloudApplicationDeploymentTask(ctx context.Context, catID string,
	deploymentParams *map[string]interface{},
) (deploymentTask *types.CloudApplicationDeploymentTask, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(PathPluginsToscaCatDeploymentTasks, catID),
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
func (imco *ClientAPI) GetCloudApplicationDeploymentTask(ctx context.Context, catID string, deploymentTaskID string,
) (deploymentTask *types.CloudApplicationDeploymentTask, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(PathPluginsToscaCatDeploymentTask, catID, deploymentTaskID),
		true,
		&deploymentTask,
	)
	if err != nil {
		return nil, err
	}
	return deploymentTask, nil
}

// ListCloudApplicationTemplates returns the list of cloud application templates as an array of CloudApplicationTemplate
func (imco *ClientAPI) ListCloudApplicationTemplates(ctx context.Context,
) (templates []*types.CloudApplicationTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathPluginsToscaCats, true, &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

// GetCloudApplicationTemplate returns a cloud application template by its ID
func (imco *ClientAPI) GetCloudApplicationTemplate(ctx context.Context, templateID string,
) (template *types.CloudApplicationTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathPluginsToscaCat, templateID), true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// CreateCloudApplicationTemplate creates a cloud application template
func (imco *ClientAPI) CreateCloudApplicationTemplate(ctx context.Context, catParams *map[string]interface{},
) (template *types.CloudApplicationTemplate, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, PathPluginsToscaCats, catParams, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// ParseMetadataCloudApplicationTemplate process cloud application template metadata
func (imco *ClientAPI) ParseMetadataCloudApplicationTemplate(ctx context.Context, templateID string,
) (template *types.CloudApplicationTemplate, err error) {
	logger.DebugFuncInfo()

	catIn := map[string]interface{}{}
	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(PathPluginsToscaCatParseMetadata, templateID), &catIn, true, &template)
	if err != nil {
		return nil, err
	}
	return template, nil
}

// DeleteCloudApplicationTemplate deletes a cloud application template by its ID
func (imco *ClientAPI) DeleteCloudApplicationTemplate(ctx context.Context, templateID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(PathPluginsToscaCat, templateID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// CreateTemporaryArchive creates a temporary archive
func (imco *ClientAPI) CreateTemporaryArchive(ctx context.Context, temporaryArchiveParams *map[string]interface{},
) (temporaryArchive *types.TemporaryArchive, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, PathPluginsToscaTemporaryArchives, temporaryArchiveParams, true, &temporaryArchive)
	if err != nil {
		return nil, err
	}
	return temporaryArchive, nil
}

// CreateTemporaryArchiveImport creates a temporary archive import
func (imco *ClientAPI) CreateTemporaryArchiveImport(ctx context.Context, temporaryArchiveID string,
	temporaryArchiveImportParams *map[string]interface{},
) (temporaryArchiveImport *types.TemporaryArchiveImport, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		fmt.Sprintf(PathPluginsToscaTemporaryArchiveImport, temporaryArchiveID),
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
func (imco *ClientAPI) GetTemporaryArchiveImport(ctx context.Context, temporaryArchiveImportID string,
) (temporaryArchiveImport *types.TemporaryArchiveImport, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(PathPluginsToscaTemporaryArchiveImport, temporaryArchiveImportID),
		true,
		&temporaryArchiveImport,
	)
	if err != nil {
		return nil, err
	}
	return temporaryArchiveImport, nil
}

// CreateTemporaryArchiveExport creates a temporary archive export
func (imco *ClientAPI) CreateTemporaryArchiveExport(ctx context.Context,
	temporaryArchiveExportParams *map[string]interface{},
) (temporaryArchiveExport *types.TemporaryArchiveExport, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx,
		PathPluginsToscaTemporaryArchivesExport,
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
func (imco *ClientAPI) GetTemporaryArchiveExportTask(ctx context.Context, temporaryArchiveID string,
) (temporaryArchiveExportTask *types.TemporaryArchiveExportTask, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(PathPluginsToscaTemporaryArchiveExport, temporaryArchiveID),
		true,
		&temporaryArchiveExportTask,
	)
	if err != nil {
		return nil, err
	}
	return temporaryArchiveExportTask, nil
}
