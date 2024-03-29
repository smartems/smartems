package commands

import (
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/smartems/smartems/pkg/bus"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/commands/datamigrations"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/logger"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/utils"
	"github.com/smartems/smartems/pkg/services/sqlstore"
	"github.com/smartems/smartems/pkg/setting"
)

func runDbCommand(command func(commandLine utils.CommandLine, sqlStore *sqlstore.SqlStore) error) func(context *cli.Context) {
	return func(context *cli.Context) {
		cmd := &utils.ContextCommandLine{Context: context}
		debug := cmd.GlobalBool("debug")

		cfg := setting.NewCfg()

		configOptions := strings.Split(cmd.GlobalString("configOverrides"), " ")
		if err := cfg.Load(&setting.CommandLineArgs{
			Config:   cmd.ConfigFile(),
			HomePath: cmd.HomePath(),
			Args:     append(configOptions, cmd.Args()...), // tailing arguments have precedence over the options string
		}); err != nil {
			logger.Errorf("\n%s: Failed to load configuration", color.RedString("Error"))
			os.Exit(1)
		}

		if debug {
			cfg.LogConfigSources()
		}

		engine := &sqlstore.SqlStore{}
		engine.Cfg = cfg
		engine.Bus = bus.GetBus()
		if err := engine.Init(); err != nil {
			logger.Errorf("\n%s: Failed to initialize SQL engine", color.RedString("Error"))
			os.Exit(1)
		}

		if err := command(cmd, engine); err != nil {
			logger.Errorf("\n%s: ", color.RedString("Error"))
			logger.Errorf("%s\n\n", err)

			if err := cmd.ShowHelp(); err != nil {
				logger.Errorf("\n%s: Failed to show help: %s %s\n\n", color.RedString("Error"),
					color.RedString("✗"), err)
			}
			os.Exit(1)
		}

		logger.Info("\n\n")
	}
}

func runPluginCommand(command func(commandLine utils.CommandLine) error) func(context *cli.Context) {
	return func(context *cli.Context) {

		cmd := &utils.ContextCommandLine{Context: context}
		if err := command(cmd); err != nil {
			logger.Errorf("\n%s: ", color.RedString("Error"))
			logger.Errorf("%s %s\n\n", color.RedString("✗"), err)

			if err := cmd.ShowHelp(); err != nil {
				logger.Errorf("\n%s: Failed to show help: %s %s\n\n", color.RedString("Error"),
					color.RedString("✗"), err)
			}
			os.Exit(1)
		}

		logger.Info("\nRestart smartems after installing plugins . <service smartems-server restart>\n\n")
	}
}

var pluginCommands = []cli.Command{
	{
		Name:   "install",
		Usage:  "install <plugin id> <plugin version (optional)>",
		Action: runPluginCommand(installCommand),
	}, {
		Name:   "list-remote",
		Usage:  "list remote available plugins",
		Action: runPluginCommand(listRemoteCommand),
	}, {
		Name:   "list-versions",
		Usage:  "list-versions <plugin id>",
		Action: runPluginCommand(listversionsCommand),
	}, {
		Name:    "update",
		Usage:   "update <plugin id>",
		Aliases: []string{"upgrade"},
		Action:  runPluginCommand(upgradeCommand),
	}, {
		Name:    "update-all",
		Aliases: []string{"upgrade-all"},
		Usage:   "update all your installed plugins",
		Action:  runPluginCommand(upgradeAllCommand),
	}, {
		Name:   "ls",
		Usage:  "list all installed plugins",
		Action: runPluginCommand(lsCommand),
	}, {
		Name:    "uninstall",
		Aliases: []string{"remove"},
		Usage:   "uninstall <plugin id>",
		Action:  runPluginCommand(removeCommand),
	},
}

var adminCommands = []cli.Command{
	{
		Name:   "reset-admin-password",
		Usage:  "reset-admin-password <new password>",
		Action: runDbCommand(resetPasswordCommand),
	},
	{
		Name:  "data-migration",
		Usage: "Runs a script that migrates or cleanups data in your db",
		Subcommands: []cli.Command{
			{
				Name:   "encrypt-datasource-passwords",
				Usage:  "Migrates passwords from unsecured fields to secure_json_data field. Return ok unless there is an error. Safe to execute multiple times.",
				Action: runDbCommand(datamigrations.EncryptDatasourcePaswords),
			},
		},
	},
}

var Commands = []cli.Command{
	{
		Name:        "plugins",
		Usage:       "Manage plugins for smartems",
		Subcommands: pluginCommands,
	},
	{
		Name:        "admin",
		Usage:       "Grafana admin commands",
		Subcommands: adminCommands,
	},
}
