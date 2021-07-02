// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCloudProviderRealms = "/cloud/cloud_providers/%s/realms"
const APIPathCloudRealm = "/cloud/realms/%s"
const APIPathCloudRealmNodePoolPlans = "/cloud/realms/%s/node_pool_plans"

// RealmService manages realm operations
type RealmService struct {
	concertoService utils.ConcertoService
}

// NewRealmService returns a Concerto realm service
func NewRealmService(concertoService utils.ConcertoService) (*RealmService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &RealmService{
		concertoService: concertoService,
	}, nil
}

// ListRealms returns the list of realms as an array of Realm
func (rs *RealmService) ListRealms(providerID string) (realms []*types.Realm, err error) {
	log.Debug("ListRealms")

	data, status, err := rs.concertoService.Get(fmt.Sprintf(APIPathCloudProviderRealms, providerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &realms); err != nil {
		return nil, err
	}

	return realms, nil
}

// GetRealm returns a realm by its ID
func (rs *RealmService) GetRealm(realmID string) (realm *types.Realm, err error) {
	log.Debug("GetRealm")

	data, status, err := rs.concertoService.Get(fmt.Sprintf(APIPathCloudRealm, realmID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &realm); err != nil {
		return nil, err
	}

	return realm, nil
}

// ListRealmNodePoolPlans returns the list of node pool plans as an array of NodePoolPlan
func (rs *RealmService) ListRealmNodePoolPlans(realmID string) (nodePoolPlans []*types.NodePoolPlan, err error) {
	log.Debug("ListRealmNodePoolPlans")

	data, status, err := rs.concertoService.Get(fmt.Sprintf(APIPathCloudRealmNodePoolPlans, realmID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodePoolPlans); err != nil {
		return nil, err
	}

	return nodePoolPlans, nil
}
