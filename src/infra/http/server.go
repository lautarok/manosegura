package http

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/lautarok/manosegura/src/infra/env"
	"github.com/lautarok/manosegura/src/internal/docs"
	"github.com/lautarok/manosegura/src/internal/exceptions"
)

type Http struct {
	App *fiber.App
}

type Controller interface {
	RegisterRoutes(app fiber.Router)
	RegisterDocs(docs *openapi3.T)
}

type HttpConfig struct {
	Env           *env.Env
	Controllers   []Controller
	DocsGenerator *docs.DocsGenerator
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

	if !config.Env.DISABLE_DOCS {
		config.DocsGenerator.GenerateOas()

		app.Get("oas.json", func(ctx *fiber.Ctx) error {
			return ctx.JSON(config.DocsGenerator.Struct())
		})

		app.Get("/swagger/*", swagger.New(swagger.Config{
			URL:         "/oas.json",
			DeepLinking: true,
		}))

		app.Get("/redoc", func(ctx *fiber.Ctx) error {
			return ctx.SendFile("./html/redoc.html")
		})
	}

	for _, controller := range config.Controllers {
		controller.RegisterRoutes(app)
	}

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

		if fiberErr, ok := (err).(*fiber.Error); ok {
			return fiberErr
		}

		if strings.Contains(err.Error(), "validation") {
			return ctx.Status(http.StatusBadRequest).JSON(exceptions.NewAppError(http.StatusBadRequest, err.Error()))
		}

		appErr := exceptions.MapError(err)
		return ctx.Status(appErr.StatusCode).JSON(appErr)
	}
}
