// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCloudProviders = "/cloud/cloud_providers"
const APIPathCloudProviderStoragePlans = "/cloud/cloud_providers/%s/storage_plans"
const APIPathCloudProviderLoadBalancerPlans = "/cloud/cloud_providers/%s/load_balancer_plans"
const APIPathCloudProviderClusterPlans = "/cloud/cloud_providers/%s/cluster_plans"

// CloudProviderService manages cloud provider operations
type CloudProviderService struct {
	concertoService utils.ConcertoService
}

// NewCloudProviderService returns a Concerto cloudProvider service
func NewCloudProviderService(concertoService utils.ConcertoService) (*CloudProviderService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CloudProviderService{
		concertoService: concertoService,
	}, nil
}

// ListCloudProviders returns the list of cloudProviders as an array of CloudProvider
func (cps *CloudProviderService) ListCloudProviders() (cloudProviders []*types.CloudProvider, err error) {
	log.Debug("ListCloudProviders")

	data, status, err := cps.concertoService.Get(APIPathCloudProviders)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudProviders); err != nil {
		return nil, err
	}

	return cloudProviders, nil
}

// ListServerStoragePlans returns the list of storage plans as an array of StoragePlan
func (cps *CloudProviderService) ListServerStoragePlans(
	providerID string,
) (storagePlans []*types.StoragePlan, err error) {
	log.Debug("ListServerStoragePlans")

	data, status, err := cps.concertoService.Get(fmt.Sprintf(APIPathCloudProviderStoragePlans, providerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &storagePlans); err != nil {
		return nil, err
	}

	return storagePlans, nil
}

// ListLoadBalancerPlans returns the list of load balancer plans as an array of LoadBalancerPlan
func (cps *CloudProviderService) ListLoadBalancerPlans(
	providerID string,
) (loadBalancerPlans []*types.LoadBalancerPlan, err error) {
	log.Debug("ListLoadBalancerPlans")

	data, status, err := cps.concertoService.Get(
		fmt.Sprintf(APIPathCloudProviderLoadBalancerPlans, providerID),
	)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancerPlans); err != nil {
		return nil, err
	}

	return loadBalancerPlans, nil
}

// ListClusterPlans returns the list of cluster plans as an array of ClusterPlan
func (cps *CloudProviderService) ListClusterPlans(providerID string) (clusterPlans []*types.ClusterPlan, err error) {
	log.Debug("ListClusterPlans")

	data, status, err := cps.concertoService.Get(fmt.Sprintf(APIPathCloudProviderClusterPlans, providerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &clusterPlans); err != nil {
		return nil, err
	}

	return clusterPlans, nil
}
