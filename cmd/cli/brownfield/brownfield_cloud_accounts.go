// Copyright (c) 2017-2022 Ingram Micro Inc.

package brownfield

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
)

func init() {
	cloudAccountsCmd := cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use: "cloud-accounts",
		Short: "Provides information about brownfield cloud accounts. " +
			"Allows querying cloud accounts to import resources from IMCO"},
	)
	cmd.NewCommand(cloudAccountsCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "Lists the cloud accounts that support importing resources",
		RunMethod: BrownfieldCloudAccountList},
	)
}

// BrownfieldCloudAccountList subcommand function
func BrownfieldCloudAccountList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	cloudAccounts, err := svc.ListBrownfieldCloudAccounts(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloud accounts data", err)
	}

	cloudProvidersMap := cli.LoadCloudProvidersMapping(cmd.GetContext())
	for id, ca := range cloudAccounts {
		cloudAccounts[id].CloudProviderName = cloudProvidersMap[ca.CloudProviderID]
	}

	if err = formatter.PrintList(cloudAccounts); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
