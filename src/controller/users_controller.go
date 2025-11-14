package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/controller/middleware"
	"github.com/lautarok/manosegura/src/dto"
	"github.com/lautarok/manosegura/src/service"
)

type UsersController struct {
	path           string
	usersService   *service.UsersService
	authMiddleware *middleware.AuthMiddleware
}

type UsersControllerDeps struct {
	UsersService   *service.UsersService
	AuthMiddleware *middleware.AuthMiddleware
}

func NewUsersController(deps *UsersControllerDeps) *UsersController {
	return &UsersController{
		path:           "users",
		usersService:   deps.UsersService,
		authMiddleware: deps.AuthMiddleware,
	}
}

func (controller *UsersController) RegisterRoutes(app fiber.Router) {
	router := app.Group(controller.path)
	router.Get("", controller.authMiddleware.AuthUser(false, "manage all"), controller.GetUsers)
	router.Post("", controller.authMiddleware.AuthUser(false, "manage all"), controller.CreateUser)
	router.Put(":id", controller.authMiddleware.AuthUser(true, "manage all"), controller.UpdateUser)
	router.Delete(":id", controller.authMiddleware.AuthUser(false, "manage all"), controller.DeleteUser)
	router.Get(":id", controller.authMiddleware.AuthUser(true, "manage all"), controller.GetUser)
}

// GetUserList godoc
// @Summary Get users
// @Description Get all user list
// @Tags Users
// @Accept json
// @Produce json
// @Router /users [get]
// @Success 200 {array} domain.User
// @Param request query dto.PaginationDto false "Pagination"
// @Security BearerAuth
func (usersController *UsersController) GetUsers(ctx *fiber.Ctx) error {
	var dto dto.PaginationDto
	err := ctx.QueryParser(&dto)
	if err != nil {
		return err
	}

	userList, err := usersController.usersService.GetUserList(&dto)
	if err != nil {
		return err
	}

	ctx.JSON(userList)
	return nil
}

// CreateUser godoc
// @Summary Post user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Router /users [post]
// @Success 201 {object} domain.User
// @Param request body dto.CreateUserDto true "User data"
// @Security BearerAuth
func (usersController *UsersController) CreateUser(ctx *fiber.Ctx) error {
	var dto dto.CreateUserDto
	ctx.BodyParser(&dto)

	if err := dto.Validate(); err != nil {
		ctx.Status(400)
		return err
	}

	userCreated, err := usersController.usersService.CreateOne(&dto)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusCreated)
	ctx.JSON(userCreated)

	return nil
}

// UpdateUser godoc
// @Summary Update user
// @Description Update existing user data
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body dto.UpdateUserDto true "New user data"
// @Success 200 {object} domain.User
// @Router /users/{id} [put]
// @Security BearerAuth
func (usersController *UsersController) UpdateUser(ctx *fiber.Ctx) error {
	var dto dto.UpdateUserDto
	ctx.BodyParser(&dto)

	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}
	dto.ID = id

	if err := dto.Validate(); err != nil {
		ctx.Status(400)
		return err
	}

	userCreated, err := usersController.usersService.UpdateOne(&dto)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusOK)
	ctx.JSON(userCreated)

	return nil
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Router /users/{id} [delete]
// @Security BearerAuth
func (usersController *UsersController) DeleteUser(ctx *fiber.Ctx) error {
	var dto dto.IdDto
	ctx.ParamsParser(&dto)

	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}
	dto.ID = id

	if err := dto.Validate(); err != nil {
		ctx.Status(400)
		return err
	}

	err = usersController.usersService.DeleteOne(&dto)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusOK)

	return nil
}

// GetUser godoc
// @Summary Get single user
// @Description Get existing user data
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Router /users/{id} [get]
// @Security BearerAuth
func (usersController *UsersController) GetUser(ctx *fiber.Ctx) error {
	var dto dto.IdDto
	ctx.ParamsParser(&dto)

	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}
	dto.ID = id

	if err := dto.Validate(); err != nil {
		ctx.Status(400)
		return err
	}

	user, err := usersController.usersService.FindOne(&dto)
	if err != nil {
		return err
	}

	ctx.Status(http.StatusOK)
	ctx.JSON(user)

	return nil
}
