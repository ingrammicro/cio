// Copyright (c) 2017-2022 Ingram Micro Inc.

package brownfield

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd/agent"

	"github.com/ingrammicro/cio/cmd"
	"github.com/spf13/viper"

	"runtime"
)

func Configure() error {
	svc, config, f := agent.WireUpAPIServer()
	if !config.CurrentUserIsAdmin {
		if runtime.GOOS == "windows" {
			f.PrintFatal("Must run as administrator user", fmt.Errorf("running as non-administrator user"))
		} else {
			f.PrintFatal("Must run as super-user", fmt.Errorf("running as non-administrator user"))
		}
	}
	ctx := cmd.GetContext()
	applySettings(ctx, svc, f, config.CurrentUserName, viper.GetString(cmd.AdminPassword))
	configureFirewall(ctx, svc, f)
	return nil
}
