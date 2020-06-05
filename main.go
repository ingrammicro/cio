// Copyright (c) 2017-2022 Ingram Micro Inc.

package main

// Agent mode

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"

	// In order to make cobra able to "discover" commands from sub-folder levels
	_ "github.com/ingrammicro/cio/cmd/agent/bootstrapping"
	_ "github.com/ingrammicro/cio/cmd/agent/brownfield"
	_ "github.com/ingrammicro/cio/cmd/agent/converge"
	_ "github.com/ingrammicro/cio/cmd/agent/dispatcher"
	_ "github.com/ingrammicro/cio/cmd/agent/firewall"
	_ "github.com/ingrammicro/cio/cmd/agent/polling"
	_ "github.com/ingrammicro/cio/cmd/agent/secret"

	//_ "github.com/ingrammicro/cio/cmd/agent"
)

func main() {
	cmd.Execute(configuration.Server)
}
