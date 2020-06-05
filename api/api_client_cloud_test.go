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

func TestListStorageVolumes(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Volume{},
			server:   NewServer(http.StatusOK, []*types.Volume{}),
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

			volumes, err := svc.ListStorageVolumes(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(volumes, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", volumes, test.expected)
			}
		})
	}
}

func TestListCloudProviders(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.CloudProvider{},
			server:   NewServer(http.StatusOK, []*types.CloudProvider{}),
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

			cloudProviders, err := svc.ListCloudProviders(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cloudProviders, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cloudProviders, test.expected)
			}
		})
	}
}

func TestListServerStoragePlans(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.StoragePlan{},
			server:   NewServer(http.StatusOK, []*types.StoragePlan{}),
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

			storagePlans, err := svc.ListServerStoragePlans(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(storagePlans, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", storagePlans, test.expected)
			}
		})
	}
}

func TestListLoadBalancerPlans(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.LoadBalancerPlan{},
			server:   NewServer(http.StatusOK, []*types.LoadBalancerPlan{}),
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

			loadBalancerPlans, err := svc.ListLoadBalancerPlans(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(loadBalancerPlans, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", loadBalancerPlans, test.expected)
			}
		})
	}
}

func TestListClusterPlans(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ClusterPlan{},
			server:   NewServer(http.StatusOK, []*types.ClusterPlan{}),
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

			clusterPlans, err := svc.ListClusterPlans(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(clusterPlans, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", clusterPlans, test.expected)
			}
		})
	}
}

func TestListGenericImages(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.GenericImage{},
			server:   NewServer(http.StatusOK, []*types.GenericImage{}),
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

			genericImages, err := svc.ListGenericImages(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(genericImages, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", genericImages, test.expected)
			}
		})
	}
}

func TestListServerArrays(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ServerArray{},
			server:   NewServer(http.StatusOK, []*types.ServerArray{}),
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

			serverArrays, err := svc.ListServerArrays(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArrays, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArrays, test.expected)
			}
		})
	}
}

func TestGetServerArray(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.GetServerArray(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArray, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArray, test.expected)
			}
		})
	}
}

func TestCreateServerArray(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.CreateServerArray(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArray, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArray, test.expected)
			}
		})
	}
}

func TestUpdateServerArray(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.UpdateServerArray(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArray, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArray, test.expected)
			}
		})
	}
}

func TestBootServerArray(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.BootServerArray(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArray, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArray, test.expected)
			}
		})
	}
}

func TestShutdownServerArray(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.ShutdownServerArray(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArray, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArray, test.expected)
			}
		})
	}
}

func TestEmptyServerArray(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.EmptyServerArray(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArray, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArray, test.expected)
			}
		})
	}
}

func TestEnlargeServerArray(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.EnlargeServerArray(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArray, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArray, test.expected)
			}
		})
	}
}

func TestListServerArrayServers(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Server{},
			server:   NewServer(http.StatusOK, []*types.Server{}),
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

			servers, err := svc.ListServerArrayServers(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(servers, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", servers, test.expected)
			}
		})
	}
}

func TestDeleteServerArray(t *testing.T) {
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

			err := svc.DeleteServerArray(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListServerPlans(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ServerPlan{},
			server:   NewServer(http.StatusOK, []*types.ServerPlan{}),
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

			serverPlans, err := svc.ListServerPlans(context.Background(), TEST, TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverPlans, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverPlans, test.expected)
			}
		})
	}
}

func TestGetServerPlan(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerPlan),
			server:   NewServer(http.StatusOK, new(types.ServerPlan)),
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

			serverPlan, err := svc.GetServerPlan(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverPlan, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverPlan, test.expected)
			}
		})
	}
}

func TestListServers(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Server{},
			server:   NewServer(http.StatusOK, []*types.Server{}),
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

			servers, err := svc.ListServers(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(servers, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", servers, test.expected)
			}
		})
	}
}

