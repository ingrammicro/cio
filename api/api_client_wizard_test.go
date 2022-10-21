package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/internal/testutils"
	"github.com/ingrammicro/cio/types"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestListApps(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.WizardApp{},
			server:   testutils.NewServer(http.StatusOK, []*types.WizardApp{}),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := InitConfigAndClientAPI()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			apps, err := svc.ListApps(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(apps, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", apps, test.expected)
			}
		})
	}
}

func TestDeployApp(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Server),
			server:   testutils.NewServer(http.StatusOK, new(types.Server)),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := InitConfigAndClientAPI()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			server, err := svc.DeployApp(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestListLocations(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Location{},
			server:   testutils.NewServer(http.StatusOK, []*types.Location{}),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := InitConfigAndClientAPI()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			locations, err := svc.ListLocations(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(locations, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", locations, test.expected)
			}
		})
	}
}

func TestListWizardCloudProviders(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.CloudProvider{},
			server:   testutils.NewServer(http.StatusOK, []*types.CloudProvider{}),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := InitConfigAndClientAPI()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			cloudProviders, err := svc.ListWizardCloudProviders(context.Background(), testutils.TEST, testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudProviders, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudProviders, test.expected)
			}
		})
	}
}

func TestListWizardServerPlans(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ServerPlan{},
			server:   testutils.NewServer(http.StatusOK, []*types.ServerPlan{}),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := InitConfigAndClientAPI()
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			serverPlans, err := svc.ListWizardServerPlans(context.Background(), testutils.TEST, testutils.TEST, testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverPlans, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverPlans, test.expected)
			}
		})
	}
}
