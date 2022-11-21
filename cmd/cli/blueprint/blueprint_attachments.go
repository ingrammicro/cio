// Copyright (c) 2017-2022 Ingram Micro Inc.

package blueprint

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Attachment Id"}

	fFilepath := cmd.FlagContext{Type: cmd.String, Name: cmd.Filepath, Required: true,
		Usage: "path and file name to download attachment file, i.e: --filename /folder-path/filename.ext"}

	attachmentsCmd := cmd.NewCommand(blueprintCmd, &cmd.CommandContext{
		Use:   "attachments",
		Short: "Allow the user to manage the attachments they want to store on the servers"},
	)
	cmd.NewCommand(attachmentsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about a specific attachment",
		RunMethod:    AttachmentShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(attachmentsCmd, &cmd.CommandContext{
		Use:          "download",
		Short:        "Downloads an attachment",
		RunMethod:    AttachmentDownload,
		FlagContexts: []cmd.FlagContext{fId, fFilepath}},
	)
	cmd.NewCommand(attachmentsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes an attachment",
		RunMethod:    AttachmentDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// AttachmentShow subcommand function
func AttachmentShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	attachment, err := svc.GetAttachment(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive attachment data", err)
		return err
	}

	if err = formatter.PrintItem(*attachment); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// AttachmentDownload subcommand function
func AttachmentDownload() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	attachment, err := svc.GetAttachment(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive attachment data", err)
		return err
	}

	realFileName, status, err := svc.DownloadFile(
		ctx,
		attachment.DownloadURL,
		viper.GetString(cmd.Filepath),
		false,
	)
	if err == nil && status != 200 {
		err = fmt.Errorf("obtained non-ok response when downloading attachment %s", attachment.DownloadURL)
	}
	if err != nil {
		formatter.PrintError("Couldn't download attachment", err)
		return err
	}
	log.Info("Available at:", realFileName)
	return nil
}

// AttachmentDelete subcommand function
func AttachmentDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteAttachment(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete attachment", err)
		return err
	}
	return nil
}
