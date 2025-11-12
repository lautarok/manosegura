package domain

import (
	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID    `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Alias       string       `bun:"type:VARCHAR(30),notnull" json:"alias"`
	Permissions []Permission `bun:"m2m:role_permissions,join:Role=Permission" json:"permissions,omitempty"`
}
