// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloud

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"
	"strings"
	"time"

	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	fServerIds := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerIds,
		Usage: "A list of comma separated server identifiers to be exported"}

	fServerArrayIds := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerArrayIds,
		Usage: "A list of comma separated server array identifiers to be exported"}

	fFilepathExport := cmd.FlagContext{Type: cmd.String, Name: cmd.Filepath, Required: true,
		Usage: "Path and file name to download infrastructure 'csar' file, i.e: --filepath /folder-path/filename.csar"}

	fTimeExport := cmd.FlagContext{Type: cmd.Int64, Name: cmd.Time, DefaultValue: DefaultTimeLapseExportStatusCheck,
		Shorthand: "t", Usage: "Time lapse -seconds- for export status check"}

	fFilepathImport := cmd.FlagContext{Type: cmd.String, Name: cmd.Filepath, Required: true,
		Usage: "Path and file name to infrastructure 'csar' file"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true,
		Usage: "New label name to be associated with infrastructure"}

	fTimeImport := cmd.FlagContext{Type: cmd.Int64, Name: cmd.Time, DefaultValue: DefaultTimeLapseImportStatusCheck,
		Shorthand: "t", Usage: "Time lapse -seconds- for import status check"}

	infrastructureCmd := cmd.NewCommand(cloudCmd, &cmd.CommandContext{
		Use:   "infrastructure",
		Short: "Provides infrastructure archives management"},
	)
	cmd.NewCommand(infrastructureCmd, &cmd.CommandContext{
		Use:          "export",
		Short:        "Exports infrastructure file from IMCO",
		RunMethod:    TemporaryArchiveExport,
		FlagContexts: []cmd.FlagContext{fServerIds, fServerArrayIds, fFilepathExport, fTimeExport}},
	)
	cmd.NewCommand(infrastructureCmd, &cmd.CommandContext{
		Use:          "import",
		Short:        "Imports infrastructure file on IMCO",
		RunMethod:    TemporaryArchiveImport,
		FlagContexts: []cmd.FlagContext{fFilepathImport, fLabel, fTimeImport}},
	)
}

const (
	DefaultTimeLapseExportStatusCheck = 15
	DefaultTimeLapseImportStatusCheck = 30
)

func setExportParams(temporaryArchiveIn map[string]interface{}) (int64, error) {
	logger.DebugFuncInfo()
	if !viper.IsSet(cmd.ServerIds) && !viper.IsSet(cmd.ServerArrayIds) {
		return 0, fmt.Errorf(
			"invalid parameters detected. Please provide at least one: 'server-ids' or 'server-array-ids'",
		)
	}
	if viper.IsSet(cmd.ServerIds) {
		temporaryArchiveIn["server_ids"] = utils.RemoveDuplicates(strings.Split(viper.GetString(cmd.ServerIds), ","))
	}
	if viper.IsSet(cmd.ServerArrayIds) {
		temporaryArchiveIn["server_array_ids"] = utils.RemoveDuplicates(
			strings.Split(viper.GetString(cmd.ServerArrayIds), ","),
		)
	}

	timeLapseExportStatusCheck := viper.GetInt64(cmd.Time)
	if timeLapseExportStatusCheck <= 0 {
		timeLapseExportStatusCheck = DefaultTimeLapseExportStatusCheck
	}
	log.Debug("Time lapse -seconds- for export status check:", timeLapseExportStatusCheck)
	return timeLapseExportStatusCheck, nil
}

func downloadTemporaryArchive(svc *api.ClientAPI, config *configuration.Config,
	temporaryArchiveExportTask *types.TemporaryArchiveExportTask,
) error {
	downloadURL := temporaryArchiveExportTask.DownloadURL(config.APIEndpoint)
	realFileName, status, err := svc.DownloadFile(cmd.GetContext(), downloadURL, viper.GetString(cmd.Filepath), false)
	if err == nil && status != 200 {
		err = fmt.Errorf("obtained non-ok response when downloading export file %s", downloadURL)
	}
	if err != nil {
		return err
	}
	log.Info("Available at:", realFileName)
	return nil
}

