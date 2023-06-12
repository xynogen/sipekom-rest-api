package handler

import (
	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get all Konsulen.
// @Description get all absen
// @Tags Konsulen
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/konsulen/ [get]
func GetAllKonsulen(c *fiber.Ctx) error {
	konsulens := new([]entity.Konsulen)
	resp := new(response.Response)
	db := database.DB

	db.Scopes(utils.Paginate(c)).Find(&konsulens)
	resp.Status = static.StatusSuccess
	resp.Message = "Return All Konsulen"
	resp.Data = konsulens

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get konsulen.
// @Description get konsulen by id.
// @Tags Konsulen
// @Produce json
// @Success 200 {object} response.Response
// @Param id path int64 true "ID"
// @Router /api/konsulen/get/{id} [get]
func GetKonsulen(c *fiber.Ctx) error {
	konsulen := new(entity.Konsulen)
	resp := new(response.Response)

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Status = static.StatusError
		resp.Message = "ID is Not Valid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB

	if err := db.Where("id = ?", id).Find(&konsulen).Error; err != nil {
		resp.Status = static.StatusError
		resp.Message = "Konsulen not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Konsulen is Found"
	resp.Data = konsulen
	return c.Status(fiber.StatusOK).JSON(resp)
}
