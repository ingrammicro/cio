// Copyright (c) 2017-2021 Ingram Micro Inc.

package brownfield

import (
	"fmt"

	"github.com/ingrammicro/cio/api"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/configuration"
	"github.com/spf13/viper"

	"runtime"

	"github.com/ingrammicro/cio/utils/format"
)

func Configure() error {
	f := format.GetFormatter()
	config, err := configuration.GetConfig()
	if err != nil {
		f.PrintFatal("Couldn't read config", err)
	}
	cs, err := api.NewIMCOClient(config)
	if err != nil {
		f.PrintFatal("Couldn't set up connection to Concerto", err)
	}
	if !config.CurrentUserIsAdmin {
		if runtime.GOOS == "windows" {
			f.PrintFatal("Must run as administrator user", fmt.Errorf("running as non-administrator user"))
		} else {
			f.PrintFatal("Must run as super-user", fmt.Errorf("running as non-administrator user"))
		}
	}
	applyConcertoSettings(cs, f, config.CurrentUserName, viper.GetString(cmd.AdminPassword))
	configureConcertoFirewall(cs, f)
	return nil
}
