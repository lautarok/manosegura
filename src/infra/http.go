package infra

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/lautarok/manosegura/src/pkg"
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

	app.Use(ErrorHandlerMiddleware())

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

func ErrorHandlerMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err == nil {
			return nil
		}

		if strings.Contains(err.Error(), "validation") {
			return ctx.Status(http.StatusBadRequest).JSON(pkg.NewAppError(http.StatusBadRequest, err.Error()))
		}

		appErr := pkg.MapError(err)
		return ctx.Status(appErr.StatusCode).JSON(appErr)
	}
}
