package configuration

import (
	"bytes"
	"encoding/xml"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

// TODO COMMON
const TEST = "test"

// TODO COMMON
func InitConfig() *Config {
	config := new(Config)
	config.XMLName = xml.Name{
		Space: "",
		Local: "",
	}
	config.APIEndpoint = "https://clients.test.imco.io/v3"
	config.LogFile = "/var/log/concerto-client.log"
	config.LogLevel = "info"
	config.Certificate = Cert{
		Cert: "testdata/ssl/cert.crt",
		Key:  "testdata/ssl/private/cert.key",
		Ca:   "testdata/ssl/ca_cert.pem",
	}
	config.BootstrapConfig = BootstrapConfig{
		IntervalSeconds:      600,
		SplaySeconds:         300,
		ApplyAfterIterations: 4,
		RunOnce:              false,
	}
	config.ConfLocation = TEST
	config.ConfFile = "testdata/client.xml"
	config.ConfFileLastLoadedAt = time.Now()
	config.IsHost = false
	config.ConcertoURL = TEST
	config.BrownfieldToken = ""
	config.CommandPollingToken = ""
	config.ServerID = TEST
	config.CurrentUserName = TEST
	config.CurrentUserIsAdmin = false
	return config
}

// TODO
func TestInitializeConfig(t *testing.T) {
	//cachedConfig = InitConfig()

	_, err := InitializeConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}
}

func TestGetConfig(t *testing.T) {
	tests := map[string]struct {
		config *Config
	}{
		"if config initialized": {
			config: InitConfig(),
		},
		"if config not initialized": {
			config: nil,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cachedConfig = test.config
			config, err := GetConfig()
			if err != nil && test.config != nil {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && config != cachedConfig {
				t.Errorf("Unexpected response: %v. Expected: %v\n", config, cachedConfig)
			}
		})
	}
}

func TestReloadConfig(t *testing.T) {
	cachedConfig = InitConfig()

	tests := map[string]struct {
		confFile             string
		confFileLastLoadedAt time.Time
		expected             any
	}{
		"if ConfFile is accessible": {
			confFile:             cachedConfig.ConfFile,
			confFileLastLoadedAt: time.Now().Add(-time.Hour * 24),
			expected:             false,
		},
		"if ConfFile is not accessible": {
			confFile: TEST,
			expected: false,
		},
		"if ConfFile has changed": {
			confFile: cachedConfig.ConfFile,
			expected: true,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cachedConfig.ConfFile = test.confFile
			if test.confFileLastLoadedAt.IsZero() {
				cachedConfig.ConfFileLastLoadedAt = test.confFileLastLoadedAt
			}

			config, configUpdated, err := ReloadConfig()
			t.Logf("ReloadConfig: %v %v %v\n", config, configUpdated, err)
			if err != nil && configUpdated != test.expected {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(config, cachedConfig)) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", config, cachedConfig)
			}
		})
	}
}

func TestDebugShowConfig(t *testing.T) {
	tests := map[string]struct {
		config   *Config
		logLevel log.Level
		expected string
	}{
		"if the log level is below the debug level": {
			config:   nil,
			logLevel: log.InfoLevel,
			expected: "",
		},
		"if config is not null": {
			config:   InitConfig(),
			logLevel: log.TraceLevel,
			expected: "LogLevel=info",
		},
		"if config is null": {
			config:   nil,
			logLevel: log.DebugLevel,
			expected: "CIO configuration not loaded",
		},
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			log.SetLevel(test.logLevel)
			cachedConfig = test.config
			debugShowConfig()
			str := buf.String()
			if !strings.Contains(str, test.expected) {
				t.Errorf("Unexpected debug message: %v. Expected: %v\n", str, test.expected)
			}
		})
	}
}

