// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// ListEvents returns the list of events as an array of Event
func (imco *IMCOClient) ListEvents() (events []*types.Event, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathAuditEvents, true, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// ListSysEvents returns the list of events as an array of Event
func (imco *IMCOClient) ListSysEvents() (events []*types.Event, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(pathAuditSystemEvents, true, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
