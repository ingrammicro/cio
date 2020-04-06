package cmd

import (
	"fmt"
	"github.com/ingrammicro/cio/api/cloudapplication"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"time"
)

// WireUpCloudApplicationDeployment prepares common resources to send request to Concerto API
func WireUpCloudApplicationDeployment(c *cli.Context) (ds *cloudapplication.CloudApplicationDeploymentService, f format.Formatter) {

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
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CloudApplicationDeploymentShow subcommand function
func CloudApplicationDeploymentShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationDeployment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	dep, err := svc.GetDeployment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloud application deployment data", err)
	}

	if err = formatter.PrintItem(*dep); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// CloudApplicationDeploymentDeploy subcommand function
func CloudApplicationDeploymentDeploy(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationDeployment(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)
	if c.IsSet("inputs") && c.IsSet("inputs-from-file") {
		return fmt.Errorf("invalid parameters detected. Please provide only one: 'inputs' or 'inputs-from-file'")
	}

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
				formatter.PrintFatal("Couldn't print/format result", err)
			}
			break
		}
		time.Sleep(5 * time.Second)
	}

	return nil
}

// CloudApplicationDeploymentDelete subcommand function
func CloudApplicationDeploymentDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpCloudApplicationDeployment(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := svc.DeleteDeployment(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete cloud application deployment", err)
	}
	return nil
}
