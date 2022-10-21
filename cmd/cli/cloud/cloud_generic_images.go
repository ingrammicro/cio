// Copyright (c) 2017-2022 Ingram Micro Inc.

package cloud

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/cli"
	"github.com/ingrammicro/cio/logger"
)

func init() {
	genericImagesCmd := cmd.NewCommand(cloudCmd, &cmd.CommandContext{
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
	svc, _, formatter := cli.WireUpAPIClient()

	genericImages, err := svc.ListGenericImages(cmd.GetContext())
	if err != nil {
		formatter.PrintError("Couldn't receive genericImage data", err)
		return err
	}
	if err = formatter.PrintList(genericImages); err != nil {
		formatter.PrintError(cmd.PrintFormatError, err)
		return err
	}
	return nil
}
