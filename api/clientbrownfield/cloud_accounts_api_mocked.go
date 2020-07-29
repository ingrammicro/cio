package clientbrownfield

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ListBrownfieldCloudAccountsMocked test mocked function
func ListBrownfieldCloudAccountsMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccounts test data corrupted")

	// call service
	cs.On("Get", "/brownfield/cloud_accounts").Return(dIn, 200, nil)
	cloudAccountsOut, err := ds.ListBrownfieldCloudAccounts()
	assert.Nil(err, "Error getting brownfield cloud account list")
	assert.Equal(cloudAccountsIn, cloudAccountsOut, "ListBrownfieldCloudAccounts returned different cloud accounts")

	return cloudAccountsOut
}

// ListBrownfieldCloudAccountsFailErrMocked test mocked function
func ListBrownfieldCloudAccountsFailErrMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccounts test data corrupted")

	// call service
	cs.On("Get", "/brownfield/cloud_accounts").Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudAccountsOut, err := ds.ListBrownfieldCloudAccounts()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountsOut
}

// ListBrownfieldCloudAccountsFailStatusMocked test mocked function
func ListBrownfieldCloudAccountsFailStatusMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccounts test data corrupted")

	// call service
	cs.On("Get", "/brownfield/cloud_accounts").Return(dIn, 499, nil)
	cloudAccountsOut, err := ds.ListBrownfieldCloudAccounts()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountsOut
}

// ListBrownfieldCloudAccountsFailJSONMocked test mocked function
func ListBrownfieldCloudAccountsFailJSONMocked(t *testing.T, cloudAccountsIn []*types.CloudAccount) []*types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/brownfield/cloud_accounts").Return(dIn, 200, nil)
	cloudAccountsOut, err := ds.ListBrownfieldCloudAccounts()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountsOut
}

// GetBrownfieldCloudAccountMocked test mocked function
func GetBrownfieldCloudAccountMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s", cloudAccountID)).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.GetBrownfieldCloudAccount(cloudAccountID)
	assert.Nil(err, "Error getting brownfield cloud account")
	assert.Equal(*cloudAccountIn, *cloudAccountOut, "GetBrownfieldCloudAccount returned different cloud account")

	return cloudAccountOut
}

// GetBrownfieldCloudAccountFailErrMocked test mocked function
func GetBrownfieldCloudAccountFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s", cloudAccountID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.GetBrownfieldCloudAccount(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// GetBrownfieldCloudAccountFailStatusMocked test mocked function
func GetBrownfieldCloudAccountFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s", cloudAccountID)).Return(dIn, 499, nil)
	cloudAccountOut, err := ds.GetBrownfieldCloudAccount(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// GetBrownfieldCloudAccountFailJSONMocked test mocked function
func GetBrownfieldCloudAccountFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s", cloudAccountID)).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.GetBrownfieldCloudAccount(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// DiscoverServersMocked test mocked function
func DiscoverServersMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.DiscoverServers(cloudAccountID)

	assert.Nil(err, "Error discovering servers for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "DiscoverServers returned different cloud account")

	return cloudAccountOut
}

// DiscoverServersFailErrMocked test mocked function
func DiscoverServersFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.DiscoverServers(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// DiscoverServersFailStatusMocked test mocked function
func DiscoverServersFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.DiscoverServers(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// DiscoverServersFailJSONMocked test mocked function
func DiscoverServersFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.DiscoverServers(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ListServersMocked test mocked function
func ListServersMocked(t *testing.T, serverImportCandidatesIn []*types.ServerImportCandidate, cloudAccountID string) []*types.ServerImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(serverImportCandidatesIn)
	assert.Nil(err, "ServersImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_candidates", cloudAccountID)).Return(dIn, 200, nil)
	serverImportCandidatesOut, err := ds.ListServers(cloudAccountID)

	assert.Nil(err, "Error getting server import candidates")
	assert.Equal(serverImportCandidatesIn, serverImportCandidatesOut, "ListServers returned different server import candidates")

	return serverImportCandidatesOut
}

// ListServersFailErrMocked test mocked function
func ListServersFailErrMocked(t *testing.T, serverImportCandidatesIn []*types.ServerImportCandidate, cloudAccountID string) []*types.ServerImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(serverImportCandidatesIn)
	assert.Nil(err, "ServersImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_candidates", cloudAccountID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	serverImportCandidatesOut, err := ds.ListServers(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverImportCandidatesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return serverImportCandidatesOut
}

// ListServersFailStatusMocked test mocked function
func ListServersFailStatusMocked(t *testing.T, serverImportCandidatesIn []*types.ServerImportCandidate, cloudAccountID string) []*types.ServerImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(serverImportCandidatesIn)
	assert.Nil(err, "ServersImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_candidates", cloudAccountID)).Return(dIn, 499, nil)
	serverImportCandidatesOut, err := ds.ListServers(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverImportCandidatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverImportCandidatesOut
}

// ListServersFailJSONMocked test mocked function
func ListServersFailJSONMocked(t *testing.T, serverImportCandidatesIn []*types.ServerImportCandidate, cloudAccountID string) []*types.ServerImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/import_candidates", cloudAccountID)).Return(dIn, 200, nil)
	serverImportCandidatesOut, err := ds.ListServers(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverImportCandidatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverImportCandidatesOut
}

// DiscoverVPCsMocked test mocked function
func DiscoverVPCsMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_vpcs", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.DiscoverVPCs(cloudAccountID)

	assert.Nil(err, "Error discovering VPCs for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "DiscoverVPCs returned different cloud account")

	return cloudAccountOut
}

// DiscoverVPCsFailErrMocked test mocked function
func DiscoverVPCsFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_vpcs", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.DiscoverVPCs(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// DiscoverVPCsFailStatusMocked test mocked function
func DiscoverVPCsFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_vpcs", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.DiscoverVPCs(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// DiscoverVPCsFailJSONMocked test mocked function
func DiscoverVPCsFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_vpcs", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.DiscoverVPCs(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ListVPCsMocked test mocked function
func ListVPCsMocked(t *testing.T, vpcImportCandidatesIn []*types.VpcImportCandidate, cloudAccountID string) []*types.VpcImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(vpcImportCandidatesIn)
	assert.Nil(err, "VpcImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/vpc_import_candidates", cloudAccountID)).Return(dIn, 200, nil)

	vpcImportCandidatesOut, err := ds.ListVPCs(cloudAccountID)
	assert.Nil(err, "Error getting VPCs import candidates")
	assert.Equal(vpcImportCandidatesIn, vpcImportCandidatesOut, "ListVPCsMocked returned different VPC import candidates")

	return vpcImportCandidatesOut
}

