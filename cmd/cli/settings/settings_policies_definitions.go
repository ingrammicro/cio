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
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Policy Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the policy definition"}

	fDescription := cmd.FlagContext{Type: cmd.String, Name: cmd.Description,
		Usage: "Description of the policy definition"}
	fDescriptionReq := fDescription
	fDescriptionReq.Required = true

	fDefinition := cmd.FlagContext{Type: cmd.String, Name: cmd.Definition,
		Usage: "The definition used to configure the policy, as a json formatted parameter. \n\t" +
			"i.e: --definition '{\"parameters\": {\"prefix\": {\"type\": \"string\"," +
			"\"metadata\": {\"description\": \"prefix data\"}},\"suffix\": {\"type\": \"string\"," +
			"\"metadata\": {\"description\": \"suffix data\"}}},\"policyRule\": {\"if\": {\"not\": {\"field\": \"name\"," +
			"\"like\": \"[concat(parameters('prefix'), '*', parameters('suffix'))]\"}},\"then\": {\"effect\": \"audit\"}}}'"}

	fDefinitionFromFile := cmd.FlagContext{Type: cmd.String, Name: cmd.DefinitionFromFile,
		Usage: "The definition used to configure the policy, from file or STDIN, as a json formatted parameter. \n\t" +
			"From file: --definition-from-file def.json \n\t" +
			"From STDIN: --definition-from-file -"}

	definitionsCmd := cmd.NewCommand(policiesCmd, &cmd.CommandContext{
		Use:   "definitions",
		Short: "Provides information about policy definitions"},
	)
	cmd.NewCommand(definitionsCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists policy definitions",
		RunMethod: PolicyDefinitionList},
	)
	cmd.NewCommand(definitionsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows policy definition",
		RunMethod:    PolicyDefinitionShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(definitionsCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates policy definition",
		RunMethod:    PolicyDefinitionCreate,
		FlagContexts: []cmd.FlagContext{fName, fDescriptionReq, fDefinition, fDefinitionFromFile}},
	)
	cmd.NewCommand(definitionsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing policy definition identified by the given id",
		RunMethod:    PolicyDefinitionUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName, fDescription}},
	)
	cmd.NewCommand(definitionsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a policy definition",
		RunMethod:    PolicyDefinitionDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(definitionsCmd, &cmd.CommandContext{
		Use:          "list-assignments",
		Short:        "lists policy assignments",
		RunMethod:    PolicyDefinitionListAssignments,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// PolicyDefinitionList subcommand function
func PolicyDefinitionList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	definitions, err := svc.ListPolicyDefinitions(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive policy definitions data", err)
		return err
	}

	if err = formatter.PrintList(definitions); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// PolicyDefinitionShow subcommand function
func PolicyDefinitionShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	definition, err := svc.GetPolicyDefinition(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive policy definition data", err)
		return err
	}

	if err = formatter.PrintItem(*definition); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// PolicyDefinitionCreate subcommand function
func PolicyDefinitionCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.Definition) && viper.IsSet(cmd.DefinitionFromFile) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'definition' or 'definition-from-file'",
		)
	}

	definitionIn := map[string]interface{}{
		"name":        viper.GetString(cmd.Name),
		"description": viper.GetString(cmd.Description),
	}
	if viper.IsSet(cmd.DefinitionFromFile) {
		defIn, err := cmd.ConvertFlagParamsJsonStringFromFileOrStdin(viper.GetString(cmd.DefinitionFromFile))
		if err != nil {
			formatter.PrintError("Cannot parse definition", err)
			return err
		}
		definitionIn["definition"] = defIn
	}
	if viper.IsSet(cmd.Definition) {
		definitionIn["definition"] = viper.GetString(cmd.Definition)
	}

	definition, err := svc.CreatePolicyDefinition(cmd.GetContext(), &definitionIn)
	if err != nil {
		formatter.PrintError("Couldn't create policy definition", err)
		return err
	}

	if err = formatter.PrintItem(*definition); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// PolicyDefinitionUpdate subcommand function
func PolicyDefinitionUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	definitionIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}
	if viper.IsSet(cmd.Description) {
		definitionIn["description"] = viper.GetString(cmd.Description)
	}

	definition, err := svc.UpdatePolicyDefinition(cmd.GetContext(), viper.GetString(cmd.Id), &definitionIn)
	if err != nil {
		formatter.PrintError("Couldn't update policy definition", err)
		return err
	}

	if err = formatter.PrintItem(*definition); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// PolicyDefinitionDelete subcommand function
func PolicyDefinitionDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeletePolicyDefinition(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete policy definition", err)
		return err
	}
	return nil
}

// PolicyDefinitionListAssignments subcommand function
func PolicyDefinitionListAssignments() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	assignments, err := svc.ListPolicyDefinitionAssignments(cmd.GetContext(), viper.GetString(cmd.Id))
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
