// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"
	"time"

	"github.com/ingrammicro/cio/api/cloudapplication"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	DefaultTimeLapseDeploymentStatusCheck = 30
	DefaultTimeLapseDeletionStatusCheck   = 30
)

// WireUpCloudApplicationDeployment prepares common resources to send request to Concerto API
func WireUpCloudApplicationDeployment(c *cli.Context) (
	ds *cloudapplication.CloudApplicationDeploymentService, f format.Formatter,
) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloudapplication.NewCloudApplicationDeploymentService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up CloudApplicationDeployment service", err)
	}

	return ds, f
}

// CloudApplicationDeploymentList subcommand function
func CloudApplicationDeploymentList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationDeployment(c)

	deps, err := svc.ListDeployments()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloud application deployments data", err)
	}

	if err = formatter.PrintList(deps); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// CloudApplicationDeploymentShow subcommand function
func CloudApplicationDeploymentShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationDeployment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	dep, _, err := svc.GetDeployment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloud application deployment data", err)
	}

	if err = formatter.PrintItem(*dep); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

func fillDeploymentInputs(c *cli.Context, formatter format.Formatter) map[string]interface{} {
	deploymentIn := map[string]interface{}{}
	deploymentIn["label_name"] = c.String("name")
	if c.IsSet("inputs-from-file") {
		caIn, err := utils.ConvertFlagParamsJsonFromFileOrStdin(c, c.String("inputs-from-file"))
		if err != nil {
			formatter.PrintFatal("Cannot parse input attributes", err)
		}
		deploymentIn["inputs"] = caIn
	}
	if c.IsSet("inputs") {
		params, err := utils.FlagConvertParamsJSON(c, []string{"inputs"})
		if err != nil {
			formatter.PrintFatal("Cannot parse input attributes", err)
		}
		deploymentIn["inputs"] = (*params)["inputs"]
	}
	return deploymentIn
}

// CloudApplicationDeploymentDeploy subcommand function
func CloudApplicationDeploymentDeploy(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationDeployment(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)
	if c.IsSet("inputs") && c.IsSet("inputs-from-file") {
		return fmt.Errorf("invalid parameters detected. Please provide only one: 'inputs' or 'inputs-from-file'")
	}

	deploymentIn := fillDeploymentInputs(c, formatter)

	timeLapseDeploymentStatusCheck := c.Int64("time")
	if timeLapseDeploymentStatusCheck <= 0 {
		timeLapseDeploymentStatusCheck = DefaultTimeLapseDeploymentStatusCheck
	}
	log.Debug("Time lapse -seconds- for deployment status check:", timeLapseDeploymentStatusCheck)

	deploymentTask, err := svc.CreateDeploymentTask(c.String("id"), &deploymentIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create cloud application deployment data", err)
	}

	log.Info("Task ID... ", deploymentTask.ID)
	log.Info("Deployment ID: ", deploymentTask.DeploymentID)
	log.Info("Deploying... ")
	for {
		deploymentTask, err = svc.GetDeploymentTask(c.String("id"), deploymentTask.ID)
		if err != nil {
			formatter.PrintFatal("Couldn't get cloud application deployment data", err)
		}
		log.Info("State: ", deploymentTask.State)

		if deploymentTask.State != "pending" {
			if err = formatter.PrintItem(*deploymentTask); err != nil {
				formatter.PrintFatal(PrintFormatError, err)
			}
			break
		}
		time.Sleep(time.Duration(timeLapseDeploymentStatusCheck) * time.Second)
	}
	return nil
}

// CloudApplicationDeploymentDelete subcommand function
func CloudApplicationDeploymentDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationDeployment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	timeLapseDeletionStatusCheck := c.Int64("time")
	if timeLapseDeletionStatusCheck <= 0 {
		timeLapseDeletionStatusCheck = DefaultTimeLapseDeletionStatusCheck
	}
	log.Debug("Time lapse -seconds- for deletion status check:", timeLapseDeletionStatusCheck)

	deployment, err := svc.DeleteDeployment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete cloud application deployment", err)
	}
	deploymentID := deployment.ID
	deploymentName := deployment.Name

	log.Info(fmt.Sprintf("Deployment: %s - %s undeploying...", deploymentID, deploymentName))
	for {
		deployment, status, err := svc.GetDeployment(deploymentID)
		if err != nil {
			if status == 404 {
				log.Info(fmt.Sprintf("Deployment: %s - %s undeployed.", deploymentID, deploymentName))
				break
			} else {
				formatter.PrintFatal("Couldn't check cloud application deployment data", err)
			}
		}
		log.Info("State: ", deployment.Value)

		// stops if something fails while undeploying
		if deployment.Value != "undeploying" {
			if err = formatter.PrintItem(*deployment); err != nil {
				formatter.PrintFatal(PrintFormatError, err)
			}
			break
		}
		time.Sleep(time.Duration(timeLapseDeletionStatusCheck) * time.Second)
	}
	return nil
}
