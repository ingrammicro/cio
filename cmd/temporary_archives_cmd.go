package cmd

import (
	"fmt"
	"github.com/ingrammicro/cio/api/cloud"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"strings"
	"time"
)

const (
	DefaultTimeLapseExportStatusCheck = 15
	DefaultTimeLapseImportStatusCheck = 30
)

// WireUpTemporaryArchive prepares common resources to send request to Concerto API
func WireUpTemporaryArchive(c *cli.Context) (tas *cloud.TemporaryArchiveService, config *utils.Config, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	tas, err = cloud.NewTemporaryArchiveService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up TemporaryArchive service", err)
	}

	return tas, config, f
}

// TemporaryArchiveExport subcommand function
func TemporaryArchiveExport(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, config, formatter := WireUpTemporaryArchive(c)

	checkRequiredFlags(c, []string{"filepath"}, formatter)
	temporaryArchiveIn := map[string]interface{}{
		"is_mock": false,
	}
	if !c.IsSet("server-ids") && !c.IsSet("server-array-ids") {
		return fmt.Errorf("invalid parameters detected. Please provide at least one: 'server-ids' or 'server-array-ids'")
	}
	if c.IsSet("server-ids") {
		temporaryArchiveIn["server_ids"] = utils.RemoveDuplicates(strings.Split(c.String("server-ids"), ","))
	}
	if c.IsSet("server-array-ids") {
		temporaryArchiveIn["server_array_ids"] = utils.RemoveDuplicates(strings.Split(c.String("server-array-ids"), ","))
	}

	timeLapseExportStatusCheck := c.Int64("time")
	if !(timeLapseExportStatusCheck > 0) {
		timeLapseExportStatusCheck = DefaultTimeLapseExportStatusCheck
	}
	log.Debug("Time lapse -seconds- for export status check:", timeLapseExportStatusCheck)

	temporaryArchiveExport, err := svc.CreateTemporaryArchiveExport(&temporaryArchiveIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create temporary archive", err)
	}

	log.Info("ID: ", temporaryArchiveExport.ID)
	log.Info("Task ID: ", temporaryArchiveExport.TaskID)
	log.Info("Exporting...")
	temporaryArchiveExportTask := new(types.TemporaryArchiveExportTask)
	for {
		temporaryArchiveExportTask, err = svc.GetTemporaryArchiveExportTask(temporaryArchiveExport.ID)
		if err != nil {
			formatter.PrintFatal("Couldn't get temporary archive", err)
		}
		log.Info("State: ", temporaryArchiveExportTask.State)

		if temporaryArchiveExportTask.State == "finished" {
			if err = formatter.PrintItem(*temporaryArchiveExportTask); err != nil {
				formatter.PrintFatal("Couldn't print/format result", err)
			}
			if temporaryArchiveExportTask.ErrorMessage != "" {
				formatter.PrintFatal("Couldn't export infrastructure file", fmt.Errorf("%s", temporaryArchiveExportTask.ErrorMessage))
			}
			break
		}
		time.Sleep(time.Duration(timeLapseExportStatusCheck) * time.Second)
	}

	downloadURL := temporaryArchiveExportTask.DownloadURL(config.APIEndpoint)
	realFileName, status, err := svc.DownloadTemporaryArchiveExport(downloadURL, c.String("filepath"))
	if err == nil && status != 200 {
		err = fmt.Errorf("obtained non-ok response when downloading export file %s", downloadURL)
	}
	if err != nil {
		return err
	}
	log.Info("Available at:", realFileName)

	return nil
}

// TemporaryArchiveImport subcommand function
func TemporaryArchiveImport(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, _, formatter := WireUpTemporaryArchive(c)

	checkRequiredFlags(c, []string{"filepath", "label"}, formatter)
	sourceFilePath := c.String("filepath")
	if !utils.FileExists(sourceFilePath) {
		formatter.PrintFatal("Invalid file path", fmt.Errorf("no such file or directory: %s", sourceFilePath))
	}
	temporaryArchiveIn := map[string]interface{}{
		"is_mock":    false,
		"label_name": c.String("label"),
	}

	timeLapseImportStatusCheck := c.Int64("time")
	if !(timeLapseImportStatusCheck > 0) {
		timeLapseImportStatusCheck = DefaultTimeLapseImportStatusCheck
	}
	log.Debug("Time lapse -seconds- for import status check:", timeLapseImportStatusCheck)

	temporaryArchive, err := svc.CreateTemporaryArchive(&temporaryArchiveIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create temporary archive", err)
	}

	err = svc.UploadTemporaryArchive(sourceFilePath, temporaryArchive.UploadURL)
	if err != nil {
		formatter.PrintFatal("Couldn't upload temporary archive", err)
	}

	temporaryArchiveImport, err := svc.CreateTemporaryArchiveImport(temporaryArchive.ID, &temporaryArchiveIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create temporary archive import", err)
	}

	log.Info("ID: ", temporaryArchiveImport.ID)
	log.Info("Archive ID: ", temporaryArchiveImport.ArchiveID)
	log.Info("Label: ", temporaryArchiveImport.LabelName)
	log.Info("Importing...")
	for {
		temporaryArchiveImport, err = svc.GetTemporaryArchiveImport(temporaryArchive.ID)
		if err != nil {
			formatter.PrintFatal("Couldn't get temporary archive import", err)
		}
		log.Info("State: ", temporaryArchiveImport.State)

		if temporaryArchiveImport.State == "finished" {
			if err = formatter.PrintItem(*temporaryArchiveImport); err != nil {
				formatter.PrintFatal("Couldn't print/format result", err)
			}
			if temporaryArchiveImport.ErrorMessage != "" {
				formatter.PrintFatal("Couldn't import infrastructure file", fmt.Errorf("%s", temporaryArchiveImport.ErrorMessage))
			}
			break
		}
		time.Sleep(time.Duration(timeLapseImportStatusCheck) * time.Second)
	}
	return nil
}
