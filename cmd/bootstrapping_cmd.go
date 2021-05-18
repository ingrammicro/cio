// Copyright (c) 2017-2021 Ingram Micro Inc.

package cmd

import (
	"github.com/ingrammicro/cio/api/blueprint"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpBootstrapping prepares common resources to send request to API
func WireUpBootstrapping(c *cli.Context) (ds *blueprint.BootstrappingService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = blueprint.NewBootstrappingService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up bootstrapping service", err)
	}

	return ds, f
}
