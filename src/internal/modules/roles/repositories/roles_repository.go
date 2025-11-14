package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra/database"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	"github.com/lautarok/manosegura/src/internal/modules/roles/domain"
)

type RolesRepository struct {
	database *database.Database
}

func NewRolesRepository(database *database.Database) *RolesRepository {
	return &RolesRepository{
		database: database,
	}
}

func (rolesRepository *RolesRepository) FindAll(dto *dto.PaginationDto) (*[]domain.Role, error) {
	var roleList []domain.Role

	offset := (dto.Page - 1) * dto.Limit

	err := rolesRepository.database.DB.
		NewSelect().
		Limit(dto.Limit).
		Offset(offset).
		Model(&roleList).
		Scan(context.Background())

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
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

	return exists, err
}

func (rolesRepository *RolesRepository) FindByAlias(alias string) (*domain.Role, error) {
	var role domain.Role

	err := rolesRepository.database.DB.
		NewSelect().
		Model(&role).
		Where("alias = ?", alias).
		Scan(context.Background())

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &role, nil
}
