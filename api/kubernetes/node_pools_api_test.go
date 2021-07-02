// Copyright (c) 2017-2021 Ingram Micro Inc.

package kubernetes

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewNodePoolServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewNodePoolService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListNodePools(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	nodePoolsIn := testdata.GetNodePoolData()
	ListNodePoolsMocked(t, clustersIn[0].ID, nodePoolsIn)
	ListNodePoolsFailErrMocked(t, clustersIn[0].ID, nodePoolsIn)
	ListNodePoolsFailStatusMocked(t, clustersIn[0].ID, nodePoolsIn)
	ListNodePoolsFailJSONMocked(t, clustersIn[0].ID, nodePoolsIn)
}

func TestGetNodePool(t *testing.T) {
	nodePoolsIn := testdata.GetNodePoolData()
	for _, nodePoolIn := range nodePoolsIn {
		GetNodePoolMocked(t, nodePoolIn)
		GetNodePoolFailErrMocked(t, nodePoolIn)
		GetNodePoolFailStatusMocked(t, nodePoolIn)
		GetNodePoolFailJSONMocked(t, nodePoolIn)
	}
}

func TestCreateNodePool(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	nodePoolsIn := testdata.GetNodePoolData()
	for _, nodePoolIn := range nodePoolsIn {
		CreateNodePoolMocked(t, clustersIn[0].ID, nodePoolIn)
		CreateNodePoolFailErrMocked(t, clustersIn[0].ID, nodePoolIn)
		CreateNodePoolFailStatusMocked(t, clustersIn[0].ID, nodePoolIn)
		CreateNodePoolFailJSONMocked(t, clustersIn[0].ID, nodePoolIn)
	}
}

func TestUpdateNodePool(t *testing.T) {
	nodePoolsIn := testdata.GetNodePoolData()
	for _, nodePoolIn := range nodePoolsIn {
		UpdateNodePoolMocked(t, nodePoolIn)
		UpdateNodePoolFailErrMocked(t, nodePoolIn)
		UpdateNodePoolFailStatusMocked(t, nodePoolIn)
		UpdateNodePoolFailJSONMocked(t, nodePoolIn)
	}
}

func TestDeleteNodePool(t *testing.T) {
	nodePoolsIn := testdata.GetNodePoolData()
	for _, nodePoolIn := range nodePoolsIn {
		DeleteNodePoolMocked(t, nodePoolIn)
		DeleteNodePoolFailErrMocked(t, nodePoolIn)
		DeleteNodePoolFailStatusMocked(t, nodePoolIn)
		DeleteNodePoolFailJSONMocked(t, nodePoolIn)
	}
}

func TestRetryNodePool(t *testing.T) {
	nodePoolsIn := testdata.GetNodePoolData()
	for _, nodePoolIn := range nodePoolsIn {
		RetryNodePoolMocked(t, nodePoolIn)
		RetryNodePoolFailErrMocked(t, nodePoolIn)
		RetryNodePoolFailStatusMocked(t, nodePoolIn)
		RetryNodePoolFailJSONMocked(t, nodePoolIn)
	}
}

func TestGetNodePoolPlan(t *testing.T) {
	nodePoolPlansIn := testdata.GetNodePoolPlanData()
	for _, nodePoolPlanIn := range nodePoolPlansIn {
		GetNodePoolPlanMocked(t, nodePoolPlanIn.ID, nodePoolPlanIn)
		GetNodePoolPlanFailErrMocked(t, nodePoolPlanIn.ID, nodePoolPlanIn)
		GetNodePoolPlanFailStatusMocked(t, nodePoolPlanIn.ID, nodePoolPlanIn)
		GetNodePoolPlanFailJSONMocked(t, nodePoolPlanIn.ID, nodePoolPlanIn)
	}
}
