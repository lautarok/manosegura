package services

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	"github.com/lautarok/manosegura/src/internal/modules/roles/domain"
	"github.com/lautarok/manosegura/src/internal/modules/roles/interfaces"
)

type RolesService struct {
	rolesRepository interfaces.RolesRepository
}

func NewRolesService(rolesRepository interfaces.RolesRepository) *RolesService {
	return &RolesService{
		rolesRepository: rolesRepository,
	}
}

func (rolesService *RolesService) CheckExists(id uuid.UUID) (bool, error) {
	return rolesService.rolesRepository.Exists(id)
}

func (rolesService *RolesService) FindAll(dto *dto.PaginationDto) (*[]domain.Role, error) {
	return rolesService.rolesRepository.FindAll(dto)
}
