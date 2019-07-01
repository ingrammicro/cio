package network

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/cio/network/firewall_profiles"
	"github.com/ingrammicro/cio/network/floating_ips"
	"github.com/ingrammicro/cio/network/subnets"
	"github.com/ingrammicro/cio/network/vpcs"
	"github.com/ingrammicro/cio/network/vpns"
)

// SubCommands returns network commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "firewall-profiles",
			Usage:       "Provides information about firewall profiles",
			Subcommands: append(firewall_profiles.SubCommands()),
		},
		{
			Name:        "floating-ips",
			Usage:       "Provides information about floating IPs",
			Subcommands: append(floating_ips.SubCommands()),
		},
		{
			Name:        "vpcs",
			Usage:       "Provides information about Virtual Private Clouds (VPCs)",
			Subcommands: append(vpcs.SubCommands()),
		},
		{
			Name:        "subnets",
			Usage:       "Provides information about VPC Subnets",
			Subcommands: append(subnets.SubCommands()),
		},
		{
			Name:        "vpns",
			Usage:       "Provides information about VPC Virtual Private Networks (VPNs)",
			Subcommands: append(vpns.SubCommands()),
		},
	}
}
