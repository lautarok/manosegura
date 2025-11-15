package users

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra/validation"
)

// swagger:model
type UpdateUserDto struct {
	ID      uuid.UUID `validate:"uuid" json:"-"`
	Name    string    `validate:"min=3,max=40" json:"name" example:"Daniel"`
	Surname string    `validate:"min=3,max=40" json:"surname" example:"Casals"`
	RoleID  uuid.UUID `validate:"uuid" json:"roleId" example:"cd1cd2d8-bff1-11f0-950b-ac8247de0e13"`
}

func (dto *UpdateUserDto) Validate() error {
	validate := validation.Validate

	if err := validate.Struct(dto); err != nil {
		return err
	}

	return nil
}
