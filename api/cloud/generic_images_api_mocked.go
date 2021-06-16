// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// ListGenericImagesMocked test mocked function
func ListGenericImagesMocked(t *testing.T, genericImagesIn []*types.GenericImage) []*types.GenericImage {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewGenericImageService(cs)
	assert.Nil(err, "Couldn't load genericImage service")
	assert.NotNil(ds, "GenericImage service not instanced")

	// to json
	dIn, err := json.Marshal(genericImagesIn)
	assert.Nil(err, "GenericImage test data corrupted")

	// call service
	cs.On("Get", APIPathCloudGenericImages).Return(dIn, 200, nil)
	genericImagesOut, err := ds.ListGenericImages()
	assert.Nil(err, "Error getting genericImage list")
	assert.Equal(genericImagesIn, genericImagesOut, "ListGenericImages returned different genericImages")

	return genericImagesOut
}

// ListGenericImagesFailErrMocked test mocked function
func ListGenericImagesFailErrMocked(t *testing.T, genericImagesIn []*types.GenericImage) []*types.GenericImage {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewGenericImageService(cs)
	assert.Nil(err, "Couldn't load genericImage service")
	assert.NotNil(ds, "GenericImage service not instanced")

	// to json
	dIn, err := json.Marshal(genericImagesIn)
	assert.Nil(err, "GenericImage test data corrupted")

	// call service
	cs.On("Get", APIPathCloudGenericImages).Return(dIn, 200, fmt.Errorf("mocked error"))
	genericImagesOut, err := ds.ListGenericImages()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(genericImagesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return genericImagesOut
}

// ListGenericImagesFailStatusMocked test mocked function
func ListGenericImagesFailStatusMocked(t *testing.T, genericImagesIn []*types.GenericImage) []*types.GenericImage {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewGenericImageService(cs)
	assert.Nil(err, "Couldn't load genericImage service")
	assert.NotNil(ds, "GenericImage service not instanced")

	// to json
	dIn, err := json.Marshal(genericImagesIn)
	assert.Nil(err, "GenericImage test data corrupted")

	// call service
	cs.On("Get", APIPathCloudGenericImages).Return(dIn, 499, nil)
	genericImagesOut, err := ds.ListGenericImages()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(genericImagesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return genericImagesOut
}

// ListGenericImagesFailJSONMocked test mocked function
func ListGenericImagesFailJSONMocked(t *testing.T, genericImagesIn []*types.GenericImage) []*types.GenericImage {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewGenericImageService(cs)
	assert.Nil(err, "Couldn't load genericImage service")
	assert.NotNil(ds, "GenericImage service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", APIPathCloudGenericImages).Return(dIn, 200, nil)
	genericImagesOut, err := ds.ListGenericImages()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(genericImagesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return genericImagesOut
}
