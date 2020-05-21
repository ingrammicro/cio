package testdata

import "github.com/ingrammicro/cio/api/types"

// GetLoadBalancerData loads test data
func GetLoadBalancerData() []*types.LoadBalancer {

	return []*types.LoadBalancer{
		{
			ID:                 "fakeID0",
			Name:               "fakeName0",
			State:              "fakeState0",
			RemoteID:           "fakeRemoteID0",
			CloudAccountID:     "fakeCloudAccountID0",
			RealmID:            "fakeRealmID0",
			VpcID:              "fakeVpcID0",
			LoadBalancerPlanID: "fakeLoadBalancerPlanID0",
			DnsName:            "fakeDnsName0",
			GlobalState:        "fakeGlobalState0",
			ErrorEventID:       "fakeErrorEventID0",
		},
	}
}

// GetLoadBalancerPlanData loads test data
func GetLoadBalancerPlanData() []*types.LoadBalancerPlan {

	return []*types.LoadBalancerPlan{
		{
			ID:                                   "fakeID0",
			Name:                                 "fakeName0",
			CloudProviderID:                      "fakeCloudProviderID0",
			CloudProviderName:                    "fakeCloudProviderName0",
			RealmID:                              "fakeRealmID0",
			RealmProviderName:                    "fakeRealmProviderName0",
			FlavourProviderName:                  "fakeFlavourProviderName0",
			Protocols:                            []string{"fakeProtocol00", "fakeProtocol01"},
			HealthCheckProtocols:                 []string{"fakeHealthCheckProtocol00", "fakeHealthCheckProtocol01"},
			RuleFields:                           []string{"fakeRuleField00", "fakeRuleField01"},
			HealthCheckIntervalValidValues:       map[string]interface{}{"fakeHealthCheckIntervalValidValue00": "x", "fakeHealthCheckIntervalValidValue01": "y"},
			HealthCheckTimeoutValidValues:        map[string]interface{}{"fakeHealthCheckTimeoutValidValue00": "x", "fakeHealthCheckTimeoutValidValue01": "y"},
			HealthCheckThresholdCountValidValues: map[string]interface{}{"fakeHealthCheckThresholdCountValidValue00": "x", "fakeHealthCheckThresholdCountValidValue01": "y"},
			Deprecated:                           false,
		},
	}
}

// GetTargetGroupData loads test data
func GetTargetGroupData() []*types.TargetGroup {

	return []*types.TargetGroup{
		{
			ID:                        "fakeID0",
			Name:                      "fakeName0",
			State:                     "fakeState0",
			RemoteID:                  "fakeRemoteID0",
			Protocol:                  "fakeProtocol0",
			Port:                      8080,
			Stickiness:                false,
			HealthCheckProtocol:       "fakeHealthCheckProtocol0",
			HealthCheckPort:           8081,
			HealthCheckInterval:       10,
			HealthCheckTimeout:        0,
			HealthCheckThresholdCount: 5,
			HealthCheckPath:           "fakeHealthCheckPath0",
			LoadBalancerID:            "fakeLoadBalancerID0",
			ErrorEventID:              "fakeErrorEventID0",
		},
	}
}

// GetTargetData loads test data
func GetTargetData() []*types.Target {

	return []*types.Target{
		{
			ID:           "fakeID0",
			ResourceType: "fakeResourceType0",
		},
	}
}

// GetListenerData loads test data
func GetListenerData() []*types.Listener {

	return []*types.Listener{
		{
			ID:                   "fakeID0",
			State:                "fakeState0",
			RemoteID:             "fakeRemoteID0",
			Protocol:             "fakeProtocol0",
			Port:                 8080,
			LoadBalancerID:       "fakeLoadBalancerID0",
			CertificateID:        "fakeCertificateID0",
			DefaultTargetGroupID: "fakeDefaultTargetGroupID0",
			ErrorEventID:         "fakeErrorEventID0",
		},
	}
}

// GetListenerRuleData loads test data
func GetListenerRuleData() []*types.ListenerRule {

	return []*types.ListenerRule{
		{
			ID:            "fakeID0",
			Field:         "fakeField0",
			Values:        []string{"fakeValue00", "fakeValue01"},
			ListenerID:    "fakeListenerID0",
			TargetGroupID: "fakeTargetGroupID0",
		},
	}
}

// GetCertificateData loads test data
func GetCertificateData() []*types.Certificate {

	return []*types.Certificate{
		{
			ID:             "fakeID0",
			Name:           "fakeName0",
			PublicKey:      "fakePublicKey0",
			Chain:          "fakeChain0",
			PrivateKey:     "fakePrivateKey0",
			LoadBalancerID: "fakeLoadBalancerID0",
		},
	}
}
