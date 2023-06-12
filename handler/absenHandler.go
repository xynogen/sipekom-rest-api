package handler

import (
	"sipekom-rest-api/database"
	"sipekom-rest-api/model"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get all absen.
// @Description get all absen
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/absen/ [get]
func GetAllAbsen(c *fiber.Ctx) error {
	absens := new([]entity.Absensi)
	resp := new(response.Response)
	db := database.DB

	db.Find(&absens)
	resp.Status = model.StatusSuccess
	resp.Message = "Return All Absen"
	resp.Data = absens

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get absen.
// @Description get absen by id.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param id path int64 true "Absen ID"
// @Router /api/absen/get/{id} [get]
func GetAbsen(c *fiber.Ctx) error {
	absen := new(entity.Absensi)
	resp := new(response.Response)

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Status = model.StatusError
		resp.Message = "ID is Not Valid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	if err := db.Where("id = ?", id).First(&absen).Error; err != nil {
		resp.Status = model.StatusError
		resp.Message = "Absen not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = model.StatusSuccess
	resp.Message = "Absen is Found"
	resp.Data = absen
	return c.Status(fiber.StatusOK).JSON(resp)
}
