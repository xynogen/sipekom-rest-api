package handler

import (
	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"

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

	type DataLokasi struct {
		ID     uint   `json:"id_lokasi"`
		Lokasi string `json:"lokasi"`
		URI    string `json:"uri"`
	}

	dataLokasi := new([]DataLokasi)
	db.Model(&entity.Lokasi{}).Scopes(utils.Paginate(c)).Find(dataLokasi)

	resp.Status = static.StatusSuccess
	resp.Message = "Return All Lokasi"
	resp.Data = dataLokasi
	return c.Status(fiber.StatusOK).JSON(resp)
}
