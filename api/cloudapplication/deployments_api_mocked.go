package cloudapplication

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ListDeploymentsMocked test mocked function
func ListDeploymentsMocked(t *testing.T, cloudApplicationDeploymentsIn []*types.CloudApplicationDeployment) []*types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentsIn)
	assert.Nil(err, "CloudApplicationDeployments test data corrupted")

	// call service
	cs.On("Get", "/labels").Return(dIn, 200, nil)
	cloudApplicationDeploymentsOut, err := ds.ListDeployments()

	assert.Nil(err, "Error getting cloud application deployments")
	assert.Equal(cloudApplicationDeploymentsIn, cloudApplicationDeploymentsOut, "ListDeployments returned different cloud application deployments")

	return cloudApplicationDeploymentsOut
}

// ListDeploymentsFailErrMocked test mocked function
func ListDeploymentsFailErrMocked(t *testing.T, cloudApplicationDeploymentsIn []*types.CloudApplicationDeployment) []*types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentsIn)
	assert.Nil(err, "CloudApplicationDeployments test data corrupted")

	// call service
	cs.On("Get", "/labels").Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudApplicationDeploymentsOut, err := ds.ListDeployments()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationDeploymentsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudApplicationDeploymentsOut
}

// ListDeploymentsFailStatusMocked test mocked function
func ListDeploymentsFailStatusMocked(t *testing.T, cloudApplicationDeploymentsIn []*types.CloudApplicationDeployment) []*types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentsIn)
	assert.Nil(err, "CloudApplicationDeployments test data corrupted")

	// call service
	cs.On("Get", "/labels").Return(dIn, 499, nil)
	cloudApplicationDeploymentsOut, err := ds.ListDeployments()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationDeploymentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudApplicationDeploymentsOut
}

// ListDeploymentsFailJSONMocked test mocked function
func ListDeploymentsFailJSONMocked(t *testing.T, cloudApplicationDeploymentsIn []*types.CloudApplicationDeployment) []*types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/labels").Return(dIn, 200, nil)
	cloudApplicationDeploymentsOut, err := ds.ListDeployments()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationDeploymentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationDeploymentsOut
}

// GetDeploymentMocked test mocked function
func GetDeploymentMocked(t *testing.T, cloudApplicationDeploymentIn *types.CloudApplicationDeployment) *types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentIn)
	assert.Nil(err, "CloudApplicationDeployment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/deployments/%s", cloudApplicationDeploymentIn.ID)).Return(dIn, 200, nil)
	cloudApplicationDeploymentOut, _, err := ds.GetDeployment(cloudApplicationDeploymentIn.ID)

	assert.Nil(err, "Error getting cloud application deployment")
	assert.Equal(*cloudApplicationDeploymentIn, *cloudApplicationDeploymentOut, "GetDeployment returned different cloud application deployment")

	return cloudApplicationDeploymentOut
}

// GetDeploymentFailErrMocked test mocked function
func GetDeploymentFailErrMocked(t *testing.T, cloudApplicationDeploymentIn *types.CloudApplicationDeployment) *types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentIn)
	assert.Nil(err, "CloudApplicationDeployment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/deployments/%s", cloudApplicationDeploymentIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudApplicationDeploymentOut, _, err := ds.GetDeployment(cloudApplicationDeploymentIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationDeploymentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudApplicationDeploymentOut
}

// GetDeploymentFailStatusMocked test mocked function
func GetDeploymentFailStatusMocked(t *testing.T, cloudApplicationDeploymentIn *types.CloudApplicationDeployment) *types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentIn)
	assert.Nil(err, "CloudApplicationDeployment test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/deployments/%s", cloudApplicationDeploymentIn.ID)).Return(dIn, 499, nil)
	cloudApplicationDeploymentOut, _, err := ds.GetDeployment(cloudApplicationDeploymentIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudApplicationDeploymentOut
}

