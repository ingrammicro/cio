package network

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTargetGroupServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewTargetGroupService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListTargetGroups(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	targetGroupsIn := testdata.GetTargetGroupData()
	ListTargetGroupsMocked(t, loadBalancersIn[0].ID, targetGroupsIn)
	ListTargetGroupsFailErrMocked(t, loadBalancersIn[0].ID, targetGroupsIn)
	ListTargetGroupsFailStatusMocked(t, loadBalancersIn[0].ID, targetGroupsIn)
	ListTargetGroupsFailJSONMocked(t, loadBalancersIn[0].ID, targetGroupsIn)
}

func TestGetTargetGroup(t *testing.T) {
	targetGroupsIn := testdata.GetTargetGroupData()
	for _, targetGroupIn := range targetGroupsIn {
		GetTargetGroupMocked(t, targetGroupIn)
		GetTargetGroupFailErrMocked(t, targetGroupIn)
		GetTargetGroupFailStatusMocked(t, targetGroupIn)
		GetTargetGroupFailJSONMocked(t, targetGroupIn)
	}
}

func TestCreateTargetGroup(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	targetGroupsIn := testdata.GetTargetGroupData()
	for _, targetGroupIn := range targetGroupsIn {
		CreateTargetGroupMocked(t, loadBalancersIn[0].ID, targetGroupIn)
		CreateTargetGroupFailErrMocked(t, loadBalancersIn[0].ID, targetGroupIn)
		CreateTargetGroupFailStatusMocked(t, loadBalancersIn[0].ID, targetGroupIn)
		CreateTargetGroupFailJSONMocked(t, loadBalancersIn[0].ID, targetGroupIn)
	}
}

func TestUpdateTargetGroup(t *testing.T) {
	targetGroupsIn := testdata.GetTargetGroupData()
	for _, targetGroupIn := range targetGroupsIn {
		UpdateTargetGroupMocked(t, targetGroupIn)
		UpdateTargetGroupFailErrMocked(t, targetGroupIn)
		UpdateTargetGroupFailStatusMocked(t, targetGroupIn)
		UpdateTargetGroupFailJSONMocked(t, targetGroupIn)
	}
}

func TestDeleteTargetGroup(t *testing.T) {
	targetGroupsIn := testdata.GetTargetGroupData()
	for _, targetGroupIn := range targetGroupsIn {
		DeleteTargetGroupMocked(t, targetGroupIn)
		DeleteTargetGroupFailErrMocked(t, targetGroupIn)
		DeleteTargetGroupFailStatusMocked(t, targetGroupIn)
		DeleteTargetGroupFailJSONMocked(t, targetGroupIn)
	}
}

func TestRetryTargetGroup(t *testing.T) {
	targetGroupsIn := testdata.GetTargetGroupData()
	for _, targetGroupIn := range targetGroupsIn {
		RetryTargetGroupMocked(t, targetGroupIn)
		RetryTargetGroupFailErrMocked(t, targetGroupIn)
		RetryTargetGroupFailStatusMocked(t, targetGroupIn)
		RetryTargetGroupFailJSONMocked(t, targetGroupIn)
	}
}

func TestListTargets(t *testing.T) {
	targetGroupsIn := testdata.GetTargetGroupData()
	targetsIn := testdata.GetTargetData()
	ListTargetsMocked(t, targetGroupsIn[0].ID, targetsIn)
	ListTargetsFailErrMocked(t, targetGroupsIn[0].ID, targetsIn)
	ListTargetsFailStatusMocked(t, targetGroupsIn[0].ID, targetsIn)
	ListTargetsFailJSONMocked(t, targetGroupsIn[0].ID, targetsIn)
}

func TestCreateTarget(t *testing.T) {
	targetGroupsIn := testdata.GetTargetGroupData()
	targetsIn := testdata.GetTargetData()
	for _, targetIn := range targetsIn {
		CreateTargetMocked(t, targetGroupsIn[0].ID, targetIn)
		CreateTargetFailErrMocked(t, targetGroupsIn[0].ID, targetIn)
		CreateTargetFailStatusMocked(t, targetGroupsIn[0].ID, targetIn)
		CreateTargetFailJSONMocked(t, targetGroupsIn[0].ID, targetIn)
	}
}

func TestDeleteTarget(t *testing.T) {
	targetsIn := testdata.GetTargetData()
	targetGroupsIn := testdata.GetTargetGroupData()
	for _, targetIn := range targetsIn {
		DeleteTargetMocked(t, targetGroupsIn[0].ID, targetIn)
		DeleteTargetFailErrMocked(t, targetGroupsIn[0].ID, targetIn)
		DeleteTargetFailStatusMocked(t, targetGroupsIn[0].ID, targetIn)
	}
}
