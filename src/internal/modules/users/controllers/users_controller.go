package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	middleware "github.com/lautarok/manosegura/src/internal/modules/auth/middlewares"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	_ "github.com/lautarok/manosegura/src/internal/modules/users/domain"
	usersDto "github.com/lautarok/manosegura/src/internal/modules/users/dto"
	"github.com/lautarok/manosegura/src/internal/modules/users/services"
)

type UsersController struct {
	path           string
	usersService   *services.UsersService
	authMiddleware *middleware.AuthMiddleware
}

type UsersControllerDeps struct {
	UsersService   *services.UsersService
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

// @Router get
// @Returns 200 []User
// @Returns 401 AppError example:"{\"statusCode\": 401, \"message\": \"unauthorized\"}"
// @Returns default AppError
// @QueryParams PaginationDto
// @BearerAuth
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

// @Router post
// @Returns 201 User
// @Returns 401 AppError example:"{\"statusCode\": 401, \"message\": \"unauthorized\"}"
// @Returns 409 AppError example:"{\"statusCode\": 409, \"message\": \"email already in use\"}"
// @Returns default AppError
// @BodyRequest CreateUserDto
// @BearerAuth
func (usersController *UsersController) CreateUser(ctx *fiber.Ctx) error {
	var dto usersDto.CreateUserDto
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

// @Router put
// @Returns 200 User
// @Returns 401 AppError example:"{\"statusCode\": 401, \"message\": \"unauthorized\"}"
// @Returns 409 AppError example:"{\"statusCode\": 409, \"message\": \"email already in use\"}"
// @Returns 404 AppError example:"{\"statusCode\": 404, \"message\": \"user not found\"}"
// @Returns default AppError
// @BodyRequest CreateUserDto
// @RouteParams IdDto
// @BearerAuth
func (usersController *UsersController) UpdateUser(ctx *fiber.Ctx) error {
	var dto usersDto.UpdateUserDto
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

// @Router delete
// @Returns 200 User
// @Returns 401 AppError example:"{\"statusCode\": 401, \"message\": \"unauthorized\"}"
// @Returns 404 AppError example:"{\"statusCode\": 404, \"message\": \"user not found\"}"
// @Returns default AppError
// @BodyRequest CreateUserDto
// @RouteParams IdDto
// @BearerAuth
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

// @Router get /[id]
// @Returns 200 User
// @Returns 401 AppError example:"{\"statusCode\": 401, \"message\": \"unauthorized\"}"
// @Returns 404 AppError example:"{\"statusCode\": 404, \"message\": \"user not found\"}"
// @Returns default AppError
// @RouteParams IdDto
// @BearerAuth
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
