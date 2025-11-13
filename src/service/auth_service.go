package service

import (
	"github.com/lautarok/manosegura/src/domain"
	"github.com/lautarok/manosegura/src/dto"
	"github.com/lautarok/manosegura/src/pkg"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	credentialsService *CredentialsService
	usersService       *UsersService
}

type AuthServiceDeps struct {
	CredentialsService *CredentialsService
	UsersService       *UsersService
}

func NewAuthService(deps *AuthServiceDeps) *AuthService {
	return &AuthService{
		credentialsService: deps.CredentialsService,
		usersService:       deps.UsersService,
	}
}

func (authService *AuthService) Login(dto *dto.LoginDto) (*domain.User, error) {
	user, err := authService.usersService.FindOneByUsername(dto.Username)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, pkg.ErrUserNotFound
	}

	credentials, err := authService.credentialsService.FindOneByUserId(user.ID)
	if err != nil {
		return nil, err
	} else if credentials == nil {
		return nil, pkg.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(credentials.Password),
		[]byte(dto.Password),
	)

	if err != nil {
		return nil, pkg.ErrUnauthorized
	}

	return user, nil
}
