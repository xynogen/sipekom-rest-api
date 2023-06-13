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
// @Summary get all Absen.
// @Description get all Absen
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
// @Summary get Absen.
// @Description get Absen by ID User.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param id_user path int64 true "ID User"
// @Router /api/absen/get/{id_user} [get]
func GetAbsen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	absen := new(entity.Absensi)

	if db.Where("id_user = ?", id).Scopes(utils.Paginate(c)).Find(&absen).RowsAffected < 1 {
		resp.Message = "Absen not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Absen is Found"
	resp.Data = absen
	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary create Absen.
// @Description get Absen by location.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param location path string true "location"
// @Router /api/absen/create/{location} [get]
func CreateAbsen(c *fiber.Ctx) error {
	jwtTokenStr := utils.GetJWTFromHeader(c)
	claims := utils.DecodeJWT(jwtTokenStr)

	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	user, err := GetUserByUsername(claims.Username)
	if err != nil {
		resp.Message = "Input Invalid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	uri := c.AllParams()["lokasi"]
	lokasi, err := GetLokasiFromUri(uri)
	if err != nil {
		resp.Message = "Input Invalid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	absenOld := new(entity.Absensi)
	absenFlag := static.FlagAbsenCheckIn
	db := database.DB
	// check latest absen
	if db.Where("id_user = ?", user.ID).Order("absen DESC").First(&absenOld).RowsAffected != 0 {
		// check if we checkin in new location
		if absenOld.AbsenFlag == static.FlagAbsenCheckIn {
			if absenOld.Lokasi != lokasi.Lokasi {
				resp.Message = "Invalid Checkin"
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

	absenNew := new(entity.Absensi)
	absenNew.Absen = time.Now()
	absenNew.AbsenFlag = absenFlag
	absenNew.Lokasi = lokasi.Lokasi
	absenNew.IDUser = user.ID

	if err := db.Create(&absenNew).Error; err != nil {
		resp.Message = "Invalid Data"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Absen successfully Created"
	resp.Data = absenNew

	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary update Absen.
// @Description update Absen by ID.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @param body body request.UpdateAbsenRequest true "body"
// @Param id_absen path int64 true "Absen ID"
// @Router /api/absen/update/{id_absen} [put]
func UpdateAbsen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	updateAbsen := new(request.UpdateAbsenRequest)
	if err := c.BodyParser(&updateAbsen); err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	absen := new(entity.Absensi)

	if err := db.First(&absen, id).Error; err != nil {
		resp.Message = "User not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	absen.Absen = utils.ParseUnitTimeInt(updateAbsen.Absen)
	absen.AbsenFlag = updateAbsen.AbsenFlag
	absen.Lokasi = updateAbsen.Lokasi

	if err := db.Save(&absen).Error; err != nil {
		resp.Message = "Duplicate Data Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "User successfully Updated"
	resp.Data = absen
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary delete Absen.
// @Description delete Absen by ID.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param id_absen path int64 true "Absen ID"
// @Router /api/absen/delete/{id_absen} [delete]
func DeleteAbsen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	absen := new(entity.Absensi)
	db := database.DB
	if db.Where("id = ?", id).Delete(&absen).RowsAffected != 1 {
		resp.Message = "Absen not Found"
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
