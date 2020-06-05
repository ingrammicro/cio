// Copyright (c) 2017-2021 Ingram Micro Inc.

package wizard

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fAppId := cmd.FlagContext{Type: cmd.String, Name: cmd.AppId, Required: true, Usage: "Identifier of the App"}

	fLocationId := cmd.FlagContext{Type: cmd.String, Name: cmd.LocationId, Required: true,
		Usage: "Identifier of the Location"}

	fCloudProviderId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudProviderId, Required: true,
		Usage: "Identifier of the Cloud Provider"}

	serverPlansCmd := cmd.NewCommand(WizardCmd, &cmd.CommandContext{
		Use:   "server-plans",
		Short: "Provides information about server plans"},
	)
	cmd.NewCommand(serverPlansCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists the available server Plans",
		RunMethod:    WizServerPlanList,
		FlagContexts: []cmd.FlagContext{fAppId, fLocationId, fCloudProviderId}},
	)
}

// WizServerPlanList subcommand function
func WizServerPlanList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	serverPlans, err := svc.ListWizardServerPlans(
		viper.GetString(cmd.AppId),
		viper.GetString(cmd.LocationId),
		viper.GetString(cmd.CloudProviderId),
	)
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	cloudProvidersMap := cmd.LoadCloudProvidersMapping()
	locationsMap := cmd.LoadLocationsMapping()

	for id, sp := range serverPlans {
		serverPlans[id].CloudProviderName = cloudProvidersMap[sp.CloudProviderID]
		serverPlans[id].LocationName = locationsMap[sp.LocationID]
	}

	if err = formatter.PrintList(serverPlans); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
