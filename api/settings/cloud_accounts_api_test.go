package settings

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewCloudAccountServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudAccountService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetCloudAccountList(t *testing.T) {
	cloudAccountsIn := testdata.GetCloudAccountData()
	GetCloudAccountListMocked(t, cloudAccountsIn)
	GetCloudAccountListFailErrMocked(t, cloudAccountsIn)
	GetCloudAccountListFailStatusMocked(t, cloudAccountsIn)
	GetCloudAccountListFailJSONMocked(t, cloudAccountsIn)
}

func TestGetCloudAccount(t *testing.T) {
	cloudAccountsIn := testdata.GetCloudAccountData()
	for _, cloudAccountIn := range cloudAccountsIn {
		GetCloudAccountMocked(t, cloudAccountIn)
		GetCloudAccountFailErrMocked(t, cloudAccountIn)
		GetCloudAccountFailStatusMocked(t, cloudAccountIn)
		GetCloudAccountFailJSONMocked(t, cloudAccountIn)
	}
}
