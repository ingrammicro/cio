// Copyright (c) 2017-2021 Ingram Micro Inc.

package wizard

import (
	"github.com/ingrammicro/cio/wizard/apps"
	"github.com/ingrammicro/cio/wizard/cloud_providers"
	"github.com/ingrammicro/cio/wizard/locations"
	"github.com/ingrammicro/cio/wizard/server_plans"
	"github.com/urfave/cli"
)

// SubCommands returns wizard commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "apps",
			Usage:       "Provides information about apps",
			Subcommands: append(apps.SubCommands()),
		},
		{
			Name:        "cloud-providers",
			Usage:       "Provides information about cloud providers",
			Subcommands: append(cloud_providers.SubCommands()),
		},
		{
			Name:        "locations",
			Usage:       "Provides information about locations",
			Subcommands: append(locations.SubCommands()),
		},
		{
			Name:        "server-plans",
			Usage:       "Provides information about server plans",
			Subcommands: append(server_plans.SubCommands()),
		},
	}
}
