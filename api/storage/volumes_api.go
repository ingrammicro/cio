package storage

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// VolumeService manages volume operations
type VolumeService struct {
	concertoService utils.ConcertoService
}

// NewVolumeService returns a Concerto Volume service
func NewVolumeService(concertoService utils.ConcertoService) (*VolumeService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &VolumeService{
		concertoService: concertoService,
	}, nil
}

// ListVolumes returns the list of Volumes as an array of Volume
func (vs *VolumeService) ListVolumes(serverID string) (volumes []*types.Volume, err error) {
	log.Debug("ListVolumes")

	path := "/storage/volumes"
	if serverID != "" {
		path = fmt.Sprintf("/cloud/servers/%s/volumes", serverID)

	}
	data, status, err := vs.concertoService.Get(path)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volumes); err != nil {
		return nil, err
	}

	return volumes, nil
}

// GetVolume returns a Volume by its ID
func (vs *VolumeService) GetVolume(volumeID string) (volume *types.Volume, err error) {
	log.Debug("GetVolume")

	data, status, err := vs.concertoService.Get(fmt.Sprintf("/storage/volumes/%s", volumeID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volume); err != nil {
		return nil, err
	}

	return volume, nil
}

// CreateVolume creates a Volume
func (vs *VolumeService) CreateVolume(volumeParams *map[string]interface{}) (volume *types.Volume, err error) {
	log.Debug("CreateVolume")

	data, status, err := vs.concertoService.Post("/storage/volumes/", volumeParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volume); err != nil {
		return nil, err
	}

	return volume, nil
}

// UpdateVolume updates a Volume by its ID
func (vs *VolumeService) UpdateVolume(volumeParams *map[string]interface{}, volumeID string) (volume *types.Volume, err error) {
	log.Debug("UpdateVolume")

	data, status, err := vs.concertoService.Put(fmt.Sprintf("/storage/volumes/%s", volumeID), volumeParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &volume); err != nil {
		return nil, err
	}

	return volume, nil
}

// AttachVolume attaches a Volume by its ID
func (vs *VolumeService) AttachVolume(volumeParams *map[string]interface{}, volumeID string) (server *types.Server, err error) {
	log.Debug("AttachVolume")

	data, status, err := vs.concertoService.Post(fmt.Sprintf("/storage/volumes/%s/attached_server", volumeID), volumeParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &server); err != nil {
		return nil, err
	}

	return server, nil
}

// DetachVolume detaches a Volume by its ID
func (vs *VolumeService) DetachVolume(volumeID string) (err error) {
	log.Debug("DetachVolume")

	data, status, err := vs.concertoService.Delete(fmt.Sprintf("/storage/volumes/%s/attached_server", volumeID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DeleteVolume deletes a Volume by its ID
func (vs *VolumeService) DeleteVolume(volumeID string) (err error) {
	log.Debug("DeleteVolume")

	data, status, err := vs.concertoService.Delete(fmt.Sprintf("/storage/volumes/%s", volumeID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// DiscardVolume discards a Volume by its ID
func (vs *VolumeService) DiscardVolume(volumeID string) (err error) {
	log.Debug("DiscardVolume")

	data, status, err := vs.concertoService.Delete(fmt.Sprintf("/storage/volumes/%s/discard", volumeID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
