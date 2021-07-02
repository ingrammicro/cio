// Copyright (c) 2017-2021 Ingram Micro Inc.

package audit

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathAuditEvents = "/audit/events"
const APIPathAuditSystemEvents = "/audit/system_events"

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

// ListEvents returns the list of events as an array of Event
func (es *EventService) ListEvents() (events []*types.Event, err error) {
	log.Debug("ListEvents")

	data, status, err := es.concertoService.Get(APIPathAuditEvents)
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

// ListSysEvents returns the list of events as an array of Event
func (es *EventService) ListSysEvents() (events []*types.Event, err error) {
	log.Debug("ListSysEvents")

	data, status, err := es.concertoService.Get(APIPathAuditSystemEvents)
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
