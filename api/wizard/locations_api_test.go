// Copyright (c) 2017-2021 Ingram Micro Inc.

package wizard

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewLocationServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewLocationService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListLocations(t *testing.T) {
	locationsIn := testdata.GetLocationData()
	ListLocationsMocked(t, locationsIn)
	ListLocationsFailErrMocked(t, locationsIn)
	ListLocationsFailStatusMocked(t, locationsIn)
	ListLocationsFailJSONMocked(t, locationsIn)
}
