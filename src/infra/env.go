package infra

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	HTTP_PORT    int
	POSTGRES_DSN string
}

func NewEnv() *Env {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	err := godotenv.Load(".env."+env, ".env")
	if err != nil {
		log.Fatalf("Error loading .env files: %v", err)
	}

	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		httpPort = 3000
	}

	return &Env{
		HTTP_PORT:    httpPort,
		POSTGRES_DSN: os.Getenv("POSTGRES_DSN"),
	}
}
