// Copyright (c) 2017-2021 Ingram Micro Inc.

package blueprint

import (
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewTemplateServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewTemplateService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListTemplates(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	ListTemplatesMocked(t, templatesIn)
	ListTemplatesFailErrMocked(t, templatesIn)
	ListTemplatesFailStatusMocked(t, templatesIn)
	ListTemplatesFailJSONMocked(t, templatesIn)
}

func TestGetTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range templatesIn {
		GetTemplateMocked(t, templateIn)
		GetTemplateFailErrMocked(t, templateIn)
		GetTemplateFailStatusMocked(t, templateIn)
		GetTemplateFailJSONMocked(t, templateIn)
	}
}

func TestCreateTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range templatesIn {
		CreateTemplateMocked(t, templateIn)
		CreateTemplateFailErrMocked(t, templateIn)
		CreateTemplateFailStatusMocked(t, templateIn)
		CreateTemplateFailJSONMocked(t, templateIn)
	}
}

func TestUpdateTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range templatesIn {
		UpdateTemplateMocked(t, templateIn)
		UpdateTemplateFailErrMocked(t, templateIn)
		UpdateTemplateFailStatusMocked(t, templateIn)
		UpdateTemplateFailJSONMocked(t, templateIn)
	}
}

func TestCompileTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range templatesIn {
		CompileTemplateMocked(t, templateIn)
		CompileTemplateFailErrMocked(t, templateIn)
		CompileTemplateFailStatusMocked(t, templateIn)
		CompileTemplateFailJSONMocked(t, templateIn)
	}
}

func TestDeleteTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range templatesIn {
		DeleteTemplateMocked(t, templateIn)
		DeleteTemplateFailErrMocked(t, templateIn)
		DeleteTemplateFailStatusMocked(t, templateIn)
	}
}

func TestListTemplateScripts(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range drsIn {
		ListTemplateScriptsMocked(t, drsIn, drIn.ID, drIn.Type)
		ListTemplateScriptsFailErrMocked(t, drsIn, drIn.ID, drIn.Type)
		ListTemplateScriptsFailStatusMocked(t, drsIn, drIn.ID, drIn.Type)
		ListTemplateScriptsFailJSONMocked(t, drsIn, drIn.ID, drIn.Type)
	}
}

func TestGetTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range drsIn {
		GetTemplateScriptMocked(t, drIn)
		GetTemplateScriptFailErrMocked(t, drIn)
		GetTemplateScriptFailStatusMocked(t, drIn)
		GetTemplateScriptFailJSONMocked(t, drIn)
	}
}

func TestCreateTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range drsIn {
		CreateTemplateScriptMocked(t, drIn)
		CreateTemplateScriptFailErrMocked(t, drIn)
		CreateTemplateScriptFailStatusMocked(t, drIn)
		CreateTemplateScriptFailJSONMocked(t, drIn)
	}
}

func TestUpdateTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range drsIn {
		UpdateTemplateScriptMocked(t, drIn)
		UpdateTemplateScriptFailErrMocked(t, drIn)
		UpdateTemplateScriptFailStatusMocked(t, drIn)
		UpdateTemplateScriptFailJSONMocked(t, drIn)
	}
}

func TestDeleteTemplateScripts(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range drsIn {
		DeleteTemplateScriptMocked(t, drIn)
		DeleteTemplateScriptFailErrMocked(t, drIn)
		DeleteTemplateScriptFailStatusMocked(t, drIn)
	}
}

func TestListTemplateServers(t *testing.T) {
	drsIn := testdata.GetTemplateServerData()
	for _, drIn := range drsIn {
		ListTemplateServersMocked(t, drsIn, drIn.ID)
		ListTemplateServersFailErrMocked(t, drsIn, drIn.ID)
		ListTemplateServersFailStatusMocked(t, drsIn, drIn.ID)
		ListTemplateServersFailJSONMocked(t, drsIn, drIn.ID)
	}
}

func TestReorderTemplateScript(t *testing.T) {
	tsIn := testdata.GetTemplateScriptData()

	// get template
	templateID := tsIn[0].TemplateID

	// reorder
	num := len(tsIn)
	reorder := make([]string, num, num)
	tsOut := make([]*types.TemplateScript, num, num)

	num--
	for i, ts := range tsIn {
		reorder[num-i] = ts.ID
		tsOut[num-i] = ts
	}

	ReorderTemplateScriptMocked(t, tsOut, templateID, reorder)
	ReorderTemplateScriptFailErrMocked(t, tsOut, templateID, reorder)
	ReorderTemplateScriptFailStatusMocked(t, tsOut, templateID, reorder)
	ReorderTemplateScriptFailJSONMocked(t, tsOut, templateID, reorder)
}
