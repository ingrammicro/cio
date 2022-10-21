// Copyright (c) 2017-2022 Ingram Micro Inc.

package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	log "github.com/sirupsen/logrus"

	"context"
	"github.com/ingrammicro/cio/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd *cobra.Command

var globalFlags []FlagContext

var serverFlags = []FlagContext{
	{
		Type:  String,
		Name:  configuration.ConcertoBrownfieldToken,
		Usage: "Concerto Brownfield Token [$CONCERTO_BROWNFIELD_TOKEN]",
	},
	{
		Type:  String,
		Name:  configuration.ConcertoCommandPollingToken,
		Usage: "Concerto CommandContext Polling Token [$CONCERTO_COMMAND_POLLING_TOKEN]",
	},
	{Type: String, Name: configuration.ConcertoServerId, Usage: "Concerto server ID [$CONCERTO_SERVER_ID]"},
}

var commonFlags = []FlagContext{
	{Type: String, Name: configuration.CaCert, Usage: "CA to verify remote connections [$CONCERTO_CA_CERT]"},
	{
		Type:  String,
		Name:  configuration.ClientCert,
		Usage: "Client cert to use for " + configuration.CloudOrchestratorPlatformName + " [$CONCERTO_CLIENT_CERT]",
	},
	{
		Type:  String,
		Name:  configuration.ClientKey,
		Usage: "Private key used in client " + configuration.CloudOrchestratorPlatformName + " auth [$CONCERTO_CLIENT_KEY]",
	},

	{Type: String, Name: configuration.ConcertoConfig, Usage: "Concerto Config File [$CONCERTO_CONFIG]"},
	{Type: String, Name: configuration.ConcertoEndpoint, Usage: "Concerto Endpoint [$CONCERTO_ENDPOINT]"},
	{Type: String, Name: configuration.ConcertoUrl, Usage: "Concerto Web URL [$CONCERTO_URL]"},
}

var outputFlags = []FlagContext{
	{
		Type:         String,
		Name:         configuration.Formatter,
		Usage:        "Output formatter [ text | json ] [$CONCERTO_FORMATTER]",
		Shorthand:    "f",
		DefaultValue: "text",
	},
	{Type: Bool, Name: configuration.Debug, Usage: "Enable debug mode", Shorthand: "D"},
}

// 1
func init() {
	// Disable command sorting
	cobra.EnableCommandSorting = false

	cobra.OnInitialize(initConfig)

	RootCmd = NewCommand(nil, &CommandContext{
		Use:     "cio",
		Short:   "Manages communication between Host and " + configuration.CloudOrchestratorPlatformName + " Platform",
		Version: configuration.VERSION,
		//Ctx:     context.Background(),
	})
	RootCmd.PersistentPreRun = persistencePreRun
	// To tell Cobra not to provide the default completion command
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

// 2

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute(mode configuration.Mode) {
	// - - - - - - - - - - - - - - - - - - - - - - - -
	// Based on 2 'main' model, cobra running steps (init > Execute > initConfig) and mode as parameter.
	// Removed from 'init' since at this moment we are not aware about Mode
	// Set flags based on contextual mode: Client-User / Server-Agent
	globalFlags = append(globalFlags, commonFlags...)
	if mode == configuration.Server {
		globalFlags = append(globalFlags, serverFlags...)
	}
	globalFlags = append(globalFlags, outputFlags...)
	attachFlags(RootCmd, globalFlags, true)
	// - - - - - - - - - - - - - - - - - - - - - - - -

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	interrupted := make(chan os.Signal, 1)
	signal.Notify(interrupted, os.Interrupt)
	defer func() {
		signal.Stop(interrupted)
		cancel()
	}()
	go func() {
		select {
		case <-interrupted:
			cancel()
		case <-ctx.Done():
		}
	}()

	var err error
	if RootCmd, err = RootCmd.ExecuteContextC(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 3
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix(configuration.ConcertoEnvVarPrefixName) // = env vars: "CONCERTO_xxx"
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// Binding on init....
	for _, f := range globalFlags {
		viper.BindPFlag(f.Name, RootCmd.Flags().Lookup(f.Name))
	}

	// sets usage template for the command
	SetUsageTemplate(RootCmd)
}

// 4
func persistencePreRun(cmd *cobra.Command, args []string) {
	// Evaluates if debug mode and apply if case
	EvaluateDebug()

	// Shows complete command
	ShowCommand(cmd, args)

	// Reads configuration
	_, err := configuration.InitializeConfig()
	if err != nil {
		log.Fatalf("Error reading configuration: %s", err)
	}

	// Formatter
	err = EvaluateFormatter()
	if err != nil {
		log.Fatalf("Error evaluating formatter: %s", err)
	}
}

// GetContext returns the current context assigned to RootCmd
func GetContext() context.Context {
	return RootCmd.Context()
}
