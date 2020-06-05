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

func TestGetCloudApplicationDeployment(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudApplicationDeployment),
			server:   NewServer(http.StatusOK, new(types.CloudApplicationDeployment)),
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

			deployment, status, err := svc.GetCloudApplicationDeployment(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(deployment, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					deployment, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestDeleteCloudApplicationDeployment(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudApplicationDeployment),
			server:   NewServer(http.StatusOK, new(types.CloudApplicationDeployment)),
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

			deployment, err := svc.DeleteCloudApplicationDeployment(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(deployment, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", deployment, test.expected)
			}
		})
	}
}

func TestCreateCloudApplicationDeploymentTask(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudApplicationDeploymentTask),
			server:   NewServer(http.StatusOK, new(types.CloudApplicationDeploymentTask)),
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

			deploymentTask, err := svc.CreateCloudApplicationDeploymentTask(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(deploymentTask, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", deploymentTask, test.expected)
			}
		})
	}
}

func TestGetCloudApplicationDeploymentTask(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudApplicationDeploymentTask),
			server:   NewServer(http.StatusOK, new(types.CloudApplicationDeploymentTask)),
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

			deploymentTask, err := svc.GetCloudApplicationDeploymentTask(context.Background(), TEST, TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(deploymentTask, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", deploymentTask, test.expected)
			}
		})
	}
}

func TestListCloudApplicationTemplates(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.CloudApplicationTemplate{},
			server:   NewServer(http.StatusOK, []*types.CloudApplicationTemplate{}),
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

			templates, err := svc.ListCloudApplicationTemplates(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(templates, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", templates, test.expected)
			}
		})
	}
}

func TestGetCloudApplicationTemplate(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudApplicationTemplate),
			server:   NewServer(http.StatusOK, new(types.CloudApplicationTemplate)),
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

			template, err := svc.GetCloudApplicationTemplate(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(template, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", template, test.expected)
			}
		})
	}
}

func TestCreateCloudApplicationTemplate(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudApplicationTemplate),
			server:   NewServer(http.StatusOK, new(types.CloudApplicationTemplate)),
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

			template, err := svc.CreateCloudApplicationTemplate(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(template, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", template, test.expected)
			}
		})
	}
}

func TestParseMetadataCloudApplicationTemplate(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CloudApplicationTemplate),
			server:   NewServer(http.StatusOK, new(types.CloudApplicationTemplate)),
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

			template, err := svc.ParseMetadataCloudApplicationTemplate(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(template, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", template, test.expected)
			}
		})
	}
}

func TestDeleteCloudApplicationTemplate(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: nil,
			server:   NewServer(http.StatusOK, nil),
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

			err := svc.DeleteCloudApplicationTemplate(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestCreateTemporaryArchive(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemporaryArchive),
			server:   NewServer(http.StatusOK, new(types.TemporaryArchive)),
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

			temporaryArchive, err := svc.CreateTemporaryArchive(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(temporaryArchive, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", temporaryArchive, test.expected)
			}
		})
	}
}

func TestCreateTemporaryArchiveImport(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemporaryArchiveImport),
			server:   NewServer(http.StatusOK, new(types.TemporaryArchiveImport)),
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

			temporaryArchiveImport, err := svc.CreateTemporaryArchiveImport(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(temporaryArchiveImport, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", temporaryArchiveImport, test.expected)
			}
		})
	}
}

func TestGetTemporaryArchiveImport(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemporaryArchiveImport),
			server:   NewServer(http.StatusOK, new(types.TemporaryArchiveImport)),
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

			temporaryArchiveImport, err := svc.GetTemporaryArchiveImport(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(temporaryArchiveImport, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", temporaryArchiveImport, test.expected)
			}
		})
	}
}

func TestCreateTemporaryArchiveExport(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemporaryArchiveExport),
			server:   NewServer(http.StatusOK, new(types.TemporaryArchiveExport)),
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

			temporaryArchiveExport, err := svc.CreateTemporaryArchiveExport(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(temporaryArchiveExport, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", temporaryArchiveExport, test.expected)
			}
		})
	}
}

func TestGetTemporaryArchiveExportTask(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemporaryArchiveExportTask),
			server:   NewServer(http.StatusOK, new(types.TemporaryArchiveExportTask)),
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

			temporaryArchiveExportTask, err := svc.GetTemporaryArchiveExportTask(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(temporaryArchiveExportTask, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", temporaryArchiveExportTask, test.expected)
			}
		})
	}
}