func TestIsAgentMode(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"if working as server mode": {
			expected: true,
		},
		"if working as client mode": {
			expected: false,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cachedConfig.IsHost = test.expected
			bOk := cachedConfig.IsAgentMode()
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

func TestIsConfigReady(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"if config is ready": {
			expected: true,
		},
		"if config is not ready": {
			expected: false,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if !test.expected {
				cachedConfig.APIEndpoint = ""
				cachedConfig.Certificate.Cert = ""
				cachedConfig.Certificate.Key = ""
				cachedConfig.Certificate.Ca = ""
			} else {
				cachedConfig.APIEndpoint = TEST
				cachedConfig.Certificate.Cert = TEST
				cachedConfig.Certificate.Key = TEST
				cachedConfig.Certificate.Ca = TEST
			}
			bOk := cachedConfig.IsConfigReady()
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

func TestIsConfigReadySetup(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"if config is ready to use setup command": {
			expected: true,
		},
		"if config is not ready to use setup command": {
			expected: false,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if !test.expected {
				cachedConfig.ConcertoURL = ""
			} else {
				cachedConfig.ConcertoURL = TEST
			}
			bOk := cachedConfig.IsConfigReadySetup()
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

func TestIsConfigReadyBrownfield(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"if config is ready for brownfield token authentication": {
			expected: true,
		},
		"if config is not ready for brownfield token authentication": {
			expected: false,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if !test.expected {
				cachedConfig.APIEndpoint = ""
				cachedConfig.BrownfieldToken = ""
			} else {
				cachedConfig.APIEndpoint = TEST
				cachedConfig.BrownfieldToken = TEST
			}
			bOk := cachedConfig.IsConfigReadyBrownfield()
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

func TestIsConfigReadyCommandPolling(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"if config is ready for polling token authentication": {
			expected: true,
		},
		"if config is not ready for polling token authentication": {
			expected: false,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if !test.expected {
				cachedConfig.APIEndpoint = ""
				cachedConfig.CommandPollingToken = ""
				cachedConfig.ServerID = ""
			} else {
				cachedConfig.APIEndpoint = TEST
				cachedConfig.CommandPollingToken = TEST
				cachedConfig.ServerID = TEST
			}
			bOk := cachedConfig.IsConfigReadyCommandPolling()
			t.Logf("IsConfigReadyCommandPolling: %v\n", bOk)
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

func TestReadConfig(t *testing.T) {
	cachedConfig = InitConfig()

	tests := map[string]struct {
		confFile string
		EnvVar   string
		expected any
		permissions []uint32
	}{
		"if config file does no exists": {
			confFile: "",
		},
		"if config file exists": {
			confFile: cachedConfig.ConfFile,
		},
		"if config file exists but cannot open": {
			confFile:    "testdata/client_protected.xml",
			expected:    "permission denied",
			permissions: []uint32{0000, 0664},
		},
		"if config does not have valid XML format": {
			confFile: "testdata/client_bad_format.xml",
			expected: "does not have valid XML format",
		},
		"if ConcertoEndpoint ENV VAR defined": {
			confFile: cachedConfig.ConfFile,
			EnvVar:   ConcertoEndpoint,
		},
		"if ConcertoEndpoint ENV VAR empty": {
			confFile: cachedConfig.ConfFile,
			EnvVar:   "",
		},
		"if ClientCert ENV VAR defined": {
			confFile: cachedConfig.ConfFile,
			EnvVar:   ClientCert,
		},
		"if ClientKey ENV VAR defined": {
			confFile: cachedConfig.ConfFile,
			EnvVar:   ClientKey,
		},
		"if CaCert ENV VAR defined": {
			confFile: cachedConfig.ConfFile,
			EnvVar:   CaCert,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cachedConfig.ConfFile = test.confFile
			cachedConfig.APIEndpoint = ""
			if test.EnvVar != "" {
				viper.Set(test.EnvVar, TEST)
			}

			if len(test.permissions) > 0 {
				err := os.Chmod(test.confFile, os.FileMode(test.permissions[0]))
				if err != nil {
					t.Errorf("Error changing permission to file: %v", err)
				}
			}

			err := cachedConfig.readConfig()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}

			if len(test.permissions) > 0 {
				err := os.Chmod(test.confFile, os.FileMode(test.permissions[1]))
				if err != nil {
					t.Errorf("Error changing permission to file: %v", err)
				}
			}
		})
	}
}

// TODO
func TestEvaluateCurrentUser(t *testing.T) {
	cachedConfig = InitConfig()
	cachedConfig.evaluateCurrentUser()
}

func TestSetConfigFile(t *testing.T) {
	tests := map[string]struct {
		BrownfieldTokenDefined bool
		expected               string
	}{
		//"if is windows in Agent mode": {
		//	BrownfieldTokenDefined: true,
		//	expected:               WindowsServerConfigFile,
		//},
		//"if is windows in User mode": {
		//	BrownfieldTokenDefined: false,
		//	expected:               "c:\\\\Users\\\\User\\\\",
		//},
		"if is nix in Agent mode": {
			BrownfieldTokenDefined: true,
			expected:               NixServerConfigFile,
		},
		"if is nix in User mode": {
			BrownfieldTokenDefined: false,
			expected:               "/home/user/",
		},
	}

	cachedConfig = InitConfig()

	user, err := cachedConfig.evaluateCurrentUser()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			user.HomeDir = test.expected
			if test.BrownfieldTokenDefined {
				cachedConfig.CurrentUserIsAdmin = true
				cachedConfig.BrownfieldToken = TEST
			} else {
				cachedConfig.CurrentUserIsAdmin = false
				cachedConfig.BrownfieldToken = ""
				test.expected = filepath.Join(test.expected, ".concerto/client.xml")
			}

			cachedConfig.setConfigFile(user)
			if cachedConfig.ConfFile != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cachedConfig.ConfFile, test.expected)
			}
		})
	}
}

// TODO + errors
func TestEvaluateConfigFile(t *testing.T) {
	tests := map[string]struct {
		EnvVar bool
	}{
		"if defined ENV VAR ConcertoConfig": {
			EnvVar: true,
		},
		"if undefined ENV VAR ConcertoConfig": {
			EnvVar: false,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.EnvVar {
				viper.Set(ConcertoConfig, TEST)
			} else {
				viper.Set(ConcertoConfig, "")
			}
			err := cachedConfig.evaluateConfigFile()
			if err != nil {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestBrownfieldTokenDefined(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"if BrownfieldToken defined": {
			expected: true,
		},
		"if BrownfieldToken not defined": {
			expected: false,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if !test.expected {
				cachedConfig.BrownfieldToken = ""
			} else {
				cachedConfig.BrownfieldToken = TEST
			}
			bOk := cachedConfig.BrownfieldTokenDefined()
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

func TestPollingTokenAndServerIdDefined(t *testing.T) {
	tests := map[string]struct {
		expected bool
	}{
		"if CommandPollingToken and ServerId defined": {
			expected: true,
		},
		"if CommandPollingToken and ServerId not defined": {
			expected: false,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if !test.expected {
				cachedConfig.CommandPollingToken = ""
				cachedConfig.ServerID = ""
			} else {
				cachedConfig.CommandPollingToken = TEST
				cachedConfig.ServerID = TEST
			}
			bOk := cachedConfig.PollingTokenAndServerIdDefined()
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

// TODO
//func TestGetUsername(t *testing.T) {
//	user := getUsername()
//	if user == "unknown" {
//		t.Errorf("getUsername() = %s; want not unknown", user)
//	}
//	t.Logf("getUsername: %v\n", user)
//}

func TestIsWinAdministrator(t *testing.T) {
	tests := map[string]struct {
		user     string
		expected bool
	}{
		"if user is win Administrator": {
			user:     "Administrator",
			expected: true,
		},
		"if user is not win Administrator": {
			user:     "user",
			expected: false,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			bOk := isWinAdministrator(test.user)
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

// TODO case "PHYSICALDRIVE0" WINDOWSS!
func TestCanPerformAdministratorTasks(t *testing.T) {
	tests := map[string]struct {
		user     string
		expected bool
	}{
		//"if can perform administrator tasks": {
		//	user:     "Administrator",
		//	expected: true,
		//},
		"if cannot perform administrator tasks": {
			user:     "user",
			expected: false,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			bOk := canPerformAdministratorTasks()
			t.Logf("canPerformAdministratorTasks: %v\n", bOk)
			if bOk != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", bOk, test.expected)
			}
		})
	}
}

func TestReadConcertoURL(t *testing.T) {
	tests := map[string]struct {
		ConcertoURL string
		EnvVar      bool
		APIEndpoint string
	}{
		"if defined config ConcertoURL": {
			ConcertoURL: TEST,
			EnvVar:      false,
			APIEndpoint: "",
		},
		"if defined ENV VAR ConcertoURL": {
			ConcertoURL: "",
			EnvVar:      true,
			APIEndpoint: "",
		},
		"if undefined ENV VAR ConcertoURL and invalid APIEndpoint": {
			ConcertoURL: "",
			EnvVar:      false,
			APIEndpoint: "http://%28",
		},
		"if undefined ENV VAR ConcertoURL and valid APIEndpoint": {
			ConcertoURL: "",
			EnvVar:      false,
			APIEndpoint: "https://clients.test.imco.io/v3",
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cachedConfig.ConcertoURL = test.ConcertoURL
			cachedConfig.APIEndpoint = test.APIEndpoint
			if test.EnvVar {
				os.Setenv("CONCERTO_URL", TEST)
			} else {
				os.Unsetenv("CONCERTO_URL")
			}

			err := cachedConfig.readConcertoURL()
			if err != nil && test.EnvVar {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestReadBrownfieldToken(t *testing.T) {
	tests := map[string]struct {
		BrownfieldToken string
		EnvVar          string
		expected        string
	}{
		"if defined config BrownfieldToken": {
			BrownfieldToken: TEST,
			EnvVar:          "",
			expected:        TEST,
		},
		"if undefined config BrownfieldToken and defined ENV VAR ConcertoBrownfieldToken": {
			BrownfieldToken: "",
			EnvVar:          ConcertoBrownfieldToken,
			expected:        TEST,
		},
	}

	cachedConfig = InitConfig()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cachedConfig.BrownfieldToken = test.BrownfieldToken
			if test.EnvVar != "" {
				viper.Set(test.EnvVar, TEST)
			}

			cachedConfig.readBrownfieldToken()
			if cachedConfig.BrownfieldToken != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cachedConfig.ConfFile, test.expected)
			}
		})
	}
}

func TestReadCommandPollingConfig(t *testing.T) {
	cachedConfig = InitConfig()

	tests := map[string]struct {
		CommandPollingToken string
		ServerID            string
		EnvVar              string
	}{
		"if defined config CommandPollingToken": {
			CommandPollingToken: TEST,
			ServerID:            "",
			EnvVar:              "",
		},
		"if undefined config CommandPollingToken and defined config ServerID": {
			CommandPollingToken: "",
			ServerID:            TEST,
			EnvVar:              "",
		},
		"if undefined config vars and defined ENV VAR CommandPollingToken": {
			CommandPollingToken: "",
			ServerID:            "",
			EnvVar:              ConcertoCommandPollingToken,
		},
		"if undefined config vars and defined ENV VAR ServerID": {
			CommandPollingToken: "",
			ServerID:            "",
			EnvVar:              ConcertoServerId,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cachedConfig.CommandPollingToken = test.CommandPollingToken
			cachedConfig.ServerID = test.ServerID
			if test.EnvVar != "" {
				viper.Set(test.EnvVar, TEST)
			}

			cachedConfig.readCommandPollingConfig()
			if cachedConfig.CommandPollingToken != "" && cachedConfig.CommandPollingToken != TEST {
				t.Errorf(TEST)
			}
			if cachedConfig.ServerID != "" && cachedConfig.ServerID != TEST {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cachedConfig.ServerID, TEST)
			}
		})
	}
}

// TODO
func TestEvaluateCertificate(t *testing.T) {
	cachedConfig = InitConfig()
	cachedConfig.evaluateCertificate()
}

func TestEvaluateAPIEndpointURL(t *testing.T) {
	cachedConfig = InitConfig()

	tests := map[string]struct {
		APIEndpoint string
	}{
		"if APIEndpoint invalid": {
			APIEndpoint: "http://%28",
		},
		"if APIEndpoint well-formed": {
			APIEndpoint: "https://clients.test.imco.io/v3",
		},
		"if APIEndpoint no path, no version": {
			APIEndpoint: "https://clients.test.imco.io",
		},
		"if APIEndpoint wrong version": {
			APIEndpoint: "https://clients.test.imco.io/v2",
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			cachedConfig.APIEndpoint = test.APIEndpoint
			err := cachedConfig.evaluateAPIEndpointURL()

			if err != nil {
				_, err = url.Parse(test.APIEndpoint)
				if err == nil {
					t.Errorf("Unexpected error: %v\n", err)
				}
			}
		})
	}
}

// TODO windows
func TestGetDefaultLogFilePath(t *testing.T) {
	tests := map[string]struct {
		expected string
	}{
		//"if GOOS is windows": {
		//	expected: WindowsServerLogFilePath,
		//},
		"if GOOS is nix": {
			expected: NixServerLogFilePath,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			filePath := GetDefaultLogFilePath()
			if filePath != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cachedConfig.ConfFile, test.expected)
			}
		})
	}
}

// TODO windows
func TestGetDefaultCaCertFilePath(t *testing.T) {
	tests := map[string]struct {
		expected string
	}{
		//"if GOOS is windows": {
		//	expected: WindowsServerCaCertPath,
		//},
		"if GOOS is nix": {
			expected: NixServerCaCertPath,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			filePath := GetDefaultCaCertFilePath()
			if filePath != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cachedConfig.ConfFile, test.expected)
			}
		})
	}
}

// TODO windows
func TestGetDefaultCertFilePath(t *testing.T) {
	tests := map[string]struct {
		expected string
	}{
		//"if GOOS is windows": {
		//	expected: WindowsServerCertPath,
		//},
		"if GOOS is nix": {
			expected: NixServerCertPath,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			filePath := GetDefaultCertFilePath()
			if filePath != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cachedConfig.ConfFile, test.expected)
			}
		})
	}
}

// TODO windows
func TestGetDefaultKeyFilePath(t *testing.T) {
	tests := map[string]struct {
		expected string
	}{
		//"if GOOS is windows": {
		//	expected: WindowsServerCertPath,
		//},
		"if GOOS is nix": {
			expected: NixServerKeyPath,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			filePath := GetDefaultKeyFilePath()
			if filePath != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cachedConfig.ConfFile, test.expected)
			}
		})
	}
}
