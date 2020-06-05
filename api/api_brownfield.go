// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListBrownfieldCloudAccounts returns the list of Brownfield Cloud Accounts as an array of CloudAccount
func (imco *IMCOClient) ListBrownfieldCloudAccounts() (cloudAccounts []*types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathBrownfieldCloudAccounts, true, &cloudAccounts)
	if err != nil {
		return nil, err
	}
	return cloudAccounts, nil
}

// GetBrownfieldCloudAccount returns a Brownfield Cloud Account by its ID
func (imco *IMCOClient) GetBrownfieldCloudAccount(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathBrownfieldCloudAccount, cloudAccountID), true, &cloudAccount)
	if err != nil {
		return nil, err
	}
	return cloudAccount, nil
}

// ImportServers imports brownfield servers
func (imco *IMCOClient) ImportServers(cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) ImportVPCs(cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) ImportFloatingIPs(cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) ImportVolumes(cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) ImportKubernetesClusters(cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
func (imco *IMCOClient) ImportPolicies(cloudAccountID string, params *map[string]interface{},
) (cloudAccount *types.CloudAccount, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(
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
