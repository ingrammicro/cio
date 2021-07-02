// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"github.com/ingrammicro/cio/api/clientbrownfield"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpBrownfieldCloudAccount prepares common resources to send request to Concerto API
func WireUpBrownfieldCloudAccount(
	c *cli.Context,
) (ds *clientbrownfield.BrownfieldCloudAccountService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = clientbrownfield.NewBrownfieldCloudAccountService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up Brownfield Cloud Account service", err)
	}

	return ds, f
}

// BrownfieldCloudAccountList subcommand function
func BrownfieldCloudAccountList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpBrownfieldCloudAccount(c)

	cloudAccounts, err := svc.ListBrownfieldCloudAccounts()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloud accounts data", err)
	}

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	for id, ca := range cloudAccounts {
		cloudAccounts[id].CloudProviderName = cloudProvidersMap[ca.CloudProviderID]
	}

	if err = formatter.PrintList(cloudAccounts); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
