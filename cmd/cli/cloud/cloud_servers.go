// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloud

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"
	"strings"

	"github.com/ingrammicro/cio/cmd/cli/network"
	"github.com/ingrammicro/cio/utils/format"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/spf13/viper"
)

func init() {
	fLabelsFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label as a query filter"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Server Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Name of the server"}
	fNameReq := fName
	fNameReq.Required = true

	fSSHProfileId := cmd.FlagContext{Type: cmd.String, Name: cmd.SSHProfileId,
		Usage: "Identifier of the ssh profile which the server shall use"}

	fSSHProfileIds := cmd.FlagContext{Type: cmd.String, Name: cmd.SSHProfileIds,
		Usage: "A list of comma separated ssh profiles ids"}

	fFirewallProfileId := cmd.FlagContext{Type: cmd.String, Name: cmd.FirewallProfileId, Required: true,
		Usage: "Identifier of the firewall profile to which the server shall use"}

	fTemplateId := cmd.FlagContext{Type: cmd.String, Name: cmd.TemplateId, Required: true,
		Usage: "Identifier of the template the server shall use"}

	fServerPlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerPlanId, Required: true,
		Usage: "Identifier of the server plan in which the server shall be deployed"}

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account in which the server shall be registered"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with server"}

	fServerId := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerId, Required: true, Usage: "Server Id"}

	fScriptId := cmd.FlagContext{Type: cmd.String, Name: cmd.ScriptId, Required: true, Usage: "Script Id"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, DefaultValue: "server", Hidden: true,
		Usage: "Resource Type"}

	serversCmd := cmd.NewCommand(cloudCmd, &cmd.CommandContext{
		Use:   "servers",
		Short: "Provides information on servers"},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists information about all the servers on this account",
		RunMethod:    ServerList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the server identified by the given id",
		RunMethod:    ServerShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:       "create",
		Short:     "Creates a new server",
		RunMethod: ServerCreate,
		FlagContexts: []cmd.FlagContext{fNameReq, fSSHProfileId, fSSHProfileIds, fFirewallProfileId, fTemplateId,
			fServerPlanId, fCloudAccountId, fLabels}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing server",
		RunMethod:    ServerUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "boot",
		Short:        "Boots a server with the given id",
		RunMethod:    ServerBoot,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "reboot",
		Short:        "Reboots a server with the given id",
		RunMethod:    ServerReboot,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "shutdown",
		Short:        "Shuts down a server with the given id",
		RunMethod:    ServerShutdown,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use: "override-server",
		Short: "This action takes the server with the given id from a stalled state " +
			"to the operational state, at the user's own risk",
		RunMethod:    ServerOverride,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use: "delete",
		Short: "This action decommissions the server with the given id. " +
			"The server must be in a inactive, stalled or commission_stalled state",
		RunMethod:    ServerDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "list-events",
		Short:        "This action returns information about the events related to the server with the given id",
		RunMethod:    EventsList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use: "list-operational-scripts",
		Short: "This action returns information about the operational scripts characterisations " +
			"related to the server with the given id",
		RunMethod:    OperationalScriptsList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use: "execute-script",
		Short: "This action initiates the execution of the script characterisation with " +
			"the given id on the server with the given id",
		RunMethod:    OperationalScriptExecute,
		FlagContexts: []cmd.FlagContext{fServerId, fScriptId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "list-floating-ips",
		Short:        "This action returns information about the floating IPs attached to the server with the given id",
		RunMethod:    ServerFloatingIPList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "list-volumes",
		Short:        "This action returns information about the volumes attached to the server with the given id",
		RunMethod:    ServerVolumesList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(serversCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// ServerList subcommand function
func ServerList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	servers, err := svc.ListServers(ctx)
	if err != nil {
		formatter.PrintError("Couldn't receive servers data", err)
		return err
	}
	if err = formatServersResponse(ctx, servers, formatter); err != nil {
		return err
	}
	return nil
}

func formatServersResponse(ctx context.Context, servers []*types.Server, formatter format.Formatter) error {
	logger.DebugFuncInfo()

	labelables := make([]types.Labelable, len(servers))
	for i := 0; i < len(servers); i++ {
		labelables[i] = types.Labelable(servers[i])
	}
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	filteredLabelables, err := labels.LabelFiltering(labelables, labelIDsByName)
	if err != nil {
		return err
	}
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	servers = make([]*types.Server, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		s, ok := labelable.(*types.Server)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.server, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		servers[i] = s
	}
	if err := formatter.PrintList(servers); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerShow subcommand function
func ServerShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	server, err := svc.GetServer(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive server data", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerCreate subcommand function
func ServerCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"template_id":      viper.GetString(cmd.TemplateId),
		"server_plan_id":   viper.GetString(cmd.ServerPlanId),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
	}

	cmd.SetParamString("ssh_profile_id", cmd.SSHProfileId, serverIn)
	if viper.IsSet(cmd.SSHProfileIds) {
		serverIn["ssh_profile_ids"] = strings.Split(viper.GetString(cmd.SSHProfileIds), ",")
	}

	ctx := cmd.GetContext()
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}

	if viper.IsSet(cmd.Labels) {
		serverIn["label_ids"], err = labels.LabelResolution(
			ctx,
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	server, err := svc.CreateServer(ctx, &serverIn)
	if err != nil {
		formatter.PrintError("Couldn't create server", err)
		return err
	}

	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerUpdate subcommand function
func ServerUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	serverIn := map[string]interface{}{}
	cmd.SetParamString("name", cmd.Name, serverIn)
	ctx := cmd.GetContext()
	server, err := svc.UpdateServer(ctx, viper.GetString(cmd.Id), &serverIn)
	if err != nil {
		formatter.PrintError("Couldn't update server", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerBoot subcommand function
func ServerBoot() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	serverIn := map[string]interface{}{}
	server, err := svc.BootServer(ctx, viper.GetString(cmd.Id), &serverIn)
	if err != nil {
		formatter.PrintError("Couldn't boot server", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerReboot subcommand function
func ServerReboot() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	serverIn := map[string]interface{}{}
	server, err := svc.RebootServer(ctx, viper.GetString(cmd.Id), &serverIn)
	if err != nil {
		formatter.PrintError("Couldn't reboot server", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerShutdown subcommand function
func ServerShutdown() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	serverIn := map[string]interface{}{}
	server, err := svc.ShutdownServer(ctx, viper.GetString(cmd.Id), &serverIn)
	if err != nil {
		formatter.PrintError("Couldn't shutdown server", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerOverride subcommand function
func ServerOverride() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	serverIn := map[string]interface{}{}
	server, err := svc.OverrideServer(ctx, viper.GetString(cmd.Id), &serverIn)
	if err != nil {
		formatter.PrintError("Couldn't override server", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ServerDelete subcommand function
func ServerDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteServer(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete server", err)
		return err
	}
	return nil
}

// ServerFloatingIPList subcommand function
func ServerFloatingIPList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	floatingIPs, err := svc.ListServerFloatingIPs(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive server floating IPs data", err)
		return err
	}
	if err = network.FormatFloatingIPsResponse(ctx, floatingIPs, formatter); err != nil {
		return err
	}
	return nil
}

// ServerVolumesList subcommand function
func ServerVolumesList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	volumes, err := svc.ListServerVolumes(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive server volumes data", err)
		return err
	}
	if err = FormatVolumesResponse(ctx, volumes, formatter); err != nil {
		return err
	}
	return nil
}

func FormatVolumesResponse(ctx context.Context, volumes []*types.Volume, formatter format.Formatter) error {
	labelables := make([]types.Labelable, len(volumes))
	for i := 0; i < len(volumes); i++ {
		labelables[i] = types.Labelable(volumes[i])
	}
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	filteredLabelables, err := labels.LabelFiltering(labelables, labelIDsByName)
	if err != nil {
		return err
	}
	labels.LabelAssignNamesForIDs(filteredLabelables, labelNamesByID)

	volumes = make([]*types.Volume, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.Volume)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.Volume, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		volumes[i] = v
	}
	if err := formatter.PrintList(volumes); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// ========= Events ========

// EventsList subcommand function
func EventsList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	events, err := svc.ListServerEvents(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive server events data", err)
		return err
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

//======= Operational Scripts ==========

// OperationalScriptsList subcommand function
func OperationalScriptsList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	scripts, err := svc.ListOperationalScripts(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive server operational scripts data", err)
		return err
	}
	if err = formatter.PrintList(scripts); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// OperationalScriptExecute subcommand function
func OperationalScriptExecute() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	in := &map[string]interface{}{}
	scriptOut, err := svc.ExecuteOperationalScript(
		cmd.GetContext(),
		viper.GetString(cmd.ServerId),
		viper.GetString(cmd.ScriptId),
		in,
	)
	if err != nil {
		formatter.PrintError("Couldn't execute server operational script", err)
		return err
	}
	if err = formatter.PrintItem(*scriptOut); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
