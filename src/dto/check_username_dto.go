package dto

import (
	"github.com/lautarok/manosegura/src/infra"
)

type CheckUsername struct {
	Username string `validate:"required,min=6,max=26" example:"usuario.prueba" json:"username"`
}

func (dto *CheckUsername) Validate() error {
	return infra.Validate.Struct(dto)
}
