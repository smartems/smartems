package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/smartems/smartems/pkg/bus"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/logger"
	"github.com/smartems/smartems/pkg/cmd/smartems-cli/utils"
	"github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/services/sqlstore"
	"github.com/smartems/smartems/pkg/util"
	"github.com/smartems/smartems/pkg/util/errutil"
)

const AdminUserId = 1

func resetPasswordCommand(c utils.CommandLine, sqlStore *sqlstore.SqlStore) error {
	newPassword := c.Args().First()

	password := models.Password(newPassword)
	if password.IsWeak() {
		return fmt.Errorf("New password is too short")
	}

	userQuery := models.GetUserByIdQuery{Id: AdminUserId}

	if err := bus.Dispatch(&userQuery); err != nil {
		return fmt.Errorf("Could not read user from database. Error: %v", err)
	}

	passwordHashed, err := util.EncodePassword(newPassword, userQuery.Result.Salt)
	if err != nil {
		return err
	}

	cmd := models.ChangeUserPasswordCommand{
		UserId:      AdminUserId,
		NewPassword: passwordHashed,
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return errutil.Wrapf(err, "Failed to update user password")
	}

	logger.Infof("\n")
	logger.Infof("Admin password changed successfully %s", color.GreenString("✔"))

	return nil
}
