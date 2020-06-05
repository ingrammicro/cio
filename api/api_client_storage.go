// Copyright (c) 2017-2022 Ingram Micro Inc.

package api

import (
	"fmt"
	"github.com/ingrammicro/cio/logger"
	"github.com/ingrammicro/cio/types"
	"golang.org/x/net/context"
)

// GetStoragePlan returns a storage plan by its ID
func (imco *ClientAPI) GetStoragePlan(ctx context.Context, storagePlanID string,
) (storagePlan *types.StoragePlan, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathStoragePlan, storagePlanID), true, &storagePlan)
	if err != nil {
		return nil, err
	}
	return storagePlan, nil
}

// GetStorageVolume returns a Volume by its ID
func (imco *ClientAPI) GetStorageVolume(ctx context.Context, volumeID string) (volume *types.Volume, err error) {
	logger.DebugFuncInfo()

	_, err = imco.GetAndCheck(ctx, fmt.Sprintf(pathStorageVolume, volumeID), true, &volume)
	if err != nil {
		return nil, err
	}
	return volume, nil
}

// CreateStorageVolume creates a Volume
func (imco *ClientAPI) CreateStorageVolume(ctx context.Context, volumeParams *map[string]interface{},
) (volume *types.Volume, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, pathStorageVolumes, volumeParams, true, &volume)
	if err != nil {
		return nil, err
	}
	return volume, nil
}

// UpdateStorageVolume updates a Volume by its ID
func (imco *ClientAPI) UpdateStorageVolume(ctx context.Context, volumeID string, volumeParams *map[string]interface{},
) (volume *types.Volume, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PutAndCheck(ctx, fmt.Sprintf(pathStorageVolume, volumeID), volumeParams, true, &volume)
	if err != nil {
		return nil, err
	}
	return volume, nil
}

// AttachStorageVolume attaches a Volume by its ID
func (imco *ClientAPI) AttachStorageVolume(ctx context.Context, volumeID string, volumeParams *map[string]interface{},
) (server *types.Server, err error) {
	logger.DebugFuncInfo()

	_, err = imco.PostAndCheck(ctx, fmt.Sprintf(pathStorageVolumeAttachedServer, volumeID), volumeParams, true, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// DetachStorageVolume detaches a Volume by its ID
func (imco *ClientAPI) DetachStorageVolume(ctx context.Context, volumeID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathStorageVolumeAttachedServer, volumeID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteStorageVolume deletes a Volume by its ID
func (imco *ClientAPI) DeleteStorageVolume(ctx context.Context, volumeID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathStorageVolume, volumeID), true, nil)
	if err != nil {
		return err
	}
	return nil
}

// DiscardStorageVolume discards a Volume by its ID
func (imco *ClientAPI) DiscardStorageVolume(ctx context.Context, volumeID string) (err error) {
	logger.DebugFuncInfo()

	_, err = imco.DeleteAndCheck(ctx, fmt.Sprintf(pathStorageVolumeDiscard, volumeID), true, nil)
	if err != nil {
		return err
	}
	return nil
}
