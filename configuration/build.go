// Copyright (c) 2017-2022 Ingram Micro Inc.

package configuration

import (
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// Config stores configuration file contents
type Config struct {
	XMLName              xml.Name        `xml:"concerto"`
	APIEndpoint          string          `xml:"server,attr"`
	LogFile              string          `xml:"log_file,attr"`
	LogLevel             string          `xml:"log_level,attr"`
	Certificate          Cert            `xml:"ssl"`
	BootstrapConfig      BootstrapConfig `xml:"bootstrap"`
	ConfLocation         string
	ConfFile             string
	ConfFileLastLoadedAt time.Time
	IsHost               bool
	ConcertoURL          string
	BrownfieldToken      string
	CommandPollingToken  string
	ServerID             string
	CurrentUserName      string
	CurrentUserIsAdmin   bool
}

// Cert stores cert files location
type Cert struct {
	Cert string `xml:"cert,attr"`
	Key  string `xml:"key,attr"`
	Ca   string `xml:"server_ca,attr"`
}

// BootstrapConfig stores configuration specific to the bootstrap command
type BootstrapConfig struct {
	IntervalSeconds      int  `xml:"interval,attr"`
	SplaySeconds         int  `xml:"splay,attr"`
	ApplyAfterIterations int  `xml:"apply_after_iterations,attr"`
	RunOnce              bool `xml:"run_once,attr"`
}

var cachedConfig *Config

// GetConfig returns configuration
func GetConfig() (*Config, error) {
	if cachedConfig == nil {
		return nil, fmt.Errorf("configuration hasn't been initialized")
	}
	return cachedConfig, nil
}

// SetConfig returns configuration
func SetConfig(config *Config) {
	cachedConfig = config
}

// InitializeConfig creates the configuration structure
func InitializeConfig() (*Config, error) {
	logger.DebugFuncInfo()

	if cachedConfig != nil {
		return cachedConfig, nil
	}

	cachedConfig = &Config{}

	cachedConfig.readBrownfieldToken()

	cachedConfig.readCommandPollingConfig()

	// where config file must me
	if err := cachedConfig.evaluateConfigFile(); err != nil {
		return nil, err
	}

	// read config contents
	log.Debugf("Reading configuration from %s", cachedConfig.ConfFile)
	if err := cachedConfig.readConfig(); err != nil {
		return nil, err
	}

	// add login URL. Needed for setup
	if err := cachedConfig.readConcertoURL(); err != nil {
		return nil, err
	}

	// check if isHost. Needed to show appropriate options
	if err := cachedConfig.evaluateCertificate(); err != nil {
		return nil, err
	}

	// evaluates API endpoint url
	if err := cachedConfig.evaluateAPIEndpointURL(); err != nil {
		return nil, err
	}

	debugShowConfig()
	return cachedConfig, nil
}

// ReloadConfig checks if the config file was modified and
// if so, attempts to reload it. It returns the resulting config
// (updated or not), whether an modification of the file happened,
// and any errors
func ReloadConfig() (*Config, bool, error) {
	fi, err := os.Stat(cachedConfig.ConfFile)
	if err != nil {
		log.Warnf("Could not stat config file %q to see if it changed: %v", cachedConfig.ConfFile, err)
		return cachedConfig, false, err
	}
	if fi.ModTime().After(cachedConfig.ConfFileLastLoadedAt) {
		log.Infof("Config file %q changed since last reading, reloading configuration...", cachedConfig.ConfFile)
		oldConfig := cachedConfig
		cachedConfig = nil
		_, err = InitializeConfig()
		if err != nil {
			log.Warnf("Could not load changes to config file %q: %v", cachedConfig.ConfFile, err)
			cachedConfig = oldConfig
		}
		return cachedConfig, true, err
	}
	return cachedConfig, false, nil
}

func debugShowConfig() {
	if log.GetLevel() < log.DebugLevel {
		return
	}
	if cachedConfig == nil {
		log.Debug("CIO configuration not loaded")
		return
	}
	debugStruct("", *cachedConfig)
}

// debugStruct iterates struct and show in debug console all items and subitems
func debugStruct(prefix string, item interface{}) {
	c := reflect.ValueOf(item)
	for i := 0; i < c.NumField(); i++ {
		if c.Type().Field(i).Type.String() != "xml.Name" && c.Field(i).CanInterface() {

			name := c.Type().Field(i).Name
			value := c.Field(i).Interface()

			// if value is struct, iterate with recursion
			if c.Type().Field(i).Type.Kind() == reflect.Struct {
				debugStruct(name, value)
			} else {
				if prefix != "" {
					name = fmt.Sprintf("%s.%s", prefix, name)
				}
				log.WithField(name, value).Debug("Configuration item")
			}
		}
	}
}

// IsAgentMode returns whether CLI is acting as server Or Client mode
func (config *Config) IsAgentMode() bool {
	bOk := config.IsHost || config.BrownfieldToken != "" || config.CommandPollingToken != ""
	if bOk {
		log.Debug("Working as server mode")
	} else {
		log.Debug("Working as client mode")
	}
	return bOk
}

// IsConfigReady returns whether configurations items are filled
func (config *Config) IsConfigReady() bool {
	if config.APIEndpoint == "" ||
		config.Certificate.Cert == "" ||
		config.Certificate.Key == "" ||
		config.Certificate.Ca == "" {
		return false
	}
	return true
}

// IsConfigReadySetup returns whether we can use setup command
func (config *Config) IsConfigReadySetup() bool {
	return config.ConcertoURL != ""
}

// IsConfigReadyBrownfield returns whether config is ready for brownfield token
// authentication
func (config *Config) IsConfigReadyBrownfield() bool {
	if config.APIEndpoint == "" ||
		config.BrownfieldToken == "" {
		return false
	}
	return true
}

// IsConfigReadyCommandPolling returns whether config is ready for polling token
// authentication
func (config *Config) IsConfigReadyCommandPolling() bool {
	if config.APIEndpoint == "" ||
		config.CommandPollingToken == "" ||
		config.ServerID == "" {
		return false
	}
	return true
}

// readConfig reads config file located at fileLocation
func (config *Config) readConfig() error {
	log.Debug("Reading Configuration")
	if utils.FileExists(config.ConfFile) {
		// file exists, read it's contents
		xmlFile, err := os.Open(config.ConfFile)
		if err != nil {
			return err
		}
		defer xmlFile.Close()
		readingTime := time.Now()
		b, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			return errors.Wrapf(err, "configuration File %s couldn't be read", config.ConfFile)
		}

		if err = xml.Unmarshal(b, &config); err != nil {
			return errors.Wrapf(err, "configuration File %s does not have valid XML format", config.ConfFile)
		}
		config.ConfFileLastLoadedAt = readingTime

	} else {
		log.Debugf("Configuration File %s does not exist. Reading environment variables", config.ConfFile)
	}

	// overwrite with environment/arguments vars
	overwEP := viper.GetString(ConcertoEndpoint)
	if overwEP != "" {
		log.Debug("Concerto APIEndpoint taken from env/args")
		config.APIEndpoint = overwEP
	}

	overwCert := viper.GetString(ClientCert)
	if overwCert != "" {
		log.Debug("Certificate path taken from env/args")
		config.Certificate.Cert = overwCert
	}

	overwKey := viper.GetString(ClientKey)
	if overwKey != "" {
		log.Debug("Certificate key path taken from env/args")
		config.Certificate.Key = overwKey
	}

	overwCa := viper.GetString(CaCert)
	if overwCa != "" {
		log.Debug("CA certificate path taken from env/args")
		config.Certificate.Ca = overwCa
	}

	// if endpoint empty set default
	// we can't set the default from flags, because it would overwrite config file
	if config.APIEndpoint == "" {
		config.APIEndpoint = DefaultEndpoint
	}
	return nil
}

