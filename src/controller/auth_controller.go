package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/manosegura/src/controller/middleware"
	"github.com/lautarok/manosegura/src/dto"
	"github.com/lautarok/manosegura/src/service"
)

type AuthController struct {
	path           string
	authService    *service.AuthService
	authMiddleware *middleware.AuthMiddleware
}

type AuthControllerDeps struct {
	AuthService    *service.AuthService
	AuthMiddleware *middleware.AuthMiddleware
}

func NewAuthController(deps *AuthControllerDeps) *AuthController {
	return &AuthController{
		path:           "auth",
		authService:    deps.AuthService,
		authMiddleware: deps.AuthMiddleware,
	}
}

func (controller *AuthController) RegisterRoutes(app fiber.Router) {
	router := app.Group(controller.path)
	router.Post("login", controller.Login)
	router.Post("signup", controller.Signup)
}

// GetUserList godoc
// @Summary Authenticate
// @Description Login with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Router /auth/login [post]
// @Success 201 {object} dto.TokenResponseDto
// @Param request body dto.LoginDto true "User credentials"
func (usersController *AuthController) Login(ctx *fiber.Ctx) error {
	var body dto.LoginDto
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	token, err := usersController.authService.Login(&body)
	if err != nil {
		return err
	}

	ctx.Status(201)
	ctx.JSON(&dto.TokenResponseDto{
		Token: *token,
	})
	return nil
}

// GetUserList godoc
// @Summary Signup
// @Description Signup with user and credentials info
// @Tags Auth
// @Accept json
// @Produce json
// @Router /auth/signup [post]
// @Success 201 {object} domain.User
// @Param request body dto.SignupDto true "Body data"
func (usersController *AuthController) Signup(ctx *fiber.Ctx) error {
	var body dto.SignupDto
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	createdUser, err := usersController.authService.Signup(&body)
	if err != nil {
		return err
	}

	ctx.Status(201)
	ctx.JSON(createdUser)
	return nil
}
