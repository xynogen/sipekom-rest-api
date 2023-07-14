package handler

import (
	"os"
	"sipekom-rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get Photo [guest🔒].
// @Param photo_name path string true "Photo Name"
// @Description get Photo from photo name
// @Tags Photo
// @Success 200
// @Router /api/photo/{photo_name} [get]
func GetPhoto(c *fiber.Ctx) error {
	placeholder, err := os.ReadFile("data/photo/placeholder.png")
	photoName := c.AllParams()["photo_name"]

	if !utils.IsExist("data/photo/" + photoName) {
		c.Set("content-type", "image/png")
		return c.Status(fiber.StatusOK).Send(placeholder)
	}

	photoBytes, err := os.ReadFile("data/photo/" + photoName)
	if err != nil {
		c.Set("content-type", "image/png")
		return c.Status(fiber.StatusOK).Send(placeholder)
	}

	c.Set("content-type", "image/jpg")
	return c.Status(fiber.StatusOK).Send(photoBytes)
}