func (config *Config) evaluateCurrentUser() (*user.User, error) {
	currUser, err := user.Current()
	if err != nil {
		log.Debugf("Couldn't use os.user to get user details: %s", err.Error())
		dir, err := homedir.Dir()
		if err != nil {
			return nil, fmt.Errorf("couldn't get home dir for current user: %s", err.Error())
		}
		currUser = &user.User{
			Username: getUsername(),
			HomeDir:  dir,
		}
	}
	if runtime.GOOS == "windows" {
		currUser.Username = currUser.Username[strings.LastIndex(currUser.Username, "\\")+1:]
		log.Debugf("Windows username is %s", currUser.Username)
		config.CurrentUserIsAdmin = currUser.Gid == "S-1-5-32-544" || isWinAdministrator(currUser.Username) ||
			canPerformAdministratorTasks()
	} else {
		config.CurrentUserIsAdmin = currUser.Uid == "0" || currUser.Username == "root"
	}
	config.CurrentUserName = currUser.Username
	return currUser, nil
}

func (config *Config) setConfigFile(currUser *user.User) {
	if runtime.GOOS == "windows" {
		if config.CurrentUserIsAdmin &&
			(config.BrownfieldTokenDefined() ||
				config.PollingTokenAndServerIdDefined() ||
				utils.FileExists(WindowsServerConfigFile)) {
			log.Debugf("Current user is administrator, setting config file as %s", WindowsServerConfigFile)
			config.ConfFile = WindowsServerConfigFile
		} else {
			// User mode Windows
			log.Debugf("Current user is regular user: %s", currUser.Username)
			config.ConfFile = filepath.Join(currUser.HomeDir, ".concerto/client.xml")
		}
	} else {
		// Server mode *nix
		if config.CurrentUserIsAdmin &&
			(config.BrownfieldTokenDefined() ||
				config.PollingTokenAndServerIdDefined() ||
				utils.FileExists(NixServerConfigFile)) {
			config.ConfFile = NixServerConfigFile
		} else {
			// User mode *nix
			config.ConfFile = filepath.Join(currUser.HomeDir, ".concerto/client.xml")
		}
	}
}

