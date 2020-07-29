package network

import (
	"github.com/ingrammicro/cio/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCertificateServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCertificateService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestListCertificates(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	certificatesIn := testdata.GetCertificateData()
	ListCertificatesMocked(t, loadBalancersIn[0].ID, certificatesIn)
	ListCertificatesFailErrMocked(t, loadBalancersIn[0].ID, certificatesIn)
	ListCertificatesFailStatusMocked(t, loadBalancersIn[0].ID, certificatesIn)
	ListCertificatesFailJSONMocked(t, loadBalancersIn[0].ID, certificatesIn)
}

func TestGetCertificate(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	certificatesIn := testdata.GetCertificateData()
	for _, certificateIn := range certificatesIn {
		GetCertificateMocked(t, loadBalancersIn[0].ID, certificateIn)
		GetCertificateFailErrMocked(t, loadBalancersIn[0].ID, certificateIn)
		GetCertificateFailStatusMocked(t, loadBalancersIn[0].ID, certificateIn)
		GetCertificateFailJSONMocked(t, loadBalancersIn[0].ID, certificateIn)
	}
}

func TestCreateCertificate(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	certificatesIn := testdata.GetCertificateData()
	for _, certificateIn := range certificatesIn {
		CreateCertificateMocked(t, loadBalancersIn[0].ID, certificateIn)
		CreateCertificateFailErrMocked(t, loadBalancersIn[0].ID, certificateIn)
		CreateCertificateFailStatusMocked(t, loadBalancersIn[0].ID, certificateIn)
		CreateCertificateFailJSONMocked(t, loadBalancersIn[0].ID, certificateIn)
	}
}

func TestUpdateCertificate(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	certificatesIn := testdata.GetCertificateData()
	for _, certificateIn := range certificatesIn {
		UpdateCertificateMocked(t, loadBalancersIn[0].ID, certificateIn)
		UpdateCertificateFailErrMocked(t, loadBalancersIn[0].ID, certificateIn)
		UpdateCertificateFailStatusMocked(t, loadBalancersIn[0].ID, certificateIn)
		UpdateCertificateFailJSONMocked(t, loadBalancersIn[0].ID, certificateIn)
	}
}

func TestDeleteCertificate(t *testing.T) {
	loadBalancersIn := testdata.GetLoadBalancerData()
	certificatesIn := testdata.GetCertificateData()
	for _, certificateIn := range certificatesIn {
		DeleteCertificateMocked(t, loadBalancersIn[0].ID, certificateIn)
		DeleteCertificateFailErrMocked(t, loadBalancersIn[0].ID, certificateIn)
		DeleteCertificateFailStatusMocked(t, loadBalancersIn[0].ID, certificateIn)
	}
}
