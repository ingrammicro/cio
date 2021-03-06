// Copyright (c) 2017-2021 Ingram Micro Inc.

package blueprint

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathBlueprintConfiguration = "/blueprint/configuration"
const APIPathBlueprintAppliedConfiguration = "/blueprint/applied_configuration"
const APIPathBlueprintBootstrapLogs = "/blueprint/bootstrap_logs"

// BootstrappingService manages bootstrapping operations
type BootstrappingService struct {
	concertoService utils.ConcertoService
}

// NewBootstrappingService returns a bootstrapping service
func NewBootstrappingService(concertoService utils.ConcertoService) (*BootstrappingService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &BootstrappingService{
		concertoService: concertoService,
	}, nil

}

// GetBootstrappingConfiguration returns the list of policy files as a JSON response with the desired configuration
// changes
func (bs *BootstrappingService) GetBootstrappingConfiguration() (
	bootstrappingConfigurations *types.BootstrappingConfiguration, status int, err error,
) {
	log.Debug("GetBootstrappingConfiguration")

	data, status, err := bs.concertoService.Get(APIPathBlueprintConfiguration)
	if err != nil {
		return nil, status, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &bootstrappingConfigurations); err != nil {
		return nil, status, err
	}

	return bootstrappingConfigurations, status, nil
}

// ReportBootstrappingAppliedConfiguration informs the platform of applied changes
func (bs *BootstrappingService) ReportBootstrappingAppliedConfiguration(
	bootstrappingAppliedConfigurationParams *map[string]interface{},
) (err error) {
	log.Debug("ReportBootstrappingAppliedConfiguration")

	data, status, err := bs.concertoService.Put(APIPathBlueprintAppliedConfiguration,
		bootstrappingAppliedConfigurationParams)

	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ReportBootstrappingLog reports a policy files application result
func (bs *BootstrappingService) ReportBootstrappingLog(
	bootstrappingContinuousReportParams *map[string]interface{},
) (command *types.BootstrappingContinuousReport, status int, err error) {
	log.Debug("ReportBootstrappingLog")

	data, status, err := bs.concertoService.Post(APIPathBlueprintBootstrapLogs, bootstrappingContinuousReportParams)

	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}

// DownloadPolicyfile gets a file from given url saving file into given file path
func (bs *BootstrappingService) DownloadPolicyfile(
	url string,
	filePath string,
) (realFileName string, status int, err error) {
	log.Debug("DownloadPolicyfile")

	realFileName, status, err = bs.concertoService.GetFile(url, filePath, false)
	if err != nil {
		return realFileName, status, err
	}

	return realFileName, status, nil
}
