// Copyright (c) 2017-2021 Ingram Micro Inc.

package bootstrapping

import (
	"github.com/urfave/cli"
)

// SubCommands returns bootstrapping commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "start",
			Usage:  "Starts a bootstrapping routine to check and execute required activities",
			Action: start,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "lines, l",
					Usage: "Maximum lines threshold per response chunk",
					Value: defaultThresholdLines,
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops the running bootstrapping process",
			Action: stop,
		},
	}
}
