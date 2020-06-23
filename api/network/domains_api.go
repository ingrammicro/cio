package network

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// DomainService manages DNS domain and record operations
type DomainService struct {
	concertoService utils.ConcertoService
}

// NewDomainService returns a Concerto Domain service
func NewDomainService(concertoService utils.ConcertoService) (*DomainService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &DomainService{
		concertoService: concertoService,
	}, nil
}

// ListDomains returns the list of domains as an array of Domain
func (ds *DomainService) ListDomains() (domains []*types.Domain, err error) {
	log.Debug("ListDomains")

	data, status, err := ds.concertoService.Get("/network/dns/domains")

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domains); err != nil {
		return nil, err
	}

	return domains, nil
}

// GetDomain returns a domain by its ID
func (ds *DomainService) GetDomain(domainID string) (domain *types.Domain, err error) {
	log.Debug("GetDomain")

	data, status, err := ds.concertoService.Get(fmt.Sprintf("/network/dns/domains/%s", domainID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domain); err != nil {
		return nil, err
	}

	return domain, nil
}

// CreateDomain creates a domain
func (ds *DomainService) CreateDomain(domainParams *map[string]interface{}) (domain *types.Domain, err error) {
	log.Debug("CreateDomain")

	data, status, err := ds.concertoService.Post("/network/dns/domains", domainParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domain); err != nil {
		return nil, err
	}

	return domain, nil
}

// DeleteDomain deletes a domain by its ID
func (ds *DomainService) DeleteDomain(domainID string) (domain *types.Domain, err error) {
	log.Debug("DeleteDomain")

	data, status, err := ds.concertoService.Delete(fmt.Sprintf("/network/dns/domains/%s", domainID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domain); err != nil {
		return nil, err
	}

	return domain, nil
}

// RetryDomain retries a domain by its ID
func (ds *DomainService) RetryDomain(domainID string, domainParams *map[string]interface{}) (domain *types.Domain, err error) {
	log.Debug("RetryDomain")

	data, status, err := ds.concertoService.Put(fmt.Sprintf("/network/dns/domains/%s/retry", domainID), domainParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domain); err != nil {
		return nil, err
	}

	return domain, nil
}

// ListRecords returns the list of records as an array of Record for given domain
func (ds *DomainService) ListRecords(domainID string) (records []*types.Record, err error) {
	log.Debug("ListRecords")

	data, status, err := ds.concertoService.Get(fmt.Sprintf("/network/dns/domains/%s/records", domainID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &records); err != nil {
		return nil, err
	}

	return records, nil
}

// GetRecord returns a record by its ID
func (ds *DomainService) GetRecord(recordID string) (record *types.Record, err error) {
	log.Debug("GetRecord")

	data, status, err := ds.concertoService.Get(fmt.Sprintf("/network/dns/records/%s", recordID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	return record, nil
}

// CreateRecord creates a record in a domain
func (ds *DomainService) CreateRecord(domainID string, recordParams *map[string]interface{}) (record *types.Record, err error) {
	log.Debug("CreateRecord")

	data, status, err := ds.concertoService.Post(fmt.Sprintf("/network/dns/domains/%s/records", domainID), recordParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	return record, nil
}

// UpdateRecord updates a record by its ID
func (ds *DomainService) UpdateRecord(recordID string, recordParams *map[string]interface{}) (record *types.Record, err error) {
	log.Debug("UpdateRecord")

	data, status, err := ds.concertoService.Put(fmt.Sprintf("/network/dns/records/%s", recordID), recordParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	return record, nil
}

// DeleteRecord deletes a record by its ID
func (ds *DomainService) DeleteRecord(recordID string) (record *types.Record, err error) {
	log.Debug("DeleteRecord")

	data, status, err := ds.concertoService.Delete(fmt.Sprintf("/network/dns/records/%s", recordID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	return record, nil
}

// RetryRecord retries a record by its ID
func (ds *DomainService) RetryRecord(recordID string, recordParams *map[string]interface{}) (record *types.Record, err error) {
	log.Debug("RetryRecord")

	data, status, err := ds.concertoService.Put(fmt.Sprintf("/network/dns/records/%s/retry", recordID), recordParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	return record, nil
}
