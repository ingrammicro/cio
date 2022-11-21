// Copyright (c) 2017-2022 Ingram Micro Inc.

package brownfield

import (
	"context"
	"github.com/ingrammicro/cio/cmd/cli"
	"time"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Cloud account Id"}

	cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use:          "import-servers",
		Short:        "Import servers for a given cloud account id",
		RunMethod:    ImportServers,
		FlagContexts: []cmd.FlagContext{fCloudAccountId}},
	)
	cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use:          "import-vpcs",
		Short:        "Import VPCs for a given cloud account id",
		RunMethod:    ImportVPCs,
		FlagContexts: []cmd.FlagContext{fCloudAccountId}},
	)
	cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use:          "import-floating-ips",
		Short:        "Import floating IPs for a given cloud account id",
		RunMethod:    ImportFloatingIPs,
		FlagContexts: []cmd.FlagContext{fCloudAccountId}},
	)
	cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use:          "import-volumes",
		Short:        "Import volumes for a given cloud account id",
		RunMethod:    ImportVolumes,
		FlagContexts: []cmd.FlagContext{fCloudAccountId}},
	)
	cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use:          "import-policies",
		Short:        "Import policies for a given cloud account id",
		RunMethod:    ImportPolicies,
		FlagContexts: []cmd.FlagContext{fCloudAccountId}},
	)
	cmd.NewCommand(brownfieldCmd, &cmd.CommandContext{
		Use:          "import-k8s-clusters",
		Short:        "Import kubernetes clusters for a given cloud account id",
		RunMethod:    ImportKubernetesClusters,
		FlagContexts: []cmd.FlagContext{fCloudAccountId}},
	)
}

func checkCloudAccountImportingState(
	ctx context.Context,
	cloudAccount *types.CloudAccount,
	state string,
) (*types.CloudAccount, error) {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	log.Info("Brownfield cloud account ID... ", cloudAccount.ID)
	log.Info("Checking importing process... ")
	for {
		ca, err := svc.GetBrownfieldCloudAccount(ctx, viper.GetString(cmd.Id))
		if err != nil {
			formatter.PrintError("Couldn't receive cloud account data", err)
			return nil, err
		}
		if (cloudAccount.State != ca.State) || (ca.State != state) {
			if ca.State == "idle" && ca.ErrorEventID != "" {
				log.Error("Error while importing: ", ca.ErrorEventID)
			} else {
				log.Info("Done!")
			}
			return ca, nil
		}
		time.Sleep(5 * time.Second)
	}
}

// ImportServers subcommand function
func ImportServers() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	cloudAccount, err := svc.ImportServers(ctx, viper.GetString(cmd.Id), &map[string]interface{}{})
	if err != nil {
		formatter.PrintError("Couldn't import servers", err)
		return err
	}

	cloudAccount, err = checkCloudAccountImportingState(ctx, cloudAccount, "importing_servers")
	if err != nil {
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(ctx)
	if err != nil {
		return err
	}
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ImportVPCs subcommand function
func ImportVPCs() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	cloudAccount, err := svc.ImportVPCs(ctx, viper.GetString(cmd.Id), &map[string]interface{}{})
	if err != nil {
		formatter.PrintError("Couldn't import vpcs", err)
		return err
	}

	cloudAccount, err = checkCloudAccountImportingState(ctx, cloudAccount, "importing_vpcs")
	if err != nil {
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(ctx)
	if err != nil {
		return err
	}
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ImportFloatingIPs subcommand function
func ImportFloatingIPs() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	cloudAccount, err := svc.ImportFloatingIPs(ctx, viper.GetString(cmd.Id), &map[string]interface{}{})
	if err != nil {
		formatter.PrintError("Couldn't import floating IPs", err)
		return err
	}

	cloudAccount, err = checkCloudAccountImportingState(ctx, cloudAccount, "importing_floating_ips")
	if err != nil {
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(ctx)
	if err != nil {
		return err
	}
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ImportVolumes subcommand function
func ImportVolumes() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	cloudAccount, err := svc.ImportVolumes(ctx, viper.GetString(cmd.Id), &map[string]interface{}{})
	if err != nil {
		formatter.PrintError("Couldn't import volumes", err)
		return err
	}

	cloudAccount, err = checkCloudAccountImportingState(ctx, cloudAccount, "importing_volumes")
	if err != nil {
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(ctx)
	if err != nil {
		return err
	}
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ImportPolicies subcommand function
func ImportPolicies() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	cloudAccount, err := svc.ImportPolicies(ctx, viper.GetString(cmd.Id), &map[string]interface{}{})
	if err != nil {
		formatter.PrintError("Couldn't import policies", err)
		return err
	}

	cloudAccount, err = checkCloudAccountImportingState(ctx, cloudAccount, "importing_policies")
	if err != nil {
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(ctx)
	if err != nil {
		return err
	}
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ImportKubernetesClusters subcommand function
func ImportKubernetesClusters() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	cloudAccount, err := svc.ImportKubernetesClusters(
		ctx,
		viper.GetString(cmd.Id),
		&map[string]interface{}{},
	)
	if err != nil {
		formatter.PrintError("Couldn't import kubernetes clusters", err)
		return err
	}

	cloudAccount, err = checkCloudAccountImportingState(ctx, cloudAccount, "importing_kubernetes_clusters")
	if err != nil {
		return err
	}

	cloudProvidersMap, err := cli.LoadCloudProvidersMapping(ctx)
	if err != nil {
		return err
	}
	cloudAccount.CloudProviderName = cloudProvidersMap[cloudAccount.CloudProviderID]

	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
