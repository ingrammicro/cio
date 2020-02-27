package audit

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// EventService manages event operations
type EventService struct {
	concertoService utils.ConcertoService
}

// NewEventService returns a Concerto event service
func NewEventService(concertoService utils.ConcertoService) (*EventService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &EventService{
		concertoService: concertoService,
	}, nil
}

// GetEventList returns the list of events as an array of Event
func (cl *EventService) GetEventList() (events []*types.Event, err error) {
	log.Debug("GetEventList")

	data, status, err := cl.concertoService.Get("/audit/events")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// GetSysEventList returns the list of events as an array of Event
func (cl *EventService) GetSysEventList() (events []*types.Event, err error) {
	log.Debug("GetSysEventList")

	data, status, err := cl.concertoService.Get("/audit/system_events")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	return events, nil
}
