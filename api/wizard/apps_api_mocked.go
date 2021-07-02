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

// TODO exclude from release compile

// ListAppsMocked test mocked function
func ListAppsMocked(t *testing.T, appsIn []*types.WizardApp) []*types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// to json
	dIn, err := json.Marshal(appsIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Get", APIPathWizardApps).Return(dIn, 200, nil)
	appsOut, err := ds.ListApps()
	assert.Nil(err, "Error getting app list")
	assert.Equal(appsIn, appsOut, "ListApps returned different apps")

	return appsOut
}

// ListAppsFailErrMocked test mocked function
func ListAppsFailErrMocked(t *testing.T, appsIn []*types.WizardApp) []*types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// to json
	dIn, err := json.Marshal(appsIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Get", APIPathWizardApps).Return(dIn, 200, fmt.Errorf("mocked error"))
	appsOut, err := ds.ListApps()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(appsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return appsOut
}

// ListAppsFailStatusMocked test mocked function
func ListAppsFailStatusMocked(t *testing.T, appsIn []*types.WizardApp) []*types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// to json
	dIn, err := json.Marshal(appsIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Get", APIPathWizardApps).Return(dIn, 499, nil)
	appsOut, err := ds.ListApps()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(appsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return appsOut
}

// ListAppsFailJSONMocked test mocked function
func ListAppsFailJSONMocked(t *testing.T, appsIn []*types.WizardApp) []*types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathWizardApps).Return(dIn, 200, nil)
	appsOut, err := ds.ListApps()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(appsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return appsOut
}

// DeployAppMocked test mocked function
func DeployAppMocked(t *testing.T, appIn *types.WizardApp) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*appIn)
	assert.Nil(err, "App test data corrupted")

	// to json
	dOut, err := json.Marshal(appIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathWizardAppDeploy, appIn.ID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.DeployApp(appIn.ID, mapIn)
	assert.Nil(err, "Error deploying app")

	return serverOut
}

// DeployAppFailErrMocked test mocked function
func DeployAppFailErrMocked(t *testing.T, appIn *types.WizardApp) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*appIn)
	assert.Nil(err, "App test data corrupted")

	// to json
	dOut, err := json.Marshal(appIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathWizardAppDeploy, appIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.DeployApp(appIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// DeployAppFailStatusMocked test mocked function
func DeployAppFailStatusMocked(t *testing.T, appIn *types.WizardApp) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*appIn)
	assert.Nil(err, "App test data corrupted")

	// to json
	dOut, err := json.Marshal(appIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathWizardAppDeploy, appIn.ID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.DeployApp(appIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// DeployAppFailJSONMocked test mocked function
func DeployAppFailJSONMocked(t *testing.T, appIn *types.WizardApp) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*appIn)
	assert.Nil(err, "App test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathWizardAppDeploy, appIn.ID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.DeployApp(appIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}
