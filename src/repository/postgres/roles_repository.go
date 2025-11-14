package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
	"github.com/lautarok/manosegura/src/infra"
)

type RolesRepository struct {
	database *infra.Database
}

func NewRolesRepository(database *infra.Database) *RolesRepository {
	return &RolesRepository{
		database: database,
	}
}

func (rolesRepository *RolesRepository) FindAll() (*[]domain.Role, error) {
	var roleList []domain.Role

	err := rolesRepository.database.DB.
		NewSelect().
		Model(&roleList).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return &roleList, nil
}

func (rolesRepository *RolesRepository) Exists(id uuid.UUID) (bool, error) {
	exists, err := rolesRepository.database.DB.
		NewSelect().
		Model((*domain.Role)(nil)).
		Where("id = ?", id).
		Exists(context.Background())

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (rolesRepository *RolesRepository) FindByAlias(alias string) (*domain.Role, error) {
	var role *domain.Role

	err := rolesRepository.database.DB.
		NewSelect().
		Model(&role).
		Where("alias = ?", alias).
		Scan(context.Background())

	return role, err
}
