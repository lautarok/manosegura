package main

import (
	"context"
	"log"
	"os"

	"github.com/lautarok/manosegura/src/infra"
)

func main() {
	args := os.Args

	if len(args) < 3 || args[1] != "migration" || args[2] != "create" {
		log.Fatal("Try \"migration create\" command")
	}

	env := infra.NewEnv()
	db := infra.NewDatabase(env)
	ctx := context.Background()

	db.CreateMigration(ctx, args[3])
}
