package utils

import (
	"sipekom-rest-api/model/static"

	"github.com/gofiber/fiber/v2"
)

func IsAdmin(c *fiber.Ctx) bool {
	userClaims := DecodeJWT(c)

	if userClaims.Role != static.RoleAdmin {
		return false
	}
	return true
}

func IsKonsulen(c *fiber.Ctx) bool {
	userClaims := DecodeJWT(c)

	if userClaims.Role != static.RoleKonsulen {
		return false
	}
	return true
}

func IsMahasiswa(c *fiber.Ctx) bool {
	userClaims := DecodeJWT(c)

	if userClaims.Role != static.RoleMahasiswa {
		return false
	}
	return true
}
