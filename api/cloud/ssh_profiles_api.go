package cloud

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// SSHProfileService manages ssh profile operations
type SSHProfileService struct {
	concertoService utils.ConcertoService
}

// NewSSHProfileService returns a Concerto sshProfile service
func NewSSHProfileService(concertoService utils.ConcertoService) (*SSHProfileService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &SSHProfileService{
		concertoService: concertoService,
	}, nil
}

// ListSSHProfiles returns the list of sshProfiles as an array of SSHProfile
func (sps *SSHProfileService) ListSSHProfiles() (sshProfiles []*types.SSHProfile, err error) {
	log.Debug("ListSSHProfiles")

	data, status, err := sps.concertoService.Get("/cloud/ssh_profiles")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &sshProfiles); err != nil {
		return nil, err
	}

	return sshProfiles, nil
}

// GetSSHProfile returns a sshProfile by its ID
func (sps *SSHProfileService) GetSSHProfile(sshProfileID string) (sshProfile *types.SSHProfile, err error) {
	log.Debug("GetSSHProfile")

	data, status, err := sps.concertoService.Get(fmt.Sprintf("/cloud/ssh_profiles/%s", sshProfileID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &sshProfile); err != nil {
		return nil, err
	}

	return sshProfile, nil
}

// CreateSSHProfile creates a sshProfile
func (sps *SSHProfileService) CreateSSHProfile(sshProfileParams *map[string]interface{}) (sshProfile *types.SSHProfile, err error) {
	log.Debug("CreateSSHProfile")

	data, status, err := sps.concertoService.Post("/cloud/ssh_profiles/", sshProfileParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &sshProfile); err != nil {
		return nil, err
	}

	return sshProfile, nil
}

// UpdateSSHProfile updates a sshProfile by its ID
func (sps *SSHProfileService) UpdateSSHProfile(sshProfileID string, sshProfileParams *map[string]interface{}) (sshProfile *types.SSHProfile, err error) {
	log.Debug("UpdateSSHProfile")

	data, status, err := sps.concertoService.Put(fmt.Sprintf("/cloud/ssh_profiles/%s", sshProfileID), sshProfileParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &sshProfile); err != nil {
		return nil, err
	}

	return sshProfile, nil
}

// DeleteSSHProfile deletes a sshProfile by its ID
func (sps *SSHProfileService) DeleteSSHProfile(sshProfileID string) (err error) {
	log.Debug("DeleteSSHProfile")

	data, status, err := sps.concertoService.Delete(fmt.Sprintf("/cloud/ssh_profiles/%s", sshProfileID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
