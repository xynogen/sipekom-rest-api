package handler

import (
	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/request"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"

	"errors"
	"strconv"
	"time"

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

	db.Scopes(utils.Paginate(c)).Find(&absens)
	resp.Status = static.StatusSuccess
	resp.Message = "Return All Absen"
	resp.Data = absens

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get absen.
// @Description get absen by id user.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param id path int64 true "ID User"
// @Router /api/absen/get/{id} [get]
func GetAbsen(c *fiber.Ctx) error {
	absen := new(entity.Absensi)
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

	if err := db.Where("id_user = ?", id).Scopes(utils.Paginate(c)).Find(&absen).Error; err != nil {
		resp.Status = static.StatusError
		resp.Message = "Absen not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Absen is Found"
	resp.Data = absen
	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary create absen.
// @Description get absen by location.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param location path string true "location"
// @Router /api/absen/create/{location} [get]
func CreateAbsen(c *fiber.Ctx) error {
	jwtTokenStr := utils.GetJWTFromHeader(c)
	claims := utils.DecodeJWT(jwtTokenStr)
	resp := new(response.Response)
	absenOld := new(entity.Absensi)
	absenNew := new(entity.Absensi)
	db := database.DB
	absenFlag := static.FlagAbsenCheckIn
	uri := c.AllParams()["lokasi"]

	user, err := GetUserByUsername(claims.Username)
	if err != nil {
		resp.Status = static.StatusError
		resp.Message = "Input Invalid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	lokasi, err := GetLokasiFromUri(uri)
	if err != nil {
		resp.Status = static.StatusError
		resp.Message = "Input Invalid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	// check latest absen
	if db.Where("id_user = ?", user.ID).Order("absen DESC").First(&absenOld).RowsAffected != 0 {
		// check if we checkin in new location
		if absenOld.AbsenFlag == static.FlagAbsenCheckIn {
			if absenOld.Lokasi != lokasi.Lokasi {
				resp.Status = static.StatusError
				resp.Message = "Invalid Checkin"
				resp.Data = nil
				return c.Status(fiber.StatusOK).JSON(resp)
			}
		}

	}

	// check if we checkout
	if absenOld.AbsenFlag == static.FlagAbsenCheckIn {
		absenFlag = static.FlagAbsenCheckOut
	}

	// check if data not exist before
	if absenOld.Lokasi == "" {
		absenFlag = static.FlagAbsenCheckIn
	}

	absenNew.Absen = time.Now()
	absenNew.AbsenFlag = absenFlag
	absenNew.Lokasi = lokasi.Lokasi
	absenNew.IDUser = user.ID

	if err := db.Create(&absenNew).Error; err != nil {
		resp.Status = static.StatusError
		resp.Message = "Invalid Data"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Absen successfully Created"
	resp.Data = nil

	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary update absen.
// @Description update absen by id.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @param body body request.UpdateAbsenRequest true "body"
// @Param id path int64 true "Absen ID"
// @Router /api/absen/update/{id} [put]
func UpdateAbsen(c *fiber.Ctx) error {
	updateAbsen := new(request.UpdateAbsenRequest)
	resp := new(response.Response)

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Status = static.StatusError
		resp.Message = "ID is Not Valid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	if err := c.BodyParser(&updateAbsen); err != nil {
		resp.Status = static.StatusError
		resp.Message = "Review your input"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	absen := new(entity.Absensi)

	if err := db.First(&absen, id).Error; err != nil {
		resp.Status = static.StatusError
		resp.Message = "User not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	absen.Absen = updateAbsen.Absen
	absen.AbsenFlag = updateAbsen.AbsenFlag
	absen.Lokasi = updateAbsen.Lokasi
	absen.IDUser = updateAbsen.IDUser

	if err := db.Save(&absen).Error; err != nil {
		resp.Status = static.StatusError
		resp.Message = "Duplicate Data Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "User successfully Updated"
	resp.Data = absen
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary delete absen.
// @Description delete absen by id.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param id path int64 true "Absen ID"
// @Router /api/absen/delete/{id} [delete]
func DeleteAbsen(c *fiber.Ctx) error {
	absen := new(entity.Absensi)
	resp := new(response.Response)

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Status = static.StatusError
		resp.Message = "ID is Not Valid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	if err := db.Where("id = ?", id).Delete(&absen).Error; err != nil {
		resp.Status = static.StatusError
		resp.Message = "Absen not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Absen has been Delete"
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}

func GetLokasiFromUri(Uri string) (*entity.Lokasi, error) {
	lokasi := new(entity.Lokasi)
	db := database.DB

	if db.Where("uri = ?", Uri).Find(&lokasi).RowsAffected < 1 {
		return nil, errors.New("location not found")
	}
	return lokasi, nil
}
