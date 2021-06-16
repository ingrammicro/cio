// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudapplication

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewCloudApplicationDeploymentServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudApplicationDeploymentService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListDeployments(t *testing.T) {
	cloudApplicationDeploymentsIn := testdata.GetCloudApplicationDeploymentData()
	ListDeploymentsMocked(t, cloudApplicationDeploymentsIn)
	ListDeploymentsFailErrMocked(t, cloudApplicationDeploymentsIn)
	ListDeploymentsFailStatusMocked(t, cloudApplicationDeploymentsIn)
	ListDeploymentsFailJSONMocked(t, cloudApplicationDeploymentsIn)
}

func TestGetDeployment(t *testing.T) {
	cloudApplicationDeploymentsIn := testdata.GetCloudApplicationDeploymentData()
	for _, cloudApplicationDeploymentIn := range cloudApplicationDeploymentsIn {
		GetDeploymentMocked(t, cloudApplicationDeploymentIn)
		GetDeploymentFailErrMocked(t, cloudApplicationDeploymentIn)
		GetDeploymentFailStatusMocked(t, cloudApplicationDeploymentIn)
		GetDeploymentFailJSONMocked(t, cloudApplicationDeploymentIn)
	}
}

func TestDeleteDeployment(t *testing.T) {
	cloudApplicationDeploymentsIn := testdata.GetCloudApplicationDeploymentData()
	for _, cloudApplicationDeploymentIn := range cloudApplicationDeploymentsIn {
		DeleteDeploymentMocked(t, cloudApplicationDeploymentIn)
		DeleteDeploymentFailErrMocked(t, cloudApplicationDeploymentIn)
		DeleteDeploymentFailStatusMocked(t, cloudApplicationDeploymentIn)
		DeleteDeploymentFailJSONMocked(t, cloudApplicationDeploymentIn)
	}
}

func TestCreateDeploymentTask(t *testing.T) {
	cloudApplicationDeploymentTasksIn := testdata.GetCloudApplicationDeploymentTaskData()
	cloudApplicationTemplatesIn := testdata.GetCloudApplicationTemplateData()
	for _, cloudApplicationDeploymentTaskIn := range cloudApplicationDeploymentTasksIn {
		CreateDeploymentTaskMocked(t, cloudApplicationTemplatesIn[0].ID, cloudApplicationDeploymentTaskIn)
		CreateDeploymentTaskFailErrMocked(t, cloudApplicationTemplatesIn[0].ID, cloudApplicationDeploymentTaskIn)
		CreateDeploymentTaskFailStatusMocked(t, cloudApplicationTemplatesIn[0].ID, cloudApplicationDeploymentTaskIn)
		CreateDeploymentTaskFailJSONMocked(t, cloudApplicationTemplatesIn[0].ID, cloudApplicationDeploymentTaskIn)
	}
}

func TestGetDeploymentTask(t *testing.T) {
	cloudApplicationDeploymentTasksIn := testdata.GetCloudApplicationDeploymentTaskData()
	cloudApplicationTemplatesIn := testdata.GetCloudApplicationTemplateData()
	for _, cloudApplicationDeploymentTaskIn := range cloudApplicationDeploymentTasksIn {
		GetDeploymentTaskMocked(t, cloudApplicationTemplatesIn[0].ID, cloudApplicationDeploymentTaskIn)
		GetDeploymentTaskFailErrMocked(t, cloudApplicationTemplatesIn[0].ID, cloudApplicationDeploymentTaskIn)
		GetDeploymentTaskFailStatusMocked(t, cloudApplicationTemplatesIn[0].ID, cloudApplicationDeploymentTaskIn)
		GetDeploymentTaskFailJSONMocked(t, cloudApplicationTemplatesIn[0].ID, cloudApplicationDeploymentTaskIn)
	}
}
