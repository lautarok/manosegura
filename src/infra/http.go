package infra

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Http struct {
	App *fiber.App
}

type Controller interface {
	RegisterRoutes(app fiber.Router)
}

type HttpConfig struct {
	Env         *Env
	Controllers []Controller
}

func NewHttp(config *HttpConfig) *Http {
	app := fiber.New()

	for _, controller := range config.Controllers {
		controller.RegisterRoutes(app)
	}

	app.Get("/swagger/*", swagger.New(swagger.Config{
		Title: "Mano Segura API documentation",
	}))

	err := app.Listen(":" + strconv.Itoa((config.Env.HTTP_PORT)))
	if err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}

	return &Http{
		App: app,
	}
}
