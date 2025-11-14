package interfaces

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	"github.com/lautarok/manosegura/src/internal/modules/roles/domain"
)

type RolesRepository interface {
	Exists(uuid uuid.UUID) (bool, error)
	FindAll(dto *dto.PaginationDto) (*[]domain.Role, error)
	FindByAlias(alias string) (*domain.Role, error)
}
