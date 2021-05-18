// Copyright (c) 2017-2021 Ingram Micro Inc.

package wizard

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewWizardCloudProviderServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewWizardCloudProviderService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListWizardCloudProviders(t *testing.T) {
	AppID := "fakeAppID"
	LocID := "fakeLocID"
	cloudProvidersIn := testdata.GetCloudProviderData()
	ListWizardCloudProvidersMocked(t, cloudProvidersIn, AppID, LocID)
	ListWizardCloudProvidersFailErrMocked(t, cloudProvidersIn, AppID, LocID)
	ListWizardCloudProvidersFailStatusMocked(t, cloudProvidersIn, AppID, LocID)
	ListWizardCloudProvidersFailJSONMocked(t, cloudProvidersIn, AppID, LocID)
}
