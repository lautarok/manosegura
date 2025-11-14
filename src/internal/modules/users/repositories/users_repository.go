package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra/database"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	"github.com/lautarok/manosegura/src/internal/modules/users/domain"
	usersDto "github.com/lautarok/manosegura/src/internal/modules/users/dto"
)

type UsersRepository struct {
	database *database.Database
}

func NewUsersRepository(database *database.Database) *UsersRepository {
	return &UsersRepository{
		database: database,
	}
}

func (usersRepository *UsersRepository) FindAll(dto *dto.PaginationDto) (*[]domain.User, error) {
	var users []domain.User

	offset := (dto.Page - 1) * dto.Limit

	err := usersRepository.database.DB.
		NewSelect().
		Model(&users).
		Relation("Role").
		Relation("Role.Permissions").
		Limit(dto.Limit).
		Offset(offset).
		Scan(context.Background())

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &users, nil
}

func (usersRepository *UsersRepository) CreateOne(user *domain.User) (*domain.User, error) {
	err := usersRepository.database.DB.NewInsert().
		Model(user).
		Returning("*").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (usersRepository *UsersRepository) UpdateOne(dto *usersDto.UpdateUserDto) (*domain.User, error) {
	user := domain.User{
		ID:      dto.ID,
		Name:    dto.Name,
		Surname: dto.Surname,
		RoleID:  dto.RoleID,
	}

	err := usersRepository.database.DB.NewUpdate().
		Model(&user).
		Where("id = ?", dto.ID).
		Returning("*").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (usersRepository *UsersRepository) DeleteOne(dto *dto.IdDto) error {
	_, err := usersRepository.database.DB.NewDelete().
		Model((*domain.User)(nil)).
		Where("id = ?", dto.ID).
		Exec(context.Background())

	return err
}

func (usersRepository *UsersRepository) Exists(id uuid.UUID) (bool, error) {
	return usersRepository.database.DB.
		NewSelect().
		Model((*domain.User)(nil)).
		Where("id = ?", id).
		Exists(context.Background())
}

func (usersRepository *UsersRepository) FindOne(dto *dto.IdDto) (*domain.User, error) {
	var user domain.User

	user.ID = dto.ID

	err := usersRepository.database.DB.
		NewSelect().
		Model(&user).
		WherePK().
		Relation("Role").
		Relation("Role.Permissions").
		Scan(context.Background())

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &user, nil
}
