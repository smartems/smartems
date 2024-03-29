package commands

import (
	"errors"

	"github.com/smartems/smartems/pkg/cmd/smartems-cli/logger"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/utils"
)

func validateVersionInput(c utils.CommandLine) error {
	arg := c.Args().First()
	if arg == "" {
		return errors.New("please specify plugin to list versions for")
	}

	return nil
}

func listversionsCommand(c utils.CommandLine) error {
	if err := validateVersionInput(c); err != nil {
		return err
	}

	pluginToList := c.Args().First()

	plugin, err := c.ApiClient().GetPlugin(pluginToList, c.GlobalString("repo"))
	if err != nil {
		return err
	}

	for _, i := range plugin.Versions {
		logger.Infof("%v\n", i.Version)
	}

	return nil
}
