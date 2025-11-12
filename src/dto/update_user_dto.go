package dto

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra"
)

type UpdateUserDto struct {
	ID      uuid.UUID `validate:"uuid" json:"-"`
	Name    string    `validate:"min=3,max=40" json:"name" example:"Daniel"`
	Surname string    `validate:"min=3,max=40" json:"surname" example:"Casals"`
	RoleID  uuid.UUID `validate:"uuid" json:"roleId" example:"cd1cd2d8-bff1-11f0-950b-ac8247de0e13"`
}

func (dto *UpdateUserDto) Validate() error {
	validate := infra.Validate

	if err := validate.Struct(dto); err != nil {
		return err
	}

	return nil
}
