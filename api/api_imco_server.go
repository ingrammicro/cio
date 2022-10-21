// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/configuration"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	pathBlueprintScriptCharacterizationsType = "/blueprint/script_characterizations?type=%s"
	pathBlueprintScriptCharacterization      = "/blueprint/script_characterizations/%s"
	pathBlueprintScriptConclusions           = "/blueprint/script_conclusions"
	pathBlueprintConfiguration               = "/blueprint/configuration"
	pathBlueprintAppliedConfiguration        = "/blueprint/applied_configuration"
	pathBlueprintBootstrapLogs               = "/blueprint/bootstrap_logs"
	pathCloudFirewallProfile                 = "/cloud/firewall_profile"
	pathCloudFirewallProfileRules            = "/cloud/firewall_profile/rules"
	pathCommandPollingPings                  = "/command_polling/pings"
	pathCommandPollingNextCommand            = "/command_polling/command"
	pathCommandPollingCommand                = "/command_polling/commands/%s"
	pathCommandPollingBootstrapLogs          = "/command_polling/bootstrap_logs"

	pathBrownfieldSslProfile = "/brownfield/ssl_profile"
	pathCommandPollingApiKey = "/command_polling/api_key"
	pathBrownfieldSettings   = "/brownfield/settings"
	pathSecretVersionContent = "/secret/secret_versions/%s"
)

// ServerAPI web service manager
type ServerAPI struct {
	HTTPClient
}

// NewIMCOServerWithToken creates new http IMCO cli based on config
func NewIMCOServerWithToken(config *configuration.Config, context configuration.Context) (svc *ServerAPI, err error) {
	if config == nil {
		return nil, fmt.Errorf(WebServiceConfigurationFailed)
	}

	if context == configuration.Brownfield && !config.IsConfigReadyBrownfield() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	if context == configuration.Polling && !config.IsConfigReadyCommandPolling() {
		return nil, fmt.Errorf(ConfigurationIsIncomplete)
	}

	// creates HTTP service with config
	svc = &ServerAPI{
		HTTPClient{
			config: config,
		},
	}

	// Creates a client with no certificates and insecure option
	svc.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * time.Duration(HttpTimeOut),
	}
	return svc, nil
}

// GetDispatcherScriptCharacterizationsByType returns script characterizations list for a given phase
func (imco *ServerAPI) GetDispatcherScriptCharacterizationsByType(ctx context.Context, phase string,
) (scriptCharacterizations []*types.ScriptCharacterization, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(pathBlueprintScriptCharacterizationsType, phase),
		true,
		&scriptCharacterizations,
	)
	if err != nil {
		return nil, err
	}
	return scriptCharacterizations, nil
}

// GetDispatcherScriptCharacterizationByUUID returns script characterizations list for a given UUID
func (imco *ServerAPI) GetDispatcherScriptCharacterizationByUUID(ctx context.Context,
	scriptCharacterizationUUID string,
) (scriptCharacterization *types.ScriptCharacterization, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(pathBlueprintScriptCharacterization, scriptCharacterizationUUID),
		true,
		&scriptCharacterization,
	)
	if err != nil {
		return nil, err
	}
	return scriptCharacterization, nil
}

// ReportScriptConclusions reports a result
func (imco *ServerAPI) ReportScriptConclusions(ctx context.Context, scriptConclusions *map[string]interface{},
) (command *types.ScriptConclusion, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.PostAndCheck(ctx, pathBlueprintScriptConclusions, scriptConclusions, true, &command)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// GetBootstrappingConfiguration returns the list of policy files as a JSON response with the desired configuration
// changes
func (imco *ServerAPI) GetBootstrappingConfiguration(ctx context.Context) (
	bootstrappingConfigurations *types.BootstrappingConfiguration, status int, err error,
) {
	logger.DebugFuncInfo()

	status, err = imco.GetAndCheck(ctx, pathBlueprintConfiguration, true, &bootstrappingConfigurations)
	if err != nil {
		return nil, status, err
	}
	return bootstrappingConfigurations, status, nil
}

// ReportBootstrappingAppliedConfiguration informs the platform of applied changes
func (imco *ServerAPI) ReportBootstrappingAppliedConfiguration(ctx context.Context,
	bootstrappingAppliedConfigurationParams *map[string]interface{},
) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		pathBlueprintAppliedConfiguration,
		bootstrappingAppliedConfigurationParams,
		true,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

// ReportBootstrappingLog reports a policy files application result
func (imco *ServerAPI) ReportBootstrappingLog(ctx context.Context,
	bootstrappingContinuousReportParams *map[string]interface{},
) (command *types.BootstrappingContinuousReport, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.PostAndCheck(ctx, pathBlueprintBootstrapLogs, bootstrappingContinuousReportParams, false, &command)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// GetPolicy returns firewall policy
func (imco *ServerAPI) GetPolicy(ctx context.Context) (policy *types.Policy, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathCloudFirewallProfile, true, &policy)
	if err != nil {
		return nil, err
	}

	var data []byte
	if data, err = json.Marshal(policy); err != nil {
		return nil, err
	}
	policy.Md5 = fmt.Sprintf("%x", md5.Sum(data))
	return policy, nil
}

// AddPolicyRule adds a new firewall policy rule
func (imco *ServerAPI) AddPolicyRule(ctx context.Context, ruleParams *map[string]interface{},
) (policyRule *types.PolicyRule, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathCloudFirewallProfileRules, ruleParams, true, &policyRule)
	if err != nil {
		return nil, err
	}
	return policyRule, nil
}

