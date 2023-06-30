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
// @Tags API
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

// Hello godoc
// @Summary get user data with limited access.
// @Description get data of user.
// @Tags API
// @Accept */*
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/data/{search_query} [get]
func GetData(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusSuccess
	resp.Data = nil

	search_query := c.AllParams()["search_query"]
	if len(search_query) > 255 {
		resp.Message = "Query is too long"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	db := database.DB
	ppds := new(entity.PPDS)

	search_query = "%" + search_query + "%"

	if db.Where("name LIKE ?", search_query).First(&ppds).RowsAffected < 1 {
		resp.Message = "User is Not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	user := new(entity.User)
	if db.Where("id = ?", ppds.IDUser).First(user).RowsAffected < 1 {
		resp.Message = "User is Not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	publicResponseData := new(response.PublicResponseData)
	publicResponseData.Name = ppds.Name
	publicResponseData.Kompetensi = ppds.Kompetensi
	publicResponseData.Foto = user.Photo

	resp.Status = static.StatusSuccess
	resp.Message = "User is found"
	resp.Data = publicResponseData
	return c.Status(fiber.StatusOK).JSON(resp)
}

// Check godoc
// @Security ApiKeyAuth
// @Summary check token validation [guest🔒].
// @Description get validation of the token.
// @Tags API
// @Accept */*
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/whoami [get]
func Whoami(c *fiber.Ctx) error {
	resp := new(response.Response)
	userClaims := utils.DecodeJWT(c)

	// make it the same as login response
	user := new(response.LoginResponseData)
	user.IDUser = userClaims.IDUser
	user.Username = userClaims.Username
	user.Role = userClaims.Role
	user.ExpireAt = userClaims.Exp

	resp.Status = static.StatusSuccess
	resp.Message = "Your Token is Valid"
	resp.Data = user

	return c.Status(fiber.StatusOK).JSON(resp)
}

// QR godoc
// @Security ApiKeyAuth
// @Summary qr code image [mahasiswa🔒, konsulen🔒, guest🔒].
// @Description get qr codes based on id_lokasi.
// @Tags API
// @Accept */*
// @Param id_lokasi path int64 true "ID Lokasi"
// @Success 200
// @Router /api/qr/get/{id_lokasi} [get]
func GetQR(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	if !utils.IsAdmin(c) {
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
