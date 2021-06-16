// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"github.com/ingrammicro/cio/api/wizard"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpWizServerPlan prepares common resources to send request to Concerto API
func WireUpWizServerPlan(c *cli.Context) (ds *wizard.WizardServerPlanService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = wizard.NewWizardServerPlanService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up wizard server plan service", err)
	}

	return ds, f
}

// WizServerPlanList subcommand function
func WizServerPlanList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpWizServerPlan(c)

	checkRequiredFlags(c, []string{"app-id", "location-id", "cloud-provider-id"}, formatter)

	serverPlans, err := serverPlanSvc.ListWizardServerPlans(
		c.String("app-id"),
		c.String("location-id"),
		c.String("cloud-provider-id"),
	)
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	locationsMap := LoadLocationsMapping(c)

	for id, sp := range serverPlans {
		serverPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
		serverPlans[id].LocationName = locationsMap[sp.LocationID]
	}

	if err = formatter.PrintList(serverPlans); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
