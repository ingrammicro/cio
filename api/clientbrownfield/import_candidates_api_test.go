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

func TestImportServers(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ImportServersMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportServersFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportServersFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportServersFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestImportVPCs(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ImportVPCsMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportVPCsFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportVPCsFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportVPCsFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestImportFloatingIPs(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ImportFloatingIPsMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportFloatingIPsFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportFloatingIPsFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportFloatingIPsFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestImportVolumes(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ImportVolumesMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportVolumesFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportVolumesFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportVolumesFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestImportKubernetesClusters(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ImportKubernetesClustersMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportKubernetesClustersFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportKubernetesClustersFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportKubernetesClustersFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}

func TestImportPolicies(t *testing.T) {
	cloudAccountsIn := testdata.GetBrownfieldCloudAccountsData()
	for _, cloudAccountIn := range cloudAccountsIn {
		ImportPoliciesMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportPoliciesFailErrMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportPoliciesFailStatusMocked(t, cloudAccountIn, cloudAccountIn.ID)
		ImportPoliciesFailJSONMocked(t, cloudAccountIn, cloudAccountIn.ID)
	}
}
