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

// ListListenersMocked test mocked function
func ListListenersMocked(t *testing.T, loadBalancerID string, listenersIn []*types.Listener) []*types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenersIn)
	assert.Nil(err, "Listeners test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerListeners, loadBalancerID)).Return(dIn, 200, nil)
	listenersOut, err := ds.ListListeners(loadBalancerID)

	assert.Nil(err, "Error getting listeners")
	assert.Equal(listenersIn, listenersOut, "ListListeners returned different listeners")

	return listenersOut
}

// ListListenersFailErrMocked test mocked function
func ListListenersFailErrMocked(t *testing.T, loadBalancerID string, listenersIn []*types.Listener) []*types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenersIn)
	assert.Nil(err, "Listeners test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerListeners, loadBalancerID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	listenersOut, err := ds.ListListeners(loadBalancerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(listenersOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return listenersOut
}

// ListListenersFailStatusMocked test mocked function
func ListListenersFailStatusMocked(
	t *testing.T,
	loadBalancerID string,
	listenersIn []*types.Listener,
) []*types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenersIn)
	assert.Nil(err, "Listeners test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerListeners, loadBalancerID)).Return(dIn, 499, nil)
	listenersOut, err := ds.ListListeners(loadBalancerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(listenersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return listenersOut
}

// ListListenersFailJSONMocked test mocked function
func ListListenersFailJSONMocked(t *testing.T, loadBalancerID string, listenersIn []*types.Listener) []*types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerListeners, loadBalancerID)).Return(dIn, 200, nil)
	listenersOut, err := ds.ListListeners(loadBalancerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(listenersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return listenersOut
}

// GetListenerMocked test mocked function
func GetListenerMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID)).Return(dIn, 200, nil)
	listenerOut, err := ds.GetListener(listenerIn.ID)

	assert.Nil(err, "Error getting listener")
	assert.Equal(*listenerIn, *listenerOut, "GetListener returned different listener")

	return listenerOut
}

// GetListenerFailErrMocked test mocked function
func GetListenerFailErrMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	listenerOut, err := ds.GetListener(listenerIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return listenerOut
}

// GetListenerFailStatusMocked test mocked function
func GetListenerFailStatusMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID)).Return(dIn, 499, nil)
	listenerOut, err := ds.GetListener(listenerIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return listenerOut
}

// GetListenerFailJSONMocked test mocked function
func GetListenerFailJSONMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID)).Return(dIn, 200, nil)
	listenerOut, err := ds.GetListener(listenerIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return listenerOut
}

// CreateListenerMocked test mocked function
func CreateListenerMocked(t *testing.T, loadBalancerID string, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkLoadBalancerListeners, loadBalancerID), mapIn).Return(dOut, 200, nil)
	listenerOut, err := ds.CreateListener(loadBalancerID, mapIn)

	assert.Nil(err, "Error creating listener")
	assert.Equal(listenerIn, listenerOut, "CreateListener returned different listener")

	return listenerOut
}

// CreateListenerFailErrMocked test mocked function
func CreateListenerFailErrMocked(t *testing.T, loadBalancerID string, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkLoadBalancerListeners, loadBalancerID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	listenerOut, err := ds.CreateListener(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return listenerOut
}

// CreateListenerFailStatusMocked test mocked function
func CreateListenerFailStatusMocked(t *testing.T, loadBalancerID string, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkLoadBalancerListeners, loadBalancerID), mapIn).Return(dOut, 499, nil)
	listenerOut, err := ds.CreateListener(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return listenerOut
}

// CreateListenerFailJSONMocked test mocked function
func CreateListenerFailJSONMocked(t *testing.T, loadBalancerID string, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkLoadBalancerListeners, loadBalancerID), mapIn).Return(dIn, 200, nil)
	listenerOut, err := ds.CreateListener(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return listenerOut
}

// UpdateListenerMocked test mocked function
func UpdateListenerMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID), mapIn).Return(dOut, 200, nil)
	listenerOut, err := ds.UpdateListener(listenerIn.ID, mapIn)

	assert.Nil(err, "Error updating listener")
	assert.Equal(listenerIn, listenerOut, "UpdateListener returned different listener")

	return listenerOut
}

