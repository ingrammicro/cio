package wizard

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// LocationService manages location operations
type LocationService struct {
	concertoService utils.ConcertoService
}

// NewLocationService returns a Concerto location service
func NewLocationService(concertoService utils.ConcertoService) (*LocationService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &LocationService{
		concertoService: concertoService,
	}, nil
}

// GetLocationList returns the list of locations as an array of Location
func (dm *LocationService) GetLocationList() (locations []*types.Location, err error) {
	log.Debug("GetLocationList")

	data, status, err := dm.concertoService.Get("/wizard/locations")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}
