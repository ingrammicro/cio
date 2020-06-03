package cloud

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSSHProfileServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewSSHProfileService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListSSHProfiles(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	ListSSHProfilesMocked(t, sshProfilesIn)
	ListSSHProfilesFailErrMocked(t, sshProfilesIn)
	ListSSHProfilesFailStatusMocked(t, sshProfilesIn)
	ListSSHProfilesFailJSONMocked(t, sshProfilesIn)
}

func TestGetSSHProfile(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	for _, sshProfileIn := range sshProfilesIn {
		GetSSHProfileMocked(t, sshProfileIn)
		GetSSHProfileFailErrMocked(t, sshProfileIn)
		GetSSHProfileFailStatusMocked(t, sshProfileIn)
		GetSSHProfileFailJSONMocked(t, sshProfileIn)
	}
}

func TestCreateSSHProfile(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	for _, sshProfileIn := range sshProfilesIn {
		CreateSSHProfileMocked(t, sshProfileIn)
		CreateSSHProfileFailErrMocked(t, sshProfileIn)
		CreateSSHProfileFailStatusMocked(t, sshProfileIn)
		CreateSSHProfileFailJSONMocked(t, sshProfileIn)
	}
}

func TestUpdateSSHProfile(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	for _, sshProfileIn := range sshProfilesIn {
		UpdateSSHProfileMocked(t, sshProfileIn)
		UpdateSSHProfileFailErrMocked(t, sshProfileIn)
		UpdateSSHProfileFailStatusMocked(t, sshProfileIn)
		UpdateSSHProfileFailJSONMocked(t, sshProfileIn)
	}
}

func TestDeleteSSHProfile(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	for _, sshProfileIn := range sshProfilesIn {
		DeleteSSHProfileMocked(t, sshProfileIn)
		DeleteSSHProfileFailErrMocked(t, sshProfileIn)
		DeleteSSHProfileFailStatusMocked(t, sshProfileIn)
	}
}
