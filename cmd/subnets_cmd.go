package cmd

import (
	"github.com/ingrammicro/cio/api/network"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpSubnet prepares common resources to send request to Concerto API
func WireUpSubnet(c *cli.Context) (ds *network.SubnetService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewSubnetService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up Subnet service", err)
	}

	return ds, f
}

// SubnetList subcommand function
func SubnetList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	subnetSvc, formatter := WireUpSubnet(c)

	checkRequiredFlags(c, []string{"vpc-id"}, formatter)
	subnets, err := subnetSvc.ListSubnets(c.String("vpc-id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive Subnet data", err)
	}

	if err = formatter.PrintList(subnets); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SubnetShow subcommand function
func SubnetShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	subnetSvc, formatter := WireUpSubnet(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	subnet, err := subnetSvc.GetSubnet(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive Subnet data", err)
	}

	if err = formatter.PrintItem(*subnet); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SubnetCreate subcommand function
func SubnetCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	subnetSvc, formatter := WireUpSubnet(c)

	checkRequiredFlags(c, []string{"vpc-id", "name", "cidr", "type"}, formatter)

	subnetIn := map[string]interface{}{
		"name":             c.String("name"),
		"cidr":             c.String("cidr"),
		"cloud_account_id": c.String("cloud-account-id"),
		"type":             c.String("type"),
	}

	subnet, err := subnetSvc.CreateSubnet(c.String("vpc-id"), &subnetIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create Subnet", err)
	}

	if err = formatter.PrintItem(*subnet); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SubnetUpdate subcommand function
func SubnetUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	subnetSvc, formatter := WireUpSubnet(c)

	checkRequiredFlags(c, []string{"id", "name"}, formatter)

	subnetIn := map[string]interface{}{
		"name": c.String("name"),
	}

	subnet, err := subnetSvc.UpdateSubnet(c.String("id"), &subnetIn)
	if err != nil {
		formatter.PrintFatal("Couldn't update Subnet", err)
	}

	if err = formatter.PrintItem(*subnet); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SubnetDelete subcommand function
func SubnetDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	subnetSvc, formatter := WireUpSubnet(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := subnetSvc.DeleteSubnet(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete Subnet", err)
	}
	return nil
}

// SubnetServerList subcommand function
func SubnetServerList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	subnetSvc, formatter := WireUpSubnet(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	servers, err := subnetSvc.ListSubnetServers(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive servers data", err)
	}

	if err = formatter.PrintList(servers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SubnetServerArrayList subcommand function
func SubnetServerArrayList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	subnetSvc, formatter := WireUpSubnet(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverArrays, err := subnetSvc.ListSubnetServerArrays(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive server arrays data", err)
	}

	if err = formatter.PrintList(serverArrays); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