// TemporaryArchiveExport subcommand function
func TemporaryArchiveExport() error {
	logger.DebugFuncInfo()
	svc, config, formatter := cli.WireUpAPIClient()

	temporaryArchiveIn := map[string]interface{}{
		"is_mock": false,
	}

	timeLapseExportStatusCheck, err := setExportParams(temporaryArchiveIn)
	if err != nil {
		return err
	}

	temporaryArchiveExport, err := svc.CreateTemporaryArchiveExport(cmd.GetContext(), &temporaryArchiveIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create temporary archive", err)
	}

	log.Info("ID: ", temporaryArchiveExport.ID)
	log.Info("Task ID: ", temporaryArchiveExport.TaskID)
	log.Info("Exporting...")
	temporaryArchiveExportTask := new(types.TemporaryArchiveExportTask)
	for {
		temporaryArchiveExportTask, err = svc.GetTemporaryArchiveExportTask(cmd.GetContext(), temporaryArchiveExport.ID)
		if err != nil {
			formatter.PrintFatal("Couldn't get temporary archive", err)
		}
		log.Info("State: ", temporaryArchiveExportTask.State)

		if temporaryArchiveExportTask.State == "finished" {
			if err = formatter.PrintItem(*temporaryArchiveExportTask); err != nil {
				formatter.PrintFatal(cmd.PrintFormatError, err)
			}
			if temporaryArchiveExportTask.ErrorMessage != "" {
				formatter.PrintFatal(
					"Couldn't export infrastructure file",
					fmt.Errorf("%s", temporaryArchiveExportTask.ErrorMessage),
				)
			}
			break
		}
		time.Sleep(time.Duration(timeLapseExportStatusCheck) * time.Second)
	}

	return downloadTemporaryArchive(svc, config, temporaryArchiveExportTask)
}

func getTimeLapseImportStatusCheckParam() int64 {
	logger.DebugFuncInfo()
	timeLapseImportStatusCheck := viper.GetInt64(cmd.Time)
	if timeLapseImportStatusCheck <= 0 {
		timeLapseImportStatusCheck = DefaultTimeLapseImportStatusCheck
	}
	log.Debug("Time lapse -seconds- for import status check:", timeLapseImportStatusCheck)

	return timeLapseImportStatusCheck
}

// TemporaryArchiveImport subcommand function
func TemporaryArchiveImport() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	sourceFilePath := viper.GetString(cmd.Filepath)
	if !utils.FileExists(sourceFilePath) {
		formatter.PrintFatal("Invalid file path", fmt.Errorf("no such file or directory: %s", sourceFilePath))
	}
	temporaryArchiveIn := map[string]interface{}{
		"is_mock":    false,
		"label_name": viper.GetString(cmd.Label),
	}

	timeLapseImportStatusCheck := getTimeLapseImportStatusCheckParam()

	temporaryArchive, err := svc.CreateTemporaryArchive(cmd.GetContext(), &temporaryArchiveIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create temporary archive", err)
	}

	err = svc.UploadFile(cmd.GetContext(), sourceFilePath, temporaryArchive.UploadURL)
	if err != nil {
		formatter.PrintFatal("Couldn't upload temporary archive", err)
	}

	temporaryArchiveImport, err := svc.CreateTemporaryArchiveImport(
		cmd.GetContext(),
		temporaryArchive.ID,
		&temporaryArchiveIn,
	)
	if err != nil {
		formatter.PrintFatal("Couldn't create temporary archive import", err)
	}

	log.Info("ID: ", temporaryArchiveImport.ID)
	log.Info("Archive ID: ", temporaryArchiveImport.ArchiveID)
	log.Info("Label: ", temporaryArchiveImport.LabelName)
	log.Info("Importing...")
	for {
		temporaryArchiveImport, err = svc.GetTemporaryArchiveImport(cmd.GetContext(), temporaryArchive.ID)
		if err != nil {
			formatter.PrintFatal("Couldn't get temporary archive import", err)
		}
		log.Info("State: ", temporaryArchiveImport.State)

		if temporaryArchiveImport.State == "finished" {
			if err = formatter.PrintItem(*temporaryArchiveImport); err != nil {
				formatter.PrintFatal(cmd.PrintFormatError, err)
			}
			if temporaryArchiveImport.ErrorMessage != "" {
				formatter.PrintFatal(
					"Couldn't import infrastructure file",
					fmt.Errorf("%s", temporaryArchiveImport.ErrorMessage),
				)
			}
			break
		}
		time.Sleep(time.Duration(timeLapseImportStatusCheck) * time.Second)
	}
	return nil
}