// evaluateConfigFile returns path to config file
func (config *Config) evaluateConfigFile() error {
	logger.DebugFuncInfo()
	currUser, err := config.evaluateCurrentUser()
	if err != nil {
		return err
	}
	if configFile := viper.GetString(ConcertoConfig); configFile != "" {
		log.Debug("Concerto configuration file location taken from env/args")
		config.ConfFile = configFile
	} else {
		config.setConfigFile(currUser)

	}
	config.ConfLocation = path.Dir(config.ConfFile)
	return nil
}

// BrownfieldTokenDefined returns whether brownfield token is defined
func (config *Config) BrownfieldTokenDefined() bool {
	return config.BrownfieldToken != ""
}

// PollingTokenAndServerIdDefined returns whether polling token and server are defined
func (config *Config) PollingTokenAndServerIdDefined() bool {
	return config.CommandPollingToken != "" && config.ServerID != ""
}

// getUsername gets username by env variable.
// os.user is dependant on cgo, so cross compiling won't work
func getUsername() string {
	logger.DebugFuncInfo()
	u := "unknown"
	osUser := ""

	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" || runtime.GOOS == "solaris" {
		osUser = os.Getenv("USER")
	}

	if runtime.GOOS == "windows" {
		osUser = os.Getenv("USERNAME")

		// remove domain
		osUser = osUser[strings.LastIndex(osUser, "\\")+1:]
		log.Debugf("Windows user has been transformed into %s", osUser)

		// HACK ugly ... if localized administrator, translate to administrator
		if isWinAdministrator(osUser) {
			osUser = "Administrator"
		}
	}

	if osUser != "" {
		u = osUser
	}
	return u
}

func isWinAdministrator(user string) bool {
	users := []string{
		"Järjestelmänvalvoja",
		"Administrateur",
		"Rendszergazda",
		"Administrador",
		"Администратор",
		"Administratör",
		"Administrator",
		"SYSTEM",
		"imco",
	}
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}

