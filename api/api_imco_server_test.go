package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/internal/testutils"
	"github.com/ingrammicro/cio/types"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// TODO COMMON
func initConfigAndServerAPI(context configuration.Context) (*configuration.Config, *ServerAPI, error) {
	config := new(configuration.Config)
	if context == configuration.Brownfield {
		config.APIEndpoint = testutils.TEST
		config.BrownfieldToken = testutils.TEST
	}
	if context == configuration.Polling {
		config.APIEndpoint = testutils.TEST
		config.CommandPollingToken = testutils.TEST
		config.ServerID = testutils.TEST
	}

	svc, err := NewIMCOServerWithToken(config, context)
	if err != nil {
		return nil, nil, err
	}
	return config, svc, nil
}

func TestNewIMCOServerWithToken(t *testing.T) {
	tests := map[string]struct {
		context   int
		setConfig bool
		config    *configuration.Config
	}{
		"if config not initialized": {
			context:   0,
			setConfig: false,
			config:    nil,
		},
		"if BrownfieldToken configuration is incomplete": {
			context:   configuration.Brownfield,
			setConfig: false,
			config:    InitConfig(),
		},
		"if BrownfieldToken configuration is complete": {
			context:   configuration.Brownfield,
			setConfig: true,
			config:    InitConfig(),
		},
		"if CommandPollingToken configuration is incomplete": {
			context:   configuration.Polling,
			setConfig: false,
			config:    InitConfig(),
		},
		"if CommandPollingToken configuration is complete": {
			context:   configuration.Polling,
			setConfig: true,
			config:    InitConfig(),
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			if test.context == configuration.Brownfield && test.setConfig {
				test.config.APIEndpoint = testutils.TEST
				test.config.BrownfieldToken = testutils.TEST
			}
			if test.context == configuration.Polling && test.setConfig {
				test.config.APIEndpoint = testutils.TEST
				test.config.CommandPollingToken = testutils.TEST
				test.config.ServerID = testutils.TEST
			}

			svc, err := NewIMCOServerWithToken(test.config, test.context)
			if err != nil && test.setConfig {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && svc != nil && !test.setConfig {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestGetDispatcherScriptCharacterizationsByType(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a ScriptCharacterization type is returned successfully": {
			expected: []*types.ScriptCharacterization{},
			server:   testutils.NewServer(http.StatusOK, []*types.ScriptCharacterization{}),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			scriptCharacterizations, err := svc.GetDispatcherScriptCharacterizationsByType(context.Background(), "boot")
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(scriptCharacterizations, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", scriptCharacterizations, test.expected)
			}
		})
	}
}

func TestGetDispatcherScriptCharacterizationByUUID(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a ScriptCharacterization type is returned successfully": {
			expected: new(types.ScriptCharacterization),
			server:   testutils.NewServer(http.StatusOK, new(types.ScriptCharacterization)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			scriptCharacterization, err := svc.GetDispatcherScriptCharacterizationByUUID(context.Background(), testutils.TEST)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(scriptCharacterization, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", scriptCharacterization, test.expected)
			}
		})
	}
}

