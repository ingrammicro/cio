// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"

	"github.com/ingrammicro/cio/api/settings"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpPolicyDefinition prepares common resources to send request to Concerto API
func WireUpPolicyDefinition(c *cli.Context) (ds *settings.PolicyDefinitionService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = settings.NewPolicyDefinitionService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up PolicyDefinition service", err)
	}

	return ds, f
}

// PolicyDefinitionList subcommand function
func PolicyDefinitionList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyDefinition(c)

	definitions, err := svc.ListDefinitions()
	if err != nil {
		formatter.PrintFatal("Couldn't receive policy definitions data", err)
	}

	if err = formatter.PrintList(definitions); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// PolicyDefinitionShow subcommand function
func PolicyDefinitionShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyDefinition(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	definition, err := svc.GetDefinition(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive policy definition data", err)
	}

	if err = formatter.PrintItem(*definition); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// PolicyDefinitionCreate subcommand function
func PolicyDefinitionCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyDefinition(c)

	checkRequiredFlags(c, []string{"name", "description"}, formatter)
	if c.IsSet("definition") && c.IsSet("definition-from-file") {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'definition' or 'definition-from-file'",
		)
	}

	definitionIn := map[string]interface{}{
		"name":        c.String("name"),
		"description": c.String("description"),
	}
	if c.IsSet("definition-from-file") {
		defIn, err := utils.ConvertFlagParamsJsonStringFromFileOrStdin(c, c.String("definition-from-file"))
		if err != nil {
			formatter.PrintFatal("Cannot parse definition", err)
		}
		definitionIn["definition"] = defIn
	}
	if c.IsSet("definition") {
		definitionIn["definition"] = c.String("definition")
	}

	definition, err := svc.CreateDefinition(&definitionIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import policy definition", err)
	}

	if err = formatter.PrintItem(*definition); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// PolicyDefinitionUpdate subcommand function
func PolicyDefinitionUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyDefinition(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)
	definitionIn := map[string]interface{}{
		"name": c.String("name"),
	}
	if c.IsSet("description") {
		definitionIn["description"] = c.String("description")
	}

	definition, err := svc.UpdateDefinition(c.String("id"), &definitionIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update policy definition", err)
	}

	if err = formatter.PrintItem(*definition); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// PolicyDefinitionDelete subcommand function
func PolicyDefinitionDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyDefinition(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := svc.DeleteDefinition(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete policy definition", err)
	}
	return nil
}

// PolicyDefinitionListAssignments subcommand function
func PolicyDefinitionListAssignments(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyDefinition(c)
	checkRequiredFlags(c, []string{"id"}, formatter)

	assignments, err := svc.ListAssignments(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive policy assignments data", err)
	}

	if err = formatter.PrintList(assignments); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
