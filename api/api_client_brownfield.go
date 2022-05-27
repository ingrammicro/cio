// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"golang.org/x/net/context"
)

// ListBrownfieldCloudAccounts returns the list of Brownfield Cloud Accounts as an array of CloudAccount
func (imco *ClientAPI) ListBrownfieldCloudAccounts(ctx context.Context,
) (cloudAccounts []*types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathBrownfieldCloudAccounts, true, &cloudAccounts)
	if err != nil {
		return nil, err
	}
	return cloudAccounts, nil
}

// GetBrownfieldCloudAccount returns a Brownfield Cloud Account by its ID
func (imco *ClientAPI) GetBrownfieldCloudAccount(ctx context.Context, cloudAccountID string,
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathBrownfieldCloudAccount, cloudAccountID), true, &cloudAccount)
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
		fmt.Sprintf(pathBrownfieldCloudAccountImportServers, cloudAccountID),
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
		fmt.Sprintf(pathBrownfieldCloudAccountImportVpcs, cloudAccountID),
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
		fmt.Sprintf(pathBrownfieldCloudAccountImportFloatingIPs, cloudAccountID),
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
		fmt.Sprintf(pathBrownfieldCloudAccountImportVolumes, cloudAccountID),
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
		fmt.Sprintf(pathBrownfieldCloudAccountImportKubernetesClusters, cloudAccountID),
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
		fmt.Sprintf(pathBrownfieldCloudAccountImportPolicies, cloudAccountID),
		params,
		true,
		&cloudAccount,
	)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}
