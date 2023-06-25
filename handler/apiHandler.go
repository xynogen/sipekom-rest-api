package handler

import (
	"strconv"

	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"

	"github.com/gofiber/fiber/v2"
	qrcode "github.com/skip2/go-qrcode"
)

// Hello godoc
// @Summary server status.
// @Description get the status of server.
// @Tags Misc
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

// Check godoc
// @Security ApiKeyAuth
// @Summary encpoint to check token validation.
// @Description get validation of the token.
// @Tags Misc
// @Accept */*
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/check [get]
func Check(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusSuccess
	resp.Message = "Your Token is Valid"
	resp.Data = nil

	return c.Status(fiber.StatusOK).JSON(resp)
}

// QR godoc
// @Summary qr code image.
// @Description get qr codes based on id_lokasi.
// @Tags Misc
// @Accept */*
// @Success 200
// @Router /api/qr/get/{id_lokasi} [get]
func GetQR(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	//if not admin return unauthorize user
	userToken := utils.GetJWTFromHeader(c)
	userClaims := utils.DecodeJWT(userToken)
	if userClaims.Role != static.RoleAdmin {
		resp.Message = "Unauthorized user"
		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	id_lokasi, err := strconv.Atoi(c.AllParams()["id_lokasi"])
	if err != nil || id_lokasi < 1 {
		resp.Message = "id_lokasi is Not Valid"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	lokasi := new(entity.Lokasi)
	db := database.DB
	if db.Where("id = ?", id_lokasi).First(&lokasi).RowsAffected != 1 {
		resp.Message = "Lokasi not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	qr_text := utils.Encode64(lokasi.Uri)

	png, err := qrcode.Encode(qr_text, qrcode.Medium, 256)
	if err != nil {
		resp.Message = "Error on Encoding to QR Code"
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	c.Set("content-type", "image/png")
	return c.Status(fiber.StatusOK).Send(png)
}
