package handler

import (
	"os"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get Photo [guest🔒].
// @Param photo_name path string true "Photo Name"
// @Description get Photo from photo name
// @Tags Photo
// @Success 200 {object} response.Response
// @Router /api/photo/{photo_name} [get]
func GetPhoto(c *fiber.Ctx) error {
	photoName := c.AllParams()["photo_name"]
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	if !utils.IsExist("data/photo/" + photoName) {
		resp.Message = "Photo not Exist"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	photoBytes, err := os.ReadFile("data/photo/" + photoName)
	if err != nil {
		resp.Message = "Failed Load Photo"
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	c.Set("content-type", "image/jpg")
	return c.Status(fiber.StatusOK).Send(photoBytes)
}
