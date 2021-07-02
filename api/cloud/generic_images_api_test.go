// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewGenericImageServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewGenericImageService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListGenericImages(t *testing.T) {
	genericImagesIn := testdata.GetGenericImageData()
	ListGenericImagesMocked(t, genericImagesIn)
	ListGenericImagesFailErrMocked(t, genericImagesIn)
	ListGenericImagesFailStatusMocked(t, genericImagesIn)
	ListGenericImagesFailJSONMocked(t, genericImagesIn)
}
