package main

import (
	"github.com/lautarok/manosegura/src/controller"
	"github.com/lautarok/manosegura/src/controller/middleware"
	"github.com/lautarok/manosegura/src/infra"
	"github.com/lautarok/manosegura/src/repository/postgres"
	"github.com/lautarok/manosegura/src/service"
)

// @title Mano Segura API
// @version 1.0
// @description Documentation of Mano Segura backend API. Use the format: Bearer {token} in the Authorization header.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	env := infra.NewEnv()

	database := infra.NewDatabase(env)

	rolesRepository := postgres.NewRolesRepository(database)
	usersRepository := postgres.NewUsersRepository(database)
	credentialsRepository := postgres.NewCredentialsRepository(database)

	rolesService := service.NewRolesService(rolesRepository)
	usersService := service.NewUsersService(&service.UsersServiceDeps{
		UsersRepository: usersRepository,
		Database:        database,
		RolesService:    rolesService,
	})
	credentialsService := service.NewCredentialsService(credentialsRepository)
	authService := service.NewAuthService(&service.AuthServiceDeps{
		CredentialsService: credentialsService,
		UsersService:       usersService,
		Database:           database,
	})

	authMiddleware := middleware.NewAuthMiddleware(&middleware.AuthMiddlewareDeps{
		AuthService:  authService,
		UsersService: usersService,
	})

	usersController := controller.NewUsersController(&controller.UsersControllerDeps{
		AuthMiddleware: authMiddleware,
		UsersService:   usersService,
	})
	authController := controller.NewAuthController(&controller.AuthControllerDeps{
		AuthMiddleware: authMiddleware,
		AuthService:    authService,
	})
	rolesController := controller.NewRolesController(&controller.RolesControllerDeps{
		AuthMiddleware: authMiddleware,
		RolesService:   rolesService,
	})

	infra.NewHttp(&infra.HttpConfig{
		Env: env,
		Controllers: []infra.Controller{
			usersController,
			authController,
			rolesController,
		},
	})
}
