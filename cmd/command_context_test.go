package cmd

import (
	"github.com/ingrammicro/cio/internal/testutils"
	"github.com/spf13/cobra"
	"testing"
)

func TestCreateCommand(t *testing.T) {
	tests := map[string]struct {
		commandContext *CommandContext
	}{
		"if command created as expected": {
			commandContext: &CommandContext{
				Use:          testutils.TEST,
				FlagContexts: []FlagContext{{Type: String, Name: testutils.TEST, Usage: "Test command"}},
			},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cmd := createCommand(test.commandContext)
			if cmd.Use != test.commandContext.Use {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cmd.Use, test.commandContext.Use)
			}
		})
	}
}

func TestNewCommand(t *testing.T) {
	tests := map[string]struct {
		parentCommand     *cobra.Command
		commandContext    *CommandContext
		childCommandFound bool
	}{
		"if parent command is nil": {
			parentCommand:  nil,
			commandContext: &CommandContext{Use: "parent"},
		},
		"if parent command not found and child command not found": {
			parentCommand:  NewCommand(nil, &CommandContext{Use: "parent"}),
			commandContext: &CommandContext{Use: "child"},
		},
		"if parent command not found and child command found": {
			parentCommand:     NewCommand(nil, &CommandContext{Use: "parent"}),
			commandContext:    &CommandContext{Use: "child"},
			childCommandFound: true,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cmd := NewCommand(test.parentCommand, test.commandContext)
			if test.childCommandFound {
				cmd = NewCommand(test.parentCommand, test.commandContext)
			}
			if cmd.Use != test.commandContext.Use {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cmd.Use, test.commandContext.Use)
			}
		})
	}
}

func TestInitCommandFlags(t *testing.T) {
	tests := map[string]struct {
		persistentFlag bool
		flagContext    FlagContext
		expected       bool
	}{
		"if PersistentFlag is defined": {
			persistentFlag: true,
		},
		"if FlagContext type is string": {
			flagContext: FlagContext{Type: String, Name: testutils.TEST, Usage: "Test flag as string", DefaultValue: testutils.TEST},
			expected:    true,
		},
		"if FlagContext type is int or int64": {
			flagContext: FlagContext{Type: Int, Name: testutils.TEST, Usage: "Test flag as integer", DefaultValue: 0},
			expected:    true,
		},
		"if FlagContext type is bool": {
			flagContext: FlagContext{Type: Bool, Name: testutils.TEST, Usage: "Test flag as bool", DefaultValue: true},
			expected:    true,
		},
		"if FlagContext type is not expected": {
			flagContext: FlagContext{Type: -1, Name: testutils.TEST, Usage: "Test flag not expected"},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			commandFlags := initCommandFlags(
				NewCommand(nil, &CommandContext{Use: testutils.TEST}),
				test.flagContext,
				test.persistentFlag)
			if commandFlags.HasFlags() != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", commandFlags.HasFlags(), test.expected)
			}
		})
	}
}

func TestAttachFlags(t *testing.T) {
	tests := map[string]struct {
		persistentFlag bool
		flagContexts   []FlagContext
	}{
		"if PersistentFlag is defined": {
			persistentFlag: true,
			flagContexts:   []FlagContext{{Type: String, Name: testutils.TEST, Usage: "Test command"}},
		},
		"if PersistentFlag is not defined": {
			flagContexts: []FlagContext{{Type: String, Name: testutils.TEST, Usage: "Test command"}},
		},
		"if PersistentFlag is not defined and flagContext is required": {
			flagContexts: []FlagContext{{Type: String, Name: testutils.TEST, Usage: "Test command", Required: true}},
		},
		"if PersistentFlag is not defined and flagContext is hidden": {
			flagContexts: []FlagContext{{Type: String, Name: testutils.TEST, Usage: "Test command", Hidden: true}},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cmd := NewCommand(nil, &CommandContext{Use: testutils.TEST})
			attachFlags(cmd, test.flagContexts, test.persistentFlag)
			if test.persistentFlag && !cmd.HasPersistentFlags() {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cmd.HasPersistentFlags(), test.persistentFlag)
			}
			if !test.persistentFlag && !cmd.HasFlags() {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cmd.HasFlags(), test.persistentFlag)
			}
		})
	}
}

func TestNameAndAliasesPadding(t *testing.T) {
	tests := map[string]struct {
		childCommandContext *CommandContext
		expected            int
	}{
		"if padding is for all commands name": {
			childCommandContext: &CommandContext{Use: "child"},
			expected:            5,
		},
		"if padding is for all commands name and aliases": {
			childCommandContext: &CommandContext{Use: "child", Aliases: []string{"c"}},
			expected:            8,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			parentCmd := NewCommand(nil, &CommandContext{})
			childCmd := NewCommand(parentCmd, test.childCommandContext)
			commandsMaxNameLen := nameAndAliasesPadding(parentCmd)
			if commandsMaxNameLen != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", commandsMaxNameLen, len(childCmd.Use))
			}
		})
	}
}

// TODO
func TestSetUsageTemplate(t *testing.T) {
	parentCmd := NewCommand(nil, &CommandContext{Use: testutils.TEST})
	SetUsageTemplate(parentCmd)
	t.Logf("*** RES: %v \n", parentCmd.UsageTemplate())
}
