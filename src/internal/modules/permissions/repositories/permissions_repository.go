package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra/database"
	"github.com/lautarok/manosegura/src/internal/modules/permissions/domain"
)

type PermissionsRepository struct {
	database *database.Database
}

func NewPermissionsRepository(database *database.Database) *PermissionsRepository {
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
