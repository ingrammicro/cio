package clientbrownfield

import (
	"github.com/ingrammicro/cio/clientbrownfield/cloud_accounts"
	"github.com/ingrammicro/cio/clientbrownfield/import_candidates"
	"github.com/urfave/cli"
)

// SubCommands returns client brownfield commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "cloud-accounts",
			Usage:       "Provides information about brownfield cloud accounts. Allows querying cloud accounts to discover (and list) candidate resources from IMCO.",
			Subcommands: append(cloud_accounts.SubCommands()),
		},
		{
			Name:        "import-candidates",
			Usage:       "Provides information about brownfield import candidates. Allows querying cloud accounts to show (and import) candidate resources from them into IMCO.",
			Subcommands: append(import_candidates.SubCommands()),
		},
	}
}
