// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"

	"github.com/ingrammicro/cio/api/settings"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpPolicyAssignment prepares common resources to send request to Concerto API
func WireUpPolicyAssignment(c *cli.Context) (ds *settings.PolicyAssignmentService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = settings.NewPolicyAssignmentService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up PolicyAssignment service", err)
	}

	return ds, f
}

// PolicyAssignmentList subcommand function
func PolicyAssignmentList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyAssignment(c)

	checkRequiredFlags(c, []string{"cloud-account-id"}, formatter)
	assignments, err := svc.ListAssignments(c.String("cloud-account-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive policy assignments data", err)
	}

	if err = formatter.PrintList(assignments); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// PolicyAssignmentShow subcommand function
func PolicyAssignmentShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyAssignment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	assignment, err := svc.GetAssignment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive policy assignment data", err)
	}

	if err = formatter.PrintItem(*assignment); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// PolicyAssignmentCreate subcommand function
func PolicyAssignmentCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyAssignment(c)

	checkRequiredFlags(c, []string{"name", "description", "cloud-account-id", "definition-id"}, formatter)
	if c.IsSet("parameters") && c.IsSet("parameters-from-file") {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'parameters' or 'parameters-from-file'",
		)
	}

	assignmentIn := map[string]interface{}{
		"name":             c.String("name"),
		"description":      c.String("description"),
		"cloud_account_id": c.String("cloud-account-id"),
	}
	if c.IsSet("parameters-from-file") {
		defIn, err := utils.ConvertFlagParamsJsonFromFileOrStdin(c, c.String("parameters-from-file"))
		if err != nil {
			formatter.PrintFatal("Cannot parse parameters", err)
		}
		assignmentIn["parameter_values"] = defIn
	}
	if c.IsSet("parameters") {
		params, err := utils.FlagConvertParamsJSON(c, []string{"parameters"})
		if err != nil {
			formatter.PrintFatal("Cannot parse parameters", err)
		}
		assignmentIn["parameter_values"] = (*params)["parameters"]
	}

	assignment, err := svc.CreateAssignment(c.String("definition-id"), &assignmentIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import policy assignment", err)
	}

	if err = formatter.PrintItem(*assignment); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// PolicyAssignmentUpdate subcommand function
func PolicyAssignmentUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyAssignment(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)
	assignmentIn := map[string]interface{}{
		"name": c.String("name"),
	}

	if c.IsSet("description") {
		assignmentIn["description"] = c.String("description")
	}

	assignment, err := svc.UpdateAssignment(c.String("id"), &assignmentIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update policy assignment", err)
	}

	if err = formatter.PrintItem(*assignment); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// PolicyAssignmentDelete subcommand function
func PolicyAssignmentDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpPolicyAssignment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	assignment, err := svc.DeleteAssignment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete policy assignment", err)
	}

	if err = formatter.PrintItem(*assignment); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
