package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/domain"
	"github.com/lautarok/manosegura/src/infra"
)

type PermissionsRepository struct {
	database *infra.Database
}

func NewPermissionsRepository(database *infra.Database) *PermissionsRepository {
	return &PermissionsRepository{
		database: database,
	}
}

func (permissionsRepository *PermissionsRepository) FindAll() (*[]domain.Permission, error) {
	var permissionList []domain.Permission

	err := permissionsRepository.database.DB.
		NewSelect().
		Model(&permissionList).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return &permissionList, nil
}

func (permissionsRepository *PermissionsRepository) Exists(id uuid.UUID) (bool, error) {
	exists, err := permissionsRepository.database.DB.
		NewSelect().
		Model((*domain.Permission)(nil)).
		Where("id = ?", id).
		Exists(context.Background())

	if err != nil {
		return false, err
	}

	return exists, nil
}
