// Copyright (c) 2017-2022 Ingram Micro Inc.

package settings

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Assignment Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the policy assignment"}

	fDescription := cmd.FlagContext{Type: cmd.String, Name: cmd.Description,
		Usage: "Description of the policy assignment"}
	fDescriptionReq := fDescription
	fDescriptionReq.Required = true

	fDefinitionId := cmd.FlagContext{Type: cmd.String, Name: cmd.DefinitionId,
		Usage: "Identifier of the policy definition to be assigned"}

	fParameters := cmd.FlagContext{Type: cmd.String, Name: cmd.Parameters,
		Usage: "The parameters used to configure the policy assignment, as a json formatted parameter. \n\t" +
			"i.e: --parameters '{\"param1\":\"val1\",\"param2\":\"val2\",\"param3\":{\"id\":\"val3\"},\"param4\":true}'"}

	fParametersFromFile := cmd.FlagContext{Type: cmd.String, Name: cmd.ParametersFromFile,
		Usage: "The parameters used to configure the policy assignment, from file or STDIN, " +
			"as a json formatted parameter. \n\t" +
			"From file: --parameters-from-file params.json \n\t" +
			"From STDIN: --parameters-from-file -"}

	assignmentsCmd := cmd.NewCommand(policiesCmd, &cmd.CommandContext{
		Use:   "assignments",
		Short: "Provides information about policy assignments"},
	)
	cmd.NewCommand(assignmentsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "lists policy assignments for a given cloud account",
		RunMethod:    PolicyAssignmentList,
		FlagContexts: []cmd.FlagContext{fCloudAccountId}},
	)
	cmd.NewCommand(assignmentsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows policy assignments",
		RunMethod:    PolicyAssignmentShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(assignmentsCmd, &cmd.CommandContext{
		Use:       "create",
		Short:     "Creates policy assignment",
		RunMethod: PolicyAssignmentCreate,
		FlagContexts: []cmd.FlagContext{
			fName,
			fDescriptionReq,
			fCloudAccountId,
			fDefinitionId,
			fParameters,
			fParametersFromFile}},
	)
	cmd.NewCommand(assignmentsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing policy assignment identified by the given id",
		RunMethod:    PolicyAssignmentUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName, fDescription}},
	)
	cmd.NewCommand(assignmentsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a policy assignment",
		RunMethod:    PolicyAssignmentDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// PolicyAssignmentList subcommand function
func PolicyAssignmentList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	assignments, err := svc.ListPolicyAssignments(cmd.GetContext(), viper.GetString(cmd.CloudAccountId))
	if err != nil {
		formatter.PrintError("Couldn't receive policy assignments data", err)
		return err
	}

	if err = formatter.PrintList(assignments); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// PolicyAssignmentShow subcommand function
func PolicyAssignmentShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	assignment, err := svc.GetPolicyAssignment(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive policy assignment data", err)
		return err
	}

	if err = formatter.PrintItem(*assignment); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// PolicyAssignmentCreate subcommand function
func PolicyAssignmentCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.Parameters) && viper.IsSet(cmd.ParametersFromFile) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'parameters' or 'parameters-from-file'",
		)
	}

	assignmentIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"description":      viper.GetString(cmd.Description),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
	}
	if viper.IsSet(cmd.ParametersFromFile) {
		defIn, err := cmd.ConvertFlagParamsJsonFromFileOrStdin(viper.GetString(cmd.ParametersFromFile))
		if err != nil {
			formatter.PrintError("Cannot parse parameters", err)
			return err
		}
		assignmentIn["parameter_values"] = defIn
	}
	if viper.IsSet(cmd.Parameters) {
		params, err := cmd.FlagConvertParamsJSON(cmd.Parameters)
		if err != nil {
			formatter.PrintError("Cannot parse parameters", err)
			return err
		}
		assignmentIn["parameter_values"] = (*params)[cmd.Parameters]
	}

	assignment, err := svc.CreatePolicyAssignment(cmd.GetContext(), viper.GetString(cmd.DefinitionId), &assignmentIn)
	if err != nil {
		formatter.PrintError("Couldn't create policy assignment", err)
		return err
	}

	if err = formatter.PrintItem(*assignment); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// PolicyAssignmentUpdate subcommand function
func PolicyAssignmentUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	assignmentIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	if viper.IsSet(cmd.Description) {
		assignmentIn["description"] = viper.GetString(cmd.Description)
	}

	assignment, err := svc.UpdatePolicyAssignment(cmd.GetContext(), viper.GetString(cmd.Id), &assignmentIn)
	if err != nil {
		formatter.PrintError("Couldn't update policy assignment", err)
		return err
	}

	if err = formatter.PrintItem(*assignment); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// PolicyAssignmentDelete subcommand function
func PolicyAssignmentDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	assignment, err := svc.DeletePolicyAssignment(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete policy assignment", err)
		return err
	}

	if err = formatter.PrintItem(*assignment); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
