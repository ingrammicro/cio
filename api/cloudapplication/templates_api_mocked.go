package cloudapplication

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ListTemplatesMocked test mocked function
func ListTemplatesMocked(t *testing.T, cloudApplicationTemplatesIn []*types.CloudApplicationTemplate) []*types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplatesIn)
	assert.Nil(err, "CloudApplicationTemplates test data corrupted")

	// call service
	cs.On("Get", "/plugins/tosca/cats").Return(dIn, 200, nil)
	cloudApplicationTemplatesOut, err := ds.ListTemplates()

	assert.Nil(err, "Error getting cloud application templates")
	assert.Equal(cloudApplicationTemplatesIn, cloudApplicationTemplatesOut, "ListTemplates returned different cloud application templates")

	return cloudApplicationTemplatesOut
}

// ListTemplatesFailErrMocked test mocked function
func ListTemplatesFailErrMocked(t *testing.T, cloudApplicationTemplatesIn []*types.CloudApplicationTemplate) []*types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplatesIn)
	assert.Nil(err, "CloudApplicationTemplates test data corrupted")

	// call service
	cs.On("Get", "/plugins/tosca/cats").Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudApplicationTemplatesOut, err := ds.ListTemplates()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationTemplatesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudApplicationTemplatesOut
}

// ListTemplatesFailStatusMocked test mocked function
func ListTemplatesFailStatusMocked(t *testing.T, cloudApplicationTemplatesIn []*types.CloudApplicationTemplate) []*types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplatesIn)
	assert.Nil(err, "CloudApplicationTemplates test data corrupted")

	// call service
	cs.On("Get", "/plugins/tosca/cats").Return(dIn, 499, nil)
	cloudApplicationTemplatesOut, err := ds.ListTemplates()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationTemplatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudApplicationTemplatesOut
}

// ListTemplatesFailJSONMocked test mocked function
func ListTemplatesFailJSONMocked(t *testing.T, cloudApplicationTemplatesIn []*types.CloudApplicationTemplate) []*types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/plugins/tosca/cats").Return(dIn, 200, nil)
	cloudApplicationTemplatesOut, err := ds.ListTemplates()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationTemplatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationTemplatesOut
}

// GetTemplateMocked test mocked function
func GetTemplateMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/cats/%s", cloudApplicationTemplateIn.ID)).Return(dIn, 200, nil)
	cloudApplicationTemplateOut, err := ds.GetTemplate(cloudApplicationTemplateIn.ID)

	assert.Nil(err, "Error getting cloud application template")
	assert.Equal(*cloudApplicationTemplateIn, *cloudApplicationTemplateOut, "GetTemplate returned different cloud application template")

	return cloudApplicationTemplateOut
}

// GetTemplateFailErrMocked test mocked function
func GetTemplateFailErrMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/cats/%s", cloudApplicationTemplateIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudApplicationTemplateOut, err := ds.GetTemplate(cloudApplicationTemplateIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudApplicationTemplateOut
}

// GetTemplateFailStatusMocked test mocked function
func GetTemplateFailStatusMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/cats/%s", cloudApplicationTemplateIn.ID)).Return(dIn, 499, nil)
	cloudApplicationTemplateOut, err := ds.GetTemplate(cloudApplicationTemplateIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudApplicationTemplateOut
}

// GetTemplateFailJSONMocked test mocked function
func GetTemplateFailJSONMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/plugins/tosca/cats/%s", cloudApplicationTemplateIn.ID)).Return(dIn, 200, nil)
	cloudApplicationTemplateOut, err := ds.GetTemplate(cloudApplicationTemplateIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationTemplateOut
}

// CreateTemplateMocked test mocked function
func CreateTemplateMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Post", "/plugins/tosca/cats", mapIn).Return(dOut, 200, nil)
	cloudApplicationTemplateOut, err := ds.CreateTemplate(mapIn)

	assert.Nil(err, "Error creating cloud application template")
	assert.Equal(cloudApplicationTemplateIn, cloudApplicationTemplateOut, "CreateTemplate returned different cloud application template")

	return cloudApplicationTemplateOut
}

// CreateTemplateFailErrMocked test mocked function
func CreateTemplateFailErrMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Post", "/plugins/tosca/cats", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudApplicationTemplateOut, err := ds.CreateTemplate(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudApplicationTemplateOut
}

// CreateTemplateFailStatusMocked test mocked function
func CreateTemplateFailStatusMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Post", "/plugins/tosca/cats", mapIn).Return(dOut, 499, nil)
	cloudApplicationTemplateOut, err := ds.CreateTemplate(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudApplicationTemplateOut
}

