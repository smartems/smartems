package commands

import (
	"github.com/fatih/color"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/logger"
	s "github.com/smartems/smartems/pkg/cmd/smartems-cli/services"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/utils"
	"github.com/smartems/smartems/pkg/util/errutil"
)

func upgradeCommand(c utils.CommandLine) error {
	pluginsDir := c.PluginDirectory()
	pluginName := c.Args().First()

	localPlugin, err := s.ReadPlugin(pluginsDir, pluginName)

	if err != nil {
		return err
	}

	plugin, err2 := c.ApiClient().GetPlugin(pluginName, c.RepoDirectory())

	if err2 != nil {
		return err2
	}

	if shouldUpgrade(localPlugin.Info.Version, &plugin) {
		if err := s.RemoveInstalledPlugin(pluginsDir, pluginName); err != nil {
			return errutil.Wrapf(err, "Failed to remove plugin '%s'", pluginName)
		}

		return InstallPlugin(pluginName, "", c)
	}

	logger.Infof("%s %s is up to date \n", color.GreenString("✔"), pluginName)
	return nil
}