// ListVPCsFailErrMocked test mocked function
func ListVPCsFailErrMocked(t *testing.T, vpcImportCandidatesIn []*types.VpcImportCandidate, cloudAccountID string) []*types.VpcImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(vpcImportCandidatesIn)
	assert.Nil(err, "VpcImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/vpc_import_candidates", cloudAccountID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	vpcImportCandidatesOut, err := ds.ListVPCs(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(vpcImportCandidatesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return vpcImportCandidatesOut
}

// ListVPCsFailStatusMocked test mocked function
func ListVPCsFailStatusMocked(t *testing.T, vpcImportCandidatesIn []*types.VpcImportCandidate, cloudAccountID string) []*types.VpcImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(vpcImportCandidatesIn)
	assert.Nil(err, "VPCImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/vpc_import_candidates", cloudAccountID)).Return(dIn, 499, nil)
	vpcImportCandidatesOut, err := ds.ListVPCs(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(vpcImportCandidatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return vpcImportCandidatesOut
}

// ListVPCsFailJSONMocked test mocked function
func ListVPCsFailJSONMocked(t *testing.T, vpcImportCandidatesIn []*types.VpcImportCandidate, cloudAccountID string) []*types.VpcImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/vpc_import_candidates", cloudAccountID)).Return(dIn, 200, nil)
	vpcImportCandidatesOut, err := ds.ListVPCs(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(vpcImportCandidatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return vpcImportCandidatesOut
}

// DiscoverFloatingIPsMocked test mocked function
func DiscoverFloatingIPsMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_floating_ips", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.DiscoverFloatingIPs(cloudAccountID)

	assert.Nil(err, "Error discovering FloatingIPs for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "DiscoverFloatingIPs returned different cloud account")

	return cloudAccountOut
}

// DiscoverFloatingIPsFailErrMocked test mocked function
func DiscoverFloatingIPsFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_floating_ips", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.DiscoverFloatingIPs(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// DiscoverFloatingIPsFailStatusMocked test mocked function
func DiscoverFloatingIPsFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_floating_ips", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.DiscoverFloatingIPs(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// DiscoverFloatingIPsFailJSONMocked test mocked function
func DiscoverFloatingIPsFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_floating_ips", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.DiscoverFloatingIPs(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ListFloatingIPsMocked test mocked function
func ListFloatingIPsMocked(t *testing.T, floatingIPImportCandidatesIn []*types.FloatingIPImportCandidate, cloudAccountID string) []*types.FloatingIPImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPImportCandidatesIn)
	assert.Nil(err, "FloatingIPImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/floating_ip_import_candidates", cloudAccountID)).Return(dIn, 200, nil)
	floatingIPImportCandidatesOut, err := ds.ListFloatingIPs(cloudAccountID)

	assert.Nil(err, "Error getting floating IPs import candidates")
	assert.Equal(floatingIPImportCandidatesIn, floatingIPImportCandidatesOut, "ListFloatingIPs returned different floating IP import candidates")

	return floatingIPImportCandidatesOut
}

