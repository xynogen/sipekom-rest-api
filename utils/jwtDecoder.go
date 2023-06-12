package utils

import (
	"log"
	"sipekom-rest-api/config"
	"sipekom-rest-api/model/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetJWTFromHeader(c *fiber.Ctx) string {
	if c.Get("Authorization") == "" {
		return ""
	}

	return c.Get("Authorization")[7:]
}

// Assuming JWT is Checked
func DecodeJWT(tokenStr string) response.Claims {
	var claims response.Claims

	token, _ := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		secretText := config.Env(config.SECRET)
		return []byte(secretText), nil
	})

	if !token.Valid {
		log.Fatal("invalid token")
	}

	return claims
}
