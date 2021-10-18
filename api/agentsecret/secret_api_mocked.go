// Copyright (c) 2017-2021 Ingram Micro Inc.

package agentsecret

import (
	"fmt"
	"testing"

	"github.com/ingrammicro/cio/utils"
	"github.com/stretchr/testify/assert"
)

func RetrieveSecretVersionMocked(t *testing.T, svID, filePath string) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	apiEndpoint := "https://clients.example.com"
	ss, err := NewSecretService(cs, apiEndpoint)
	assert.Nil(err, "Couldn't load secret service")
	assert.NotNil(ss, "Secret service not instanced")

	urlPath := fmt.Sprintf("/secret/secret_versions/%s", svID)
	url := fmt.Sprintf("%s%s", apiEndpoint, urlPath)

	// call service
	cs.On("GetFile", url, filePath).Return(filePath, 200, nil)
	status, err := ss.RetrieveSecretVersion(svID, filePath)
	assert.Nil(err, "Error downloading attachment file")
	assert.Equal(status, 200, "RetrieveSecretVersion returned invalid response")
}

// RetrieveSecretVersionFailErrMocked test mocked function
func RetrieveSecretVersionFailErrMocked(t *testing.T, svID, filePath string) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	apiEndpoint := "https://clients.example.com"
	ss, err := NewSecretService(cs, apiEndpoint)
	assert.Nil(err, "Couldn't load secret service")
	assert.NotNil(ss, "Secret service not instanced")

	urlPath := fmt.Sprintf("/secret/secret_versions/%s", svID)
	url := fmt.Sprintf("%s%s", apiEndpoint, urlPath)

	// call service
	cs.On("GetFile", url, filePath).Return("", 499, fmt.Errorf("mocked error"))
	status, err := ss.RetrieveSecretVersion(svID, filePath)
	assert.NotNil(err, "We are expecting an error")
	assert.Equal(status, 499, "RetrieveSecretVersion returned an unexpected status code")
	assert.Equal(err.Error(), "mocked error", "Error should be 'mocked error'")
}
