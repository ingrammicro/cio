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

func TestListCertificates(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Certificate{},
			server:   testutils.NewServer(http.StatusOK, []*types.Certificate{}),
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

			certificates, err := svc.ListCertificates(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(certificates, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", certificates, test.expected)
			}
		})
	}
}

func TestGetCertificate(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Certificate),
			server:   testutils.NewServer(http.StatusOK, new(types.Certificate)),
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

			certificate, err := svc.GetCertificate(context.Background(), testutils.TEST, testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(certificate, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", certificate, test.expected)
			}
		})
	}
}

func TestCreateCertificate(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Certificate),
			server:   testutils.NewServer(http.StatusOK, new(types.Certificate)),
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

			certificate, err := svc.CreateCertificate(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(certificate, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", certificate, test.expected)
			}
		})
	}
}

func TestUpdateCertificate(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Certificate),
			server:   testutils.NewServer(http.StatusOK, new(types.Certificate)),
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

			certificate, err := svc.UpdateCertificate(context.Background(), testutils.TEST, testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(certificate, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", certificate, test.expected)
			}
		})
	}
}

func TestDeleteCertificate(t *testing.T) {
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

			err := svc.DeleteCertificate(context.Background(), testutils.TEST, testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListFirewallProfiles(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.FirewallProfile{},
			server:   testutils.NewServer(http.StatusOK, []*types.FirewallProfile{}),
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

			firewallProfiles, err := svc.ListFirewallProfiles(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(firewallProfiles, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", firewallProfiles, test.expected)
			}
		})
	}
}

func TestGetFirewallProfile(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.FirewallProfile),
			server:   testutils.NewServer(http.StatusOK, new(types.FirewallProfile)),
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

			firewallProfile, err := svc.GetFirewallProfile(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(firewallProfile, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", firewallProfile, test.expected)
			}
		})
	}
}

func TestCreateFirewallProfile(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.FirewallProfile),
			server:   testutils.NewServer(http.StatusOK, new(types.FirewallProfile)),
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

			firewallProfile, err := svc.CreateFirewallProfile(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(firewallProfile, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", firewallProfile, test.expected)
			}
		})
	}
}

func TestUpdateFirewallProfile(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.FirewallProfile),
			server:   testutils.NewServer(http.StatusOK, new(types.FirewallProfile)),
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

			firewallProfile, err := svc.UpdateFirewallProfile(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(firewallProfile, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", firewallProfile, test.expected)
			}
		})
	}
}

func TestDeleteFirewallProfile(t *testing.T) {
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

			err := svc.DeleteFirewallProfile(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestGetFloatingIP(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.FloatingIP),
			server:   testutils.NewServer(http.StatusOK, new(types.FloatingIP)),
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

			floatingIP, err := svc.GetFloatingIP(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(floatingIP, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", floatingIP, test.expected)
			}
		})
	}
}

func TestCreateFloatingIP(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.FloatingIP),
			server:   testutils.NewServer(http.StatusOK, new(types.FloatingIP)),
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

			floatingIP, err := svc.CreateFloatingIP(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(floatingIP, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", floatingIP, test.expected)
			}
		})
	}
}

func TestUpdateFloatingIP(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.FloatingIP),
			server:   testutils.NewServer(http.StatusOK, new(types.FloatingIP)),
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

			floatingIP, err := svc.UpdateFloatingIP(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(floatingIP, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", floatingIP, test.expected)
			}
		})
	}
}

func TestAttachFloatingIP(t *testing.T) {
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

			server, err := svc.AttachFloatingIP(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(server, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", server, test.expected)
			}
		})
	}
}

func TestDetachFloatingIP(t *testing.T) {
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

			err := svc.DetachFloatingIP(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestDeleteFloatingIP(t *testing.T) {
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

			err := svc.DeleteFloatingIP(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestDiscardFloatingIP(t *testing.T) {
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

			err := svc.DiscardFloatingIP(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListListeners(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Listener{},
			server:   testutils.NewServer(http.StatusOK, []*types.Listener{}),
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

			listeners, err := svc.ListListeners(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listeners, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listeners, test.expected)
			}
		})
	}
}

func TestGetListener(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Listener),
			server:   testutils.NewServer(http.StatusOK, new(types.Listener)),
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

			listener, err := svc.GetListener(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listener, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listener, test.expected)
			}
		})
	}
}

