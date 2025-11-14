package factory

import (
	"github.com/lautarok/manosegura/src/infra/database"
	credentialsInterface "github.com/lautarok/manosegura/src/internal/modules/credentials/interfaces"
	credentialsRepository "github.com/lautarok/manosegura/src/internal/modules/credentials/repositories"
	permissionsInterface "github.com/lautarok/manosegura/src/internal/modules/permissions/interfaces"
	permissionsRepository "github.com/lautarok/manosegura/src/internal/modules/permissions/repositories"
	rolesInterface "github.com/lautarok/manosegura/src/internal/modules/roles/interfaces"
	rolesRepository "github.com/lautarok/manosegura/src/internal/modules/roles/repositories"
	usersInterface "github.com/lautarok/manosegura/src/internal/modules/users/interfaces"
	usersRepository "github.com/lautarok/manosegura/src/internal/modules/users/repositories"
)

type RepositoriesFactory struct {
	database *database.Database
}

func NewRepositoriesFactory(database *database.Database) *RepositoriesFactory {
	return &RepositoriesFactory{
		database: database,
	}
}

func (repositoriesFactory *RepositoriesFactory) UsersRepository() usersInterface.UsersRepository {
	return usersRepository.NewUsersRepository(repositoriesFactory.database)
}

func (repositoriesFactory *RepositoriesFactory) RolesRepository() rolesInterface.RolesRepository {
	return rolesRepository.NewRolesRepository(repositoriesFactory.database)
}

func (repositoriesFactory *RepositoriesFactory) CredentialsRepository() credentialsInterface.CredentialsRepository {
	return credentialsRepository.NewCredentialsRepository(repositoriesFactory.database)
}

func (repositoriesFactory *RepositoriesFactory) PermissionsRepository() permissionsInterface.PermissionsRepository {
	return permissionsRepository.NewPermissionsRepository(repositoriesFactory.database)
}
