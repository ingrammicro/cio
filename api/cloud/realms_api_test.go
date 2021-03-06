// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewRealmServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewRealmService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListRealms(t *testing.T) {
	realmsIn := testdata.GetRealmData()
	ListRealmsMocked(t, realmsIn, realmsIn[0].CloudProviderID)
	ListRealmsFailErrMocked(t, realmsIn, realmsIn[0].CloudProviderID)
	ListRealmsFailStatusMocked(t, realmsIn, realmsIn[0].CloudProviderID)
	ListRealmsFailJSONMocked(t, realmsIn, realmsIn[0].CloudProviderID)
}

func TestGetRealm(t *testing.T) {
	realmsIn := testdata.GetRealmData()
	for _, realmIn := range realmsIn {
		GetRealmMocked(t, realmIn)
		GetRealmFailErrMocked(t, realmIn)
		GetRealmFailStatusMocked(t, realmIn)
		GetRealmFailJSONMocked(t, realmIn)
	}
}

func TestListRealmNodePoolPlans(t *testing.T) {
	nodePoolPlansIn := testdata.GetNodePoolPlanData()
	ListRealmNodePoolPlansMocked(t, nodePoolPlansIn, nodePoolPlansIn[0].RealmID)
	ListRealmNodePoolPlansFailErrMocked(t, nodePoolPlansIn, nodePoolPlansIn[0].RealmID)
	ListRealmNodePoolPlansFailStatusMocked(t, nodePoolPlansIn, nodePoolPlansIn[0].RealmID)
	ListRealmNodePoolPlansFailJSONMocked(t, nodePoolPlansIn, nodePoolPlansIn[0].RealmID)
}
