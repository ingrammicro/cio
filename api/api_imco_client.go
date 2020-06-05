// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ingrammicro/cio/configuration"
)

// IMCOClient web service manager
type IMCOClient struct {
	config *configuration.Config
	client *http.Client
}

// NewIMCOClient creates new http Concerto client based on config
func NewIMCOClient(config *configuration.Config) (ic *IMCOClient, err error) {
	if config == nil {
		return nil, fmt.Errorf(WebServiceConfigurationFailed)
	}

	if !config.IsConfigReady() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	// creates HTTP Concerto service with config
	ic = &IMCOClient{
		config: config,
	}

	// Loads CA Certificate
	caCert, err := ioutil.ReadFile(ic.config.Certificate.Ca)
	if err != nil {
		return nil, fmt.Errorf("cannot read IMCO CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Loads Clients Certificates and creates and 509KeyPair
	cert, err := tls.LoadX509KeyPair(ic.config.Certificate.Cert, ic.config.Certificate.Key)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot read IMCO API key (from '%s' and '%s'): %v",
			ic.config.Certificate.Cert,
			ic.config.Certificate.Key,
			err,
		)
	}

	// Creates a client with specific transport configurations
	ic.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	return ic, nil
}

// NewIMCOClientWithBrownfieldToken creates new http Concerto client based on config
func NewIMCOClientWithBrownfieldToken(config *configuration.Config) (ic *IMCOClient, err error) {
	if config == nil {
		return nil, fmt.Errorf(WebServiceConfigurationFailed)
	}

	if !config.IsConfigReadyBrownfield() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	// creates HTTP Concerto service with config
	ic = &IMCOClient{
		config: config,
	}
	// Creates a client with no certificates and insecure option
	ic.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return ic, nil
}

// NewIMCOClientWithCommandPolling creates new http Concerto client based on config
func NewIMCOClientWithCommandPolling(config *configuration.Config) (ic *IMCOClient, err error) {
	if config == nil {
		return nil, fmt.Errorf(WebServiceConfigurationFailed)
	}

	if !config.IsConfigReadyCommandPolling() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	// creates HTTP Concerto service with config
	ic = &IMCOClient{
		config: config,
	}
	// Creates a client with no certificates and insecure option
	ic.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return ic, nil
}
