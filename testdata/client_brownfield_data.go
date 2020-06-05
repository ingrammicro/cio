// Copyright (c) 2017-2022 Ingram Micro Inc.

package testdata

import (
	"github.com/ingrammicro/cio/types"
)

// GetBrownfieldCloudAccountsData loads test data
func GetBrownfieldCloudAccountsData() []*types.CloudAccount {
	return []*types.CloudAccount{
		{
			ID:                           "fakeID0",
			Name:                         "fakeName0",
			SubscriptionID:               "fakeSubscriptionID0",
			RemoteID:                     "fakeRemoteID0",
			CloudProviderID:              "CloudProviderID0",
			CloudProviderName:            "CloudProviderName0",
			SupportsImporting:            true,
			SupportsImportingVPCs:        true,
			SupportsImportingFloatingIPs: true,
			SupportsImportingVolumes:     true,
		},
		{
			ID:                           "fakeID1",
			Name:                         "fakeName1",
			SubscriptionID:               "fakeSubscriptionID1",
			RemoteID:                     "fakeRemoteID1",
			CloudProviderID:              "CloudProviderID1",
			CloudProviderName:            "CloudProviderName1",
			SupportsImporting:            false,
			SupportsImportingVPCs:        false,
			SupportsImportingFloatingIPs: false,
			SupportsImportingVolumes:     false,
		},
	}
}
