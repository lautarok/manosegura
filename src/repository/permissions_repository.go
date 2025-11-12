package repository

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
)

type PermissionsRepository interface {
	Exists(uuid uuid.UUID) (bool, error)
	FindAll() (*[]domain.Permission, error)
}
