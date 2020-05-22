package cmd

import (
	"github.com/ingrammicro/cio/api/cloud"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/urfave/cli"
)

// WireUpGenericImage prepares common resources to send request to Concerto API
func WireUpGenericImage(c *cli.Context) (ns *cloud.GenericImageService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = cloud.NewGenericImageService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up genericImage service", err)
	}

	return ns, f
}

// GenericImageList subcommand function
func GenericImageList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	genericImageSvc, formatter := WireUpGenericImage(c)

	genericImages, err := genericImageSvc.ListGenericImages()
	if err != nil {
		formatter.PrintFatal("Couldn't receive genericImage data", err)
	}
	if err = formatter.PrintList(genericImages); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
