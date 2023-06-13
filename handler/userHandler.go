package handler

import (
	"errors"
	"strconv"

	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/request"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get all User.
// @Description get all User
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/ [get]
func GetAllUser(c *fiber.Ctx) error {
	users := new([]entity.User)
	var resp response.Response
	db := database.DB

	db.Find(&users)

	resp.Status = static.StatusSuccess
	resp.Message = "Return all Users"
	resp.Data = users

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get User.
// @Description get User by id.
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @Param id_user path int64 true "User ID"
// @Router /api/user/get/{id_user} [get]
func GetUser(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	user := new(entity.User)
	db := database.DB
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		resp.Message = "User not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "User is Found"
	resp.Data = user
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary delete User.
// @Description delete User by id.
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @Param id_user path int64 true "User ID"
// @Router /api/user/delete/{id_user} [delete]
func DeleteUser(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	user := new(entity.User)
	db := database.DB

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Message = "User not Found"
			return c.Status(fiber.StatusOK).JSON(resp)
		}
		resp.Message = "Query Error"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	// delete user account in general
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		resp.Message = "Query Error"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	// delete konsulen data if konsulen
	if user.Level == static.LevelKonsulen {
		konsulen := new(entity.Konsulen)
		if err := db.Where("id_user = ?", user.ID).Delete(&konsulen).Error; err != nil {
			resp.Message = "Query Error"
			return c.Status(fiber.StatusOK).JSON(resp)
		}
	}

	// delete mahasiwa data if mahasiswa
	if user.Level == static.LevelMahasiswa {
		mahasiwa := new(entity.PPDS)
		if err := db.Where("id_user = ?", user.ID).Delete(&mahasiwa).Error; err != nil {
			resp.Message = "Query Error"
			return c.Status(fiber.StatusOK).JSON(resp)
		}
	}

	resp.Status = static.StatusSuccess
	resp.Message = "User has been Delete"
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary update User.
// @Description update User by id.
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @param body body request.UpdateUserRequest true "body"
// @Param id_user path int64 true "User ID"
// @Router /api/user/update/{id_user} [put]
func UpdateUser(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	updateUser := new(request.UpdateUserRequest)
	if err := c.BodyParser(&updateUser); err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	user := new(entity.User)

	if err := db.First(&user, id).Error; err != nil {
		resp.Message = "User not Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	user.Username = updateUser.Username
	user.Level = updateUser.Level
	user.IsActivated = updateUser.IsActivated

	if err := db.Save(&user).Error; err != nil {
		resp.Message = "Duplicate Data Found"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "User successfully Updated"
	resp.Data = user
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary create User.
// @Description create new User.
// @Tags User
// @Accept json
// @Produce json
// @param body body request.CreateUserRequest true "body"
// @Success 200 {object} response.Response
// @Router /api/user/create [post]
func CreateUser(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	newUser := new(request.CreateUserRequest)
	if err := c.BodyParser(&newUser); err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	var err error
	newUser.Password, err = utils.HashPassword(newUser.Password)
	if err != nil {
		resp.Message = "Hashing Failed"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	newUserModel := new(entity.User)
	newUserModel.Username = newUser.Username
	newUserModel.Password = newUser.Password
	newUserModel.Level = newUser.Level
	newUserModel.IsActivated = static.Activated

	if err := db.Create(&newUserModel).Error; err != nil {
		resp.Message = "Invalid Data"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "User successfully Created"
	resp.Data = newUserModel
	return c.Status(fiber.StatusOK).JSON(resp)
}
