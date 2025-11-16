package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Variables *Env

type Env struct {
	HTTP_PORT    int
	POSTGRES_DSN string
	JWT_SECRET   string
	DISABLE_DOCS bool
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

	Variables = &Env{
		HTTP_PORT:    httpPort,
		POSTGRES_DSN: os.Getenv("POSTGRES_DSN"),
		JWT_SECRET:   os.Getenv("JWT_SECRET"),
		DISABLE_DOCS: os.Getenv("DISABLE_DOCS") == "true",
	}

	return Variables
}
