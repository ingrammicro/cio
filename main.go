// Copyright (c) 2017-2021 Ingram Micro Inc.

package main

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"

	// In order to make cobra able to "discover" commands from sub-folder levels

	// User
	_ "github.com/ingrammicro/cio/cmd/client/audit"
	_ "github.com/ingrammicro/cio/cmd/client/blueprint"
	_ "github.com/ingrammicro/cio/cmd/client/brownfield"
	_ "github.com/ingrammicro/cio/cmd/client/cloud"
	_ "github.com/ingrammicro/cio/cmd/client/cloudapplications"
	_ "github.com/ingrammicro/cio/cmd/client/cloudspecificextensions"
	_ "github.com/ingrammicro/cio/cmd/client/kubernetes"
	_ "github.com/ingrammicro/cio/cmd/client/labels"
	_ "github.com/ingrammicro/cio/cmd/client/network"
	_ "github.com/ingrammicro/cio/cmd/client/settings"
	_ "github.com/ingrammicro/cio/cmd/client/storage"
	_ "github.com/ingrammicro/cio/cmd/client/wizard"

	// Server
	//_ "github.com/ingrammicro/cio/cmd/server/bootstrapping"
	//_ "github.com/ingrammicro/cio/cmd/server/brownfield"
	//_ "github.com/ingrammicro/cio/cmd/server/converge"
	//_ "github.com/ingrammicro/cio/cmd/server/dispatcher"
	//_ "github.com/ingrammicro/cio/cmd/server/firewall"
	//_ "github.com/ingrammicro/cio/cmd/server/polling"
)

// TODO 2 main? uno para modo: User/Server
// TODO so, TO BE REMOVED!?
func main() {
	cmd.Execute(configuration.Client)
}
