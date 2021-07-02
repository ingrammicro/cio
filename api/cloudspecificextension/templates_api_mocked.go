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

// ListTemplatesMocked test mocked function
func ListTemplatesMocked(
	t *testing.T,
	cloudSpecificExtensionTemplatesIn []*types.CloudSpecificExtensionTemplate,
) []*types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplatesIn)
	assert.Nil(err, "CloudSpecificExtensionTemplates test data corrupted")

	// call service
	cs.On("Get", APIPathCseTemplates).Return(dIn, 200, nil)
	cloudSpecificExtensionTemplatesOut, err := ds.ListTemplates()

	assert.Nil(err, "Error getting cloud specific extension templates")
	assert.Equal(
		cloudSpecificExtensionTemplatesIn,
		cloudSpecificExtensionTemplatesOut,
		"ListTemplates returned different cloud specific extension templates",
	)

	return cloudSpecificExtensionTemplatesOut
}

// ListTemplatesFailErrMocked test mocked function
func ListTemplatesFailErrMocked(
	t *testing.T,
	cloudSpecificExtensionTemplatesIn []*types.CloudSpecificExtensionTemplate,
) []*types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplatesIn)
	assert.Nil(err, "CloudSpecificExtensionTemplates test data corrupted")

	// call service
	cs.On("Get", APIPathCseTemplates).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionTemplatesOut, err := ds.ListTemplates()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionTemplatesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionTemplatesOut
}

// ListTemplatesFailStatusMocked test mocked function
func ListTemplatesFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionTemplatesIn []*types.CloudSpecificExtensionTemplate,
) []*types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplatesIn)
	assert.Nil(err, "CloudSpecificExtensionTemplates test data corrupted")

	// call service
	cs.On("Get", APIPathCseTemplates).Return(dIn, 499, nil)
	cloudSpecificExtensionTemplatesOut, err := ds.ListTemplates()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionTemplatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionTemplatesOut
}

// ListTemplatesFailJSONMocked test mocked function
func ListTemplatesFailJSONMocked(
	t *testing.T,
	cloudSpecificExtensionTemplatesIn []*types.CloudSpecificExtensionTemplate,
) []*types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathCseTemplates).Return(dIn, 200, nil)
	cloudSpecificExtensionTemplatesOut, err := ds.ListTemplates()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionTemplatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionTemplatesOut
}

// GetTemplateMocked test mocked function
func GetTemplateMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID)).Return(dIn, 200, nil)
	cloudSpecificExtensionTemplateOut, err := ds.GetTemplate(cloudSpecificExtensionTemplateIn.ID)

	assert.Nil(err, "Error getting cloud specific extension template")
	assert.Equal(
		*cloudSpecificExtensionTemplateIn,
		*cloudSpecificExtensionTemplateOut,
		"GetTemplate returned different cloud specific extension template",
	)

	return cloudSpecificExtensionTemplateOut
}

// GetTemplateFailErrMocked test mocked function
func GetTemplateFailErrMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionTemplateOut, err := ds.GetTemplate(cloudSpecificExtensionTemplateIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionTemplateOut
}

// GetTemplateFailStatusMocked test mocked function
func GetTemplateFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID)).Return(dIn, 499, nil)
	cloudSpecificExtensionTemplateOut, err := ds.GetTemplate(cloudSpecificExtensionTemplateIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionTemplateOut
}

// GetTemplateFailJSONMocked test mocked function
func GetTemplateFailJSONMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID)).Return(dIn, 200, nil)
	cloudSpecificExtensionTemplateOut, err := ds.GetTemplate(cloudSpecificExtensionTemplateIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionTemplateOut
}

// CreateTemplateMocked test mocked function
func CreateTemplateMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Post", APIPathCseTemplates, mapIn).Return(dOut, 200, nil)
	cloudSpecificExtensionTemplateOut, err := ds.CreateTemplate(mapIn)

	assert.Nil(err, "Error creating cloud specific extension template")
	assert.Equal(
		cloudSpecificExtensionTemplateIn,
		cloudSpecificExtensionTemplateOut,
		"CreateTemplate returned different cloud specific extension template",
	)

	return cloudSpecificExtensionTemplateOut
}

// CreateTemplateFailErrMocked test mocked function
func CreateTemplateFailErrMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Post", APIPathCseTemplates, mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionTemplateOut, err := ds.CreateTemplate(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionTemplateOut
}

// CreateTemplateFailStatusMocked test mocked function
func CreateTemplateFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Post", APIPathCseTemplates, mapIn).Return(dOut, 499, nil)
	cloudSpecificExtensionTemplateOut, err := ds.CreateTemplate(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionTemplateOut
}

// CreateTemplateFailJSONMocked test mocked function
func CreateTemplateFailJSONMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", APIPathCseTemplates, mapIn).Return(dIn, 200, nil)
	cloudSpecificExtensionTemplateOut, err := ds.CreateTemplate(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionTemplateOut
}

