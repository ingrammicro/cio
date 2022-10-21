package brownfield

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/internal/testutils"
	"github.com/ingrammicro/cio/types"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBrownfieldCloudAccountList(t *testing.T) {
	tests := map[string]struct {
		expected                 any
		serverRequired           bool
		cloudProvidersResolution bool
	}{
		"if defined endpoint for API service is resolving properly": {
			serverRequired:           true,
			cloudProvidersResolution: true,
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
		},
		"if defined endpoint for API service is resolving properly but not for cloud providers": {
			expected:       "HTTP request failed",
			serverRequired: true,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = testutils.TEST
			viper.Set(cmd.Labels, "")
			if test.serverRequired {
				mux := http.NewServeMux()
				testutils.AddHandleFunc(mux, api.PathBrownfieldCloudAccounts, http.StatusOK, []*types.CloudAccount{{
					CloudProviderName: testutils.TEST,
				}})
				if test.cloudProvidersResolution {
					testutils.AddHandleFunc(mux, api.PathCloudCloudProviders, http.StatusOK, []*types.CloudProvider{})
				} else {
					testutils.AddHandleFunc(mux, api.PathCloudCloudProviders, http.StatusConflict, nil)
				}

				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := BrownfieldCloudAccountList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}
