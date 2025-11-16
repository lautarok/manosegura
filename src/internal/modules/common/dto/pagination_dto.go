package dto

import "github.com/lautarok/manosegura/src/infra/validation"

type PaginationDto struct {
	Page  int `json:"page"`
	Limit int `validate:"gte=1,lte=30" json:"limit"`
}

func (dto *PaginationDto) Validate() error {
	validate := validation.Validate

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
