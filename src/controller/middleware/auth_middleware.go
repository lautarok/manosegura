package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lautarok/manosegura/src/dto"
	"github.com/lautarok/manosegura/src/pkg"
	"github.com/lautarok/manosegura/src/service"
)

type AuthMiddleware struct {
	usersService *service.UsersService
	authService  *service.AuthService
}

type AuthMiddlewareDeps struct {
	UsersService *service.UsersService
	AuthService  *service.AuthService
}

func NewAuthMiddleware(deps *AuthMiddlewareDeps) *AuthMiddleware {
	return &AuthMiddleware{
		usersService: deps.UsersService,
		authService:  deps.AuthService,
	}
}

func (authMiddleware *AuthMiddleware) AuthUser(owner bool, permissionAlias ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return pkg.ErrBearerTokenNotFound
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		tokenPayload, err := authMiddleware.authService.GetAuthTokenPayload(tokenStr)
		if err != nil {
			return err
		}

		user, err := authMiddleware.usersService.FindOne(&dto.IdDto{
			ID: tokenPayload.UserID,
		})
		if err != nil {
			return err
		}

		validRole := false
		if owner {
			idParam, err := uuid.Parse(ctx.Params("id"))
			if err != nil {
				return err
			} else if user.ID == idParam {
				validRole = true
			}
		}

		if !validRole {
			if len(permissionAlias) > 0 {
				for _, requiredPermissionAlias := range permissionAlias {
					for _, userPermission := range user.Role.Permissions {
						if requiredPermissionAlias == userPermission.Alias {
							validRole = true
							break
						}
					}
					if validRole {
						break
					}
				}
			} else {
				validRole = true
			}
		}

		if !validRole {
			return pkg.ErrInsufficientPermissions
		}

		ctx.Locals("auth user", user)

		return ctx.Next()
	}
}
