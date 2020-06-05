// Copyright (c) 2017-2022 Ingram Micro Inc.

package main

// Client / User mode

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"

	// In order to make cobra able to "discover" commands from sub-folder levels
	_ "github.com/ingrammicro/cio/cmd/cli/audit"
	_ "github.com/ingrammicro/cio/cmd/cli/blueprint"
	_ "github.com/ingrammicro/cio/cmd/cli/brownfield"
	_ "github.com/ingrammicro/cio/cmd/cli/cloud"
	_ "github.com/ingrammicro/cio/cmd/cli/cloudapplications"
	_ "github.com/ingrammicro/cio/cmd/cli/cloudspecificextensions"
	_ "github.com/ingrammicro/cio/cmd/cli/kubernetes"
	_ "github.com/ingrammicro/cio/cmd/cli/labels"
	_ "github.com/ingrammicro/cio/cmd/cli/network"
	_ "github.com/ingrammicro/cio/cmd/cli/settings"
	_ "github.com/ingrammicro/cio/cmd/cli/storage"
	_ "github.com/ingrammicro/cio/cmd/cli/wizard"
)

func main() {
	cmd.Execute(configuration.Client)
}
