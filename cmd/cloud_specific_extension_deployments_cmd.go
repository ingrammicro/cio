// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"

	"github.com/ingrammicro/cio/api/cloudspecificextension"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpCloudSpecificExtensionDeployment prepares common resources to send request to Concerto API
func WireUpCloudSpecificExtensionDeployment(
	c *cli.Context,
) (ds *cloudspecificextension.CloudSpecificExtensionDeploymentService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloudspecificextension.NewCloudSpecificExtensionDeploymentService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up CloudSpecificExtensionDeployment service", err)
	}

	return ds, f
}

// CloudSpecificExtensionDeploymentList subcommand function
func CloudSpecificExtensionDeploymentList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionDeployment(c)

	cseds, err := svc.ListDeployments()
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE deployment data", err)
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
			formatter.PrintFatal(LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.CloudSpecificExtensionDeployment, got a %T", labelable))
		}
		cseds[i] = v
	}
	if err = formatter.PrintList(cseds); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionDeploymentShow subcommand function
func CloudSpecificExtensionDeploymentShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionDeployment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	csed, err := svc.GetDeployment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE deployment data", err)
	}
	_, labelNamesByID := LabelLoadsMapping(c)
	csed.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*csed); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionDeploymentCreate subcommand function
func CloudSpecificExtensionDeploymentCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionDeployment(c)

	checkRequiredFlags(c, []string{"id", "name", "cloud-account-id", "realm-id"}, formatter)
	if c.IsSet("parameters") && c.IsSet("parameters-from-file") {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'parameters' or 'parameters-from-file'",
		)
	}

	cseDeploymentIn := map[string]interface{}{
		"name":             c.String("name"),
		"cloud_account_id": c.String("cloud-account-id"),
		"realm_id":         c.String("realm-id"),
	}

	if c.IsSet("parameters-from-file") {
		defIn, err := utils.ConvertFlagParamsJsonFromFileOrStdin(c, c.String("parameters-from-file"))
		if err != nil {
			formatter.PrintFatal("Cannot parse parameters", err)
		}
		cseDeploymentIn["parameter_values"] = defIn
	}
	if c.IsSet("parameters") {
		params, err := utils.FlagConvertParamsJSON(c, []string{"parameters"})
		if err != nil {
			formatter.PrintFatal("Cannot parse parameters", err)
		}
		cseDeploymentIn["parameter_values"] = (*params)["parameters"]
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	if c.IsSet("labels") {
		cseDeploymentIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	cseDeployment, err := svc.CreateDeployment(c.String("id"), &cseDeploymentIn)
	if err != nil {
		formatter.PrintFatal("Couldn't import CSE deployment", err)
	}

	cseDeployment.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseDeployment); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionDeploymentUpdate subcommand function
func CloudSpecificExtensionDeploymentUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionDeployment(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)
	cseDeploymentIn := map[string]interface{}{
		"name": c.String("name"),
	}

	cseDeployment, err := svc.UpdateDeployment(c.String("id"), &cseDeploymentIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update CSE deployment", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	cseDeployment.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseDeployment); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionDeploymentDelete subcommand function
func CloudSpecificExtensionDeploymentDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudSpecificExtensionDeployment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cseDeployment, err := svc.DeleteDeployment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete CSE deployment", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	cseDeployment.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseDeployment); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}

	return nil
}
