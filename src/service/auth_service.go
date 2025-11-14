package service

import (
	"context"
	"database/sql"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lautarok/manosegura/src/domain"
	"github.com/lautarok/manosegura/src/dto"
	"github.com/lautarok/manosegura/src/factory"
	"github.com/lautarok/manosegura/src/infra"
	"github.com/lautarok/manosegura/src/pkg"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	credentialsService *CredentialsService
	usersService       *UsersService
	database           *infra.Database
}

type AuthServiceDeps struct {
	CredentialsService *CredentialsService
	UsersService       *UsersService
	Database           *infra.Database
}

func NewAuthService(deps *AuthServiceDeps) *AuthService {
	return &AuthService{
		credentialsService: deps.CredentialsService,
		usersService:       deps.UsersService,
		database:           deps.Database,
	}
}

func (authService *AuthService) Login(dto *dto.LoginDto) (*string, error) {
	credential, err := authService.credentialsService.FindOneByUsername(dto.Username)
	if err != nil {
		return nil, err
	} else if credential == nil {
		return nil, pkg.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(credential.Password),
		[]byte(dto.Password),
	)

	if err != nil {
		return nil, pkg.ErrUnauthorized
	}

	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     credential.UserID,
		"subscriber": credential.UserID,
	})

	token, err := tokenData.SignedString([]byte(infra.Variables.JWT_SECRET))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (authService *AuthService) Signup(dto *dto.SignupDto) (*domain.User, error) {
	ctx := context.Background()

	var newUser *domain.User

	err := authService.database.DB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		repositoriesFactory := factory.NewRepositoriesFactory(&infra.Database{
			DB: &tx,
		})

		rolesRepository := repositoriesFactory.RolesRepository()
		credentialsRepository := repositoriesFactory.CredentialsRepository()
		usersRepository := repositoriesFactory.UsersRepository()

		usernameExists, err := credentialsRepository.UsernameExists(dto.Username)
		if err != nil {
			return err
		} else if usernameExists {
			return pkg.ErrUsernameAlreadyExists
		}

		emailExists, err := credentialsRepository.EmailExists(dto.Email)
		if err != nil {
			return err
		} else if emailExists {
			return pkg.ErrEmailAlreadyExists
		}

		role, err := rolesRepository.FindByAlias("regular user")
		if err != nil {
			return err
		}

		newUser, err = usersRepository.CreateOne(&domain.User{
			Name:    dto.Name,
			Surname: dto.Surname,
			RoleID:  role.ID,
		})
		if err != nil {
			return err
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		_, err = credentialsRepository.CreateOne(&domain.Credential{
			Email:    dto.Email,
			Username: dto.Username,
			Password: string(passwordHash),
			UserID:   newUser.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return newUser, err
}
