package users

import "github.com/lautarok/manosegura/src/infra/validation"

// swagger:model
type CheckUsername struct {
	Username string `validate:"required,min=6,max=26" example:"usuario.prueba" json:"username"`
}

func (dto *CheckUsername) Validate() error {
	return validation.Validate.Struct(dto)
}