// UpdateListenerFailErrMocked test mocked function
func UpdateListenerFailErrMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	listenerOut, err := ds.UpdateListener(listenerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return listenerOut
}

// UpdateListenerFailStatusMocked test mocked function
func UpdateListenerFailStatusMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID), mapIn).Return(dOut, 499, nil)
	listenerOut, err := ds.UpdateListener(listenerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return listenerOut
}

// UpdateListenerFailJSONMocked test mocked function
func UpdateListenerFailJSONMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID), mapIn).Return(dIn, 200, nil)
	listenerOut, err := ds.UpdateListener(listenerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return listenerOut
}

// DeleteListenerMocked test mocked function
func DeleteListenerMocked(t *testing.T, listenerIn *types.Listener) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID)).Return(dIn, 200, nil)
	listenerOut, err := ds.DeleteListener(listenerIn.ID)

	assert.Nil(err, "Error deleting listener")
	assert.Equal(listenerIn, listenerOut, "DeleteListener returned different listener")

}

// DeleteListenerFailErrMocked test mocked function
func DeleteListenerFailErrMocked(t *testing.T, listenerIn *types.Listener) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	listenerOut, err := ds.DeleteListener(listenerIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteListenerFailStatusMocked test mocked function
func DeleteListenerFailStatusMocked(t *testing.T, listenerIn *types.Listener) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID)).Return(dIn, 499, nil)
	listenerOut, err := ds.DeleteListener(listenerIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteListenerFailJSONMocked test mocked function
func DeleteListenerFailJSONMocked(t *testing.T, listenerIn *types.Listener) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkListener, listenerIn.ID)).Return(dIn, 200, nil)
	listenerOut, err := ds.DeleteListener(listenerIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")
}

// RetryListenerMocked test mocked function
func RetryListenerMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListenerRetry, listenerIn.ID), mapIn).Return(dOut, 200, nil)
	listenerOut, err := ds.RetryListener(listenerIn.ID, mapIn)

	assert.Nil(err, "Error retrying listener")
	assert.Equal(listenerIn, listenerOut, "RetryListener returned different listener")

	return listenerOut
}

// RetryListenerFailErrMocked test mocked function
func RetryListenerFailErrMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListenerRetry, listenerIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	listenerOut, err := ds.RetryListener(listenerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return listenerOut
}

// RetryListenerFailStatusMocked test mocked function
func RetryListenerFailStatusMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// to json
	dOut, err := json.Marshal(listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListenerRetry, listenerIn.ID), mapIn).Return(dOut, 499, nil)
	listenerOut, err := ds.RetryListener(listenerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return listenerOut
}

// RetryListenerFailJSONMocked test mocked function
func RetryListenerFailJSONMocked(t *testing.T, listenerIn *types.Listener) *types.Listener {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*listenerIn)
	assert.Nil(err, "Listener test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListenerRetry, listenerIn.ID), mapIn).Return(dIn, 200, nil)
	listenerOut, err := ds.RetryListener(listenerIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(listenerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return listenerOut
}

// ListRulesMocked test mocked function
func ListRulesMocked(t *testing.T, listenerID string, rulesIn []*types.ListenerRule) []*types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(rulesIn)
	assert.Nil(err, "Rules test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkListenerRules, listenerID)).Return(dIn, 200, nil)
	rulesOut, err := ds.ListRules(listenerID)

	assert.Nil(err, "Error getting rules")
	assert.Equal(rulesIn, rulesOut, "ListRules returned different rules")

	return rulesOut
}

// ListRulesFailErrMocked test mocked function
func ListRulesFailErrMocked(t *testing.T, listenerID string, rulesIn []*types.ListenerRule) []*types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(rulesIn)
	assert.Nil(err, "Rules test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkListenerRules, listenerID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	rulesOut, err := ds.ListRules(listenerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(rulesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return rulesOut
}

// ListRulesFailStatusMocked test mocked function
func ListRulesFailStatusMocked(t *testing.T, listenerID string, rulesIn []*types.ListenerRule) []*types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(rulesIn)
	assert.Nil(err, "Rules test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkListenerRules, listenerID)).Return(dIn, 499, nil)
	rulesOut, err := ds.ListRules(listenerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(rulesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return rulesOut
}

// ListRulesFailJSONMocked test mocked function
func ListRulesFailJSONMocked(t *testing.T, listenerID string, rulesIn []*types.ListenerRule) []*types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkListenerRules, listenerID)).Return(dIn, 200, nil)
	rulesOut, err := ds.ListRules(listenerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(rulesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return rulesOut
}

// CreateRuleMocked test mocked function
func CreateRuleMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) *types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// to json
	dOut, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkListenerRules, listenerID), mapIn).Return(dOut, 200, nil)
	ruleOut, err := ds.CreateRule(listenerID, mapIn)

	assert.Nil(err, "Error creating rule")
	assert.Equal(ruleIn, ruleOut, "CreateRule returned different rule")

	return ruleOut
}