// ListFloatingIPsFailErrMocked test mocked function
func ListFloatingIPsFailErrMocked(t *testing.T, floatingIPImportCandidatesIn []*types.FloatingIPImportCandidate, cloudAccountID string) []*types.FloatingIPImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPImportCandidatesIn)
	assert.Nil(err, "FloatingIPImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/floating_ip_import_candidates", cloudAccountID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	floatingIPImportCandidatesOut, err := ds.ListFloatingIPs(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(floatingIPImportCandidatesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return floatingIPImportCandidatesOut
}

// ListFloatingIPsFailStatusMocked test mocked function
func ListFloatingIPsFailStatusMocked(t *testing.T, floatingIPImportCandidatesIn []*types.FloatingIPImportCandidate, cloudAccountID string) []*types.FloatingIPImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(floatingIPImportCandidatesIn)
	assert.Nil(err, "FloatingIPImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/floating_ip_import_candidates", cloudAccountID)).Return(dIn, 499, nil)
	floatingIPImportCandidatesOut, err := ds.ListFloatingIPs(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(floatingIPImportCandidatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return floatingIPImportCandidatesOut
}

// ListFloatingIPsFailJSONMocked test mocked function
func ListFloatingIPsFailJSONMocked(t *testing.T, floatingIPImportCandidatesIn []*types.FloatingIPImportCandidate, cloudAccountID string) []*types.FloatingIPImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/floating_ip_import_candidates", cloudAccountID)).Return(dIn, 200, nil)
	floatingIPImportCandidatesOut, err := ds.ListFloatingIPs(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(floatingIPImportCandidatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return floatingIPImportCandidatesOut
}

// DiscoverVolumesMocked test mocked function
func DiscoverVolumesMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_volumes", cloudAccountID), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := ds.DiscoverVolumes(cloudAccountID)

	assert.Nil(err, "Error discovering volumes for cloud account")
	assert.Equal(cloudAccountIn, cloudAccountOut, "DiscoverVolumes returned different cloud account")

	return cloudAccountOut
}

// DiscoverVolumesFailErrMocked test mocked function
func DiscoverVolumesFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_volumes", cloudAccountID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	cloudAccountOut, err := ds.DiscoverVolumes(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return cloudAccountOut
}

// DiscoverVolumesFailStatusMocked test mocked function
func DiscoverVolumesFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_volumes", cloudAccountID), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := ds.DiscoverVolumes(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// DiscoverVolumesFailJSONMocked test mocked function
func DiscoverVolumesFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount, cloudAccountID string) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/brownfield/cloud_accounts/%s/discover_volumes", cloudAccountID), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := ds.DiscoverVolumes(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// ListVolumesMocked test mocked function
func ListVolumesMocked(t *testing.T, volumeImportCandidatesIn []*types.VolumeImportCandidate, cloudAccountID string) []*types.VolumeImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(volumeImportCandidatesIn)
	assert.Nil(err, "VolumeImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/volume_import_candidates", cloudAccountID)).Return(dIn, 200, nil)
	volumeImportCandidatesOut, err := ds.ListVolumes(cloudAccountID)

	assert.Nil(err, "Error getting volumes import candidates")
	assert.Equal(volumeImportCandidatesIn, volumeImportCandidatesOut, "ListVolumes returned different volume import candidates")

	return volumeImportCandidatesOut
}

// ListVolumesFailErrMocked test mocked function
func ListVolumesFailErrMocked(t *testing.T, volumeImportCandidatesIn []*types.VolumeImportCandidate, cloudAccountID string) []*types.VolumeImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(volumeImportCandidatesIn)
	assert.Nil(err, "VolumeImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/volume_import_candidates", cloudAccountID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	volumeImportCandidatesOut, err := ds.ListVolumes(cloudAccountID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(volumeImportCandidatesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return volumeImportCandidatesOut
}

// ListVolumesFailStatusMocked test mocked function
func ListVolumesFailStatusMocked(t *testing.T, volumeImportCandidatesIn []*types.VolumeImportCandidate, cloudAccountID string) []*types.VolumeImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(volumeImportCandidatesIn)
	assert.Nil(err, "VolumeImportCandidates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/volume_import_candidates", cloudAccountID)).Return(dIn, 499, nil)
	volumeImportCandidatesOut, err := ds.ListVolumes(cloudAccountID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(volumeImportCandidatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return volumeImportCandidatesOut
}

// ListVolumesFailJSONMocked test mocked function
func ListVolumesFailJSONMocked(t *testing.T, volumeImportCandidatesIn []*types.VolumeImportCandidate, cloudAccountID string) []*types.VolumeImportCandidate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewBrownfieldCloudAccountService(cs)
	assert.Nil(err, "Couldn't load brownfieldCloudAccount service")
	assert.NotNil(ds, "BrownfieldCloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/brownfield/cloud_accounts/%s/volume_import_candidates", cloudAccountID)).Return(dIn, 200, nil)
	volumeImportCandidatesOut, err := ds.ListVolumes(cloudAccountID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(volumeImportCandidatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return volumeImportCandidatesOut
}
