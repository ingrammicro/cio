// Copyright (c) 2017-2022 Ingram Micro Inc.

package secret

import (
	"fmt"
	"github.com/ingrammicro/cio/cmd"
	"github.com/ingrammicro/cio/cmd/agent"
	"github.com/ingrammicro/cio/logger"
)

func init() {
	cmd.NewCommand(cmd.RootCmd, &cmd.CommandContext{
		Use: "retrieve",
		Short: "Dumps contents of the secret version with the ID given as " +
			"first argument into the file given as second argument",
		RunMethod: RetrieveSecret},
	)
}

func RetrieveSecret(params []string) error {
	logger.DebugFuncInfo()

	svID := params[0]
	filePath := params[1]
	svc, _, formatter := agent.WireUpAPIServer()

	status, err := svc.RetrieveSecretVersion(cmd.GetContext(), svID, filePath)
	if err == nil && (status > 299 || status < 200) {
		err = fmt.Errorf("Secret content download failed with status %d", status)
	}
	if err != nil {
		formatter.PrintError("Failed", err)
	}
	return err
}
