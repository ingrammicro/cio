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

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "SSH profile id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Name of the SSH profile"}
	fNameReq := fName
	fNameReq.Required = true

	fPublicKey := cmd.FlagContext{Type: cmd.String, Name: cmd.PublicKey, Usage: "Public key of the SSH profile"}
	fPublicKeyReq := fPublicKey
	fPublicKeyReq.Required = true

	fPrivateKey := cmd.FlagContext{Type: cmd.String, Name: cmd.PrivateKey, Usage: "Private key of the SSH profile"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with SSH profile"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{
		Type:         cmd.String,
		Name:         cmd.ResourceType,
		DefaultValue: "ssh_profile",
		Hidden:       true,
		Usage:        "Resource Type",
	}

	sshProfilesCmd := cmd.NewCommand(cloudCmd, &cmd.CommandContext{
		Use:   "ssh-profiles",
		Short: "Provides information on SSH profiles"},
	)
	cmd.NewCommand(sshProfilesCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all available SSH profiles",
		RunMethod:    SSHProfileList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(sshProfilesCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the SSH profile identified by the given id",
		RunMethod:    SSHProfileShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(sshProfilesCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new SSH profile",
		RunMethod:    SSHProfileCreate,
		FlagContexts: []cmd.FlagContext{fNameReq, fPublicKeyReq, fPrivateKey, fLabels}},
	)
	cmd.NewCommand(sshProfilesCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing SSH profile",
		RunMethod:    SSHProfileUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName, fPublicKey, fPrivateKey}},
	)
	cmd.NewCommand(sshProfilesCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a SSH profile",
		RunMethod:    SSHProfileDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(sshProfilesCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(sshProfilesCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// SSHProfileList subcommand function
func SSHProfileList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	sshProfiles, err := svc.ListSSHProfiles(ctx)
	if err != nil {
		formatter.PrintError("Couldn't receive ssh profiles data", err)
		return err
	}

	labelables := make([]types.Labelable, len(sshProfiles))
	for i := 0; i < len(sshProfiles); i++ {
		labelables[i] = types.Labelable(sshProfiles[i])
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
	sshProfiles = make([]*types.SSHProfile, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		sshP, ok := labelable.(*types.SSHProfile)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.SSHProfile, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		sshProfiles[i] = sshP
	}

	if err = formatter.PrintList(sshProfiles); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SSHProfileShow subcommand function
func SSHProfileShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	sshProfile, err := svc.GetSSHProfile(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive ssh profile data", err)
		return err
	}
	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	sshProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SSHProfileCreate subcommand function
func SSHProfileCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	sshProfileIn := map[string]interface{}{
		"name":       viper.GetString(cmd.Name),
		"public_key": viper.GetString(cmd.PublicKey),
	}
	cmd.SetParamString("private_key", cmd.PrivateKey, sshProfileIn)

	ctx := cmd.GetContext()
	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}

	if viper.IsSet(cmd.Labels) {
		sshProfileIn["label_ids"], err = labels.LabelResolution(
			ctx,
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	sshProfile, err := svc.CreateSSHProfile(ctx, &sshProfileIn)
	if err != nil {
		formatter.PrintError("Couldn't create ssh profile", err)
		return err
	}

	sshProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SSHProfileUpdate subcommand function
func SSHProfileUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	sshProfileIn := map[string]interface{}{}
	cmd.SetParamString("name", cmd.Name, sshProfileIn)
	cmd.SetParamString("public_key", cmd.PublicKey, sshProfileIn)
	cmd.SetParamString("private_key", cmd.PrivateKey, sshProfileIn)

	ctx := cmd.GetContext()
	sshProfile, err := svc.UpdateSSHProfile(ctx, viper.GetString(cmd.Id), &sshProfileIn)
	if err != nil {
		formatter.PrintError("Couldn't update ssh profile", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	sshProfile.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SSHProfileDelete subcommand function
func SSHProfileDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteSSHProfile(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete ssh profile", err)
		return err
	}
	return nil
}
