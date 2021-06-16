// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// ListLoadBalancersMocked test mocked function
func ListLoadBalancersMocked(t *testing.T, loadBalancersIn []*types.LoadBalancer) []*types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancersIn)
	assert.Nil(err, "LoadBalancers test data corrupted")

	// call service
	cs.On("Get", APIPathNetworkLoadBalancers).Return(dIn, 200, nil)
	loadBalancersOut, err := ds.ListLoadBalancers()

	assert.Nil(err, "Error getting load balancers")
	assert.Equal(loadBalancersIn, loadBalancersOut, "ListLoadBalancers returned different load balancers")

	return loadBalancersOut
}

// ListLoadBalancersFailErrMocked test mocked function
func ListLoadBalancersFailErrMocked(t *testing.T, loadBalancersIn []*types.LoadBalancer) []*types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancersIn)
	assert.Nil(err, "LoadBalancers test data corrupted")

	// call service
	cs.On("Get", APIPathNetworkLoadBalancers).Return(dIn, 200, fmt.Errorf("mocked error"))
	loadBalancersOut, err := ds.ListLoadBalancers()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancersOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return loadBalancersOut
}

// ListLoadBalancersFailStatusMocked test mocked function
func ListLoadBalancersFailStatusMocked(t *testing.T, loadBalancersIn []*types.LoadBalancer) []*types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancersIn)
	assert.Nil(err, "LoadBalancers test data corrupted")

	// call service
	cs.On("Get", APIPathNetworkLoadBalancers).Return(dIn, 499, nil)
	loadBalancersOut, err := ds.ListLoadBalancers()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancersOut
}

// ListLoadBalancersFailJSONMocked test mocked function
func ListLoadBalancersFailJSONMocked(t *testing.T, loadBalancersIn []*types.LoadBalancer) []*types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathNetworkLoadBalancers).Return(dIn, 200, nil)
	loadBalancersOut, err := ds.ListLoadBalancers()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancersOut
}

// GetLoadBalancerMocked test mocked function
func GetLoadBalancerMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID)).Return(dIn, 200, nil)
	loadBalancerOut, err := ds.GetLoadBalancer(loadBalancerIn.ID)

	assert.Nil(err, "Error getting load balancer")
	assert.Equal(*loadBalancerIn, *loadBalancerOut, "GetLoadBalancer returned different load balancer")

	return loadBalancerOut
}

// GetLoadBalancerFailErrMocked test mocked function
func GetLoadBalancerFailErrMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	loadBalancerOut, err := ds.GetLoadBalancer(loadBalancerIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return loadBalancerOut
}

// GetLoadBalancerFailStatusMocked test mocked function
func GetLoadBalancerFailStatusMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID)).Return(dIn, 499, nil)
	loadBalancerOut, err := ds.GetLoadBalancer(loadBalancerIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancerOut
}

// GetLoadBalancerFailJSONMocked test mocked function
func GetLoadBalancerFailJSONMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID)).Return(dIn, 200, nil)
	loadBalancerOut, err := ds.GetLoadBalancer(loadBalancerIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerOut
}

// CreateLoadBalancerMocked test mocked function
func CreateLoadBalancerMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Post", APIPathNetworkLoadBalancers, mapIn).Return(dOut, 200, nil)
	loadBalancerOut, err := ds.CreateLoadBalancer(mapIn)

	assert.Nil(err, "Error creating load balancer")
	assert.Equal(loadBalancerIn, loadBalancerOut, "CreateLoadBalancer returned different load balancer")

	return loadBalancerOut
}

// CreateLoadBalancerFailErrMocked test mocked function
func CreateLoadBalancerFailErrMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Post", APIPathNetworkLoadBalancers, mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	loadBalancerOut, err := ds.CreateLoadBalancer(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return loadBalancerOut
}

// CreateLoadBalancerFailStatusMocked test mocked function
func CreateLoadBalancerFailStatusMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Post", APIPathNetworkLoadBalancers, mapIn).Return(dOut, 499, nil)
	loadBalancerOut, err := ds.CreateLoadBalancer(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancerOut
}

// CreateLoadBalancerFailJSONMocked test mocked function
func CreateLoadBalancerFailJSONMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", APIPathNetworkLoadBalancers, mapIn).Return(dIn, 200, nil)
	loadBalancerOut, err := ds.CreateLoadBalancer(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerOut
}

// UpdateLoadBalancerMocked test mocked function
func UpdateLoadBalancerMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID), mapIn).Return(dOut, 200, nil)
	loadBalancerOut, err := ds.UpdateLoadBalancer(loadBalancerIn.ID, mapIn)

	assert.Nil(err, "Error updating load balancer")
	assert.Equal(loadBalancerIn, loadBalancerOut, "UpdateLoadBalancer returned different load balancer")

	return loadBalancerOut
}

// UpdateLoadBalancerFailErrMocked test mocked function
func UpdateLoadBalancerFailErrMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	loadBalancerOut, err := ds.UpdateLoadBalancer(loadBalancerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return loadBalancerOut
}

// UpdateLoadBalancerFailStatusMocked test mocked function
func UpdateLoadBalancerFailStatusMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID), mapIn).Return(dOut, 499, nil)
	loadBalancerOut, err := ds.UpdateLoadBalancer(loadBalancerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancerOut
}

