package domain

import (
	"time"

	"github.com/google/uuid"
	rolesDomain "github.com/lautarok/manosegura/src/internal/modules/roles/domain"
)

// swagger:model
type User struct {
	ID        uuid.UUID         `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name      string            `bun:"type:VARCHAR(40),notnull" json:"name"`
	Surname   string            `bun:"type:VARCHAR(40),notnull" json:"surname"`
	CreatedAt time.Time         `bun:"default:current_timestamp,notnull" json:"createdAt"`
	Role      *rolesDomain.Role `bun:"rel:belongs-to,join:role_id=id" json:"role,omitempty"`
	RoleID    uuid.UUID         `bun:"role_id,type:uuid" json:"-"`
}
