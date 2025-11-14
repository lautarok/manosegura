package service

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
	"github.com/lautarok/manosegura/src/dto"
	"github.com/lautarok/manosegura/src/repository"
)

type RolesService struct {
	rolesRepository repository.RolesRepository
}

func NewRolesService(rolesRepository repository.RolesRepository) *RolesService {
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
