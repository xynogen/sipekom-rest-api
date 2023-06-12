package handler

import (
	"errors"
	"time"

	"sipekom-rest-api/config"
	"sipekom-rest-api/database"
	"sipekom-rest-api/model/entity"
	"sipekom-rest-api/model/request"
	"sipekom-rest-api/model/response"
	"sipekom-rest-api/model/static"
	"sipekom-rest-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// auth godoc
// @Summary authorization.
// @Description login.
// @Tags Authorization
// @param body body request.LoginRequest true "body"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/login [post]
func Login(c *fiber.Ctx) error {
	input := new(request.LoginRequest)
	resp := new(response.Response)
	respToken := new(response.TokenResponse)

	if err := c.BodyParser(&input); err != nil {
		resp.Status = static.StatusError
		resp.Message = "Error on login request"
		resp.Data = nil
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	user, err := GetUserByUsername(input.Username)

	if err != nil {
		resp.Status = static.StatusError
		resp.Message = "Error on Input Data"
		resp.Data = nil
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	if user.IsActivated != static.Activated {
		resp.Status = static.StatusError
		resp.Message = "User is Disabled"
		resp.Data = nil
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	if !utils.IsPasswordValid(input.Password, user.Password) {
		resp.Status = static.StatusError
		resp.Message = "Invalid password"
		resp.Data = nil
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	// create claims
	expire := time.Now().Add(time.Hour * 24).Unix()

	claims := response.Claims{
		Username: user.Username,
		Level:    user.Level,
		Exp:      expire,
	}

	// set signing method and create token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Env(config.SECRET)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sendUserData := response.LoginResponseData{
		Username: user.Username,
		Level:    user.Level,
		ExpireAt: expire,
	}

	respToken.Status = static.StatusSuccess
	respToken.Message = "Login Success"
	respToken.Data = sendUserData
	respToken.Token = token

	return c.Status(fiber.StatusOK).JSON(respToken)
}

func GetUserByUsername(username string) (*entity.User, error) {
	db := database.DB
	user := new(entity.User)

	if err := db.Where("username = ?", username).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return user, nil
}
