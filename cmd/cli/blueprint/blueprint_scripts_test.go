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

func TestScriptsList(t *testing.T) {
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
				testutils.AddHandleFunc(mux, api.PathBlueprintScripts, http.StatusOK, []*types.Script{{LabelableFields: types.LabelableFields{}}})
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
			err := ScriptsList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestScriptShow(t *testing.T) {
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
				testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintScript, testutils.TEST), http.StatusOK, new(types.Script))
				if test.labelResolution {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusOK, []*types.Label{})
				} else {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusConflict, nil)
				}
				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := ScriptShow()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestScriptCreate(t *testing.T) {
	tests := map[string]struct {
		expected        any
		server          *httptest.Server
		labelResolution bool
		creatable       bool
	}{
		"if defined endpoint for API service is resolving properly": {
			expected:        new(types.Script),
			server:          testutils.NewServer(http.StatusOK, new(types.Script)),
			labelResolution: true,
			creatable:       true,
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
			server:   nil,
		},
		"if defined endpoint for API service is resolving properly but not for labelling resolution": {
			expected: "HTTP request failed",
			server:   testutils.NewServer(0, nil),
		},
		"if defined endpoint for API service is resolving properly but not for script resolution": {
			expected:        "HTTP request failed",
			server:          testutils.NewServer(0, nil),
			labelResolution: true,
			creatable:       false,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(cmd.Name, testutils.TEST)
			viper.Set(cmd.Description, testutils.TEST)
			viper.Set(cmd.Code, testutils.TEST)
			viper.Set(cmd.Parameters, "param1, param2")
			viper.Set(cmd.Labels, testutils.TEST)

			config.APIEndpoint = testutils.TEST
			if test.server != nil {
				mux := http.NewServeMux()

				variants := map[string]testutils.Value{"GET": {Status: http.StatusOK, V: []*types.Label{}}}
				if test.labelResolution {
					variants["POST"] = testutils.Value{Status: http.StatusOK, V: types.Label{ID: ""}}
				} else {
					variants["POST"] = testutils.Value{Status: http.StatusConflict, V: nil}
				}
				testutils.AddHandleFuncMultiple(mux, api.PathLabels, variants)

				if test.creatable {
					testutils.AddHandleFunc(mux, api.PathBlueprintScripts, http.StatusOK, new(types.Script))
				} else {
					testutils.AddHandleFunc(mux, api.PathBlueprintScripts, http.StatusConflict, nil)
				}
				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := ScriptCreate()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestScriptUpdate(t *testing.T) {
	tests := map[string]struct {
		expected        any
		serverRequired  bool
		labelResolution bool
	}{
		"if defined endpoint for API service is resolving properly": {
			serverRequired:  true,
			labelResolution: true,
		},
		"if defined endpoint for API service is invalid or cannot be reached": {
			expected: "Cannot execute request",
		},
		"if defined endpoint for API service is resolving properly but not for labelling resolution": {
			expected:       "HTTP request failed",
			serverRequired: true,
		},
	}

	config := testutils.InitConfig()
	configuration.SetConfig(config)
	cmd.RootCmd.SetContext(context.Background())

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			viper.Set(cmd.Id, testutils.TEST)
			viper.Set(cmd.Name, testutils.TEST)
			viper.Set(cmd.Description, testutils.TEST)
			viper.Set(cmd.Code, testutils.TEST)
			viper.Set(cmd.Parameters, "param1, param2")

			config.APIEndpoint = testutils.TEST
			if test.serverRequired {
				mux := http.NewServeMux()
				testutils.AddHandleFunc(mux, fmt.Sprintf(api.PathBlueprintScript, testutils.TEST), http.StatusOK, new(types.Script))
				if test.labelResolution {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusOK, []*types.Label{})
				} else {
					testutils.AddHandleFunc(mux, api.PathLabels, http.StatusConflict, nil)
				}

				server := httptest.NewServer(mux)
				defer server.Close()
				config.APIEndpoint = server.URL
			}
			err := ScriptUpdate()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestScriptDelete(t *testing.T) {
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

			err := ScriptDelete()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

// TODO
//func TestScriptAttachmentAdd(t *testing.T) {
//	tests := map[string]struct {
//		expected any
//		server   *httptest.Server
//	}{
//		"if defined endpoint for API service is resolving properly": {
//			expected: new(types.Attachment),
//			server:   testutils.NewServer(http.StatusOK, new(types.Attachment)),
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
//			viper.Set(cmd.Id, testutils.TEST)
//			viper.Set(cmd.Filepath, "../../testdata/test_utils.txt")
//			viper.Set(cmd.Name, testutils.TEST)
//
//			config.APIEndpoint = testutils.TEST
//			if test.server != nil {
//				mux := http.NewServeMux()
//				mux.HandleFunc(fmt.Sprintf("/blueprint/scripts/%s/attachments", testutils.TEST), func(res http.ResponseWriter, req *http.Request) {
//					res.WriteHeader(http.StatusOK)
//					d, _ := json.Marshal(&types.Attachment{
//						UploadURL: "URL",
//					})
//					res.Write(d)
//				})
//				//mux.HandleFunc("/labels", func(res http.ResponseWriter, req *http.Request) {
//				//	res.WriteHeader(http.StatusOK)
//				//	var d []byte
//				//	if req.Method == "GET" {
//				//		d, _ = json.Marshal([]*types.Label{})
//				//	}
//				//	if req.Method == "POST" {
//				//		d, _ = json.Marshal(types.Label{ID: ""})
//				//	}
//				//
//				//	res.Write(d)
//				//})
//				server := httptest.NewServer(mux)
//				defer server.Close()
//				config.APIEndpoint = server.URL
//			}
//			err := ScriptAttachmentAdd()
//			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
//				t.Errorf("Unexpected error: %v\n", err)
//			}
//		})
//	}
//}

func TestCleanAttachment(t *testing.T) {
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

			err := cleanAttachment(context.Background(), svc, formatter, testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestScriptAttachmentList(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Attachment{},
			server:   testutils.NewServer(http.StatusOK, []*types.Attachment{}),
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
			err := ScriptAttachmentList()
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}
