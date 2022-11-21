package blueprint

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/internal/testutils"
	"github.com/ingrammicro/cio/types"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCookbookVersionList(t *testing.T) {
	tests := map[string]struct {
		expected        any
		serverRequired  bool
		labelResolution bool
	}{
		"if defined endpoint for API service is resolving properly": {
			serverRequired:  true,
			labelResolution: true,
		},
		"if defined endpoint for API service is resolving properly, but provided label format is invalid": {
			expected:        "invalid label format",
			serverRequired:  true,
			labelResolution: true,
		},
		"if defined endpoint for API service is resolving properly but not for labelling": {
			expected:       "HTTP request failed",
			serverRequired: true,
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
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
				testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusOK, []*types.CookbookVersion{{LabelableFields: types.LabelableFields{}}})
				if test.labelResolution {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusOK, []*types.Label{})
					if test.expected == "invalid label format" {
						viper.Set(cmd.Labels, "|")
					}
				} else {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusConflict, nil)
				}
				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := CookbookVersionList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestCookbookVersionShow(t *testing.T) {
	tests := map[string]struct {
		expected        any
		serverRequired  bool
		labelResolution bool
	}{
		"if defined endpoint for API service is resolving properly": {
			serverRequired:  true,
			labelResolution: true,
		},
		"if defined endpoint for API service is resolving properly but not for labelling": {
			expected:       "HTTP request failed",
			serverRequired: true,
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(cmd.Id, testutils.TEST)
			config.APIEndpoint = testutils.TEST
			if test.serverRequired {
				mux := http.NewServeMux()
				testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintCookbookVersion, testutils.TEST), http.StatusOK, new(types.CookbookVersion))
				if test.labelResolution {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusOK, []*types.Label{})
				} else {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusConflict, nil)
				}
				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := CookbookVersionShow()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

//// TODO
//func TestCookbookVersionUpload(t *testing.T) {
//	tests := map[string]struct {
//		expected any
//		server   *httptest.Server
//	}{
//		"if defined endpoint for API service is resolving properly": {
//			expected: new(types.CookbookVersion),
//			server:   testutils.NewServer(http.StatusOK, new(types.CookbookVersion)),
//		},
//		//"if defined endpoint for API service is invalid or cannot be reached": {
//		//	expected: "Cannot execute request",
//		//	server:   nil,
//		//},
//	}
//
//	config := testutils.InitConfig()
//	configuration.SetConfig(config)
//	cmd.RootCmd.SetContext(context.Background())
//
//	for title, test := range tests {
//		t.Run(title, func(t *testing.T) {
//			viper.Set(cmd.Id, "123")
//			viper.Set(cmd.Filepath, "../../testdata/test_utils.txt")
//			config.APIEndpoint = testutils.TEST
//			if test.server != nil {
//				mux := http.NewServeMux()
//				mux.HandleFunc("/labels", func(res http.ResponseWriter, req *http.Request) {
//					res.WriteHeader(http.StatusOK)
//					d, _ := json.Marshal([]*types.Label{})
//					res.Write(d)
//				})
//				mux.HandleFunc("/blueprint/cookbook_versions", func(res http.ResponseWriter, req *http.Request) {
//					res.WriteHeader(http.StatusOK)
//					d, _ := json.Marshal(new(types.CookbookVersion))
//					res.Write(d)
//				})
//				server := httptest.NewServer(mux)
//				defer server.Close()
//				config.APIEndpoint = server.URL
//			}
//			//UploadURL
//			err := CookbookVersionUpload()
//			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
//				t.Errorf("Unexpected error: %v\n", err)
//			}
//		})
//	}
//}

func TestCleanCookbookVersion(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			server: testutils.NewServer(http.StatusOK, nil),
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	ds, err := api.NewHTTPClient(config)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
		return
	}
	svc := new(api.ClientAPI)
	svc.HTTPClient = *ds

	formatter := format.GetFormatter()

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			config.APIEndpoint = testutils.TEST
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			err := cleanCookbookVersion(context.Background(), svc, formatter, testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestCookbookVersionDelete(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			server: testutils.NewServer(http.StatusOK, nil),
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
			viper.Set(cmd.Id, testutils.TEST)
			config.APIEndpoint = testutils.TEST
			if test.server != nil {
				server := test.server
				defer server.Close()
				config.APIEndpoint = server.URL
			}

			err := CookbookVersionDelete()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}
