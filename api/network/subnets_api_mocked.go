package network

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetSubnetListMocked test mocked function
func GetSubnetListMocked(t *testing.T, subnetsIn []*types.Subnet) []*types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetsIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/subnets", subnetsIn[0].VpcID)).Return(dIn, 200, nil)
	subnetsOut, err := ds.GetSubnetList(subnetsIn[0].VpcID)
	assert.Nil(err, "Error getting Subnet list")
	assert.Equal(subnetsIn, subnetsOut, "GetSubnetList returned different Subnets")

	return subnetsOut
}

// GetSubnetListFailErrMocked test mocked function
func GetSubnetListFailErrMocked(t *testing.T, subnetsIn []*types.Subnet) []*types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetsIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/subnets", subnetsIn[0].VpcID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	subnetsOut, err := ds.GetSubnetList(subnetsIn[0].VpcID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(subnetsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return subnetsOut
}

// GetSubnetListFailStatusMocked test mocked function
func GetSubnetListFailStatusMocked(t *testing.T, subnetsIn []*types.Subnet) []*types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetsIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/subnets", subnetsIn[0].VpcID)).Return(dIn, 499, nil)
	subnetsOut, err := ds.GetSubnetList(subnetsIn[0].VpcID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(subnetsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return subnetsOut
}

// GetSubnetListFailJSONMocked test mocked function
func GetSubnetListFailJSONMocked(t *testing.T, subnetsIn []*types.Subnet) []*types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/vpcs/%s/subnets", subnetsIn[0].VpcID)).Return(dIn, 200, nil)
	subnetsOut, err := ds.GetSubnetList(subnetsIn[0].VpcID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(subnetsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return subnetsOut
}

// GetSubnetMocked test mocked function
func GetSubnetMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/subnets/%s", subnetIn.ID)).Return(dIn, 200, nil)
	subnetOut, err := ds.GetSubnet(subnetIn.ID)
	assert.Nil(err, "Error getting Subnet")
	assert.Equal(*subnetIn, *subnetOut, "GetSubnet returned different Subnets")

	return subnetOut
}

// GetSubnetFailErrMocked test mocked function
func GetSubnetFailErrMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/subnets/%s", subnetIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	subnetOut, err := ds.GetSubnet(subnetIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return subnetOut
}

// GetSubnetFailStatusMocked test mocked function
func GetSubnetFailStatusMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/subnets/%s", subnetIn.ID)).Return(dIn, 499, nil)
	subnetOut, err := ds.GetSubnet(subnetIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return subnetOut
}

// GetSubnetFailJSONMocked test mocked function
func GetSubnetFailJSONMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/subnets/%s", subnetIn.ID)).Return(dIn, 200, nil)
	subnetOut, err := ds.GetSubnet(subnetIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return subnetOut
}

// CreateSubnetMocked test mocked function
func CreateSubnetMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// to json
	dOut, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/vpcs/%s/subnets", subnetIn.VpcID), mapIn).Return(dOut, 200, nil)
	subnetOut, err := ds.CreateSubnet(mapIn, subnetIn.VpcID)
	assert.Nil(err, "Error creating Subnet list")
	assert.Equal(subnetIn, subnetOut, "CreateSubnet returned different Subnets")

	return subnetOut
}

// CreateSubnetFailErrMocked test mocked function
func CreateSubnetFailErrMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// to json
	dOut, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/vpcs/%s/subnets", subnetIn.VpcID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	subnetOut, err := ds.CreateSubnet(mapIn, subnetIn.VpcID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return subnetOut
}

// CreateSubnetFailStatusMocked test mocked function
func CreateSubnetFailStatusMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// to json
	dOut, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/vpcs/%s/subnets", subnetIn.VpcID), mapIn).Return(dOut, 499, nil)
	subnetOut, err := ds.CreateSubnet(mapIn, subnetIn.VpcID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return subnetOut
}

