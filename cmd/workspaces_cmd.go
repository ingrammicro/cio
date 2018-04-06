package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/cloud"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpWorkspace prepares common resources to send request to Concerto API
func WireUpWorkspace(c *cli.Context) (ds *cloud.WorkspaceService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloud.NewWorkspaceService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up workspace service", err)
	}

	return ds, f
}

// WorkspaceList subcommand function
func WorkspaceList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	workspaces, err := workspaceSvc.GetWorkspaceList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive workspace data", err)
	}
	if err = formatter.PrintList(workspaces); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// WorkspaceShow subcommand function
func WorkspaceShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	workspace, err := workspaceSvc.GetWorkspace(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive workspace data", err)
	}
	if err = formatter.PrintItem(*workspace); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// WorkspaceCreate subcommand function
func WorkspaceCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"name", "ssh_profile_id", "firewall_profile_id"}, formatter)
	workspace, err := workspaceSvc.CreateWorkspace(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create workspace", err)
	}
	if err = formatter.PrintItem(*workspace); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// WorkspaceUpdate subcommand function
func WorkspaceUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	workspace, err := workspaceSvc.UpdateWorkspace(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update workspace", err)
	}
	if err = formatter.PrintItem(*workspace); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// WorkspaceDelete subcommand function
func WorkspaceDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := workspaceSvc.DeleteWorkspace(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete workspace", err)
	}
	return nil
}

// WorkspaceServerList subcommand function
func WorkspaceServerList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"workspace_id"}, formatter)
	workspaceServers, err := workspaceSvc.GetWorkspaceServerList(c.String("workspace_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list workspace records", err)
	}
	if err = formatter.PrintList(*workspaceServers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
