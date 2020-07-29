package cloudspecificextension

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCloudSpecificExtensionTemplateServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudSpecificExtensionTemplateService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestLisTemplates(t *testing.T) {
	cloudSpecificExtensionTemplatesIn := testdata.GetCloudSpecificExtensionTemplateData()
	ListTemplatesMocked(t, cloudSpecificExtensionTemplatesIn)
	ListTemplatesFailErrMocked(t, cloudSpecificExtensionTemplatesIn)
	ListTemplatesFailStatusMocked(t, cloudSpecificExtensionTemplatesIn)
	ListTemplatesFailJSONMocked(t, cloudSpecificExtensionTemplatesIn)
}

func TestGeTemplate(t *testing.T) {
	cloudSpecificExtensionTemplatesIn := testdata.GetCloudSpecificExtensionTemplateData()
	for _, cloudSpecificExtensionTemplateIn := range cloudSpecificExtensionTemplatesIn {
		GetTemplateMocked(t, cloudSpecificExtensionTemplateIn)
		GetTemplateFailErrMocked(t, cloudSpecificExtensionTemplateIn)
		GetTemplateFailStatusMocked(t, cloudSpecificExtensionTemplateIn)
		GetTemplateFailJSONMocked(t, cloudSpecificExtensionTemplateIn)
	}
}

func TestCreateTemplate(t *testing.T) {
	cloudSpecificExtensionTemplatesIn := testdata.GetCloudSpecificExtensionTemplateData()
	for _, cloudSpecificExtensionTemplateIn := range cloudSpecificExtensionTemplatesIn {
		CreateTemplateMocked(t, cloudSpecificExtensionTemplateIn)
		CreateTemplateFailErrMocked(t, cloudSpecificExtensionTemplateIn)
		CreateTemplateFailStatusMocked(t, cloudSpecificExtensionTemplateIn)
		CreateTemplateFailJSONMocked(t, cloudSpecificExtensionTemplateIn)
	}
}

func TestUpdateTemplate(t *testing.T) {
	cloudSpecificExtensionTemplatesIn := testdata.GetCloudSpecificExtensionTemplateData()
	for _, cloudSpecificExtensionTemplateIn := range cloudSpecificExtensionTemplatesIn {
		UpdateTemplateMocked(t, cloudSpecificExtensionTemplateIn)
		UpdateTemplateFailErrMocked(t, cloudSpecificExtensionTemplateIn)
		UpdateTemplateFailStatusMocked(t, cloudSpecificExtensionTemplateIn)
		UpdateTemplateFailJSONMocked(t, cloudSpecificExtensionTemplateIn)
	}
}

func TestListTemplateDeployments(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	cloudSpecificExtensionTemplatesIn := testdata.GetCloudSpecificExtensionDeploymentData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ListTemplateDeploymentsMocked(t, cloudAccountIn.ID, cloudSpecificExtensionTemplatesIn)
		ListTemplateDeploymentsFailErrMocked(t, cloudAccountIn.ID, cloudSpecificExtensionTemplatesIn)
		ListTemplateDeploymentsFailStatusMocked(t, cloudAccountIn.ID, cloudSpecificExtensionTemplatesIn)
		ListTemplateDeploymentsFailJSONMocked(t, cloudAccountIn.ID, cloudSpecificExtensionTemplatesIn)
	}
}

func TestDeleteTemplate(t *testing.T) {
	cloudSpecificExtensionTemplatesIn := testdata.GetCloudSpecificExtensionTemplateData()
	for _, cloudSpecificExtensionTemplateIn := range cloudSpecificExtensionTemplatesIn {
		DeleteTemplateMocked(t, cloudSpecificExtensionTemplateIn)
		DeleteTemplateFailErrMocked(t, cloudSpecificExtensionTemplateIn)
		DeleteTemplateFailStatusMocked(t, cloudSpecificExtensionTemplateIn)
	}
}
