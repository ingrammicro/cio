package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/types"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestListBrownfieldCloudAccounts(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.CloudAccount{},
			server:   NewServer(http.StatusOK, []*types.CloudAccount{}),
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

			cloudAccounts, err := svc.ListBrownfieldCloudAccounts(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudAccounts, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudAccounts, test.expected)
			}
		})
	}
}

func TestGetBrownfieldCloudAccount(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudAccount),
			server:   NewServer(http.StatusOK, new(types.CloudAccount)),
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

			cloudAccount, err := svc.GetBrownfieldCloudAccount(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudAccount, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudAccount, test.expected)
			}
		})
	}
}

func TestImportServers(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudAccount),
			server:   NewServer(http.StatusOK, new(types.CloudAccount)),
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

			cloudAccount, err := svc.ImportServers(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudAccount, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudAccount, test.expected)
			}
		})
	}
}

func TestImportVPCs(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudAccount),
			server:   NewServer(http.StatusOK, new(types.CloudAccount)),
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

			cloudAccount, err := svc.ImportVPCs(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudAccount, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudAccount, test.expected)
			}
		})
	}
}

func TestImportFloatingIPs(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudAccount),
			server:   NewServer(http.StatusOK, new(types.CloudAccount)),
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

			cloudAccount, err := svc.ImportFloatingIPs(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudAccount, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudAccount, test.expected)
			}
		})
	}
}

func TestImportVolumes(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudAccount),
			server:   NewServer(http.StatusOK, new(types.CloudAccount)),
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

			cloudAccount, err := svc.ImportVolumes(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudAccount, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudAccount, test.expected)
			}
		})
	}
}

func TestImportKubernetesClusters(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudAccount),
			server:   NewServer(http.StatusOK, new(types.CloudAccount)),
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

			cloudAccount, err := svc.ImportKubernetesClusters(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudAccount, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudAccount, test.expected)
			}
		})
	}
}

func TestImportPolicies(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudAccount),
			server:   NewServer(http.StatusOK, new(types.CloudAccount)),
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

			cloudAccount, err := svc.ImportPolicies(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudAccount, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudAccount, test.expected)
			}
		})
	}
}
