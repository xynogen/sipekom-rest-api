package utils

import "github.com/gofiber/fiber/v2"

func GetJWTFromHeader(c *fiber.Ctx) string {

	return c.Get("Authorization")[7:]
}
