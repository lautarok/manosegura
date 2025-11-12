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
	rolesService := service.NewRolesService(rolesRepository)

	usersRepository := postgres.NewUsersRepository(database)
	usersService := service.NewUsersService(&service.UsersServiceDeps{
		UsersRepository: usersRepository,
		Database:        database,
		RolesService:    rolesService,
	})
	usersController := controller.NewUsersController(&controller.UsersControllerDeps{
		UsersService: usersService,
		RolesService: rolesService,
	})

	infra.NewHttp(&infra.HttpConfig{
		Env: env,
		Controllers: []infra.Controller{
			usersController,
		},
	})
}