func TestGetServer(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Server),
			server:   NewServer(http.StatusOK, new(types.Server)),
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

			server, err := svc.GetServer(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestCreateServer(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Server),
			server:   NewServer(http.StatusOK, new(types.Server)),
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

			server, err := svc.CreateServer(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestUpdateServer(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Server),
			server:   NewServer(http.StatusOK, new(types.Server)),
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

			server, err := svc.UpdateServer(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestBootServer(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Server),
			server:   NewServer(http.StatusOK, new(types.Server)),
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

			server, err := svc.BootServer(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestRebootServer(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Server),
			server:   NewServer(http.StatusOK, new(types.Server)),
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

			server, err := svc.RebootServer(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestShutdownServer(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Server),
			server:   NewServer(http.StatusOK, new(types.Server)),
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

			server, err := svc.ShutdownServer(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestOverrideServer(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Server),
			server:   NewServer(http.StatusOK, new(types.Server)),
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

			server, err := svc.OverrideServer(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestDeleteServer(t *testing.T) {
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

			err := svc.DeleteServer(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListServerFloatingIPs(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.FloatingIP{},
			server:   NewServer(http.StatusOK, []*types.FloatingIP{}),
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

			floatingIPs, err := svc.ListServerFloatingIPs(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(floatingIPs, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", floatingIPs, test.expected)
			}
		})
	}
}

func TestListServerVolumes(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Volume{},
			server:   NewServer(http.StatusOK, []*types.Volume{}),
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

			volumes, err := svc.ListServerVolumes(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(volumes, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", volumes, test.expected)
			}
		})
	}
}

func TestListServerEvents(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Event{},
			server:   NewServer(http.StatusOK, []*types.Event{}),
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

			events, err := svc.ListServerEvents(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(events, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", events, test.expected)
			}
		})
	}
}

func TestListOperationalScripts(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ScriptChar{},
			server:   NewServer(http.StatusOK, []*types.ScriptChar{}),
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

			scripts, err := svc.ListOperationalScripts(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(scripts, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", scripts, test.expected)
			}
		})
	}
}

func TestExecuteOperationalScript(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Event),
			server:   NewServer(http.StatusOK, new(types.Event)),
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

			event, err := svc.ExecuteOperationalScript(context.Background(), TEST, TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(event, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", event, test.expected)
			}
		})
	}
}

func TestListSSHProfiles(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.SSHProfile{},
			server:   NewServer(http.StatusOK, []*types.SSHProfile{}),
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

			sshProfiles, err := svc.ListSSHProfiles(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(sshProfiles, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", sshProfiles, test.expected)
			}
		})
	}
}

func TestGetSSHProfile(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.SSHProfile),
			server:   NewServer(http.StatusOK, new(types.SSHProfile)),
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

			sshProfile, err := svc.GetSSHProfile(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(sshProfile, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", sshProfile, test.expected)
			}
		})
	}
}

func TestCreateSSHProfile(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.SSHProfile),
			server:   NewServer(http.StatusOK, new(types.SSHProfile)),
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

			sshProfile, err := svc.CreateSSHProfile(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(sshProfile, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", sshProfile, test.expected)
			}
		})
	}
}

func TestUpdateSSHProfile(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.SSHProfile),
			server:   NewServer(http.StatusOK, new(types.SSHProfile)),
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

			sshProfile, err := svc.UpdateSSHProfile(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(sshProfile, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", sshProfile, test.expected)
			}
		})
	}
}

func TestDeleteSSHProfile(t *testing.T) {
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

			err := svc.DeleteSSHProfile(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListFloatingIPs(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.FloatingIP{},
			server:   NewServer(http.StatusOK, []*types.FloatingIP{}),
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

			floatingIPs, err := svc.ListFloatingIPs(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(floatingIPs, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", floatingIPs, test.expected)
			}
		})
	}
}

func TestListRealms(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Realm{},
			server:   NewServer(http.StatusOK, []*types.Realm{}),
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

			realms, err := svc.ListRealms(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(realms, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", realms, test.expected)
			}
		})
	}
}

func TestGetRealm(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Realm),
			server:   NewServer(http.StatusOK, new(types.Realm)),
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

			realm, err := svc.GetRealm(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(realm, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", realm, test.expected)
			}
		})
	}
}

func TestListRealmNodePoolPlans(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.NodePoolPlan{},
			server:   NewServer(http.StatusOK, []*types.NodePoolPlan{}),
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

			nodePoolPlans, err := svc.ListRealmNodePoolPlans(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(nodePoolPlans, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", nodePoolPlans, test.expected)
			}
		})
	}
}
