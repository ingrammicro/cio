package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/cloud"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpSaasProvider prepares common resources to send request to Concerto API
func WireUpSaasProvider(c *cli.Context) (cs *cloud.SaasProviderService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	cs, err = cloud.NewSaasProviderService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up saasProvider service", err)
	}

	return cs, f
}

// SaasProviderList subcommand function
func SaasProviderList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	saasProviderSvc, formatter := WireUpSaasProvider(c)

	saasProviders, err := saasProviderSvc.GetSaasProviderList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive saasProvider data", err)
	}
	if err = formatter.PrintList(saasProviders); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
