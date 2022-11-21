// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

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

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Firewall profile Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Logical name of the firewall profile"}
	fNameReq := fName
	fNameReq.Required = true

	fDescription := cmd.FlagContext{Type: cmd.String, Name: cmd.Description,
		Usage: "Description of the firewall profile"}
	fDescriptionReq := fDescription
	fDescriptionReq.Required = true

	fRules := cmd.FlagContext{Type: cmd.String, Name: cmd.Rules,
		Usage: "Set of rules of the firewall profile, " +
			"i.e: --rules TCP/8080-8081:0.0.0.0/0,TCP/9090-9091:any,UDP/3456:1.2.3.4\n\t" +
			"Rule format: [PROTOCOL/MIN_PORT[-MAX_PORT]:CIDR_IP]"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with firewall profile"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, DefaultValue: "firewall_profile",
		Hidden: true, Usage: "Resource Type"}

	firewallProfilesCmd := cmd.NewCommand(networkCmd, &cmd.CommandContext{
		Use:   "firewall-profiles",
		Short: "Provides information about firewall profiles"},
	)
	cmd.NewCommand(firewallProfilesCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all existing firewall profiles",
		RunMethod:    FirewallProfileList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(firewallProfilesCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the firewall profile identified by the given id",
		RunMethod:    FirewallProfileShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(firewallProfilesCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a a firewall profile with the given parameters",
		RunMethod:    FirewallProfileCreate,
		FlagContexts: []cmd.FlagContext{fNameReq, fDescriptionReq, fRules, fLabels}},
	)
	cmd.NewCommand(firewallProfilesCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates the firewall profile identified by the given id with the given parameters",
		RunMethod:    FirewallProfileUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName, fDescription, fRules}},
	)
	cmd.NewCommand(firewallProfilesCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a firewall profile",
		RunMethod:    FirewallProfileDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(firewallProfilesCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(firewallProfilesCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// FirewallProfileList subcommand function
func FirewallProfileList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	firewallProfiles, err := svc.ListFirewallProfiles(ctx)
	if err != nil {
		formatter.PrintError("Couldn't receive firewall profiles data", err)
		return err
	}

	labelables := make([]types.Labelable, len(firewallProfiles))
	for i := 0; i < len(firewallProfiles); i++ {
		labelables[i] = types.Labelable(firewallProfiles[i])
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

	firewallProfiles = make([]*types.FirewallProfile, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		fw, ok := labelable.(*types.FirewallProfile)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.FirewallProfile, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		firewallProfiles[i] = fw
	}
	if err = formatter.PrintList(firewallProfiles); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// FirewallProfileShow subcommand function
func FirewallProfileShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	firewallProfile, err := svc.GetFirewallProfile(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive firewall profile data", err)
		return err
	}
	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	firewallProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// FirewallProfileCreate subcommand function
func FirewallProfileCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	firewallProfileIn := map[string]interface{}{
		"name":        viper.GetString(cmd.Name),
		"description": viper.GetString(cmd.Description),
	}

	if viper.IsSet(cmd.Rules) {
		fw := new(types.FirewallProfile)
		if err := fw.ConvertFlagParamsToRules(viper.GetString(cmd.Rules)); err != nil {
			formatter.PrintError("Error parsing parameters", err)
			return err
		}
		firewallProfileIn["rules"] = fw.Rules
	}

	ctx := cmd.GetContext()
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}

	if viper.IsSet(cmd.Labels) {
		firewallProfileIn["label_ids"], err = labels.LabelResolution(
			ctx,
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	firewallProfile, err := svc.CreateFirewallProfile(ctx, &firewallProfileIn)
	if err != nil {
		formatter.PrintError("Couldn't create firewall profile", err)
		return err
	}

	firewallProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// FirewallProfileUpdate subcommand function
func FirewallProfileUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	firewallProfileIn := map[string]interface{}{}
	if viper.IsSet(cmd.Name) {
		firewallProfileIn["name"] = viper.GetString(cmd.Name)
	}
	if viper.IsSet(cmd.Description) {
		firewallProfileIn["description"] = viper.GetString(cmd.Description)
	}
	if viper.IsSet(cmd.Rules) {
		fw := new(types.FirewallProfile)
		if err := fw.ConvertFlagParamsToRules(viper.GetString(cmd.Rules)); err != nil {
			formatter.PrintError("Error parsing parameters", err)
			return err
		}
		firewallProfileIn["rules"] = fw.Rules
	}

	ctx := cmd.GetContext()
	firewallProfile, err := svc.UpdateFirewallProfile(ctx, viper.GetString(cmd.Id), &firewallProfileIn)
	if err != nil {
		formatter.PrintError("Couldn't update firewall profile", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	firewallProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// FirewallProfileDelete subcommand function
func FirewallProfileDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteFirewallProfile(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete firewall profile", err)
		return err
	}
	return nil
}
