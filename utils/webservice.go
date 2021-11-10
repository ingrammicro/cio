// Copyright (c) 2017-2021 Ingram Micro Inc.

package utils

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const WebServiceConfigurationFailed = "web service configuration failed. No data in configuration"
const ContentTypeApplicationJson = "application/json"
const ConfigurationIsIncomplete = "configuration is incomplete"
const HttpTimeOut = 30

// ConcertoService defines actions to be performed by web service manager
type ConcertoService interface {
	Post(path string, payload *map[string]interface{}) ([]byte, int, error)
	Put(path string, payload *map[string]interface{}) ([]byte, int, error)
	Delete(path string) ([]byte, int, error)
	Get(path string) ([]byte, int, error)
	GetFile(url string, filePath string, discoveryFileName bool) (string, int, error)
	PutFile(sourceFilePath string, targetURL string) ([]byte, int, error)
}

// HTTPConcertoservice web service manager.
type HTTPConcertoservice struct {
	config *Config
	client *http.Client
}

// NewHTTPConcertoService creates new http Concerto client based on config
func NewHTTPConcertoService(config *Config) (hcs *HTTPConcertoservice, err error) {

	if config == nil {
		return nil, fmt.Errorf(WebServiceConfigurationFailed)
	}

	if !config.IsConfigReady() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	// creates HTTP Concerto service with config
	hcs = &HTTPConcertoservice{
		config: config,
	}

	// Loads CA Certificate
	caCert, err := ioutil.ReadFile(hcs.config.Certificate.Ca)
	if err != nil {
		return nil, fmt.Errorf("cannot read IMCO CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Loads Clients Certificates and creates and 509KeyPair
	cert, err := tls.LoadX509KeyPair(hcs.config.Certificate.Cert, hcs.config.Certificate.Key)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot read IMCO API key (from '%s' and '%s'): %v",
			hcs.config.Certificate.Cert,
			hcs.config.Certificate.Key,
			err,
		)
	}

	// Creates a client with specific transport configurations
	hcs.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
		Timeout: time.Second * time.Duration(HttpTimeOut),
	}
	return hcs, nil
}

// NewHTTPConcertoServiceWithBrownfieldToken creates new http Concerto client based on config
func NewHTTPConcertoServiceWithBrownfieldToken(config *Config) (hcs *HTTPConcertoservice, err error) {

	if config == nil {
		return nil, fmt.Errorf(WebServiceConfigurationFailed)
	}

	if !config.IsConfigReadyBrownfield() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	// creates HTTP Concerto service with config
	hcs = &HTTPConcertoservice{
		config: config,
	}
	// Creates a client with no certificates and insecure option
	hcs.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * time.Duration(HttpTimeOut),
	}
	return hcs, nil
}

// NewHTTPConcertoServiceWithCommandPolling creates new http Concerto client based on config
func NewHTTPConcertoServiceWithCommandPolling(config *Config) (hcs *HTTPConcertoservice, err error) {

	if config == nil {
		return nil, fmt.Errorf(WebServiceConfigurationFailed)
	}

	if !config.IsConfigReadyCommandPolling() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	// creates HTTP Concerto service with config
	hcs = &HTTPConcertoservice{
		config: config,
	}
	// Creates a client with no certificates and insecure option
	hcs.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * time.Duration(HttpTimeOut),
	}
	return hcs, nil
}

// Post sends POST request to Concerto API
func (hcs *HTTPConcertoservice) Post(path string, payload *map[string]interface{}) ([]byte, int, error) {

	url, jsPayload, err := hcs.prepareCall(path, payload)
	if err != nil {
		return nil, 0, err
	}

	log.Debugf("Sending POST request to %s with payload %v ", url, jsPayload)
	req, err := http.NewRequest("POST", url, jsPayload)
	req.Header.Add("Content-Type", ContentTypeApplicationJson)
	if hcs.config.BrownfieldToken != "" {
		log.Debugf(
			"Including brownfield token %s in POST request as X-Concerto-Brownfield-Token header ",
			hcs.config.BrownfieldToken,
		)
		req.Header.Add("X-Concerto-Brownfield-Token", hcs.config.BrownfieldToken)
	}
	if hcs.config.CommandPollingToken != "" && hcs.config.ServerID != "" {
		log.Debugf(
			"Including command polling token %s in POST request as X-IMCO-Command-Polling-Token header ",
			hcs.config.CommandPollingToken,
		)
		req.Header.Add("X-IMCO-Command-Polling-Token", hcs.config.CommandPollingToken)
		log.Debugf("Including Server id %s in POST request as X-IMCO-Server-ID header ", hcs.config.ServerID)
		req.Header.Add("X-IMCO-Server-ID", hcs.config.ServerID)
	}
	response, err := hcs.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	return hcs.receiveResponse(response)
}

