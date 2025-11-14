package interfaces

import (
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	"github.com/lautarok/manosegura/src/internal/modules/users/domain"
	usersDto "github.com/lautarok/manosegura/src/internal/modules/users/dto"
)

type UsersRepository interface {
	FindAll(dto *dto.PaginationDto) (*[]domain.User, error)
	CreateOne(user *domain.User) (*domain.User, error)
	UpdateOne(dto *usersDto.UpdateUserDto) (*domain.User, error)
	DeleteOne(dto *dto.IdDto) error
	Exists(id uuid.UUID) (bool, error)
	FindOne(dto *dto.IdDto) (*domain.User, error)
}
