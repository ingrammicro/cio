// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/cmd/cli"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/spf13/viper"
)

func init() {
	fServerId := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerId,
		Usage: "Identifier of a server to return only the floating IPs that are attached with that server"}

	fLabelsFilter := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label as a query filter"}

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Floating IP Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Name of the floating IP"}
	fNameReq := fName
	fNameReq.Required = true

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account in which the floating IP is"}

	fRealmId := cmd.FlagContext{Type: cmd.String, Name: cmd.RealmId, Required: true,
		Usage: "Identifier of the realm in which the floating IP is"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with floating IP"}

	fServerIdAttach := cmd.FlagContext{Type: cmd.String, Name: cmd.ServerId, Required: true,
		Usage: "Identifier of the server to attach the floating IP"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{
		Type:         cmd.String,
		Name:         cmd.ResourceType,
		DefaultValue: "floating_ip",
		Hidden:       true,
		Usage:        "Resource Type",
	}

	floatingIpsCmd := cmd.NewCommand(networkCmd, &cmd.CommandContext{
		Use:   "floating-ips",
		Short: "Provides information about floating IPs"},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all existing floating IPs",
		RunMethod:    FloatingIPList,
		FlagContexts: []cmd.FlagContext{fServerId, fLabelsFilter}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the floating IP identified by the given id",
		RunMethod:    FloatingIPShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new floating IP",
		RunMethod:    FloatingIPCreate,
		FlagContexts: []cmd.FlagContext{fNameReq, fCloudAccountId, fRealmId, fLabels}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing floating IP identified by the given id",
		RunMethod:    FloatingIPUpdate,
		FlagContexts: []cmd.FlagContext{fId, fNameReq}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "attach",
		Short:        "Attaches the floating IP to server",
		RunMethod:    FloatingIPAttach,
		FlagContexts: []cmd.FlagContext{fId, fServerIdAttach}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "detach",
		Short:        "Detaches a floating IP from server",
		RunMethod:    FloatingIPDetach,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a floating IP",
		RunMethod:    FloatingIPDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "discard",
		Short:        "Discards a floating IP but does not delete it from the cloud provider",
		RunMethod:    FloatingIPDiscard,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(floatingIpsCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// FloatingIPList subcommand function
func FloatingIPList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	floatingIPs, err := svc.ListFloatingIPs(ctx, viper.GetString(cmd.ServerId))
	if err != nil {
		formatter.PrintError("Couldn't receive floating IPs data", err)
		return err
	}
	if err = FormatFloatingIPsResponse(ctx, floatingIPs, formatter); err != nil {
		return err
	}
	return nil
}

func FormatFloatingIPsResponse(ctx context.Context, floatingIPs []*types.FloatingIP, formatter format.Formatter) error {
	labelables := make([]types.Labelable, len(floatingIPs))
	for i := 0; i < len(floatingIPs); i++ {
		labelables[i] = types.Labelable(floatingIPs[i])
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

	floatingIPs = make([]*types.FloatingIP, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		fIP, ok := labelable.(*types.FloatingIP)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.FloatingIP, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		floatingIPs[i] = fIP
	}
	if err := formatter.PrintList(floatingIPs); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// FloatingIPShow subcommand function
func FloatingIPShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	floatingIP, err := svc.GetFloatingIP(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive floating IP data", err)
		return err
	}
	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	floatingIP.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*floatingIP); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// FloatingIPCreate subcommand function
func FloatingIPCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	floatingIPIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
		"realm_id":         viper.GetString(cmd.RealmId),
	}

	ctx := cmd.GetContext()
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}

	if viper.IsSet(cmd.Labels) {
		floatingIPIn["label_ids"], err = labels.LabelResolution(
			ctx,
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	floatingIP, err := svc.CreateFloatingIP(ctx, &floatingIPIn)
	if err != nil {
		formatter.PrintError("Couldn't create floating IP", err)
		return err
	}

	floatingIP.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*floatingIP); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// FloatingIPUpdate subcommand function
func FloatingIPUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	floatingIPIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	ctx := cmd.GetContext()
	floatingIP, err := svc.UpdateFloatingIP(ctx, viper.GetString(cmd.Id), &floatingIPIn)
	if err != nil {
		formatter.PrintError("Couldn't update floating IP", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	floatingIP.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*floatingIP); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// FloatingIPAttach subcommand function
func FloatingIPAttach() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	floatingIPIn := map[string]interface{}{
		"attached_server_id": viper.GetString(cmd.ServerId),
	}

	ctx := cmd.GetContext()
	server, err := svc.AttachFloatingIP(ctx, viper.GetString(cmd.Id), &floatingIPIn)
	if err != nil {
		formatter.PrintError("Couldn't attach floating IP", err)
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

// FloatingIPDetach subcommand function
func FloatingIPDetach() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DetachFloatingIP(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't detach floating IP", err)
		return err
	}
	return nil
}

// FloatingIPDelete subcommand function
func FloatingIPDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteFloatingIP(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete floating IP", err)
		return err
	}
	return nil
}

// FloatingIPDiscard subcommand function
func FloatingIPDiscard() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DiscardFloatingIP(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't discard floating IP", err)
		return err
	}
	return nil
}
