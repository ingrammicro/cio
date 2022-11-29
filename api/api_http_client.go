// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/configuration"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ingrammicro/cio/logger"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context/ctxhttp"
)

const CannotReadHttpResponseBody = "Cannot read http response body"
const WebServiceConfigurationFailed = "web service configuration failed. No data in configuration"
const ConfigurationIsIncomplete = "configuration is incomplete"
const HttpTimeOut = 30

// HTTPClient web service manager
type HTTPClient struct {
	config *configuration.Config
	client *http.Client
}

// NewHTTPClient creates new http cli based on config
func NewHTTPClient(config *configuration.Config) (svc *HTTPClient, err error) {
	if config == nil {
		return nil, fmt.Errorf(WebServiceConfigurationFailed)
	}

	if !config.IsConfigReady() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	// creates HTTP service with config
	svc = &HTTPClient{
		config: config,
	}

	// Loads CA Certificate
	caCert, err := os.ReadFile(svc.config.Certificate.Ca)
	if err != nil {
		return nil, fmt.Errorf("cannot read "+configuration.CloudOrchestratorPlatformName+" CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Loads Clients Certificates and creates and 509KeyPair
	cert, err := tls.LoadX509KeyPair(svc.config.Certificate.Cert, svc.config.Certificate.Key)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot read "+configuration.CloudOrchestratorPlatformName+" API key (from '%s' and '%s'): %v",
			svc.config.Certificate.Cert,
			svc.config.Certificate.Key,
			err,
		)
	}

	// Creates a client with specific transport configurations
	svc.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
		Timeout: time.Second * time.Duration(HttpTimeOut),
	}
	return svc, nil
}

// GetAndCheck executes Get method, then optionally evaluates response status and response unmarshalling
func (imco *HTTPClient) GetAndCheck(ctx context.Context, path string, bCheck bool, v interface{}) (int, error) {
	logger.DebugFuncInfo()

	data, status, err := imco.request(ctx, "GET", path, nil)
	if err != nil {
		return status, err
	}
	return imco.checkAndUnmarshal(status, data, bCheck, v)
}

// PutAndCheck executes Put method, then optionally evaluates response status and response unmarshalling
func (imco *HTTPClient) PutAndCheck(ctx context.Context, path string, payload *map[string]interface{},
	bCheck bool, v interface{},
) (int, error) {
	logger.DebugFuncInfo()

	data, status, err := imco.request(ctx, "PUT", path, payload)
	if err != nil {
		return status, err
	}
	return imco.checkAndUnmarshal(status, data, bCheck, v)
}

// PostAndCheck executes Post method, then optionally evaluates response status and response unmarshalling
func (imco *HTTPClient) PostAndCheck(ctx context.Context, path string, payload *map[string]interface{},
	bCheck bool, v interface{},
) (int, error) {
	logger.DebugFuncInfo()

	data, status, err := imco.request(ctx, "POST", path, payload)
	if err != nil {
		return status, err
	}
	return imco.checkAndUnmarshal(status, data, bCheck, v)
}

// DeleteAndCheck executes Delete method, then optionally evaluates response status and response unmarshalling
func (imco *HTTPClient) DeleteAndCheck(ctx context.Context, path string, bCheck bool, v interface{}) (int, error) {
	logger.DebugFuncInfo()

	data, status, err := imco.request(ctx, "DELETE", path, nil)
	if err != nil {
		return status, err
	}
	return imco.checkAndUnmarshal(status, data, bCheck, v)
}

// DownloadFile gets a file from given url saving file into given file path
func (imco *HTTPClient) DownloadFile(ctx context.Context, url string, filepath string, discoveryFileName bool,
) (string, int, error) {
	logger.DebugFuncInfo()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, errors.Wrap(err, "Cannot create download request")
	}
	log.Debugf("Sending GET request to %s", url)

	req = req.WithContext(ctx)
	httpResponse, err := ctxhttp.Do(ctx, imco.client, req)
	if err != nil {
		return "", 0, errors.Wrap(err, "Cannot download file")
	}
	defer httpResponse.Body.Close()
	log.Debugf("Status code:%d message:%s", httpResponse.StatusCode, httpResponse.Status)
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		return "", httpResponse.StatusCode, fmt.Errorf("HTTP request failed with status %s", httpResponse.Status)
	}

	buf, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return "", httpResponse.StatusCode, errors.Wrap(err, CannotReadHttpResponseBody)
	}

	realFileName := filepath
	if discoveryFileName {
		r, err := regexp.Compile("filename=\\\"([^\\\"]*){1}\\\"")
		if err != nil {
			return "", httpResponse.StatusCode, err
		}
		realFileName = fmt.Sprintf("%s/%s", filepath,
			r.FindStringSubmatch(httpResponse.Header.Get("Content-Disposition"))[1],
		)
	}

	outputFile, err := os.Create(realFileName)
	if err != nil {
		return realFileName, httpResponse.StatusCode, errors.Wrap(err, "Cannot create file")
	}
	defer outputFile.Close()

	n, err := io.Copy(outputFile, bytes.NewReader(buf))
	if err != nil {
		return realFileName, httpResponse.StatusCode, errors.Wrap(err, "Cannot copy file data")
	}
	log.Debugf("%#v bytes downloaded", n)
	return realFileName, httpResponse.StatusCode, nil
}