// GetDeploymentFailJSONMocked test mocked function
func GetDeploymentFailJSONMocked(t *testing.T, cloudApplicationDeploymentIn *types.CloudApplicationDeployment) *types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/deployments/%s", cloudApplicationDeploymentIn.ID)).Return(dIn, 200, nil)
	cloudApplicationDeploymentOut, _, err := ds.GetDeployment(cloudApplicationDeploymentIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationDeploymentOut
}

// DeleteDeploymentMocked test mocked function
func DeleteDeploymentMocked(t *testing.T, cloudApplicationDeploymentIn *types.CloudApplicationDeployment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentIn)
	assert.Nil(err, "CloudApplicationDeployment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/plugins/tosca/deployments/%s", cloudApplicationDeploymentIn.ID)).Return(dIn, 200, nil)
	cloudApplicationDeploymentOut, err := ds.DeleteDeployment(cloudApplicationDeploymentIn.ID)
	assert.Nil(err, "Error deleting cloud application deployment")
	assert.Equal(*cloudApplicationDeploymentIn, *cloudApplicationDeploymentOut, "DeleteDeployment returned different cloud application deployment")
}

// DeleteDeploymentFailErrMocked test mocked function
func DeleteDeploymentFailErrMocked(t *testing.T, cloudApplicationDeploymentIn *types.CloudApplicationDeployment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentIn)
	assert.Nil(err, "CloudApplicationDeployment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/plugins/tosca/deployments/%s", cloudApplicationDeploymentIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudApplicationDeploymentOut, err := ds.DeleteDeployment(cloudApplicationDeploymentIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationDeploymentOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteDeploymentFailStatusMocked test mocked function
func DeleteDeploymentFailStatusMocked(t *testing.T, cloudApplicationDeploymentIn *types.CloudApplicationDeployment) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentIn)
	assert.Nil(err, "CloudApplicationDeployment test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/plugins/tosca/deployments/%s", cloudApplicationDeploymentIn.ID)).Return(dIn, 499, nil)
	cloudApplicationDeploymentOut, err := ds.DeleteDeployment(cloudApplicationDeploymentIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteDeploymentFailJSONMocked test mocked function
func DeleteDeploymentFailJSONMocked(t *testing.T, cloudApplicationDeploymentIn *types.CloudApplicationDeployment) *types.CloudApplicationDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf("/plugins/tosca/deployments/%s", cloudApplicationDeploymentIn.ID)).Return(dIn, 200, nil)
	cloudApplicationDeploymentOut, err := ds.DeleteDeployment(cloudApplicationDeploymentIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationDeploymentOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationDeploymentOut
}

// CreateDeploymentTaskMocked test mocked function
func CreateDeploymentTaskMocked(t *testing.T, catID string, cloudApplicationDeploymentTaskIn *types.CloudApplicationDeploymentTask) *types.CloudApplicationDeploymentTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudApplicationDeploymentTaskIn)

	// to json
	dOut, err := json.Marshal(cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/plugins/tosca/cats/%s/deployment_tasks", catID), mapIn).Return(dOut, 200, nil)
	cloudApplicationDeploymentTaskOut, err := ds.CreateDeploymentTask(catID, mapIn)

	assert.Nil(err, "Error creating cloud application deployment task")
	assert.Equal(cloudApplicationDeploymentTaskIn, cloudApplicationDeploymentTaskOut, "CreateDeploymentTask returned different cloud application deployment task")

	return cloudApplicationDeploymentTaskOut
}

// CreateDeploymentTaskFailErrMocked test mocked function
func CreateDeploymentTaskFailErrMocked(t *testing.T, catID string, cloudApplicationDeploymentTaskIn *types.CloudApplicationDeploymentTask) *types.CloudApplicationDeploymentTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/plugins/tosca/cats/%s/deployment_tasks", catID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudApplicationDeploymentTaskOut, err := ds.CreateDeploymentTask(catID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationDeploymentTaskOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudApplicationDeploymentTaskOut
}

