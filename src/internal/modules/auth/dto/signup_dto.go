package dto

import (
	"github.com/lautarok/manosegura/src/infra/validation"
	"github.com/lautarok/manosegura/src/internal/exceptions"
)

// swagger:model
type SignupDto struct {
	Name           string `validate:"required,min=3,max=40" json:"name" example:"Maia Abigail"`
	Surname        string `validate:"required,min=3,max=40" json:"surname" example:"Freytes"`
	Email          string `validate:"required,email" json:"email" example:"maiaamorosa1@gmail.com"`
	Username       string `validate:"required,min=6,max=26" json:"username" example:"maiaamorosa"`
	Password       string `validate:"required,password" json:"password" example:"#Manosegura.1"`
	RepeatPassword string `validate:"required,password" json:"repeatPassword" example:"#Manosegura.1"`
}

func (dto *SignupDto) Validate() error {
	err := validation.Validate.Struct(dto)
	if err != nil {
		return err
	} else if dto.Password != dto.RepeatPassword {
		return exceptions.ErrRepeatPasswordNotMatch
	}
	return nil
}