// CreateSubnetFailJSONMocked test mocked function
func CreateSubnetFailJSONMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/network/vpcs/%s/subnets", subnetIn.VpcID), mapIn).Return(dIn, 200, nil)
	subnetOut, err := ds.CreateSubnet(mapIn, subnetIn.VpcID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return subnetOut
}

// UpdateSubnetMocked test mocked function
func UpdateSubnetMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// to json
	dOut, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/subnets/%s", subnetIn.ID), mapIn).Return(dOut, 200, nil)
	subnetOut, err := ds.UpdateSubnet(mapIn, subnetIn.ID)
	assert.Nil(err, "Error updating Subnet list")
	assert.Equal(subnetIn, subnetOut, "UpdateSubnet returned different Subnets")

	return subnetOut
}

// UpdateSubnetFailErrMocked test mocked function
func UpdateSubnetFailErrMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// to json
	dOut, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/subnets/%s", subnetIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	subnetOut, err := ds.UpdateSubnet(mapIn, subnetIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return subnetOut
}

// UpdateSubnetFailStatusMocked test mocked function
func UpdateSubnetFailStatusMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// to json
	dOut, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/subnets/%s", subnetIn.ID), mapIn).Return(dOut, 499, nil)
	subnetOut, err := ds.UpdateSubnet(mapIn, subnetIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return subnetOut
}

// UpdateSubnetFailJSONMocked test mocked function
func UpdateSubnetFailJSONMocked(t *testing.T, subnetIn *types.Subnet) *types.Subnet {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/network/subnets/%s", subnetIn.ID), mapIn).Return(dIn, 200, nil)
	subnetOut, err := ds.UpdateSubnet(mapIn, subnetIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(subnetOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return subnetOut
}

// DeleteSubnetMocked test mocked function
func DeleteSubnetMocked(t *testing.T, subnetIn *types.Subnet) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/subnets/%s", subnetIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteSubnet(subnetIn.ID)
	assert.Nil(err, "Error deleting Subnet")
}

// DeleteSubnetFailErrMocked test mocked function
func DeleteSubnetFailErrMocked(t *testing.T, subnetIn *types.Subnet) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/subnets/%s", subnetIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteSubnet(subnetIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteSubnetFailStatusMocked test mocked function
func DeleteSubnetFailStatusMocked(t *testing.T, subnetIn *types.Subnet) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(subnetIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/subnets/%s", subnetIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteSubnet(subnetIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// GetSubnetServersListMocked test mocked function
func GetSubnetServersListMocked(t *testing.T, serversIn []*types.Server) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/subnets/%s/servers", serversIn[0].VpcID)).Return(dIn, 200, nil)
	serversOut, err := ds.GetSubnetServersList(serversIn[0].VpcID)
	assert.Nil(err, "Error getting Subnet servers list")
	assert.Equal(serversIn, serversOut, "GetSubnetServersList returned different Servers")

	return serversOut
}

// GetSubnetServersListFailErrMocked test mocked function
func GetSubnetServersListFailErrMocked(t *testing.T, serversIn []*types.Server) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/subnets/%s/servers", serversIn[0].VpcID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	serversOut, err := ds.GetSubnetServersList(serversIn[0].VpcID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serversOut
}

// GetSubnetServersListFailStatusMocked test mocked function
func GetSubnetServersListFailStatusMocked(t *testing.T, serversIn []*types.Server) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Subnet test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/subnets/%s/servers", serversIn[0].VpcID)).Return(dIn, 499, nil)
	serversOut, err := ds.GetSubnetServersList(serversIn[0].VpcID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serversOut
}

// GetSubnetServersListFailJSONMocked test mocked function
func GetSubnetServersListFailJSONMocked(t *testing.T, serversIn []*types.Server) []*types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSubnetService(cs)
	assert.Nil(err, "Couldn't load Subnet service")
	assert.NotNil(ds, "Subnet service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/subnets/%s/servers", serversIn[0].VpcID)).Return(dIn, 200, nil)
	serversOut, err := ds.GetSubnetServersList(serversIn[0].VpcID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serversOut
}
