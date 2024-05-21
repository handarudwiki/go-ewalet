package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/golang-ewalet/domain"
)

func Authenticate(userService domain.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")

		if !strings.Contains(token, "Bearer") {
			return c.SendStatus(401)
		}

		token = strings.Split(token, " ")[1]

		user, err := userService.ValidateToken(c.Context(), token)

		if err != nil {
			return c.SendStatus(401)
		}

		c.Locals("x-user", user)

		return c.Next()
	}
}
