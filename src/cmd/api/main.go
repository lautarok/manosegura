package main

import (
	"github.com/lautarok/manosegura/src/infra/database"
	"github.com/lautarok/manosegura/src/infra/env"
	"github.com/lautarok/manosegura/src/infra/http"
	"github.com/lautarok/manosegura/src/internal/docs"
	authControllerPackage "github.com/lautarok/manosegura/src/internal/modules/auth/controllers"
	authMiddleware "github.com/lautarok/manosegura/src/internal/modules/auth/middlewares"
	authServicePackage "github.com/lautarok/manosegura/src/internal/modules/auth/services"
	credentialsRepositoryPackage "github.com/lautarok/manosegura/src/internal/modules/credentials/repositories"
	credentialsServicePackage "github.com/lautarok/manosegura/src/internal/modules/credentials/services"
	rolesControllerPackage "github.com/lautarok/manosegura/src/internal/modules/roles/controllers"
	rolesRepositoryPackage "github.com/lautarok/manosegura/src/internal/modules/roles/repositories"
	rolesServicePackage "github.com/lautarok/manosegura/src/internal/modules/roles/services"
	usersControllerPackage "github.com/lautarok/manosegura/src/internal/modules/users/controllers"
	usersRepositoryPackage "github.com/lautarok/manosegura/src/internal/modules/users/repositories"
	usersServicePackage "github.com/lautarok/manosegura/src/internal/modules/users/services"
)

// @title Mano Segura API
// @version 1.0
// @description Documentation of Mano Segura backend API. Use the format: Bearer {token} in the Authorization header.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	env := env.NewEnv()

	database := database.NewDatabase(env)

	rolesRepository := rolesRepositoryPackage.NewRolesRepository(database)
	usersRepository := usersRepositoryPackage.NewUsersRepository(database)
	credentialsRepository := credentialsRepositoryPackage.NewCredentialsRepository(database)

	rolesService := rolesServicePackage.NewRolesService(rolesRepository)
	usersService := usersServicePackage.NewUsersService(&usersServicePackage.UsersServiceDeps{
		UsersRepository: usersRepository,
		Database:        database,
		RolesService:    rolesService,
	})
	credentialsService := credentialsServicePackage.NewCredentialsService(credentialsRepository)
	authService := authServicePackage.NewAuthService(&authServicePackage.AuthServiceDeps{
		CredentialsService: credentialsService,
		UsersService:       usersService,
		Database:           database,
	})

	authMiddleware := authMiddleware.NewAuthMiddleware(&authMiddleware.AuthMiddlewareDeps{
		AuthService:  authService,
		UsersService: usersService,
	})

	usersController := usersControllerPackage.NewUsersController(&usersControllerPackage.UsersControllerDeps{
		AuthMiddleware: authMiddleware,
		UsersService:   usersService,
	})
	authController := authControllerPackage.NewAuthController(&authControllerPackage.AuthControllerDeps{
		AuthMiddleware: authMiddleware,
		AuthService:    authService,
	})
	rolesController := rolesControllerPackage.NewRolesController(&rolesControllerPackage.RolesControllerDeps{
		AuthMiddleware: authMiddleware,
		RolesService:   rolesService,
	})

	docsGenerator := docs.NewDocsGenerator(&docs.DocsGeneratorConfig{
		Title:       "Mano Segura",
		Description: "Mano Segura API documentation",
		Version:     "1.0.0",
	})

	http.NewHttp(&http.HttpConfig{
		Env:           env,
		DocsGenerator: docsGenerator,
		Controllers: []http.Controller{
			usersController,
			authController,
			rolesController,
		},
	})
}
