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

// Absen godoc
// @Security ApiKeyAuth
// @Summary get all Absen [mahasiswa limit, guest🔒].
// @Description get all Absen, mahasiswa have limited access
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param id_user path int64 true "ID User"
// @Router /api/absen/{id_user} [get]
func GetAllAbsen(c *fiber.Ctx) error {
	absens := new([]entity.Absensi)
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id_user, err := strconv.Atoi(c.AllParams()["id_user"])
	if err != nil || id_user < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusNotAcceptable).JSON(resp)
	}

	db := database.DB

	jwtToken := utils.GetJWTFromHeader(c)
	userClaims := utils.DecodeJWT(jwtToken)

	// if mahasiswa return data according to user
	if userClaims.Role == static.RoleMahasiswa {
		if db.Scopes(utils.Paginate(c)).Where("id_user = ?", userClaims.IDUser).Find(absens).RowsAffected < 1 {
			resp.Status = static.StatusSuccess
			resp.Message = "ID does not have any absen yet."
			resp.Data = nil
			return c.Status(fiber.StatusOK).JSON(resp)
		}

		resp.Status = static.StatusSuccess
		resp.Message = "Return All Absen From ID"
		resp.Data = absens
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	if db.Scopes(utils.Paginate(c)).Where("id_user = ?", id_user).Find(absens).RowsAffected < 1 {
		resp.Status = static.StatusSuccess
		resp.Message = "ID does not have any absen yet."
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Return All Absen"
	resp.Data = absens
	return c.Status(fiber.StatusOK).JSON(resp)
}

// Absen godoc
// @Security ApiKeyAuth
// @Summary get Absen [mahasiswa limit, guest🔒].
// @Description get Absen by ID Absen, mahasiswa have limited access.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param id_absen path int64 true "ID Absen"
// @Router /api/absen/get/{id_absen} [get]
func GetAbsen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id_absen, err := strconv.Atoi(c.AllParams()["id_absen"])
	if err != nil || id_absen < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusNotAcceptable).JSON(resp)
	}

	db := database.DB
	absen := new(entity.Absensi)

	if db.Where("id = ?", id_absen).First(&absen).RowsAffected < 1 {
		resp.Message = "Absen not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	//if mahasiswa return data according to user
	userToken := utils.GetJWTFromHeader(c)
	userClaims := utils.DecodeJWT(userToken)
	if userClaims.Role == static.RoleMahasiswa {
		id_user := int(userClaims.IDUser)

		if id_user != int(absen.IDUser) {
			resp.Message = "Unauthorized user"
			return c.Status(fiber.StatusForbidden).JSON(resp)
		}

	}

	resp.Status = static.StatusSuccess
	resp.Message = "Absen is Found"
	resp.Data = absen
	return c.Status(fiber.StatusOK).JSON(resp)
}

// Absen godoc
// @Security ApiKeyAuth
// @Summary create Absen [guest🔒].
// @Description get Absen by location.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param location_base64 path string true "location base64"
// @Router /api/absen/create/{uri_base64} [get]
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

	uri_base64 := c.AllParams()["uri_base64"]
	uri := utils.Decode64(uri_base64)
	if uri == "" {
		resp.Message = "Input Invalid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	lokasi, err := GetLokasiFromUri(uri)
	if err != nil {
		resp.Message = "Input Invalid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	absenOld := new(entity.Absensi)
	absenFlag := static.AbsenCheckIn
	db := database.DB
	// check latest absen
	if db.Where("id_user = ?", user.ID).Order("absen DESC").First(&absenOld).RowsAffected != 0 {
		// check if we checkin in new location
		if absenOld.AbsenFlag == static.AbsenCheckIn {
			if absenOld.Lokasi != lokasi.Lokasi {
				resp.Message = "Invalid Checkin"
				return c.Status(fiber.StatusOK).JSON(resp)
			}
		}
	}

	// check if we checkout
	if absenOld.AbsenFlag == static.AbsenCheckIn {
		absenFlag = static.AbsenCheckOut
	}

	// check if data not exist before
	if absenOld.Lokasi == "" {
		absenFlag = static.AbsenCheckIn
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

// Absen godoc
// @Security ApiKeyAuth
// @Summary update Absen [konsulen🔒, mahasiswa🔒, guest🔒].
// @Description update Absen by ID, only Admin can update Absen.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @param body body request.UpdateAbsenRequest true "body"
// @Param id_absen path int64 true "ID Absen"
// @Router /api/absen/update/{id_absen} [put]
func UpdateAbsen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	if !utils.IsAdmin(c) {
		resp.Message = "Unauthorized user"
		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	id_absen, err := strconv.Atoi(c.AllParams()["id_absen"])
	if err != nil || id_absen < 1 {
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

	if db.First(&absen, id_absen).RowsAffected < 1 {
		resp.Message = "Absen not Found"
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

// Absen godoc
// @Security ApiKeyAuth
// @Summary delete Absen [konsulen🔒, mahasiswa🔒, guest🔒].
// @Description delete Absen by ID, only admin can delete absen.
// @Tags Absen
// @Produce json
// @Success 200 {object} response.Response
// @Param id_absen path int64 true "ID Absen"
// @Router /api/absen/delete/{id_absen} [delete]
func DeleteAbsen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	if !utils.IsAdmin(c) {
		resp.Message = "Unauthorized user"
		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	id_absen, err := strconv.Atoi(c.AllParams()["id_absen"])
	if err != nil || id_absen < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	absen := new(entity.Absensi)
	db := database.DB
	if db.Where("id = ?", id_absen).Delete(&absen).RowsAffected != 1 {
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
