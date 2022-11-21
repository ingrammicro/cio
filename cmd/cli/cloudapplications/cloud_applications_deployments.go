// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloudapplications

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"
	"time"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Deployment Id"}

	fIdCAT := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "CAT Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the deployment"}

	fInputs := cmd.FlagContext{Type: cmd.String, Name: cmd.Inputs,
		Usage: "The inputs used to configure the cloud application deployment, as a json formatted parameter. \n\t" +
			"i.e: --inputs " +
			"'{\"region\":{\"cloud_provider\":\"Azure\",\"name\":\"US\"}," +
			"\"server_plan\":\"Standard_D2_v3\",\"admin_user\":\"admin\",\"admin_password\":\"abc$1\"}'"}

	fInputsFromFile := cmd.FlagContext{Type: cmd.String, Name: cmd.InputsFromFile,
		Usage: "The inputs used to configure the cloud application deployment, from file or STDIN, " +
			"as a json formatted parameter. \n\t" +
			"From file: --inputs-from-file attrs.json \n\t" +
			"From STDIN: --inputs-from-file -"}

	fTimeDeployment := cmd.FlagContext{Type: cmd.Int64, Name: cmd.Time,
		DefaultValue: DefaultTimeLapseDeploymentStatusCheck, Shorthand: "t",
		Usage: "Time lapse -seconds- for deployment status check"}

	fTimeDeletion := cmd.FlagContext{Type: cmd.Int64, Name: cmd.Time, DefaultValue: DefaultTimeLapseDeletionStatusCheck,
		Shorthand: "t", Usage: "Time lapse -seconds- for deletion status check"}

	deploymentsCmd := cmd.NewCommand(cloudApplicationsCmd, &cmd.CommandContext{
		Use:   "deployments",
		Short: "Provides information about CAT deployments"},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists deployments",
		RunMethod: CloudApplicationDeploymentList},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows deployment",
		RunMethod:    CloudApplicationDeploymentShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:          "deploy",
		Short:        "Deploys a CAT",
		RunMethod:    CloudApplicationDeploymentDeploy,
		FlagContexts: []cmd.FlagContext{fIdCAT, fName, fInputs, fInputsFromFile, fTimeDeployment}},
	)
	cmd.NewCommand(deploymentsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a deployment",
		RunMethod:    CloudApplicationDeploymentDelete,
		FlagContexts: []cmd.FlagContext{fId, fTimeDeletion}},
	)
}

const (
	DefaultTimeLapseDeploymentStatusCheck = 30
	DefaultTimeLapseDeletionStatusCheck   = 30
)

// CloudApplicationDeploymentList subcommand function
func CloudApplicationDeploymentList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	deps, err := svc.ListCloudApplicationDeployments(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive cloud application deployments data", err)
		return err
	}

	if err = formatter.PrintList(deps); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// CloudApplicationDeploymentShow subcommand function
func CloudApplicationDeploymentShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	dep, _, err := svc.GetCloudApplicationDeployment(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive cloud application deployment data", err)
		return err
	}

	if err = formatter.PrintItem(*dep); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

func fillDeploymentInputs(formatter format.Formatter) (map[string]interface{}, error) {
	deploymentIn := map[string]interface{}{}
	deploymentIn["label_name"] = viper.GetString(cmd.Name)
	if viper.IsSet(cmd.InputsFromFile) {
		caIn, err := cmd.ConvertFlagParamsJsonFromFileOrStdin(viper.GetString(cmd.InputsFromFile))
		if err != nil {
			formatter.PrintError("Cannot parse input attributes from file", err)
			return nil, err
		}
		deploymentIn["inputs"] = caIn
	}
	if viper.IsSet(cmd.Inputs) {
		params, err := cmd.FlagConvertParamsJSON(cmd.Inputs)
		if err != nil {
			formatter.PrintError("Cannot parse input attributes", err)
			return nil, err
		}
		deploymentIn["inputs"] = (*params)["inputs"]
	}
	return deploymentIn, nil
}

// CloudApplicationDeploymentDeploy subcommand function
func CloudApplicationDeploymentDeploy() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	if viper.IsSet(cmd.Inputs) && viper.IsSet(cmd.InputsFromFile) {
		return fmt.Errorf("invalid parameters detected. Please provide only one: 'inputs' or 'inputs-from-file'")
	}

	deploymentIn, err := fillDeploymentInputs(formatter)
	if err != nil {
		return err
	}

	timeLapseDeploymentStatusCheck := viper.GetInt64(cmd.Time)
	if timeLapseDeploymentStatusCheck <= 0 {
		timeLapseDeploymentStatusCheck = DefaultTimeLapseDeploymentStatusCheck
	}
	log.Debug("Time lapse -seconds- for deployment status check:", timeLapseDeploymentStatusCheck)

	ctx := cmd.GetContext()
	deploymentTask, err := svc.CreateCloudApplicationDeploymentTask(
		ctx,
		viper.GetString(cmd.Id),
		&deploymentIn,
	)
	if err != nil {
		formatter.PrintError("Couldn't create cloud application deployment data", err)
		return err
	}

	log.Info("Task ID... ", deploymentTask.ID)
	log.Info("Deployment ID: ", deploymentTask.DeploymentID)
	log.Info("Deploying... ")
	for {
		deploymentTask, err = svc.GetCloudApplicationDeploymentTask(
			ctx,
			viper.GetString(cmd.Id),
			deploymentTask.ID,
		)
		if err != nil {
			formatter.PrintError("Couldn't get cloud application deployment data", err)
			return err
		}
		log.Info("State: ", deploymentTask.State)

		if deploymentTask.State != "pending" {
			if err = formatter.PrintItem(*deploymentTask); err != nil {
				formatter.PrintError(cmd.PrintFormatError, err)
				return err
			}
			break
		}
		time.Sleep(time.Duration(timeLapseDeploymentStatusCheck) * time.Second)
	}
	return nil
}

// CloudApplicationDeploymentDelete subcommand function
func CloudApplicationDeploymentDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	timeLapseDeletionStatusCheck := viper.GetInt64(cmd.Time)
	if timeLapseDeletionStatusCheck <= 0 {
		timeLapseDeletionStatusCheck = DefaultTimeLapseDeletionStatusCheck
	}
	log.Debug("Time lapse -seconds- for deletion status check:", timeLapseDeletionStatusCheck)

	ctx := cmd.GetContext()
	deployment, err := svc.DeleteCloudApplicationDeployment(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete cloud application deployment", err)
		return err
	}
	deploymentID := deployment.ID
	deploymentName := deployment.Name

	log.Info(fmt.Sprintf("Deployment: %s - %s undeploying...", deploymentID, deploymentName))
	for {
		deployment, status, err := svc.GetCloudApplicationDeployment(ctx, deploymentID)
		if err != nil {
			if status == 404 {
				log.Info(fmt.Sprintf("Deployment: %s - %s undeployed.", deploymentID, deploymentName))
				break
			} else {
				formatter.PrintError("Couldn't check cloud application deployment data", err)
				return err
			}
		}
		log.Info("State: ", deployment.Value)

		// stops if something fails while undeploying
		if deployment.Value != "undeploying" {
			if err = formatter.PrintItem(*deployment); err != nil {
				formatter.PrintError(cmd.PrintFormatError, err)
				return err
			}
			break
		}
		time.Sleep(time.Duration(timeLapseDeletionStatusCheck) * time.Second)
	}
	return nil
}
