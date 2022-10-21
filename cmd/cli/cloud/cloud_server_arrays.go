// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloud

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/spf13/viper"
)

func init() {
	fLabelsFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label as a query filter"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Server Array Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Name of the server array"}
	fNameReq := fName
	fNameReq.Required = true

	fTemplateId := cmd.FlagContext{Type: cmd.String, Name: cmd.TemplateId, Required: true,
		Usage: "Identifier of the template the server array shall use"}

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account in which the server array shall be registered"}

	fServerPlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerPlanId, Required: true,
		Usage: "Identifier of the server plan in which the server array shall be deployed"}

	fSize := cmd.FlagContext{Type: cmd.Int, Name: cmd.Size,
		Usage: "Number of initial servers in the server array. Value by default is 0"}

	fSizeAdd := cmd.FlagContext{Type: cmd.Int, Name: cmd.Size, Required: true,
		Usage: "The number of servers to add to the array, a number between 1 and 5"}

	fFirewallProfileId := cmd.FlagContext{Type: cmd.String, Name: cmd.FirewallProfileId,
		Usage: "Identifier of the firewall profile to which the server array belongs. " +
			"It will take default firewall profile if it is not given"}

	fSSHProfileId := cmd.FlagContext{Type: cmd.String, Name: cmd.SSHProfileId,
		Usage: "Identifier of the ssh profile to which the server array belongs. " +
			"It will take default ssh profile if it is not given"}

	fSubnetId := cmd.FlagContext{Type: cmd.String, Name: cmd.SubnetId,
		Usage: "Identifier of the subnet to which the server array belongs. " +
			"It will not be on any subnet managed by IMCO if not given"}

	fPrivateness := cmd.FlagContext{Type: cmd.Bool, Name: cmd.Privateness,
		Usage: "If the server array is private, set this flag, i.e: --privateness. " +
			"If it's public, do not set this flag"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with server array"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{
		Type:         cmd.String,
		Name:         cmd.ResourceType,
		DefaultValue: "server_array",
		Hidden:       true,
		Usage:        "Resource Type",
	}

	serverArraysCmd := cmd.NewCommand(cloudCmd, &cmd.CommandContext{
		Use:   "server-arrays",
		Short: "Provides information on server arrays"},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists information about all the server arrays on this account",
		RunMethod:    ServerArrayList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the server array identified by the given id",
		RunMethod:    ServerArrayShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:       "create",
		Short:     "Creates a new server array",
		RunMethod: ServerArrayCreate,
		FlagContexts: []cmd.FlagContext{
			fNameReq,
			fTemplateId,
			fCloudAccountId,
			fServerPlanId,
			fSize,
			fFirewallProfileId,
			fSSHProfileId,
			fSubnetId,
			fPrivateness,
			fLabels}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing server array",
		RunMethod:    ServerArrayUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use: "boot",
		Short: "This action boots all the servers in the server array with the given id. " +
			"The server array must be in an inactive state",
		RunMethod:    ServerArrayBoot,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use: "shutdown",
		Short: "This action shuts down all the servers in the server array identified by the given id. " +
			"The server must be in a bootstrap",
		RunMethod:    ServerArrayShutdown,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:          "empty",
		Short:        "This action empties all servers in server array with the given id",
		RunMethod:    ServerArrayEmpty,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:          "enlarge",
		Short:        "This action add servers in server array with the given id",
		RunMethod:    ServerArrayEnlarge,
		FlagContexts: []cmd.FlagContext{fId, fSizeAdd}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:          "list-servers",
		Short:        "This action list servers in server array with the given id",
		RunMethod:    ServerArrayServerList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use: "delete",
		Short: "This action decommissions the server array with the given id. " +
			"This action will only be allowed if the server array is empty",
		RunMethod:    ServerArrayDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(serverArraysCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// ServerArrayList subcommand function
func ServerArrayList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverArrays, err := svc.ListServerArrays(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive server array data", err)
		return err
	}

	labelables := make([]types.Labelable, len(serverArrays))
	for i := 0; i < len(serverArrays); i++ {
		labelables[i] = types.Labelable(serverArrays[i])
	}
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	filteredLabelables, err := labels.LabelFiltering(labelables, labelIDsByName)
	if err != nil {
		return err
	}
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	serverArrays = make([]*types.ServerArray, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		sa, ok := labelable.(*types.ServerArray)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.ServerArray, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		serverArrays[i] = sa
	}
	if err = formatter.PrintList(serverArrays); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerArrayShow subcommand function
func ServerArrayShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverArray, err := svc.GetServerArray(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive server array data", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerArrayCreate subcommand function
func ServerArrayCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverArrayIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"template_id":      viper.GetString(cmd.TemplateId),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
		"server_plan_id":   viper.GetString(cmd.ServerPlanId),
	}

	if viper.IsSet(cmd.Size) {
		serverArrayIn["size"] = viper.GetInt(cmd.Size)
	}
	if viper.IsSet(cmd.FirewallProfileId) {
		serverArrayIn["firewall_profile_id"] = viper.GetString(cmd.FirewallProfileId)
	}
	if viper.IsSet(cmd.SSHProfileId) {
		serverArrayIn["ssh_profile_id"] = viper.GetString(cmd.SSHProfileId)
	}
	if viper.IsSet(cmd.SubnetId) {
		serverArrayIn["subnet_id"] = viper.GetString(cmd.SubnetId)
	}
	if viper.IsSet(cmd.Privateness) {
		serverArrayIn["privateness"] = viper.GetBool(cmd.Privateness)
	}

	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}

	if viper.IsSet(cmd.Labels) {
		serverArrayIn["label_ids"], err = labels.LabelResolution(
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	serverArray, err := svc.CreateServerArray(cmd.GetContext(), &serverArrayIn)
	if err != nil {
		formatter.PrintError("Couldn't create server array", err)
		return err
	}

	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerArrayUpdate subcommand function
func ServerArrayUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverArrayIn := map[string]interface{}{}
	if viper.IsSet(cmd.Name) {
		serverArrayIn["name"] = viper.GetString(cmd.Name)
	}
	serverArray, err := svc.UpdateServerArray(cmd.GetContext(), viper.GetString(cmd.Id), &serverArrayIn)
	if err != nil {
		formatter.PrintError("Couldn't update server array", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerArrayBoot subcommand function
func ServerArrayBoot() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverArray, err := svc.BootServerArray(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't boot server array", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerArrayShutdown subcommand function
func ServerArrayShutdown() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverArray, err := svc.ShutdownServerArray(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't shutdown server array", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerArrayEmpty subcommand function
func ServerArrayEmpty() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverArray, err := svc.EmptyServerArray(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't empty server array", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerArrayEnlarge subcommand function
func ServerArrayEnlarge() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverArrayIn := map[string]interface{}{
		"size_increase": viper.GetInt(cmd.Size),
	}
	serverArray, err := svc.EnlargeServerArray(cmd.GetContext(), viper.GetString(cmd.Id), &serverArrayIn)
	if err != nil {
		formatter.PrintError("Couldn't enlarge server array", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerArrayServerList subcommand function
func ServerArrayServerList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	servers, err := svc.ListServerArrayServers(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive server data", err)
		return err
	}
	if err = formatServersResponse(servers, formatter); err != nil {
		return err
	}
	return nil
}

// ServerArrayDelete subcommand function
func ServerArrayDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteServerArray(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete server array", err)
		return err
	}
	return nil
}
