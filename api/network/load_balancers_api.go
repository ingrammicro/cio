// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathNetworkLoadBalancers = "/network/load_balancers"
const APIPathNetworkLoadBalancer = "/network/load_balancers/%s"
const APIPathNetworkLoadBalancerRetry = "/network/load_balancers/%s/retry"
const APIPathNetworkLoadBalancerPlan = "/network/load_balancer_plans/%s"

// LoadBalancerService manages load balancer operations
type LoadBalancerService struct {
	concertoService utils.ConcertoService
}

// NewLoadBalancerService returns a Concerto load balancer service
func NewLoadBalancerService(concertoService utils.ConcertoService) (*LoadBalancerService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &LoadBalancerService{
		concertoService: concertoService,
	}, nil
}

// ListLoadBalancers returns the list of load balancers as an array of LoadBalancer
func (lbs *LoadBalancerService) ListLoadBalancers() (loadBalancers []*types.LoadBalancer, err error) {
	log.Debug("ListLoadBalancers")

	data, status, err := lbs.concertoService.Get(APIPathNetworkLoadBalancers)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancers); err != nil {
		return nil, err
	}

	return loadBalancers, nil
}

// GetLoadBalancer returns a load balancer by its ID
func (lbs *LoadBalancerService) GetLoadBalancer(loadBalancerID string) (loadBalancer *types.LoadBalancer, err error) {
	log.Debug("GetLoadBalancer")

	data, status, err := lbs.concertoService.Get(fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// CreateLoadBalancer creates a load balancer
func (lbs *LoadBalancerService) CreateLoadBalancer(
	loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	log.Debug("CreateLoadBalancer")

	data, status, err := lbs.concertoService.Post(APIPathNetworkLoadBalancers, loadBalancerParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// UpdateLoadBalancer updates a load balancer by its ID
func (lbs *LoadBalancerService) UpdateLoadBalancer(
	loadBalancerID string,
	loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	log.Debug("UpdateLoadBalancer")

	data, status, err := lbs.concertoService.Put(
		fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerID),
		loadBalancerParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// DeleteLoadBalancer deletes a load balancer by its ID
func (lbs *LoadBalancerService) DeleteLoadBalancer(
	loadBalancerID string,
) (loadBalancer *types.LoadBalancer, err error) {
	log.Debug("DeleteLoadBalancer")

	data, status, err := lbs.concertoService.Delete(fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// RetryLoadBalancer retries a load balancer by its ID
func (lbs *LoadBalancerService) RetryLoadBalancer(
	loadBalancerID string,
	loadBalancerParams *map[string]interface{},
) (loadBalancer *types.LoadBalancer, err error) {
	log.Debug("RetryLoadBalancer")

	data, status, err := lbs.concertoService.Put(
		fmt.Sprintf(APIPathNetworkLoadBalancerRetry, loadBalancerID),
		loadBalancerParams,
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// GetLoadBalancerPlan returns a load balancer plan by its ID
func (lbs *LoadBalancerService) GetLoadBalancerPlan(
	loadBalancerPlanID string,
) (loadBalancerPlan *types.LoadBalancerPlan, err error) {
	log.Debug("GetLoadBalancerPlan")

	data, status, err := lbs.concertoService.Get(fmt.Sprintf(APIPathNetworkLoadBalancerPlan, loadBalancerPlanID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancerPlan); err != nil {
		return nil, err
	}

	return loadBalancerPlan, nil
}
