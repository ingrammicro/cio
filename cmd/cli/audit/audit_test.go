package audit

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/internal/testutils"
	"github.com/ingrammicro/cio/types"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEventList(t *testing.T) {
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

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = testutils.TEST
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := EventList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestSystemEventList(t *testing.T) {
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

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = testutils.TEST
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := SystemEventList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}
