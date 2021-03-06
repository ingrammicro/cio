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

// TODO exclude from release compile

// ListFloatingIPsMocked test mocked function
func ListFloatingIPsMocked(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", APIPathNetworkFloatingIPs).Return(dIn, 200, nil)
	floatingIPOut, err := ds.ListFloatingIPs("")
	assert.Nil(err, "Error getting floating IP list")
	assert.Equal(floatingIPIn, floatingIPOut, "ListFloatingIPs returned different floating IPs")

	return floatingIPOut
}

func ListFloatingIPsMockedFilteredByServer(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCloudServerFloatingIPs, floatingIPIn[0].AttachedServerID)).Return(dIn, 200, nil)
	floatingIPOut, err := ds.ListFloatingIPs(floatingIPIn[0].AttachedServerID)
	assert.Nil(err, "Error getting floating IP list filtered by server")
	assert.Equal(floatingIPIn, floatingIPOut, "ListFloatingIPs returned different floating IPs")

	return floatingIPOut
}

// ListFloatingIPsFailErrMocked test mocked function
func ListFloatingIPsFailErrMocked(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", APIPathNetworkFloatingIPs).Return(dIn, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.ListFloatingIPs("")

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// ListFloatingIPsFailStatusMocked test mocked function
func ListFloatingIPsFailStatusMocked(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", APIPathNetworkFloatingIPs).Return(dIn, 499, nil)
	floatingIPOut, err := ds.ListFloatingIPs("")

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPOut
}

// ListFloatingIPsFailJSONMocked test mocked function
func ListFloatingIPsFailJSONMocked(t *testing.T, floatingIPIn []*types.FloatingIP) []*types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathNetworkFloatingIPs).Return(dIn, 200, nil)
	floatingIPOut, err := ds.ListFloatingIPs("")

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// GetFloatingIPMocked test mocked function
func GetFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID)).Return(dIn, 200, nil)
	floatingIPOut, err := ds.GetFloatingIP(floatingIPIn.ID)
	assert.Nil(err, "Error getting floating IP")
	assert.Equal(*floatingIPIn, *floatingIPOut, "GetFloatingIP returned different floating IPs")

	return floatingIPOut
}

// GetFloatingIPFailErrMocked test mocked function
func GetFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.GetFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// GetFloatingIPFailStatusMocked test mocked function
func GetFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID)).Return(dIn, 499, nil)
	floatingIPOut, err := ds.GetFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPOut
}

// GetFloatingIPFailJSONMocked test mocked function
func GetFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID)).Return(dIn, 200, nil)
	floatingIPOut, err := ds.GetFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// CreateFloatingIPMocked test mocked function
func CreateFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", APIPathNetworkFloatingIPs, mapIn).Return(dOut, 200, nil)
	floatingIPOut, err := ds.CreateFloatingIP(mapIn)
	assert.Nil(err, "Error creating floating IP list")
	assert.Equal(floatingIPIn, floatingIPOut, "CreateFloatingIP returned different floating IPs")

	return floatingIPOut
}

// CreateFloatingIPFailErrMocked test mocked function
func CreateFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", APIPathNetworkFloatingIPs, mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.CreateFloatingIP(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// CreateFloatingIPFailStatusMocked test mocked function
func CreateFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", APIPathNetworkFloatingIPs, mapIn).Return(dOut, 499, nil)
	floatingIPOut, err := ds.CreateFloatingIP(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPOut
}

// CreateFloatingIPFailJSONMocked test mocked function
func CreateFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", APIPathNetworkFloatingIPs, mapIn).Return(dIn, 200, nil)
	floatingIPOut, err := ds.CreateFloatingIP(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// UpdateFloatingIPMocked test mocked function
func UpdateFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID), mapIn).Return(dOut, 200, nil)
	floatingIPOut, err := ds.UpdateFloatingIP(floatingIPIn.ID, mapIn)
	assert.Nil(err, "Error updating floating IP list")
	assert.Equal(floatingIPIn, floatingIPOut, "UpdateFloatingIP returned different floating IPs")

	return floatingIPOut
}

// UpdateFloatingIPFailErrMocked test mocked function
func UpdateFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.UpdateFloatingIP(floatingIPIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// UpdateFloatingIPFailStatusMocked test mocked function
func UpdateFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID), mapIn).Return(dOut, 499, nil)
	floatingIPOut, err := ds.UpdateFloatingIP(floatingIPIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return floatingIPOut
}

// UpdateFloatingIPFailJSONMocked test mocked function
func UpdateFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID), mapIn).Return(dIn, 200, nil)
	floatingIPOut, err := ds.UpdateFloatingIP(floatingIPIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// AttachFloatingIPMocked test mocked function
func AttachFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: floatingIPIn.AttachedServerID})
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPIn.ID), mapIn).
		Return(dOut, 200, nil)
	serverOut, err := ds.AttachFloatingIP(floatingIPIn.ID, mapIn)
	assert.Nil(err, "Error attaching floating IP")
	assert.Equal(floatingIPIn.AttachedServerID, serverOut.ID, "AttachFloatingIP returned invalid values")

	return serverOut
}

// AttachFloatingIPFailErrMocked test mocked function
func AttachFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: floatingIPIn.AttachedServerID})
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.AttachFloatingIP(floatingIPIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// AttachFloatingIPFailStatusMocked test mocked function
func AttachFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// to json
	dOut, err := json.Marshal(types.Server{ID: floatingIPIn.AttachedServerID})
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPIn.ID), mapIn).
		Return(dOut, 499, nil)
	serverOut, err := ds.AttachFloatingIP(floatingIPIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return serverOut
}

// AttachFloatingIPFailJSONMocked test mocked function
func AttachFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.AttachFloatingIP(floatingIPIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// DetachFloatingIPMocked test mocked function
func DetachFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPIn.ID)).Return(dIn, 200, nil)
	err = ds.DetachFloatingIP(floatingIPIn.ID)
	assert.Nil(err, "Error detaching floating IP")
}

// DetachFloatingIPFailErrMocked test mocked function
func DetachFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DetachFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DetachFloatingIPFailStatusMocked test mocked function
func DetachFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIPAttachedServer, floatingIPIn.ID)).Return(dIn, 499, nil)
	err = ds.DetachFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteFloatingIPMocked test mocked function
func DeleteFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteFloatingIP(floatingIPIn.ID)
	assert.Nil(err, "Error deleting floating IP")
}

// DeleteFloatingIPFailErrMocked test mocked function
func DeleteFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteFloatingIPFailStatusMocked test mocked function
func DeleteFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIP, floatingIPIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DiscardFloatingIPMocked test mocked function
func DiscardFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIPDiscard, floatingIPIn.ID)).Return(dIn, 200, nil)
	err = ds.DiscardFloatingIP(floatingIPIn.ID)
	assert.Nil(err, "Error discarding floating IP")
}

// DiscardFloatingIPFailErrMocked test mocked function
func DiscardFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIPDiscard, floatingIPIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DiscardFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DiscardFloatingIPFailStatusMocked test mocked function
func DiscardFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFloatingIPService(cs)
	assert.Nil(err, "Couldn't load floating IP service")
	assert.NotNil(ds, "FloatingIP service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "FloatingIP test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkFloatingIPDiscard, floatingIPIn.ID)).Return(dIn, 499, nil)
	err = ds.DiscardFloatingIP(floatingIPIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
