// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"golang.org/x/net/context"
)

// ListEvents returns the list of events as an array of Event
func (imco *ClientAPI) ListEvents(ctx context.Context) (events []*types.Event, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathAuditEvents, true, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// ListSysEvents returns the list of events as an array of Event
func (imco *ClientAPI) ListSysEvents(ctx context.Context) (events []*types.Event, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, pathAuditSystemEvents, true, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
