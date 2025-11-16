package dto

import "github.com/lautarok/manosegura/src/infra/validation"

type LoginDto struct {
	Username string `validate:"required,min=6,max=26" example:"usuario.prueba" json:"username"`
	Password string `validate:"required,min=3,max=40" json:"password" example:"#Manosegura.1"`
}

func (dto *LoginDto) Validate() error {
	return validation.Validate.Struct(dto)
}
