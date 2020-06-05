// Copyright (c) 2017-2021 Ingram Micro Inc.

package main

// Server / Agent mode

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"

	// In order to make cobra able to "discover" commands from sub-folder levels
	_ "github.com/ingrammicro/cio/cmd/server/bootstrapping"
	_ "github.com/ingrammicro/cio/cmd/server/brownfield"
	_ "github.com/ingrammicro/cio/cmd/server/converge"
	_ "github.com/ingrammicro/cio/cmd/server/dispatcher"
	_ "github.com/ingrammicro/cio/cmd/server/firewall"
	_ "github.com/ingrammicro/cio/cmd/server/polling"
)

func main() {
	cmd.Execute(configuration.Server)
}
