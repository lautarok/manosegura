package dto

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra"
)

type IdDto struct {
	ID uuid.UUID `validate:"required,uuid" json:"id"`
}

func (dto *IdDto) Validate() error {
	validate := infra.Validate

	if err := validate.Struct(dto); err != nil {
		return err
	}

	return nil
}
