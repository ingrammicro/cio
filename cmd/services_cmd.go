package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/blueprint"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpService prepares common resources to send request to Concerto API
func WireUpService(c *cli.Context) (sv *blueprint.ServicesService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	sv, err = blueprint.NewServicesService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up service service", err)
	}

	return sv, f
}

// ServiceList subcommand function
func ServiceList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serviceSvc, formatter := WireUpService(c)

	services, err := serviceSvc.GetServiceList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive service data", err)
	}
	if err = formatter.PrintList(services); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServiceShow subcommand function
func ServiceShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serviceSvc, formatter := WireUpService(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	service, err := serviceSvc.GetService(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive service data", err)
	}
	if err = formatter.PrintItem(*service); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
