// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloudapplications

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/utils"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "CAT Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the CAT"}

	fFilepath := cmd.FlagContext{Type: cmd.String, Name: cmd.Filepath, Required: true, Usage: "path to CAT file"}

	templatesCmd := cmd.NewCommand(cloudApplicationsCmd, &cmd.CommandContext{
		Use:   "templates",
		Short: "Provides information about CAT"},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists CATs",
		RunMethod: CloudApplicationTemplateList},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows CAT",
		RunMethod:    CloudApplicationTemplateShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "upload",
		Short:        "Uploads a CAT",
		RunMethod:    CloudApplicationTemplateUpload,
		FlagContexts: []cmd.FlagContext{fName, fFilepath}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a CAT",
		RunMethod:    CloudApplicationTemplateDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// CloudApplicationTemplateList subcommand function
func CloudApplicationTemplateList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cats, err := svc.ListCloudApplicationTemplates(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive cloud application templates data", err)
		return err
	}

	if err = formatter.PrintList(cats); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CloudApplicationTemplateShow subcommand function
func CloudApplicationTemplateShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cat, err := svc.GetCloudApplicationTemplate(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive cloud application template data", err)
		return err
	}

	if err = formatter.PrintItem(*cat); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CloudApplicationTemplateUpload subcommand function
func CloudApplicationTemplateUpload() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	sourceFilePath := viper.GetString(cmd.Filepath)
	if !utils.FileExists(sourceFilePath) {
		e := fmt.Errorf("no such file or directory: %s", sourceFilePath)
		formatter.PrintError("Invalid file path", e)
		return e
	}

	catIn := map[string]interface{}{
		"is_mock": false,
	}
	cat, err := svc.CreateCloudApplicationTemplate(cmd.GetContext(), &catIn)
	if err != nil {
		formatter.PrintError("Couldn't receive cloud application template data", err)
		return err
	}

	catID := cat.ID
	err = svc.UploadFile(cmd.GetContext(), sourceFilePath, cat.UploadURL)
	if err != nil {
		cleanTemplate(catID)
		formatter.PrintError("Couldn't upload cloud application template data", err)
		return err
	}

	cat, err = svc.ParseMetadataCloudApplicationTemplate(cmd.GetContext(), catID)
	if err != nil {
		cleanTemplate(catID)
		formatter.PrintError("Couldn't parse cloud application template metadata", err)
		return err
	}

	if err = formatter.PrintItem(*cat); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// cleanTemplate deletes cloud application template. Ideally for cleaning at uploading error cases
func cleanTemplate(catID string) {
	svc, _, formatter := cli.WireUpAPIClient()
	if err := svc.DeleteCloudApplicationTemplate(cmd.GetContext(), catID); err != nil {
		formatter.PrintError("Couldn't clean failed cloud application template", err)
	}
}

// CloudApplicationTemplateDelete subcommand function
func CloudApplicationTemplateDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteCloudApplicationTemplate(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete cloud application template", err)
		return err
	}
	return nil
}
