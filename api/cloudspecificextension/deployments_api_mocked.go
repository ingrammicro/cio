// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudspecificextension

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// ListDeploymentsMocked test mocked function
func ListDeploymentsMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentsIn []*types.CloudSpecificExtensionDeployment,
) []*types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentsIn)
	assert.Nil(err, "CloudSpecificExtensionDeployments test data corrupted")

	// call service
	cs.On("Get", APIPathCseDeployments).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentsOut, err := ds.ListDeployments()

	assert.Nil(err, "Error getting cloud specific extension deployments")
	assert.Equal(
		cloudSpecificExtensionDeploymentsIn,
		cloudSpecificExtensionDeploymentsOut,
		"ListDeployments returned different cloud specific extension deployments",
	)

	return cloudSpecificExtensionDeploymentsOut
}

// ListDeploymentsFailErrMocked test mocked function
func ListDeploymentsFailErrMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentsIn []*types.CloudSpecificExtensionDeployment,
) []*types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentsIn)
	assert.Nil(err, "CloudSpecificExtensionDeployments test data corrupted")

	// call service
	cs.On("Get", APIPathCseDeployments).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionDeploymentsOut, err := ds.ListDeployments()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionDeploymentsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionDeploymentsOut
}

// ListDeploymentsFailStatusMocked test mocked function
func ListDeploymentsFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentsIn []*types.CloudSpecificExtensionDeployment,
) []*types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentsIn)
	assert.Nil(err, "CloudSpecificExtensionDeployments test data corrupted")

	// call service
	cs.On("Get", APIPathCseDeployments).Return(dIn, 499, nil)
	cloudSpecificExtensionDeploymentsOut, err := ds.ListDeployments()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionDeploymentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionDeploymentsOut
}

// ListDeploymentsFailJSONMocked test mocked function
func ListDeploymentsFailJSONMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentsIn []*types.CloudSpecificExtensionDeployment,
) []*types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathCseDeployments).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentsOut, err := ds.ListDeployments()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionDeploymentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionDeploymentsOut
}

// GetDeploymentMocked test mocked function
func GetDeploymentMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID)).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.GetDeployment(cloudSpecificExtensionDeploymentIn.ID)

	assert.Nil(err, "Error getting cloud specific extension deployment")
	assert.Equal(
		*cloudSpecificExtensionDeploymentIn,
		*cloudSpecificExtensionDeploymentOut,
		"GetDeployment returned different cloud specific extension deployment",
	)

	return cloudSpecificExtensionDeploymentOut
}

// GetDeploymentFailErrMocked test mocked function
func GetDeploymentFailErrMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionDeploymentOut, err := ds.GetDeployment(cloudSpecificExtensionDeploymentIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionDeploymentOut
}

// GetDeploymentFailStatusMocked test mocked function
func GetDeploymentFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID)).Return(dIn, 499, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.GetDeployment(cloudSpecificExtensionDeploymentIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionDeploymentOut
}

// GetDeploymentFailJSONMocked test mocked function
func GetDeploymentFailJSONMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID)).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.GetDeployment(cloudSpecificExtensionDeploymentIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionDeploymentOut
}

// CreateDeploymentMocked test mocked function
func CreateDeploymentMocked(
	t *testing.T,
	templateID string,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathCseTemplateDeployments, templateID), mapIn).Return(dOut, 200, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.CreateDeployment(templateID, mapIn)

	assert.Nil(err, "Error creating cloud specific extension deployment")
	assert.Equal(
		cloudSpecificExtensionDeploymentIn,
		cloudSpecificExtensionDeploymentOut,
		"CreateDeployment returned different cloud specific extension deployment",
	)

	return cloudSpecificExtensionDeploymentOut
}

// CreateDeploymentFailErrMocked test mocked function
func CreateDeploymentFailErrMocked(
	t *testing.T,
	templateID string,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathCseTemplateDeployments, templateID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionDeploymentOut, err := ds.CreateDeployment(templateID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionDeploymentOut
}

// CreateDeploymentFailStatusMocked test mocked function
func CreateDeploymentFailStatusMocked(
	t *testing.T,
	templateID string,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathCseTemplateDeployments, templateID), mapIn).Return(dOut, 499, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.CreateDeployment(templateID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionDeploymentOut
}

// CreateDeploymentFailJSONMocked test mocked function
func CreateDeploymentFailJSONMocked(
	t *testing.T,
	templateID string,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathCseTemplateDeployments, templateID), mapIn).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.CreateDeployment(templateID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionDeploymentOut
}

// UpdateDeploymentMocked test mocked function
func UpdateDeploymentMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID), mapIn).
		Return(dOut, 200, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.UpdateDeployment(cloudSpecificExtensionDeploymentIn.ID, mapIn)

	assert.Nil(err, "Error updating cloud specific extension deployment")
	assert.Equal(
		cloudSpecificExtensionDeploymentIn,
		cloudSpecificExtensionDeploymentOut,
		"UpdateDeployment returned different cloud specific extension deployment",
	)

	return cloudSpecificExtensionDeploymentOut
}

// UpdateDeploymentFailErrMocked test mocked function
func UpdateDeploymentFailErrMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionDeploymentOut, err := ds.UpdateDeployment(cloudSpecificExtensionDeploymentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionDeploymentOut
}

// UpdateDeploymentFailStatusMocked test mocked function
func UpdateDeploymentFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID), mapIn).
		Return(dOut, 499, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.UpdateDeployment(cloudSpecificExtensionDeploymentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionDeploymentOut
}

// UpdateDeploymentFailJSONMocked test mocked function
func UpdateDeploymentFailJSONMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID), mapIn).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.UpdateDeployment(cloudSpecificExtensionDeploymentIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionDeploymentOut
}

// DeleteDeploymentMocked test mocked function
func DeleteDeploymentMocked(t *testing.T, cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID)).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.DeleteDeployment(cloudSpecificExtensionDeploymentIn.ID)

	assert.Nil(err, "Error deleting cloud specific extension template")
	assert.Equal(
		cloudSpecificExtensionDeploymentIn,
		cloudSpecificExtensionDeploymentOut,
		"DeleteDeployment returned different cloud specific extension deployment",
	)
}

// DeleteDeploymentFailErrMocked test mocked function
func DeleteDeploymentFailErrMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionDeploymentOut, err := ds.DeleteDeployment(cloudSpecificExtensionDeploymentIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteDeploymentFailStatusMocked test mocked function
func DeleteDeploymentFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentIn)
	assert.Nil(err, "CloudSpecificExtensionDeployment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID)).Return(dIn, 499, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.DeleteDeployment(cloudSpecificExtensionDeploymentIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteDeploymentFailJSONMocked test mocked function
func DeleteDeploymentFailJSONMocked(
	t *testing.T,
	cloudSpecificExtensionDeploymentIn *types.CloudSpecificExtensionDeployment,
) *types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionDeployment service")
	assert.NotNil(ds, "CloudSpecificExtensionDeployment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathCseDeployment, cloudSpecificExtensionDeploymentIn.ID)).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentOut, err := ds.DeleteDeployment(cloudSpecificExtensionDeploymentIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionDeploymentOut
}
