package cmd

import (
	"github.com/ingrammicro/cio/api/cloud"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpRealm prepares common resources to send request to Concerto API
func WireUpRealm(c *cli.Context) (cs *cloud.RealmService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	cs, err = cloud.NewRealmService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up realm service", err)
	}

	return cs, f
}

// RealmList subcommand function
func RealmList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpRealm(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	realms, err := svc.ListRealms(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive realms data", err)
	}

	if err = formatter.PrintList(realms); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// RealmShow subcommand function
func RealmShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpRealm(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	realm, err := svc.GetRealm(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive realm data", err)
	}

	if err = formatter.PrintItem(*realm); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// RealmNodePoolPlansList subcommand function
func RealmNodePoolPlansList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpRealm(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	realms, err := svc.ListRealmNodePoolPlans(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive realms data", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	for id, sp := range realms {
		realms[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
	}

	if err = formatter.PrintList(realms); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
