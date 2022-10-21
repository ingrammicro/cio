package blueprint

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

func TestTemplateList(t *testing.T) {
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
				testutils.AddHandleFunc(mux, api.PathBlueprintTemplates, http.StatusOK, []*types.Template{{LabelableFields: types.LabelableFields{}}})
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
			err := TemplateList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestTemplateShow(t *testing.T) {
	tests := map[string]struct {
		expected                  any
		serverRequired            bool
		labelResolution           bool
		cookbookVersionResolution bool
	}{
		"if defined endpoint for API service is resolving properly": {
			serverRequired:            true,
			labelResolution:           true,
			cookbookVersionResolution: true,
		},
		"if defined endpoint for API service is resolving properly but not for cookbook versions": {
			expected:       "HTTP request failed",
			serverRequired: true,
		},
		"if defined endpoint for API service is resolving properly but not for labelling": {
			expected:                  "HTTP request failed",
			serverRequired:            true,
			cookbookVersionResolution: true,
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
				testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintTemplate, testutils.TEST), http.StatusOK, new(types.Template))

				if test.cookbookVersionResolution {
					testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusOK, []*types.CookbookVersion{})
				} else {
					testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusConflict, nil)
				}

				if test.labelResolution {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusOK, []*types.Label{})
				} else {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusConflict, nil)
				}
				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := TemplateShow()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

//setTemplateParams