// UpdateTemplateMocked test mocked function
func UpdateTemplateMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID), mapIn).Return(dOut, 200, nil)
	cloudSpecificExtensionTemplateOut, err := ds.UpdateTemplate(cloudSpecificExtensionTemplateIn.ID, mapIn)

	assert.Nil(err, "Error updating cloud specific extension template")
	assert.Equal(
		cloudSpecificExtensionTemplateIn,
		cloudSpecificExtensionTemplateOut,
		"UpdateTemplate returned different cloud specific extension template",
	)

	return cloudSpecificExtensionTemplateOut
}

// UpdateTemplateFailErrMocked test mocked function
func UpdateTemplateFailErrMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionTemplateOut, err := ds.UpdateTemplate(cloudSpecificExtensionTemplateIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionTemplateOut
}

// UpdateTemplateFailStatusMocked test mocked function
func UpdateTemplateFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID), mapIn).Return(dOut, 499, nil)
	cloudSpecificExtensionTemplateOut, err := ds.UpdateTemplate(cloudSpecificExtensionTemplateIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionTemplateOut
}

// UpdateTemplateFailJSONMocked test mocked function
func UpdateTemplateFailJSONMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) *types.CloudSpecificExtensionTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID), mapIn).Return(dIn, 200, nil)
	cloudSpecificExtensionTemplateOut, err := ds.UpdateTemplate(cloudSpecificExtensionTemplateIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionTemplateOut
}

// ListTemplateDeploymentsMocked test mocked function
func ListTemplateDeploymentsMocked(
	t *testing.T,
	cloudAccountID string,
	cloudSpecificExtensionDeploymentsIn []*types.CloudSpecificExtensionDeployment,
) []*types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentsIn)
	assert.Nil(err, "CloudSpecificExtensionDeployments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseTemplateDeployments, cloudAccountID)).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentsOut, err := ds.ListDeployments(cloudAccountID)

	assert.Nil(err, "Error getting cloud specific extension deployments")
	assert.Equal(
		cloudSpecificExtensionDeploymentsIn,
		cloudSpecificExtensionDeploymentsOut,
		"ListDeployments returned different cloud specific extension deployments",
	)

	return cloudSpecificExtensionDeploymentsOut
}

// ListTemplateDeploymentsFailErrMocked test mocked function
func ListTemplateDeploymentsFailErrMocked(
	t *testing.T,
	cloudAccountID string,
	cloudSpecificExtensionDeploymentsIn []*types.CloudSpecificExtensionDeployment,
) []*types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentsIn)
	assert.Nil(err, "CloudSpecificExtensionDeployments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseTemplateDeployments, cloudAccountID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudSpecificExtensionDeploymentsOut, err := ds.ListDeployments(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudSpecificExtensionDeploymentsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudSpecificExtensionDeploymentsOut
}

// ListTemplateDeploymentsFailStatusMocked test mocked function
func ListTemplateDeploymentsFailStatusMocked(
	t *testing.T,
	cloudAccountID string,
	cloudSpecificExtensionDeploymentsIn []*types.CloudSpecificExtensionDeployment,
) []*types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionDeploymentsIn)
	assert.Nil(err, "CloudSpecificExtensionDeployments test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseTemplateDeployments, cloudAccountID)).Return(dIn, 499, nil)
	cloudSpecificExtensionDeploymentsOut, err := ds.ListDeployments(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudSpecificExtensionDeploymentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudSpecificExtensionDeploymentsOut
}

// ListTemplateDeploymentsFailJSONMocked test mocked function
func ListTemplateDeploymentsFailJSONMocked(
	t *testing.T,
	cloudAccountID string,
	cloudSpecificExtensionDeploymentsIn []*types.CloudSpecificExtensionDeployment,
) []*types.CloudSpecificExtensionDeployment {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathCseTemplateDeployments, cloudAccountID)).Return(dIn, 200, nil)
	cloudSpecificExtensionDeploymentsOut, err := ds.ListDeployments(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudSpecificExtensionDeploymentsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudSpecificExtensionDeploymentsOut
}

// DeleteTemplateMocked test mocked function
func DeleteTemplateMocked(t *testing.T, cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteTemplate(cloudSpecificExtensionTemplateIn.ID)

	assert.Nil(err, "Error deleting cloud specific extension template")
}

// DeleteTemplateFailErrMocked test mocked function
func DeleteTemplateFailErrMocked(t *testing.T, cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteTemplate(cloudSpecificExtensionTemplateIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteTemplateFailStatusMocked test mocked function
func DeleteTemplateFailStatusMocked(
	t *testing.T,
	cloudSpecificExtensionTemplateIn *types.CloudSpecificExtensionTemplate,
) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudSpecificExtensionTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudSpecificExtensionTemplate service")
	assert.NotNil(ds, "CloudSpecificExtensionTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudSpecificExtensionTemplateIn)
	assert.Nil(err, "CloudSpecificExtensionTemplate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathCseTemplate, cloudSpecificExtensionTemplateIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteTemplate(cloudSpecificExtensionTemplateIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
