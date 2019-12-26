package dtos

import (
	m "github.com/smartems/smartems/pkg/models"
)

type UpdateDashboardAclCommand struct {
	Items []DashboardAclUpdateItem `json:"items"`
}

type DashboardAclUpdateItem struct {
	UserId     int64            `json:"userId"`
	TeamId     int64            `json:"teamId"`
	Role       *m.RoleType      `json:"role,omitempty"`
	Permission m.PermissionType `json:"permission"`
}
