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

func TestListStorageVolumes(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Volume{},
			server:   testutils.NewServer(http.StatusOK, []*types.Volume{}),
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

			volumes, err := svc.ListStorageVolumes(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.StoragePlan{},
			server:   testutils.NewServer(http.StatusOK, []*types.StoragePlan{}),
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

			storagePlans, err := svc.ListServerStoragePlans(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.LoadBalancerPlan{},
			server:   testutils.NewServer(http.StatusOK, []*types.LoadBalancerPlan{}),
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

			loadBalancerPlans, err := svc.ListLoadBalancerPlans(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ClusterPlan{},
			server:   testutils.NewServer(http.StatusOK, []*types.ClusterPlan{}),
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

			clusterPlans, err := svc.ListClusterPlans(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.GenericImage{},
			server:   testutils.NewServer(http.StatusOK, []*types.GenericImage{}),
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ServerArray{},
			server:   testutils.NewServer(http.StatusOK, []*types.ServerArray{}),
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   testutils.NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.GetServerArray(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   testutils.NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.CreateServerArray(context.Background(), new(map[string]any))
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   testutils.NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.UpdateServerArray(context.Background(), testutils.TEST, new(map[string]any))
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   testutils.NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.BootServerArray(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   testutils.NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.ShutdownServerArray(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   testutils.NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.EmptyServerArray(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerArray),
			server:   testutils.NewServer(http.StatusOK, new(types.ServerArray)),
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

			serverArray, err := svc.EnlargeServerArray(context.Background(), testutils.TEST, new(map[string]any))
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Server{},
			server:   testutils.NewServer(http.StatusOK, []*types.Server{}),
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

			servers, err := svc.ListServerArrayServers(context.Background(), testutils.TEST)
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

			err := svc.DeleteServerArray(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListServerPlans(t *testing.T) {
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

			serverPlans, err := svc.ListServerPlans(context.Background(), testutils.TEST, testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ServerPlan),
			server:   testutils.NewServer(http.StatusOK, new(types.ServerPlan)),
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

			serverPlan, err := svc.GetServerPlan(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Server{},
			server:   testutils.NewServer(http.StatusOK, []*types.Server{}),
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

			server, err := svc.GetServer(context.Background(), testutils.TEST)
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

			server, err := svc.CreateServer(context.Background(), new(map[string]any))
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

			server, err := svc.UpdateServer(context.Background(), testutils.TEST, new(map[string]any))
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

			server, err := svc.BootServer(context.Background(), testutils.TEST, new(map[string]any))
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

			server, err := svc.RebootServer(context.Background(), testutils.TEST, new(map[string]any))
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

			server, err := svc.ShutdownServer(context.Background(), testutils.TEST, new(map[string]any))
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

			server, err := svc.OverrideServer(context.Background(), testutils.TEST, new(map[string]any))
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

			err := svc.DeleteServer(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListServerFloatingIPs(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.FloatingIP{},
			server:   testutils.NewServer(http.StatusOK, []*types.FloatingIP{}),
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

			floatingIPs, err := svc.ListServerFloatingIPs(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Volume{},
			server:   testutils.NewServer(http.StatusOK, []*types.Volume{}),
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

			volumes, err := svc.ListServerVolumes(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Event{},
			server:   testutils.NewServer(http.StatusOK, []*types.Event{}),
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

			events, err := svc.ListServerEvents(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ScriptChar{},
			server:   testutils.NewServer(http.StatusOK, []*types.ScriptChar{}),
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

			scripts, err := svc.ListOperationalScripts(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Event),
			server:   testutils.NewServer(http.StatusOK, new(types.Event)),
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

			event, err := svc.ExecuteOperationalScript(context.Background(), testutils.TEST, testutils.TEST, new(map[string]any))
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.SSHProfile{},
			server:   testutils.NewServer(http.StatusOK, []*types.SSHProfile{}),
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.SSHProfile),
			server:   testutils.NewServer(http.StatusOK, new(types.SSHProfile)),
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

			sshProfile, err := svc.GetSSHProfile(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.SSHProfile),
			server:   testutils.NewServer(http.StatusOK, new(types.SSHProfile)),
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

			sshProfile, err := svc.CreateSSHProfile(context.Background(), new(map[string]any))
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.SSHProfile),
			server:   testutils.NewServer(http.StatusOK, new(types.SSHProfile)),
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

			sshProfile, err := svc.UpdateSSHProfile(context.Background(), testutils.TEST, new(map[string]any))
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

			err := svc.DeleteSSHProfile(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListFloatingIPs(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.FloatingIP{},
			server:   testutils.NewServer(http.StatusOK, []*types.FloatingIP{}),
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

			floatingIPs, err := svc.ListFloatingIPs(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Realm{},
			server:   testutils.NewServer(http.StatusOK, []*types.Realm{}),
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

			realms, err := svc.ListRealms(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Realm),
			server:   testutils.NewServer(http.StatusOK, new(types.Realm)),
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

			realm, err := svc.GetRealm(context.Background(), testutils.TEST)
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
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.NodePoolPlan{},
			server:   testutils.NewServer(http.StatusOK, []*types.NodePoolPlan{}),
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

			nodePoolPlans, err := svc.ListRealmNodePoolPlans(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(nodePoolPlans, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", nodePoolPlans, test.expected)
			}
		})
	}
}
