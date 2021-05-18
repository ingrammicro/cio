// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"github.com/ingrammicro/cio/api/wizard"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpLocation prepares common resources to send request to Concerto API
func WireUpLocation(c *cli.Context) (ds *wizard.LocationService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = wizard.NewLocationService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up location service", err)
	}

	return ds, f
}

// LocationList subcommand function
func LocationList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	locationSvc, formatter := WireUpLocation(c)

	locations, err := locationSvc.ListLocations()
	if err != nil {
		formatter.PrintFatal("Couldn't receive location data", err)
	}
	if err = formatter.PrintList(locations); err != nil {
		formatter.PrintFatal(PrintFormatError, err)
	}
	return nil
}

// LoadLocationsMapping retrieves Locations and create a map between ID and Name
func LoadLocationsMapping(c *cli.Context) map[string]string {
	debugCmdFuncInfo(c)

	locationSvc, formatter := WireUpLocation(c)
	locations, err := locationSvc.ListLocations()
	if err != nil {
		formatter.PrintFatal("Couldn't receive location data", err)
	}
	locationsMap := make(map[string]string)
	for _, location := range locations {
		locationsMap[location.ID] = location.Name
	}

	return locationsMap
}
