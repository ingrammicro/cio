// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/cmd/cli/labels"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"
)

func init() {
	fId := cmd.FlagContext{Type: cmd.String, Name: cmd.Id, Required: true, Usage: "Subnet Id"}

	fVpcId := cmd.FlagContext{Type: cmd.String, Name: cmd.VpcId, Required: true, Usage: "VPC Id"}

	fName := cmd.FlagContext{Type: cmd.String, Name: cmd.Name, Required: true, Usage: "Name of the Subnet"}

	fCidr := cmd.FlagContext{Type: cmd.String, Name: cmd.Cidr, Required: true, Usage: "CIDR of the Subnet"}

	fType := cmd.FlagContext{Type: cmd.String, Name: cmd.Type, Required: true,
		Usage: "Type of the Subnet (among 'only_public', 'only_private' and 'public_and_private')"}

	subnetsCmd := cmd.NewCommand(networkCmd, &cmd.CommandContext{
		Use:   "subnets",
		Short: "Provides information about VPC subnets"},
	)
	cmd.NewCommand(subnetsCmd, &cmd.CommandContext{
		Use:          "list",
		Short:        "Lists all subnets of a VPC",
		RunMethod:    SubnetList,
		FlagContexts: []cmd.FlagContext{fVpcId}},
	)
	cmd.NewCommand(subnetsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the subnet identified by the given id",
		RunMethod:    SubnetShow,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(subnetsCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new subnet inside the specified VPC",
		RunMethod:    SubnetCreate,
		FlagContexts: []cmd.FlagContext{fVpcId, fName, fCidr, fType}},
	)
	cmd.NewCommand(subnetsCmd, &cmd.CommandContext{
		Use:          "update",
		Short:        "Updates an existing subnet identified by the given id",
		RunMethod:    SubnetUpdate,
		FlagContexts: []cmd.FlagContext{fId, fName}},
	)
	cmd.NewCommand(subnetsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a subnet",
		RunMethod:    SubnetDelete,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(subnetsCmd, &cmd.CommandContext{
		Use:          "list-servers",
		Short:        "Lists servers belonging to the subnet identified by the given id",
		RunMethod:    SubnetServerList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
	cmd.NewCommand(subnetsCmd, &cmd.CommandContext{
		Use:          "list-server-arrays",
		Short:        "Lists server arrays belonging to the subnet identified by the given id",
		RunMethod:    SubnetServerArrayList,
		FlagContexts: []cmd.FlagContext{fId}},
	)
}

// SubnetList subcommand function
func SubnetList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	subnets, err := svc.ListSubnets(cmd.GetContext(), viper.GetString(cmd.VpcId))
	if err != nil {
		formatter.PrintError("Couldn't receive subnets data", err)
		return err
	}

	if err = formatter.PrintList(subnets); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SubnetShow subcommand function
func SubnetShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	subnet, err := svc.GetSubnet(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive subnet data", err)
		return err
	}

	if err = formatter.PrintItem(*subnet); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SubnetCreate subcommand function
func SubnetCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	subnetIn := map[string]interface{}{
		"name":             viper.GetString(cmd.Name),
		"cidr":             viper.GetString(cmd.Cidr),
		"cloud_account_id": viper.GetString(cmd.CloudAccountId),
		"type":             viper.GetString(cmd.Type),
	}

	subnet, err := svc.CreateSubnet(cmd.GetContext(), viper.GetString(cmd.VpcId), &subnetIn)
	if err != nil {
		formatter.PrintError("Couldn't create subnet", err)
		return err
	}

	if err = formatter.PrintItem(*subnet); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SubnetUpdate subcommand function
func SubnetUpdate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	subnetIn := map[string]interface{}{
		"name": viper.GetString(cmd.Name),
	}

	subnet, err := svc.UpdateSubnet(cmd.GetContext(), viper.GetString(cmd.Id), &subnetIn)
	if err != nil {
		formatter.PrintError("Couldn't update subnet", err)
		return err
	}

	if err = formatter.PrintItem(*subnet); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SubnetDelete subcommand function
func SubnetDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteSubnet(cmd.GetContext(), viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't delete subnet", err)
		return err
	}
	return nil
}

// SubnetServerList subcommand function
func SubnetServerList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	servers, err := svc.ListSubnetServers(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive subnet servers data", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	for _, server := range servers {
		server.FillInLabelNames(labelNamesByID)
	}

	if err = formatter.PrintList(servers); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}

// SubnetServerArrayList subcommand function
func SubnetServerArrayList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	ctx := cmd.GetContext()
	serverArrays, err := svc.ListSubnetServerArrays(ctx, viper.GetString(cmd.Id))
	if err != nil {
		formatter.PrintError("Couldn't receive subnet server arrays data", err)
		return err
	}

	_, labelNamesByID, err := labels.LabelLoadsMapping(ctx)
	if err != nil {
		return err
	}
	for _, serverArray := range serverArrays {
		serverArray.FillInLabelNames(labelNamesByID)
	}

	if err = formatter.PrintList(serverArrays); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
