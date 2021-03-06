// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const PrintFormatError = "Couldn't print/format result"
const LabelFilteringUnexpected = "Label filtering returned unexpected result"

// debugCmdFuncInfo writes context info about the calling function
func debugCmdFuncInfo(c *cli.Context) {
	if log.GetLevel() < log.DebugLevel {
		return
	}

	// get function name
	dbgMsg := ""
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		dbgMsg = runtime.FuncForPC(pc).Name()
		i := strings.LastIndex(dbgMsg, "/")
		if i != -1 {
			dbgMsg = dbgMsg[i+1:]
		}
	} else {
		dbgMsg = "<unknown function name>"
	}
	dbgMsg = fmt.Sprintf("func %s", dbgMsg)

	// get used flags
	for _, flag := range c.FlagNames() {
		dbgMsg = fmt.Sprintf("%s\n\t%s=%+v", dbgMsg, flag, c.Generic(flag))
	}
	log.Debugf(dbgMsg)
}

// checkRequiredFlags checks for required flags, and show usage if requirements not met
func checkRequiredFlags(c *cli.Context, flags []string, f format.Formatter) {
	missing := ""
	for _, flag := range flags {
		if !c.IsSet(flag) {
			missing = fmt.Sprintf("%s\n\t--%s", missing, flag)
		}
	}

	if missing != "" {
		f.PrintError("Incorrect usage.", fmt.Errorf("Mandatory parameters missing: %s\n", missing))
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(2)
	}
}

// checkRequiredFlagsOr checks that at least one of required flags is present, and show usage if requirements not met
func checkRequiredFlagsOr(c *cli.Context, flags []string, f format.Formatter) {
	missing := ""
	for _, flag := range flags {
		if c.IsSet(flag) {
			return
		}
		missing = fmt.Sprintf("%s\n\t--%s", missing, flag)
	}

	f.PrintError("Incorrect usage.", fmt.Errorf("Please use one of these parameters: %s\n", missing))
	cli.ShowCommandHelp(c, c.Command.Name)
	os.Exit(2)
}

func setParamString(c *cli.Context, name string, flag string, paramsIn map[string]interface{}) {
	if c.IsSet(flag) {
		paramsIn[name] = c.String(flag)
	}
}

func setParamInt(c *cli.Context, name string, flag string, paramsIn map[string]interface{}) {
	if c.IsSet(flag) {
		paramsIn[name] = c.Int(flag)
	}
}
