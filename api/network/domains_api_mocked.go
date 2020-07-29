package network

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ListDomainsMocked test mocked function
func ListDomainsMocked(t *testing.T, domainsIn []*types.Domain) []*types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainsIn)
	assert.Nil(err, "Domains test data corrupted")

	// call service
	cs.On("Get", "/network/dns/domains").Return(dIn, 200, nil)
	domainsOut, err := ds.ListDomains()

	assert.Nil(err, "Error getting domains")
	assert.Equal(domainsIn, domainsOut, "ListDomains returned different domains")

	return domainsOut
}

// ListDomainsFailErrMocked test mocked function
func ListDomainsFailErrMocked(t *testing.T, domainsIn []*types.Domain) []*types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainsIn)
	assert.Nil(err, "Domains test data corrupted")

	// call service
	cs.On("Get", "/network/dns/domains").Return(dIn, 200, fmt.Errorf("mocked error"))
	domainsOut, err := ds.ListDomains()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return domainsOut
}

// ListDomainsFailStatusMocked test mocked function
func ListDomainsFailStatusMocked(t *testing.T, domainsIn []*types.Domain) []*types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainsIn)
	assert.Nil(err, "Domains test data corrupted")

	// call service
	cs.On("Get", "/network/dns/domains").Return(dIn, 499, nil)
	domainsOut, err := ds.ListDomains()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return domainsOut
}

// ListDomainsFailJSONMocked test mocked function
func ListDomainsFailJSONMocked(t *testing.T, domainsIn []*types.Domain) []*types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/network/dns/domains").Return(dIn, 200, nil)
	domainsOut, err := ds.ListDomains()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return domainsOut
}

// GetDomainMocked test mocked function
func GetDomainMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/domains/%s", domainIn.ID)).Return(dIn, 200, nil)
	domainOut, err := ds.GetDomain(domainIn.ID)

	assert.Nil(err, "Error getting domain")
	assert.Equal(*domainIn, *domainOut, "GetDomain returned different domain")

	return domainOut
}

// GetDomainFailErrMocked test mocked function
func GetDomainFailErrMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/domains/%s", domainIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	domainOut, err := ds.GetDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return domainOut
}

// GetDomainFailStatusMocked test mocked function
func GetDomainFailStatusMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/domains/%s", domainIn.ID)).Return(dIn, 499, nil)
	domainOut, err := ds.GetDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return domainOut
}

// GetDomainFailJSONMocked test mocked function
func GetDomainFailJSONMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/domains/%s", domainIn.ID)).Return(dIn, 200, nil)
	domainOut, err := ds.GetDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return domainOut
}

// CreateDomainMocked test mocked function
func CreateDomainMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// to json
	dOut, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Post", "/network/dns/domains", mapIn).Return(dOut, 200, nil)
	domainOut, err := ds.CreateDomain(mapIn)

	assert.Nil(err, "Error creating domain")
	assert.Equal(domainIn, domainOut, "CreateDomain returned different domain")

	return domainOut
}

// CreateDomainFailErrMocked test mocked function
func CreateDomainFailErrMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// to json
	dOut, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Post", "/network/dns/domains", mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	domainOut, err := ds.CreateDomain(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return domainOut
}

// CreateDomainFailStatusMocked test mocked function
func CreateDomainFailStatusMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// to json
	dOut, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Post", "/network/dns/domains", mapIn).Return(dOut, 499, nil)
	domainOut, err := ds.CreateDomain(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return domainOut
}

// CreateDomainFailJSONMocked test mocked function
func CreateDomainFailJSONMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/network/dns/domains", mapIn).Return(dIn, 200, nil)
	domainOut, err := ds.CreateDomain(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return domainOut
}

// DeleteDomainMocked test mocked function
func DeleteDomainMocked(t *testing.T, domainIn *types.Domain) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/dns/domains/%s", domainIn.ID)).Return(dIn, 200, nil)
	domainOut, err := ds.DeleteDomain(domainIn.ID)

	assert.Nil(err, "Error deleting domain")
	assert.Equal(domainIn, domainOut, "DeleteDomain returned different domain")

}

// DeleteDomainFailErrMocked test mocked function
func DeleteDomainFailErrMocked(t *testing.T, domainIn *types.Domain) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/dns/domains/%s", domainIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	domainOut, err := ds.DeleteDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteDomainFailStatusMocked test mocked function
func DeleteDomainFailStatusMocked(t *testing.T, domainIn *types.Domain) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/dns/domains/%s", domainIn.ID)).Return(dIn, 499, nil)
	domainOut, err := ds.DeleteDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteDomainFailJSONMocked test mocked function
