package api

import (
	"fmt"

	"github.com/smartems/smartems/pkg/api/dtos"
	"github.com/smartems/smartems/pkg/bus"
	"github.com/smartems/smartems/pkg/events"
	"github.com/smartems/smartems/pkg/infra/metrics"
	m "github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/setting"
	"github.com/smartems/smartems/pkg/util"
)

func GetPendingOrgInvites(c *m.ReqContext) Response {
	query := m.GetTempUsersQuery{OrgId: c.OrgId, Status: m.TmpUserInvitePending}

	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get invites from db", err)
	}

	for _, invite := range query.Result {
		invite.Url = setting.ToAbsUrl("invite/" + invite.Code)
	}

	return JSON(200, query.Result)
}

func AddOrgInvite(c *m.ReqContext, inviteDto dtos.AddInviteForm) Response {
	if !inviteDto.Role.IsValid() {
		return Error(400, "Invalid role specified", nil)
	}

	// first try get existing user
	userQuery := m.GetUserByLoginQuery{LoginOrEmail: inviteDto.LoginOrEmail}
	if err := bus.Dispatch(&userQuery); err != nil {
		if err != m.ErrUserNotFound {
			return Error(500, "Failed to query db for existing user check", err)
		}
	} else {
		return inviteExistingUserToOrg(c, userQuery.Result, &inviteDto)
	}

	if setting.DisableLoginForm {
		return Error(400, "Cannot invite when login is disabled.", nil)
	}

	cmd := m.CreateTempUserCommand{}
	cmd.OrgId = c.OrgId
	cmd.Email = inviteDto.LoginOrEmail
	cmd.Name = inviteDto.Name
	cmd.Status = m.TmpUserInvitePending
	cmd.InvitedByUserId = c.UserId
	var err error
	cmd.Code, err = util.GetRandomString(30)
	if err != nil {
		return Error(500, "Could not generate random string", err)
	}
	cmd.Role = inviteDto.Role
	cmd.RemoteAddr = c.Req.RemoteAddr

	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to save invite to database", err)
	}

	// send invite email
	if inviteDto.SendEmail && util.IsEmail(inviteDto.LoginOrEmail) {
		emailCmd := m.SendEmailCommand{
			To:       []string{inviteDto.LoginOrEmail},
			Template: "new_user_invite.html",
			Data: map[string]interface{}{
				"Name":      util.StringsFallback2(cmd.Name, cmd.Email),
				"OrgName":   c.OrgName,
				"Email":     c.Email,
				"LinkUrl":   setting.ToAbsUrl("invite/" + cmd.Code),
				"InvitedBy": util.StringsFallback3(c.Name, c.Email, c.Login),
			},
		}

		if err := bus.Dispatch(&emailCmd); err != nil {
			if err == m.ErrSmtpNotEnabled {
				return Error(412, err.Error(), err)
			}
			return Error(500, "Failed to send email invite", err)
		}

		emailSentCmd := m.UpdateTempUserWithEmailSentCommand{Code: cmd.Result.Code}
		if err := bus.Dispatch(&emailSentCmd); err != nil {
			return Error(500, "Failed to update invite with email sent info", err)
		}

		return Success(fmt.Sprintf("Sent invite to %s", inviteDto.LoginOrEmail))
	}

	return Success(fmt.Sprintf("Created invite for %s", inviteDto.LoginOrEmail))
}

func inviteExistingUserToOrg(c *m.ReqContext, user *m.User, inviteDto *dtos.AddInviteForm) Response {
	// user exists, add org role
	createOrgUserCmd := m.AddOrgUserCommand{OrgId: c.OrgId, UserId: user.Id, Role: inviteDto.Role}
	if err := bus.Dispatch(&createOrgUserCmd); err != nil {
		if err == m.ErrOrgUserAlreadyAdded {
			return Error(412, fmt.Sprintf("User %s is already added to organization", inviteDto.LoginOrEmail), err)
		}
		return Error(500, "Error while trying to create org user", err)
	}

	if inviteDto.SendEmail && util.IsEmail(user.Email) {
		emailCmd := m.SendEmailCommand{
			To:       []string{user.Email},
			Template: "invited_to_org.html",
			Data: map[string]interface{}{
				"Name":      user.NameOrFallback(),
				"OrgName":   c.OrgName,
				"InvitedBy": util.StringsFallback3(c.Name, c.Email, c.Login),
			},
		}

		if err := bus.Dispatch(&emailCmd); err != nil {
			return Error(500, "Failed to send email invited_to_org", err)
		}
	}

	return Success(fmt.Sprintf("Existing Grafana user %s added to org %s", user.NameOrFallback(), c.OrgName))
}