func TestCreateListener(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Listener),
			server:   testutils.NewServer(http.StatusOK, new(types.Listener)),
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

			listener, err := svc.CreateListener(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listener, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listener, test.expected)
			}
		})
	}
}

func TestUpdateListener(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Listener),
			server:   testutils.NewServer(http.StatusOK, new(types.Listener)),
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

			listener, err := svc.UpdateListener(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listener, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listener, test.expected)
			}
		})
	}
}

func TestDeleteListener(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Listener),
			server:   testutils.NewServer(http.StatusOK, new(types.Listener)),
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

			listener, err := svc.DeleteListener(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listener, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listener, test.expected)
			}
		})
	}
}

func TestRetryListener(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Listener),
			server:   testutils.NewServer(http.StatusOK, new(types.Listener)),
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

			listener, err := svc.RetryListener(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listener, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listener, test.expected)
			}
		})
	}
}

func TestListRules(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.ListenerRule{},
			server:   testutils.NewServer(http.StatusOK, []*types.ListenerRule{}),
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

			listenerRules, err := svc.ListRules(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listenerRules, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listenerRules, test.expected)
			}
		})
	}
}

func TestCreateRule(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ListenerRule),
			server:   testutils.NewServer(http.StatusOK, new(types.ListenerRule)),
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

			listenerRule, err := svc.CreateRule(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listenerRule, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listenerRule, test.expected)
			}
		})
	}
}

func TestUpdateRule(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.ListenerRule),
			server:   testutils.NewServer(http.StatusOK, new(types.ListenerRule)),
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

			listenerRule, err := svc.UpdateRule(context.Background(), testutils.TEST, testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(listenerRule, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", listenerRule, test.expected)
			}
		})
	}
}

func TestDeleteRule(t *testing.T) {
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

			err := svc.DeleteRule(context.Background(), testutils.TEST, testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListLoadBalancers(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.LoadBalancer{},
			server:   testutils.NewServer(http.StatusOK, []*types.LoadBalancer{}),
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

			loadBalancers, err := svc.ListLoadBalancers(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(loadBalancers, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", loadBalancers, test.expected)
			}
		})
	}
}

func TestGetLoadBalancer(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.LoadBalancer),
			server:   testutils.NewServer(http.StatusOK, new(types.LoadBalancer)),
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

			loadBalancer, err := svc.GetLoadBalancer(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(loadBalancer, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", loadBalancer, test.expected)
			}
		})
	}
}

func TestCreateLoadBalancer(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.LoadBalancer),
			server:   testutils.NewServer(http.StatusOK, new(types.LoadBalancer)),
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

			loadBalancer, err := svc.CreateLoadBalancer(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(loadBalancer, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", loadBalancer, test.expected)
			}
		})
	}
}

func TestUpdateLoadBalancer(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.LoadBalancer),
			server:   testutils.NewServer(http.StatusOK, new(types.LoadBalancer)),
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

			loadBalancer, err := svc.UpdateLoadBalancer(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(loadBalancer, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", loadBalancer, test.expected)
			}
		})
	}
}

func TestDeleteLoadBalancer(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.LoadBalancer),
			server:   testutils.NewServer(http.StatusOK, new(types.LoadBalancer)),
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

			loadBalancer, err := svc.DeleteLoadBalancer(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(loadBalancer, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", loadBalancer, test.expected)
			}
		})
	}
}

func TestRetryLoadBalancer(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.LoadBalancer),
			server:   testutils.NewServer(http.StatusOK, new(types.LoadBalancer)),
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

			loadBalancer, err := svc.RetryLoadBalancer(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(loadBalancer, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", loadBalancer, test.expected)
			}
		})
	}
}

func TestGetLoadBalancerPlan(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.LoadBalancerPlan),
			server:   testutils.NewServer(http.StatusOK, new(types.LoadBalancerPlan)),
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

			loadBalancerPlan, err := svc.GetLoadBalancerPlan(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(loadBalancerPlan, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", loadBalancerPlan, test.expected)
			}
		})
	}
}

