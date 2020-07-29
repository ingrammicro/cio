package wizard

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// ListLocationsMocked test mocked function
func ListLocationsMocked(t *testing.T, locationsIn []*types.Location) []*types.Location {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLocationService(cs)
	assert.Nil(err, "Couldn't load location service")
	assert.NotNil(ds, "Location service not instanced")

	// to json
	dIn, err := json.Marshal(locationsIn)
	assert.Nil(err, "Location test data corrupted")

	// call service
	cs.On("Get", "/wizard/locations").Return(dIn, 200, nil)
	locationsOut, err := ds.ListLocations()
	assert.Nil(err, "Error getting location list")
	assert.Equal(locationsIn, locationsOut, "ListLocations returned different locations")

	return locationsOut
}

// ListLocationsFailErrMocked test mocked function
func ListLocationsFailErrMocked(t *testing.T, locationsIn []*types.Location) []*types.Location {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLocationService(cs)
	assert.Nil(err, "Couldn't load location service")
	assert.NotNil(ds, "Location service not instanced")

	// to json
	dIn, err := json.Marshal(locationsIn)
	assert.Nil(err, "Location test data corrupted")

	// call service
	cs.On("Get", "/wizard/locations").Return(dIn, 200, fmt.Errorf("mocked error"))
	locationsOut, err := ds.ListLocations()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(locationsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return locationsOut
}

// ListLocationsFailStatusMocked test mocked function
func ListLocationsFailStatusMocked(t *testing.T, locationsIn []*types.Location) []*types.Location {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLocationService(cs)
	assert.Nil(err, "Couldn't load location service")
	assert.NotNil(ds, "Location service not instanced")

	// to json
	dIn, err := json.Marshal(locationsIn)
	assert.Nil(err, "Location test data corrupted")

	// call service
	cs.On("Get", "/wizard/locations").Return(dIn, 499, nil)
	locationsOut, err := ds.ListLocations()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(locationsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return locationsOut
}

// ListLocationsFailJSONMocked test mocked function
func ListLocationsFailJSONMocked(t *testing.T, locationsIn []*types.Location) []*types.Location {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLocationService(cs)
	assert.Nil(err, "Couldn't load location service")
	assert.NotNil(ds, "Location service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/wizard/locations").Return(dIn, 200, nil)
	locationsOut, err := ds.ListLocations()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(locationsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return locationsOut
}
