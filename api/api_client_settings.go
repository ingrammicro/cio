// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListCloudAccounts returns the list of cloudAccounts as an array of CloudAccount
func (imco *ClientAPI) ListCloudAccounts(ctx context.Context) (cloudAccounts []*types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathSettingsCloudAccounts, true, &cloudAccounts)
	if err != nil {
		return nil, err
	}
	return cloudAccounts, nil
}

// GetCloudAccount returns a cloudAccount by its ID
func (imco *ClientAPI) GetCloudAccount(ctx context.Context, cloudAccountID string,
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathSettingsCloudAccount, cloudAccountID), true, &cloudAccount)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ListPolicyAssignments returns the list of policy assignments as an array of PolicyAssignment
func (imco *ClientAPI) ListPolicyAssignments(ctx context.Context, cloudAccountID string,
) (assignments []*types.PolicyAssignment, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx,
		fmt.Sprintf(PathSettingsCloudAccountPolicyAssignments, cloudAccountID),
		true,
		&assignments,
	)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}
