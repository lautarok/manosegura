package repository

import "github.com/lautarok/manosegura/src/domain"

type CredentialsRepository interface {
	EmailExists(email string) (bool, error)
	UsernameExists(username string) (bool, error)
	CreateOne(*domain.Credential) (*domain.Credential, error)
}
