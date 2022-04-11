// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"github.com/ingrammicro/cio/api/cloud"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpServerPlan prepares common resources to send request to Concerto API
func WireUpServerPlan(c *cli.Context) (ds *cloud.ServerPlanService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloud.NewServerPlanService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up server plan service", err)
	}

	return ds, f
}

// ServerPlanList subcommand function
func ServerPlanList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpServerPlan(c)

	checkRequiredFlags(c, []string{"cloud-provider-id", "realm-id"}, formatter)
	serverPlans, err := serverPlanSvc.ListServerPlans(c.String("cloud-provider-id"), c.String("realm-id"))
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

// ServerPlanShow subcommand function
func ServerPlanShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpServerPlan(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverPlan, err := serverPlanSvc.GetServerPlan(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	locationsMap := LoadLocationsMapping(c)
	serverPlan.LocationName = locationsMap[serverPlan.LocationID]

	if err = formatter.PrintItem(*serverPlan); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