// CreateTemplateFailJSONMocked test mocked function
func CreateTemplateFailJSONMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/plugins/tosca/cats", mapIn).Return(dIn, 200, nil)
	cloudApplicationTemplateOut, err := ds.CreateTemplate(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationTemplateOut
}

// UpdateTemplateMocked test mocked function
func UpdateTemplateMocked(t *testing.T, cbIn *types.CloudApplicationTemplate) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := cbIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 200, nil)
	err = ds.UploadTemplate(sourceFilePath, targetURL)
	assert.Nil(err, "Error uploading cloud application template file")
}

// UpdateTemplateFailErrMocked test mocked function
func UpdateTemplateFailErrMocked(t *testing.T, cbIn *types.CloudApplicationTemplate) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := cbIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 403, nil)
	err = ds.UploadTemplate(sourceFilePath, targetURL)
	assert.NotNil(err, "We are expecting an error")
}

// UpdateTemplateFailStatusMocked test mocked function
func UpdateTemplateFailStatusMocked(t *testing.T, cbIn *types.CloudApplicationTemplate) {
	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	sourceFilePath := "fakeURLToFile"
	targetURL := cbIn.UploadURL

	// call service
	var noBytes []uint8
	cs.On("PutFile", sourceFilePath, targetURL).Return(noBytes, 403, fmt.Errorf("mocked error"))
	err = ds.UploadTemplate(sourceFilePath, targetURL)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// ParseMetadataTemplateMocked test mocked function
func ParseMetadataTemplateMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	mapIn := map[string]interface{}{}

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/plugins/tosca/cats/%s/parse_metadata", cloudApplicationTemplateIn.ID), &mapIn).Return(dIn, 200, nil)
	cloudApplicationTemplateOut, err := ds.ParseMetadataTemplate(cloudApplicationTemplateIn.ID)

	assert.Nil(err, "Error parsing cloud application template metadata")
	assert.Equal(*cloudApplicationTemplateIn, *cloudApplicationTemplateOut, "ParseMetadataTemplate returned different cloud application template")

	return cloudApplicationTemplateOut
}

// ParseMetadataTemplateFailErrMocked test mocked function
func ParseMetadataTemplateFailErrMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	mapIn := map[string]interface{}{}

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/plugins/tosca/cats/%s/parse_metadata", cloudApplicationTemplateIn.ID), &mapIn).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudApplicationTemplateOut, err := ds.ParseMetadataTemplate(cloudApplicationTemplateIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudApplicationTemplateOut
}

// ParseMetadataTemplateFailStatusMocked test mocked function
func ParseMetadataTemplateFailStatusMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	mapIn := map[string]interface{}{}

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/plugins/tosca/cats/%s/parse_metadata", cloudApplicationTemplateIn.ID), &mapIn).Return(dIn, 499, nil)
	cloudApplicationTemplateOut, err := ds.ParseMetadataTemplate(cloudApplicationTemplateIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudApplicationTemplateOut
}

// ParseMetadataTemplateFailJSONMocked test mocked function
func ParseMetadataTemplateFailJSONMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) *types.CloudApplicationTemplate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	mapIn := map[string]interface{}{}

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/plugins/tosca/cats/%s/parse_metadata", cloudApplicationTemplateIn.ID), &mapIn).Return(dIn, 200, nil)
	cloudApplicationTemplateOut, err := ds.ParseMetadataTemplate(cloudApplicationTemplateIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudApplicationTemplateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudApplicationTemplateOut
}

// DeleteTemplateMocked test mocked function
func DeleteTemplateMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/plugins/tosca/cats/%s", cloudApplicationTemplateIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteTemplate(cloudApplicationTemplateIn.ID)
	assert.Nil(err, "Error deleting cloud application template")
}

// DeleteTemplateFailErrMocked test mocked function
func DeleteTemplateFailErrMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/plugins/tosca/cats/%s", cloudApplicationTemplateIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteTemplate(cloudApplicationTemplateIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteTemplateFailStatusMocked test mocked function
func DeleteTemplateFailStatusMocked(t *testing.T, cloudApplicationTemplateIn *types.CloudApplicationTemplate) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudApplicationTemplateService(cs)
	assert.Nil(err, "Couldn't load cloudApplicationTemplate service")
	assert.NotNil(ds, "CloudApplicationTemplate service not instanced")

	// to json
	dIn, err := json.Marshal(cloudApplicationTemplateIn)
	assert.Nil(err, "CloudApplicationTemplate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/plugins/tosca/cats/%s", cloudApplicationTemplateIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteTemplate(cloudApplicationTemplateIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
