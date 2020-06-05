// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloudspecificextensions

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "CSE deployment Id"}

	fIdTemplate := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "CSE template Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the CSE deployment"}

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account in which is deployed"}

	fRealmId := cmd.FlagContext{Type: cmd.String, Name: cmd.RealmId, Required: true,
		Usage: "Identifier of the realm in which is deployed"}

	fParameters := cmd.FlagContext{Type: cmd.String, Name: cmd.Parameters,
		Usage: "The parameters used to configure the CSE deployment, as a json formatted parameter. \n\t" +
			"i.e: --parameters '{\"param1\":\"val1\",\"param2\":\"val2\",\"param3\":{\"id\":\"val3\"},\"param4\":true}'"}

	fParametersFromFile := cmd.FlagContext{Type: cmd.String, Name: cmd.ParametersFromFile,
		Usage: "The parameters used to configure the CSE deployment, from file or STDIN, " +
			"as a json formatted parameter. \n\t" +
			"From file: --parameters-from-file params.json \n\t" +
			"From STDIN: --parameters-from-file -"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with CSE deployment"}

	deploymentsCmd := cmd.NewCommand(cloudSpecificExtensionsCmd, &cmd.CommandContext{
		Use:   "deployments",
		Short: "Provides information about CSE deployments"},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "List CSE deployments",
		RunMethod: CloudSpecificExtensionDeploymentList},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows CSE deployment",
		RunMethod:    CloudSpecificExtensionDeploymentShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:       "deploy",
		Short:     "Deploys a new CSE deployment from CSE template",
		RunMethod: CloudSpecificExtensionDeploymentCreate,
		FlagContexts: []cmd.FlagContext{
			fIdTemplate,
			fName,
			fCloudAccountId,
			fRealmId,
			fParameters,
			fParametersFromFile,
			fLabels}},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing CSE deployment identified by the given id",
		RunMethod:    CloudSpecificExtensionDeploymentUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName}},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a CSE deployment",
		RunMethod:    CloudSpecificExtensionDeploymentDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// CloudSpecificExtensionDeploymentList subcommand function
func CloudSpecificExtensionDeploymentList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cseds, err := svc.ListCloudSpecificExtensionDeployments(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE deployments data", err)
	}
	if err = formatDeploymentsResponse(cseds, formatter); err != nil {
		return err
	}
	return nil
}

// CloudSpecificExtensionDeploymentShow subcommand function
func CloudSpecificExtensionDeploymentShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	csed, err := svc.GetCloudSpecificExtensionDeployment(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't receive CSE deployment data", err)
	}
	_, labelNamesByID := labels.LabelLoadsMapping()
	csed.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*csed); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionDeploymentCreate subcommand function
func CloudSpecificExtensionDeploymentCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.Parameters) && viper.IsSet(cmd.ParametersFromFile) {
		return fmt.Errorf(
			"invalid parameters detected. Please provide only one: 'parameters' or 'parameters-from-file'",
		)
	}

	cseDeploymentIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
		"realm_id":         viper.GetString(cmd.RealmId),
	}

	if viper.IsSet(cmd.ParametersFromFile) {
		defIn, err := cmd.ConvertFlagParamsJsonFromFileOrStdin(viper.GetString(cmd.ParametersFromFile))
		if err != nil {
			formatter.PrintFatal("Cannot parse parameters", err)
		}
		cseDeploymentIn["parameter_values"] = defIn
	}
	if viper.IsSet(cmd.Parameters) {
		params, err := cmd.FlagConvertParamsJSON(cmd.Parameters)
		if err != nil {
			formatter.PrintFatal("Cannot parse parameters", err)
		}
		cseDeploymentIn["parameter_values"] = (*params)["parameters"]
	}

	labelIDsByName, labelNamesByID := labels.LabelLoadsMapping()
	if viper.IsSet(cmd.Labels) {
		cseDeploymentIn["label_ids"] = labels.LabelResolution(
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName,
		)
	}

	cseDeployment, err := svc.CreateCloudSpecificExtensionDeployment(
		cmd.GetContext(),
		viper.GetString(cmd.Id),
		&cseDeploymentIn,
	)
	if err != nil {
		formatter.PrintFatal("Couldn't import CSE deployment", err)
	}

	cseDeployment.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseDeployment); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionDeploymentUpdate subcommand function
func CloudSpecificExtensionDeploymentUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cseDeploymentIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	cseDeployment, err := svc.UpdateCloudSpecificExtensionDeployment(
		cmd.GetContext(),
		viper.GetString(cmd.Id),
		&cseDeploymentIn,
	)
	if err != nil {
		formatter.PrintFatal("Couldn't update CSE deployment", err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	cseDeployment.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseDeployment); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// CloudSpecificExtensionDeploymentDelete subcommand function
func CloudSpecificExtensionDeploymentDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cseDeployment, err := svc.DeleteCloudSpecificExtensionDeployment(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintFatal("Couldn't delete CSE deployment", err)
	}

	_, labelNamesByID := labels.LabelLoadsMapping()
	cseDeployment.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*cseDeployment); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}

	return nil
}
