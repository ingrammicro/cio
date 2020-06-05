// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloudspecificextensions

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/utils/format"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "CSE template Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the CSE template"}

	fSyntax := cmd.FlagContext{Type: cmd.String, Name: cmd.Syntax, Required: true,
		Usage: "Cloud provider syntax of the CSE template"}

	fDefinition := cmd.FlagContext{Type: cmd.String, Name: cmd.Definition,
		Usage: "The definition used to configure the CSE template, as a json formatted parameter. \n\t" +
			"i.e: --definition " +
			"'{\"$schema\":\"https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#\"," +
			"\"contentVersion\":\"1.0.0.0\",\"parameters\":{\"vmName\":{\"type\":\"string\"," +
			"\"defaultValue\": \"simpleLinuxVM\",\"metadata\":{\"description\": \"The name of you Virtual Machine.\"}}}}'"}

	fDefinitionFromFile := cmd.FlagContext{Type: cmd.String, Name: cmd.DefinitionFromFile,
		Usage: "The definition used to configure the CSE template, from file or STDIN, as a json formatted parameter. \n\t" +
			"From file: --definition-from-file def.json \n\t" +
			"From STDIN: --definition-from-file -"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with CSE template"}

	templatesCmd := cmd.NewCommand(cloudSpecificExtensionsCmd, &cmd.CommandContext{
		Use:   "templates",
		Short: "Provides information about CSE templates"},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "List CSE templates",
		RunMethod: CloudSpecificExtensionTemplateList},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows CSE template",
		RunMethod:    CloudSpecificExtensionTemplateShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "import",
		Short:        "Imports a CSE template",
		RunMethod:    CloudSpecificExtensionTemplateImport,
		FlagContexts: []cmd.FlagContext{fName, fSyntax, fDefinition, fDefinitionFromFile, fLabels}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing CSE template identified by the given id",
		RunMethod:    CloudSpecificExtensionTemplateUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "list-deployments",
		Short:        "List CSE deployments of a CSE template",
		RunMethod:    CloudSpecificExtensionTemplateListDeployments,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a CSE template",
		RunMethod:    CloudSpecificExtensionTemplateDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// CloudSpecificExtensionTemplateList subcommand function
func CloudSpecificExtensionTemplateList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	csets, err := svc.ListCloudSpecificExtensionTemplates(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE template data", err)
	}

	labelables := make([]types.Labelable, len(csets))
	for i := 0; i < len(csets); i++ {
		labelables[i] = types.Labelable(csets[i])
	}
	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()
	filteredLabelables := labels.LabelFiltering(labelables, labelIDsByName)
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	csets = make([]*types.CloudSpecificExtensionTemplate, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.CloudSpecificExtensionTemplate)
		if !ok {
			formatter.PrintFatal(cmd.LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.CloudSpecificExtensionTemplate, got a %T", labelable))
		}
		csets[i] = v
	}
	if err = formatter.PrintList(csets); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionTemplateShow subcommand function
func CloudSpecificExtensionTemplateShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cset, err := svc.GetCloudSpecificExtensionTemplate(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE template data", err)
	}
	_, labelNamesByID := labels.LabelLoadsMapping()
	cset.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cset); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionTemplateImport subcommand function
func CloudSpecificExtensionTemplateImport() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.Definition) && viper.IsSet(cmd.DefinitionFromFile) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'definition' or 'definition-from-file'",
		)
	}

	cseTemplateIn := map[string]interface{}{
		"name":   viper.GetString(cmd.Name),
		"syntax": viper.GetString(cmd.Syntax),
	}
	if viper.IsSet(cmd.DefinitionFromFile) {
		defIn, err := cmd.ConvertFlagParamsJsonStringFromFileOrStdin(viper.GetString(cmd.DefinitionFromFile))
		if err != nil {
			formatter.PrintFatal("Cannot parse definition", err)
		}
		cseTemplateIn["definition"] = defIn
	}
	if viper.IsSet(cmd.Definition) {
		cseTemplateIn["definition"] = viper.GetString(cmd.Definition)
	}

	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()
	if viper.IsSet(cmd.Labels) {
		cseTemplateIn["label_ids"] = labels.LabelResolution(
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName,
		)
	}

	cseTemplate, err := svc.CreateCloudSpecificExtensionTemplate(cmd.GetContext(), &cseTemplateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import CSE template", err)
	}

	cseTemplate.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseTemplate); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionTemplateUpdate subcommand function
func CloudSpecificExtensionTemplateUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cseTemplateIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	cseTemplate, err := svc.UpdateCloudSpecificExtensionTemplate(
		cmd.GetContext(),
		viper.GetString(cmd.Id),
		&cseTemplateIn,
	)
	if err != nil {
		formatter.PrintFatal("Couldn't update CSE template", err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	cseTemplate.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseTemplate); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionTemplateListDeployments subcommand function
func CloudSpecificExtensionTemplateListDeployments() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cseds, err := svc.ListCloudSpecificExtensionTemplateDeployments(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE template deployments data", err)
	}
	if err = formatDeploymentsResponse(cseds, formatter); err != nil {
		return err
	}
	return nil
}

func formatDeploymentsResponse(cseds []*types.CloudSpecificExtensionDeployment, formatter format.Formatter) error {
	labelables := make([]types.Labelable, len(cseds))
	for i := 0; i < len(cseds); i++ {
		labelables[i] = types.Labelable(cseds[i])
	}
	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()
	filteredLabelables := labels.LabelFiltering(labelables, labelIDsByName)
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	cseds = make([]*types.CloudSpecificExtensionDeployment, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.CloudSpecificExtensionDeployment)
		if !ok {
			formatter.PrintFatal(cmd.LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.CloudSpecificExtensionDeployment, got a %T", labelable))
		}
		cseds[i] = v
	}
	if err := formatter.PrintList(cseds); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionTemplateDelete subcommand function
func CloudSpecificExtensionTemplateDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteCloudSpecificExtensionTemplate(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't delete CSE template", err)
	}
	return nil
}
