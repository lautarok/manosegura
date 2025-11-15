package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	usersDomain "github.com/lautarok/manosegura/src/internal/modules/users/domain"
)

// swagger:model
type Credential struct {
	ID        uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Email     string            `bun:",unique,type:VARCHAR(50)" json:"email"`
	Username  string            `bun:",unique,type:VARCHAR(26)" json:"username"`
	Password  string            `bun:"type:VARCHAR(60)" json:"-"`
	User      *usersDomain.User `bun:"rel:belongs-to,join:user_id=id" json:"user"`
	UserID    uuid.UUID         `bun:"user_id" json:"user_id"`
	CreatedAt time.Time         `bun:"default:current_timestamp,notnull" json:"createdAt"`
	UpdatedAt time.Time         `bun:"default:current_timestamp,notnull" json:"updatedAt"`
}

func (credential *Credential) BeforeUpdate(ctx context.Context) (context.Context, error) {
	credential.UpdatedAt = time.Now()
	return ctx, nil
}
