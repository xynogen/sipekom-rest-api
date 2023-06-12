package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		size := 15
		queryPage := c.Query("page")
		page, _ := strconv.Atoi(queryPage)

		offset := (page) * size
		return db.Offset(offset).Limit(size)
	}
}
