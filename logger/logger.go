// Copyright (c) 2017-2022 Ingram Micro Inc.

package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

// DebugFuncInfo writes context info about the calling function
func DebugFuncInfo() {
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
	log.Debugf(dbgMsg)
}
