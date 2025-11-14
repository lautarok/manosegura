package interfaces

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/internal/modules/permissions/domain"
)

type PermissionsRepository interface {
	Exists(uuid uuid.UUID) (bool, error)
	FindAll() (*[]domain.Permission, error)
}
