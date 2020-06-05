// Copyright (c) 2017-2022 Ingram Micro Inc.

package network

import (
	"github.com/ingrammicro/cio/cmd/cli"
	"strings"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/utils"
	"github.com/spf13/viper"
)

func init() {
	fVpcId := cmd.FlagContext{Type: cmd.String, Name: cmd.VpcId, Required: true, Usage: "VPC Id"}

	fPublicIp := cmd.FlagContext{Type: cmd.String, Name: cmd.PublicIp, Required: true, Usage: "Public Ip for the VPN"}

	fPsk := cmd.FlagContext{Type: cmd.String, Name: cmd.Psk, Required: true, Usage: "Pass key of the VPN"}

	fExposedCidrs := cmd.FlagContext{Type: cmd.String, Name: cmd.ExposedCidrs, Required: true,
		Usage: "A list of comma separated exposed CIDRs of the VPN"}

	fVpnPlanId := cmd.FlagContext{Type: cmd.String, Name: cmd.VpnPlanId, Required: true,
		Usage: "Identifier of the VPN plan"}

	vpnsCmd := cmd.NewCommand(networkCmd, &cmd.CommandContext{
		Use:   "vpns",
		Short: "Provides information about VPC Virtual Private Networks (VPNs)"},
	)
	cmd.NewCommand(vpnsCmd, &cmd.CommandContext{
		Use:          "show",
		Short:        "Shows information about the VPN identified by the given VPC id",
		RunMethod:    VPNShow,
		FlagContexts: []cmd.FlagContext{fVpcId}},
	)
	cmd.NewCommand(vpnsCmd, &cmd.CommandContext{
		Use:          "create",
		Short:        "Creates a new VPN for the specified VPC",
		RunMethod:    VPNCreate,
		FlagContexts: []cmd.FlagContext{fVpcId, fPublicIp, fPsk, fExposedCidrs, fVpnPlanId}},
	)
	cmd.NewCommand(vpnsCmd, &cmd.CommandContext{
		Use:          "delete",
		Short:        "Deletes a VPN of the specified VPC",
		RunMethod:    VPNDelete,
		FlagContexts: []cmd.FlagContext{fVpcId}},
	)
	cmd.NewCommand(vpnsCmd, &cmd.CommandContext{
		Use:          "list-plans",
		Short:        "Lists VPN plans of the specified VPC",
		RunMethod:    VPNPlanList,
		FlagContexts: []cmd.FlagContext{fVpcId}},
	)
}

// VPNShow subcommand function
func VPNShow() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	vpn, err := svc.GetVPN(cmd.GetContext(), viper.GetString(cmd.VpcId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive VPN data", err)
	}
	if err = formatter.PrintItem(*vpn); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// VPNCreate subcommand function
func VPNCreate() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	vpnIn := map[string]interface{}{
		"public_ip":     viper.GetString(cmd.PublicIp),
		"psk":           viper.GetString(cmd.Psk),
		"exposed_cidrs": utils.RemoveDuplicates(strings.Split(viper.GetString(cmd.ExposedCidrs), ",")),
		"vpn_plan_id":   viper.GetString(cmd.VpnPlanId),
	}

	vpn, err := svc.CreateVPN(cmd.GetContext(), viper.GetString(cmd.VpcId), &vpnIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create VPN", err)
	}

	if err = formatter.PrintItem(*vpn); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// VPNDelete subcommand function
func VPNDelete() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	err := svc.DeleteVPN(cmd.GetContext(), viper.GetString(cmd.VpcId))
	if err != nil {
		formatter.PrintFatal("Couldn't delete VPN", err)
	}
	return nil
}

// VPNPlanList subcommand function
func VPNPlanList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	vpns, err := svc.ListVPNPlans(cmd.GetContext(), viper.GetString(cmd.VpcId))
	if err != nil {
		formatter.PrintFatal("Couldn't receive VPN data", err)
	}

	if err = formatter.PrintList(vpns); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
