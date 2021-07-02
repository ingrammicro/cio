// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewLoadBalancerServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewLoadBalancerService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListLoadBalancers(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	ListLoadBalancersMocked(t, loadBalancersIn)
	ListLoadBalancersFailErrMocked(t, loadBalancersIn)
	ListLoadBalancersFailStatusMocked(t, loadBalancersIn)
	ListLoadBalancersFailJSONMocked(t, loadBalancersIn)
}

func TestGetLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range loadBalancersIn {
		GetLoadBalancerMocked(t, loadBalancerIn)
		GetLoadBalancerFailErrMocked(t, loadBalancerIn)
		GetLoadBalancerFailStatusMocked(t, loadBalancerIn)
		GetLoadBalancerFailJSONMocked(t, loadBalancerIn)
	}
}

func TestCreateLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range loadBalancersIn {
		CreateLoadBalancerMocked(t, loadBalancerIn)
		CreateLoadBalancerFailErrMocked(t, loadBalancerIn)
		CreateLoadBalancerFailStatusMocked(t, loadBalancerIn)
		CreateLoadBalancerFailJSONMocked(t, loadBalancerIn)
	}
}

func TestUpdateLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range loadBalancersIn {
		UpdateLoadBalancerMocked(t, loadBalancerIn)
		UpdateLoadBalancerFailErrMocked(t, loadBalancerIn)
		UpdateLoadBalancerFailStatusMocked(t, loadBalancerIn)
		UpdateLoadBalancerFailJSONMocked(t, loadBalancerIn)
	}
}

func TestDeleteLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range loadBalancersIn {
		DeleteLoadBalancerMocked(t, loadBalancerIn)
		DeleteLoadBalancerFailErrMocked(t, loadBalancerIn)
		DeleteLoadBalancerFailStatusMocked(t, loadBalancerIn)
		DeleteLoadBalancerFailJSONMocked(t, loadBalancerIn)
	}
}

func TestRetryLoadBalancer(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	for _, loadBalancerIn := range loadBalancersIn {
		RetryLoadBalancerMocked(t, loadBalancerIn)
		RetryLoadBalancerFailErrMocked(t, loadBalancerIn)
		RetryLoadBalancerFailStatusMocked(t, loadBalancerIn)
		RetryLoadBalancerFailJSONMocked(t, loadBalancerIn)
	}
}

func TestGetLoadBalancerPlan(t *testing.T) {
	loadBalancerPlansIn := testdata.GetLoadBalancerPlanData()
	for _, loadBalancerPlanIn := range loadBalancerPlansIn {
		GetLoadBalancerPlanMocked(t, loadBalancerPlanIn.ID, loadBalancerPlanIn)
		GetLoadBalancerPlanFailErrMocked(t, loadBalancerPlanIn.ID, loadBalancerPlanIn)
		GetLoadBalancerPlanFailStatusMocked(t, loadBalancerPlanIn.ID, loadBalancerPlanIn)
		GetLoadBalancerPlanFailJSONMocked(t, loadBalancerPlanIn.ID, loadBalancerPlanIn)
	}
}
