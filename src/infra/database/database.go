package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/lautarok/manosegura/src/infra/env"
	credentialsDomain "github.com/lautarok/manosegura/src/internal/modules/credentials/domain"
	permissionsDomain "github.com/lautarok/manosegura/src/internal/modules/permissions/domain"
	rolesDomain "github.com/lautarok/manosegura/src/internal/modules/roles/domain"
	usersDomain "github.com/lautarok/manosegura/src/internal/modules/users/domain"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

type Database struct {
	DB       bun.IDB
	migrator *migrate.Migrator
}

func NewDatabase(env *env.Env) *Database {
	dsn := env.POSTGRES_DSN
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	ctx := context.Background()

	migrations := migrate.NewMigrations(
		migrate.WithMigrationsDirectory("./database/migrations"),
	)

	if err := migrations.Discover(os.DirFS("./database/migrations")); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	migrator := migrate.NewMigrator(
		db,
		migrations,
	)

	migrator.Init(ctx)
	migrator.Lock(ctx)

	_, err := migrator.Migrate(ctx)
	if err != nil {
		log.Fatal(err)
	}

	migrator.Unlock(ctx)

	db.RegisterModel(
		(*credentialsDomain.Credential)(nil),
		(*rolesDomain.RolePermission)(nil),
		(*permissionsDomain.Permission)(nil),
		(*rolesDomain.Role)(nil),
		(*usersDomain.User)(nil),
	)

	return &Database{
		DB:       db,
		migrator: migrator,
	}
}

func (database *Database) CreateMigration(context context.Context, name string) {
	database.migrator.CreateTxSQLMigrations(context, name)
}