// Put sends PUT request to Concerto API
func (hcs *HTTPConcertoservice) Put(path string, payload *map[string]interface{}) ([]byte, int, error) {
	url, jsPayload, err := hcs.prepareCall(path, payload)
	if err != nil {
		return nil, 0, err
	}

	log.Debugf("Sending PUT request to %s with payload %v ", url, jsPayload)
	request, err := http.NewRequest("PUT", url, jsPayload)
	if err != nil {
		return nil, 0, err
	}
	request.Header = map[string][]string{"Content-type": {ContentTypeApplicationJson}}
	response, err := hcs.client.Do(request)
	if err != nil {
		return nil, 0, err
	}

	return hcs.receiveResponse(response)
}

// Delete sends DELETE request to Concerto API
func (hcs *HTTPConcertoservice) Delete(path string) ([]byte, int, error) {
	url, _, err := hcs.prepareCall(path, nil)
	if err != nil {
		return nil, 0, err
	}

	log.Debugf("Sending DELETE request to %s", url)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, 0, err
	}
	request.Header = map[string][]string{"Content-type": {ContentTypeApplicationJson}}
	response, err := hcs.client.Do(request)
	if err != nil {
		return nil, 0, err
	}

	return hcs.receiveResponse(response)
}

// Get sends GET request to Concerto API
func (hcs *HTTPConcertoservice) Get(path string) ([]byte, int, error) {

	url, _, err := hcs.prepareCall(path, nil)
	if err != nil {
		return nil, 0, err
	}

	log.Debugf("Sending GET request to %s", url)
	response, err := hcs.client.Get(url)
	if err != nil {
		return nil, 0, err
	}

	return hcs.receiveResponse(response)
}

// GetFile sends GET request to Concerto API and receives a file
func (hcs *HTTPConcertoservice) GetFile(url string, filePath string, discoveryFileName bool) (string, int, error) {

	log.Debugf("Sending GET request to %s", url)
	response, err := hcs.client.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()
	log.Debugf("Status code:%d message:%s", response.StatusCode, response.Status)
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return "", response.StatusCode, fmt.Errorf("HTTP request failed with status %s", response.Status)
	}

	realFileName := filePath
	if discoveryFileName {
		r, err := regexp.Compile("filename=\\\"([^\\\"]*){1}\\\"")
		if err != nil {
			return "", 0, err
		}
		realFileName = fmt.Sprintf(
			"%s/%s",
			filePath,
			r.FindStringSubmatch(response.Header.Get("Content-Disposition"))[1],
		)
	}

	output, err := os.Create(realFileName)
	if err != nil {
		return "", response.StatusCode, err
	}
	defer output.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return "", response.StatusCode, err
	}

	log.Debugf("%#v bytes downloaded", n)
	return realFileName, response.StatusCode, nil
}

// PutFile sends PUT request to send a file
func (hcs *HTTPConcertoservice) PutFile(sourceFilePath string, targetURL string) ([]byte, int, error) {

	data, err := os.Open(sourceFilePath)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PUT", targetURL, data)
	if err != nil {
		return nil, 0, err
	}

	res, err := hcs.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	status := res.StatusCode
	return buf, status, nil
}

func (hcs *HTTPConcertoservice) prepareCall(
	path string,
	payload *map[string]interface{},
) (url string, jsPayload *strings.Reader, err error) {

	if hcs.config == nil || hcs.client == nil {
		return "", nil, fmt.Errorf("Can not call web service without loading configuration")
	}

	url = fmt.Sprintf("%s%s", hcs.config.APIEndpoint, path)

	if payload == nil {
		return url, nil, nil
	}

	// payload to json
	json, err := json.Marshal(payload)
	if err != nil {
		return "", nil, err
	}

	return url, strings.NewReader(string(json)), err
}

func (hcs *HTTPConcertoservice) receiveResponse(response *http.Response) (body []byte, status int, err error) {

	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	log.Debugf("Response : %s", body)
	log.Debugf("Status code: (%d) %s", response.StatusCode, response.Status)

	return body, response.StatusCode, nil
}
