package audit

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// ListEventsMocked test mocked function
func ListEventsMocked(t *testing.T, eventsIn []*types.Event) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Event test data corrupted")

	// call service
	cs.On("Get", "/audit/events").Return(dIn, 200, nil)
	eventsOut, err := ds.ListEvents()
	assert.Nil(err, "Error getting event list")
	assert.Equal(eventsIn, eventsOut, "ListEvents returned different events")

	return eventsOut
}

// ListEventsFailErrMocked test mocked function
func ListEventsFailErrMocked(t *testing.T, eventsIn []*types.Event) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Event test data corrupted")

	// call service
	cs.On("Get", "/audit/events").Return(dIn, 200, fmt.Errorf("mocked error"))
	eventsOut, err := ds.ListEvents()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return eventsOut
}

// ListEventsFailStatusMocked test mocked function
func ListEventsFailStatusMocked(t *testing.T, eventsIn []*types.Event) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Event test data corrupted")

	// call service
	cs.On("Get", "/audit/events").Return(dIn, 499, nil)
	eventsOut, err := ds.ListEvents()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return eventsOut
}

// ListEventsFailJSONMocked test mocked function
func ListEventsFailJSONMocked(t *testing.T, eventsIn []*types.Event) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/audit/events").Return(dIn, 200, nil)
	eventsOut, err := ds.ListEvents()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return eventsOut
}

// ListSysEventsMocked test mocked function
func ListSysEventsMocked(t *testing.T, eventsIn []*types.Event) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Event test data corrupted")

	// call service
	cs.On("Get", "/audit/system_events").Return(dIn, 200, nil)
	eventsOut, err := ds.ListSysEvents()
	assert.Nil(err, "Error getting event list")
	assert.Equal(eventsIn, eventsOut, "ListSysEvents returned different events")

	return eventsOut
}

// ListSysEventsFailErrMocked test mocked function
func ListSysEventsFailErrMocked(t *testing.T, eventsIn []*types.Event) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Event test data corrupted")

	// call service
	cs.On("Get", "/audit/system_events").Return(dIn, 200, fmt.Errorf("mocked error"))
	eventsOut, err := ds.ListSysEvents()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return eventsOut
}

// ListSysEventsFailStatusMocked test mocked function
func ListSysEventsFailStatusMocked(t *testing.T, eventsIn []*types.Event) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Event test data corrupted")

	// call service
	cs.On("Get", "/audit/system_events").Return(dIn, 499, nil)
	eventsOut, err := ds.ListSysEvents()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return eventsOut
}

// ListSysEventsFailJSONMocked test mocked function
func ListSysEventsFailJSONMocked(t *testing.T, eventsIn []*types.Event) []*types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/audit/system_events").Return(dIn, 200, nil)
	eventsOut, err := ds.ListSysEvents()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return eventsOut
}
