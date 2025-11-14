package services

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/internal/modules/permissions/interfaces"
)

type PermissionsService struct {
	permissionsRepository interfaces.PermissionsRepository
}

func NewPermissionsService(permissionsRepository interfaces.PermissionsRepository) *PermissionsService {
	return &PermissionsService{
		permissionsRepository: permissionsRepository,
	}
}

func (permissionsService *PermissionsService) CheckExists(id uuid.UUID) (bool, error) {
	return permissionsService.permissionsRepository.Exists(id)
}
