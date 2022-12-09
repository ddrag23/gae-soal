package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RoleAccess() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		roleName := claims["role_name"].(string)
		if roleName != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "Error", "message": "Anda tidak memiliki akses ke link ini", "data": nil})
		}
		return c.Next()
	}

}
