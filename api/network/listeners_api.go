package network

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// ListenerService manages listener operations
type ListenerService struct {
	concertoService utils.ConcertoService
}

// NewListenerService returns a Concerto listener service
func NewListenerService(concertoService utils.ConcertoService) (*ListenerService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ListenerService{
		concertoService: concertoService,
	}, nil
}

// ListListeners returns the list of listeners in a load balancer by its ID, as an array of Listener
func (ls *ListenerService) ListListeners(loadBalancerID string) (listeners []*types.Listener, err error) {
	log.Debug("ListListeners")

	data, status, err := ls.concertoService.Get(fmt.Sprintf("/network/load_balancers/%s/listeners", loadBalancerID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &listeners); err != nil {
		return nil, err
	}

	return listeners, nil
}

// GetListener returns a listener by its ID
func (ls *ListenerService) GetListener(listenerID string) (listener *types.Listener, err error) {
	log.Debug("GetListener")

	data, status, err := ls.concertoService.Get(fmt.Sprintf("/network/listeners/%s", listenerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &listener); err != nil {
		return nil, err
	}

	return listener, nil
}

// CreateListener creates a listener in a load balancer by its ID
func (ls *ListenerService) CreateListener(loadBalancerID string, listenerParams *map[string]interface{}) (listener *types.Listener, err error) {
	log.Debug("CreateListener")

	data, status, err := ls.concertoService.Post(fmt.Sprintf("/network/load_balancers/%s/listeners", loadBalancerID), listenerParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &listener); err != nil {
		return nil, err
	}

	return listener, nil
}

// UpdateListener updates a listener by its ID
func (ls *ListenerService) UpdateListener(listenerID string, listenerParams *map[string]interface{}) (listener *types.Listener, err error) {
	log.Debug("UpdateListener")

	data, status, err := ls.concertoService.Put(fmt.Sprintf("/network/listeners/%s", listenerID), listenerParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &listener); err != nil {
		return nil, err
	}

	return listener, nil
}

// DeleteListener deletes a listener by its ID
func (ls *ListenerService) DeleteListener(listenerID string) (err error) {
	log.Debug("DeleteListener")

	data, status, err := ls.concertoService.Delete(fmt.Sprintf("/network/listeners/%s", listenerID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// RetryListener retries a listener by its ID
func (ls *ListenerService) RetryListener(listenerID string, listenerParams *map[string]interface{}) (listener *types.Listener, err error) {
	log.Debug("RetryListener")

	data, status, err := ls.concertoService.Put(fmt.Sprintf("/network/listeners/%s/retry", listenerID), listenerParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &listener); err != nil {
		return nil, err
	}

	return listener, nil
}

// ListRules returns the list of rules in a listener by its ID, as an array of ListenerRule
func (ls *ListenerService) ListRules(listenerID string) (listenerRules []*types.ListenerRule, err error) {
	log.Debug("ListRules")

	data, status, err := ls.concertoService.Get(fmt.Sprintf("/network/listeners/%s/rules", listenerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &listenerRules); err != nil {
		return nil, err
	}

	return listenerRules, nil
}

// CreateRule creates a rule in a listener by its ID
func (ls *ListenerService) CreateRule(listenerID string, listenerRuleParams *map[string]interface{}) (listenerRule *types.ListenerRule, err error) {
	log.Debug("CreateRule")

	data, status, err := ls.concertoService.Post(fmt.Sprintf("/network/listeners/%s/rules", listenerID), listenerRuleParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &listenerRule); err != nil {
		return nil, err
	}

	return listenerRule, nil
}

// UpdateRule updates a rule in a listener by its ID
func (ls *ListenerService) UpdateRule(listenerID string, listenerRuleID string, listenerRuleParams *map[string]interface{}) (listenerRule *types.ListenerRule, err error) {
	log.Debug("UpdateRule")

	data, status, err := ls.concertoService.Put(fmt.Sprintf("/network/listeners/%s/rules/%s", listenerID, listenerRuleID), listenerRuleParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &listenerRule); err != nil {
		return nil, err
	}

	return listenerRule, nil
}

// DeleteRule deletes a rule in a listener by given IDs
func (ls *ListenerService) DeleteRule(listenerID string, listenerRuleID string) (err error) {
	log.Debug("DeleteRule")

	data, status, err := ls.concertoService.Delete(fmt.Sprintf("/network/listeners/%s/rules/%s", listenerID, listenerRuleID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
