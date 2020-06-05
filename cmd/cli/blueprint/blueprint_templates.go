// Copyright (c) 2017-2022 Ingram Micro Inc.

package blueprint

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"
	"regexp"
	"strings"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/spf13/viper"
)

func init() {
	fLabelsFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label as a query filter"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Template Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Name of the template"}
	fNameReq := fName
	fNameReq.Required = true

	fGenericImageId := cmd.FlagContext{Type: cmd.String, Name: cmd.GenericImageId, Required: true,
		Usage: "Identifier of the OS image that the template builds on"}

	fRunList := cmd.FlagContext{Type: cmd.String, Name: cmd.RunList,
		Usage: "A list of comma separated cookbook recipes that is run on the servers at start-up, " +
			"i.e: --run-list imco::client,1password,joomla"}

	fCookbookVersions := cmd.FlagContext{Type: cmd.String, Name: cmd.CookbookVersions,
		Usage: "The cookbook versions used to configure the service recipes in the run-list, " +
			"i.e: --cookbook-versions \"imco:3.0.3,1password~>1.3.0,joomla:0.11.0\" \n\t" +
			"Cookbook version format: [NAME<OPERATOR>VERSION] \n\t" +
			"Supported Operators:\n\t\t" +
			"Chef supermarket cookbook '~>','=','>=','>','<','<='\n\t\tUploaded cookbook ':'"}

	fConfigurationAttributes := cmd.FlagContext{Type: cmd.String, Name: cmd.ConfigurationAttributes,
		Usage: "The attributes used to configure the service recipes in the run-list, as a json formatted parameter. " +
			"i.e: --configuration-attributes '{\"joomla\":{\"db\":{\"password\":\"my_pass\"},\"port\":\"8080\"}}'"}

	fConfigurationAttributesFromFile := cmd.FlagContext{Type: cmd.String, Name: cmd.ConfigurationAttributesFromFile,
		Usage: "The attributes used to configure the service recipes in the run-list, from file or STDIN, " +
			"as a json formatted parameter. \n\t" +
			"From file: --configuration-attributes-from-file attrs.json \n\t" +
			"From STDIN: --configuration-attributes-from-file -"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with template"}

	fTypeFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Type, Required: true,
		Usage: "Shows the script characterisations of a template: \"operational\", \"boot\" or \"shutdown\""}

	fTemplateId := cmd.FlagContext{Type: cmd.String, Name: cmd.TemplateId, Required: true, Usage: "Template Id"}

	fIdScript := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Script Id"}

	fType := cmd.FlagContext{Type: cmd.String, Name: cmd.Type, Required: true,
		Usage: "Must be \"operational\", \"boot\" or \"shutdown\""}

	fScriptId := cmd.FlagContext{Type: cmd.String, Name: cmd.ScriptId, Required: true,
		Usage: "Identifier for the script that is parameterised by the script characterisation"}

	fParameterValues := cmd.FlagContext{Type: cmd.String, Name: cmd.ParameterValues,
		Usage: "A map that assigns a value to each script parameter, as a json formatted parameter; " +
			"i.e: '{\"param1\":\"val1\",\"param2\":\"val2\"}'"}

	fParameterValuesFromFile := cmd.FlagContext{Type: cmd.String, Name: cmd.ParameterValuesFromFile,
		Usage: "A map that assigns a value to each script parameter, from file or STDIN, " +
			"as a json formatted parameter. \n\t" +
			"From file: --parameter-values-from-file params.json \n\t" +
			"From STDIN: --parameter-values-from-file -"}

	fIdTemplateScript := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true,
		Usage: "Identifier for the template-script that is parameterised by the script characterisation"}

	fScriptIds := cmd.FlagContext{Type: cmd.String, Name: cmd.ScriptIds, Required: true,
		Usage: "A list of comma separated scripts ids that must contain all the ids of scripts " +
			"of the given template and type in the desired execution order"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, DefaultValue: "template", Hidden: true,
		Usage: "Resource Type"}

	templatesCmd := cmd.NewCommand(blueprintCmd, &cmd.CommandContext{
		Use:   "templates",
		Short: "Provides information on templates"},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all available templates",
		RunMethod:    TemplateList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about a specific template",
		RunMethod:    TemplateShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:       "create",
		Short:     "Creates a new template",
		RunMethod: TemplateCreate,
		FlagContexts: []cmd.FlagContext{
			fNameReq,
			fGenericImageId,
			fRunList,
			fCookbookVersions,
			fConfigurationAttributes,
			fConfigurationAttributesFromFile,
			fLabels,
		}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:       "update",
		Short:     "Updates an existing template",
		RunMethod: TemplateUpdate,
		FlagContexts: []cmd.FlagContext{
			fId,
			fName,
			fRunList,
			fCookbookVersions,
			fConfigurationAttributes,
			fConfigurationAttributesFromFile,
		}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "compile",
		Short:        "Compiles an existing template",
		RunMethod:    TemplateCompile,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a template",
		RunMethod:    TemplateDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "list-template-scripts",
		Short:        "Shows the script characterisations of a template",
		RunMethod:    TemplateScriptList,
		FlagContexts: []cmd.FlagContext{fTemplateId, fTypeFilter}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "show-template-script",
		Short:        "Shows information about a specific script characterisation",
		RunMethod:    TemplateScriptShow,
		FlagContexts: []cmd.FlagContext{fTemplateId, fIdScript}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use: "create-template-script",
		Short: "Creates a new script characterisation for a template and appends it " +
			"to the list of script characterisations of the same type",
		RunMethod:    TemplateScriptCreate,
		FlagContexts: []cmd.FlagContext{fTemplateId, fType, fScriptId, fParameterValues, fParameterValuesFromFile}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "update-template-script",
		Short:        "Updates an existing script characterisation for a template",
		RunMethod:    TemplateScriptUpdate,
		FlagContexts: []cmd.FlagContext{fTemplateId, fIdTemplateScript, fParameterValues, fParameterValuesFromFile}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use: "reorder-template-scripts",
		Short: "Reorders the scripts of the template and type specified according to the provided order, " +
			"changing their execution order as corresponds",
		RunMethod:    TemplateScriptReorder,
		FlagContexts: []cmd.FlagContext{fTemplateId, fType, fScriptIds}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "delete-template-script",
		Short:        "Removes a parametrized script from a template",
		RunMethod:    TemplateScriptDelete,
		FlagContexts: []cmd.FlagContext{fTemplateId, fIdTemplateScript}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "list-template-servers",
		Short:        "Returns information about the servers that use a specific template",
		RunMethod:    TemplateServersList,
		FlagContexts: []cmd.FlagContext{fTemplateId}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(templatesCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

var templateCookbookVersionValueRegexp = regexp.MustCompile(`^([a-zA-Z0-9_-]+)(~>|=|>=|<=|>|<|:)(\d+(?:\.\d+){0,2})$`)

const CannotResolveCookbookVersionsData = "cannot resolve cookbook versions data"
const CannotParseInputParameterValues = "Cannot parse input parameter values"

// TemplateList subcommand function
func TemplateList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	templates, err := svc.ListTemplates(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive template data", err)
	}

	labelables := make([]types.Labelable, len(templates))
	for i := 0; i < len(templates); i++ {
		labelables[i] = types.Labelable(templates[i])
	}
	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()
	filteredLabelables := labels.LabelFiltering(labelables, labelIDsByName)
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	templates = make([]*types.Template, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		tpl, ok := labelable.(*types.Template)
		if !ok {
			formatter.PrintFatal(cmd.LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.Template, got a %T", labelable))
		}
		templates[i] = tpl
	}

	if err = formatter.PrintList(templates); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// TemplateShow subcommand function
func TemplateShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	template, err := svc.GetTemplate(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive template data", err)
	}

	if err = resolveCookbookVersions(template); err != nil {
		formatter.PrintFatal(CannotResolveCookbookVersionsData, err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	template.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

func setTemplateParams(formatter format.Formatter, templateIn map[string]interface{}) {
	logger.DebugFuncInfo()
	if viper.IsSet(cmd.ConfigurationAttributesFromFile) {
		caIn, err := cmd.ConvertFlagParamsJsonFromFileOrStdin(viper.GetString(cmd.ConfigurationAttributesFromFile))
		if err != nil {
			formatter.PrintFatal("Cannot parse input configuration attributes", err)
		}
		templateIn["configuration_attributes"] = caIn
	}
	if viper.IsSet(cmd.ConfigurationAttributes) {
		params, err := cmd.FlagConvertParamsJSON(cmd.ConfigurationAttributes)
		if err != nil {
			formatter.PrintFatal("Cannot parse input configuration attributes", err)
		}
		templateIn["configuration_attributes"] = (*params)[cmd.ConfigurationAttributes]
	}
	if viper.IsSet(cmd.RunList) {
		templateIn["run_list"] = utils.RemoveDuplicates(strings.Split(viper.GetString(cmd.RunList), ","))
	}
	if viper.IsSet(cmd.CookbookVersions) {
		cbIn, err := convertFlagParamsToCookbookVersions(viper.GetString(cmd.CookbookVersions))
		if err != nil {
			formatter.PrintFatal("Cannot parse input cookbook versions", err)
		}
		templateIn["cookbook_versions"] = cbIn
	}
}

// TemplateCreate subcommand function
func TemplateCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.ConfigurationAttributes) && viper.IsSet(cmd.ConfigurationAttributesFromFile) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: " +
				"'configuration-attributes' or 'configuration-attributes-from-file'",
		)
	}

	templateIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"generic_image_id": viper.GetString(cmd.GenericImageId),
	}
	setTemplateParams(formatter, templateIn)

	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()

	if viper.IsSet("labels") {
		templateIn["label_ids"] = labels.LabelResolution(viper.GetString(cmd.Labels), &labelNamesByID, &labelIDsByName)
	}

	template, err := svc.CreateTemplate(cmd.GetContext(), &templateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create template", err)
	}

	if err = resolveCookbookVersions(template); err != nil {
		formatter.PrintFatal(CannotResolveCookbookVersionsData, err)
	}

	template.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// TemplateUpdate subcommand function
func TemplateUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.ConfigurationAttributes) && viper.IsSet(cmd.ConfigurationAttributesFromFile) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: " +
				"'configuration-attributes' or 'configuration-attributes-from-file'",
		)
	}

	templateIn := map[string]interface{}{}
	cmd.SetParamString("name", cmd.Name, templateIn)
	setTemplateParams(formatter, templateIn)

	template, err := svc.UpdateTemplate(cmd.GetContext(), viper.GetString(cmd.Id), &templateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update template", err)
	}

	if err = resolveCookbookVersions(template); err != nil {
		formatter.PrintFatal(CannotResolveCookbookVersionsData, err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	template.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// TemplateCompile subcommand function
func TemplateCompile() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	templateIn := map[string]interface{}{}
	template, err := svc.CompileTemplate(cmd.GetContext(), viper.GetString(cmd.Id), &templateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't compile template", err)
	}

	if err = resolveCookbookVersions(template); err != nil {
		formatter.PrintFatal(CannotResolveCookbookVersionsData, err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	template.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// TemplateDelete subcommand function
func TemplateDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteTemplate(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't delete template", err)
	}
	return nil
}

// =========== Template Scripts =============

// TemplateScriptList subcommand function
func TemplateScriptList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	templateScripts, err := svc.ListTemplateScripts(
		cmd.GetContext(),
		viper.GetString(cmd.TemplateId),
		viper.GetString(cmd.Type),
	)
	if err != nil {
		formatter.PrintFatal("Couldn't receive templateScript data", err)
	}
	if err = formatter.PrintList(templateScripts); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// TemplateScriptShow subcommand function
func TemplateScriptShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	templateScript, err := svc.GetTemplateScript(
		cmd.GetContext(),
		viper.GetString(cmd.TemplateId),
		viper.GetString(cmd.Id),
	)
	if err != nil {
		formatter.PrintFatal("Couldn't receive templateScript data", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// TemplateScriptCreate subcommand function
func TemplateScriptCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.ParameterValues) && viper.IsSet(cmd.ParameterValuesFromFile) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'parameter-values' or 'parameter-values-from-file'",
		)
	}

	templateScriptIn := map[string]interface{}{
		"type":      viper.GetString(cmd.Type),
		"script_id": viper.GetString(cmd.ScriptId),
	}
	if viper.IsSet(cmd.ParameterValuesFromFile) {
		pvIn, err := cmd.ConvertFlagParamsJsonFromFileOrStdin(viper.GetString(cmd.ParameterValuesFromFile))
		if err != nil {
			formatter.PrintFatal(CannotParseInputParameterValues, err)
		}
		templateScriptIn["parameter_values"] = pvIn
	}
	if viper.IsSet(cmd.ParameterValues) {
		params, err := cmd.FlagConvertParamsJSON(cmd.ParameterValues)
		if err != nil {
			formatter.PrintFatal(CannotParseInputParameterValues, err)
		}
		templateScriptIn["parameter_values"] = (*params)[cmd.ParameterValues]
	}

	templateScript, err := svc.CreateTemplateScript(
		cmd.GetContext(),
		viper.GetString(cmd.TemplateId),
		&templateScriptIn,
	)
	if err != nil {
		formatter.PrintFatal("Couldn't create templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// TemplateScriptUpdate subcommand function
func TemplateScriptUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.ParameterValues) && viper.IsSet(cmd.ParameterValuesFromFile) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'parameter-values' or 'parameter-values-from-file'",
		)
	}

	templateScriptIn := map[string]interface{}{}

	if viper.IsSet(cmd.ParameterValuesFromFile) {
		pvIn, err := cmd.ConvertFlagParamsJsonFromFileOrStdin(viper.GetString(cmd.ParameterValuesFromFile))
		if err != nil {
			formatter.PrintFatal(CannotParseInputParameterValues, err)
		}
		templateScriptIn["parameter_values"] = pvIn
	}
	if viper.IsSet(cmd.ParameterValues) {
		params, err := cmd.FlagConvertParamsJSON(cmd.ParameterValues)
		if err != nil {
			formatter.PrintFatal(CannotParseInputParameterValues, err)
		}
		templateScriptIn["parameter_values"] = (*params)[cmd.ParameterValues]
	}

	templateScript, err := svc.UpdateTemplateScript(cmd.GetContext(),
		viper.GetString(cmd.TemplateId),
		viper.GetString(cmd.Id),
		&templateScriptIn,
	)
	if err != nil {
		formatter.PrintFatal("Couldn't update templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// TemplateScriptDelete subcommand function
func TemplateScriptDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteTemplateScript(cmd.GetContext(), viper.GetString(cmd.TemplateId), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't delete templateScript", err)
	}
	return nil
}

// TemplateScriptReorder subcommand function
func TemplateScriptReorder() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	templateScriptIn := map[string]interface{}{
		"type":       viper.GetString(cmd.Type),
		"script_ids": utils.RemoveDuplicates(strings.Split(viper.GetString(cmd.ScriptIds), ",")),
	}

	templateScript, err := svc.ReorderTemplateScript(
		cmd.GetContext(),
		viper.GetString(cmd.TemplateId),
		&templateScriptIn,
	)
	if err != nil {
		formatter.PrintFatal("Couldn't reorder templateScript", err)
	}
	if err = formatter.PrintList(templateScript); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// =========== Template Servers =============

// TemplateServersList subcommand function
func TemplateServersList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	templateServers, err := svc.ListTemplateServers(cmd.GetContext(), viper.GetString(cmd.TemplateId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive template servers data", err)
	}
	if err = formatter.PrintList(templateServers); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// =========== Template helpers =============

func processCookbookVersionItem(
	cbvIn string,
	cookbookVersions []*types.CookbookVersion,
	name, operator, version string,
	result map[string]interface{},
) error {
	// uploaded. It requires to map adequate version_id
	if operator == ":" {
		if len(cookbookVersions) == 0 {
			// data is loaded only once
			svc, _, formatter := cli.WireUpAPIClient()
			cbvs, err := svc.ListCookbookVersions(cmd.GetContext())
			if err != nil {
				formatter.PrintFatal("cannot receive uploaded cookbook versions data", err)
			}
			cookbookVersions = cbvs
		}
		for _, cbv := range cookbookVersions {
			if name == cbv.Name && version == cbv.Version {
				result[name] = map[string]interface{}{"version_id": cbv.ID}
			}
		}
		// provided cookbook version does not match the available uploaded
		if _, found := result[name]; !found {
			return fmt.Errorf(
				"invalid cookbook version: %s does not match any of the cookbook versions uploaded to the platform",
				cbvIn,
			)
		}
	} else {
		//supermarket
		// at any case, it should leave a space between operator and version
		result[name] = map[string]interface{}{"version": operator + " " + version}
	}
	return nil
}

// convertFlagParamsToCookbookVersions returns the json representation for the given friendly input format f cookbook
//versions assignation
// i.e: "wordpress:0.1.0,nano=2.0.1,1password~>1.3.0"
func convertFlagParamsToCookbookVersions(cbvsIn string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	cookbookVersions := make([]*types.CookbookVersion, 0)
	for _, cbvIn := range strings.Split(cbvsIn, ",") {
		values := templateCookbookVersionValueRegexp.FindStringSubmatch(cbvIn)
		if len(values) == 0 {
			return nil, fmt.Errorf("invalid input cookbook version format %s", cbvIn)
		}
		name, operator, version := values[1], values[2], values[3]
		if _, found := result[name]; found {
			return nil, fmt.Errorf("detected duplicated cookbook version name: %s", name)
		}
		if err := processCookbookVersionItem(cbvIn, cookbookVersions, name, operator, version, result); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// resolveCookbookVersions resolves adequate cookbook version ids
func resolveCookbookVersions(template *types.Template) error {
	svc, _, _ := cli.WireUpAPIClient()
	cbvs, err := svc.ListCookbookVersions(cmd.GetContext())
	if err != nil {
		return err
	}

	customCookbookVersionsByVersionID := make(map[string]string)
	for _, cbv := range cbvs {
		customCookbookVersionsByVersionID[cbv.ID] = cbv.Version
	}
	template.FillInCookbookVersionComposite(customCookbookVersionsByVersionID)

	return nil
}
