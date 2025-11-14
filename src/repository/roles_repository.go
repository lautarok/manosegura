package repository

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
	"github.com/lautarok/manosegura/src/dto"
)

type RolesRepository interface {
	Exists(uuid uuid.UUID) (bool, error)
	FindAll(dto *dto.PaginationDto) (*[]domain.Role, error)
	FindByAlias(alias string) (*domain.Role, error)
}
