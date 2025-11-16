package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/manosegura/src/internal/exceptions"
	"github.com/lautarok/manosegura/src/internal/modules/auth/dto"
	"github.com/lautarok/manosegura/src/internal/modules/auth/middlewares"
	"github.com/lautarok/manosegura/src/internal/modules/auth/services"
	"github.com/lautarok/manosegura/src/internal/modules/users/domain"
)

type AuthController struct {
	path           string
	authService    *services.AuthService
	authMiddleware *middlewares.AuthMiddleware
}

type AuthControllerDeps struct {
	AuthService    *services.AuthService
	AuthMiddleware *middlewares.AuthMiddleware
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
	router.Get("me", controller.authMiddleware.AuthUser(false), controller.GetMyUser)
}

// @Router post /login
// @Returns 201 TokenResponseDto
// @Returns 401 AppError example:"{\"statusCode\": 401, \"message\": \"invalid credentials\"}"
// @Returns default AppError
// @BodyRequest LoginDto
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

// @Router post /signup
// @Returns 201 User
// @Returns 409 AppError example:"{\"statusCode\": 409, \"message\": \"email already in use\"}"
// @Returns default AppError
// @BodyRequest SignupDto
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

// @Router get /me
// @Returns 200 User
// @Returns 401 AppError example:"{\"statusCode\": 401, \"message\": \"token not found\"}"
// @Returns default AppError
// @BearerAuth
func (usersController *AuthController) GetMyUser(ctx *fiber.Ctx) error {
	user, ok := (ctx.Locals("auth user")).(*domain.User)
	if !ok {
		return exceptions.ErrUserNotFound
	}
	ctx.JSON(user)
	return nil
}
