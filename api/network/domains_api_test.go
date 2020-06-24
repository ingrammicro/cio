package network

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDomainServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewDomainService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListDomains(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	ListDomainsMocked(t, domainsIn)
	ListDomainsFailErrMocked(t, domainsIn)
	ListDomainsFailStatusMocked(t, domainsIn)
	ListDomainsFailJSONMocked(t, domainsIn)
}

func TestGetDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range domainsIn {
		GetDomainMocked(t, domainIn)
		GetDomainFailErrMocked(t, domainIn)
		GetDomainFailStatusMocked(t, domainIn)
		GetDomainFailJSONMocked(t, domainIn)
	}
}

func TestCreateDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range domainsIn {
		CreateDomainMocked(t, domainIn)
		CreateDomainFailErrMocked(t, domainIn)
		CreateDomainFailStatusMocked(t, domainIn)
		CreateDomainFailJSONMocked(t, domainIn)
	}
}

func TestDeleteDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range domainsIn {
		DeleteDomainMocked(t, domainIn)
		DeleteDomainFailErrMocked(t, domainIn)
		DeleteDomainFailStatusMocked(t, domainIn)
		DeleteDomainFailJSONMocked(t, domainIn)
	}
}

func TestRetryDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range domainsIn {
		RetryDomainMocked(t, domainIn)
		RetryDomainFailErrMocked(t, domainIn)
		RetryDomainFailStatusMocked(t, domainIn)
		RetryDomainFailJSONMocked(t, domainIn)
	}
}

func TestListRecords(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	recordsIn := testdata.GetRecordData()
	ListRecordsMocked(t, domainsIn[0].ID, recordsIn)
	ListRecordsFailErrMocked(t, domainsIn[0].ID, recordsIn)
	ListRecordsFailStatusMocked(t, domainsIn[0].ID, recordsIn)
	ListRecordsFailJSONMocked(t, domainsIn[0].ID, recordsIn)
}

func TestGetRecord(t *testing.T) {
	recordsIn := testdata.GetRecordData()
	for _, recordIn := range recordsIn {
		GetRecordMocked(t, recordIn)
		GetRecordFailErrMocked(t, recordIn)
		GetRecordFailStatusMocked(t, recordIn)
		GetRecordFailJSONMocked(t, recordIn)
	}
}

func TestCreateRecord(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	recordsIn := testdata.GetRecordData()
	for _, recordIn := range recordsIn {
		CreateRecordMocked(t, domainsIn[0].ID, recordIn)
		CreateRecordFailErrMocked(t, domainsIn[0].ID, recordIn)
		CreateRecordFailStatusMocked(t, domainsIn[0].ID, recordIn)
		CreateRecordFailJSONMocked(t, domainsIn[0].ID, recordIn)
	}
}

func TestUpdateRecord(t *testing.T) {
	recordsIn := testdata.GetRecordData()
	for _, recordIn := range recordsIn {
		UpdateRecordMocked(t, recordIn)
		UpdateRecordFailErrMocked(t, recordIn)
		UpdateRecordFailStatusMocked(t, recordIn)
		UpdateRecordFailJSONMocked(t, recordIn)
	}
}

func TestDeleteRecord(t *testing.T) {
	recordsIn := testdata.GetRecordData()
	for _, recordIn := range recordsIn {
		DeleteRecordMocked(t, recordIn)
		DeleteRecordFailErrMocked(t, recordIn)
		DeleteRecordFailStatusMocked(t, recordIn)
		DeleteRecordFailJSONMocked(t, recordIn)
	}
}

func TestRetryRecord(t *testing.T) {
	recordsIn := testdata.GetRecordData()
	for _, recordIn := range recordsIn {
		RetryRecordMocked(t, recordIn)
		RetryRecordFailErrMocked(t, recordIn)
		RetryRecordFailStatusMocked(t, recordIn)
		RetryRecordFailJSONMocked(t, recordIn)
	}
}
