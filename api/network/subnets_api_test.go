// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewSubnetServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewSubnetService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListSubnets(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	ListSubnetsMocked(t, subnetsIn)
	ListSubnetsFailErrMocked(t, subnetsIn)
	ListSubnetsFailStatusMocked(t, subnetsIn)
	ListSubnetsFailJSONMocked(t, subnetsIn)
}

func TestGetSubnet(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	for _, subnetIn := range subnetsIn {
		GetSubnetMocked(t, subnetIn)
		GetSubnetFailErrMocked(t, subnetIn)
		GetSubnetFailStatusMocked(t, subnetIn)
		GetSubnetFailJSONMocked(t, subnetIn)
	}
}

func TestCreateSubnet(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	for _, subnetIn := range subnetsIn {
		CreateSubnetMocked(t, subnetIn)
		CreateSubnetFailErrMocked(t, subnetIn)
		CreateSubnetFailStatusMocked(t, subnetIn)
		CreateSubnetFailJSONMocked(t, subnetIn)
	}
}

func TestUpdateSubnet(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	for _, subnetIn := range subnetsIn {
		UpdateSubnetMocked(t, subnetIn)
		UpdateSubnetFailErrMocked(t, subnetIn)
		UpdateSubnetFailStatusMocked(t, subnetIn)
		UpdateSubnetFailJSONMocked(t, subnetIn)
	}
}

func TestDeleteSubnet(t *testing.T) {
	subnetsIn := testdata.GetSubnetData()
	for _, subnetIn := range subnetsIn {
		DeleteSubnetMocked(t, subnetIn)
		DeleteSubnetFailErrMocked(t, subnetIn)
		DeleteSubnetFailStatusMocked(t, subnetIn)
	}
}

func TestListSubnetServers(t *testing.T) {
	serversIn := testdata.GetSubnetServersData()
	ListSubnetServersMocked(t, serversIn)
	ListSubnetServersFailErrMocked(t, serversIn)
	ListSubnetServersFailStatusMocked(t, serversIn)
	ListSubnetServersFailJSONMocked(t, serversIn)
}

func TestListSubnetServerArrays(t *testing.T) {
	serverArraysIn := testdata.GetServerArrayData()
	ListSubnetServerArraysMocked(t, serverArraysIn)
	ListSubnetServerArraysFailErrMocked(t, serverArraysIn)
	ListSubnetServerArraysFailStatusMocked(t, serverArraysIn)
	ListSubnetServerArraysFailJSONMocked(t, serverArraysIn)
}
