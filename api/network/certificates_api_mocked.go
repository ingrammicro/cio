// Copyright (c) 2017-2021 Ingram Micro Inc.

package network

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

// ListCertificatesMocked test mocked function
func ListCertificatesMocked(
	t *testing.T,
	loadBalancerID string,
	certificatesIn []*types.Certificate,
) []*types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificatesIn)
	assert.Nil(err, "Certificates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerCertificates, loadBalancerID)).Return(dIn, 200, nil)
	certificatesOut, err := ds.ListCertificates(loadBalancerID)

	assert.Nil(err, "Error getting certificates")
	assert.Equal(certificatesIn, certificatesOut, "ListCertificates returned different certificates")

	return certificatesOut
}

// ListCertificatesFailErrMocked test mocked function
func ListCertificatesFailErrMocked(
	t *testing.T,
	loadBalancerID string,
	certificatesIn []*types.Certificate,
) []*types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificatesIn)
	assert.Nil(err, "Certificates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerCertificates, loadBalancerID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	certificatesOut, err := ds.ListCertificates(loadBalancerID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(certificatesOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return certificatesOut
}

// ListCertificatesFailStatusMocked test mocked function
func ListCertificatesFailStatusMocked(
	t *testing.T,
	loadBalancerID string,
	certificatesIn []*types.Certificate,
) []*types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificatesIn)
	assert.Nil(err, "Certificates test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerCertificates, loadBalancerID)).Return(dIn, 499, nil)
	certificatesOut, err := ds.ListCertificates(loadBalancerID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(certificatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return certificatesOut
}

// ListCertificatesFailJSONMocked test mocked function
func ListCertificatesFailJSONMocked(
	t *testing.T,
	loadBalancerID string,
	certificatesIn []*types.Certificate,
) []*types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerCertificates, loadBalancerID)).Return(dIn, 200, nil)
	certificatesOut, err := ds.ListCertificates(loadBalancerID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(certificatesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return certificatesOut
}

// GetCertificateMocked test mocked function
func GetCertificateMocked(t *testing.T, loadBalancerID string, certificateIn *types.Certificate) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID)).
		Return(dIn, 200, nil)
	certificateOut, err := ds.GetCertificate(loadBalancerID, certificateIn.ID)

	assert.Nil(err, "Error getting certificate")
	assert.Equal(*certificateIn, *certificateOut, "GetCertificate returned different certificate")

	return certificateOut
}

// GetCertificateFailErrMocked test mocked function
func GetCertificateFailErrMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	certificateOut, err := ds.GetCertificate(loadBalancerID, certificateIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return certificateOut
}

// GetCertificateFailStatusMocked test mocked function
func GetCertificateFailStatusMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID)).
		Return(dIn, 499, nil)
	certificateOut, err := ds.GetCertificate(loadBalancerID, certificateIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return certificateOut
}

// GetCertificateFailJSONMocked test mocked function
func GetCertificateFailJSONMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID)).
		Return(dIn, 200, nil)
	certificateOut, err := ds.GetCertificate(loadBalancerID, certificateIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return certificateOut
}

// CreateCertificateMocked test mocked function
func CreateCertificateMocked(t *testing.T, loadBalancerID string, certificateIn *types.Certificate) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// to json
	dOut, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkLoadBalancerCertificates, loadBalancerID), mapIn).Return(dOut, 200, nil)
	certificateOut, err := ds.CreateCertificate(loadBalancerID, mapIn)

	assert.Nil(err, "Error creating certificate")
	assert.Equal(certificateIn, certificateOut, "CreateCertificate returned different certificate")

	return certificateOut
}

// CreateCertificateFailErrMocked test mocked function
func CreateCertificateFailErrMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// to json
	dOut, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkLoadBalancerCertificates, loadBalancerID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	certificateOut, err := ds.CreateCertificate(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return certificateOut
}

// CreateCertificateFailStatusMocked test mocked function
func CreateCertificateFailStatusMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// to json
	dOut, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkLoadBalancerCertificates, loadBalancerID), mapIn).Return(dOut, 499, nil)
	certificateOut, err := ds.CreateCertificate(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return certificateOut
}

// CreateCertificateFailJSONMocked test mocked function
func CreateCertificateFailJSONMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", fmt.Sprintf(APIPathNetworkLoadBalancerCertificates, loadBalancerID), mapIn).Return(dIn, 200, nil)
	certificateOut, err := ds.CreateCertificate(loadBalancerID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return certificateOut
}

// UpdateCertificateMocked test mocked function
func UpdateCertificateMocked(t *testing.T, loadBalancerID string, certificateIn *types.Certificate) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// to json
	dOut, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID), mapIn).
		Return(dOut, 200, nil)
	certificateOut, err := ds.UpdateCertificate(loadBalancerID, certificateIn.ID, mapIn)

	assert.Nil(err, "Error updating certificate")
	assert.Equal(certificateIn, certificateOut, "UpdateCertificate returned different certificate")

	return certificateOut
}

// UpdateCertificateFailErrMocked test mocked function
func UpdateCertificateFailErrMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// to json
	dOut, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID), mapIn).
		Return(dOut, 200, fmt.Errorf("mocked error"))
	certificateOut, err := ds.UpdateCertificate(loadBalancerID, certificateIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")

	return certificateOut
}

// UpdateCertificateFailStatusMocked test mocked function
func UpdateCertificateFailStatusMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// to json
	dOut, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID), mapIn).
		Return(dOut, 499, nil)
	certificateOut, err := ds.UpdateCertificate(loadBalancerID, certificateIn.ID, mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return certificateOut
}

// UpdateCertificateFailJSONMocked test mocked function
func UpdateCertificateFailJSONMocked(
	t *testing.T,
	loadBalancerID string,
	certificateIn *types.Certificate,
) *types.Certificate {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID), mapIn).
		Return(dIn, 200, nil)
	certificateOut, err := ds.UpdateCertificate(loadBalancerID, certificateIn.ID, mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(certificateOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return certificateOut
}

// DeleteCertificateMocked test mocked function
func DeleteCertificateMocked(t *testing.T, loadBalancerID string, certificateIn *types.Certificate) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID)).
		Return(dIn, 200, nil)
	err = ds.DeleteCertificate(loadBalancerID, certificateIn.ID)

	assert.Nil(err, "Error deleting certificate")
}

// DeleteCertificateFailErrMocked test mocked function
func DeleteCertificateFailErrMocked(t *testing.T, loadBalancerID string, certificateIn *types.Certificate) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID)).
		Return(dIn, 200, fmt.Errorf("mocked error"))
	err = ds.DeleteCertificate(loadBalancerID, certificateIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}

// DeleteCertificateFailStatusMocked test mocked function
func DeleteCertificateFailStatusMocked(t *testing.T, loadBalancerID string, certificateIn *types.Certificate) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCertificateService(cs)
	assert.Nil(err, "Couldn't load certificate service")
	assert.NotNil(ds, "Certificate service not instanced")

	// to json
	dIn, err := json.Marshal(certificateIn)
	assert.Nil(err, "Certificate test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf(APIPathNetworkLoadBalancerCertificate, loadBalancerID, certificateIn.ID)).
		Return(dIn, 499, nil)
	err = ds.DeleteCertificate(loadBalancerID, certificateIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
