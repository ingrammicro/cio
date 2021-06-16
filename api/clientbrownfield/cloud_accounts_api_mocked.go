// Copyright (c) 2017-2021 Ingram Micro Inc.

package clientbrownfield

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// ListBrownfieldCloudAccountsMocked test mocked function
func ListBrownfieldCloudAccountsMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccounts test data corrupted")

	// call service
	cs.On("Get", APIPathBlueprintCloudAccounts).Return(dIn, 200, nil)
	cloudAccountsOut, err := ds.ListBrownfieldCloudAccounts()
	assert.Nil(err, "Error getting brownfield cloud account list")
	assert.Equal(cloudAccountsIn, cloudAccountsOut, "ListBrownfieldCloudAccounts returned different cloud accounts")

	return cloudAccountsOut
}

// ListBrownfieldCloudAccountsFailErrMocked test mocked function
func ListBrownfieldCloudAccountsFailErrMocked(
	t *testing.T,
	cloudAccountsIn []*types.CloudAccount,
) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccounts test data corrupted")

	// call service
	cs.On("Get", APIPathBlueprintCloudAccounts).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudAccountsOut, err := ds.ListBrownfieldCloudAccounts()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountsOut
}

// ListBrownfieldCloudAccountsFailStatusMocked test mocked function
func ListBrownfieldCloudAccountsFailStatusMocked(
	t *testing.T,
	cloudAccountsIn []*types.CloudAccount,
) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccounts test data corrupted")

	// call service
	cs.On("Get", APIPathBlueprintCloudAccounts).Return(dIn, 499, nil)
	cloudAccountsOut, err := ds.ListBrownfieldCloudAccounts()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountsOut
}

// ListBrownfieldCloudAccountsFailJSONMocked test mocked function
func ListBrownfieldCloudAccountsFailJSONMocked(
	t *testing.T,
	cloudAccountsIn []*types.CloudAccount,
) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathBlueprintCloudAccounts).Return(dIn, 200, nil)
	cloudAccountsOut, err := ds.ListBrownfieldCloudAccounts()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountsOut
}

// GetBrownfieldCloudAccountMocked test mocked function
func GetBrownfieldCloudAccountMocked(
	t *testing.T,
	cloudAccountIn *types.CloudAccount,
	cloudAccountID string,
) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintCloudAccount, cloudAccountID)).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.GetBrownfieldCloudAccount(cloudAccountID)
	assert.Nil(err, "Error getting brownfield cloud account")
	assert.Equal(*cloudAccountIn, *cloudAccountOut, "GetBrownfieldCloudAccount returned different cloud account")

	return cloudAccountOut
}

// GetBrownfieldCloudAccountFailErrMocked test mocked function
func GetBrownfieldCloudAccountFailErrMocked(
	t *testing.T,
	cloudAccountIn *types.CloudAccount,
	cloudAccountID string,
) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintCloudAccount, cloudAccountID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.GetBrownfieldCloudAccount(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// GetBrownfieldCloudAccountFailStatusMocked test mocked function
func GetBrownfieldCloudAccountFailStatusMocked(
	t *testing.T,
	cloudAccountIn *types.CloudAccount,
	cloudAccountID string,
) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintCloudAccount, cloudAccountID)).Return(dIn, 499, nil)
	cloudAccountOut, err := ds.GetBrownfieldCloudAccount(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// GetBrownfieldCloudAccountFailJSONMocked test mocked function
func GetBrownfieldCloudAccountFailJSONMocked(
	t *testing.T,
	cloudAccountIn *types.CloudAccount,
	cloudAccountID string,
) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathBlueprintCloudAccount, cloudAccountID)).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.GetBrownfieldCloudAccount(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}
