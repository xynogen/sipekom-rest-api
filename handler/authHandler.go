package handler

import (
	"errors"
	"sipekom-rest-api/config"
	"sipekom-rest-api/database"
	"sipekom-rest-api/model"
	"sipekom-rest-api/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type SendUserData struct {
		Username string `json:"username"`
		Level    uint8  `json:"level"`
		ExpireAt int64  `json:"expireAt"`
	}

	input := new(LoginInput)
	userData := new(model.User)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	user, err := getUserByUsername(input.Username)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}

	if user != nil {
		userData.Username = user.Username
		userData.Password = user.Password
		userData.Level = user.Level
	}

	if !utils.IsPasswordValid(input.Password, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	// create claims
	expire := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"username": userData.Username,
		"level":    userData.Level,
		"exp":      expire,
	}

	// set signing method and create token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Env(config.SECRET)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sendUserData := SendUserData{
		Username: userData.Username,
		Level:    userData.Level,
		ExpireAt: expire,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success Login", "token": token, "data": sendUserData})
}

func getUserByUsername(username string) (*model.User, error) {
	db := database.DB
	var user model.User

	if err := db.Where("username = ?", username).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
