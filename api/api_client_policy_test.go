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

func TestListPolicyDefinitions(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.PolicyDefinition{},
			server:   testutils.NewServer(http.StatusOK, []*types.PolicyDefinition{}),
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

			definitions, err := svc.ListPolicyDefinitions(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(definitions, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", definitions, test.expected)
			}
		})
	}
}

func TestGetPolicyDefinition(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.PolicyDefinition),
			server:   testutils.NewServer(http.StatusOK, new(types.PolicyDefinition)),
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

			definition, err := svc.GetPolicyDefinition(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(definition, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", definition, test.expected)
			}
		})
	}
}

func TestCreatePolicyDefinition(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.PolicyDefinition),
			server:   testutils.NewServer(http.StatusOK, new(types.PolicyDefinition)),
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

			definition, err := svc.CreatePolicyDefinition(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(definition, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", definition, test.expected)
			}
		})
	}
}

func TestUpdatePolicyDefinition(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.PolicyDefinition),
			server:   testutils.NewServer(http.StatusOK, new(types.PolicyDefinition)),
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

			definition, err := svc.UpdatePolicyDefinition(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(definition, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", definition, test.expected)
			}
		})
	}
}

func TestDeletePolicyDefinition(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: nil,
			server:   testutils.NewServer(http.StatusOK, nil),
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

			err := svc.DeletePolicyDefinition(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListPolicyDefinitionAssignments(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.PolicyAssignment{},
			server:   testutils.NewServer(http.StatusOK, []*types.PolicyAssignment{}),
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

			assignments, err := svc.ListPolicyDefinitionAssignments(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(assignments, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", assignments, test.expected)
			}
		})
	}
}

func TestGetPolicyAssignment(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.PolicyAssignment),
			server:   testutils.NewServer(http.StatusOK, new(types.PolicyAssignment)),
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

			assignment, err := svc.GetPolicyAssignment(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(assignment, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", assignment, test.expected)
			}
		})
	}
}

func TestCreatePolicyAssignment(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.PolicyAssignment),
			server:   testutils.NewServer(http.StatusOK, new(types.PolicyAssignment)),
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

			assignment, err := svc.CreatePolicyAssignment(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(assignment, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", assignment, test.expected)
			}
		})
	}
}

func TestUpdatePolicyAssignment(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.PolicyAssignment),
			server:   testutils.NewServer(http.StatusOK, new(types.PolicyAssignment)),
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

			assignment, err := svc.UpdatePolicyAssignment(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(assignment, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", assignment, test.expected)
			}
		})
	}
}

func TestDeletePolicyAssignment(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.PolicyAssignment),
			server:   testutils.NewServer(http.StatusOK, new(types.PolicyAssignment)),
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

			assignment, err := svc.DeletePolicyAssignment(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(assignment, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", assignment, test.expected)
			}
		})
	}
}
