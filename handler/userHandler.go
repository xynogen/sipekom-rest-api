package handler

import (
	"strconv"

	"sipekom-rest-api/database"
	"sipekom-rest-api/model"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/request"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/utils"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Security ApiKeyAuth
// @Summary get all user.
// @Description get all user
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/user/ [get]
func GetAllUser(c *fiber.Ctx) error {
	users := new([]entity.User)
	var resp response.Response
	db := database.DB

	db.Find(&users)

	resp.Status = model.StatusSuccess
	resp.Message = "Return all Users"
	resp.Data = users

	return c.Status(fiber.StatusOK).JSON(resp)
}

// User godoc
// @Security ApiKeyAuth
// @Summary get user.
// @Description get user by id.
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @Param id path int64 true "User ID"
// @Router /api/user/get/{id} [get]
func GetUser(c *fiber.Ctx) error {
	user := new(entity.User)
	resp := new(response.Response)

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Status = model.StatusError
		resp.Message = "ID is Not Valid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		resp.Status = model.StatusError
		resp.Message = "User not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = model.StatusSuccess
	resp.Message = "User is Found"
	resp.Data = user
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary delete user.
// @Description delete user by id.
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @Param id path int64 true "User ID"
// @Router /api/user/delete/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	users := new([]entity.User)
	resp := new(response.Response)

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Status = model.StatusError
		resp.Message = "ID is Not Valid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	if err := db.Where("id = ?", id).Delete(&users).Error; err != nil {
		resp.Status = model.StatusError
		resp.Message = "User not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = model.StatusSuccess
	resp.Message = "User has been Delete"
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary update user.
// @Description update user by id.
// @Tags User
// @Produce json
// @Success 200 {object} response.Response
// @param body body request.UpdateUserRequest true "body"
// @Param id path int64 true "User ID"
// @Router /api/user/update/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	updateUser := new(request.UpdateUserRequest)
	resp := new(response.Response)

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Status = model.StatusError
		resp.Message = "ID is Not Valid"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	if err := c.BodyParser(&updateUser); err != nil {
		resp.Status = model.StatusError
		resp.Message = "Review your input"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	user := new(entity.User)

	if err := db.First(&user, id).Error; err != nil {
		resp.Status = model.StatusError
		resp.Message = "User not Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	user.Username = updateUser.Username
	user.Level = updateUser.Level
	user.IsActivated = updateUser.IsActivated

	if err := db.Save(&user).Error; err != nil {
		resp.Status = model.StatusError
		resp.Message = "Duplicate Data Found"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = model.StatusSuccess
	resp.Message = "User successfully Updated"
	resp.Data = user
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @User godoc
// @Security ApiKeyAuth
// @Summary create user.
// @Description create new user.
// @Tags User
// @Accept json
// @Produce json
// @param body body request.CreateUserRequest true "body"
// @Success 200 {object} response.Response
// @Router /api/user/create [post]
func CreateUser(c *fiber.Ctx) error {
	newUser := new(request.CreateUserRequest)
	resp := new(response.Response)

	if err := c.BodyParser(&newUser); err != nil {
		resp.Status = model.StatusError
		resp.Message = "Review your input"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	var err error
	newUser.Password, err = utils.HashPassword(newUser.Password)
	if err != nil {
		resp.Status = model.StatusError
		resp.Message = "Hashing Failed"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db := database.DB
	newUserModel := new(entity.User)
	newUserModel.Username = newUser.Username
	newUserModel.Password = newUser.Password
	newUserModel.Level = newUser.Level
	newUserModel.IsActivated = model.Activated

	if err := db.Create(&newUserModel).Error; err != nil {
		resp.Status = model.StatusError
		resp.Message = "Invalid Data"
		resp.Data = nil
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	resp.Status = model.StatusSuccess
	resp.Message = "User successfully Created"
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}
