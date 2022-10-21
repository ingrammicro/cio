package cli

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/internal/testutils"
	"github.com/ingrammicro/cio/types"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// TODO FatalError/osExit(1)
func TestWireUpAPIClient(t *testing.T) {
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
				c.Certificate.Cert = "../../configuration/testdata/ssl/cert.crt"
				c.Certificate.Key = "../../configuration/testdata/ssl/private/cert.key"
				c.Certificate.Ca = "../../configuration/testdata/ssl/ca_cert.pem"

				test.config = c
			}
			configuration.SetConfig(test.config)
			svc, config, f := WireUpAPIClient()
			if svc != nil && f != nil && config != test.config {
				t.Errorf("Unexpected response: %v. Expected: %v\n", config, test.config)
			}
		})
	}
}

func TestLoadCloudProvidersMapping(t *testing.T) {
	cpID := testutils.TEST
	cpName := "Sample Cloud Provider"

	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: map[string]string{cpID: cpName},
			server:   testutils.NewServer(http.StatusOK, []*types.CloudProvider{{ID: cpID, Name: cpName}}),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	c, err := configuration.InitializeConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	svc, config, f := WireUpAPIClient()
	if svc != nil && f != nil && config != c {
		t.Errorf("Unexpected response: %v. Expected: %v\n", config, c)
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = "TEST"
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			cloudProvidersMap, err := LoadCloudProvidersMapping(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudProvidersMap, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudProvidersMap, test.expected)
			}
		})
	}
}

func TestLoadLocationsMapping(t *testing.T) {
	cpID := testutils.TEST
	cpName := "Sample Location"

	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: map[string]string{cpID: cpName},
			server:   testutils.NewServer(http.StatusOK, []*types.Location{{ID: cpID, Name: cpName}}),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	c, err := configuration.InitializeConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	svc, config, f := WireUpAPIClient()
	if svc != nil && f != nil && config != c {
		t.Errorf("Unexpected response: %v. Expected: %v\n", config, c)
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = "TEST"
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			locationsMap, err := LoadLocationsMapping(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(locationsMap, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", locationsMap, test.expected)
			}
		})
	}
}
