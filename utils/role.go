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

func IsKonsulen(c *fiber.Ctx) bool {
	jwtToken := GetJWTFromHeader(c)
	userClaims := DecodeJWT(jwtToken)

	if userClaims.Role != static.RoleKonsulen {
		return false
	}
	return true
}

func IsMahasiswa(c *fiber.Ctx) bool {
	jwtToken := GetJWTFromHeader(c)
	userClaims := DecodeJWT(jwtToken)

	if userClaims.Role != static.RoleMahasiswa {
		return false
	}
	return true
}
