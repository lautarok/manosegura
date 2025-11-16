package domain

import (
	"github.com/google/uuid"
	permissionsDomain "github.com/lautarok/manosegura/src/internal/modules/permissions/domain"
)

type RolePermission struct {
	ID           uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Role         *Role     `bun:"rel:belongs-to,join:role_id=id,type:uuid" json:"role"`
	RoleID       uuid.UUID
	Permission   *permissionsDomain.Permission `bun:"rel:belongs-to,join:permission_id=id,type:uuid" json:"permission"`
	PermissionID uuid.UUID
}
