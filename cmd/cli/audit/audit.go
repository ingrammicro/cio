// Copyright (c) 2017-2022 Ingram Micro Inc.

package audit

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
)

func init() {
	auditCmd := cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:     "events",
		Short:   "Events allow the user to track their actions and the state of their servers",
		Aliases: []string{"ev"}},
	)
	cmd.NewCommand(auditCmd, &cmd.CommandContext{
		Use:       "list-events",
		Short:     "Returns information about the events related to the account group",
		RunMethod: EventList},
	)
	cmd.NewCommand(auditCmd, &cmd.CommandContext{
		Use:       "list-system-events",
		Short:     "Returns information about system-wide events",
		RunMethod: SystemEventList},
	)
}

// EventList function
func EventList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	events, err := svc.ListEvents(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive event data", err)
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}

// SystemEventList function
func SystemEventList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cli.WireUpAPIClient()

	events, err := svc.ListSysEvents(cmd.GetContext())
	if err != nil {
		formatter.PrintFatal("Couldn't receive system event data", err)
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
