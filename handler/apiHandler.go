package handler

import "github.com/gofiber/fiber/v2"

// Hello godoc
// @Summary server status.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api [get]
func Hello(c *fiber.Ctx) error {
	res := c.JSON(fiber.Map{"status": "success", "message": "Server is UP!", "data": nil})

	if err := res; err != nil {
		return err
	}
	return nil
}
