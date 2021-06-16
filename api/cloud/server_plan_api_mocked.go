// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// ListServerPlansMocked test mocked function
func ListServerPlansMocked(
	t *testing.T,
	serverPlansIn []*types.ServerPlan,
	cloudProviderId string,
) []*types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlansIn)
	assert.Nil(err, "ServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudProviderServerPlans, cloudProviderId)).Return(dIn, 200, nil)
	serverPlansOut, err := ds.ListServerPlans(cloudProviderId)
	assert.Nil(err, "Error getting serverPlan list")
	assert.Equal(serverPlansIn, serverPlansOut, "ListServerPlans returned different serverPlans")

	return serverPlansOut
}

// ListServerPlansFailErrMocked test mocked function
func ListServerPlansFailErrMocked(
	t *testing.T,
	serverPlansIn []*types.ServerPlan,
	cloudProviderId string,
) []*types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlansIn)
	assert.Nil(err, "ServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudProviderServerPlans, cloudProviderId)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	serverPlansOut, err := ds.ListServerPlans(cloudProviderId)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverPlansOut
}

// ListServerPlansFailStatusMocked test mocked function
func ListServerPlansFailStatusMocked(
	t *testing.T,
	serverPlansIn []*types.ServerPlan,
	cloudProviderId string,
) []*types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlansIn)
	assert.Nil(err, "ServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudProviderServerPlans, cloudProviderId)).Return(dIn, 499, nil)
	serverPlansOut, err := ds.ListServerPlans(cloudProviderId)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverPlansOut
}

// ListServerPlansFailJSONMocked test mocked function
func ListServerPlansFailJSONMocked(
	t *testing.T,
	serverPlansIn []*types.ServerPlan,
	cloudProviderId string,
) []*types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudProviderServerPlans, cloudProviderId)).Return(dIn, 200, nil)
	serverPlansOut, err := ds.ListServerPlans(cloudProviderId)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverPlansOut
}

// GetServerPlanMocked test mocked function
func GetServerPlanMocked(t *testing.T, serverPlan *types.ServerPlan) *types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlan)
	assert.Nil(err, "ServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudServerPlans, serverPlan.ID)).Return(dIn, 200, nil)
	serverPlanOut, err := ds.GetServerPlan(serverPlan.ID)
	assert.Nil(err, "Error getting serverPlan")
	assert.Equal(*serverPlan, *serverPlanOut, "GetServerPlan returned different serverPlans")

	return serverPlanOut
}

// GetServerPlanFailErrMocked test mocked function
func GetServerPlanFailErrMocked(t *testing.T, serverPlan *types.ServerPlan) *types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlan)
	assert.Nil(err, "ServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudServerPlans, serverPlan.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	serverPlanOut, err := ds.GetServerPlan(serverPlan.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverPlanOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverPlanOut
}

// GetServerPlanFailStatusMocked test mocked function
func GetServerPlanFailStatusMocked(t *testing.T, serverPlan *types.ServerPlan) *types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlan)
	assert.Nil(err, "ServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudServerPlans, serverPlan.ID)).Return(dIn, 499, nil)
	serverPlanOut, err := ds.GetServerPlan(serverPlan.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverPlanOut
}

// GetServerPlanFailJSONMocked test mocked function
func GetServerPlanFailJSONMocked(t *testing.T, serverPlan *types.ServerPlan) *types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudServerPlans, serverPlan.ID)).Return(dIn, 200, nil)
	serverPlanOut, err := ds.GetServerPlan(serverPlan.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverPlanOut
}
