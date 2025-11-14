package repository

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
)

type CredentialsRepository interface {
	EmailExists(email string) (bool, error)
	UsernameExists(username string) (bool, error)
	CreateOne(*domain.Credential) (*domain.Credential, error)
	FindOneByUserId(userId uuid.UUID) (*domain.Credential, error)
	FindOneByUsername(username string) (*domain.Credential, error)
}
