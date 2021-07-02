// Copyright (c) 2017-2021 Ingram Micro Inc.

package storage

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathStorageVolumes = "/storage/volumes"
const APIPathStorageVolume = "/storage/volumes/%s"
const APIPathCloudServerVolumes = "/cloud/servers/%s/volumes"
const APIPathStorageVolumeAttachedServer = "/storage/volumes/%s/attached_server"
const APIPathStorageVolumeDiscard = "/storage/volumes/%s/discard"

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

	path := APIPathStorageVolumes
	if serverID != "" {
		path = fmt.Sprintf(APIPathCloudServerVolumes, serverID)

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

	data, status, err := vs.concertoService.Get(fmt.Sprintf(APIPathStorageVolume, volumeID))
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

	data, status, err := vs.concertoService.Post(APIPathStorageVolumes, volumeParams)

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
func (vs *VolumeService) UpdateVolume(
	volumeID string,
	volumeParams *map[string]interface{},
) (volume *types.Volume, err error) {
	log.Debug("UpdateVolume")

	data, status, err := vs.concertoService.Put(fmt.Sprintf(APIPathStorageVolume, volumeID), volumeParams)

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
func (vs *VolumeService) AttachVolume(
	volumeID string,
	volumeParams *map[string]interface{},
) (server *types.Server, err error) {
	log.Debug("AttachVolume")

	data, status, err := vs.concertoService.Post(
		fmt.Sprintf(APIPathStorageVolumeAttachedServer, volumeID),
		volumeParams,
	)

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

	data, status, err := vs.concertoService.Delete(fmt.Sprintf(APIPathStorageVolumeAttachedServer, volumeID))
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

	data, status, err := vs.concertoService.Delete(fmt.Sprintf(APIPathStorageVolume, volumeID))
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

	data, status, err := vs.concertoService.Delete(fmt.Sprintf(APIPathStorageVolumeDiscard, volumeID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
