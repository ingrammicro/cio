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

// WireUpBrownfieldCloudAccount prepares common resources to send request to Concerto API
func WireUpBrownfieldCloudAccount(c *cli.Context) (ds *clientbrownfield.BrownfieldCloudAccountService, f format.Formatter) {

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
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

func checkCloudAccountLastDiscoveredStatus(c *cli.Context, cloudAccount *types.CloudAccount) {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	log.Info("Brownfield cloud account ID... ", cloudAccount.ID)
	log.Info("Checking discovery... ")
	for {
		ca, err := cloudAccountSvc.GetBrownfieldCloudAccount(c.String("id"))
		if err != nil {
			formatter.PrintFatal("Couldn't get cloud account data", err)
		}
		if (cloudAccount.LastDiscoveredAt != ca.LastDiscoveredAt) ||
			(cloudAccount.LastDiscoveredFailedAt != ca.LastDiscoveredFailedAt) {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

// BrownfieldCloudAccountServersDiscover subcommand function
func BrownfieldCloudAccountServersDiscover(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cloudAccount, err := cloudAccountSvc.DiscoverServers(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't discover servers", err)
	}

	checkCloudAccountLastDiscoveredStatus(c, cloudAccount)

	serversImportCandidates, err := cloudAccountSvc.ListServers(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list servers import candidates", err)
	}

	if err = formatter.PrintList(serversImportCandidates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// BrownfieldCloudAccountVPCsDiscover subcommand function
func BrownfieldCloudAccountVPCsDiscover(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cloudAccount, err := cloudAccountSvc.DiscoverVPCs(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't discover VPCs", err)
	}

	checkCloudAccountLastDiscoveredStatus(c, cloudAccount)

	vpcsImportCandidates, err := cloudAccountSvc.ListVPCs(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list VPCs import candidates", err)
	}

	if err = formatter.PrintList(vpcsImportCandidates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// BrownfieldCloudAccountFloatingIPsDiscover subcommand function
func BrownfieldCloudAccountFloatingIPsDiscover(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cloudAccount, err := cloudAccountSvc.DiscoverFloatingIPs(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't discover floating IPs", err)
	}

	checkCloudAccountLastDiscoveredStatus(c, cloudAccount)

	floatingIPsImportCandidates, err := cloudAccountSvc.ListFloatingIPs(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list floating IPs import candidates", err)
	}

	if err = formatter.PrintList(floatingIPsImportCandidates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// BrownfieldCloudAccountVolumesDiscover subcommand function
func BrownfieldCloudAccountVolumesDiscover(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpBrownfieldCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cloudAccount, err := cloudAccountSvc.DiscoverVolumes(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't discover volumes", err)
	}

	checkCloudAccountLastDiscoveredStatus(c, cloudAccount)

	volumesImportCandidates, err := cloudAccountSvc.ListVolumes(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list volumes import candidates", err)
	}

	if err = formatter.PrintList(volumesImportCandidates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
