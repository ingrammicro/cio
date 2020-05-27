package network

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// CertificateService manages certificate operations
type CertificateService struct {
	concertoService utils.ConcertoService
}

// NewCertificateService returns a Concerto certificate service
func NewCertificateService(concertoService utils.ConcertoService) (*CertificateService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CertificateService{
		concertoService: concertoService,
	}, nil
}

// ListCertificates returns the list of certificates in a load balancer by its ID, as an array of Certificate
func (cs *CertificateService) ListCertificates(loadBalancerID string) (certificates []*types.Certificate, err error) {
	log.Debug("ListCertificates")

	data, status, err := cs.concertoService.Get(fmt.Sprintf("/network/load_balancers/%s/certificates", loadBalancerID))

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &certificates); err != nil {
		return nil, err
	}

	return certificates, nil
}

// GetCertificate returns a certificate by its ID
func (cs *CertificateService) GetCertificate(loadBalancerID string, certificateID string) (certificate *types.Certificate, err error) {
	log.Debug("GetCertificate")

	data, status, err := cs.concertoService.Get(fmt.Sprintf("/network/load_balancers/%s/certificates/%s", loadBalancerID, certificateID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &certificate); err != nil {
		return nil, err
	}

	return certificate, nil
}

// CreateCertificate creates a certificate in a load balancer by its ID
func (cs *CertificateService) CreateCertificate(loadBalancerID string, certificateParams *map[string]interface{}) (certificate *types.Certificate, err error) {
	log.Debug("CreateCertificate")

	data, status, err := cs.concertoService.Post(fmt.Sprintf("/network/load_balancers/%s/certificates", loadBalancerID), certificateParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &certificate); err != nil {
		return nil, err
	}

	return certificate, nil
}

// UpdateCertificate updates a certificate by its ID
func (cs *CertificateService) UpdateCertificate(loadBalancerID string, certificateID string, certificateParams *map[string]interface{}) (certificate *types.Certificate, err error) {
	log.Debug("UpdateCertificate")

	data, status, err := cs.concertoService.Put(fmt.Sprintf("/network/load_balancers/%s/certificates/%s", loadBalancerID, certificateID), certificateParams)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &certificate); err != nil {
		return nil, err
	}

	return certificate, nil
}

// DeleteCertificate deletes a certificate by its ID
func (cs *CertificateService) DeleteCertificate(loadBalancerID string, certificateID string) (err error) {
	log.Debug("DeleteCertificate")

	data, status, err := cs.concertoService.Delete(fmt.Sprintf("/network/load_balancers/%s/certificates/%s", loadBalancerID, certificateID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
