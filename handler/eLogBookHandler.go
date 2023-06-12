package handler

import (
	"strconv"

	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get all E-Log Book.
// @Description get all E-Log Book
// @Tags ELogBook
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/elogbook/ [get]
func GetAllELogBook(c *fiber.Ctx) error {
	eLogBooks := new([]entity.ELogBook)
	resp := new(response.Response)
	db := database.DB

	db.Scopes(utils.Paginate(c)).Find(&eLogBooks)
	resp.Status = static.StatusSuccess
	resp.Message = "Return All E-Log Book"
	resp.Data = eLogBooks

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get elogbook.
// @Description get elogbook by id user.
// @Tags ELogBook
// @Produce json
// @Success 200 {object} response.Response
// @Param id path int64 true "ID User"
// @Router /api/elogbook/get/{id} [get]
func GetELogBook(c *fiber.Ctx) error {
	eLogBook := new(entity.ELogBook)
	user := new(entity.User)
	resp := new(response.Response)

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Status = static.StatusError
		resp.Message = "ID is Not Valid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		resp.Status = static.StatusError
		resp.Message = "User not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	if user.Level != static.LevelMahasiswa {
		resp.Status = static.StatusError
		resp.Message = "User not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	if err := db.Where("id_user = ?", id).Scopes(utils.Paginate(c)).Find(&eLogBook).Error; err != nil {
		resp.Status = static.StatusError
		resp.Message = "E-Log Book not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "E-Log Book is Found"
	resp.Data = eLogBook
	return c.Status(fiber.StatusOK).JSON(resp)
}
