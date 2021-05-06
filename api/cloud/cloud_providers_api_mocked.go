package cloud

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ListCloudProvidersMocked test mocked function
func ListCloudProvidersMocked(t *testing.T, cloudProvidersIn []*types.CloudProvider) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/cloud/cloud_providers").Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.ListCloudProviders()
	assert.Nil(err, "Error getting cloudProvider list")
	assert.Equal(cloudProvidersIn, cloudProvidersOut, "ListCloudProviders returned different cloudProviders")

	return cloudProvidersOut
}

// ListCloudProvidersFailErrMocked test mocked function
func ListCloudProvidersFailErrMocked(t *testing.T, cloudProvidersIn []*types.CloudProvider) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/cloud/cloud_providers").Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudProvidersOut, err := ds.ListCloudProviders()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudProvidersOut
}

// ListCloudProvidersFailStatusMocked test mocked function
func ListCloudProvidersFailStatusMocked(t *testing.T, cloudProvidersIn []*types.CloudProvider) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/cloud/cloud_providers").Return(dIn, 499, nil)
	cloudProvidersOut, err := ds.ListCloudProviders()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudProvidersOut
}

// ListCloudProvidersFailJSONMocked test mocked function
func ListCloudProvidersFailJSONMocked(t *testing.T, cloudProvidersIn []*types.CloudProvider) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/cloud/cloud_providers").Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.ListCloudProviders()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudProvidersOut
}

// ListServerStoragePlansMocked test mocked function
func ListServerStoragePlansMocked(t *testing.T, storagePlansIn []*types.StoragePlan, providerID string) []*types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlansIn)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID)).Return(dIn, 200, nil)
	storagePlansOut, err := ds.ListServerStoragePlans(providerID)
	assert.Nil(err, "Error getting storage plan list")
	assert.Equal(storagePlansIn, storagePlansOut, "ListServerStoragePlans returned different storage plans")

	return storagePlansOut
}

// ListServerStoragePlansFailErrMocked test mocked function
func ListServerStoragePlansFailErrMocked(t *testing.T, storagePlansIn []*types.StoragePlan, providerID string) []*types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlansIn)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	storagePlansOut, err := ds.ListServerStoragePlans(providerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(storagePlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return storagePlansOut
}

// ListServerStoragePlansFailStatusMocked test mocked function
func ListServerStoragePlansFailStatusMocked(t *testing.T, storagePlansIn []*types.StoragePlan, providerID string) []*types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(storagePlansIn)
	assert.Nil(err, "Storage plan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID)).Return(dIn, 499, nil)
	storagePlansOut, err := ds.ListServerStoragePlans(providerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(storagePlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return storagePlansOut
}

// ListServerStoragePlansFailJSONMocked test mocked function
func ListServerStoragePlansFailJSONMocked(t *testing.T, storagePlansIn []*types.StoragePlan, providerID string) []*types.StoragePlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/storage_plans", providerID)).Return(dIn, 200, nil)
	storagePlansOut, err := ds.ListServerStoragePlans(providerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(storagePlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return storagePlansOut
}

// ListLoadBalancerPlansMocked test mocked function
func ListLoadBalancerPlansMocked(t *testing.T, loadBalancerPlansIn []*types.LoadBalancerPlan, providerID string) []*types.LoadBalancerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerPlansIn)
	assert.Nil(err, "LoadBalancerPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/load_balancer_plans", providerID)).Return(dIn, 200, nil)
	loadBalancerPlansOut, err := ds.ListLoadBalancerPlans(providerID)

	assert.Nil(err, "Error getting load balancer plan list")
	assert.Equal(loadBalancerPlansIn, loadBalancerPlansOut, "ListLoadBalancerPlans returned different load balancer plans")

	return loadBalancerPlansOut
}

// ListLoadBalancerPlansFailErrMocked test mocked function
func ListLoadBalancerPlansFailErrMocked(t *testing.T, loadBalancerPlansIn []*types.LoadBalancerPlan, providerID string) []*types.LoadBalancerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerPlansIn)
	assert.Nil(err, "LoadBalancerPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/load_balancer_plans", providerID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	loadBalancerPlansOut, err := ds.ListLoadBalancerPlans(providerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return loadBalancerPlansOut
}

// ListLoadBalancerPlansFailStatusMocked test mocked function
func ListLoadBalancerPlansFailStatusMocked(t *testing.T, loadBalancerPlansIn []*types.LoadBalancerPlan, providerID string) []*types.LoadBalancerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerPlansIn)
	assert.Nil(err, "LoadBalancerPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/load_balancer_plans", providerID)).Return(dIn, 499, nil)
	loadBalancerPlansOut, err := ds.ListLoadBalancerPlans(providerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancerPlansOut
}

