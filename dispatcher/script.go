// Copyright (c) 2017-2021 Ingram Micro Inc.

package dispatcher

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ingrammicro/cio/api/dispatcher"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/utils/format"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func cmdBoot(c *cli.Context) error {
	execute(c, "boot", "")
	return nil
}

func cmdOperational(c *cli.Context) error {
	execute(c, "operational", c.Args().Get(0))
	return nil
}

func cmdShutdown(c *cli.Context) error {
	execute(c, "shutdown", "")
	return nil
}

func getDispatcherScriptCharacterization(
	dispatcherSvc *dispatcher.DispatcherService, formatter format.Formatter, phase, scriptCharacterizationUUID string,
) []*types.ScriptCharacterization {
	var scriptChars []*types.ScriptCharacterization
	var err error
	log.Debugf("Current Script Characterization %s (UUID=%s)", phase, scriptCharacterizationUUID)
	if scriptCharacterizationUUID == "" {
		scriptChars, err = dispatcherSvc.GetDispatcherScriptCharacterizationsByType(phase)
	} else {
		var scriptChar *types.ScriptCharacterization
		scriptChar, err = dispatcherSvc.GetDispatcherScriptCharacterizationByUUID(scriptCharacterizationUUID)
		scriptChars = []*types.ScriptCharacterization{scriptChar}
	}
	if err != nil {
		formatter.PrintFatal("Couldn't receive Script Characterization data", err)
	}
	return scriptChars
}

func setEnvironmentVariables(formatter format.Formatter, parameters map[string]string) {
	log.Infof("Environment Variables")
	for index, value := range parameters {
		if err := os.Setenv(index, value); err != nil {
			formatter.PrintFatal(fmt.Sprintf("Couldn't set environment variable %s:%s", index, value), err)
		}
		log.Infof("\t - %s=%s", index, value)
	}
}

func getAttachmentDir(formatter format.Formatter) string {
	attachmentDir := os.Getenv("ATTACHMENT_DIR")
	if err := os.Mkdir(attachmentDir, 0777); err != nil {
		formatter.PrintFatal("Couldn't create attachments directory", err)
	}
	return attachmentDir
}

func downloadAttachments(
	attachmentDir string,
	attachmentPaths []string,
	dispatcherSvc *dispatcher.DispatcherService,
	config *utils.Config,
	formatter format.Formatter,
) {
	log.Infof("Attachment Folder: %s", attachmentDir)
	log.Infof("Attachments")
	for _, endpoint := range attachmentPaths {
		realFileName, _, err := dispatcherSvc.DownloadAttachment(
			fmt.Sprintf("%s%s", config.APIEndpoint, endpoint),
			attachmentDir,
		)
		if err != nil {
			formatter.PrintFatal("Couldn't download attachment", err)
		}
		log.Infof("\t - %s --> %s", endpoint, realFileName)
	}
}

func execute(c *cli.Context, phase string, scriptCharacterizationUUID string) {
	dispatcherSvc, config, formatter := cmd.WireUpDispatcher(c)
	scriptChars := getDispatcherScriptCharacterization(dispatcherSvc, formatter, phase, scriptCharacterizationUUID)

	for _, sc := range scriptChars {
		log.Infof("------------------------------------------------------------------------------------------------")
		path, err := ioutil.TempDir("", "cio")
		if err != nil {
			formatter.PrintFatal("Couldn't create temporary directory", err)
		}
		defer os.RemoveAll(path)

		if err = os.Setenv("ATTACHMENT_DIR", fmt.Sprintf("%s/%s", path, "attachments")); err != nil {
			formatter.PrintFatal("Couldn't set attachments directory as environment variable", err)
		}

		log.Infof("UUID: %s", sc.UUID)
		log.Infof("Home Folder: %s", path)

		attachmentDir := getAttachmentDir(formatter)

		// Setting up environment Variables
		setEnvironmentVariables(formatter, sc.Parameters)

		if len(sc.Script.AttachmentPaths) > 0 {
			downloadAttachments(attachmentDir, sc.Script.AttachmentPaths, dispatcherSvc, config, formatter)
		}

		output, exitCode, startedAt, finishedAt := utils.ExecCode(sc.Script.Code, path, sc.Script.UUID)
		scriptConclusionIn := map[string]interface{}{
			"script_characterization_id": sc.UUID,
			"output":                     output,
			"exit_code":                  exitCode,
			"started_at":                 startedAt.Format(utils.TimeStampLayout),
			"finished_at":                finishedAt.Format(utils.TimeStampLayout),
		}
		scriptConclusionRootIn := map[string]interface{}{
			"script_conclusion": scriptConclusionIn,
		}
		_, _, err = dispatcherSvc.ReportScriptConclusions(&scriptConclusionRootIn)
		if err != nil {
			formatter.PrintFatal("Couldn't send script_conclusions report data", err)
		}
		log.Infof("------------------------------------------------------------------------------------------------")
	}
}
