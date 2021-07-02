// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewVPNServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewVPNService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetVPN(t *testing.T) {
	vpcsIn := testdata.GetVPNData()
	for _, vpcIn := range vpcsIn {
		GetVPNMocked(t, vpcIn)
		GetVPNFailErrMocked(t, vpcIn)
		GetVPNFailStatusMocked(t, vpcIn)
		GetVPNFailJSONMocked(t, vpcIn)
	}
}

func TestCreateVPN(t *testing.T) {
	vpcsIn := testdata.GetVPNData()
	for _, vpcIn := range vpcsIn {
		CreateVPNMocked(t, vpcIn)
		CreateVPNFailErrMocked(t, vpcIn)
		CreateVPNFailStatusMocked(t, vpcIn)
		CreateVPNFailJSONMocked(t, vpcIn)
	}
}

func TestDeleteVPN(t *testing.T) {
	vpcsIn := testdata.GetVPNData()
	for _, vpcIn := range vpcsIn {
		DeleteVPNMocked(t, vpcIn)
		DeleteVPNFailErrMocked(t, vpcIn)
		DeleteVPNFailStatusMocked(t, vpcIn)
	}
}

func TestListVPNPlans(t *testing.T) {
	vpnPlansIn := testdata.GetVPNPlanData()
	ListVPNPlansMocked(t, vpnPlansIn, "fakeVpcID0")
	ListVPNPlansFailErrMocked(t, vpnPlansIn, "fakeVpcID0")
	ListVPNPlansFailStatusMocked(t, vpnPlansIn, "fakeVpcID0")
	ListVPNPlansFailJSONMocked(t, vpnPlansIn, "fakeVpcID0")
}
