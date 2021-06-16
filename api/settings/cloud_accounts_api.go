// Copyright (c) 2017-2021 Ingram Micro Inc.

package settings

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathSettingsCloudAccounts = "/settings/cloud_accounts"
const APIPathSettingsCloudAccount = "/settings/cloud_accounts/%s"

// CloudAccountService manages cloud account operations
type CloudAccountService struct {
	concertoService utils.ConcertoService
}

// NewCloudAccountService returns a Concerto cloudAccount service
func NewCloudAccountService(concertoService utils.ConcertoService) (*CloudAccountService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudAccountService{
		concertoService: concertoService,
	}, nil
}

// ListCloudAccounts returns the list of cloudAccounts as an array of CloudAccount
func (cas *CloudAccountService) ListCloudAccounts() (cloudAccounts []*types.CloudAccount, err error) {
	log.Debug("ListCloudAccounts")

	data, status, err := cas.concertoService.Get(APIPathSettingsCloudAccounts)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccounts); err != nil {
		return nil, err
	}

	return cloudAccounts, nil
}

// GetCloudAccount returns a cloudAccount by its ID
func (cas *CloudAccountService) GetCloudAccount(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("GetCloudAccount")

	data, status, err := cas.concertoService.Get(fmt.Sprintf(APIPathSettingsCloudAccount, cloudAccountID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}
