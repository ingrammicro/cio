// Copyright (c) 2017-2022 Ingram Micro Inc.

package polling

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ingrammicro/cio/cmd/agent"

	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
	"github.com/spf13/viper"

	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
)

func init() {
	// flags defined with camel case style!: longTime, shortTime!!!
	fLongTime := cmd.FlagContext{
		Type:         cmd.Int64,
		Name:         cmd.LongTime,
		DefaultValue: DefaultPollingPingTimingIntervalLong,
		Shorthand:    "l",
		Usage:        "Polling ping long time interval (seconds)",
	}

	fShortTime := cmd.FlagContext{Type: cmd.Int64, Name: cmd.ShortTime,
		DefaultValue: DefaultPollingPingTimingIntervalShort, Shorthand: "s",
		Usage: "Polling ping short time interval (seconds)"}

	fTime := cmd.FlagContext{Type: cmd.Int, Name: cmd.Time, DefaultValue: DefaultThresholdTime, Shorthand: "t",
		Usage: "Maximum time -seconds- threshold per response chunk"}

	pollingCmd := cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use:   "polling",
		Short: "Manages polling commands"})
	cmd.NewCommand(pollingCmd, &cmd.CommandContext{
		Use:       "register",
		Short:     "Registers orchestrator agent within an imported host",
		RunMethod: agent.RegisterPolling},
	)
	cmd.NewCommand(pollingCmd, &cmd.CommandContext{
		Use:          "start",
		Short:        "Starts a polling routine to check and execute pending scripts",
		RunMethod:    Start,
		FlagContexts: []cmd.FlagContext{fLongTime, fShortTime}},
	)
	cmd.NewCommand(pollingCmd, &cmd.CommandContext{
		Use:       "stop",
		Short:     "Stops the running polling process",
		RunMethod: Stop},
	)
	cmd.NewCommand(pollingCmd, &cmd.CommandContext{
		Use:          "continuous-report-run",
		Short:        "Runs a script and gradually report its output",
		RunMethod:    ContinuousReportRun,
		FlagContexts: []cmd.FlagContext{fTime}},
	)
}

const (
	DefaultPollingPingTimingIntervalLong  = 30
	DefaultPollingPingTimingIntervalShort = 5
	ProcessIdFile                         = "cio-polling.pid"
)

// Handle signals
func handleSysSignals(cancelFunc context.CancelFunc) {
	logger.DebugFuncInfo()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	log.Debug("Ending, signal detected:", <-gracefulStop)
	cancelFunc()
}

// Returns the full path to the tmp folder joined with pid management file name
func getProcessIdFilePath() string {
	return strings.Join([]string{os.TempDir(), string(os.PathSeparator), ProcessIdFile}, "")
}

// Start starts the polling process
func Start() error {
	logger.DebugFuncInfo()

	formatter := format.GetFormatter()
	if err := agent.SetProcessIdToFile(getProcessIdFilePath()); err != nil {
		formatter.PrintError("cannot create the pid file", err)
		return err
	}

	pollingPingTimingIntervalLong := viper.GetInt64(cmd.LongTime)
	if pollingPingTimingIntervalLong <= 0 {
		pollingPingTimingIntervalLong = DefaultPollingPingTimingIntervalLong
	}
	log.Debug("Ping long time interval:", pollingPingTimingIntervalLong)

	pollingPingTimingIntervalShort := viper.GetInt64(cmd.ShortTime)
	if pollingPingTimingIntervalShort <= 0 {
		pollingPingTimingIntervalShort = DefaultPollingPingTimingIntervalShort
	}
	log.Debug("Ping short time interval:", pollingPingTimingIntervalShort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleSysSignals(cancel)

	pingRoutine(ctx, pollingPingTimingIntervalLong, pollingPingTimingIntervalShort)

	return nil
}

// Stop stops the polling process
func Stop() error {
	logger.DebugFuncInfo()

	formatter := format.GetFormatter()
	if err := agent.StopProcess(getProcessIdFilePath()); err != nil {
		formatter.PrintError("cannot stop the polling process", err)
		return err
	}

	log.Info("CIO polling successfully stopped")
	return nil
}

// Main polling background routine
func pingRoutine(ctx context.Context, longTimePeriod int64, shortTimePeriod int64) {
	logger.DebugFuncInfo()

	svc, _, formatter := agent.WireUpAPIServer()
	commandProcessed := make(chan bool, 1)

	// initialization
	isRunningCommandRoutine := false
	longTicker := time.NewTicker(time.Duration(longTimePeriod) * time.Second)
	currentTicker := longTicker
	for {
		log.Debug("Requesting for candidate commands status")
		ping, status, err := svc.Ping(ctx)
		if err != nil {
			formatter.PrintError("Couldn't receive polling ping data", err)
		} else {
			// One command is available, and no process running
			if status == 201 && ping.PendingCommands && !isRunningCommandRoutine {
				log.Debug("Detected a candidate command")
				isRunningCommandRoutine = true
				go processingCommandRoutine(ctx, svc, formatter, commandProcessed)
			}
		}

		log.Debug("Waiting...", currentTicker)

		select {
		case <-commandProcessed:
			isRunningCommandRoutine = false
			if currentTicker != longTicker {
				currentTicker.Stop()
			}
			log.Debug("Ticker assigned: short")
			currentTicker = time.NewTicker(time.Duration(shortTimePeriod) * time.Second)
		case <-currentTicker.C:
			if currentTicker != longTicker {
				currentTicker.Stop()
				log.Debug("Ticker assigned: Long")
				currentTicker = longTicker
			}
		case <-ctx.Done():
			log.Debug(ctx.Err())
			log.Debug("closing polling")
			return
		}
	}
}

// Subsidiary routine for commands processing
func processingCommandRoutine(
	ctx context.Context,
	svc *api.ServerAPI,
	formatter format.Formatter,
	commandProcessed chan bool,
) {
	logger.DebugFuncInfo()

	// 1. Request for the new command available
	log.Debug("Retrieving available command")
	command, status, err := svc.GetNextCommand(ctx)
	if err != nil {
		formatter.PrintError("Couldn't receive polling command candidate data", err)
	}

	// 2. Execute the retrieved command
	if status == 200 {
		log.Debug("Running the retrieved command")
		command.ExitCode, command.Stdout, command.Stderr, _, _ = agent.RunTracedCmd(command.Script)

		// 3. then status is propagated to orchestration platform
		log.Debug("Reporting command execution status")

		commandIn := map[string]interface{}{
			"id":        command.ID,
			"script":    command.Script,
			"stdout":    command.Stdout,
			"stderr":    command.Stderr,
			"exit_code": command.ExitCode,
		}

		_, status, err := svc.UpdateCommand(ctx, command.ID, &commandIn)
		if err != nil {
			formatter.PrintError("Couldn't send polling command report data", err)
		}

		if status == 200 {
			log.Debug("Command execution results successfully reported")
		} else {
			log.Error("Cannot report the command execution results")
		}
	} else {
		log.Error("Cannot retrieve the next command")
	}

	commandProcessed <- true
}