func canPerformAdministratorTasks() bool {
	f, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

// readConcertoURL reads URL from CONCERTO_URL environment or calculates using API URL
func (config *Config) readConcertoURL() error {
	if config.ConcertoURL != "" {
		return nil
	}

	if overwURL := os.Getenv("CONCERTO_URL"); overwURL != "" {
		config.ConcertoURL = overwURL
		log.Debug("Concerto URL taken from CONCERTO_URL")
		return nil
	}

	cURL, err := url.Parse(config.APIEndpoint)
	if err != nil {
		return err
	}

	tokenHost := strings.Split(cURL.Host, ":")
	tokenFqdn := strings.Split(tokenHost[0], ".")

	if !strings.Contains(cURL.Host, "staging") {
		tokenFqdn[0] = "start"
	}

	config.ConcertoURL = fmt.Sprintf("%s://%s/", cURL.Scheme, strings.Join(tokenFqdn, "."))
	return nil
}

func (config *Config) readBrownfieldToken() {
	if config.BrownfieldToken != "" {
		return
	}

	// overwrite with environment/arguments vars
	overwBrownfieldToken := viper.GetString(ConcertoBrownfieldToken)
	if overwBrownfieldToken != "" {
		log.Debug("Concerto Brownfield token taken from env/args")
		config.BrownfieldToken = overwBrownfieldToken
	}
	return
}

func (config *Config) readCommandPollingConfig() {
	if config.CommandPollingToken != "" || config.ServerID != "" {
		return
	}

	// overwrite with environment/arguments vars
	overwCommandPollingToken := viper.GetString(ConcertoCommandPollingToken)
	if overwCommandPollingToken != "" {
		log.Debug("Concerto Command Polling token taken from env/args")
		config.CommandPollingToken = overwCommandPollingToken
	}

	// overwrite with environment/arguments vars
	overwServerID := viper.GetString(ConcertoServerId)
	if overwServerID != "" {
		log.Debug("Concerto server ID taken from env/args")
		config.ServerID = overwServerID
	}
	return
}

// evaluateCertificate determines if a certificate has been issued for a host
func (config *Config) evaluateCertificate() error {
	if utils.FileExists(config.Certificate.Cert) {
		data, err := ioutil.ReadFile(config.Certificate.Cert)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(data)
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return err
		}

		if len(cert.Subject.OrganizationalUnit) > 0 && cert.Subject.OrganizationalUnit[0] == "Hosts" {
			config.IsHost = true
			return nil
		}
		if len(cert.Issuer.Organization) > 0 && cert.Issuer.Organization[0] == "Tapp" {
			config.IsHost = true
			return nil
		}
	}
	config.IsHost = false
	return nil
}

// evaluateAPIEndpointURL evaluates if API endpoint url is valid,
//advising if invalid version defined, and adapting if required
func (config *Config) evaluateAPIEndpointURL() error {
	logger.DebugFuncInfo()

	// remove ending slash if exist
	config.APIEndpoint = strings.TrimRight(config.APIEndpoint, "/")

	// In User mode, endpoint url should include API version
	if !config.IsAgentMode() {
		cURL, err := url.Parse(config.APIEndpoint)
		if err != nil {
			return err
		}
		fmt.Printf("cURL.Path: %v\n", cURL.Path)
		if cURL.Path == "" {
			config.APIEndpoint = strings.Join([]string{config.APIEndpoint, VERSION_API_USER_MODE}, "/")
			log.Warnf(
				"Defined API server endpoint url does not include API version. Normalized to latest version (%s): %s",
				VERSION_API_USER_MODE, config.APIEndpoint,
			)
		}
		if cURL.Path != "" && cURL.Path != strings.Join([]string{"/", VERSION_API_USER_MODE}, "") {
			log.Warnf(
				"Defined API server endpoint url does not match the latest supported API version (%s). Found %s",
				VERSION_API_USER_MODE, cURL.Path,
			)
		}
	}
	return nil
}

// GetDefaultLogFilePath returns default configuration path file
func GetDefaultLogFilePath() string {
	if runtime.GOOS == "windows" {
		return WindowsServerLogFilePath
	}
	return NixServerLogFilePath
}

// GetDefaultCaCertFilePath returns default configuration path file
func GetDefaultCaCertFilePath() string {
	if runtime.GOOS == "windows" {
		return WindowsServerCaCertPath
	}
	return NixServerCaCertPath
}

// GetDefaultCertFilePath returns default configuration path file
func GetDefaultCertFilePath() string {
	if runtime.GOOS == "windows" {
		return WindowsServerCertPath
	}
	return NixServerCertPath
}

// GetDefaultKeyFilePath returns default configuration path file
func GetDefaultKeyFilePath() string {
	if runtime.GOOS == "windows" {
		return WindowsServerKeyPath
	}
	return NixServerKeyPath
}