// UploadFile uploads a file to target url
func (imco *HTTPClient) UploadFile(ctx context.Context, sourceFilePath, targetURL string) error {
	logger.DebugFuncInfo()

	sourceData, err := os.Open(sourceFilePath)
	if err != nil {
		return errors.Wrap(err, "Cannot open source file path")
	}

	req, err := http.NewRequest("PUT", targetURL, sourceData)
	if err != nil {
		return errors.Wrap(err, "Cannot create upload request")
	}
	log.Debugf("Sending PUT request to %s", targetURL)

	req = req.WithContext(ctx)
	httpResponse, err := ctxhttp.Do(ctx, imco.client, req)
	if err != nil {
		return errors.Wrap(err, "Cannot upload file")
	}
	defer httpResponse.Body.Close()
	buf, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return errors.Wrap(err, CannotReadHttpResponseBody)
	}

	if err = imco.checkStandardStatus(httpResponse.StatusCode, buf); err != nil {
		return err
	}
	return nil
}

func (imco *HTTPClient) request(ctx context.Context, method string, path string, payload *map[string]interface{},
) ([]byte, int, error) {
	if imco.config == nil || imco.client == nil {
		return nil, 0, fmt.Errorf("cannot call web service without loading configuration")
	}

	url := fmt.Sprintf("%s%s", imco.config.APIEndpoint, path)
	var body []byte
	var err error
	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return nil, 0, errors.Wrap(err, "Cannot create body")
		}
	}

	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Cannot create request")
	}
	log.Debugf("Sending %s request to %s with payload %v ", method, url, strings.NewReader(string(body)))
	req = req.WithContext(ctx)
	if body != nil {
		// Set Content-Type header. Assumption: This call is only made for JSON bodies.
		req.Header.Add("Content-Type", "application/json")
		if imco.config.BrownfieldToken != "" {
			log.Debugf(
				"Including brownfield token %s in POST request as X-Concerto-Brownfield-Token header ",
				imco.config.BrownfieldToken,
			)
			req.Header.Add("X-Concerto-Brownfield-Token", imco.config.BrownfieldToken)
		}
		if imco.config.CommandPollingToken != "" && imco.config.ServerID != "" {
			log.Debugf(
				"Including command polling token %s in POST request as X-IMCO-Command-Polling-Token header ",
				imco.config.CommandPollingToken,
			)
			req.Header.Add("X-IMCO-Command-Polling-Token", imco.config.CommandPollingToken)
			log.Debugf("Including server id %s in POST request as X-IMCO-server-ID header ", imco.config.ServerID)
			req.Header.Add("X-IMCO-server-ID", imco.config.ServerID)
		}
	}
	log.Debugf("Executing HTTP Request wit header %s", req.Header)
	httpResponse, err := ctxhttp.Do(ctx, imco.client, req)
	if err != nil {
		log.Info("Executing request failed to url: ", req.URL.String())
		return nil, 0, errors.Wrap(err, "Cannot execute request")
	}
	defer httpResponse.Body.Close()

	buf, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, httpResponse.StatusCode, errors.Wrap(err, CannotReadHttpResponseBody)
	}
	log.Debugf("Response: %s", buf)
	log.Debugf("Status code: (%d) %s", httpResponse.StatusCode, httpResponse.Status)
	return buf, httpResponse.StatusCode, nil
}

// checkStandardStatus return error if status is not OK
func (imco *HTTPClient) checkStandardStatus(status int, response []byte) error {
	if status < 300 {
		return nil
	}

	// default, raw, not parsing
	message := string(response[:])

	var responseContent map[string]interface{}
	err := json.Unmarshal(response, &responseContent)
	if err == nil {
		if responseContent["errors"] != nil {
			message = ""
			for key, value := range responseContent["errors"].(map[string]interface{}) {
				subMessages := make([]string, len(value.([]interface{})))
				for i, v := range value.([]interface{}) {
					subMessages[i] = fmt.Sprint(v)
				}
				composedMsg := strings.Join(subMessages, ",")
				message = fmt.Sprintf("%s#%s:%s", message, key, composedMsg)
			}
		}
		if responseContent["errors"] == nil && responseContent["error"] != nil {
			message = responseContent["error"].(string)
		}
	}
	return fmt.Errorf("HTTP request failed: (%d) [%s]", status, message)
}

// checkAndUnmarshal if defined, evaluates status and unmarshal json data
func (imco *HTTPClient) checkAndUnmarshal(status int, data []byte, bCheck bool, v interface{}) (int, error) {
	logger.DebugFuncInfo()

	if bCheck {
		if err := imco.checkStandardStatus(status, data); err != nil {
			return status, err
		}
	}
	if v != nil {
		if err := json.Unmarshal(data, &v); err != nil {
			return status, err
		}
	}
	return status, nil
}
