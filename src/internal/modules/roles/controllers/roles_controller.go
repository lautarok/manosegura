package controllers

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/manosegura/src/internal/modules/auth/middlewares"
	"github.com/lautarok/manosegura/src/internal/modules/common/dto"
	_ "github.com/lautarok/manosegura/src/internal/modules/roles/domain"
	"github.com/lautarok/manosegura/src/internal/modules/roles/services"
)

type RolesController struct {
	path           string
	rolesService   *services.RolesService
	authMiddleware *middlewares.AuthMiddleware
}

type RolesControllerDeps struct {
	RolesService   *services.RolesService
	AuthMiddleware *middlewares.AuthMiddleware
}

func NewRolesController(deps *RolesControllerDeps) *RolesController {
	return &RolesController{
		path:           "roles",
		rolesService:   deps.RolesService,
		authMiddleware: deps.AuthMiddleware,
	}
}

func (controller *RolesController) RegisterRoutes(app fiber.Router) {
	router := app.Group(controller.path)
	router.Get("", controller.authMiddleware.AuthUser(false, "manage all"), controller.GetRoles)
}

func (rolesController *RolesController) RegisterDocs(docs *openapi3.T) {

}

// GetRoles godoc
// @Summary Get roles
// @Description Get all roles
// @Tags Roles
// @Accept json
// @Produce json
// @Router get /roles
// @Success 200 {array} domain.Role
// @Param request query dto.PaginationDto false "Pagination"
// @Security BearerAuth
func (rolesController *RolesController) GetRoles(ctx *fiber.Ctx) error {
	var dto dto.PaginationDto
	err := ctx.QueryParser(&dto)
	if err != nil {
		return err
	}

	roleList, err := rolesController.rolesService.FindAll(&dto)
	if err != nil {
		return err
	}

	ctx.JSON(roleList)
	return nil
}
