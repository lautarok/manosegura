package services

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/infra/database"
	"github.com/lautarok/manosegura/src/infra/factory"
	"github.com/lautarok/manosegura/src/internal/exceptions"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	credentialsDomain "github.com/lautarok/manosegura/src/internal/modules/credentials/domain"
	rolesService "github.com/lautarok/manosegura/src/internal/modules/roles/services"
	"github.com/lautarok/manosegura/src/internal/modules/users/domain"
	usersDto "github.com/lautarok/manosegura/src/internal/modules/users/dto"
	"github.com/lautarok/manosegura/src/internal/modules/users/interfaces"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepository interfaces.UsersRepository
	database        *database.Database
	rolesService    *rolesService.RolesService
}

type UsersServiceDeps struct {
	UsersRepository interfaces.UsersRepository
	Database        *database.Database
	RolesService    *rolesService.RolesService
}

func NewUsersService(deps *UsersServiceDeps) *UsersService {
	return &UsersService{
		usersRepository: deps.UsersRepository,
		database:        deps.Database,
		rolesService:    deps.RolesService,
	}
}

func (usersService *UsersService) GetUserList(dto *dto.PaginationDto) (*[]domain.User, error) {
	list, err := usersService.usersRepository.FindAll(dto)
	if err != nil {
		return &[]domain.User{}, err
	}

	return list, nil
}

func (usersService *UsersService) CreateOne(dto *usersDto.CreateUserDto) (*domain.User, error) {
	ctx := context.Background()

	var createdUser *domain.User

	err := usersService.database.DB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		repositoriesFactory := factory.NewRepositoriesFactory(&database.Database{
			DB: &tx,
		})

		usersRepository := repositoriesFactory.UsersRepository()
		credentialsRepository := repositoriesFactory.CredentialsRepository()
		rolesRepository := repositoriesFactory.RolesRepository()

		emailExists, err := credentialsRepository.EmailExists(dto.Email)
		if err != nil {
			return err
		}

		if emailExists {
			return exceptions.ErrEmailAlreadyExists
		}

		usernameExists, err := credentialsRepository.UsernameExists(dto.Username)
		if err != nil {
			return err
		}

		if usernameExists {
			return exceptions.ErrUsernameAlreadyExists
		}

		roleExists, err := rolesRepository.Exists(dto.RoleId)
		if err != nil {
			return err
		}

		if !roleExists {
			return exceptions.ErrRoleNotFound
		}

		createdUser, err = usersRepository.CreateOne(&domain.User{
			Name:    dto.Name,
			Surname: dto.Surname,
			RoleID:  dto.RoleId,
		})
		if err != nil {
			return err
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		_, err = credentialsRepository.CreateOne(&credentialsDomain.Credential{
			Email:    dto.Email,
			Password: string(passwordHash),
			Username: dto.Username,
			UserID:   createdUser.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (usersService *UsersService) Exists(id uuid.UUID) (bool, error) {
	return usersService.usersRepository.Exists(id)
}

func (usersService *UsersService) UpdateOne(dto *usersDto.UpdateUserDto) (*domain.User, error) {
	userExists, err := usersService.Exists(dto.ID)
	if err != nil {
		return nil, err
	} else if !userExists {
		return nil, exceptions.ErrUserNotFound
	}

	roleExists, err := usersService.rolesService.CheckExists(dto.RoleID)
	if err != nil {
		return nil, err
	} else if !roleExists {
		return nil, exceptions.ErrRoleNotFound
	}

	return usersService.usersRepository.UpdateOne(dto)
}

func (usersService *UsersService) DeleteOne(dto *dto.IdDto) error {
	exists, err := usersService.Exists(dto.ID)
	if err != nil {
		return err
	} else if !exists {
		return exceptions.ErrUserNotFound
	}

	err = usersService.usersRepository.DeleteOne(dto)
	return err
}

func (usersService *UsersService) FindOne(dto *dto.IdDto) (*domain.User, error) {
	exists, err := usersService.Exists(dto.ID)
	if err != nil {
		return nil, err
	} else if !exists {
		return nil, exceptions.ErrUserNotFound
	}

	return usersService.usersRepository.FindOne(dto)
}
