// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCloudProviderServerPlansByRealm = "/cloud/cloud_providers/%s/server_plans?realm_id=%s"
const APIPathCloudServerPlans = "/cloud/server_plans/%s"

// ServerPlanService manages server plan operations
type ServerPlanService struct {
	concertoService utils.ConcertoService
}

// NewServerPlanService returns a Concerto serverPlan service
func NewServerPlanService(concertoService utils.ConcertoService) (*ServerPlanService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ServerPlanService{
		concertoService: concertoService,
	}, nil
}

// ListServerPlans returns the list of serverPlans as an array of ServerPlan
func (sps *ServerPlanService) ListServerPlans(providerID string, realmID string) (serverPlans []*types.ServerPlan, err error) {
	log.Debug("ListServerPlans")

	data, status, err := sps.concertoService.Get(fmt.Sprintf(APIPathCloudProviderServerPlansByRealm, providerID, realmID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverPlans); err != nil {
		return nil, err
	}

	return serverPlans, nil
}

// GetServerPlan returns a serverPlan by its ID
func (sps *ServerPlanService) GetServerPlan(planID string) (serverPlan *types.ServerPlan, err error) {
	log.Debug("GetServerPlan")

	data, status, err := sps.concertoService.Get(fmt.Sprintf(APIPathCloudServerPlans, planID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverPlan); err != nil {
		return nil, err
	}

	return serverPlan, nil
}
