// Copyright (c) 2017-2022 Ingram Micro Inc.

package agent

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/utils/format"
	log "github.com/sirupsen/logrus"

	"github.com/ingrammicro/cio/configuration"

	"text/template"

	//_ "github.com/ingrammicro/cio/cmd/agent/bootstrapping"
	//_ "github.com/ingrammicro/cio/cmd/agent/brownfield"
	//_ "github.com/ingrammicro/cio/cmd/agent/converge"
	//_ "github.com/ingrammicro/cio/cmd/agent/dispatcher"
	//_ "github.com/ingrammicro/cio/cmd/agent/firewall"
	//_ "github.com/ingrammicro/cio/cmd/agent/polling"
	//_ "github.com/ingrammicro/cio/cmd/agent/secret"
)

var configFileTemplate = template.Must(template.New("configFile").Parse(
	`<concerto version="1.0" server="{{.APIEndpoint}}" log_file="{{.LogFile}}" log_level="{{.LogLevel}}">
<ssl cert="{{.CertPath}}" key="{{.KeyPath}}" server_ca="{{.CaCertPath}}" />
</concerto>
`))

// RegisterBrownfield registration entry point for the brownfield process
func RegisterBrownfield() {
	Register(configuration.Brownfield)
}

// RegisterPolling registration entry point for the polling process
func RegisterPolling() {
	Register(configuration.Polling)
}

// Register registers the brownfield/polling process
func Register(context configuration.Context) {
	log.Info("Register")
	f := format.GetFormatter()
	config, err := configuration.GetConfig()
	if err != nil {
		f.PrintFatal("Couldn't read config", err)
	}
	if !config.CurrentUserIsAdmin {
		if runtime.GOOS == "windows" {
			f.PrintFatal("Must run as administrator user", fmt.Errorf("running as non-administrator user"))
		} else {
			f.PrintFatal("Must run as super-user", fmt.Errorf("running as non-administrator user"))
		}
	}
	rootCACert, cert, key, err := obtainServerKeys(config, context)
	if err != nil {
		f.PrintFatal("Couldn't obtain server keys", err)
	}
	err = configureServerKeys(config, rootCACert, cert, key)
	if err != nil {
		f.PrintFatal("Couldn't configure server keys", err)
	}
	fmt.Printf("CIO agent successfully registered, configuration file placed at %s\n", config.ConfFile)
}

func obtainServerKeys(config *configuration.Config, context configuration.Context) (
	rootCACert string, cert string, key string, err error,
) {
	cs, err := api.NewIMCOServerWithToken(config, context)
	if err != nil {
		return
	}

	payload := make(map[string]interface{})
	var status int
	responseData := make(map[string]interface{})

	if context == configuration.Brownfield {
		responseData, status, err = cs.ObtainBrownfieldSslProfile(cmd.GetContext(), &payload)
	}
	if context == configuration.Polling {
		responseData, status, err = cs.ObtainPollingApiKey(cmd.GetContext(), &payload)
	}
	if err != nil {
		return
	}

	if status == 403 {
		err = fmt.Errorf("server responded with 403 code: the token is not valid, maybe it expired")
		return
	}
	if status >= 300 {
		err = fmt.Errorf("server responded with %d code: %s", status, responseData)
		return
	}

	iRootCACert, ok := responseData["root_ca_cert"]
	if !ok {
		err = fmt.Errorf("server response did not include root CA cert: %v", responseData)
		return
	}
	rootCACert, ok = iRootCACert.(string)
	if !ok {
		err = fmt.Errorf("server response returned a %T as root CA cert, expected a string", iRootCACert)
		return
	}
	iCert, ok := responseData["cert"]
	if !ok {
		err = fmt.Errorf("server response did not include server cert: %v", responseData)
		return
	}
	cert, ok = iCert.(string)
	if !ok {
		err = fmt.Errorf("server response returned a %T as server cert, expected a string", iCert)
		return
	}
	iKey, ok := responseData["key"]
	if !ok {
		err = fmt.Errorf("server response did not include server private key: %v", responseData)
		return
	}
	key, ok = iKey.(string)
	if !ok {
		err = fmt.Errorf("server response returned a %T as server private key, expected a string", iKey)
	}
	return
}

func configureServerKeys(config *configuration.Config, rootCACert, cert, key string) error {
	configFileData := &struct {
		APIEndpoint string
		LogFile     string
		LogLevel    string
		CertPath    string
		KeyPath     string
		CaCertPath  string
	}{config.APIEndpoint, config.LogFile, config.LogLevel,
		config.Certificate.Cert, config.Certificate.Key, config.Certificate.Ca}

	if configFileData.LogLevel == "" {
		configFileData.LogLevel = "info"
	}
	if configFileData.LogFile == "" {
		configFileData.LogFile = configuration.GetDefaultLogFilePath()
	}
	if configFileData.CaCertPath == "" {
		configFileData.CaCertPath = configuration.GetDefaultCaCertFilePath()
	}
	if configFileData.CertPath == "" {
		configFileData.CertPath = configuration.GetDefaultCertFilePath()
	}
	if configFileData.KeyPath == "" {
		configFileData.KeyPath = configuration.GetDefaultKeyFilePath()
	}
	err := os.MkdirAll(filepath.Dir(configFileData.CaCertPath), 0644)
	if err != nil {
		return fmt.Errorf("cannot create directory to place root CA cert: %v", err)
	}
	err = ioutil.WriteFile(configFileData.CaCertPath, []byte(rootCACert), 0644)
	if err != nil {
		return fmt.Errorf("cannot write root CA cert: %v", err)
	}
	err = os.MkdirAll(filepath.Dir(configFileData.CertPath), 0644)
	if err != nil {
		return fmt.Errorf("cannot create directory to place server cert: %v", err)
	}
	err = ioutil.WriteFile(configFileData.CertPath, []byte(cert), 0644)
	if err != nil {
		return fmt.Errorf("cannot write server cert: %v", err)
	}
	err = os.MkdirAll(filepath.Dir(configFileData.KeyPath), 0644)
	if err != nil {
		return fmt.Errorf("cannot create directory to place server key: %v", err)
	}
	err = ioutil.WriteFile(configFileData.KeyPath, []byte(key), 0600)
	if err != nil {
		return fmt.Errorf("cannot write server key: %v", err)
	}
	err = os.MkdirAll(config.ConfLocation, 0644)
	if err != nil {
		return fmt.Errorf("cannot create directory to place config file: %v", err)
	}
	f, err := os.OpenFile(config.ConfFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("could not open config file for writing: %v", err)
	}
	defer f.Close()
	err = configFileTemplate.Execute(f, configFileData)
	if err != nil {
		return fmt.Errorf("could not generate config file contents: %v", err)
	}
	return nil
}

// WireUpAPIServer prepares common resources to send request to Orchestrator API
func WireUpAPIServer() (svc *api.ServerAPI, config *configuration.Config, f format.Formatter) {
	ds, config, f := cmd.WireUpAPI()
	svc = new(api.ServerAPI)
	svc.HTTPClient = *ds
	return svc, config, f
}