// CreateRuleFailErrMocked test mocked function
func CreateRuleFailErrMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) *types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// to json
	dOut, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkListenerRules, listenerID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	ruleOut, err := ds.CreateRule(listenerID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(ruleOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return ruleOut
}

// CreateRuleFailStatusMocked test mocked function
func CreateRuleFailStatusMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) *types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// to json
	dOut, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkListenerRules, listenerID), mapIn).Return(dOut, 499, nil)
	ruleOut, err := ds.CreateRule(listenerID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(ruleOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return ruleOut
}

// CreateRuleFailJSONMocked test mocked function
func CreateRuleFailJSONMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) *types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkListenerRules, listenerID), mapIn).Return(dIn, 200, nil)
	ruleOut, err := ds.CreateRule(listenerID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(ruleOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return ruleOut
}

// UpdateRuleMocked test mocked function
func UpdateRuleMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) *types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// to json
	dOut, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListenerRule, listenerID, ruleIn.ID), mapIn).Return(dOut, 200, nil)
	ruleOut, err := ds.UpdateRule(listenerID, ruleIn.ID, mapIn)

	assert.Nil(err, "Error updating rule")
	assert.Equal(ruleIn, ruleOut, "UpdateRule returned different rule")

	return ruleOut
}

// UpdateRuleFailErrMocked test mocked function
func UpdateRuleFailErrMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) *types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// to json
	dOut, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListenerRule, listenerID, ruleIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	ruleOut, err := ds.UpdateRule(listenerID, ruleIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(ruleOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return ruleOut
}

// UpdateRuleFailStatusMocked test mocked function
func UpdateRuleFailStatusMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) *types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// to json
	dOut, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListenerRule, listenerID, ruleIn.ID), mapIn).Return(dOut, 499, nil)
	ruleOut, err := ds.UpdateRule(listenerID, ruleIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(ruleOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return ruleOut
}

// UpdateRuleFailJSONMocked test mocked function
func UpdateRuleFailJSONMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) *types.ListenerRule {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkListenerRule, listenerID, ruleIn.ID), mapIn).Return(dIn, 200, nil)
	ruleOut, err := ds.UpdateRule(listenerID, ruleIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(ruleOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return ruleOut
}

// DeleteRuleMocked test mocked function
func DeleteRuleMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkListenerRule, listenerID, ruleIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteRule(listenerID, ruleIn.ID)

	assert.Nil(err, "Error deleting rule")
}

// DeleteRuleFailErrMocked test mocked function
func DeleteRuleFailErrMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkListenerRule, listenerID, ruleIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteRule(listenerID, ruleIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteRuleFailStatusMocked test mocked function
func DeleteRuleFailStatusMocked(t *testing.T, listenerID string, ruleIn *types.ListenerRule) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewListenerService(cs)
	assert.Nil(err, "Couldn't load listener service")
	assert.NotNil(ds, "Listener service not instanced")

	// to json
	dIn, err := json.Marshal(ruleIn)
	assert.Nil(err, "Rule test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkListenerRule, listenerID, ruleIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteRule(listenerID, ruleIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
