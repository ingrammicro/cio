package cmd

import (
	"github.com/ingrammicro/cio/api/wizard"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpApp prepares common resources to send request to Concerto API
func WireUpApp(c *cli.Context) (ds *wizard.AppService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = wizard.NewAppService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up app service", err)
	}

	return ds, f
}

// AppList subcommand function
func AppList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	appSvc, formatter := WireUpApp(c)

	apps, err := appSvc.ListApps()
	if err != nil {
		formatter.PrintFatal("Couldn't receive app data", err)
	}
	if err = formatter.PrintList(apps); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// AppDeploy subcommand function
func AppDeploy(c *cli.Context) error {
	debugCmdFuncInfo(c)
	appSvc, formatter := WireUpApp(c)

	checkRequiredFlags(c, []string{"id", "location-id", "cloud-account-id", "hostname"}, formatter)

	appIn := map[string]interface{}{
		"location_id":      c.String("location-id"),
		"cloud_account_id": c.String("cloud-account-id"),
		"hostname":         c.String("hostname"),
	}
	if c.IsSet("server-plan-id") {
		appIn["server_plan_id"] = c.String("server-plan-id")
	}

	app, err := appSvc.DeployApp(c.String("id"), &appIn)
	if err != nil {
		formatter.PrintFatal("Couldn't deploy app", err)
	}
	if err = formatter.PrintItem(*app); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
