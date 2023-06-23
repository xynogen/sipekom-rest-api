package handler

import (
	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/request"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get all Konsulen.
// @Description get all absen
// @Tags Konsulen
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/konsulen/ [get]
func GetAllKonsulen(c *fiber.Ctx) error {
	resp := new(response.Response)

	konsulens := new([]entity.Konsulen)
	db := database.DB
	db.Scopes(utils.Paginate(c)).Find(&konsulens)

	resp.Status = static.StatusSuccess
	resp.Message = "Return All Konsulen"
	resp.Data = konsulens

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get Konsulen.
// @Description get Konsulen by id.
// @Tags Konsulen
// @Produce json
// @Success 200 {object} response.Response
// @Param id_user path int64 true "ID User"
// @Router /api/konsulen/get/{id_user} [get]
func GetKonsulen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	konsulen := new(entity.Konsulen)
	db := database.DB

	if db.Where("id_user = ?", id).Find(&konsulen).RowsAffected != 1 {
		resp.Message = "Konsulen not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Konsulen is Found"
	resp.Data = konsulen
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary create Konsulen.
// @Description create new Konsulen.
// @Tags Konsulen
// @Accept json
// @Produce json
// @param body body request.CreateKonsulenRequest true "body"
// @Success 200 {object} response.Response
// @Router /api/konsulen/create [post]
func CreateKonsulen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	newKonsulenData := new(request.CreateKonsulenRequest)
	if err := c.BodyParser(&newKonsulenData); err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	var err error
	newKonsulenData.Password, err = utils.HashPassword(newKonsulenData.Password)
	if err != nil {
		resp.Message = "Hashing Failed"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	newUserModel := new(entity.User)
	newUserModel.Username = newKonsulenData.Username
	newUserModel.Password = newKonsulenData.Password
	newUserModel.Role = static.RoleKonsulen
	newUserModel.IsActivated = static.Activated

	if err := db.Create(&newUserModel).Error; err != nil {
		resp.Message = "Username Already Exist"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	newKonsulen := new(entity.Konsulen)
	newKonsulen.Name = newKonsulenData.Name
	newKonsulen.Spesialis = newKonsulenData.Spesialis
	newKonsulen.IDUser = newUserModel.ID

	if err := db.Create(&newKonsulen).Error; err != nil {
		resp.Message = "Invalid Data"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "Konsulen successfully Created"
	resp.Data = newKonsulen
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary update Konsulen.
// @Description update Konsulen by id.
// @Tags Konsulen
// @Produce json
// @Success 200 {object} response.Response
// @param body body request.UpdateKonsulenRequest true "body"
// @Param id_user path int64 true "User ID"
// @Router /api/konsulen/update/{id_user} [put]
func UpdateKonsulen(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	updateKonsulen := new(request.UpdateKonsulenRequest)
	if err := c.BodyParser(&updateKonsulen); err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	konsulen := new(entity.Konsulen)

	if err := db.Where("id_user = ?", id).First(&konsulen).Error; err != nil {
		resp.Message = "Konsulen not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	konsulen.Name = updateKonsulen.Name
	konsulen.Spesialis = updateKonsulen.Spesialis

	if err := db.Save(&konsulen).Error; err != nil {
		resp.Message = "Duplicate Data Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "User successfully Updated"
	resp.Data = konsulen
	return c.Status(fiber.StatusOK).JSON(resp)
}
