package blueprint

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// CookbookVersionService manages cookbook version operations
type CookbookVersionService struct {
	concertoService utils.ConcertoService
}

// NewCookbookVersionService returns a Concerto cookbook version service
func NewCookbookVersionService(concertoService utils.ConcertoService) (*CookbookVersionService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &CookbookVersionService{
		concertoService: concertoService,
	}, nil
}

// ListCookbookVersions returns the list of cookbook versions as an array of CookbookVersion
func (cvs *CookbookVersionService) ListCookbookVersions() (cookbookVersions []*types.CookbookVersion, err error) {
	log.Debug("ListCookbookVersions")

	data, status, err := cvs.concertoService.Get("/blueprint/cookbook_versions")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cookbookVersions); err != nil {
		return nil, err
	}

	return cookbookVersions, nil
}

// GetCookbookVersion returns a cookbook version by its ID
func (cvs *CookbookVersionService) GetCookbookVersion(cookbookVersionID string) (cookbookVersion *types.CookbookVersion, err error) {
	log.Debug("GetCookbookVersion")

	data, status, err := cvs.concertoService.Get(fmt.Sprintf("/blueprint/cookbook_versions/%s", cookbookVersionID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cookbookVersion); err != nil {
		return nil, err
	}

	return cookbookVersion, nil
}

// CreateCookbookVersion creates a new cookbook version
func (cvs *CookbookVersionService) CreateCookbookVersion(cookbookVersionParams *map[string]interface{}) (cookbookVersion *types.CookbookVersion, err error) {
	log.Debug("CreateCookbookVersion")

	data, status, err := cvs.concertoService.Post("/blueprint/cookbook_versions", cookbookVersionParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cookbookVersion); err != nil {
		return nil, err
	}

	return cookbookVersion, nil
}

// UploadCookbookVersion uploads a cookbook version file
func (cvs *CookbookVersionService) UploadCookbookVersion(sourceFilePath string, targetURL string) error {
	log.Debug("UploadCookbookVersion")

	data, status, err := cvs.concertoService.PutFile(sourceFilePath, targetURL)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ProcessCookbookVersion process a cookbook version by its ID
func (cvs *CookbookVersionService) ProcessCookbookVersion(cookbookVersionID string, cookbookVersionParams *map[string]interface{}) (cookbookVersion *types.CookbookVersion, err error) {
	log.Debug("ProcessCookbookVersion")

	data, status, err := cvs.concertoService.Post(fmt.Sprintf("/blueprint/cookbook_versions/%s/process", cookbookVersionID), cookbookVersionParams)

	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cookbookVersion); err != nil {
		return nil, err
	}

	return cookbookVersion, nil
}

// DeleteCookbookVersion deletes a cookbook version by its ID
func (cvs *CookbookVersionService) DeleteCookbookVersion(cookbookVersionID string) (err error) {
	log.Debug("DeleteCookbookVersion")

	data, status, err := cvs.concertoService.Delete(fmt.Sprintf("/blueprint/cookbook_versions/%s", cookbookVersionID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
