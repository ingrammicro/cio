// Copyright (c) 2017-2021 Ingram Micro Inc.

package storage

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewVolumeServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewVolumeService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListVolumes(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	ListVolumesMocked(t, volumesIn)
	ListVolumesMockedFilteredByServer(t, volumesIn)
	ListVolumesFailErrMocked(t, volumesIn)
	ListVolumesFailStatusMocked(t, volumesIn)
	ListVolumesFailJSONMocked(t, volumesIn)
}

func TestGetVolume(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	for _, volumeIn := range volumesIn {
		GetVolumeMocked(t, volumeIn)
		GetVolumeFailErrMocked(t, volumeIn)
		GetVolumeFailStatusMocked(t, volumeIn)
		GetVolumeFailJSONMocked(t, volumeIn)
	}
}

func TestCreateVolume(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	for _, volumeIn := range volumesIn {
		CreateVolumeMocked(t, volumeIn)
		CreateVolumeFailErrMocked(t, volumeIn)
		CreateVolumeFailStatusMocked(t, volumeIn)
		CreateVolumeFailJSONMocked(t, volumeIn)
	}
}

func TestUpdateVolume(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	for _, volumeIn := range volumesIn {
		UpdateVolumeMocked(t, volumeIn)
		UpdateVolumeFailErrMocked(t, volumeIn)
		UpdateVolumeFailStatusMocked(t, volumeIn)
		UpdateVolumeFailJSONMocked(t, volumeIn)
	}
}

func TestAttachVolume(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	for _, volumeIn := range volumesIn {
		AttachVolumeMocked(t, volumeIn)
		AttachVolumeFailErrMocked(t, volumeIn)
		AttachVolumeFailStatusMocked(t, volumeIn)
		AttachVolumeFailJSONMocked(t, volumeIn)
	}
}

func TestDetachVolume(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	for _, volumeIn := range volumesIn {
		DetachVolumeMocked(t, volumeIn)
		DetachVolumeFailErrMocked(t, volumeIn)
		DetachVolumeFailStatusMocked(t, volumeIn)
	}
}

func TestDeleteVolume(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	for _, volumeIn := range volumesIn {
		DeleteVolumeMocked(t, volumeIn)
		DeleteVolumeFailErrMocked(t, volumeIn)
		DeleteVolumeFailStatusMocked(t, volumeIn)
	}
}

func TestDiscardVolume(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	for _, volumeIn := range volumesIn {
		DiscardVolumeMocked(t, volumeIn)
		DiscardVolumeFailErrMocked(t, volumeIn)
		DiscardVolumeFailStatusMocked(t, volumeIn)
	}
}
