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
		"if running as Agent": {
			mode: configuration.Server,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			globalFlags = nil
			RootCmd.ResetFlags()
			Execute(test.mode)
		})
	}
}
