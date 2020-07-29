package wizard

import (
	"encoding/json"
	"fmt"

	"github.com/ingrammicro/cio/api/types"
	"github.com/ingrammicro/cio/utils"
	log "github.com/sirupsen/logrus"
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

// ListLocations returns the list of locations as an array of Location
func (ls *LocationService) ListLocations() (locations []*types.Location, err error) {
	log.Debug("ListLocations")

	data, status, err := ls.concertoService.Get("/wizard/locations")
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