func DeleteDomainFailJSONMocked(t *testing.T, domainIn *types.Domain) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf("/network/dns/domains/%s", domainIn.ID)).Return(dIn, 200, nil)
	domainOut, err := ds.DeleteDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")
}

// RetryDomainMocked test mocked function
func RetryDomainMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/domains/%s/retry", domainIn.ID), mapIn).Return(dOut, 200, nil)
	domainOut, err := ds.RetryDomain(domainIn.ID)

	assert.Nil(err, "Error retrying domain")
	assert.Equal(domainIn, domainOut, "RetryDomain returned different domain")

	return domainOut
}

// RetryDomainFailErrMocked test mocked function
func RetryDomainFailErrMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/domains/%s/retry", domainIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	domainOut, err := ds.RetryDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return domainOut
}

// RetryDomainFailStatusMocked test mocked function
func RetryDomainFailStatusMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/domains/%s/retry", domainIn.ID), mapIn).Return(dOut, 499, nil)
	domainOut, err := ds.RetryDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return domainOut
}

// RetryDomainFailJSONMocked test mocked function
func RetryDomainFailJSONMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/domains/%s/retry", domainIn.ID), mapIn).Return(dIn, 200, nil)
	domainOut, err := ds.RetryDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return domainOut
}

// ListRecordsMocked test mocked function
func ListRecordsMocked(t *testing.T, domainID string, recordsIn []*types.Record) []*types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordsIn)
	assert.Nil(err, "Records test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/domains/%s/records", domainID)).Return(dIn, 200, nil)
	recordsOut, err := ds.ListRecords(domainID)

	assert.Nil(err, "Error getting records")
	assert.Equal(recordsIn, recordsOut, "ListRecords returned different records")

	return recordsOut
}

// ListRecordsFailErrMocked test mocked function
func ListRecordsFailErrMocked(t *testing.T, domainID string, recordsIn []*types.Record) []*types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordsIn)
	assert.Nil(err, "Records test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/domains/%s/records", domainID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	recordsOut, err := ds.ListRecords(domainID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(recordsOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return recordsOut
}

// ListRecordsFailStatusMocked test mocked function
func ListRecordsFailStatusMocked(t *testing.T, domainID string, recordsIn []*types.Record) []*types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordsIn)
	assert.Nil(err, "Records test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/domains/%s/records", domainID)).Return(dIn, 499, nil)
	recordsOut, err := ds.ListRecords(domainID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(recordsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return recordsOut
}

// ListRecordsFailJSONMocked test mocked function
func ListRecordsFailJSONMocked(t *testing.T, domainID string, recordsIn []*types.Record) []*types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/domains/%s/records", domainID)).Return(dIn, 200, nil)
	recordsOut, err := ds.ListRecords(domainID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(recordsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return recordsOut
}

// GetRecordMocked test mocked function
func GetRecordMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/records/%s", recordIn.ID)).Return(dIn, 200, nil)
	recordOut, err := ds.GetRecord(recordIn.ID)

	assert.Nil(err, "Error getting record")
	assert.Equal(recordIn, recordOut, "GetRecord returned different record")

	return recordOut
}

// GetRecordFailErrMocked test mocked function
func GetRecordFailErrMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/records/%s", recordIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	recordOut, err := ds.GetRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return recordOut
}

// GetRecordFailStatusMocked test mocked function
func GetRecordFailStatusMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/records/%s", recordIn.ID)).Return(dIn, 499, nil)
	recordOut, err := ds.GetRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return recordOut
}

// GetRecordFailJSONMocked test mocked function
func GetRecordFailJSONMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/network/dns/records/%s", recordIn.ID)).Return(dIn, 200, nil)
	recordOut, err := ds.GetRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return recordOut
}

// CreateRecordMocked test mocked function
func CreateRecordMocked(t *testing.T, domainID string, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*recordIn)
	assert.Nil(err, "Record test data corrupted")

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/dns/domains/%s/records", domainID), mapIn).Return(dOut, 200, nil)
	recordOut, err := ds.CreateRecord(domainID, mapIn)

	assert.Nil(err, "Error creating record")
	assert.Equal(recordIn, recordOut, "CreateRecord returned different record")

	return recordOut
}

// CreateRecordFailErrMocked test mocked function
func CreateRecordFailErrMocked(t *testing.T, domainID string, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*recordIn)
	assert.Nil(err, "Record test data corrupted")

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/dns/domains/%s/records", domainID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	recordOut, err := ds.CreateRecord(domainID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return recordOut
}

