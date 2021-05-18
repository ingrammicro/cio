// Copyright (c) 2017-2021 Ingram Micro Inc.

package wizard

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// ListWizardCloudProvidersMocked test mocked function
func ListWizardCloudProvidersMocked(
	t *testing.T,
	cloudProvidersIn []*types.CloudProvider,
	AppID string,
	LocID string,
) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizardCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "WizCloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "WizCloudProvider test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathWizardCloudProvidersByAppLocation, AppID, LocID)).Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.ListWizardCloudProviders(AppID, LocID)
	assert.Nil(err, "Error getting cloudProvider list")
	assert.Equal(cloudProvidersIn, cloudProvidersOut, "ListWizardCloudProviders returned different cloudProviders")

	return cloudProvidersOut
}

// ListWizardCloudProvidersFailErrMocked test mocked function
func ListWizardCloudProvidersFailErrMocked(
	t *testing.T,
	cloudProvidersIn []*types.CloudProvider,
	AppID string,
	LocID string,
) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizardCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "WizCloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "WizCloudProvider test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathWizardCloudProvidersByAppLocation, AppID, LocID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudProvidersOut, err := ds.ListWizardCloudProviders(AppID, LocID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudProvidersOut
}

// ListWizardCloudProvidersFailStatusMocked test mocked function
func ListWizardCloudProvidersFailStatusMocked(
	t *testing.T,
	cloudProvidersIn []*types.CloudProvider,
	AppID string,
	LocID string,
) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizardCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "WizCloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "WizCloudProvider test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathWizardCloudProvidersByAppLocation, AppID, LocID)).Return(dIn, 499, nil)
	cloudProvidersOut, err := ds.ListWizardCloudProviders(AppID, LocID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudProvidersOut
}

// ListWizardCloudProvidersFailJSONMocked test mocked function
func ListWizardCloudProvidersFailJSONMocked(
	t *testing.T,
	cloudProvidersIn []*types.CloudProvider,
	AppID string,
	LocID string,
) []*types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizardCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "WizCloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathWizardCloudProvidersByAppLocation, AppID, LocID)).Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.ListWizardCloudProviders(AppID, LocID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudProvidersOut
}
