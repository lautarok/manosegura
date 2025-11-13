package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
	"github.com/lautarok/manosegura/src/infra"
)

type CredentialsRepository struct {
	database *infra.Database
}

func NewCredentialsRepository(database *infra.Database) *CredentialsRepository {
	return &CredentialsRepository{
		database: database,
	}
}

func (credentialsRepository *CredentialsRepository) UsernameExists(username string) (bool, error) {
	exists, err := credentialsRepository.database.DB.
		NewSelect().
		Model((*domain.Credential)(nil)).
		Where("username = ?", username).
		Exists(context.Background())

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (credentialsRepository *CredentialsRepository) EmailExists(email string) (bool, error) {
	exists, err := credentialsRepository.database.DB.
		NewSelect().
		Model((*domain.Credential)(nil)).
		Where("email = ?", email).
		Exists(context.Background())

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (credentialsRepository *CredentialsRepository) CreateOne(credential *domain.Credential) (*domain.Credential, error) {
	_, err := credentialsRepository.database.DB.
		NewInsert().
		Model(credential).
		Returning("*").
		Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return credential, nil
}

func (credentialsRepository *CredentialsRepository) FindOneByUserId(userId uuid.UUID) (*domain.Credential, error) {
	var credential domain.Credential

	err := credentialsRepository.database.DB.
		NewSelect().
		Model(&credential).
		Where("user_id = ?", userId).
		Scan(context.Background())

	return &credential, err
}
