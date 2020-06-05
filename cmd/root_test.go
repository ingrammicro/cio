package cmd

import (
	"github.com/ingrammicro/cio/configuration"
	"testing"
)

func TestRootCmd(t *testing.T) {
	tests := map[string]struct {
		mode configuration.Mode
	}{
		"if running as Client": {
			mode: configuration.Client,
		},
		//"if running as Agent": {
		//	mode: configuration.Server,
		//},
	}

	for title, test := range tests {
		rootCmd := NewCommand(nil, &CommandContext{
			Use:     "cio",
			Short:   "Manages communication between Host and " + configuration.CloudOrchestratorPlatformName + " Platform",
			Version: configuration.VERSION,
			//Ctx:     context.Background(),
		})
		rootCmd.PersistentPreRun = persistencePreRun
		// To tell Cobra not to provide the default completion command
		rootCmd.CompletionOptions.DisableDefaultCmd = true
		t.Run(title, func(t *testing.T) {
			Execute(test.mode)
		})
	}
}
