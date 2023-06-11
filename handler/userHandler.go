package handler

import (
	"sipekom-rest-api/database"
	jsonmodel "sipekom-rest-api/json_model"
	"sipekom-rest-api/model"
	"sipekom-rest-api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Summary get all user.
// @Description get all user
// @Tags User
// @Produce json
// @Success 200 {object} []model.User
// @Router /api/user/ [get]
func GetAllUser(c *fiber.Ctx) error {
	var user []model.User
	db := database.DB

	db.Find(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Return all Users", "data": user})
}

// User godoc
// @Summary get user.
// @Description get user by id.
// @Tags User
// @Produce json
// @Success 200 {object} model.User
// @Param id path int64 true "User ID"
// @Router /api/user/get/{id} [get]
func GetUser(c *fiber.Ctx) error {
	var user model.User

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "ID is Not Valid", "data": nil})
	}

	db := database.DB
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "User not Found", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User is Found", "data": user})
}

// @User godoc
// @Summary delete user.
// @Description delete user by id.
// @Tags User
// @Produce json
// @Success 200 {object} model.User
// @Param id path int64 true "User ID"
// @Router /api/user/delete/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	var user []model.User

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "ID is Not Valid", "data": nil})
	}

	db := database.DB
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "User not Found", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User has been Delete", "data": nil})
}

// @User godoc
// @Summary update user.
// @Description update user by id.
// @Tags User
// @Produce json
// @Success 200 {object} model.User
// @Param id path int64 true "User ID"
// @Router /api/user/update/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	var updateUser jsonmodel.UpdateUserInput

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "ID is Not Valid", "data": nil})
	}

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	db := database.DB
	var user model.User

	if err := db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "User not Found", "data": nil})
	}

	user.Username = updateUser.Username
	user.Level = updateUser.Level

	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "Duplicate Data Found", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User successfully Updated", "data": user})
}

// @User godoc
// @Summary create user.
// @Description create new user.
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /api/user/create [get]
func CreateUser(c *fiber.Ctx) error {
	var newUser model.User

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	var err error
	newUser.Password, err = utils.HashPassword(newUser.Password)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "Hashing Failed", "data": err})
	}

	db := database.DB
	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "Invalid Data", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User successfully Created", "data": newUser})
}
