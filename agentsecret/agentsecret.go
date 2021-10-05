package agentsecret

import (
	"fmt"

	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns dispatcher commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "retrieve",
			Usage:  "Renders content of a secret",
			Action: cmdRetrieveSecret,
		},
	}
}

func cmdRetrieveSecret(c *cli.Context) error {
	svID := c.Args().Get(0)
	filePath := c.Args().Get(1)
	secretSvc, _, formatter := cmd.WireUpSecret(c)

	status, err := secretSvc.RetrieveSecretVersion(svID, filePath)
	if err == nil && (status > 299 || status < 200) {
		err = fmt.Errorf("Secret content download failed with status %d", status)
	}
	if err != nil {
		formatter.PrintError("Failed", err)
	}
	return err
}
