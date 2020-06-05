// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"os"
	"strings"

	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const PrintFormatError = "Couldn't print/format result"
const LabelFilteringUnexpected = "Label filtering returned unexpected result"

//const CouldNotReceiveData  = "Couldn't receive %s data"
//formatter.PrintFatal(fmt.Sprintf(cmd.CouldNotReceiveData, "system event"), err)

// WireUpAPI prepares common resources to send request to Concerto API
func WireUpAPI() (ds *api.IMCOClient, config *configuration.Config, f format.Formatter) {
	f = format.GetFormatter()

	config, err := configuration.GetConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}

	ds, err = api.NewIMCOClient(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	return ds, config, f
}

// LoadCloudProvidersMapping retrieves Cloud Providers and create a map between ID and Name
func LoadCloudProvidersMapping() map[string]string {
	logger.DebugFuncInfo()
	svc, _, formatter := WireUpAPI()

	cloudProviders, err := svc.ListCloudProviders()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	cloudProvidersMap := make(map[string]string)
	for _, cloudProvider := range cloudProviders {
		cloudProvidersMap[cloudProvider.ID] = cloudProvider.Name
	}

	return cloudProvidersMap
}

// LoadLocationsMapping retrieves Locations and create a map between ID and Name
func LoadLocationsMapping() map[string]string {
	logger.DebugFuncInfo()
	svc, _, formatter := WireUpAPI()

	locations, err := svc.ListLocations()
	if err != nil {
		formatter.PrintFatal("Couldn't receive location data", err)
	}
	locationsMap := make(map[string]string)
	for _, location := range locations {
		locationsMap[location.ID] = location.Name
	}

	return locationsMap
}

func ShowCommand(cmd *cobra.Command, args []string) {
	command := []string{cmd.Name()}

	// Loop up parents, prepending name
	for p := cmd.Parent(); p != nil; p = p.Parent() {
		command = append([]string{p.Name()}, command...)
	}

	// For every set flag, append its name and value
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		command = append(command, "--"+flag.Name+"="+flag.Value.String())
	})

	// Append the args
	command = append(command, args...)
	log.Debug("CommandContext: ", strings.Join(command, " "))
}

func EvaluateDebug() {
	debug := viper.GetBool(configuration.Debug)
	if debug {
		if err := os.Setenv("DEBUG", "1"); err != nil {
			log.Fatalf("Couldn't set environment debug mode: %s", err)
		}
		log.SetOutput(os.Stderr)
		log.SetLevel(log.DebugLevel)
	}
}

func EvaluateFormatter() error {
	formatter := viper.GetString(configuration.Formatter)
	if formatter != "" {
		if formatter != "text" && formatter != "json" {
			log.Fatalf("Unrecognized formatter %s. Please, use one of [ text | json ]", formatter)
		}
		format.InitializeFormatter(formatter, os.Stdout)
	}
	return nil
}

func SetParamString(name string, flag string, paramsIn map[string]interface{}) {
	if viper.IsSet(flag) {
		paramsIn[name] = viper.GetString(flag)
	}
}

func SetParamInt(name string, flag string, paramsIn map[string]interface{}) {
	if viper.IsSet(flag) {
		paramsIn[name] = viper.GetInt(flag)
	}
}
