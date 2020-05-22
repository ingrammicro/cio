package network

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// FirewallProfileService manages firewall profile operations
type FirewallProfileService struct {
	concertoService utils.ConcertoService
}

// NewFirewallProfileService returns a Concerto firewallProfile service
func NewFirewallProfileService(concertoService utils.ConcertoService) (*FirewallProfileService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &FirewallProfileService{
		concertoService: concertoService,
	}, nil
}

// ListFirewallProfiles returns the list of firewallProfiles as an array of FirewallProfile
func (fps *FirewallProfileService) ListFirewallProfiles() (firewallProfiles []*types.FirewallProfile, err error) {
	log.Debug("ListFirewallProfiles")

	data, status, err := fps.concertoService.Get("/network/firewall_profiles")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &firewallProfiles); err != nil {
		return nil, err
	}

	return firewallProfiles, nil
}

// GetFirewallProfile returns a firewallProfile by its ID
func (fps *FirewallProfileService) GetFirewallProfile(firewallProfileID string) (firewallProfile *types.FirewallProfile, err error) {
	log.Debug("GetFirewallProfile")

	data, status, err := fps.concertoService.Get(fmt.Sprintf("/network/firewall_profiles/%s", firewallProfileID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &firewallProfile); err != nil {
		return nil, err
	}

	return firewallProfile, nil
}

// CreateFirewallProfile creates a firewallProfile
func (fps *FirewallProfileService) CreateFirewallProfile(firewallProfileParams *map[string]interface{}) (firewallProfile *types.FirewallProfile, err error) {
	log.Debug("CreateFirewallProfile")

	data, status, err := fps.concertoService.Post("/network/firewall_profiles/", firewallProfileParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &firewallProfile); err != nil {
		return nil, err
	}

	return firewallProfile, nil
}

// UpdateFirewallProfile updates a firewallProfile by its ID
func (fps *FirewallProfileService) UpdateFirewallProfile(firewallProfileParams *map[string]interface{}, ID string) (firewallProfile *types.FirewallProfile, err error) {
	log.Debug("UpdateFirewallProfile")

	data, status, err := fps.concertoService.Put(fmt.Sprintf("/network/firewall_profiles/%s", ID), firewallProfileParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &firewallProfile); err != nil {
		return nil, err
	}

	return firewallProfile, nil
}

// DeleteFirewallProfile deletes a firewallProfile by its ID
func (fps *FirewallProfileService) DeleteFirewallProfile(firewallProfileID string) (err error) {
	log.Debug("DeleteFirewallProfile")

	data, status, err := fps.concertoService.Delete(fmt.Sprintf("/network/firewall_profiles/%s", firewallProfileID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
