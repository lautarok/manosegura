package dto

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra"
)

type CreateUserDto struct {
	Username string    `validate:"required,min=6,max=26" example:"usuario.prueba" json:"username"`
	Name     string    `validate:"required,min=3,max=40" json:"name" example:"Lautaro"`
	Surname  string    `validate:"required,min=3,max=40" json:"surname" example:"Kazalukian"`
	Email    string    `validate:"required,email" json:"email" example:"lautaro.kazalukian@manosegura2.com"`
	Password string    `validate:"required,min=3,max=40" json:"password" example:"#Manosegura.1"`
	RoleId   uuid.UUID `validate:"required,uuid" json:"roleId" example:"cd1cd2d8-bff1-11f0-950b-ac8247de0e13"`
}

func (dto *CreateUserDto) Validate() error {
	return infra.Validate.Struct(dto)
}
