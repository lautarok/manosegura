package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/manosegura/src/dto"
	"github.com/lautarok/manosegura/src/service"
)

type AuthController struct {
	path        string
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		path:        "auth",
		authService: authService,
	}
}

func (controller *AuthController) RegisterRoutes(app fiber.Router) {
	router := app.Group(controller.path)
	router.Post("login", controller.Login)
}

// GetUserList godoc
// @Summary Authenticate
// @Description Login with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Router /auth/login [post]
// @Success 200 {object} domain.User
// @Param request body dto.LoginDto true "User credentials"
func (usersController *AuthController) Login(ctx *fiber.Ctx) error {
	var dto dto.LoginDto
	err := ctx.BodyParser(&dto)
	if err != nil {
		return err
	} else if err := dto.Validate(); err != nil {
		return err
	}

	token, err := usersController.authService.Login(&dto)
	if err != nil {
		return err
	}

	ctx.JSON(map[string]string{
		"token": *token,
	})
	return nil
}
