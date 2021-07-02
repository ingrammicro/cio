// Copyright (c) 2017-2021 Ingram Micro Inc.

package audit

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewEventServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewEventService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListEvents(t *testing.T) {
	eventsIn := testdata.GetEventData()
	ListEventsMocked(t, eventsIn)
	ListEventsFailErrMocked(t, eventsIn)
	ListEventsFailStatusMocked(t, eventsIn)
	ListEventsFailJSONMocked(t, eventsIn)
}

func TestListSysEvents(t *testing.T) {
	eventsIn := testdata.GetEventData()
	ListSysEventsMocked(t, eventsIn)
	ListSysEventsFailErrMocked(t, eventsIn)
	ListSysEventsFailStatusMocked(t, eventsIn)
	ListSysEventsFailJSONMocked(t, eventsIn)
}