func TestReportScriptConclusions(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a ScriptConclusion type is returned successfully": {
			expected: new(types.ScriptConclusion),
			server:   testutils.NewServer(http.StatusOK, new(types.ScriptConclusion)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			scriptConclusion, status, err := svc.ReportScriptConclusions(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(scriptConclusion, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					scriptConclusion, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestGetBootstrappingConfiguration(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a BootstrappingConfiguration type is returned successfully": {
			expected: new(types.BootstrappingConfiguration),
			server:   testutils.NewServer(http.StatusOK, new(types.BootstrappingConfiguration)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			bootstrappingConfigurations, status, err := svc.GetBootstrappingConfiguration(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(bootstrappingConfigurations, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					bootstrappingConfigurations, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestReportBootstrappingAppliedConfiguration(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a BootstrappingConfiguration type is returned successfully": {
			expected: new(types.BootstrappingConfiguration),
			server:   testutils.NewServer(http.StatusOK, new(types.BootstrappingConfiguration)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			err := svc.ReportBootstrappingAppliedConfiguration(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
		})
	}
}

func TestReportBootstrappingLog(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a BootstrappingContinuousReport type is returned successfully": {
			expected: new(types.BootstrappingContinuousReport),
			server:   testutils.NewServer(http.StatusOK, new(types.BootstrappingContinuousReport)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			bootstrappingContinuousReport, status, err := svc.ReportBootstrappingLog(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(bootstrappingContinuousReport, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					bootstrappingContinuousReport, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestGetPolicy(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a Policy type is returned successfully": {
			expected: new(types.Policy),
			server:   testutils.NewServer(http.StatusOK, new(types.Policy)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			policy, err := svc.GetPolicy(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && policy.Md5 == "" {
				t.Errorf("Unexpected response: %v. Expected: %v\n", policy, test.expected)
			}
		})
	}
}

func TestAddPolicyRule(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a PolicyRule is added successfully": {
			expected: new(types.PolicyRule),
			server:   testutils.NewServer(http.StatusOK, new(types.PolicyRule)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			policyRule, err := svc.AddPolicyRule(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(policyRule, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", policyRule, test.expected)
			}
		})
	}
}

func TestUpdatePolicy(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a Policy type is returned successfully": {
			expected: new(types.Policy),
			server:   testutils.NewServer(http.StatusOK, new(types.Policy)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			policy, err := svc.UpdatePolicy(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && !reflect.DeepEqual(policy, test.expected) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", policy, test.expected)
			}
		})
	}
}

func TestPing(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a PollingPing type is returned successfully": {
			expected: new(types.PollingPing),
			server:   testutils.NewServer(http.StatusOK, new(types.PollingPing)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Polling)
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

			ping, status, err := svc.Ping(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(ping, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					ping, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestGetNextCommand(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a PollingCommand type is returned successfully": {
			expected: new(types.PollingCommand),
			server:   testutils.NewServer(http.StatusOK, new(types.PollingCommand)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Polling)
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

			command, status, err := svc.GetNextCommand(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(command, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					command, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestUpdateCommand(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a PollingCommand type is returned successfully": {
			expected: new(types.PollingCommand),
			server:   testutils.NewServer(http.StatusOK, new(types.PollingCommand)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Polling)
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

			command, status, err := svc.UpdateCommand(context.Background(), testutils.TEST, new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(command, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					command, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestReportBootstrapLog(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a PollingContinuousReport type is returned successfully": {
			expected: new(types.PollingContinuousReport),
			server:   testutils.NewServer(http.StatusOK, new(types.PollingContinuousReport)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Polling)
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

			command, status, err := svc.ReportBootstrapLog(context.Background(), new(map[string]any))
			t.Logf("command: %v status: %v err: %v\n", command, status, err)
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(command, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					command, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestObtainBrownfieldSslProfile(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a BrownfieldSslProfile type is returned successfully": {
			expected: make(map[string]any),
			server:   testutils.NewServer(http.StatusOK, make(map[string]any)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			response, status, err := svc.ObtainBrownfieldSslProfile(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(response, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					response, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestObtainPollingApiKey(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a PollingApiKey type is returned successfully": {
			expected: make(map[string]any),
			server:   testutils.NewServer(http.StatusOK, make(map[string]any)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Polling)
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

			response, status, err := svc.ObtainPollingApiKey(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(response, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					response, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestSetFirewallProfile(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a Firewall type is returned successfully": {
			expected: new(types.Firewall),
			server:   testutils.NewServer(http.StatusOK, new(types.Firewall)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			firewall, status, err := svc.SetFirewallProfile(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(firewall, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					firewall, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestGetBrownfieldSettings(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a Settings type is returned successfully": {
			expected: new(types.Settings),
			server:   testutils.NewServer(http.StatusOK, new(types.Settings)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			settings, status, err := svc.GetBrownfieldSettings(context.Background())
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(settings, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					settings, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestSetBrownfieldSettings(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a Settings type type is returned successfully": {
			expected: new(types.Settings),
			server:   testutils.NewServer(http.StatusOK, new(types.Settings)),
		},
		"if returns an error": {
			expected: "Cannot execute request",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			settings, status, err := svc.SetBrownfieldSettings(context.Background(), new(map[string]any))
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(settings, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v (%v). Expected: %v(%v)\n",
					settings, status, test.expected, http.StatusOK)
			}
		})
	}
}

func TestRetrieveSecretVersion(t *testing.T) {
	tests := map[string]struct {
		expected any
		server   *httptest.Server
	}{
		"if a secret type is downloaded successfully": {
			expected: http.StatusOK,
			server:   testutils.NewServer(http.StatusOK, nil),
		},
		"if returns an error": {
			expected: "Cannot download file",
			server:   nil,
		},
	}

	config, svc, err := initConfigAndServerAPI(configuration.Brownfield)
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

			status, err := svc.RetrieveSecretVersion(context.Background(), "1234567890", "testdata/test_utils.txt")
			if err != nil && !strings.Contains(err.Error(), fmt.Sprintf("%s", test.expected)) {
				t.Errorf("Unexpected error: %v\n", err)
			}
			if err == nil && (!reflect.DeepEqual(status, test.expected) || status != http.StatusOK) {
				t.Errorf("Unexpected response: %v. Expected: %v\n", status, test.expected)
			}
		})
	}
}