// UpdateLoadBalancerFailJSONMocked test mocked function
func UpdateLoadBalancerFailJSONMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID), mapIn).Return(dIn, 200, nil)
	loadBalancerOut, err := ds.UpdateLoadBalancer(loadBalancerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerOut
}

// DeleteLoadBalancerMocked test mocked function
func DeleteLoadBalancerMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID)).Return(dIn, 200, nil)
	loadBalancerOut, err := ds.DeleteLoadBalancer(loadBalancerIn.ID)

	assert.Nil(err, "Error deleting load balancer")
	assert.Equal(loadBalancerIn, loadBalancerOut, "DeleteLoadBalancer returned different load balancer")

}

// DeleteLoadBalancerFailErrMocked test mocked function
func DeleteLoadBalancerFailErrMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	loadBalancerOut, err := ds.DeleteLoadBalancer(loadBalancerIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteLoadBalancerFailStatusMocked test mocked function
func DeleteLoadBalancerFailStatusMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID)).Return(dIn, 499, nil)
	loadBalancerOut, err := ds.DeleteLoadBalancer(loadBalancerIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteLoadBalancerFailJSONMocked test mocked function
func DeleteLoadBalancerFailJSONMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkLoadBalancer, loadBalancerIn.ID)).Return(dIn, 200, nil)
	loadBalancerOut, err := ds.DeleteLoadBalancer(loadBalancerIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")
}

// RetryLoadBalancerMocked test mocked function
func RetryLoadBalancerMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancerRetry, loadBalancerIn.ID), mapIn).Return(dOut, 200, nil)
	loadBalancerOut, err := ds.RetryLoadBalancer(loadBalancerIn.ID, mapIn)

	assert.Nil(err, "Error retrying load balancer")
	assert.Equal(loadBalancerIn, loadBalancerOut, "RetryLoadBalancer returned different load balancer")

	return loadBalancerOut
}

// RetryLoadBalancerFailErrMocked test mocked function
func RetryLoadBalancerFailErrMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancerRetry, loadBalancerIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	loadBalancerOut, err := ds.RetryLoadBalancer(loadBalancerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return loadBalancerOut
}

// RetryLoadBalancerFailStatusMocked test mocked function
func RetryLoadBalancerFailStatusMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancerRetry, loadBalancerIn.ID), mapIn).Return(dOut, 499, nil)
	loadBalancerOut, err := ds.RetryLoadBalancer(loadBalancerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancerOut
}

// RetryLoadBalancerFailJSONMocked test mocked function
func RetryLoadBalancerFailJSONMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancerRetry, loadBalancerIn.ID), mapIn).Return(dIn, 200, nil)
	loadBalancerOut, err := ds.RetryLoadBalancer(loadBalancerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerOut
}

// GetLoadBalancerPlanMocked test mocked function
func GetLoadBalancerPlanMocked(
	t *testing.T,
	loadBalancerPlanID string,
	loadBalancerPlanIn *types.LoadBalancerPlan,
) *types.LoadBalancerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerPlanIn)
	assert.Nil(err, "LoadBalancerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerPlan, loadBalancerPlanID)).Return(dIn, 200, nil)
	loadBalancerPlanOut, err := ds.GetLoadBalancerPlan(loadBalancerPlanID)

	assert.Nil(err, "Error getting load balancer plan")
	assert.Equal(loadBalancerPlanIn, loadBalancerPlanOut, "GetLoadBalancerPlan returned different load balancer plan")

	return loadBalancerPlanOut
}

// GetLoadBalancerPlanFailErrMocked test mocked function
func GetLoadBalancerPlanFailErrMocked(
	t *testing.T,
	loadBalancerPlanID string,
	loadBalancerPlanIn *types.LoadBalancerPlan,
) *types.LoadBalancerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerPlanIn)
	assert.Nil(err, "LoadBalancerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerPlan, loadBalancerPlanID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	loadBalancerPlanOut, err := ds.GetLoadBalancerPlan(loadBalancerPlanID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerPlanOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return loadBalancerPlanOut
}

// GetLoadBalancerPlanFailStatusMocked test mocked function
func GetLoadBalancerPlanFailStatusMocked(
	t *testing.T,
	loadBalancerPlanID string,
	loadBalancerPlanIn *types.LoadBalancerPlan,
) *types.LoadBalancerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// to json
	dIn, err := json.Marshal(loadBalancerPlanIn)
	assert.Nil(err, "LoadBalancerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerPlan, loadBalancerPlanID)).Return(dIn, 499, nil)
	loadBalancerPlanOut, err := ds.GetLoadBalancerPlan(loadBalancerPlanID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancerPlanOut
}

// GetLoadBalancerPlanFailJSONMocked test mocked function
func GetLoadBalancerPlanFailJSONMocked(
	t *testing.T,
	loadBalancerPlanID string,
	loadBalancerPlanIn *types.LoadBalancerPlan,
) *types.LoadBalancerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(ds, "LoadBalancer service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerPlan, loadBalancerPlanID)).Return(dIn, 200, nil)
	loadBalancerPlanOut, err := ds.GetLoadBalancerPlan(loadBalancerPlanID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerPlanOut
}
