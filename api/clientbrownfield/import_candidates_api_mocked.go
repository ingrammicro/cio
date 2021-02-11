package clientbrownfield

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ImportServersMocked test mocked function
func ImportServersMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportServers test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_servers", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.ImportServers(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing servers for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "ImportServers returned different cloud account")

	return cloudAccountOut
}

// ImportServersFailErrMocked test mocked function
func ImportServersFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportServers test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_servers", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.ImportServers(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// ImportServersFailStatusMocked test mocked function
func ImportServersFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportServers test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_servers", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.ImportServers(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// ImportServersFailJSONMocked test mocked function
func ImportServersFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

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
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_servers", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.ImportServers(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ImportVPCsMocked test mocked function
func ImportVPCsMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportVPCs test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_vpcs", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.ImportVPCs(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing VPCs for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "ImportVPCs returned different cloud account")

	return cloudAccountOut
}

// ImportVPCsFailErrMocked test mocked function
func ImportVPCsFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportVPCs test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_vpcs", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.ImportVPCs(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// ImportVPCsFailStatusMocked test mocked function
func ImportVPCsFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportVPCs test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_vpcs", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.ImportVPCs(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// ImportVPCsFailJSONMocked test mocked function
func ImportVPCsFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

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
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_vpcs", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.ImportVPCs(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ImportFloatingIPsMocked test mocked function
func ImportFloatingIPsMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportFloatingIPs test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_floating_ips", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.ImportFloatingIPs(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing floating IPs for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "ImportFloatingIPs returned different cloud account")

	return cloudAccountOut
}

// ImportFloatingIPsFailErrMocked test mocked function
func ImportFloatingIPsFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportFloatingIPs test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_floating_ips", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.ImportFloatingIPs(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// ImportFloatingIPsFailStatusMocked test mocked function
func ImportFloatingIPsFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportFloatingIPs test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_floating_ips", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.ImportFloatingIPs(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// ImportFloatingIPsFailJSONMocked test mocked function
func ImportFloatingIPsFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

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
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_floating_ips", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.ImportFloatingIPs(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ImportVolumesMocked test mocked function
func ImportVolumesMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportVolumes test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_volumes", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.ImportVolumes(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing volumes for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "ImportVolumes returned different cloud account")

	return cloudAccountOut
}

// ImportVolumesFailErrMocked test mocked function
func ImportVolumesFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportVolumes test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_volumes", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.ImportVolumes(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// ImportVolumesFailStatusMocked test mocked function
func ImportVolumesFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportVolumes test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_volumes", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.ImportVolumes(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// ImportVolumesFailJSONMocked test mocked function
func ImportVolumesFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

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
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_volumes", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.ImportVolumes(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ImportKubernetesClustersMocked test mocked function
func ImportKubernetesClustersMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportKubernetesClusters test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_kubernetes_clusters", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.ImportKubernetesClusters(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing kubernetes clusters for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "ImportKubernetesClusters returned different cloud account")

	return cloudAccountOut
}

// ImportKubernetesClustersFailErrMocked test mocked function
func ImportKubernetesClustersFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportKubernetesClusters test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_kubernetes_clusters", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.ImportKubernetesClusters(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// ImportKubernetesClustersFailStatusMocked test mocked function
func ImportKubernetesClustersFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportKubernetesClusters test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_kubernetes_clusters", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.ImportKubernetesClusters(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// ImportKubernetesClustersFailJSONMocked test mocked function
func ImportKubernetesClustersFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

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
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_kubernetes_clusters", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.ImportKubernetesClusters(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ImportPoliciesMocked test mocked function
func ImportPoliciesMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportPolicies test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_policies", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.ImportPolicies(cloudAccountID, mapIn)

	assert.Nil(err, "Error importing policies for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "ImportPolicies returned different cloud account")

	return cloudAccountOut
}

// ImportPoliciesFailErrMocked test mocked function
func ImportPoliciesFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportPolicies test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_policies", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.ImportPolicies(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// ImportPoliciesFailStatusMocked test mocked function
func ImportPoliciesFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewImportCandidateService(cs)
	assert.Nil(err, "Couldn't load ImportCandidate service")
	assert.NotNil(ds, "ImportCandidate service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "ImportPolicies test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_policies", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.ImportPolicies(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// ImportPoliciesFailJSONMocked test mocked function
func ImportPoliciesFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

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
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_policies", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.ImportPolicies(cloudAccountID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}
