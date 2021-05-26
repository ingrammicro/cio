package cloud

import (
	"github.com/ingrammicro/cio/cloud/cloud_providers"
	"github.com/ingrammicro/cio/cloud/generic_images"
	"github.com/ingrammicro/cio/cloud/realms"
	"github.com/ingrammicro/cio/cloud/server_arrays"
	"github.com/ingrammicro/cio/cloud/server_plan"
	"github.com/ingrammicro/cio/cloud/servers"
	"github.com/ingrammicro/cio/cloud/ssh_profiles"
	"github.com/ingrammicro/cio/cloud/temporary_archives"
	"github.com/urfave/cli"
)

// SubCommands returns cloud commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "server-arrays",
			Usage:       "Provides information on server arrays",
			Subcommands: append(server_arrays.SubCommands()),
		},
		{
			Name:        "servers",
			Usage:       "Provides information on servers",
			Subcommands: append(servers.SubCommands()),
		},
		{
			Name:        "generic-images",
			Usage:       "Provides information on generic images",
			Subcommands: append(generic_images.SubCommands()),
		},
		{
			Name:        "ssh-profiles",
			Usage:       "Provides information on SSH profiles",
			Subcommands: append(ssh_profiles.SubCommands()),
		},
		{
			Name:        "cloud-providers",
			Usage:       "Provides information on cloud providers",
			Subcommands: append(cloud_providers.SubCommands()),
		},
		{
			Name:        "realms",
			Usage:       "Provides information on realms",
			Subcommands: append(realms.SubCommands()),
		},
		{
			Name:        "server-plans",
			Usage:       "Provides information on server plans",
			Subcommands: append(server_plan.SubCommands()),
		},
		{
			Name:        "infrastructure",
			Usage:       "Provides infrastructure archives management",
			Subcommands: append(temporary_archives.SubCommands()),
		},
	}
}