// CreateRecordFailStatusMocked test mocked function
func CreateRecordFailStatusMocked(t *testing.T, domainID string, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*recordIn)
	assert.Nil(err, "Record test data corrupted")

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/network/dns/domains/%s/records", domainID), mapIn).Return(dOut, 499, nil)
	recordOut, err := ds.CreateRecord(domainID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return recordOut
}

// CreateRecordFailJSONMocked test mocked function
func CreateRecordFailJSONMocked(t *testing.T, domainID string, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*recordIn)
	assert.Nil(err, "Record test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf("/network/dns/domains/%s/records", domainID), mapIn).Return(dIn, 200, nil)
	recordOut, err := ds.CreateRecord(domainID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return recordOut
}

// UpdateRecordMocked test mocked function
func UpdateRecordMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*recordIn)
	assert.Nil(err, "Record test data corrupted")

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/records/%s", recordIn.ID), mapIn).Return(dOut, 200, nil)
	recordOut, err := ds.UpdateRecord(recordIn.ID, mapIn)

	assert.Nil(err, "Error updating record")
	assert.Equal(recordIn, recordOut, "UpdateRecord returned different record")

	return recordOut
}

// UpdateRecordFailErrMocked test mocked function
func UpdateRecordFailErrMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*recordIn)
	assert.Nil(err, "Record test data corrupted")

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/records/%s", recordIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	recordOut, err := ds.UpdateRecord(recordIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return recordOut
}

// UpdateRecordFailStatusMocked test mocked function
func UpdateRecordFailStatusMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*recordIn)
	assert.Nil(err, "Record test data corrupted")

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/records/%s", recordIn.ID), mapIn).Return(dOut, 499, nil)
	recordOut, err := ds.UpdateRecord(recordIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return recordOut
}

// UpdateRecordFailJSONMocked test mocked function
func UpdateRecordFailJSONMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*recordIn)
	assert.Nil(err, "Record test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/records/%s", recordIn.ID), mapIn).Return(dIn, 200, nil)
	recordOut, err := ds.UpdateRecord(recordIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return recordOut
}

// DeleteRecordMocked test mocked function
func DeleteRecordMocked(t *testing.T, recordIn *types.Record) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/dns/records/%s", recordIn.ID)).Return(dIn, 200, nil)
	recordOut, err := ds.DeleteRecord(recordIn.ID)

	assert.Nil(err, "Error deleting record")
	assert.Equal(recordIn, recordOut, "DeleteRecord returned different record")
}

// DeleteRecordFailErrMocked test mocked function
func DeleteRecordFailErrMocked(t *testing.T, recordIn *types.Record) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/dns/records/%s", recordIn.ID)).Return(dIn, 200, fmt.Errorf("mocked error"))
	recordOut, err := ds.DeleteRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteRecordFailStatusMocked test mocked function
func DeleteRecordFailStatusMocked(t *testing.T, recordIn *types.Record) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/network/dns/records/%s", recordIn.ID)).Return(dIn, 499, nil)
	recordOut, err := ds.DeleteRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// DeleteRecordFailJSONMocked test mocked function
func DeleteRecordFailJSONMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Delete", fmt.Sprintf("/network/dns/records/%s", recordIn.ID)).Return(dIn, 200, nil)
	recordOut, err := ds.DeleteRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return recordOut
}

// RetryRecordMocked test mocked function
func RetryRecordMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/records/%s/retry", recordIn.ID), mapIn).Return(dOut, 200, nil)
	recordOut, err := ds.RetryRecord(recordIn.ID)

	assert.Nil(err, "Error retrying record")
	assert.Equal(recordIn, recordOut, "RetryRecord returned different record")

	return recordOut
}

// RetryRecordFailErrMocked test mocked function
func RetryRecordFailErrMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/records/%s/retry", recordIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("mocked error"))
	recordOut, err := ds.RetryRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return recordOut
}

// RetryRecordFailStatusMocked test mocked function
func RetryRecordFailStatusMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	mapIn := new(map[string]interface{})

	// to json
	dOut, err := json.Marshal(recordIn)
	assert.Nil(err, "Record test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/records/%s/retry", recordIn.ID), mapIn).Return(dOut, 499, nil)
	recordOut, err := ds.RetryRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return recordOut
}

// RetryRecordFailJSONMocked test mocked function
func RetryRecordFailJSONMocked(t *testing.T, recordIn *types.Record) *types.Record {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	mapIn := new(map[string]interface{})

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/network/dns/records/%s/retry", recordIn.ID), mapIn).Return(dIn, 200, nil)
	recordOut, err := ds.RetryRecord(recordIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(recordOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return recordOut
}
