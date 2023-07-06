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
// @Summary get all PPDS [guest🔒].
// @Description get all PPDS
// @Tags PPDS
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/ppds/ [get]
func GetAllPPDS(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	ppds := new([]entity.PPDS)
	db := database.DB

	db.Scopes(utils.Paginate(c)).Find(&ppds)

	resp.Status = static.StatusSuccess
	resp.Message = "Return All PPDS"
	resp.Data = ppds

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get all PPDS [guest🔒].
// @Description get all PPDS by id
// @Tags PPDS
// @Produce json
// @Success 200 {object} response.Response
// @Param id_user path int64 true "User ID"
// @Router /api/ppds/{id_user} [get]
func GetPPDS(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id_user, err := strconv.Atoi(c.AllParams()["id_user"])
	if err != nil || id_user < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	ppds := new(entity.PPDS)
	db := database.DB

	if db.Where("id_user = ?", id_user).First(&ppds).RowsAffected < 1 {
		resp.Message = "PPDS not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "PPDS is Found"
	resp.Data = ppds

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get Autocomplete for PPDS [guest🔒].
// @Description get Autocomplete for PPDS
// @Tags PPDS
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/ppds/auto [get]
func GetAutoCompletePPDS(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	type AutoCompletePPDS struct {
		IDUser uint   `json:"id_user"`
		Name   string `json:"name"`
	}

	ppdsAutocomplete := new([]AutoCompletePPDS)
	db := database.DB

	db.Model(&entity.PPDS{}).Find(&ppdsAutocomplete)

	resp.Status = static.StatusSuccess
	resp.Message = "Return PPDS Autocomplete"
	resp.Data = ppdsAutocomplete

	return c.Status(fiber.StatusOK).JSON(resp)
}
