// Copyright (c) 2017-2021 Ingram Micro Inc.

package polling

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCommandPollingPings = "/command_polling/pings"
const APIPathCommandPollingNextCommand = "/command_polling/command"
const APIPathCommandPollingCommand = "/command_polling/commands/%s"
const APIPathCommandPollingBootstrapLogs = "/command_polling/bootstrap_logs"

// PollingService manages polling operations
type PollingService struct {
	concertoService utils.ConcertoService
}

// NewPollingService returns a Concerto polling service
func NewPollingService(concertoService utils.ConcertoService) (*PollingService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &PollingService{
		concertoService: concertoService,
	}, nil
}

// Ping resolves if new command is waiting for execution
func (ps *PollingService) Ping() (ping *types.PollingPing, status int, err error) {
	log.Debug("Ping")

	payload := make(map[string]interface{})
	data, status, err := ps.concertoService.Post(APIPathCommandPollingPings, &payload)
	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &ping); err != nil {
		return nil, status, err
	}

	return ping, status, nil
}

// GetNextCommand returns the command to be executed
func (ps *PollingService) GetNextCommand() (command *types.PollingCommand, status int, err error) {
	log.Debug("GetNextCommand")

	data, status, err := ps.concertoService.Get(APIPathCommandPollingNextCommand)
	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}

// UpdateCommand updates a command by its ID
func (ps *PollingService) UpdateCommand(
	commandID string,
	pollingCommandParams *map[string]interface{},
) (command *types.PollingCommand, status int, err error) {
	log.Debug("UpdateCommand")

	data, status, err := ps.concertoService.Put(
		fmt.Sprintf(APIPathCommandPollingCommand, commandID),
		pollingCommandParams,
	)

	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}

// ReportBootstrapLog reports a command result
func (ps *PollingService) ReportBootstrapLog(
	pollingContinuousReportParams *map[string]interface{},
) (command *types.PollingContinuousReport, status int, err error) {
	log.Debug("ReportBootstrapLog")

	data, status, err := ps.concertoService.Post(APIPathCommandPollingBootstrapLogs, pollingContinuousReportParams)

	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}
