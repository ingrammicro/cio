package cmd

import (
	"github.com/ingrammicro/cio/api/agentsecret"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpSecret prepares common resources to send request to API
func WireUpSecret(c *cli.Context) (ss *agentsecret.SecretService, config *utils.Config, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ss, err = agentsecret.NewSecretService(hcs, config.APIEndpoint)
	if err != nil {
		f.PrintFatal("Couldn't wire up secret service", err)
	}

	return ss, config, f
}
