package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/settings"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpSaasAccount prepares common resources to send request to Concerto API
func WireUpSaasAccount(c *cli.Context) (ds *settings.SaasAccountService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = settings.NewSaasAccountService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up saasAccount service", err)
	}

	return ds, f
}

// SaasAccountList subcommand function
func SaasAccountList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	saasAccountSvc, formatter := WireUpSaasAccount(c)

	saasAccounts, err := saasAccountSvc.GetSaasAccountList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive saasAccount data", err)
	}
	if err = formatter.PrintList(saasAccounts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SaasAccountCreate subcommand function
func SaasAccountCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	saasAccountSvc, formatter := WireUpSaasAccount(c)

	checkRequiredFlags(c, []string{"saas_provider_id", "account_data"}, formatter)
	saasAccount, err := saasAccountSvc.CreateSaasAccount(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create saasAccount", err)
	}
	if err = formatter.PrintItem(*saasAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SaasAccountUpdate subcommand function
func SaasAccountUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	saasAccountSvc, formatter := WireUpSaasAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	saasAccount, err := saasAccountSvc.UpdateSaasAccount(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update saasAccount", err)
	}
	if err = formatter.PrintItem(*saasAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SaasAccountDelete subcommand function
func SaasAccountDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	saasAccountSvc, formatter := WireUpSaasAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := saasAccountSvc.DeleteSaasAccount(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete saasAccount", err)
	}
	return nil
}
