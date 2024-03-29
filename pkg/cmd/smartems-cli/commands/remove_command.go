package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/smartems/smartems/pkg/cmd/smartems-cli/services"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/utils"
)

var removePlugin func(pluginPath, id string) error = services.RemoveInstalledPlugin

func removeCommand(c utils.CommandLine) error {
	pluginPath := c.PluginDirectory()

	plugin := c.Args().First()
	if plugin == "" {
		return errors.New("Missing plugin parameter")
	}

	err := removePlugin(pluginPath, plugin)

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return fmt.Errorf("Plugin does not exist")
		}

		return err
	}

	return nil
}
