package http

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/lautarok/manosegura/docs"
	"github.com/lautarok/manosegura/src/infra/env"
	"github.com/lautarok/manosegura/src/internal/exceptions"
)

type Http struct {
	App *fiber.App
}

type Controller interface {
	RegisterRoutes(app fiber.Router)
}

type HttpConfig struct {
	Env         *env.Env
	Controllers []Controller
}

// @title Mano Segura API
// @version 1.0
// @description Documentation of Mano Segura backend API
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewHttp(config *HttpConfig) *Http {
	app := fiber.New()

	app.Use(ErrorHandlerMiddleware())

	for _, controller := range config.Controllers {
		controller.RegisterRoutes(app)
	}

	app.Get("/swagger/*", swagger.HandlerDefault)

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
			return ctx.Status(http.StatusBadRequest).JSON(exceptions.NewAppError(http.StatusBadRequest, err.Error()))
		}

		appErr := exceptions.MapError(err)
		return ctx.Status(appErr.StatusCode).JSON(appErr)
	}
}
