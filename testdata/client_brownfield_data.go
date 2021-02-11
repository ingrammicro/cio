package testdata

import (
	"github.com/ingrammicro/cio/api/types"
)

// GetBrownfieldCloudAccountsData loads test data
func GetBrownfieldCloudAccountsData() []*types.CloudAccount {
	return []*types.CloudAccount{
		{
			ID:                                  "fakeID0",
			Name:                                "fakeName0",
			SubscriptionID:                      "fakeSubscriptionID0",
			RemoteID:                            "fakeRemoteID0",
			CloudProviderID:                     "CloudProviderID0",
			CloudProviderName:                   "CloudProviderName0",
			SupportsImporting:                   true,
			SupportsImportingVPCs:               true,
			SupportsImportingFloatingIPs:        true,
			SupportsImportingVolumes:            true,
			SupportsImportingPolicies:           true,
			SupportsImportingKubernetesClusters: true,
			State:                               "fakeState0",
		},
		{
			ID:                                  "fakeID1",
			Name:                                "fakeName1",
			SubscriptionID:                      "fakeSubscriptionID1",
			RemoteID:                            "fakeRemoteID1",
			CloudProviderID:                     "CloudProviderID1",
			CloudProviderName:                   "CloudProviderName1",
			SupportsImporting:                   false,
			SupportsImportingVPCs:               false,
			SupportsImportingFloatingIPs:        false,
			SupportsImportingVolumes:            false,
			SupportsImportingPolicies:           false,
			SupportsImportingKubernetesClusters: false,
			State:                               "fakeState1",
		},
	}
}

// GetBrownfieldServerImportCandidatesData loads test data
func GetBrownfieldServerImportCandidatesData() []*types.ServerImportCandidate {
	return []*types.ServerImportCandidate{
		{
			ID:       "fakeID0",
			Name:     "fakeName0",
			Fqdn:     "fakeFqdn0",
			State:    "fakeState0",
			RemoteID: "fakeRemoteID0",
			Image: types.Image{
				Name:     "fakeImageName0",
				RemoteID: "fakeImageRemoteID0",
			},
			ServerPlanID:                "fakeServerPlanID0",
			PublicIP:                    "fakePublicIP0",
			CloudAccountID:              "fakeCloudAccountID0",
			FloatingIPsImportCandidates: nil,
			VolumesImportCandidates:     nil,
		},
		{
			ID:       "fakeID1",
			Name:     "fakeName1",
			Fqdn:     "fakeFqdn1",
			State:    "fakeState1",
			RemoteID: "fakeRemoteID1",
			Image: types.Image{
				Name:     "fakeImageName1",
				RemoteID: "fakeImageRemoteID1",
			},
			ServerPlanID:                "fakeServerPlanID1",
			PublicIP:                    "fakePublicIP1",
			CloudAccountID:              "fakeCloudAccountID1",
			FloatingIPsImportCandidates: nil,
			VolumesImportCandidates:     nil,
		},
	}
}

// GetBrownfieldVPCImportCandidatesData loads test data
func GetBrownfieldVPCImportCandidatesData() []*types.VpcImportCandidate {
	return []*types.VpcImportCandidate{
		{
			ID:                "fakeID0",
			Name:              "fakeName0",
			Cidr:              "fakeCidr0",
			RemoteID:          "fakeRemoteID0",
			CloudAccountID:    "fakeCloudAccountID0",
			RealmID:           "fakeRealmID0",
			SubnetsCandidates: nil,
		},
		{
			ID:                "fakeID1",
			Name:              "fakeName1",
			Cidr:              "fakeCidr1",
			RemoteID:          "fakeRemoteID1",
			CloudAccountID:    "fakeCloudAccountID1",
			RealmID:           "fakeRealmID1",
			SubnetsCandidates: nil,
		},
	}
}

// GetBrownfieldFloatingIPImportCandidatesData loads test data
func GetBrownfieldFloatingIPImportCandidatesData() []*types.FloatingIPImportCandidate {
	return []*types.FloatingIPImportCandidate{
		{
			ID:               "fakeID0",
			Name:             "fakeName0",
			Address:          "fakeAddress0",
			RemoteID:         "fakeRemoteID0",
			CloudAccountID:   "fakeCloudAccountID0",
			RealmID:          "fakeRealmID0",
			AttachedServerID: "fakeAttachedServerID0",
		},
		{
			ID:               "fakeID1",
			Name:             "fakeName1",
			Address:          "fakeAddress1",
			RemoteID:         "fakeRemoteID1",
			CloudAccountID:   "fakeCloudAccountID1",
			RealmID:          "fakeRealmID1",
			AttachedServerID: "fakeAttachedServerID1",
		},
	}
}

// GetBrownfieldVolumeImportCandidatesData loads test data
func GetBrownfieldVolumeImportCandidatesData() []*types.VolumeImportCandidate {
	return []*types.VolumeImportCandidate{
		{
			ID:               "fakeID0",
			Name:             "fakeName0",
			Size:             0,
			RemoteID:         "fakeRemoteID0",
			CloudAccountID:   "fakeCloudAccountID0",
			StoragePlanID:    "fakeStoragePlanID0",
			AttachedServerID: "fakeAttachedServerID0",
		},
		{
			ID:               "fakeID1",
			Name:             "fakeName1",
			Size:             1,
			RemoteID:         "fakeRemoteID1",
			CloudAccountID:   "fakeCloudAccountID1",
			StoragePlanID:    "fakeStoragePlanID1",
			AttachedServerID: "fakeAttachedServerID1",
		},
	}
}
