package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lautarok/manosegura/src/dto"
	"github.com/lautarok/manosegura/src/infra"
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

func (authService *AuthService) Login(dto *dto.LoginDto) (*string, error) {
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

	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":     user.ID,
		"subscriber": user.ID,
	})

	token, err := tokenData.SignedString([]byte(infra.Variables.JWT_SECRET))
	if err != nil {
		return nil, err
	}

	return &token, nil
}
