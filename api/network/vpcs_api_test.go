// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewVPCServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewVPCService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListVPCs(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	ListVPCsMocked(t, vpcsIn)
	ListVPCsFailErrMocked(t, vpcsIn)
	ListVPCsFailStatusMocked(t, vpcsIn)
	ListVPCsFailJSONMocked(t, vpcsIn)
}

func TestGetVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		GetVPCMocked(t, vpcIn)
		GetVPCFailErrMocked(t, vpcIn)
		GetVPCFailStatusMocked(t, vpcIn)
		GetVPCFailJSONMocked(t, vpcIn)
	}
}

func TestCreateVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		CreateVPCMocked(t, vpcIn)
		CreateVPCFailErrMocked(t, vpcIn)
		CreateVPCFailStatusMocked(t, vpcIn)
		CreateVPCFailJSONMocked(t, vpcIn)
	}
}

func TestUpdateVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		UpdateVPCMocked(t, vpcIn)
		UpdateVPCFailErrMocked(t, vpcIn)
		UpdateVPCFailStatusMocked(t, vpcIn)
		UpdateVPCFailJSONMocked(t, vpcIn)
	}
}

func TestDeleteVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		DeleteVPCMocked(t, vpcIn)
		DeleteVPCFailErrMocked(t, vpcIn)
		DeleteVPCFailStatusMocked(t, vpcIn)
	}
}

func TestDiscardVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	for _, vpcIn := range vpcsIn {
		DiscardVPCMocked(t, vpcIn)
		DiscardVPCFailErrMocked(t, vpcIn)
		DiscardVPCFailStatusMocked(t, vpcIn)
	}
}
