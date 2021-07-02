// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"

	"github.com/ingrammicro/cio/api/cloud"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpServerArray prepares common resources to send request to Concerto API
func WireUpServerArray(c *cli.Context) (ds *cloud.ServerArrayService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloud.NewServerArrayService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up server array service", err)
	}

	return ds, f
}

// ServerArrayList subcommand function
func ServerArrayList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	serverArrays, err := serverArraySvc.ListServerArrays()
	if err != nil {
		formatter.PrintFatal("Couldn't receive server array data", err)
	}

	labelables := make([]types.Labelable, len(serverArrays))
	for i := 0; i < len(serverArrays); i++ {
		labelables[i] = types.Labelable(serverArrays[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	serverArrays = make([]*types.ServerArray, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		sa, ok := labelable.(*types.ServerArray)
		if !ok {
			formatter.PrintFatal(LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.ServerArray, got a %T", labelable))
		}
		serverArrays[i] = sa
	}
	if err = formatter.PrintList(serverArrays); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayShow subcommand function
func ServerArrayShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverArray, err := serverArraySvc.GetServerArray(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive server array data", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayCreate subcommand function
func ServerArrayCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"name", "template-id", "cloud-account-id", "server-plan-id"}, formatter)
	serverArrayIn := map[string]interface{}{
		"name":             c.String("name"),
		"template_id":      c.String("template-id"),
		"cloud_account_id": c.String("cloud-account-id"),
		"server_plan_id":   c.String("server-plan-id"),
	}

	if c.IsSet("size") {
		serverArrayIn["size"] = c.Int("size")
	}
	if c.IsSet("firewall-profile-id") {
		serverArrayIn["firewall_profile_id"] = c.String("firewall-profile-id")
	}
	if c.IsSet("ssh-profile-id") {
		serverArrayIn["ssh_profile_id"] = c.String("ssh-profile-id")
	}
	if c.IsSet("subnet-id") {
		serverArrayIn["subnet_id"] = c.String("subnet-id")
	}
	if c.IsSet("privateness") {
		serverArrayIn["privateness"] = c.Bool("privateness")
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		serverArrayIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	serverArray, err := serverArraySvc.CreateServerArray(&serverArrayIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create server array", err)
	}

	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayUpdate subcommand function
func ServerArrayUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverArray, err := serverArraySvc.UpdateServerArray(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't update server array", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayBoot subcommand function
func ServerArrayBoot(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverArray, err := serverArraySvc.BootServerArray(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't boot server array", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayShutdown subcommand function
func ServerArrayShutdown(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverArray, err := serverArraySvc.ShutdownServerArray(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't shutdown server array", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayEmpty subcommand function
func ServerArrayEmpty(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverArray, err := serverArraySvc.EmptyServerArray(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't empty server array", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayEnlarge subcommand function
func ServerArrayEnlarge(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"id", "size"}, formatter)
	serverArrayEnlargeIn := map[string]interface{}{
		"size_increase": c.Int("size"),
	}
	serverArray, err := serverArraySvc.EnlargeServerArray(c.String("id"), &serverArrayEnlargeIn)
	if err != nil {
		formatter.PrintFatal("Couldn't enlarge server array", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	serverArray.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*serverArray); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayServerList subcommand function
func ServerArrayServerList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	servers, err := serverArraySvc.ListServerArrayServers(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive server data", err)
	}

	labelables := make([]types.Labelable, len(servers))
	for i := 0; i < len(servers); i++ {
		labelables[i] = types.Labelable(servers[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	servers = make([]*types.Server, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		s, ok := labelable.(*types.Server)
		if !ok {
			formatter.PrintFatal(LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.Server, got a %T", labelable))
		}
		servers[i] = s
	}
	if err = formatter.PrintList(servers); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerArrayDelete subcommand function
func ServerArrayDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverArraySvc, formatter := WireUpServerArray(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := serverArraySvc.DeleteServerArray(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete server array", err)
	}
	return nil
}
