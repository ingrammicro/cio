// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"

	"github.com/ingrammicro/cio/api/cloudapplication"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpCloudApplicationTemplate prepares common resources to send request to Concerto API
func WireUpCloudApplicationTemplate(
	c *cli.Context,
) (ds *cloudapplication.CloudApplicationTemplateService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloudapplication.NewCloudApplicationTemplateService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up CloudApplicationTemplate service", err)
	}

	return ds, f
}

// CloudApplicationTemplateList subcommand function
func CloudApplicationTemplateList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationTemplate(c)

	cats, err := svc.ListTemplates()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloud application templates data", err)
	}

	if err = formatter.PrintList(cats); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CloudApplicationTemplateShow subcommand function
func CloudApplicationTemplateShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cat, err := svc.GetTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloud application template data", err)
	}

	if err = formatter.PrintItem(*cat); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CloudApplicationTemplateUpload subcommand function
func CloudApplicationTemplateUpload(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationTemplate(c)

	checkRequiredFlags(c, []string{"name", "filepath"}, formatter)

	sourceFilePath := c.String("filepath")
	if !utils.FileExists(sourceFilePath) {
		formatter.PrintFatal("Invalid file path", fmt.Errorf("no such file or directory: %s", sourceFilePath))
	}

	catIn := map[string]interface{}{
		"is_mock": false,
	}
	cat, err := svc.CreateTemplate(&catIn)
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloud application template data", err)
	}

	catID := cat.ID
	err = svc.UploadTemplate(sourceFilePath, cat.UploadURL)
	if err != nil {
		cleanTemplate(c, catID)
		formatter.PrintFatal("Couldn't upload cloud application template data", err)
	}

	cat, err = svc.ParseMetadataTemplate(catID)
	if err != nil {
		cleanTemplate(c, catID)
		formatter.PrintFatal("Couldn't parse cloud application template metadata", err)
	}

	if err = formatter.PrintItem(*cat); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// cleanTemplate deletes cloud application template. Ideally for cleaning at uploading error cases
func cleanTemplate(c *cli.Context, catID string) {
	svc, formatter := WireUpCloudApplicationTemplate(c)
	if err := svc.DeleteTemplate(catID); err != nil {
		formatter.PrintError("Couldn't clean failed cloud application template", err)
	}
}

// CloudApplicationTemplateDelete subcommand function
func CloudApplicationTemplateDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := svc.DeleteTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete cloud application template", err)
	}
	return nil
}