func RevokeInvite(c *m.ReqContext) Response {
	if ok, rsp := updateTempUserStatus(c.Params(":code"), m.TmpUserRevoked); !ok {
		return rsp
	}

	return Success("Invite revoked")
}

// GetInviteInfoByCode gets a pending user invite corresponding to a certain code.
// A response containing an InviteInfo object is returned if the invite is found.
// If a (pending) invite is not found, 404 is returned.
func GetInviteInfoByCode(c *m.ReqContext) Response {
	query := m.GetTempUserByCodeQuery{Code: c.Params(":code")}
	if err := bus.Dispatch(&query); err != nil {
		if err == m.ErrTempUserNotFound {
			return Error(404, "Invite not found", nil)
		}
		return Error(500, "Failed to get invite", err)
	}

	invite := query.Result
	if invite.Status != m.TmpUserInvitePending {
		return Error(404, "Invite not found", nil)
	}

	return JSON(200, dtos.InviteInfo{
		Email:     invite.Email,
		Name:      invite.Name,
		Username:  invite.Email,
		InvitedBy: util.StringsFallback3(invite.InvitedByName, invite.InvitedByLogin, invite.InvitedByEmail),
	})
}

func (hs *HTTPServer) CompleteInvite(c *m.ReqContext, completeInvite dtos.CompleteInviteForm) Response {
	query := m.GetTempUserByCodeQuery{Code: completeInvite.InviteCode}

	if err := bus.Dispatch(&query); err != nil {
		if err == m.ErrTempUserNotFound {
			return Error(404, "Invite not found", nil)
		}
		return Error(500, "Failed to get invite", err)
	}

	invite := query.Result
	if invite.Status != m.TmpUserInvitePending {
		return Error(412, fmt.Sprintf("Invite cannot be used in status %s", invite.Status), nil)
	}

	cmd := m.CreateUserCommand{
		Email:        completeInvite.Email,
		Name:         completeInvite.Name,
		Login:        completeInvite.Username,
		Password:     completeInvite.Password,
		SkipOrgSetup: true,
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "failed to create user", err)
	}

	user := &cmd.Result

	if err := bus.Publish(&events.SignUpCompleted{
		Name:  user.NameOrFallback(),
		Email: user.Email,
	}); err != nil {
		return Error(500, "failed to publish event", err)
	}

	if ok, rsp := applyUserInvite(user, invite, true); !ok {
		return rsp
	}

	hs.loginUserWithUser(user, c)

	metrics.MApiUserSignUpCompleted.Inc()
	metrics.MApiUserSignUpInvite.Inc()

	return Success("User created and logged in")
}

func updateTempUserStatus(code string, status m.TempUserStatus) (bool, Response) {
	// update temp user status
	updateTmpUserCmd := m.UpdateTempUserStatusCommand{Code: code, Status: status}
	if err := bus.Dispatch(&updateTmpUserCmd); err != nil {
		return false, Error(500, "Failed to update invite status", err)
	}

	return true, nil
}

func applyUserInvite(user *m.User, invite *m.TempUserDTO, setActive bool) (bool, Response) {
	// add to org
	addOrgUserCmd := m.AddOrgUserCommand{OrgId: invite.OrgId, UserId: user.Id, Role: invite.Role}
	if err := bus.Dispatch(&addOrgUserCmd); err != nil {
		if err != m.ErrOrgUserAlreadyAdded {
			return false, Error(500, "Error while trying to create org user", err)
		}
	}

	// update temp user status
	if ok, rsp := updateTempUserStatus(invite.Code, m.TmpUserCompleted); !ok {
		return false, rsp
	}

	if setActive {
		// set org to active
		if err := bus.Dispatch(&m.SetUsingOrgCommand{OrgId: invite.OrgId, UserId: user.Id}); err != nil {
			return false, Error(500, "Failed to set org as active", err)
		}
	}

	return true, nil
}
