package clientbrownfield

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportCandidateServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewImportCandidateService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestImportServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, serverIn := range serversIn {
		ImportServerMocked(t, serverIn, cloudAccountsIn[0].ID)
		ImportServerFailErrMocked(t, serverIn, cloudAccountsIn[0].ID)
		ImportServerFailStatusMocked(t, serverIn, cloudAccountsIn[0].ID)
		ImportServerFailJSONMocked(t, serverIn, cloudAccountsIn[0].ID)
	}
}

func TestImportVPC(t *testing.T) {
	vpcsIn := testdata.GetVPCData()
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, vpcIn := range vpcsIn {
		ImportVPCMocked(t, vpcIn, cloudAccountsIn[0].ID)
		ImportVPCFailErrMocked(t, vpcIn, cloudAccountsIn[0].ID)
		ImportVPCFailStatusMocked(t, vpcIn, cloudAccountsIn[0].ID)
		ImportVPCFailJSONMocked(t, vpcIn, cloudAccountsIn[0].ID)
	}
}

func TestImportFloatingIP(t *testing.T) {
	floatingIPsIn := testdata.GetFloatingIPData()
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, floatingIPIn := range floatingIPsIn {
		ImportFloatingIPMocked(t, floatingIPIn, cloudAccountsIn[0].ID)
		ImportFloatingIPFailErrMocked(t, floatingIPIn, cloudAccountsIn[0].ID)
		ImportFloatingIPFailStatusMocked(t, floatingIPIn, cloudAccountsIn[0].ID)
		ImportFloatingIPFailJSONMocked(t, floatingIPIn, cloudAccountsIn[0].ID)
	}
}

func TestImportVolume(t *testing.T) {
	volumesIn := testdata.GetVolumeData()
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, volumeIn := range volumesIn {
		ImportVolumeMocked(t, volumeIn, cloudAccountsIn[0].ID)
		ImportVolumeFailErrMocked(t, volumeIn, cloudAccountsIn[0].ID)
		ImportVolumeFailStatusMocked(t, volumeIn, cloudAccountsIn[0].ID)
		ImportVolumeFailJSONMocked(t, volumeIn, cloudAccountsIn[0].ID)
	}
}