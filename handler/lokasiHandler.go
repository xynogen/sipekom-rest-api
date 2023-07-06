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

// Absen godoc
// @Security ApiKeyAuth
// @Summary get all Lokasi [guest🔒].
// @Description get all Lokasi
// @Tags Lokasi
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/lokasi/ [get]
func GetAllLokasi(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil
	db := database.DB

	dataLokasi := new([]response.GetLokasiResponse)
	db.Model(&entity.Lokasi{}).Scopes(utils.Paginate(c)).Find(dataLokasi)

	resp.Status = static.StatusSuccess
	resp.Message = "Return All Lokasi"
	resp.Data = dataLokasi
	return c.Status(fiber.StatusOK).JSON(resp)
}

// Absen godoc
// @Security ApiKeyAuth
// @Summary get Lokasi [guest🔒].
// @Description get Lokasi by ID_lokasi
// @Tags Lokasi
// @Produce json
// @Success 200 {object} response.Response
// @Param id_user path int64 true "User ID"
// @Router /api/lokasi/ [get]
func GetLokasi(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id_lokasi, err := strconv.Atoi(c.AllParams()["id_lokasi"])
	if err != nil || id_lokasi < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	dataLokasi := new(response.GetLokasiResponse)
	db := database.DB

	if db.Model(&entity.Lokasi{}).Where("id = ?", id_lokasi).First(&dataLokasi).RowsAffected < 1 {
		resp.Message = "Lokasi not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Lokasi is Found"
	resp.Data = dataLokasi
	return c.Status(fiber.StatusOK).JSON(resp)
}
