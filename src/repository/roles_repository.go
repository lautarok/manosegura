package repository

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
)

type RolesRepository interface {
	Exists(uuid uuid.UUID) (bool, error)
	FindAll() (*[]domain.Role, error)
}
