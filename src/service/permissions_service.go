package service

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/repository"
)

type PermissionsService struct {
	permissionsRepository repository.PermissionsRepository
}

func NewPermissionsService(permissionsRepository repository.PermissionsRepository) *PermissionsService {
	return &PermissionsService{
		permissionsRepository: permissionsRepository,
	}
}

func (permissionsService *PermissionsService) CheckExists(id uuid.UUID) (bool, error) {
	return permissionsService.permissionsRepository.Exists(id)
}
