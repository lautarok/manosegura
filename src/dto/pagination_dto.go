package dto

import (
	"github.com/lautarok/manosegura/src/infra"
)

type PaginationDto struct {
	Page  int `validate:"min=1,max=30" json:"page"`
	Limit int `validate:"gte=1,lte=30" json:"limit"`
}

func (dto *PaginationDto) Validate() error {
	validate := infra.Validate

	if dto.Page == 0 {
		dto.Page = 1
	}

	if dto.Limit == 0 {
		dto.Limit = 20
	}

	if err := validate.Struct(dto); err != nil {
		return err
	}

	return nil
}
