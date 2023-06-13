package handler

import (
	"strconv"

	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/request"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get all ELogBook.
// @Description get all ELogBook
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
// @Summary get ELogBook.
// @Description get ELogBook by id user.
// @Tags ELogBook
// @Produce json
// @Success 200 {object} response.Response
// @Param id path int64 true "ID User"
// @Router /api/elogbook/get/{id_user} [get]
func GetELogBook(c *fiber.Ctx) error {
	eLogBook := new(entity.ELogBook)
	user := new(entity.User)
	resp := new(response.Response)

	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		resp.Message = "User not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	if user.Level != static.LevelMahasiswa {
		resp.Message = "User not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	if err := db.Where("id_user = ?", id).Scopes(utils.Paginate(c)).Find(&eLogBook).Error; err != nil {
		resp.Message = "E-Log Book not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "E-Log Book is Found"
	resp.Data = eLogBook
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary create ELogBook.
// @Description create new ELogBook.
// @Tags ELogBook
// @Accept json
// @Produce json
// @param body body request.CreateELogBookRequest true "body"
// @Success 200 {object} response.Response
// @Router /api/elogbook/create [post]
func CreateELogBook(c *fiber.Ctx) error {
	newELogBook := new(request.CreateELogBookRequest)
	resp := new(response.Response)

	resp.Status = static.StatusError
	resp.Data = nil

	if err := c.BodyParser(&newELogBook); err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	jwtTokenStr := utils.GetJWTFromHeader(c)
	claims := utils.DecodeJWT(jwtTokenStr)

	user, err := GetUserByUsername(claims.Username)
	if err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB

	newELogBookModel := new(entity.ELogBook)
	newELogBookModel.IDUser = user.ID
	newELogBookModel.Title = newELogBook.Title
	newELogBookModel.Jumlah = newELogBook.Jumlah
	newELogBookModel.StartTime = utils.ParseUnitTimeInt(newELogBook.StartTime)
	newELogBookModel.EndTime = utils.ParseUnitTimeInt(newELogBook.EndTime)
	newELogBookModel.Deskripsi = newELogBook.Deskripsi
	newELogBookModel.Medical_Record = newELogBook.Medical_Record

	if err := db.Create(&newELogBookModel).Error; err != nil {
		resp.Message = "Invalid Data"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "User successfully Created"
	resp.Data = newELogBookModel
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary delete ELogBook.
// @Description delet ELogBook by ID.
// @Tags ELogBook
// @Accept json
// @Produce json
// @Param id path int64 true "ELogBook ID"
// @Success 200 {object} response.Response
// @Router /api/elogbook/delete/{id_elogbook} [delete]
func DeleteELogBook(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	eLogBook := new(entity.ELogBook)
	db := database.DB
	if err := db.Where("id = ?", id).Delete(&eLogBook).Error; err != nil {
		resp.Message = "E-Log Book not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "E-Log Book has been Delete"
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary update ELogBook.
// @Description update user by id.
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @param body body request.UpdateUserRequest true "body"
// @Param id path int64 true "User ID"
// @Router /api/user/update/{id_user} [put]
