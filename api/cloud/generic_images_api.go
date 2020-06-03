package cloud

import (
	"encoding/json"
	"fmt"
	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
)

// GenericImageService manages generic image operations
type GenericImageService struct {
	concertoService utils.ConcertoService
}

// NewGenericImageService returns a Concerto genericImage service
func NewGenericImageService(concertoService utils.ConcertoService) (*GenericImageService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &GenericImageService{
		concertoService: concertoService,
	}, nil
}

// ListGenericImages returns the list of generic images as an array of GenericImage
func (gis *GenericImageService) ListGenericImages() (genericImages []*types.GenericImage, err error) {
	log.Debug("ListGenericImages")

	data, status, err := gis.concertoService.Get("/cloud/generic_images")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &genericImages); err != nil {
		return nil, err
	}

	return genericImages, nil
}
