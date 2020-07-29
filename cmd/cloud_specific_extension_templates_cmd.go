package cmd

import (
	"fmt"
	"github.com/ingrammicro/cio/api/cloudspecificextension"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpCloudSpecificExtensionTemplate prepares common resources to send request to Concerto API
func WireUpCloudSpecificExtensionTemplate(c *cli.Context) (ds *cloudspecificextension.CloudSpecificExtensionTemplateService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloudspecificextension.NewCloudSpecificExtensionTemplateService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up CloudSpecificExtensionTemplate service", err)
	}

	return ds, f
}

// CloudSpecificExtensionTemplateList subcommand function
func CloudSpecificExtensionTemplateList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionTemplate(c)

	csets, err := svc.ListTemplates()
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE template data", err)
	}

	labelables := make([]types.Labelable, len(csets))
	for i := 0; i < len(csets); i++ {
		labelables[i] = types.Labelable(csets[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	csets = make([]*types.CloudSpecificExtensionTemplate, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.CloudSpecificExtensionTemplate)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.CloudSpecificExtensionTemplate, got a %T", labelable))
		}
		csets[i] = v
	}
	if err = formatter.PrintList(csets); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CloudSpecificExtensionTemplateShow subcommand function
func CloudSpecificExtensionTemplateShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cset, err := svc.GetTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE template data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	cset.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cset); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CloudSpecificExtensionTemplateImport subcommand function
func CloudSpecificExtensionTemplateImport(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionTemplate(c)

	checkRequiredFlags(c, []string{"name", "syntax"}, formatter)
	if c.IsSet("definition") && c.IsSet("definition-from-file") {
		return fmt.Errorf("invalid parameters detected. Please provide only one: 'definition' or 'definition-from-file'")
	}

	cseTemplateIn := map[string]interface{}{
		"name":   c.String("name"),
		"syntax": c.String("syntax"),
	}
	if c.IsSet("definition-from-file") {
		defIn, err := utils.ConvertFlagParamsJsonStringFromFileOrStdin(c, c.String("definition-from-file"))
		if err != nil {
			formatter.PrintFatal("Cannot parse definition", err)
		}
		cseTemplateIn["definition"] = defIn
	}
	if c.IsSet("definition") {
		cseTemplateIn["definition"] = c.String("definition")
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		cseTemplateIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	cseTemplate, err := svc.CreateTemplate(&cseTemplateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import CSE template", err)
	}

	cseTemplate.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseTemplate); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CloudSpecificExtensionTemplateUpdate subcommand function
func CloudSpecificExtensionTemplateUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionTemplate(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)
	cseTemplateIn := map[string]interface{}{
		"name": c.String("name"),
	}

	cseTemplate, err := svc.UpdateTemplate(c.String("id"), &cseTemplateIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update CSE template", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	cseTemplate.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseTemplate); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CloudSpecificExtensionTemplateListDeployments subcommand function
func CloudSpecificExtensionTemplateListDeployments(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionTemplate(c)
	checkRequiredFlags(c, []string{"id"}, formatter)

	cseds, err := svc.ListDeployments(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE deployments data", err)
	}

	labelables := make([]types.Labelable, len(cseds))
	for i := 0; i < len(cseds); i++ {
		labelables[i] = types.Labelable(cseds[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	cseds = make([]*types.CloudSpecificExtensionDeployment, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.CloudSpecificExtensionDeployment)
		if !ok {
			formatter.PrintFatal("Label filtering returned unexpected result",
				fmt.Errorf("expected labelable to be a *types.CloudSpecificExtensionDeployment, got a %T", labelable))
		}
		cseds[i] = v
	}
	if err = formatter.PrintList(cseds); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CloudSpecificExtensionTemplateDelete subcommand function
func CloudSpecificExtensionTemplateDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := svc.DeleteTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete CSE template", err)
	}
	return nil
}
