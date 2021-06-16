// Copyright (c) 2017-2021 Ingram Micro Inc.

package clientbrownfield

import (
	"testing"

	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
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