// CreateDeploymentTaskFailStatusMocked test mocked function
func CreateDeploymentTaskFailStatusMocked(t *testing.T, catID string, cloudApplicationDeploymentTaskIn *types.CloudApplicationDeploymentTask) *types.CloudApplicationDeploymentTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/plugins/tosca/cats/%s/deployment_tasks", catID), mapIn).Return(dOut, 499, nil)
	cloudApplicationDeploymentTaskOut, err := ds.CreateDeploymentTask(catID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationDeploymentTaskOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudApplicationDeploymentTaskOut
}

// CreateDeploymentTaskFailJSONMocked test mocked function
func CreateDeploymentTaskFailJSONMocked(t *testing.T, catID string, cloudApplicationDeploymentTaskIn *types.CloudApplicationDeploymentTask) *types.CloudApplicationDeploymentTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/plugins/tosca/cats/%s/deployment_tasks", catID), mapIn).Return(dIn, 200, nil)
	cloudApplicationDeploymentTaskOut, err := ds.CreateDeploymentTask(catID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationDeploymentTaskOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationDeploymentTaskOut
}

// GetDeploymentTaskMocked test mocked function
func GetDeploymentTaskMocked(t *testing.T, catID string, cloudApplicationDeploymentTaskIn *types.CloudApplicationDeploymentTask) *types.CloudApplicationDeploymentTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/cats/%s/deployment_tasks/%s", catID, cloudApplicationDeploymentTaskIn.ID)).Return(dIn, 200, nil)
	cloudApplicationDeploymentTaskOut, err := ds.GetDeploymentTask(catID, cloudApplicationDeploymentTaskIn.ID)

	assert.Nil(err, "Error getting cloud application deployment task")
	assert.Equal(*cloudApplicationDeploymentTaskIn, *cloudApplicationDeploymentTaskOut, "GetDeploymentTask returned different cloud application deployment task")

	return cloudApplicationDeploymentTaskOut
}

// GetDeploymentTaskFailErrMocked test mocked function
func GetDeploymentTaskFailErrMocked(t *testing.T, catID string, cloudApplicationDeploymentTaskIn *types.CloudApplicationDeploymentTask) *types.CloudApplicationDeploymentTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/cats/%s/deployment_tasks/%s", catID, cloudApplicationDeploymentTaskIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudApplicationDeploymentTaskOut, err := ds.GetDeploymentTask(catID, cloudApplicationDeploymentTaskIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationDeploymentTaskOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudApplicationDeploymentTaskOut
}

// GetDeploymentTaskFailStatusMocked test mocked function
func GetDeploymentTaskFailStatusMocked(t *testing.T, catID string, cloudApplicationDeploymentTaskIn *types.CloudApplicationDeploymentTask) *types.CloudApplicationDeploymentTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationDeploymentTaskIn)
	assert.Nil(err, "CloudApplicationDeploymentTask test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/cats/%s/deployment_tasks/%s", catID, cloudApplicationDeploymentTaskIn.ID)).Return(dIn, 499, nil)
	cloudApplicationDeploymentTaskOut, err := ds.GetDeploymentTask(catID, cloudApplicationDeploymentTaskIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationDeploymentTaskOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudApplicationDeploymentTaskOut
}

// GetDeploymentTaskFailJSONMocked test mocked function
func GetDeploymentTaskFailJSONMocked(t *testing.T, catID string, cloudApplicationDeploymentTaskIn *types.CloudApplicationDeploymentTask) *types.CloudApplicationDeploymentTask {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationDeploymentService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationDeployment service")
	assert.NotNil(ds, "CloudApplicationDeployment service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/cats/%s/deployment_tasks/%s", catID, cloudApplicationDeploymentTaskIn.ID)).Return(dIn, 200, nil)
	cloudApplicationDeploymentTaskOut, err := ds.GetDeploymentTask(catID, cloudApplicationDeploymentTaskIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationDeploymentTaskOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationDeploymentTaskOut
}
