// Copyright (c) 2017-2021 Ingram Micro Inc.

package firewall

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathCloudFirewallProfile = "/cloud/firewall_profile"
const APIPathCloudFirewallProfileRules = "/cloud/firewall_profile/rules"

// FirewallService manages firewall operations
type FirewallService struct {
	concertoService utils.ConcertoService
}

// NewFirewallService returns a Concerto firewall service
func NewFirewallService(concertoService utils.ConcertoService) (*FirewallService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &FirewallService{
		concertoService: concertoService,
	}, nil
}

// GetPolicy returns firewall policy
func (fs *FirewallService) GetPolicy() (policy *types.Policy, err error) {
	log.Debug("GetPolicy")

	data, status, err := fs.concertoService.Get(APIPathCloudFirewallProfile)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &policy); err != nil {
		return nil, err
	}
	policy.Md5 = fmt.Sprintf("%x", md5.Sum(data))

	return policy, nil
}

// AddPolicyRule adds a new firewall policy rule
func (fs *FirewallService) AddPolicyRule(ruleParams *map[string]interface{}) (policyRule *types.PolicyRule, err error) {
	log.Debug("AddPolicyRule")

	data, status, err := fs.concertoService.Post(APIPathCloudFirewallProfileRules, ruleParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &policyRule); err != nil {
		return nil, err
	}
	return policyRule, nil
}

// UpdatePolicy update firewall policy
func (fs *FirewallService) UpdatePolicy(policyParams *map[string]interface{}) (policy *types.Policy, err error) {
	log.Debug("UpdatePolicy")

	data, status, err := fs.concertoService.Put(APIPathCloudFirewallProfile, policyParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &policy); err != nil {
		return nil, err
	}
	return policy, nil
}
