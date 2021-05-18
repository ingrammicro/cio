// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudspecificextension

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewCloudSpecificExtensionDeploymentServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudSpecificExtensionDeploymentService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestLisDeployments(t *testing.T) {
	cloudSpecificExtensionDeploymentsIn := testdata.GetCloudSpecificExtensionDeploymentData()
	ListDeploymentsMocked(t, cloudSpecificExtensionDeploymentsIn)
	ListDeploymentsFailErrMocked(t, cloudSpecificExtensionDeploymentsIn)
	ListDeploymentsFailStatusMocked(t, cloudSpecificExtensionDeploymentsIn)
	ListDeploymentsFailJSONMocked(t, cloudSpecificExtensionDeploymentsIn)
}

func TestGeDeployment(t *testing.T) {
	cloudSpecificExtensionDeploymentsIn := testdata.GetCloudSpecificExtensionDeploymentData()
	for _, cloudSpecificExtensionDeploymentIn := range cloudSpecificExtensionDeploymentsIn {
		GetDeploymentMocked(t, cloudSpecificExtensionDeploymentIn)
		GetDeploymentFailErrMocked(t, cloudSpecificExtensionDeploymentIn)
		GetDeploymentFailStatusMocked(t, cloudSpecificExtensionDeploymentIn)
		GetDeploymentFailJSONMocked(t, cloudSpecificExtensionDeploymentIn)
	}
}

func TestCreateDeployment(t *testing.T) {
	cloudSpecificExtensionDeploymentsIn := testdata.GetCloudSpecificExtensionDeploymentData()
	cloudSpecificExtensionTemplatesIn := testdata.GetCloudSpecificExtensionTemplateData()
	for _, cloudSpecificExtensionDeploymentIn := range cloudSpecificExtensionDeploymentsIn {
		CreateDeploymentMocked(t, cloudSpecificExtensionTemplatesIn[0].ID, cloudSpecificExtensionDeploymentIn)
		CreateDeploymentFailErrMocked(t, cloudSpecificExtensionTemplatesIn[0].ID, cloudSpecificExtensionDeploymentIn)
		CreateDeploymentFailStatusMocked(t, cloudSpecificExtensionTemplatesIn[0].ID, cloudSpecificExtensionDeploymentIn)
		CreateDeploymentFailJSONMocked(t, cloudSpecificExtensionTemplatesIn[0].ID, cloudSpecificExtensionDeploymentIn)
	}
}

func TestUpdateDeployment(t *testing.T) {
	cloudSpecificExtensionDeploymentsIn := testdata.GetCloudSpecificExtensionDeploymentData()
	for _, cloudSpecificExtensionDeploymentIn := range cloudSpecificExtensionDeploymentsIn {
		UpdateDeploymentMocked(t, cloudSpecificExtensionDeploymentIn)
		UpdateDeploymentFailErrMocked(t, cloudSpecificExtensionDeploymentIn)
		UpdateDeploymentFailStatusMocked(t, cloudSpecificExtensionDeploymentIn)
		UpdateDeploymentFailJSONMocked(t, cloudSpecificExtensionDeploymentIn)
	}
}

func TestDeleteDeployment(t *testing.T) {
	cloudSpecificExtensionDeploymentsIn := testdata.GetCloudSpecificExtensionDeploymentData()
	for _, cloudSpecificExtensionDeploymentIn := range cloudSpecificExtensionDeploymentsIn {
		DeleteDeploymentMocked(t, cloudSpecificExtensionDeploymentIn)
		DeleteDeploymentFailErrMocked(t, cloudSpecificExtensionDeploymentIn)
		DeleteDeploymentFailStatusMocked(t, cloudSpecificExtensionDeploymentIn)
		DeleteDeploymentFailJSONMocked(t, cloudSpecificExtensionDeploymentIn)
	}
}
