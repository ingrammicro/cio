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

// ListWizardServerPlansMocked test mocked function
func ListWizardServerPlansMocked(
	t *testing.T,
	serverPlansIn []*types.ServerPlan,
	AppID string,
	LocID string,
	ProviderID string,
) []*types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizardServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "WizServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlansIn)
	assert.Nil(err, "WizServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathWizardServerPlansByAppLocationCloudProvider, AppID, LocID, ProviderID)).
		Return(dIn, 200, nil)
	serverPlansOut, err := ds.ListWizardServerPlans(AppID, LocID, ProviderID)
	assert.Nil(err, "Error getting serverPlan list")
	assert.Equal(serverPlansIn, serverPlansOut, "ListWizardServerPlans returned different serverPlans")

	return serverPlansOut
}

// ListWizardServerPlansFailErrMocked test mocked function
func ListWizardServerPlansFailErrMocked(
	t *testing.T,
	serverPlansIn []*types.ServerPlan,
	AppID string,
	LocID string,
	ProviderID string,
) []*types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizardServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "WizServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlansIn)
	assert.Nil(err, "WizServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathWizardServerPlansByAppLocationCloudProvider, AppID, LocID, ProviderID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	serverPlansOut, err := ds.ListWizardServerPlans(AppID, LocID, ProviderID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverPlansOut
}

// ListWizardServerPlansFailStatusMocked test mocked function
func ListWizardServerPlansFailStatusMocked(
	t *testing.T,
	serverPlansIn []*types.ServerPlan,
	AppID string,
	LocID string,
	ProviderID string,
) []*types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizardServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "WizServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlansIn)
	assert.Nil(err, "WizServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathWizardServerPlansByAppLocationCloudProvider, AppID, LocID, ProviderID)).
		Return(dIn, 499, nil)
	serverPlansOut, err := ds.ListWizardServerPlans(AppID, LocID, ProviderID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverPlansOut
}

// ListWizardServerPlansFailJSONMocked test mocked function
func ListWizardServerPlansFailJSONMocked(
	t *testing.T,
	serverPlansIn []*types.ServerPlan,
	AppID string,
	LocID string,
	ProviderID string,
) []*types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizardServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "WizServerPlan service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathWizardServerPlansByAppLocationCloudProvider, AppID, LocID, ProviderID)).
		Return(dIn, 200, nil)
	serverPlansOut, err := ds.ListWizardServerPlans(AppID, LocID, ProviderID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverPlansOut
}
