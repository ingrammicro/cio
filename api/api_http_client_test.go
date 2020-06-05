package api

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/ingrammicro/cio/configuration"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

//TODO COMMON
const TEST = "test"

//TODO COMMON
func InitConfig() *configuration.Config {
	config := new(configuration.Config)
	config.XMLName = xml.Name{
		Space: "",
		Local: "",
	}
	config.APIEndpoint = "https://clients.test.imco.io/v3"
	config.LogFile = "/var/log/concerto-client.log"
	config.LogLevel = "info"
	config.Certificate = configuration.Cert{
		Cert: "testdata/ssl/cert.crt",
		Key:  "testdata/ssl/private/cert.key",
		Ca:   "testdata/ssl/ca_cert.pem",
	}
	config.BootstrapConfig = configuration.BootstrapConfig{
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

func InitNewHTTPClient() (*HTTPClient, *configuration.Config, error) {
	config := new(configuration.Config)
	config = InitConfig()
	svc, err := NewHTTPClient(config)
	if err != nil {
		return nil, nil, err
	}
	return svc, config, nil
}

func InitConfigAndClientAPI() (*configuration.Config, *ClientAPI, error) {
	config := InitConfig()
	//config := new(configuration.Config) // para forzar error de config
	ds, err := NewHTTPClient(config)
	if err != nil {
		return nil, nil, err
	}

	svc := new(ClientAPI)
	svc.HTTPClient = *ds
	return config, svc, nil
}

func NewServer(returnStatus int, returnData interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//if r.URL.Path != "/storage/volumes" {
		//	t.Errorf("Expected to request '/storage/volumes', got: %s", r.URL.Path)
		//}
		//if r.Header.Get("Accept") != "application/json" {
		//	t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		//}
		w.WriteHeader(returnStatus)
		d, _ := json.Marshal(returnData)
		w.Write(d)
	}))
	return server
}

type TestData struct {
	ID          string    `json:"id"          header:"ID"`
	Timestamp   time.Time `json:"timestamp"   header:"TIMESTAMP"`
	Name        string    `json:"name"       header:"NAME"`
	Description string    `json:"description" header:"DESCRIPTION"`
}

type APIError struct {
	Error  string                 `json:"error"       header:"ERROR"`
	Errors map[string]interface{} `json:"errors"       header:"ERRORS"`
}

func TestNewHTTPClient(t *testing.T) {
	tests := map[string]struct {
		config   *configuration.Config
		expected interface{}
	}{
		"if config not initialized": {
			config:   nil,
			expected: WebServiceConfigurationFailed,
		},
		"if config not ready": {
			config:   InitConfig(),
			expected: ConfigurationIsIncomplete,
		},
		"if cannot read config API CA cert": {
			config:   InitConfig(),
			expected: "cannot read IMCO CA cert",
		},
		"if cannot read config API key": {
			config:   InitConfig(),
			expected: "cannot read IMCO API key",
		},
		"if config initialized": {
			config:   InitConfig(),
			expected: new(HTTPClient),
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.config != nil {
				if test.expected == ConfigurationIsIncomplete {
					test.config.APIEndpoint = ""
					test.config.Certificate.Cert = ""
					test.config.Certificate.Key = ""
					test.config.Certificate.Ca = ""
				}
				if test.expected == "cannot read IMCO CA cert" {
					test.config.Certificate.Ca = TEST
				}
				if test.expected == "cannot read IMCO API key" {
					test.config.Certificate.Key = TEST
				}
			}
			svc, err := NewHTTPClient(test.config)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			//if err == nil && fmt.Sprintf("%T", svc) != "*api.HTTPClient" {
			if err == nil && reflect.TypeOf(svc) != reflect.TypeOf(test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", svc, test.expected)
			}
		})
	}
}

func TestGetAndCheck(t *testing.T) {
	tests := map[string]struct {
		expected int
		server   *httptest.Server
	}{
		"if get method resolves with no issues": {
			expected: http.StatusOK,
			server:   NewServer(http.StatusOK, TestData{ID: TEST}),
		},
		"if get method ends with error": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, APIError{Error: "Invalid request"}),
		},
		"if get method ends with errors": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, APIError{Errors: map[string]interface{}{"name": []string{"is already taken"}}}),
		},
		"if get method ends with invalid url": {
			expected: 0,
			server:   nil,
		},
	}

	svc, config, err := InitNewHTTPClient()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = ""
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			status, err := svc.GetAndCheck(context.Background(), pathStorageVolumes, true, &TestData{ID: TEST})
			if err != nil && status != test.expected {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if status != test.expected {
				t.Errorf("Expected to return %v, got %v", test.expected, status)
			}
		})
	}
}

func TestPutAndCheck(t *testing.T) {
	tests := map[string]struct {
		expected int
		server   *httptest.Server
	}{
		"if put method resolves with no issues": {
			expected: http.StatusOK,
			server:   NewServer(http.StatusOK, TestData{ID: TEST}),
		},
		"if put method ends with error": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, APIError{Error: "Invalid request"}),
		},
		"if put method ends with errors": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, APIError{Errors: map[string]interface{}{"name": []string{"is already taken"}}}),
		},
		"if put method ends with invalid url": {
			expected: 0,
			server:   nil,
		},
	}

	svc, config, err := InitNewHTTPClient()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = ""
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			payload := map[string]interface{}{
				"name": TEST,
			}
			status, err := svc.PutAndCheck(context.Background(), pathStorageVolumes, &payload, true, &TestData{ID: TEST})
			if err != nil && status != test.expected {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if status != test.expected {
				t.Errorf("Expected to return %v, got %v", test.expected, status)
			}
		})
	}
}

