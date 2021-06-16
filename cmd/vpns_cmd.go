// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"strings"

	"github.com/ingrammicro/cio/api/network"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpVPN prepares common resources to send request to Concerto API
func WireUpVPN(c *cli.Context) (ds *network.VPNService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewVPNService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up VPN service", err)
	}

	return ds, f
}

// VPNShow subcommand function
func VPNShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpnSvc, formatter := WireUpVPN(c)

	checkRequiredFlags(c, []string{"vpc-id"}, formatter)
	vpn, err := vpnSvc.GetVPN(c.String("vpc-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive VPN data", err)
	}
	if err = formatter.PrintItem(*vpn); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// VPNCreate subcommand function
func VPNCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpnSvc, formatter := WireUpVPN(c)

	checkRequiredFlags(c, []string{"vpc-id", "public-ip", "psk", "exposed-cidrs", "vpn-plan-id"}, formatter)

	vpnIn := map[string]interface{}{
		"public_ip":     c.String("public-ip"),
		"psk":           c.String("psk"),
		"exposed_cidrs": utils.RemoveDuplicates(strings.Split(c.String("exposed-cidrs"), ",")),
		"vpn_plan_id":   c.String("vpn-plan-id"),
	}

	vpn, err := vpnSvc.CreateVPN(c.String("vpc-id"), &vpnIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create VPN", err)
	}

	if err = formatter.PrintItem(*vpn); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// VPNDelete subcommand function
func VPNDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpnSvc, formatter := WireUpVPN(c)

	checkRequiredFlags(c, []string{"vpc-id"}, formatter)
	err := vpnSvc.DeleteVPN(c.String("vpc-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete VPN", err)
	}
	return nil
}

// VPNPlanList subcommand function
func VPNPlanList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	vpcSvc, formatter := WireUpVPN(c)
	checkRequiredFlags(c, []string{"vpc-id"}, formatter)

	vpns, err := vpcSvc.ListVPNPlans(c.String("vpc-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive VPN data", err)
	}

	if err = formatter.PrintList(vpns); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
