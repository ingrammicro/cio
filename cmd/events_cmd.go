// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"github.com/ingrammicro/cio/api/audit"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpEvent prepares common resources to send request to Concerto API
func WireUpEvent(c *cli.Context) (ns *audit.EventService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = audit.NewEventService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up event service", err)
	}

	return ns, f
}

// EventList subcommand function
func EventList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	eventSvc, formatter := WireUpEvent(c)

	events, err := eventSvc.ListEvents()
	if err != nil {
		formatter.PrintFatal("Couldn't receive event data", err)
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// SysEventList subcommand function
func SysEventList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	eventSvc, formatter := WireUpEvent(c)

	events, err := eventSvc.ListSysEvents()
	if err != nil {
		formatter.PrintFatal("Couldn't receive system event data", err)
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}
