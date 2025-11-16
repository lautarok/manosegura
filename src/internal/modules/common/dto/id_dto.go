package dto

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra/validation"
)

type IdDto struct {
	ID uuid.UUID `validate:"required,uuid" json:"id"`
}

func (dto *IdDto) Validate() error {
	return validation.Validate.Struct(dto)
}
