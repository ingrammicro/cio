package cloud

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ListRealmsMocked test mocked function
func ListRealmsMocked(t *testing.T, realmsIn []*types.Realm, cloudProviderID string) []*types.Realm {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(realmsIn)
	assert.Nil(err, "Realms test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/realms", cloudProviderID)).Return(dIn, 200, nil)
	realmsOut, err := ds.ListRealms(cloudProviderID)

	assert.Nil(err, "Error getting realms")
	assert.Equal(realmsIn, realmsOut, "ListRealms returned different realms")

	return realmsOut
}

// ListRealmsFailErrMocked test mocked function
func ListRealmsFailErrMocked(t *testing.T, realmsIn []*types.Realm, cloudProviderID string) []*types.Realm {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(realmsIn)
	assert.Nil(err, "Realms test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/realms", cloudProviderID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	realmsOut, err := ds.ListRealms(cloudProviderID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(realmsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return realmsOut
}

// ListRealmsFailStatusMocked test mocked function
func ListRealmsFailStatusMocked(t *testing.T, realmsIn []*types.Realm, cloudProviderID string) []*types.Realm {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(realmsIn)
	assert.Nil(err, "Realms test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/realms", cloudProviderID)).Return(dIn, 499, nil)
	realmsOut, err := ds.ListRealms(cloudProviderID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(realmsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return realmsOut
}

// ListRealmsFailJSONMocked test mocked function
func ListRealmsFailJSONMocked(t *testing.T, realmsIn []*types.Realm, cloudProviderID string) []*types.Realm {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/cloud_providers/%s/realms", cloudProviderID)).Return(dIn, 200, nil)
	realmsOut, err := ds.ListRealms(cloudProviderID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(realmsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return realmsOut
}

// GetRealmMocked test mocked function
func GetRealmMocked(t *testing.T, realmIn *types.Realm) *types.Realm {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(realmIn)
	assert.Nil(err, "Realm test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/realms/%s", realmIn.ID)).Return(dIn, 200, nil)
	realmOut, err := ds.GetRealm(realmIn.ID)

	assert.Nil(err, "Error getting realm")
	assert.Equal(*realmIn, *realmOut, "GetRealm returned different realm")

	return realmOut
}

// GetRealmFailErrMocked test mocked function
func GetRealmFailErrMocked(t *testing.T, realmIn *types.Realm) *types.Realm {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(realmIn)
	assert.Nil(err, "Realm test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/realms/%s", realmIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	realmOut, err := ds.GetRealm(realmIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(realmOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return realmOut
}

// GetRealmFailStatusMocked test mocked function
func GetRealmFailStatusMocked(t *testing.T, realmIn *types.Realm) *types.Realm {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(realmIn)
	assert.Nil(err, "Realm test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/realms/%s", realmIn.ID)).Return(dIn, 499, nil)
	realmOut, err := ds.GetRealm(realmIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(realmOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return realmOut
}

// GetRealmFailJSONMocked test mocked function
func GetRealmFailJSONMocked(t *testing.T, realmIn *types.Realm) *types.Realm {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/realms/%s", realmIn.ID)).Return(dIn, 200, nil)
	realmOut, err := ds.GetRealm(realmIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(realmOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return realmOut
}

// ListRealmNodePoolPlansMocked test mocked function
func ListRealmNodePoolPlansMocked(
	t *testing.T, nodePoolPlansIn []*types.NodePoolPlan, cloudProviderID string,
) []*types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlansIn)
	assert.Nil(err, "NodePoolPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/realms/%s/node_pool_plans", cloudProviderID)).Return(dIn, 200, nil)
	nodePoolPlansOut, err := ds.ListRealmNodePoolPlans(cloudProviderID)

	assert.Nil(err, "Error getting node pool plan list")
	assert.Equal(nodePoolPlansIn, nodePoolPlansOut, "ListRealmNodePoolPlans returned different node pool plans")

	return nodePoolPlansOut
}

// ListRealmNodePoolPlansFailErrMocked test mocked function
func ListRealmNodePoolPlansFailErrMocked(
	t *testing.T, nodePoolPlansIn []*types.NodePoolPlan, cloudProviderID string,
) []*types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlansIn)
	assert.Nil(err, "NodePoolPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/realms/%s/node_pool_plans", cloudProviderID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	nodePoolPlansOut, err := ds.ListRealmNodePoolPlans(cloudProviderID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodePoolPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return nodePoolPlansOut
}

// ListRealmNodePoolPlansFailStatusMocked test mocked function
func ListRealmNodePoolPlansFailStatusMocked(
	t *testing.T, nodePoolPlansIn []*types.NodePoolPlan, cloudProviderID string,
) []*types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// to json
	dIn, err := json.Marshal(nodePoolPlansIn)
	assert.Nil(err, "NodePoolPlans test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/realms/%s/node_pool_plans", cloudProviderID)).Return(dIn, 499, nil)
	nodePoolPlansOut, err := ds.ListRealmNodePoolPlans(cloudProviderID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodePoolPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodePoolPlansOut
}

// ListRealmNodePoolPlansFailJSONMocked test mocked function
func ListRealmNodePoolPlansFailJSONMocked(
	t *testing.T, nodePoolPlansIn []*types.NodePoolPlan, cloudProviderID string,
) []*types.NodePoolPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewRealmService(cs)
	assert.Nil(err, "Couldn't load realm service")
	assert.NotNil(ds, "Realm service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/cloud/realms/%s/node_pool_plans", cloudProviderID)).Return(dIn, 200, nil)
	nodePoolPlansOut, err := ds.ListRealmNodePoolPlans(cloudProviderID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodePoolPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return nodePoolPlansOut
}
