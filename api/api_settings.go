// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListCloudAccounts returns the list of cloudAccounts as an array of CloudAccount
func (imco *IMCOClient) ListCloudAccounts() (cloudAccounts []*types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathSettingsCloudAccounts, true, &cloudAccounts)
	if err != nil {
		return nil, err
	}
	return cloudAccounts, nil
}

// GetCloudAccount returns a cloudAccount by its ID
func (imco *IMCOClient) GetCloudAccount(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathSettingsCloudAccount, cloudAccountID), true, &cloudAccount)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ListPolicyAssignments returns the list of policy assignments as an array of PolicyAssignment
func (imco *IMCOClient) ListPolicyAssignments(cloudAccountID string,
) (assignments []*types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(
		fmt.Sprintf(pathSettingsCloudAccountPolicyAssignments, cloudAccountID),
		true,
		&assignments,
	)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}
