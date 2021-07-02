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

// WireUpServer prepares common resources to send request to Concerto API
func WireUpServer(c *cli.Context) (ds *cloud.ServerService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloud.NewServerService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up server service", err)
	}

	return ds, f
}

// ServerList subcommand function
func ServerList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	servers, err := serverSvc.ListServers()
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

// ServerShow subcommand function
func ServerShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.GetServer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive server data", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerCreate subcommand function
func ServerCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(
		c,
		[]string{"name", "ssh-profile-id", "firewall-profile-id", "template-id", "server-plan-id", "cloud-account-id"},
		formatter,
	)
	serverIn := map[string]interface{}{
		"name":                c.String("name"),
		"ssh_profile_id":      c.String("ssh-profile-id"),
		"firewall_profile_id": c.String("firewall-profile-id"),
		"template_id":         c.String("template-id"),
		"server_plan_id":      c.String("server-plan-id"),
		"cloud_account_id":    c.String("cloud-account-id"),
	}

	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)

	if c.IsSet("labels") {
		serverIn["label_ids"] = LabelResolution(c, c.String("labels"), &labelNamesByID, &labelIDsByName)
	}

	server, err := serverSvc.CreateServer(&serverIn)
	if err != nil {
		formatter.PrintFatal("Couldn't create server", err)
	}

	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerUpdate subcommand function
func ServerUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.UpdateServer(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't update server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerBoot subcommand function
func ServerBoot(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.BootServer(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't boot server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerReboot subcommand function
func ServerReboot(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.RebootServer(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't reboot server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerShutdown subcommand function
func ServerShutdown(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.ShutdownServer(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't shutdown server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerOverride subcommand function
func ServerOverride(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.OverrideServer(c.String("id"), utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't override server", err)
	}

	_, labelNamesByID := LabelLoadsMapping(c)
	server.FillInLabelNames(labelNamesByID)
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerDelete subcommand function
func ServerDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := serverSvc.DeleteServer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete server", err)
	}
	return nil
}

// ServerFloatingIPList subcommand function
func ServerFloatingIPList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	floatingIPs, err := serverSvc.ListServerFloatingIPs(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive floating IPs data", err)
	}

	labelables := make([]types.Labelable, len(floatingIPs))
	for i := 0; i < len(floatingIPs); i++ {
		labelables[i] = types.Labelable(floatingIPs[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	floatingIPs = make([]*types.FloatingIP, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		fIP, ok := labelable.(*types.FloatingIP)
		if !ok {
			formatter.PrintFatal(LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.FloatingIP, got a %T", labelable))
		}
		floatingIPs[i] = fIP
	}
	if err = formatter.PrintList(floatingIPs); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ServerVolumesList subcommand function
func ServerVolumesList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	volumes, err := serverSvc.ListServerVolumes(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive volumes data", err)
	}

	labelables := make([]types.Labelable, len(volumes))
	for i := 0; i < len(volumes); i++ {
		labelables[i] = types.Labelable(volumes[i])
	}
	labelIDsByName, labelNamesByID := LabelLoadsMapping(c)
	filteredLabelables := LabelFiltering(c, labelables, labelIDsByName)
	LabelAssignNamesForIDs(c, filteredLabelables, labelNamesByID)

	volumes = make([]*types.Volume, len(filteredLabelables))
	for i, labelable := range filteredLabelables {
		v, ok := labelable.(*types.Volume)
		if !ok {
			formatter.PrintFatal(LabelFilteringUnexpected,
				fmt.Errorf("expected labelable to be a *types.Volume, got a %T", labelable))
		}
		volumes[i] = v
	}
	if err = formatter.PrintList(volumes); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// ========= Events ========

// EventsList subcommand function
func EventsList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	events, err := svc.ListEvents(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive event data", err)
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

//======= Operational Scripts ==========

// OperationalScriptsList subcommand function
func OperationalScriptsList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	svc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	scripts, err := svc.ListOperationalScripts(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive script data", err)
	}
	if err = formatter.PrintList(scripts); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// OperationalScriptExecute subcommand function
func OperationalScriptExecute(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"server-id", "script-id"}, formatter)
	in := &map[string]interface{}{}
	scriptOut, err := serverSvc.ExecuteOperationalScript(c.String("server-id"), c.String("script-id"), in)
	if err != nil {
		formatter.PrintFatal("Couldn't execute operational script", err)
	}
	if err = formatter.PrintItem(*scriptOut); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
