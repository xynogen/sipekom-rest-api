package utils

import (
	"sipekom-rest-api/model/static"

	"github.com/gofiber/fiber/v2"
)

func IsAdmin(c *fiber.Ctx) bool {
	jwtToken := GetJWTFromHeader(c)
	userClaims := DecodeJWT(jwtToken)

	if userClaims.Role != static.RoleAdmin {
		return false
	}
	return true
}
