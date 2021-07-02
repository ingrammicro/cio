// Copyright (c) 2017-2021 Ingram Micro Inc.

package wizard

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewAppServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewAppService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListApps(t *testing.T) {
	appsIn := testdata.GetAppData()
	ListAppsMocked(t, appsIn)
	ListAppsFailErrMocked(t, appsIn)
	ListAppsFailStatusMocked(t, appsIn)
	ListAppsFailJSONMocked(t, appsIn)
}

func TestDeployApp(t *testing.T) {
	appsIn := testdata.GetAppData()
	for _, appIn := range appsIn {
		DeployAppMocked(t, appIn)
		DeployAppFailErrMocked(t, appIn)
		DeployAppFailStatusMocked(t, appIn)
		DeployAppFailJSONMocked(t, appIn)
	}
}