func TestPostAndCheck(t *testing.T) {
	tests := map[string]struct {
		expected int
		server   *httptest.Server
	}{
		"if post method resolves with no issues": {
			expected: http.StatusOK,
			server:   NewServer(http.StatusOK, TestData{ID: TEST}),
		},
		"if post method ends with error": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, APIError{Error: "Invalid request"}),
		},
		"if post method ends with errors": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, APIError{Errors: map[string]interface{}{"name": []string{"is already taken"}}}),
		},
		"if post method ends with invalid url": {
			expected: 0,
			server:   nil,
		},
	}

	svc, config, err := InitNewHTTPClient()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = ""
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			payload := map[string]interface{}{
				"name": TEST,
			}
			status, err := svc.PostAndCheck(context.Background(), pathStorageVolumes, &payload, true, &TestData{ID: TEST})
			if err != nil && status != test.expected {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if status != test.expected {
				t.Errorf("Expected to return %v, got %v", test.expected, status)
			}
		})
	}
}

func TestDeleteAndCheck(t *testing.T) {
	tests := map[string]struct {
		expected int
		server   *httptest.Server
	}{
		"if delete method resolves with no issues": {
			expected: http.StatusOK,
			server:   NewServer(http.StatusOK, TestData{ID: TEST}),
		},
		"if delete method ends with error": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, APIError{Error: "Invalid request"}),
		},
		"if delete method ends with errors": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, APIError{Errors: map[string]interface{}{"name": []string{"is already taken"}}}),
		},
		"if delete method ends with invalid url": {
			expected: 0,
			server:   nil,
		},
	}

	svc, config, err := InitNewHTTPClient()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = ""
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			status, err := svc.DeleteAndCheck(context.Background(), pathStorageVolumes, true, &TestData{ID: TEST})
			if err != nil && status != test.expected {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if status != test.expected {
				t.Errorf("Expected to return %v, got %v", test.expected, status)
			}
		})
	}
}

func TestDownloadFile(t *testing.T) {
	tests := map[string]struct {
		expected int
		server   *httptest.Server
		URL      string
		filePath string
	}{
		"if cannot create download request": {
			expected: 0,
			server:   nil,
			URL:      "hg ttp:",
			filePath: "testdata/test_utils.txt",
		},
		"if cannot download file": {
			expected: 0,
			server:   nil,
			filePath: "testdata/test_utils.txt",
		},
		"if HTTP request failed with status": {
			expected: http.StatusBadRequest,
			server:   NewServer(http.StatusBadRequest, TestData{ID: TEST}),
			filePath: "testdata/test_utils.txt",
		},
		"if cannot create file": {
			expected: http.StatusOK,
			server:   NewServer(http.StatusOK, TestData{ID: TEST}),
			filePath: "testdata//",
		},
		"if no discovery": {
			expected: http.StatusOK,
			server:   NewServer(http.StatusOK, TestData{ID: TEST}),
			filePath: "testdata/test_utils.txt",
		},
	}

	svc, _, err := InitNewHTTPClient()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			URL := ""
			if test.server != nil {
				server := test.server
				defer server.Close()
				URL = server.URL + "/testdata/client.xml"
			}
			if test.URL != "" {
				URL = test.URL
			}

			realFileName, status, err := svc.DownloadFile(context.Background(), URL, test.filePath, false)
			t.Logf("DownloadFile: %v %v %v\n", realFileName, status, err)
			if err != nil && status != test.expected {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if status != test.expected {
				t.Errorf("Expected to return %v, got %v", test.expected, status)
			}
		})
	}
}

func TestUploadFile(t *testing.T) {
	tests := map[string]struct {
		targetURL      string
		sourceFilePath string
		server         *httptest.Server
	}{
		"if completed successfully": {
			targetURL:      TEST,
			sourceFilePath: "testdata/test_utils.txt",
			server:         NewServer(http.StatusOK, TestData{ID: TEST}),
		},
		"if cannot open source file path": {
			targetURL:      "",
			sourceFilePath: "",
			server:         nil,
		},
		"if cannot upload file": {
			targetURL:      "",
			sourceFilePath: "testdata/test_utils.txt",
			server:         NewServer(http.StatusBadRequest, TestData{ID: TEST}),
		},
	}

	svc, _, err := InitNewHTTPClient()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.server != nil {
				server := test.server
				defer server.Close()
				if test.targetURL != "" {
					test.targetURL = server.URL + "/testdata/client.xml"
				}
			}

			err := svc.UploadFile(context.Background(), test.sourceFilePath, test.targetURL)
			if err != nil && test.targetURL != "" {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && test.targetURL == "" {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestCheckAndUnmarshal(t *testing.T) {
	var v interface{}
	v = TestData{ID: TEST}
	data, err := json.Marshal(v)

	tests := map[string]struct {
		status   int
		data     []byte
		check    bool
		v        any
		expected interface{}
	}{
		"if status and data are correct": {
			status:   http.StatusOK,
			data:     data,
			check:    true,
			v:        TestData{ID: TEST},
			expected: http.StatusOK,
		},
		"if HTTP request failed": {
			status:   http.StatusBadRequest,
			data:     data,
			check:    true,
			v:        TestData{ID: TEST},
			expected: "HTTP request failed",
		},
		"if invalid JSON data": {
			status:   http.StatusOK,
			data:     nil,
			check:    false,
			v:        TestData{ID: TEST},
			expected: "unexpected end of JSON input",
		},
	}

	svc, _, err := InitNewHTTPClient()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			status, err := svc.checkAndUnmarshal(test.status, test.data, test.check, test.v)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && status != test.expected {
				t.Errorf("Unexpected response: %v. Expected: %v\n", status, test.expected)
			}
		})
	}
}
