package clientbrownfield

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ImportServerMocked test mocked function
func ImportServerMocked(t *testing.T, serverIn *types.Server, cloudAccountID string) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.ImportServer(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing server for cloud account")
	assert.Equal(serverIn, serverOut, "ImportServer returned different server")

	return serverOut
}

// ImportServerFailErrMocked test mocked function
func ImportServerFailErrMocked(t *testing.T, serverIn *types.Server, cloudAccountID string) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	serverOut, err := ds.ImportServer(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverOut
}

// ImportServerFailStatusMocked test mocked function
func ImportServerFailStatusMocked(t *testing.T, serverIn *types.Server, cloudAccountID string) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.ImportServer(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// ImportServerFailJSONMocked test mocked function
func ImportServerFailJSONMocked(t *testing.T, serverIn *types.Server, cloudAccountID string) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/import_candidates/%s/import", cloudAccountID), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.ImportServer(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverOut
}

// ImportVPCMocked test mocked function
func ImportVPCMocked(t *testing.T, vpcIn *types.Vpc, cloudAccountID string) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/vpc_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 200, nil)
	vpcOut, err := ds.ImportVPC(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing VPC for cloud account")
	assert.Equal(vpcIn, vpcOut, "ImportVPC returned different VPC")

	return vpcOut
}

// ImportVPCFailErrMocked test mocked function
func ImportVPCFailErrMocked(t *testing.T, vpcIn *types.Vpc, cloudAccountID string) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/vpc_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	vpcOut, err := ds.ImportVPC(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpcOut
}

// ImportVPCFailStatusMocked test mocked function
func ImportVPCFailStatusMocked(t *testing.T, vpcIn *types.Vpc, cloudAccountID string) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(vpcIn)
	assert.Nil(err, "VPC test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/vpc_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 499, nil)
	vpcOut, err := ds.ImportVPC(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vpcOut
}

// ImportVPCFailJSONMocked test mocked function
func ImportVPCFailJSONMocked(t *testing.T, vpcIn *types.Vpc, cloudAccountID string) *types.Vpc {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/vpc_import_candidates/%s/import", cloudAccountID), mapIn).Return(dIn, 200, nil)
	vpcOut, err := ds.ImportVPC(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpcOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpcOut
}

// ImportFloatingIPMocked test mocked function
func ImportFloatingIPMocked(t *testing.T, floatingIPIn *types.FloatingIP, cloudAccountID string) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "Floating IP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "Floating IP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/floating_ip_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 200, nil)
	floatingIPOut, err := ds.ImportFloatingIP(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing floating IP for cloud account")
	assert.Equal(floatingIPIn, floatingIPOut, "ImportFloatingIP returned different floating IP")

	return floatingIPOut
}

// ImportFloatingIPFailErrMocked test mocked function
func ImportFloatingIPFailErrMocked(t *testing.T, floatingIPIn *types.FloatingIP, cloudAccountID string) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*floatingIPIn)
	assert.Nil(err, "Floating IP test data corrupted")

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "Floating IP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/floating_ip_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	floatingIPOut, err := ds.ImportFloatingIP(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPOut
}

// ImportFloatingIPFailStatusMocked test mocked function
func ImportFloatingIPFailStatusMocked(t *testing.T, floatingIPIn *types.FloatingIP, cloudAccountID string) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(floatingIPIn)
	assert.Nil(err, "Floating IP test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/floating_ip_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 499, nil)
	floatingIPOut, err := ds.ImportFloatingIP(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPOut
}

// ImportFloatingIPFailJSONMocked test mocked function
func ImportFloatingIPFailJSONMocked(t *testing.T, floatingIPIn *types.FloatingIP, cloudAccountID string) *types.FloatingIP {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/floating_ip_import_candidates/%s/import", cloudAccountID), mapIn).Return(dIn, 200, nil)
	floatingIPOut, err := ds.ImportFloatingIP(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPOut
}

// ImportVolumeMocked test mocked function
func ImportVolumeMocked(t *testing.T, volumeIn *types.Volume, cloudAccountID string) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/volume_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 200, nil)
	volumeOut, err := ds.ImportVolume(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing volume for cloud account")
	assert.Equal(volumeIn, volumeOut, "ImportVolume returned different volume")

	return volumeOut
}

// ImportVolumeFailErrMocked test mocked function
func ImportVolumeFailErrMocked(t *testing.T, volumeIn *types.Volume, cloudAccountID string) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/volume_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	volumeOut, err := ds.ImportVolume(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return volumeOut
}

// ImportVolumeFailStatusMocked test mocked function
func ImportVolumeFailStatusMocked(t *testing.T, volumeIn *types.Volume, cloudAccountID string) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(volumeIn)
	assert.Nil(err, "Volume test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/volume_import_candidates/%s/import", cloudAccountID), mapIn).Return(dOut, 499, nil)
	volumeOut, err := ds.ImportVolume(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return volumeOut
}

// ImportVolumeFailJSONMocked test mocked function
func ImportVolumeFailJSONMocked(t *testing.T, volumeIn *types.Volume, cloudAccountID string) *types.Volume {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/brownfield/volume_import_candidates/%s/import", cloudAccountID), mapIn).Return(dIn, 200, nil)
	volumeOut, err := ds.ImportVolume(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(volumeOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return volumeOut
}
