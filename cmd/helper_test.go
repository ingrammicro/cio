package cmd

import (
	"github.com/ingrammicro/cio/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"testing"
)

// TODO FatalError/osExit(1)
func TestWireUpAPI(t *testing.T) {
	tests := map[string]struct {
		config *configuration.Config
	}{
		"if running with config": {
			config: new(configuration.Config),
		},
		//"if running with no config": {
		//	config: nil,
		//},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.config != nil {
				c, err := configuration.InitializeConfig()
				if err != nil {
					t.Errorf("Unexpected error: %v\n", err)
				}
				c.APIEndpoint = "https://clients.test.imco.io/v3"
				c.Certificate.Cert = "../configuration/testdata/ssl/cert.crt"
				c.Certificate.Key = "../configuration/testdata/ssl/private/cert.key"
				c.Certificate.Ca = "../configuration/testdata/ssl/ca_cert.pem"

				test.config = c
			}
			configuration.SetConfig(test.config)
			svc, config, f := WireUpAPI()
			if svc != nil && f != nil && config != test.config {
				t.Errorf("Unexpected response: %v. Expected: %v\n", config, test.config)
			}
		})
	}
}

// TODO +flags / result assert?
func TestShowCommand(t *testing.T) {
	cmdChild := new(cobra.Command)
	tests := map[string]struct {
		parent *cobra.Command
	}{
		"if no command parent": {
			parent: nil,
		},
		"if command parent": {
			parent: new(cobra.Command),
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.parent != nil {
				test.parent.AddCommand(cmdChild)
			}

			ShowCommand(cmdChild, []string{})
		})
	}
}

// TODO FatalError/osExit(1)
func TestEvaluateDebug(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"if debug mode inactive": {
			expected: false,
		},
		"if debug mode active": {
			expected: true,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(configuration.Debug, test.expected)
			EvaluateDebug()
			if viper.GetBool(configuration.Debug) != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", viper.GetBool(configuration.Debug), test.expected)
			}
		})
	}
}

func TestEvaluateFormatter(t *testing.T) {
	tests := map[string]struct {
		expected string
	}{
		"if no formatter defined": {
			expected: "",
		},
		"if formatter is defined": {
			expected: "json",
		},
		"if formatter is defined but unexpected": {
			expected: "Unrecognized formatter test. Please, use one of [ text | json ]",
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(configuration.Formatter, test.expected)
			EvaluateFormatter()
			if viper.GetString(configuration.Formatter) != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", viper.GetString(configuration.Formatter), test.expected)
			}
		})
	}
}

func TestSetParamString(t *testing.T) {
	tests := map[string]struct {
		expected string
	}{
		"if flag defined": {
			expected: "json",
		},
		"if flag no defined": {
			expected: "",
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(configuration.Formatter, test.expected)

			paramsIn := map[string]any{}
			SetParamString(configuration.Formatter, configuration.Formatter, paramsIn)
			if paramsIn[configuration.Formatter] != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", paramsIn[configuration.Formatter], test.expected)
			}
		})
	}
}

func TestSetParamInt(t *testing.T) {
	tests := map[string]struct {
		expected int
	}{
		"if flag defined": {
			expected: 8080,
		},
		"if flag no defined": {
			expected: 0,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(Port, test.expected)

			paramsIn := map[string]any{}
			SetParamInt(Port, Port, paramsIn)
			if paramsIn[Port] != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", paramsIn[Port], test.expected)
			}
		})
	}
}
