package middleware

import (
	"net/http"
	"github.com/labstack/echo/v5"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
	RoleGuest = "guest"
)

func RequiresRole(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			userRoleVal := c.Get("user_role")
			userRole, ok := userRoleVal.(string)
			if !ok {
				userRole = RoleGuest
			}

			isAllowed := false
			for _, role := range allowedRoles {
				if userRole == role {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "ACCESS_DENIED",
				})
			}

			return next(c)
		}
	}
}