package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/ingrammicro/cio/audit"
	"github.com/ingrammicro/cio/blueprint"
	"github.com/ingrammicro/cio/bootstrapping"
	"github.com/ingrammicro/cio/brownfield"
	"github.com/ingrammicro/cio/clientbrownfield"
	"github.com/ingrammicro/cio/cloud"
	"github.com/ingrammicro/cio/cloudspecificextensions"
	"github.com/ingrammicro/cio/cmdpolling"
	"github.com/ingrammicro/cio/converge"
	"github.com/ingrammicro/cio/dispatcher"
	"github.com/ingrammicro/cio/firewall"
	"github.com/ingrammicro/cio/labels"
	"github.com/ingrammicro/cio/network"
	"github.com/ingrammicro/cio/settings"
	"github.com/ingrammicro/cio/storage"
	"github.com/ingrammicro/cio/utils"
	"github.com/ingrammicro/cio/utils/format"
	"github.com/ingrammicro/cio/wizard"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var serverCommands = []cli.Command{
	{
		Name:        "bootstrap",
		Usage:       "Manages bootstrapping commands",
		Subcommands: append(bootstrapping.SubCommands()),
	},
	{
		Name:        "brownfield",
		Usage:       "Manages registration and configuration within an imported brownfield Host",
		Subcommands: append(brownfield.SubCommands()),
	},
	{
		Name:   "converge",
		Usage:  "Converges Host to original Blueprint",
		Action: converge.CmbConverge,
	},
	{
		Name:        "firewall",
		Usage:       "Manages Firewall Policies within a Host",
		Subcommands: append(firewall.SubCommands()),
	},
	{
		Name:        "polling",
		Usage:       "Manages polling commands",
		Subcommands: append(cmdpolling.SubCommands()),
	},
	{
		Name:        "scripts",
		Usage:       "Manages Execution Scripts within a Host",
		Subcommands: append(dispatcher.SubCommands()),
	},
}

var clientCommands = []cli.Command{
	{
		Name:        "blueprint",
		ShortName:   "bl",
		Usage:       "Manages blueprint commands for scripts, services and templates",
		Subcommands: append(blueprint.SubCommands()),
	},
	{
		Name:        "brownfield",
		ShortName:   "bf",
		Usage:       "Manages brownfield resources, allowing users to discover and import servers, VPCs, floating IPs and volumes from different cloud accounts into the system.",
		Subcommands: append(clientbrownfield.SubCommands()),
	},
	{
		Name:        "cloud",
		ShortName:   "clo",
		Usage:       "Manages cloud related commands for server arrays, servers, generic images, ssh profiles, cloud providers and server plans",
		Subcommands: append(cloud.SubCommands()),
	},
	{
		Name:        "cloud-specific-extensions",
		ShortName:   "cse",
		Usage:       "Manages cloud specific extensions -CSEs- templates and deployments",
		Subcommands: append(cloudspecificextensions.SubCommands()),
	},
	{
		Name:        "events",
		ShortName:   "ev",
		Usage:       "Events allow the user to track their actions and the state of their servers",
		Subcommands: append(audit.SubCommands()),
	},
	{
		Name:        "labels",
		ShortName:   "lbl",
		Usage:       "Provides information about labels",
		Subcommands: append(labels.SubCommands()),
	},
	{
		Name:        "network",
		ShortName:   "net",
		Usage:       "Manages network related commands for firewall profiles",
		Subcommands: append(network.SubCommands()),
	},
	{
		Name:        "storage",
		ShortName:   "st",
		Usage:       "Manages storage commands for plans and volumes",
		Subcommands: append(storage.SubCommands()),
	},
	{
		Name:        "settings",
		ShortName:   "set",
		Usage:       "Provides settings for cloud accounts",
		Subcommands: append(settings.SubCommands()),
	},
	{
		Name:        "wizard",
		ShortName:   "wiz",
		Usage:       "Manages wizard related commands for apps, locations, cloud providers, server plans",
		Subcommands: append(wizard.SubCommands()),
	},
}

var appFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "debug, D",
		Usage: "Enable debug mode",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_CA_CERT",
		Name:   "ca-cert",
		Usage:  "CA to verify remote connections",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_CLIENT_CERT",
		Name:   "client-cert",
		Usage:  "Client cert to use for Concerto",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_CLIENT_KEY",
		Name:   "client-key",
		Usage:  "Private key used in client Concerto auth",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_CONFIG",
		Name:   "concerto-config",
		Usage:  "Concerto Config File",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_ENDPOINT",
		Name:   "concerto-endpoint",
		Usage:  "Concerto Endpoint",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_URL",
		Name:   "concerto-url",
		Usage:  "Concerto Web URL",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_BROWNFIELD_TOKEN",
		Name:   "concerto-brownfield-token",
		Usage:  "Concerto Brownfield Token",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_COMMAND_POLLING_TOKEN",
		Name:   "concerto-command-polling-token",
		Usage:  "Concerto Command Polling Token",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_SERVER_ID",
		Name:   "concerto-server-id",
		Usage:  "Concerto Server ID",
	},
	cli.StringFlag{
		EnvVar: "CONCERTO_FORMATTER",
		Name:   "formatter",
		Usage:  "Output formatter [ text | json ] ",
		Value:  "text",
	},
}

func excludeFlags(visibleFlags []cli.Flag, arr []string) (flags []cli.Flag) {
	for _, flag := range visibleFlags {
		bFound := false
		for _, a := range arr {
			if a == flag.GetName() {
				bFound = true
				break
			}
		}
		if !bFound {
			flags = append(flags, flag)
		}
	}
	return
}

func cmdNotFound(c *cli.Context, command string) {
	log.Fatalf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		c.App.Name,
	)
}

func prepareFlags(c *cli.Context) error {
	if c.Bool("debug") {
		if err := os.Setenv("DEBUG", "1"); err != nil {
			log.Errorf("Couldn't set environment debug mode: %s", err)
			return err
		}
		log.SetOutput(os.Stderr)
		log.SetLevel(log.DebugLevel)
	}

	// try to read configuration
	config, err := utils.InitializeConcertoConfig(c)
	if err != nil {
		log.Errorf("Error reading Concerto configuration: %s", err)
		return err
	}

	// validate formatter
	if c.String("formatter") != "text" && c.String("formatter") != "json" {
		log.Errorf("Unrecognized formatter %s. Please, use one of [ text | json ]", c.String("formatter"))
		return fmt.Errorf("unrecognized formatter %s. Please, use one of [ text | json ]", c.String("formatter"))
	}
	format.InitializeFormatter(c.String("formatter"), os.Stdout)

	if config.IsAgentMode() {
		log.Debug("Setting server commands to concerto")
		c.App.Commands = serverCommands
	} else {
		log.Debug("Setting client commands to concerto")
		c.App.Commands = clientCommands

		// Excluding Server/Agent contextual flags
		c.App.Flags = excludeFlags(c.App.VisibleFlags(), []string{"concerto-brownfield-token", "concerto-command-polling-token", "concerto-server-id"})
	}

	sort.Sort(cli.CommandsByName(c.App.Commands))
	sort.Sort(cli.FlagsByName(c.App.Flags))

	// hack: substitute commands in category ... we should evaluate cobra/viper
	cat := c.App.Categories()

	for _, category := range cat {
		category.Commands = category.Commands[:0]
	}

	for _, command := range c.App.Commands {
		cat = cat.AddCommand(command.Category, command)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "cio"
	app.Author = "Ingram Micro"
	app.Email = "https://github.com/ingrammicro/cio"

	app.CommandNotFound = cmdNotFound
	app.Usage = "Manages communication between Host and IMCO Platform"
	app.Version = utils.VERSION

	// set client commands by default to populate categories
	app.Commands = clientCommands

	app.Flags = appFlags

	app.Before = prepareFlags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