func TestListSubnets(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Subnet{},
			server:   testutils.NewServer(http.StatusOK, []*types.Subnet{}),
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

			subnets, err := svc.ListSubnets(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(subnets, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", subnets, test.expected)
			}
		})
	}
}

func TestGetSubnet(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Subnet),
			server:   testutils.NewServer(http.StatusOK, new(types.Subnet)),
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

			subnet, err := svc.GetSubnet(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(subnet, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", subnet, test.expected)
			}
		})
	}
}

func TestCreateSubnet(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Subnet),
			server:   testutils.NewServer(http.StatusOK, new(types.Subnet)),
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

			subnet, err := svc.CreateSubnet(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(subnet, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", subnet, test.expected)
			}
		})
	}
}

func TestUpdateSubnet(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Subnet),
			server:   testutils.NewServer(http.StatusOK, new(types.Subnet)),
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

			subnet, err := svc.UpdateSubnet(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(subnet, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", subnet, test.expected)
			}
		})
	}
}

func TestDeleteSubnet(t *testing.T) {
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

			err := svc.DeleteSubnet(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListSubnetServers(t *testing.T) {
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

			servers, err := svc.ListSubnetServers(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(servers, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", servers, test.expected)
			}
		})
	}
}

func TestListSubnetServerArrays(t *testing.T) {
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

			serverArrays, err := svc.ListSubnetServerArrays(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(serverArrays, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", serverArrays, test.expected)
			}
		})
	}
}

func TestListTargetGroups(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.TargetGroup{},
			server:   testutils.NewServer(http.StatusOK, []*types.TargetGroup{}),
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

			targetGroups, err := svc.ListTargetGroups(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(targetGroups, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", targetGroups, test.expected)
			}
		})
	}
}

func TestGetTargetGroup(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TargetGroup),
			server:   testutils.NewServer(http.StatusOK, new(types.TargetGroup)),
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

			targetGroup, err := svc.GetTargetGroup(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(targetGroup, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", targetGroup, test.expected)
			}
		})
	}
}

func TestCreateTargetGroup(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TargetGroup),
			server:   testutils.NewServer(http.StatusOK, new(types.TargetGroup)),
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

			targetGroup, err := svc.CreateTargetGroup(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(targetGroup, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", targetGroup, test.expected)
			}
		})
	}
}

func TestUpdateTargetGroup(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TargetGroup),
			server:   testutils.NewServer(http.StatusOK, new(types.TargetGroup)),
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

			targetGroup, err := svc.UpdateTargetGroup(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(targetGroup, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", targetGroup, test.expected)
			}
		})
	}
}

func TestDeleteTargetGroup(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TargetGroup),
			server:   testutils.NewServer(http.StatusOK, new(types.TargetGroup)),
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

			targetGroup, err := svc.DeleteTargetGroup(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(targetGroup, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", targetGroup, test.expected)
			}
		})
	}
}

func TestRetryTargetGroup(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TargetGroup),
			server:   testutils.NewServer(http.StatusOK, new(types.TargetGroup)),
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

			targetGroup, err := svc.RetryTargetGroup(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(targetGroup, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", targetGroup, test.expected)
			}
		})
	}
}

func TestListTargets(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Target{},
			server:   testutils.NewServer(http.StatusOK, []*types.Target{}),
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

			targets, err := svc.ListTargets(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(targets, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", targets, test.expected)
			}
		})
	}
}

func TestCreateTarget(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Target),
			server:   testutils.NewServer(http.StatusOK, new(types.Target)),
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

			target, err := svc.CreateTarget(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(target, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", target, test.expected)
			}
		})
	}
}

func TestDeleteTarget(t *testing.T) {
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

			err := svc.DeleteTarget(context.Background(), testutils.TEST, testutils.TEST, testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListVPCs(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Vpc{},
			server:   testutils.NewServer(http.StatusOK, []*types.Vpc{}),
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

			vpcs, err := svc.ListVPCs(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(vpcs, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", vpcs, test.expected)
			}
		})
	}
}

