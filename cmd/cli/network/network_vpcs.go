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

	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "VPC Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Usage: "Name of the VPC"}
	fNameReq := fName
	fNameReq.Required = true

	fCidr := cmd.FlagContext{Type: cmd.String, Name: cmd.Cidr, Required: true, Usage: "CIDR of the VPC"}

	fCloudAccountId := cmd.FlagContext{Type: cmd.String, Name: cmd.CloudAccountId, Required: true,
		Usage: "Identifier of the cloud account in which the VPC is"}

	fRealmProviderName := cmd.FlagContext{Type: cmd.String, Name: cmd.RealmProviderName, Required: true,
		Usage: "Name of the provider realm in which the VPC is"}

	fLabels := cmd.FlagContext{Type: cmd.String, Name: cmd.Labels,
		Usage: "A list of comma separated label names to be associated with VPC"}

	fLabel := cmd.FlagContext{Type: cmd.String, Name: cmd.Label, Required: true, Usage: "Label name"}

	fResourceType := cmd.FlagContext{Type: cmd.String, Name: cmd.ResourceType, DefaultValue: "vpc", Hidden: true,
		Usage: "Resource Type"}

	vpcsCmd := cmd.NewCommand(networkCmd, &cmd.CommandContext{
		Use:   "vpcs",
		Short: "Provides information about Virtual Private Clouds (VPCs)"},
	)
	cmd.NewCommand(vpcsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all existing VPCs",
		RunMethod:    VPCList,
		FlagContexts: []cmd.FlagContext{fLabelsFilter}},
	)
	cmd.NewCommand(vpcsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the VPC identified by the given id",
		RunMethod:    VPCShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(vpcsCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new VPC",
		RunMethod:    VPCCreate,
		FlagContexts: []cmd.FlagContext{fNameReq, fCidr, fCloudAccountId, fRealmProviderName, fLabels}},
	)
	cmd.NewCommand(vpcsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing VPC identified by the given id",
		RunMethod:    VPCUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName}},
	)
	cmd.NewCommand(vpcsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a VPC",
		RunMethod:    VPCDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(vpcsCmd, &cmd.CommandContext{
		Use:          "discard",
		Short:        "Discards a VPC but does not delete it from the cloud provider",
		RunMethod:    VPCDiscard,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(vpcsCmd, &cmd.CommandContext{
		Use:          "add-label",
		Short:        "This action assigns a single label from a single labelable resource",
		RunMethod:    labels.LabelAdd,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
	cmd.NewCommand(vpcsCmd, &cmd.CommandContext{
		Use:          "remove-label",
		Short:        "This action unassigns a single label from a single labelable resource",
		RunMethod:    labels.LabelRemove,
		FlagContexts: []cmd.FlagContext{fId, fLabel, fResourceType}},
	)
}

// VPCList subcommand function
func VPCList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	vpcs, err := svc.ListVPCs(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive VPC data", err)
		return err
	}

	labelables := make([]types.Labelable, len(vpcs))
	for i := 0; i < len(vpcs); i++ {
		labelables[i] = types.Labelable(vpcs[i])
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

	vpcs = make([]*types.Vpc, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.Vpc)
		if !ok {
			e := fmt.Errorf("expected labelable to be a *types.Vpc, got a %T", labelable)
			formatter.PrintError(cmd.LabelFilteringUnexpected, e)
			return e
		}
		vpcs[i] = v
	}
	if err = formatter.PrintList(vpcs); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// VPCShow subcommand function
func VPCShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	vpc, err := svc.GetVPC(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive VPC data", err)
		return err
	}
	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	vpc.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*vpc); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// VPCCreate subcommand function
func VPCCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	vpcIn := map[string]interface{}{
		"name":                viper.GetString(cmd.Name),
		"cidr":                viper.GetString(cmd.Cidr),
		"cloud_account_id":    viper.GetString(cmd.CloudAccountId),
		"realm_provider_name": viper.GetString(cmd.RealmProviderName),
	}

	labelIDsByName, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}

	if viper.IsSet(cmd.Labels) {
		vpcIn["label_ids"], err = labels.LabelResolution(
			viper.GetString(cmd.Labels),
			&labelNamesByID,
			&labelIDsByName)
		if err != nil {
			return err
		}
	}

	vpc, err := svc.CreateVPC(cmd.GetContext(), &vpcIn)
	if err != nil {
		formatter.PrintError("Couldn't create VPC", err)
		return err
	}

	vpc.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*vpc); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// VPCUpdate subcommand function
func VPCUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	vpcIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	vpc, err := svc.UpdateVPC(cmd.GetContext(), viper.GetString(cmd.Id), &vpcIn)
	if err != nil {
		formatter.PrintError("Couldn't update VPC", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping()
	if err != nil {
		return err
	}
	vpc.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*vpc); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// VPCDelete subcommand function
func VPCDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteVPC(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete VPC", err)
		return err
	}
	return nil
}

// VPCDiscard subcommand function
func VPCDiscard() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DiscardVPC(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't discard VPC", err)
		return err
	}
	return nil
}
