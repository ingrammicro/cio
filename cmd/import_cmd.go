// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"time"

	"github.com/ingrammicro/cio/api/clientbrownfield"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// WireUpImport prepares common resources to send request to Concerto API
func WireUpImport(c *cli.Context) (ds *clientbrownfield.ImportService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = clientbrownfield.NewImportService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up Import service", err)
	}

	return ds, f
}

func checkCloudAccountImportingState(
	c *cli.Context,
	cloudAccount *types.CloudAccount,
	state string,
) *types.CloudAccount {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	log.Info("Brownfield cloud account ID... ", cloudAccount.ID)
	log.Info("Checking importing process... ")
	for {
		ca, err := cloudAccountSvc.GetBrownfieldCloudAccount(c.String("id"))
		if err != nil {
			formatter.PrintFatal("Couldn't get cloud account data", err)
		}
		if (cloudAccount.State != ca.State) || (ca.State != state) {
			if ca.State == "idle" && ca.ErrorEventID != "" {
				log.Error("Error while importing: ", ca.ErrorEventID)
			} else {
				log.Info("Done!")
			}
			return ca
		}
		time.Sleep(5 * time.Second)
	}
}

// ImportServers subcommand function
func ImportServers(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importSvc, formatter := WireUpImport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importSvc.ImportServers(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import servers", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_servers")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ImportVPCs subcommand function
func ImportVPCs(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importSvc, formatter := WireUpImport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importSvc.ImportVPCs(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import vpcs", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_vpcs")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ImportFloatingIPs subcommand function
func ImportFloatingIPs(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importSvc, formatter := WireUpImport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importSvc.ImportFloatingIPs(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import floating IPs", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_floating_ips")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ImportVolumes subcommand function
func ImportVolumes(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importSvc, formatter := WireUpImport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importSvc.ImportVolumes(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import volumes", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_volumes")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ImportKubernetesClusters subcommand function
func ImportKubernetesClusters(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importSvc, formatter := WireUpImport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importSvc.ImportKubernetesClusters(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import kubernetes clusters", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_kubernetes_clusters")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ImportPolicies subcommand function
func ImportPolicies(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importSvc, formatter := WireUpImport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importSvc.ImportPolicies(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import policies", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_policies")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
