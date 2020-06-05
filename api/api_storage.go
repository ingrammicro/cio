// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"fmt"

	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
)

// GetStoragePlan returns a storage plan by its ID
func (imco *IMCOClient) GetStoragePlan(storagePlanID string) (storagePlan *types.StoragePlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathStoragePlan, storagePlanID), true, &storagePlan)
	if err != nil {
		return nil, err
	}
	return storagePlan, nil
}

// GetStorageVolume returns a Volume by its ID
func (imco *IMCOClient) GetStorageVolume(volumeID string) (volume *types.Volume, err error) {
	logger.DebugFuncInfo()

	_, err = imco.getAndCheck(fmt.Sprintf(pathStorageVolume, volumeID), true, &volume)
	if err != nil {
		return nil, err
	}
	return volume, nil
}

// CreateStorageVolume creates a Volume
func (imco *IMCOClient) CreateStorageVolume(volumeParams *map[string]interface{}) (volume *types.Volume, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(pathStorageVolumes, volumeParams, true, &volume)
	if err != nil {
		return nil, err
	}
	return volume, nil
}

// UpdateStorageVolume updates a Volume by its ID
func (imco *IMCOClient) UpdateStorageVolume(volumeID string, volumeParams *map[string]interface{},
) (volume *types.Volume, err error) {
	logger.DebugFuncInfo()

	_, err = imco.putAndCheck(fmt.Sprintf(pathStorageVolume, volumeID), volumeParams, true, &volume)
	if err != nil {
		return nil, err
	}
	return volume, nil
}

// AttachStorageVolume attaches a Volume by its ID
func (imco *IMCOClient) AttachStorageVolume(volumeID string, volumeParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.postAndCheck(fmt.Sprintf(pathStorageVolumeAttachedServer, volumeID), volumeParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// DetachStorageVolume detaches a Volume by its ID
func (imco *IMCOClient) DetachStorageVolume(volumeID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathStorageVolumeAttachedServer, volumeID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteStorageVolume deletes a Volume by its ID
func (imco *IMCOClient) DeleteStorageVolume(volumeID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathStorageVolume, volumeID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DiscardStorageVolume discards a Volume by its ID
func (imco *IMCOClient) DiscardStorageVolume(volumeID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.deleteAndCheck(fmt.Sprintf(pathStorageVolumeDiscard, volumeID), true, nil)
	if err != nil {
		return err
	}
	return nil
}