func TestTemplateCreate(t *testing.T) {
	tests := map[string]struct {
		expected                  any
		serverRequired            bool
		labelResolution           bool
		creatable                 bool
		cookbookVersionResolution bool
		incompatibleParameters    bool
	}{
		"if invalid parameters defined": {
			expected:               "invalid parameters detected",
			incompatibleParameters: true,
		},
		"if defined endpoint for API service is resolving properly": {
			serverRequired:            true,
			labelResolution:           true,
			creatable:                 true,
			cookbookVersionResolution: true,
		},
		"if defined endpoint for API service is invalid or cannot be reached for labels mapping": {
			expected: "Cannot execute request",
		},
		"if defined endpoint for API service is resolving properly but not for labelling resolution": {
			expected:       "HTTP request failed",
			serverRequired: true,
		},
		"if defined endpoint for API service is resolving properly but not for template creation": {
			expected:                  "HTTP request failed",
			serverRequired:            true,
			labelResolution:           true,
			cookbookVersionResolution: true,
		},
		"if defined endpoint for API service is resolving properly but not for cookbook versions": {
			expected:        "HTTP request failed",
			serverRequired:  true,
			labelResolution: true,
			creatable:       true,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(cmd.Name, testutils.TEST)
			viper.Set(cmd.GenericImageId, testutils.TEST)
			viper.Set(cmd.Labels, testutils.TEST)

			config.APIEndpoint = testutils.TEST
			if test.incompatibleParameters {
				viper.Set(cmd.ConfigurationAttributes, testutils.TEST)
				viper.Set(cmd.ConfigurationAttributesFromFile, testutils.TEST)
			} else {
				viper.Set(cmd.ConfigurationAttributes, nil)
				viper.Set(cmd.ConfigurationAttributesFromFile, nil)
			}
			if test.serverRequired {
				mux := http.NewServeMux()

				if test.cookbookVersionResolution {
					testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusOK, []*types.CookbookVersion{})
				} else {
					testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusConflict, nil)
				}

				variants := map[string]testutils.Value{"GET": {Status: http.StatusOK, V: []*types.Label{}}}
				if test.labelResolution {
					variants["POST"] = testutils.Value{Status: http.StatusOK, V: types.Label{ID: ""}}
				} else {
					variants["POST"] = testutils.Value{Status: http.StatusConflict, V: nil}
				}
				testutils.AddHandleFuncMultiple(mux, api.PathLabels, variants)

				if test.creatable {
					testutils.AddHandleFunc(mux, api.PathBlueprintTemplates, http.StatusOK, new(types.Template))
				} else {
					testutils.AddHandleFunc(mux, api.PathBlueprintTemplates, http.StatusConflict, nil)
				}
				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := TemplateCreate()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestTemplateUpdate(t *testing.T) {
	tests := map[string]struct {
		expected                  any
		serverRequired            bool
		labelResolution           bool
		creatable                 bool
		cookbookVersionResolution bool
		incompatibleParameters    bool
	}{
		"if invalid parameters defined": {
			expected:               "invalid parameters detected",
			incompatibleParameters: true,
		},
		"if defined endpoint for API service is resolving properly": {
			serverRequired:            true,
			creatable:                 true,
			cookbookVersionResolution: true,
			labelResolution:           true,
		},
		"if defined endpoint for API service is invalid": {
			expected: "Cannot execute request",
		},
		"if defined endpoint for API service is resolving properly but not for cookbook versions": {
			expected:        "HTTP request failed",
			serverRequired:  true,
			creatable:       true,
			labelResolution: true,
		},
		"if defined endpoint for API service is invalid or cannot be reached for labels mapping": {
			expected:                  "HTTP request failed",
			creatable:                 true,
			serverRequired:            true,
			cookbookVersionResolution: true,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(cmd.Id, testutils.TEST)
			viper.Set(cmd.Name, testutils.TEST)

			config.APIEndpoint = testutils.TEST
			if test.incompatibleParameters {
				viper.Set(cmd.ConfigurationAttributes, testutils.TEST)
				viper.Set(cmd.ConfigurationAttributesFromFile, testutils.TEST)
			} else {
				viper.Set(cmd.ConfigurationAttributes, nil)
				viper.Set(cmd.ConfigurationAttributesFromFile, nil)
			}
			if test.serverRequired {
				mux := http.NewServeMux()

				if test.creatable {
					testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintTemplate, testutils.TEST), http.StatusOK, new(types.Template))
				} else {
					testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintTemplate, testutils.TEST), http.StatusConflict, nil)
				}

				if test.cookbookVersionResolution {
					testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusOK, []*types.CookbookVersion{})
				} else {
					testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusConflict, nil)
				}

				if test.labelResolution {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusOK, []*types.Label{})
				} else {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusConflict, nil)
				}

				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := TemplateUpdate()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestTemplateCompile(t *testing.T) {
	tests := map[string]struct {
		expected                  any
		serverRequired            bool
		labelResolution           bool
		cookbookVersionResolution bool
	}{
		"if defined endpoint for API service is resolving properly": {
			serverRequired:            true,
			labelResolution:           true,
			cookbookVersionResolution: true,
		},
		"if defined endpoint for API service is resolving properly but not for cookbook versions": {
			expected:       "HTTP request failed",
			serverRequired: true,
		},
		"if defined endpoint for API service is resolving properly but not for labelling": {
			expected:                  "HTTP request failed",
			serverRequired:            true,
			cookbookVersionResolution: true,
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
				testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintTemplateCompile, testutils.TEST), http.StatusOK, new(types.Template))

				if test.cookbookVersionResolution {
					testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusOK, []*types.CookbookVersion{})
				} else {
					testutils.AddHandleFunc(mux, api.PathBlueprintCookbookVersions, http.StatusConflict, nil)
				}

				if test.labelResolution {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusOK, []*types.Label{})
				} else {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusConflict, nil)
				}
				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := TemplateCompile()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestTemplateDelete(t *testing.T) {
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

			err := TemplateDelete()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestTemplateScriptList(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.TemplateScript{},
			server:   testutils.NewServer(http.StatusOK, []*types.TemplateScript{}),
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
			err := TemplateScriptList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestTemplateScriptShow(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemplateScript),
			server:   testutils.NewServer(http.StatusOK, new(types.TemplateScript)),
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
			err := TemplateScriptShow()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

// TODO refactor
func TestTemplateScriptCreate(t *testing.T) {
	tests := map[string]struct {
		expected                       any
		serverRequired                 bool
		creatable                      bool
		incompatibleParameters         bool
		invalidParameterValues         bool
		invalidParameterValuesFromFile bool
	}{
		"if invalid parameters defined": {
			expected:               "invalid parameters detected",
			incompatibleParameters: true,
		},
		"if defined endpoint for API service is resolving properly": {
			serverRequired: true,
			creatable:      true,
		},
		"if defined endpoint for API service is invalid or cannot be reached for labels mapping": {
			expected: "Cannot execute request",
		},
		"if invalid parameter values defined": {
			expected:               "flag parameter-values isn't a valid JSON",
			invalidParameterValues: true,
		},
		"if invalid parameter values from file defined": {
			expected:                       "cannot open test to read JSON",
			invalidParameterValuesFromFile: true,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(cmd.Type, testutils.TEST)
			viper.Set(cmd.ScriptId, testutils.TEST)
			viper.Set(cmd.TemplateId, testutils.TEST)

			config.APIEndpoint = testutils.TEST
			viper.Set(cmd.ParameterValues, nil)
			viper.Set(cmd.ParameterValuesFromFile, nil)

			if test.incompatibleParameters {
				viper.Set(cmd.ParameterValues, testutils.TEST)
				viper.Set(cmd.ParameterValuesFromFile, testutils.TEST)
			} else if test.invalidParameterValues {
				viper.Set(cmd.ParameterValues, testutils.TEST)
			} else if test.invalidParameterValuesFromFile {
				viper.Set(cmd.ParameterValuesFromFile, testutils.TEST)
			} else if !test.invalidParameterValues {
				viper.Set(cmd.ParameterValues, "{\"One\":\"123\"}")
			} else if !test.invalidParameterValuesFromFile {
				viper.Set(cmd.ParameterValuesFromFile, "../../testdata/test_utils.txt")
			}

			if test.serverRequired {
				mux := http.NewServeMux()

				if test.creatable {
					testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintTemplateScripts, testutils.TEST), http.StatusOK, new(types.TemplateScript))
				} else {
					testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintTemplateScripts, testutils.TEST), http.StatusConflict, nil)
				}
				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := TemplateScriptCreate()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

//TemplateScriptUpdate

func TestTemplateScriptDelete(t *testing.T) {
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

			err := TemplateScriptDelete()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestTemplateScriptReorder(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.TemplateScript{},
			server:   testutils.NewServer(http.StatusOK, []*types.TemplateScript{}),
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
			err := TemplateScriptReorder()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestTemplateServersList(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.TemplateServer{},
			server:   testutils.NewServer(http.StatusOK, []*types.TemplateServer{}),
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
			err := TemplateServersList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

//processCookbookVersionItem
//convertFlagParamsToCookbookVersions
//resolveCookbookVersions
