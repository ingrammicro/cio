package clientbrownfield

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBrownfieldCloudAccountServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewBrownfieldCloudAccountService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListBrownfieldCloudAccounts(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	ListBrownfieldCloudAccountsMocked(t, cloudAccountsIn)
	ListBrownfieldCloudAccountsFailErrMocked(t, cloudAccountsIn)
	ListBrownfieldCloudAccountsFailStatusMocked(t, cloudAccountsIn)
	ListBrownfieldCloudAccountsFailJSONMocked(t, cloudAccountsIn)
}

func TestGetBrownfieldCloudAccount(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		GetBrownfieldCloudAccountMocked(t, cloudAccountIn, cloudAccountIn.ID)
		GetBrownfieldCloudAccountFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		GetBrownfieldCloudAccountFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		GetBrownfieldCloudAccountFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestDiscoverServers(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		DiscoverServersMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverServersFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverServersFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverServersFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestListServers(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	serverImportCandidatesIn := testdata.GetBrownfieldServerImportCandidatesData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ListServersMocked(t, serverImportCandidatesIn, cloudAccountIn.ID)
		ListServersFailErrMocked(t, serverImportCandidatesIn, cloudAccountIn.ID)
		ListServersFailStatusMocked(t, serverImportCandidatesIn, cloudAccountIn.ID)
		ListServersFailJSONMocked(t, serverImportCandidatesIn, cloudAccountIn.ID)
	}
}

func TestDiscoverVPCs(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		DiscoverVPCsMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverVPCsFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverVPCsFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverVPCsFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestListVPCs(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	vpcImportCandidatesIn := testdata.GetBrownfieldVPCImportCandidatesData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ListVPCsMocked(t, vpcImportCandidatesIn, cloudAccountIn.ID)
		ListVPCsFailErrMocked(t, vpcImportCandidatesIn, cloudAccountIn.ID)
		ListVPCsFailStatusMocked(t, vpcImportCandidatesIn, cloudAccountIn.ID)
		ListVPCsFailJSONMocked(t, vpcImportCandidatesIn, cloudAccountIn.ID)
	}
}

func TestDiscoverFloatingIPs(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		DiscoverFloatingIPsMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverFloatingIPsFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverFloatingIPsFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverFloatingIPsFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestListFloatingIPs(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	floatingIPImportCandidatesIn := testdata.GetBrownfieldFloatingIPImportCandidatesData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ListFloatingIPsMocked(t, floatingIPImportCandidatesIn, cloudAccountIn.ID)
		ListFloatingIPsFailErrMocked(t, floatingIPImportCandidatesIn, cloudAccountIn.ID)
		ListFloatingIPsFailStatusMocked(t, floatingIPImportCandidatesIn, cloudAccountIn.ID)
		ListFloatingIPsFailJSONMocked(t, floatingIPImportCandidatesIn, cloudAccountIn.ID)
	}
}

func TestDiscoverVolumes(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		DiscoverVolumesMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverVolumesFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverVolumesFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		DiscoverVolumesFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestListVolumes(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	volumeImportCandidatesIn := testdata.GetBrownfieldVolumeImportCandidatesData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ListVolumesMocked(t, volumeImportCandidatesIn, cloudAccountIn.ID)
		ListVolumesFailErrMocked(t, volumeImportCandidatesIn, cloudAccountIn.ID)
		ListVolumesFailStatusMocked(t, volumeImportCandidatesIn, cloudAccountIn.ID)
		ListVolumesFailJSONMocked(t, volumeImportCandidatesIn, cloudAccountIn.ID)
	}
}
