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

func TestGetAttachment(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Attachment),
			server:   NewServer(http.StatusOK, new(types.Attachment)),
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

			attachment, err := svc.GetAttachment(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(attachment, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", attachment, test.expected)
			}
		})
	}
}

func TestDeleteAttachment(t *testing.T) {
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

			err := svc.DeleteAttachment(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListCookbookVersions(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.CookbookVersion{},
			server:   NewServer(http.StatusOK, []*types.CookbookVersion{}),
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

			cookbookVersions, err := svc.ListCookbookVersions(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cookbookVersions, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cookbookVersions, test.expected)
			}
		})
	}
}

func TestGetCookbookVersion(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CookbookVersion),
			server:   NewServer(http.StatusOK, new(types.CookbookVersion)),
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

			cookbookVersion, err := svc.GetCookbookVersion(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cookbookVersion, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cookbookVersion, test.expected)
			}
		})
	}
}

func TestCreateCookbookVersion(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CookbookVersion),
			server:   NewServer(http.StatusOK, new(types.CookbookVersion)),
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

			cookbookVersion, err := svc.CreateCookbookVersion(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cookbookVersion, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cookbookVersion, test.expected)
			}
		})
	}
}

func TestProcessCookbookVersion(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.CookbookVersion),
			server:   NewServer(http.StatusOK, new(types.CookbookVersion)),
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

			cookbookVersion, err := svc.ProcessCookbookVersion(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(cookbookVersion, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", cookbookVersion, test.expected)
			}
		})
	}
}

func TestDeleteCookbookVersion(t *testing.T) {
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

			err := svc.DeleteCookbookVersion(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListScripts(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Script{},
			server:   NewServer(http.StatusOK, []*types.Script{}),
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

			scripts, err := svc.ListScripts(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(scripts, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", scripts, test.expected)
			}
		})
	}
}

func TestGetScript(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Script),
			server:   NewServer(http.StatusOK, new(types.Script)),
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

			script, err := svc.GetScript(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(script, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", script, test.expected)
			}
		})
	}
}

func TestCreateScript(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Script),
			server:   NewServer(http.StatusOK, new(types.Script)),
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

			script, err := svc.CreateScript(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(script, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", script, test.expected)
			}
		})
	}
}

func TestUpdateScript(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Script),
			server:   NewServer(http.StatusOK, new(types.Script)),
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

			script, err := svc.UpdateScript(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(script, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", script, test.expected)
			}
		})
	}
}

func TestDeleteScript(t *testing.T) {
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

			err := svc.DeleteScript(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestAddScriptAttachment(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Attachment),
			server:   NewServer(http.StatusOK, new(types.Attachment)),
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

			attachment, err := svc.AddScriptAttachment(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(attachment, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", attachment, test.expected)
			}
		})
	}
}

func TestUploadedScriptAttachment(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Attachment),
			server:   NewServer(http.StatusOK, new(types.Attachment)),
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

			attachment, err := svc.UploadedScriptAttachment(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(attachment, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", attachment, test.expected)
			}
		})
	}
}

func TestListScriptAttachments(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Attachment{},
			server:   NewServer(http.StatusOK, []*types.Attachment{}),
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

			attachments, err := svc.ListScriptAttachments(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(attachments, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", attachments, test.expected)
			}
		})
	}
}

func TestListTemplates(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.Template{},
			server:   NewServer(http.StatusOK, []*types.Template{}),
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

			templates, err := svc.ListTemplates(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(templates, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", templates, test.expected)
			}
		})
	}
}

func TestGetTemplate(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Template),
			server:   NewServer(http.StatusOK, new(types.Template)),
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

			template, err := svc.GetTemplate(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(template, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", template, test.expected)
			}
		})
	}
}

func TestCreateTemplate(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Template),
			server:   NewServer(http.StatusOK, new(types.Template)),
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

			template, err := svc.CreateTemplate(context.Background(), new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(template, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", template, test.expected)
			}
		})
	}
}

func TestUpdateTemplate(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Template),
			server:   NewServer(http.StatusOK, new(types.Template)),
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

			template, err := svc.UpdateTemplate(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(template, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", template, test.expected)
			}
		})
	}
}

func TestCompileTemplate(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.Template),
			server:   NewServer(http.StatusOK, new(types.Template)),
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

			template, err := svc.CompileTemplate(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(template, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", template, test.expected)
			}
		})
	}
}

func TestDeleteTemplate(t *testing.T) {
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

			err := svc.DeleteTemplate(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestListTemplateScripts(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.TemplateScript{},
			server:   NewServer(http.StatusOK, []*types.TemplateScript{}),
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

			templateScript, err := svc.ListTemplateScripts(context.Background(), TEST, TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(templateScript, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", templateScript, test.expected)
			}
		})
	}
}

func TestGetTemplateScript(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemplateScript),
			server:   NewServer(http.StatusOK, new(types.TemplateScript)),
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

			templateScript, err := svc.GetTemplateScript(context.Background(), TEST, TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(templateScript, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", templateScript, test.expected)
			}
		})
	}
}

func TestCreateTemplateScript(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemplateScript),
			server:   NewServer(http.StatusOK, new(types.TemplateScript)),
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

			templateScript, err := svc.CreateTemplateScript(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(templateScript, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", templateScript, test.expected)
			}
		})
	}
}

func TestUpdateTemplateScript(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: new(types.TemplateScript),
			server:   NewServer(http.StatusOK, new(types.TemplateScript)),
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

			templateScript, err := svc.UpdateTemplateScript(context.Background(), TEST, TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(templateScript, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", templateScript, test.expected)
			}
		})
	}
}

func TestDeleteTemplateScript(t *testing.T) {
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

			err := svc.DeleteTemplateScript(context.Background(), TEST, TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestReorderTemplateScript(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.TemplateScript{},
			server:   NewServer(http.StatusOK, []*types.TemplateScript{}),
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

			templateScript, err := svc.ReorderTemplateScript(context.Background(), TEST, new(map[string]interface{}))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(templateScript, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", templateScript, test.expected)
			}
		})
	}
}

func TestListTemplateServers(t *testing.T) {
	tests := map[string]struct {
		expected interface{}
		server   *httptest.Server
	}{
		"if defined endpoint for API service is resolving properly": {
			expected: []*types.TemplateServer{},
			server:   NewServer(http.StatusOK, []*types.TemplateServer{}),
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

			templateServer, err := svc.ListTemplateServers(context.Background(), TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(templateServer, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", templateServer, test.expected)
			}
		})
	}
}
