// Copyright (c) 2017-2021 Ingram Micro Inc.
package agentsecret

import (
	"fmt"

	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

const APIPathSecretVersionContent = "/secret/secret_versions/%s"

// SecretService manages secret retrieval operations
type SecretService struct {
	apiEndpoint     string
	concertoService utils.ConcertoService
}

// NewSecretService returns a dispatcher service
func NewSecretService(concertoService utils.ConcertoService, apiEndpoint string) (*SecretService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &SecretService{
		apiEndpoint:     apiEndpoint,
		concertoService: concertoService,
	}, nil
}

// RetrieveSecretVersion returns script characterizations list for a given UUID
func (ss *SecretService) RetrieveSecretVersion(svID, filePath string) (int, error) {
	log.Debug("RetrieveSecretVersion")

	_, status, err := ss.concertoService.GetFile(
		fmt.Sprintf("%s"+APIPathSecretVersionContent, ss.apiEndpoint, svID), filePath, false)
	if err != nil {
		return status, err
	}
	return status, nil
}
