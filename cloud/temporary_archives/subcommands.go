package temporary_archives

import (
	"github.com/ingrammicro/cio/cmd"
	"github.com/urfave/cli"
)

// SubCommands returns temporary archives commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "export",
			Usage:  "Exports infrastructure file from IMCO",
			Action: cmd.TemporaryArchiveExport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "server-ids",
					Usage: "A list of comma separated server identifiers to be exported",
				},
				cli.StringFlag{
					Name:  "server-array-ids",
					Usage: "A list of comma separated server array identifiers to be exported",
				},
				cli.StringFlag{
					Name:  "filepath",
					Usage: "Path and file name to download infrastructure 'csar' file, i.e: --filename /folder-path/filename.csar",
				},
				cli.IntFlag{
					Name:  "time, t",
					Usage: "Time lapse -seconds- for export status check",
					Value: cmd.DefaultTimeLapseExportStatusCheck,
				},
			},
		},
		{
			Name:   "import",
			Usage:  "Imports infrastructure file on IMCO",
			Action: cmd.TemporaryArchiveImport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "filepath",
					Usage: "Path and file name to infrastructure 'csar' file",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "New label name to be associated with infrastructure",
				},
				cli.IntFlag{
					Name:  "time, t",
					Usage: "Time lapse -seconds- for import status check",
					Value: cmd.DefaultTimeLapseImportStatusCheck,
				},
			},
		},
	}
}
