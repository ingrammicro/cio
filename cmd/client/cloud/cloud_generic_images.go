// Copyright (c) 2017-2021 Ingram Micro Inc.

package cloud

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/logger"
)

func init() {
	genericImagesCmd := cmd.NewCommand(CloudCmd, &cmd.CommandContext{
		Use:   "generic-images",
		Short: "Provides information on generic images"},
	)
	cmd.NewCommand(genericImagesCmd, &cmd.CommandContext{
		Use:       "list",
		Short:     "This action lists the available generic images",
		RunMethod: GenericImageList},
	)
}

// GenericImageList subcommand function
func GenericImageList() error {
	logger.DebugFuncInfo()
	svc, _, formatter := cmd.WireUpAPI()

	genericImages, err := svc.ListGenericImages()
	if err != nil {
		formatter.PrintFatal("Couldn't receive genericImage data", err)
	}
	if err = formatter.PrintList(genericImages); err != nil {
		formatter.PrintFatal(cmd.PrintFormatError, err)
	}
	return nil
}