// UpdatePolicy updates firewall policy
func (imco *ServerAPI) UpdatePolicy(ctx context.Context, policyParams *map[string]interface{},
) (policy *types.Policy, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, pathCloudFirewallProfile, policyParams, true, &policy)
	if err != nil {
		return nil, err
	}
	return policy, nil
}

// Ping resolves if new command is waiting for execution
func (imco *ServerAPI) Ping(ctx context.Context) (ping *types.PollingPing, status int, err error) {
	logger.DebugFuncInfo()

	payload := make(map[string]interface{})
	status, err = imco.PostAndCheck(ctx, pathCommandPollingPings, &payload, false, &ping)
	if err != nil {
		return nil, status, err
	}
	return ping, status, nil
}

// GetNextCommand returns the command to be executed
func (imco *ServerAPI) GetNextCommand(ctx context.Context) (command *types.PollingCommand, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.GetAndCheck(ctx, pathCommandPollingNextCommand, false, &command)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// UpdateCommand updates a command by its ID
func (imco *ServerAPI) UpdateCommand(ctx context.Context, commandID string,
	pollingCommandParams *map[string]interface{},
) (command *types.PollingCommand, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(pathCommandPollingCommand, commandID),
		pollingCommandParams,
		false,
		&command,
	)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// ReportBootstrapLog reports a command result
func (imco *ServerAPI) ReportBootstrapLog(ctx context.Context, pollingContinuousReportParams *map[string]interface{},
) (command *types.PollingContinuousReport, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.PostAndCheck(ctx,
		pathCommandPollingBootstrapLogs,
		pollingContinuousReportParams,
		false,
		&command,
	)
	if err != nil {
		return nil, status, err
	}
	return command, status, nil
}

// ObtainBrownfieldSslProfile obtains server brownfield ssl profile
func (imco *ServerAPI) ObtainBrownfieldSslProfile(ctx context.Context, payload *map[string]interface{},
) (response map[string]interface{}, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.PostAndCheck(ctx, pathBrownfieldSslProfile, payload, false, &response)
	if err != nil {
		return nil, status, err
	}
	return response, status, nil
}

// ObtainPollingApiKey obtains server polling api key
func (imco *ServerAPI) ObtainPollingApiKey(ctx context.Context, payload *map[string]interface{},
) (response map[string]interface{}, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.PostAndCheck(ctx, pathCommandPollingApiKey, payload, false, &response)
	if err != nil {
		return nil, status, err
	}
	return response, status, nil
}

// SetFirewallProfile creates firewall policy
func (imco *ServerAPI) SetFirewallProfile(ctx context.Context, policyParams *map[string]interface{},
) (firewall *types.Firewall, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.PostAndCheck(ctx, pathCloudFirewallProfile, policyParams, false, &firewall)
	if err != nil {
		return nil, status, err
	}
	return firewall, status, nil
}

// GetBrownfieldSettings obtains brownfield settings
func (imco *ServerAPI) GetBrownfieldSettings(ctx context.Context) (settings *types.Settings, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.GetAndCheck(ctx, pathBrownfieldSettings, false, &settings)
	if err != nil {
		return nil, status, err
	}
	return settings, status, nil
}

// SetBrownfieldSettings sets brownfield settings
func (imco *ServerAPI) SetBrownfieldSettings(ctx context.Context, payload *map[string]interface{},
) (settings *types.Settings, status int, err error) {
	logger.DebugFuncInfo()

	status, err = imco.PutAndCheck(ctx, pathBrownfieldSettings, payload, false, &settings)
	if err != nil {
		return nil, status, err
	}
	return settings, status, nil
}

// RetrieveSecretVersion returns script characterizations list for a given UUID
func (imco *ServerAPI) RetrieveSecretVersion(ctx context.Context, svID, filePath string) (int, error) {
	log.Debug("RetrieveSecretVersion")

	_, status, err := imco.DownloadFile(ctx,
		fmt.Sprintf("%s"+pathSecretVersionContent, imco.config.APIEndpoint, svID),
		filePath,
		false)
	if err != nil {
		return status, err
	}
	return status, nil
}