func TestGetVPC(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Vpc),
			server:   testutils.NewServer(http.StatusOK, new(types.Vpc)),
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

			vpc, err := svc.GetVPC(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(vpc, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", vpc, test.expected)
			}
		})
	}
}

func TestCreateVPC(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Vpc),
			server:   testutils.NewServer(http.StatusOK, new(types.Vpc)),
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

			vpc, err := svc.CreateVPC(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(vpc, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", vpc, test.expected)
			}
		})
	}
}

func TestUpdateVPC(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Vpc),
			server:   testutils.NewServer(http.StatusOK, new(types.Vpc)),
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

			vpc, err := svc.UpdateVPC(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(vpc, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", vpc, test.expected)
			}
		})
	}
}

func TestDeleteVPC(t *testing.T) {
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

			err := svc.DeleteVPC(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestDiscardVPC(t *testing.T) {
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

			err := svc.DiscardVPC(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestGetVPN(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Vpn),
			server:   testutils.NewServer(http.StatusOK, new(types.Vpn)),
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

			vpn, err := svc.GetVPN(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(vpn, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", vpn, test.expected)
			}
		})
	}
}

func TestCreateVPN(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Vpn),
			server:   testutils.NewServer(http.StatusOK, new(types.Vpn)),
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

			vpn, err := svc.CreateVPN(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(vpn, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", vpn, test.expected)
			}
		})
	}
}

func TestDeleteVPN(t *testing.T) {
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

			err := svc.DeleteVPN(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListVPNPlans(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.VpnPlan{},
			server:   testutils.NewServer(http.StatusOK, []*types.VpnPlan{}),
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

			vpnPlans, err := svc.ListVPNPlans(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(vpnPlans, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", vpnPlans, test.expected)
			}
		})
	}
}

func TestListDomains(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Domain{},
			server:   testutils.NewServer(http.StatusOK, []*types.Domain{}),
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

			domains, err := svc.ListDomains(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(domains, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", domains, test.expected)
			}
		})
	}
}

func TestGetDomain(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Domain),
			server:   testutils.NewServer(http.StatusOK, new(types.Domain)),
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

			domain, err := svc.GetDomain(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(domain, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", domain, test.expected)
			}
		})
	}
}

func TestCreateDomain(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Domain),
			server:   testutils.NewServer(http.StatusOK, new(types.Domain)),
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

			domain, err := svc.CreateDomain(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(domain, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", domain, test.expected)
			}
		})
	}
}

func TestDeleteDomain(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Domain),
			server:   testutils.NewServer(http.StatusOK, new(types.Domain)),
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

			domain, err := svc.DeleteDomain(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(domain, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", domain, test.expected)
			}
		})
	}
}

func TestRetryDomain(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Domain),
			server:   testutils.NewServer(http.StatusOK, new(types.Domain)),
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

			domain, err := svc.RetryDomain(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(domain, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", domain, test.expected)
			}
		})
	}
}

func TestListRecords(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Record{},
			server:   testutils.NewServer(http.StatusOK, []*types.Record{}),
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

			records, err := svc.ListRecords(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(records, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", records, test.expected)
			}
		})
	}
}

func TestGetRecord(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Record),
			server:   testutils.NewServer(http.StatusOK, new(types.Record)),
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

			record, err := svc.GetRecord(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(record, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", record, test.expected)
			}
		})
	}
}

func TestCreateRecord(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Record),
			server:   testutils.NewServer(http.StatusOK, new(types.Record)),
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

			record, err := svc.CreateRecord(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(record, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", record, test.expected)
			}
		})
	}
}

func TestUpdateRecord(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Record),
			server:   testutils.NewServer(http.StatusOK, new(types.Record)),
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

			record, err := svc.UpdateRecord(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(record, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", record, test.expected)
			}
		})
	}
}

func TestDeleteRecord(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Record),
			server:   testutils.NewServer(http.StatusOK, new(types.Record)),
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

			record, err := svc.DeleteRecord(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(record, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", record, test.expected)
			}
		})
	}
}

func TestRetryRecord(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Record),
			server:   testutils.NewServer(http.StatusOK, new(types.Record)),
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

			record, err := svc.RetryRecord(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(record, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", record, test.expected)
			}
		})
	}
}
