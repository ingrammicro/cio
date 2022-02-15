// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewServerPlanServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewServerPlanService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListServerPlans(t *testing.T) {
	serverPlansIn := testdata.GetServerPlanData()
	for _, serverPlanIn := range serverPlansIn {
		ListServerPlansMocked(t, serverPlansIn, serverPlanIn.CloudProviderID, serverPlanIn.RealmID)
		ListServerPlansFailErrMocked(t, serverPlansIn, serverPlanIn.CloudProviderID, serverPlanIn.RealmID)
		ListServerPlansFailStatusMocked(t, serverPlansIn, serverPlanIn.CloudProviderID, serverPlanIn.RealmID)
		ListServerPlansFailJSONMocked(t, serverPlansIn, serverPlanIn.CloudProviderID, serverPlanIn.RealmID)
	}
}

func TestGetServerPlan(t *testing.T) {
	serverPlansIn := testdata.GetServerPlanData()
	for _, serverPlanIn := range serverPlansIn {
		GetServerPlanMocked(t, serverPlanIn)
		GetServerPlanFailErrMocked(t, serverPlanIn)
		GetServerPlanFailStatusMocked(t, serverPlanIn)
		GetServerPlanFailJSONMocked(t, serverPlanIn)
	}
}
