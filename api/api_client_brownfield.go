// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"context"
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListBrownfieldCloudAccounts returns the list of Brownfield Cloud Accounts as an array of CloudAccount
func (imco *ClientAPI) ListBrownfieldCloudAccounts(ctx context.Context,
) (cloudAccounts []*types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, PathBrownfieldCloudAccounts, true, &cloudAccounts)
	if err != nil {
		return nil, err
	}
	return cloudAccounts, nil
}

// GetBrownfieldCloudAccount returns a Brownfield Cloud Account by its ID
func (imco *ClientAPI) GetBrownfieldCloudAccount(ctx context.Context, cloudAccountID string,
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(PathBrownfieldCloudAccount, cloudAccountID), true, &cloudAccount)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ImportServers imports brownfield servers
func (imco *ClientAPI) ImportServers(ctx context.Context, cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBrownfieldCloudAccountImportServers, cloudAccountID),
		params,
		true,
		&cloudAccount,
	)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ImportVPCs imports brownfield vpcs
func (imco *ClientAPI) ImportVPCs(ctx context.Context, cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBrownfieldCloudAccountImportVpcs, cloudAccountID),
		params,
		true,
		&cloudAccount,
	)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ImportFloatingIPs imports brownfield floating ips
func (imco *ClientAPI) ImportFloatingIPs(ctx context.Context, cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBrownfieldCloudAccountImportFloatingIPs, cloudAccountID),
		params,
		true,
		&cloudAccount,
	)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ImportVolumes imports brownfield volumes
func (imco *ClientAPI) ImportVolumes(ctx context.Context, cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBrownfieldCloudAccountImportVolumes, cloudAccountID),
		params,
		true,
		&cloudAccount,
	)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ImportKubernetesClusters imports brownfield kubernetes clusters
func (imco *ClientAPI) ImportKubernetesClusters(ctx context.Context, cloudAccountID string,
	params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBrownfieldCloudAccountImportKubernetesClusters, cloudAccountID),
		params,
		true,
		&cloudAccount,
	)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ImportPolicies imports brownfield policies
func (imco *ClientAPI) ImportPolicies(ctx context.Context, cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx,
		fmt.Sprintf(PathBrownfieldCloudAccountImportPolicies, cloudAccountID),
		params,
		true,
		&cloudAccount,
	)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}
