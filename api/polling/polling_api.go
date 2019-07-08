package polling

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
)

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
func (p *PollingService) Ping() (ping *types.PollingPing, status int, err error) {
	log.Debug("Ping")

	payload := make(map[string]interface{})
	data, status, err := p.concertoService.Post("/command_polling/pings", &payload)
	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &ping); err != nil {
		return nil, status, err
	}

	return ping, status, nil
}

// GetNextCommand returns the command to be executed
func (p *PollingService) GetNextCommand() (command *types.PollingCommand, status int, err error) {
	log.Debug("GetNextCommand")

	data, status, err := p.concertoService.Get("/command_polling/command")
	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}

// UpdateCommand updates a command by its ID
func (p *PollingService) UpdateCommand(pollingCommandVector *map[string]interface{}, ID string) (command *types.PollingCommand, status int, err error) {
	log.Debug("UpdateCommand")

	data, status, err := p.concertoService.Put(fmt.Sprintf("/command_polling/commands/%s", ID), pollingCommandVector)
	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}

// ReportBootstrapLog reports a command result
func (p *PollingService) ReportBootstrapLog(PollingContinuousReportVector *map[string]interface{}) (command *types.PollingContinuousReport, status int, err error) {
	log.Debug("ReportBootstrapLog")

	data, status, err := p.concertoService.Post("/command_polling/bootstrap_logs", PollingContinuousReportVector)
	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}
