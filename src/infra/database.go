package infra

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/lautarok/manosegura/src/domain"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

type Database struct {
	DB       bun.IDB
	migrator *migrate.Migrator
}

func NewDatabase(env *Env) *Database {
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
		(*domain.Credential)(nil),
		(*domain.RolePermission)(nil),
		(*domain.Permission)(nil),
		(*domain.Role)(nil),
		(*domain.User)(nil),
	)

	return &Database{
		DB:       db,
		migrator: migrator,
	}
}

func (database *Database) CreateMigration(context context.Context, name string) {
	database.migrator.CreateTxSQLMigrations(context, name)
}
