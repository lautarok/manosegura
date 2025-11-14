package domain

import (
	"github.com/google/uuid"
	permissionsDomain "github.com/lautarok/manosegura/src/internal/modules/permissions/domain"
)

type Role struct {
	ID          uuid.UUID                      `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Alias       string                         `bun:"type:VARCHAR(30),notnull" json:"alias" example:"Human resources"`
	Permissions []permissionsDomain.Permission `bun:"m2m:role_permissions,join:Role=Permission" json:"permissions,omitempty"`
}
