package domain

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID    uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Alias string    `bun:"type:VARCHAR(30),notnull" json:"alias" example:"manage users"`
}
