// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewFirewallProfileServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewFirewallProfileService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListFirewallProfiles(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	ListFirewallProfilesMocked(t, firewallProfilesIn)
	ListFirewallProfilesFailErrMocked(t, firewallProfilesIn)
	ListFirewallProfilesFailStatusMocked(t, firewallProfilesIn)
	ListFirewallProfilesFailJSONMocked(t, firewallProfilesIn)
}

func TestGetFirewallProfile(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	for _, firewallProfileIn := range firewallProfilesIn {
		GetFirewallProfileMocked(t, firewallProfileIn)
		GetFirewallProfileFailErrMocked(t, firewallProfileIn)
		GetFirewallProfileFailStatusMocked(t, firewallProfileIn)
		GetFirewallProfileFailJSONMocked(t, firewallProfileIn)
	}
}

func TestCreateFirewallProfile(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	for _, firewallProfileIn := range firewallProfilesIn {
		CreateFirewallProfileMocked(t, firewallProfileIn)
		CreateFirewallProfileFailErrMocked(t, firewallProfileIn)
		CreateFirewallProfileFailStatusMocked(t, firewallProfileIn)
		CreateFirewallProfileFailJSONMocked(t, firewallProfileIn)
	}
}

func TestUpdateFirewallProfile(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	for _, firewallProfileIn := range firewallProfilesIn {
		UpdateFirewallProfileMocked(t, firewallProfileIn)
		UpdateFirewallProfileFailErrMocked(t, firewallProfileIn)
		UpdateFirewallProfileFailStatusMocked(t, firewallProfileIn)
		UpdateFirewallProfileFailJSONMocked(t, firewallProfileIn)
	}
}

func TestDeleteFirewallProfile(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	for _, firewallProfileIn := range firewallProfilesIn {
		DeleteFirewallProfileMocked(t, firewallProfileIn)
		DeleteFirewallProfileFailErrMocked(t, firewallProfileIn)
		DeleteFirewallProfileFailStatusMocked(t, firewallProfileIn)
	}
}
