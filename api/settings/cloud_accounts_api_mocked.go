package settings

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// ListCloudAccountsMocked test mocked function
func ListCloudAccountsMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloud account service")
	assert.NotNil(ds, "Cloud account service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "Cloud account test data corrupted")

	// call service
	cs.On("Get", "/settings/cloud_accounts").Return(dIn, 200, nil)
	cloudAccountsOut, err := ds.ListCloudAccounts()
	assert.Nil(err, "Error getting cloud account list")
	assert.Equal(cloudAccountsIn, cloudAccountsOut, "ListCloudAccounts returned different cloud accounts")

	return cloudAccountsOut
}

// ListCloudAccountsFailErrMocked test mocked function
func ListCloudAccountsFailErrMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloud account service")
	assert.NotNil(ds, "Cloud account service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "Cloud account test data corrupted")

	// call service
	cs.On("Get", "/settings/cloud_accounts").Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudAccountsOut, err := ds.ListCloudAccounts()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountsOut
}

// ListCloudAccountsFailStatusMocked test mocked function
func ListCloudAccountsFailStatusMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloud account service")
	assert.NotNil(ds, "Cloud account service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "Cloud account test data corrupted")

	// call service
	cs.On("Get", "/settings/cloud_accounts").Return(dIn, 499, nil)
	cloudAccountsOut, err := ds.ListCloudAccounts()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountsOut
}

// ListCloudAccountsFailJSONMocked test mocked function
func ListCloudAccountsFailJSONMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloud account service")
	assert.NotNil(ds, "Cloud account service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/settings/cloud_accounts").Return(dIn, 200, nil)
	cloudAccountsOut, err := ds.ListCloudAccounts()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountsOut
}

// GetCloudAccountMocked test mocked function
func GetCloudAccountMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloud account service")
	assert.NotNil(ds, "Cloud account service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "Cloud account test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/settings/cloud_accounts/%s", cloudAccountIn.ID)).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.GetCloudAccount(cloudAccountIn.ID)
	assert.Nil(err, "Error getting cloud account")
	assert.Equal(*cloudAccountIn, *cloudAccountOut, "GetCloudAccount returned different cloud account")

	return cloudAccountOut
}

// GetCloudAccountFailErrMocked test mocked function
func GetCloudAccountFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloud account service")
	assert.NotNil(ds, "Cloud account service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "Cloud account test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/settings/cloud_accounts/%s", cloudAccountIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.GetCloudAccount(cloudAccountIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// GetCloudAccountFailStatusMocked test mocked function
func GetCloudAccountFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloud account service")
	assert.NotNil(ds, "Cloud account service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "Cloud account test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/settings/cloud_accounts/%s", cloudAccountIn.ID)).Return(dIn, 499, nil)
	cloudAccountOut, err := ds.GetCloudAccount(cloudAccountIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// GetCloudAccountFailJSONMocked test mocked function
func GetCloudAccountFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloud account service")
	assert.NotNil(ds, "Cloud account service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/settings/cloud_accounts/%s", cloudAccountIn.ID)).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.GetCloudAccount(cloudAccountIn.ID)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}
