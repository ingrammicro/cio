package cmd

import (
	"github.com/ingrammicro/cio/api/clientbrownfield"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"time"
)

// WireUpImportCandidate prepares common resources to send request to Concerto API
func WireUpImportCandidate(c *cli.Context) (ds *clientbrownfield.ImportCandidateService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = clientbrownfield.NewImportCandidateService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up Import Candidate service", err)
	}

	return ds, f
}

func checkCloudAccountImportingState(c *cli.Context, cloudAccount *types.CloudAccount, state string) *types.CloudAccount {
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

// ImportCandidateServers subcommand function
func ImportCandidateServers(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importCandidateSvc.ImportServers(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import servers", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_servers")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateVPCs subcommand function
func ImportCandidateVPCs(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importCandidateSvc.ImportVPCs(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import vpcs", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_vpcs")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateFloatingIPs subcommand function
func ImportCandidateFloatingIPs(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importCandidateSvc.ImportFloatingIPs(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import floating IPs", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_floating_ips")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateVolumes subcommand function
func ImportCandidateVolumes(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importCandidateSvc.ImportVolumes(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import volumes", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_volumes")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidateKubernetesClusters subcommand function
func ImportCandidateKubernetesClusters(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importCandidateSvc.ImportKubernetesClusters(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import kubernetes clusters", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_kubernetes_clusters")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ImportCandidatePolicies subcommand function
func ImportCandidatePolicies(c *cli.Context) error {
	debugCmdFuncInfo(c)
	importCandidateSvc, formatter := WireUpImportCandidate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	cloudAccount, err := importCandidateSvc.ImportPolicies(c.String("id"), &map[string]interface{}{})
	if err != nil {
		formatter.PrintFatal("Couldn't import policies", err)
	}

	cloudAccount = checkCloudAccountImportingState(c, cloudAccount, "importing_policies")

	cloudProvidersMap := LoadCloudProvidersMapping(c)
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
