package clientbrownfield

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// BrownfieldCloudAccountService manages brownfield cloud account operations
type BrownfieldCloudAccountService struct {
	concertoService utils.ConcertoService
}

// NewBrownfieldCloudAccountService returns a Concerto BrownfieldCloudAccount service
func NewBrownfieldCloudAccountService(concertoService utils.ConcertoService) (*BrownfieldCloudAccountService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &BrownfieldCloudAccountService{
		concertoService: concertoService,
	}, nil
}

// ListBrownfieldCloudAccounts returns the list of Brownfield Cloud Accounts as an array of CloudAccount
func (bcas *BrownfieldCloudAccountService) ListBrownfieldCloudAccounts() (cloudAccounts []*types.CloudAccount, err error) {
	log.Debug("ListBrownfieldCloudAccounts")

	data, status, err := bcas.concertoService.Get("/brownfield/cloud_accounts")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccounts); err != nil {
		return nil, err
	}

	return cloudAccounts, nil
}

// GetBrownfieldCloudAccount returns a Brownfield Cloud Account by its ID
func (bcas *BrownfieldCloudAccountService) GetBrownfieldCloudAccount(cloudAccountID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("GetBrownfieldCloudAccount")

	data, status, err := bcas.concertoService.Get(fmt.Sprintf("/brownfield/cloud_accounts/%s", cloudAccountID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}
