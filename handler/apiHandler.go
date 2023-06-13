package handler

import (
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"

	"github.com/gofiber/fiber/v2"
)

// Hello godoc
// @Summary server status.
// @Description get the status of server.
// @Tags Root
// @Accept */*
// @Produce json
// @Success 200 {object} response.Response
// @Router /api [get]
func Hello(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusSuccess
	resp.Message = "Server is UP!"
	resp.Data = nil

	return c.Status(fiber.StatusOK).JSON(resp)
}
