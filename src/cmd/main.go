package main

import (
	_ "github.com/lautarok/manosegura/docs"
	"github.com/lautarok/manosegura/src/controller"
	"github.com/lautarok/manosegura/src/infra"
	"github.com/lautarok/manosegura/src/repository/postgres"
	"github.com/lautarok/manosegura/src/service"
)

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
	})

	usersController := controller.NewUsersController(usersService)
	authController := controller.NewAuthController(authService)

	infra.NewHttp(&infra.HttpConfig{
		Env: env,
		Controllers: []infra.Controller{
			usersController,
			authController,
		},
	})
}
