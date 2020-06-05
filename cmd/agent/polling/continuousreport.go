// Copyright (c) 2017-2022 Ingram Micro Inc.

package polling

import (
	"errors"
	"fmt"
	"github.com/ingrammicro/cio/cmd/agent"
	"os"
	"time"

	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

const (
	RetriesNumber        = 5
	DefaultThresholdTime = 10
)

// ContinuousReportRun Starts the continuous report run process
func ContinuousReportRun(cmdArgs []string) error {
	logger.DebugFuncInfo()

	svc, _, formatter := agent.WireUpAPIServer()

	// Workaround, receiving as argument
	if cmdArgs[0] == "" {
		formatter.PrintFatal("argument missing", errors.New("a script or command is required"))
	}
	cmdArg := cmdArgs[0]

	// cli command threshold flag
	thresholdTime := viper.GetInt(cmd.Time)
	if thresholdTime <= 0 {
		thresholdTime = DefaultThresholdTime
	}
	log.Debug("Time threshold:", thresholdTime)

	// Custom method for chunks processing
	fn := func(chunk string) error {
		log.Debug("sendChunks")
		err := agent.Retry(RetriesNumber, time.Second, func() error {
			log.Debug("Sending: ", chunk)

			commandIn := map[string]interface{}{
				"stdout": chunk,
			}

			_, statusCode, err := svc.ReportBootstrapLog(cmd.GetContext(), &commandIn)
			switch {
			// 0<100 error cases??
			case statusCode == 0:
				return fmt.Errorf("communication error %v %v", statusCode, err)
			case statusCode >= 500:
				return fmt.Errorf("server error %v %v", statusCode, err)
			case statusCode >= 400:
				return fmt.Errorf("client error %v %v", statusCode, err)
			default:
				return nil
			}
		})

		if err != nil {
			return fmt.Errorf("cannot send the chunk data, %v", err)
		}
		return nil
	}

	exitCode, err := agent.RunContinuousCmd(fn, cmdArg, thresholdTime, -1)
	if err != nil {
		formatter.PrintFatal("cannot process continuous report command", err)
	}

	log.Info("completed: ", exitCode)
	os.Exit(exitCode)
	return nil
}
