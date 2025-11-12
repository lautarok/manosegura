package factory

import (
	"github.com/lautarok/manosegura/src/infra"
	"github.com/lautarok/manosegura/src/repository"
	"github.com/lautarok/manosegura/src/repository/postgres"
)

type RepositoriesFactory struct {
	database *infra.Database
}

func NewRepositoriesFactory(database *infra.Database) *RepositoriesFactory {
	return &RepositoriesFactory{
		database: database,
	}
}

func (repositoriesFactory *RepositoriesFactory) UsersRepository() repository.UsersRepository {
	return postgres.NewUsersRepository(repositoriesFactory.database)
}

func (repositoriesFactory *RepositoriesFactory) RolesRepository() repository.RolesRepository {
	return postgres.NewRolesRepository(repositoriesFactory.database)
}

func (repositoriesFactory *RepositoriesFactory) CredentialsRepository() repository.CredentialsRepository {
	return postgres.NewCredentialsRepository(repositoriesFactory.database)
}

func (repositoriesFactory *RepositoriesFactory) PermissionsRepository() repository.PermissionsRepository {
	return postgres.NewPermissionsRepository(repositoriesFactory.database)
}
