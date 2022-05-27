// Copyright (c) 2017-2022 Ingram Micro Inc.

package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CommandContext struct {
	Name                   string
	Use                    string
	Short                  string
	Long                   string
	Aliases                []string
	FlagContexts           []FlagContext
	RunMethod              interface{}
	PersistentPreRunMethod interface{}
	Version                string
	//Ctx                    context.Context
}

func (c *CommandContext) AddFlag(f FlagContext) {
	f.cmd = c
	c.FlagContexts = append(c.FlagContexts, f)
}

func createCommand(commandContext *CommandContext) *cobra.Command {
	return &cobra.Command{
		Use:     commandContext.Use,
		Short:   commandContext.Short,
		Long:    commandContext.Long,
		Aliases: commandContext.Aliases,
		Version: commandContext.Version,
		Run: func(cmd *cobra.Command, args []string) {
			for _, flagContext := range commandContext.FlagContexts {
				// It seems that due to the overlap of the same name in different commands, we need to bind so that we
				// have the value later in viper.GetString
				viper.BindPFlag(flagContext.Name, cmd.Flags().Lookup(flagContext.Name))
			}

			// Run method
			if commandContext.RunMethod != nil {
				// Workaround focused on server mode methods, where no flag name but value (argument) received!
				nRequiredArgsIn := reflect.TypeOf(commandContext.RunMethod).NumIn()
				nReceivedArgsIn := len(args)

				in := []reflect.Value{reflect.ValueOf(args)}
				if nRequiredArgsIn == 0 {
					in = nil
				}
				if nRequiredArgsIn > 0 && nReceivedArgsIn < nRequiredArgsIn {
					log.Fatalf("Not enough parameters!")
				} else {
					reflect.ValueOf(commandContext.RunMethod).Call(in)
				}
			} else {
				//default, if no method defined
				cmd.Help()
			}
		},
	}
}

// NewCommand Builds a new cobra command to a parent one. Based on received context
func NewCommand(parentCommand *cobra.Command, commandContext *CommandContext) *cobra.Command {
	// COMMAND
	newCommand := createCommand(commandContext)

	// By default, deactivate sorting. Shown based on definition order
	newCommand.Flags().SortFlags = false

	// FLAGS
	attachFlags(newCommand, commandContext.FlagContexts, false)

	// HIERARCHY
	// Add command to parent -if required-
	if parentCommand != nil {
		bFound := false
		for _, c := range parentCommand.Commands() {
			if c.Name() == newCommand.Name() {
				bFound = true
				break
			}
		}
		if !bFound {
			parentCommand.AddCommand(newCommand)
		}
	}
	return newCommand
}

func initCommandFlags(command *cobra.Command, flagContext FlagContext, bPersistentFlag bool) *pflag.FlagSet {
	commandFlags := command.Flags()
	if bPersistentFlag {
		commandFlags = command.PersistentFlags()
	}
	switch flagContext.Type {
	case String:
		value := ""
		if flagContext.DefaultValue != nil {
			value = flagContext.DefaultValue.(string)
		}
		commandFlags.StringP(flagContext.Name, flagContext.Shorthand, value, flagContext.Usage)
	case Int, Int64:
		value := 0
		if flagContext.DefaultValue != nil {
			value = flagContext.DefaultValue.(int)
		}
		commandFlags.IntP(flagContext.Name, flagContext.Shorthand, value, flagContext.Usage)
	case Bool:
		value := false
		if flagContext.DefaultValue != nil {
			value = flagContext.DefaultValue.(bool)
		}
		commandFlags.BoolP(flagContext.Name, flagContext.Shorthand, value, flagContext.Usage)
	default:

	}
	return commandFlags
}

func attachFlags(command *cobra.Command, flags []FlagContext, bPersistentFlag bool) {
	// FLAGS
	for _, flagContext := range flags {
		commandFlags := initCommandFlags(command, flagContext, bPersistentFlag)

		if !bPersistentFlag {
			// if required
			if flagContext.Required {
				command.MarkFlagRequired(flagContext.Name)
			}

			// if hidden
			if flagContext.Hidden {
				commandFlags.Lookup(flagContext.Name).Hidden = true
			}

			// Bind flag to command
			viper.BindPFlag(flagContext.Name, commandFlags.Lookup(flagContext.Name))
		}
	}
}

// nameAndAliasesPadding returns padding for the name and all aliases
func nameAndAliasesPadding(cmd *cobra.Command) int {
	commandsMaxNameLen := 0
	for _, x := range cmd.Commands() {
		nameLen := len(strings.Join(append([]string{x.Name()}, x.Aliases...), ", "))
		if nameLen > commandsMaxNameLen {
			commandsMaxNameLen = nameLen
		}
	}
	return commandsMaxNameLen
}

// SetUsageTemplate sets usage template for the command
func SetUsageTemplate(command *cobra.Command) {
	s := `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .NameAndAliases %d }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
	// sets name and aliases padding
	command.SetUsageTemplate(fmt.Sprintf(s, nameAndAliasesPadding(command)))
}