// ListLoadBalancerPlansFailJSONMocked test mocked function
func ListLoadBalancerPlansFailJSONMocked(t *testing.T, loadBalancerPlansIn []*types.LoadBalancerPlan, providerID string) []*types.LoadBalancerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/load_balancer_plans", providerID)).Return(dIn, 200, nil)
	loadBalancerPlansOut, err := ds.ListLoadBalancerPlans(providerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerPlansOut
}

// ListClusterPlansMocked test mocked function
func ListClusterPlansMocked(t *testing.T, clusterPlansIn []*types.ClusterPlan, providerID string) []*types.ClusterPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(clusterPlansIn)
	assert.Nil(err, "ClusterPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/cluster_plans", providerID)).Return(dIn, 200, nil)
	clusterPlansOut, err := ds.ListClusterPlans(providerID)

	assert.Nil(err, "Error getting cluster plan list")
	assert.Equal(clusterPlansIn, clusterPlansOut, "ListClusterPlans returned different cluster plans")

	return clusterPlansOut
}

// ListClusterPlansFailErrMocked test mocked function
func ListClusterPlansFailErrMocked(t *testing.T, clusterPlansIn []*types.ClusterPlan, providerID string) []*types.ClusterPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(clusterPlansIn)
	assert.Nil(err, "ClusterPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/cluster_plans", providerID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	clusterPlansOut, err := ds.ListClusterPlans(providerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clusterPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return clusterPlansOut
}

// ListClusterPlansFailStatusMocked test mocked function
func ListClusterPlansFailStatusMocked(t *testing.T, clusterPlansIn []*types.ClusterPlan, providerID string) []*types.ClusterPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(clusterPlansIn)
	assert.Nil(err, "ClusterPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/cluster_plans", providerID)).Return(dIn, 499, nil)
	clusterPlansOut, err := ds.ListClusterPlans(providerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clusterPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return clusterPlansOut
}

// ListClusterPlansFailJSONMocked test mocked function
func ListClusterPlansFailJSONMocked(t *testing.T, clusterPlansIn []*types.ClusterPlan, providerID string) []*types.ClusterPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/cluster_plans", providerID)).Return(dIn, 200, nil)
	clusterPlansOut, err := ds.ListClusterPlans(providerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clusterPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return clusterPlansOut
}

// ListNodePoolPlansMocked test mocked function
func ListNodePoolPlansMocked(t *testing.T, nodePoolPlansIn []*types.NodePoolPlan, providerID string) []*types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlansIn)
	assert.Nil(err, "NodePoolPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/node_pool_plans", providerID)).Return(dIn, 200, nil)
	nodePoolPlansOut, err := ds.ListNodePoolPlans(providerID)

	assert.Nil(err, "Error getting node pool plan list")
	assert.Equal(nodePoolPlansIn, nodePoolPlansOut, "ListNodePoolPlans returned different node pool plans")

	return nodePoolPlansOut
}

// ListNodePoolPlansFailErrMocked test mocked function
func ListNodePoolPlansFailErrMocked(t *testing.T, nodePoolPlansIn []*types.NodePoolPlan, providerID string) []*types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlansIn)
	assert.Nil(err, "NodePoolPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/node_pool_plans", providerID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	nodePoolPlansOut, err := ds.ListNodePoolPlans(providerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return nodePoolPlansOut
}

// ListNodePoolPlansFailStatusMocked test mocked function
func ListNodePoolPlansFailStatusMocked(t *testing.T, nodePoolPlansIn []*types.NodePoolPlan, providerID string) []*types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlansIn)
	assert.Nil(err, "NodePoolPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/node_pool_plans", providerID)).Return(dIn, 499, nil)
	nodePoolPlansOut, err := ds.ListNodePoolPlans(providerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodePoolPlansOut
}

// ListNodePoolPlansFailJSONMocked test mocked function
func ListNodePoolPlansFailJSONMocked(t *testing.T, nodePoolPlansIn []*types.NodePoolPlan, providerID string) []*types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/node_pool_plans", providerID)).Return(dIn, 200, nil)
	nodePoolPlansOut, err := ds.ListNodePoolPlans(providerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return nodePoolPlansOut
}
