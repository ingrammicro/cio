package testdata

import "github.com/ingrammicro/cio/api/types"

// GetDomainData loads test data
func GetDomainData() []*types.Domain {

	return []*types.Domain{
		{
			ID:             "fakeID0",
			Name:           "fakeName0",
			State:          "fakeState0",
			RemoteID:       "fakeRemoteID0",
			CloudAccountID: "fakeCloudAccountID0",
			Nameservers:    []string{"fakeNameserver0", "fakeNameserver1", "fakeNameserver2", "fakeNameserver3"},
			GlobalState:    "fakeGlobalState0",
			ErrorEventID:   "fakeErrorEventID0",
		},
	}
}

// GetRecordData loads test data
func GetRecordData() []*types.Record {

	return []*types.Record{
		{
			ID:             "fakeID0",
			Name:           "fakeName0",
			State:          "fakeState0",
			Content:        "fakeContent0",
			RemoteID:       "fakeRemoteID0",
			Type:           "fakeType0",
			TTL:            3600,
			DomainID:       "fakeDomainID0",
			InstanceID:     "fakeInstanceID0",
			FloatingIpID:   "fakeFloatingIpID0",
			LoadBalancerID: "fakeLoadBalancerID0",
			Priority:       0,
			Weight:         0,
			Port:           0,
			ErrorEventID:   "fakeErrorEventID0",
		},
	}
}
