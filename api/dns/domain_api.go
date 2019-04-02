package dns

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// DomainService manages domain operations
type DomainService struct {
	concertoService utils.ConcertoService
}

// NewDomainService returns a Concerto domain service
func NewDomainService(concertoService utils.ConcertoService) (*DomainService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &DomainService{
		concertoService: concertoService,
	}, nil
}

// GetDomainList returns the list of domains as an array of Domain
func (dm *DomainService) GetDomainList() (domains []types.Domain, err error) {
	log.Debug("GetDomainList")

	data, status, err := dm.concertoService.Get("/v2/dns/domains")
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
func (dm *DomainService) GetDomain(ID string) (domain *types.Domain, err error) {
	log.Debug("GetDomain")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v2/dns/domains/%s", ID))
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
func (dm *DomainService) CreateDomain(domainVector *map[string]interface{}) (domain *types.Domain, err error) {
	log.Debug("CreateDomain")

	data, status, err := dm.concertoService.Post("/v2/dns/domains/", domainVector)
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

// UpdateDomain updates a domain by its ID
func (dm *DomainService) UpdateDomain(domainVector *map[string]interface{}, ID string) (domain *types.Domain, err error) {
	log.Debug("UpdateDomain")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/v2/dns/domains/%s", ID), domainVector)
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
func (dm *DomainService) DeleteDomain(ID string) (err error) {
	log.Debug("DeleteDomain")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/v2/dns/domains/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// GetDomainRecordList returns a list of domainRecord by domain ID
func (dm *DomainService) GetDomainRecordList(domainID string) (domainRecord *[]types.DomainRecord, err error) {
	log.Debug("ListDomainRecords")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v2/dns/domains/%s/records", domainID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domainRecord); err != nil {
		return nil, err
	}

	return domainRecord, nil
}

// GetDomainRecord returns a domainRecord
func (dm *DomainService) GetDomainRecord(domID string, ID string) (domainRecord *types.DomainRecord, err error) {
	log.Debug("GetDomainRecord")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v2/dns/domains/%s/records/%s", domID, ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domainRecord); err != nil {
		return nil, err
	}

	return domainRecord, nil
}

// CreateDomainRecord returns a list of domainRecord
func (dm *DomainService) CreateDomainRecord(domainRecordVector *map[string]interface{}, domID string) (domainRecord *types.DomainRecord, err error) {
	log.Debug("CreateDomainRecord")

	data, status, err := dm.concertoService.Post(fmt.Sprintf("/v2/dns/domains/%s/records", domID), domainRecordVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domainRecord); err != nil {
		return nil, err
	}

	return domainRecord, nil
}

// UpdateDomainRecord returns a list of domainRecord
func (dm *DomainService) UpdateDomainRecord(domainRecordVector *map[string]interface{}, domID string, ID string) (domainRecord *types.DomainRecord, err error) {
	log.Debug("UpdateDomainRecord")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/v2/dns/domains/%s/records/%s", domID, ID), domainRecordVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domainRecord); err != nil {
		return nil, err
	}

	return domainRecord, nil
}

// DeleteDomainRecord deletes a domain record
func (dm *DomainService) DeleteDomainRecord(domID string, ID string) (err error) {
	log.Debug("DeleteDomainRecord")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/v2/dns/domains/%s/records/%s", domID, ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
