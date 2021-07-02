// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloudapplication

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewCloudApplicationTemplateServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudApplicationTemplateService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListTemplates(t *testing.T) {
	cloudApplicationTemplatesIn := testdata.GetCloudApplicationTemplateData()
	ListTemplatesMocked(t, cloudApplicationTemplatesIn)
	ListTemplatesFailErrMocked(t, cloudApplicationTemplatesIn)
	ListTemplatesFailStatusMocked(t, cloudApplicationTemplatesIn)
	ListTemplatesFailJSONMocked(t, cloudApplicationTemplatesIn)
}

func TestGetTemplate(t *testing.T) {
	cloudApplicationTemplatesIn := testdata.GetCloudApplicationTemplateData()
	for _, cloudApplicationTemplateIn := range cloudApplicationTemplatesIn {
		GetTemplateMocked(t, cloudApplicationTemplateIn)
		GetTemplateFailErrMocked(t, cloudApplicationTemplateIn)
		GetTemplateFailStatusMocked(t, cloudApplicationTemplateIn)
		GetTemplateFailJSONMocked(t, cloudApplicationTemplateIn)
	}
}

func TestCreateTemplate(t *testing.T) {
	cloudApplicationTemplatesIn := testdata.GetCloudApplicationTemplateData()
	for _, cloudApplicationTemplateIn := range cloudApplicationTemplatesIn {
		CreateTemplateMocked(t, cloudApplicationTemplateIn)
		CreateTemplateFailErrMocked(t, cloudApplicationTemplateIn)
		CreateTemplateFailStatusMocked(t, cloudApplicationTemplateIn)
		CreateTemplateFailJSONMocked(t, cloudApplicationTemplateIn)
	}
}

func TestUploadTemplate(t *testing.T) {
	cloudApplicationTemplatesIn := testdata.GetCloudApplicationTemplateData()
	for _, cloudApplicationTemplateIn := range cloudApplicationTemplatesIn {
		UpdateTemplateMocked(t, cloudApplicationTemplateIn)
		UpdateTemplateFailErrMocked(t, cloudApplicationTemplateIn)
		UpdateTemplateFailStatusMocked(t, cloudApplicationTemplateIn)
	}
}

func TestParseMetadataTemplate(t *testing.T) {
	cloudApplicationTemplatesIn := testdata.GetCloudApplicationTemplateData()
	for _, cloudApplicationTemplateIn := range cloudApplicationTemplatesIn {
		ParseMetadataTemplateMocked(t, cloudApplicationTemplateIn)
		ParseMetadataTemplateFailErrMocked(t, cloudApplicationTemplateIn)
		ParseMetadataTemplateFailStatusMocked(t, cloudApplicationTemplateIn)
		ParseMetadataTemplateFailJSONMocked(t, cloudApplicationTemplateIn)
	}
}

func TestDeleteTemplate(t *testing.T) {
	cloudApplicationTemplatesIn := testdata.GetCloudApplicationTemplateData()
	for _, cloudApplicationTemplateIn := range cloudApplicationTemplatesIn {
		DeleteTemplateMocked(t, cloudApplicationTemplateIn)
		DeleteTemplateFailErrMocked(t, cloudApplicationTemplateIn)
		DeleteTemplateFailStatusMocked(t, cloudApplicationTemplateIn)
	}
}
