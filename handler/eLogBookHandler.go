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

// ELogBook godoc
// @Security ApiKeyAuth
// @Summary get all ELogBook [mahasiswa 🧱, guest🔒].
// @Description get all ELogBook
// @Tags ELogBook
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/elogbook/ [get]
func GetAllELogBook(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	eLogBooks := new([]entity.ELogBook)
	db := database.DB

	// If Role Mahasiswa then return according to user
	userClaims := utils.DecodeJWT(c)
	if userClaims.Role == static.RoleMahasiswa {
		if db.Scopes(utils.Paginate(c)).Where("id_user = ?", userClaims.IDUser).Find(&eLogBooks).RowsAffected < 1 {
			resp.Status = static.StatusSuccess
			resp.Message = "ID does not have any ELogBook yet."
			resp.Data = nil
			return c.Status(fiber.StatusOK).JSON(resp)
		}
		resp.Status = static.StatusSuccess
		resp.Message = "Return All ELogBook From ID"
		resp.Data = eLogBooks
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	db.Scopes(utils.Paginate(c)).Find(&eLogBooks)
	resp.Status = static.StatusSuccess
	resp.Message = "Return All ELogBook"
	resp.Data = eLogBooks

	return c.Status(fiber.StatusOK).JSON(resp)
}

// ELogBook godoc
// @Security ApiKeyAuth
// @Summary get ELogBook [mahasiswa 🧱, guest🔒].
// @Description get ELogBook by id user.
// @Tags ELogBook
// @Produce json
// @Success 200 {object} response.Response
// @Param id_elogbook path int64 true "ID Elogbook"
// @Router /api/elogbook/get/{id_elogbook} [get]
func GetELogBook(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	db := database.DB
	eLogBook := new(entity.ELogBook)

	if db.Where("id = ?", id).Find(&eLogBook).RowsAffected < 1 {
		resp.Message = "ELogBook not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	//if mahasiswa return data according to user
	userClaims := utils.DecodeJWT(c)
	if userClaims.Role == static.RoleMahasiswa {
		if userClaims.IDUser != eLogBook.IDUser {
			resp.Message = "Unauthorized user"
			return c.Status(fiber.StatusForbidden).JSON(resp)
		}
	}

	resp.Status = static.StatusSuccess
	resp.Message = "ELogBook is Found"
	resp.Data = eLogBook
	return c.Status(fiber.StatusOK).JSON(resp)
}

// ELogBook godoc
// @Security ApiKeyAuth
// @Summary create ELogBook [konsulen🔒, mahasiswa 🧱, guest🔒].
// @Description create new ELogBook.
// @Tags ELogBook
// @Accept json
// @Produce json
// @param body body request.CreateELogBookRequest true "body"
// @Success 200 {object} response.Response
// @Router /api/elogbook/create [post]
func CreateELogBook(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	userClaims := utils.DecodeJWT(c)
	// konsulen should not created elogbook
	if userClaims.Role == static.RoleKonsulen {
		resp.Message = "Konsulen Should not Create ELogBook"
		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	newELogBook := new(request.CreateELogBookRequest)
	if err := c.BodyParser(&newELogBook); err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	db := database.DB
	newELogBookModel := new(entity.ELogBook)
	newELogBookModel.IDUser = userClaims.IDUser
	newELogBookModel.IDKonsulen = newELogBook.IDkonsulen
	newELogBookModel.Title = newELogBook.Title
	newELogBookModel.Jumlah = newELogBook.Jumlah
	newELogBookModel.StartTime = utils.ParseUnitTimeInt(newELogBook.StartTime)
	newELogBookModel.EndTime = utils.ParseUnitTimeInt(newELogBook.EndTime)
	newELogBookModel.Deskripsi = newELogBook.Deskripsi
	newELogBookModel.MedicalRecord = newELogBook.Medical_Record
	newELogBookModel.IsAccepted = static.AccOnReview

	if err := db.Create(&newELogBookModel).Error; err != nil {
		resp.Message = "Invalid Data"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "ELogBook successfully Created"
	resp.Data = newELogBookModel
	return c.Status(fiber.StatusOK).JSON(resp)
}

// ELogBook godoc
// @Security ApiKeyAuth
// @Summary delete ELogBook [konsulen🔒, mahasiswa 🧱, guest🔒].
// @Description delete ELogBook by ID.
// @Tags ELogBook
// @Accept json
// @Produce json
// @Param id_elogbook path int64 true "ID ELogBook"
// @Success 200 {object} response.Response
// @Router /api/elogbook/delete/{id_elogbook} [delete]
func DeleteELogBook(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id_elogbook, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil || id_elogbook < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	eLogBook := new(entity.ELogBook)
	db := database.DB

	if db.Where("id = ?", id_elogbook).Find(&eLogBook).RowsAffected < 1 {
		resp.Message = "ELogBook not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	userClaims := utils.DecodeJWT(c)
	//if mahasiswa process data according to user
	if userClaims.Role == static.RoleMahasiswa {
		if userClaims.IDUser != eLogBook.IDUser {
			resp.Message = "Unauthorized user"
			return c.Status(fiber.StatusForbidden).JSON(resp)
		}
	}

	// if konsulen process data according to konsulen
	if userClaims.Role == static.RoleKonsulen {
		resp.Message = "Konsulen Should not Delete ELogBook"
		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	db.Where("id = ?", id_elogbook).Delete(&eLogBook)
	resp.Status = static.StatusSuccess
	resp.Message = "E-Log Book has been Delete"
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}

// ELogBook godoc
// @Security ApiKeyAuth
// @Summary update ELogBook [konsulen🔒, mahasiswa 🧱, guest🔒].
// @Description update ELogBook by ID.
// @Tags ELogBook
// @Produce json
// @Success 200 {object} response.Response
// @param body body request.UpdateELogBookRequest true "body"
// @Param id_elogbook path int64 true "ID ELogBook"
// @Router /api/elogbook/update/{id_elogbook} [put]
func UpdateElogBook(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil

	id_elogbook, err := strconv.Atoi(c.AllParams()["id_elogbook"])
	if err != nil || id_elogbook < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	db := database.DB
	eLogBook := new(entity.ELogBook)

	if db.Where("id = ?", id_elogbook).First(&eLogBook).RowsAffected < 1 {
		resp.Message = "E-Log Book not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	userClaims := utils.DecodeJWT(c)
	//if mahasiswa try to update other user elogbook
	if userClaims.Role == static.RoleMahasiswa {
		if userClaims.IDUser != eLogBook.IDUser {
			resp.Message = "Unauthorized user"
			return c.Status(fiber.StatusForbidden).JSON(resp)
		}
	}

	// konsulen should not directly update elogbook
	if userClaims.Role == static.RoleKonsulen {
		resp.Message = "Konsulen Should not Update ElogBook directly"
		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	updateELogBook := new(request.UpdateELogBookRequest)
	if err := c.BodyParser(&updateELogBook); err != nil {
		resp.Message = "Review your input"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	eLogBook.Title = updateELogBook.Title
	eLogBook.Jumlah = updateELogBook.Jumlah
	eLogBook.StartTime = utils.ParseUnitTimeInt(updateELogBook.StartTime)
	eLogBook.EndTime = utils.ParseUnitTimeInt(updateELogBook.EndTime)
	eLogBook.Deskripsi = updateELogBook.Deskripsi
	eLogBook.MedicalRecord = updateELogBook.Medical_Record

	if err := db.Save(&eLogBook).Error; err != nil {
		resp.Message = "Duplicate Data Found"
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "ELogBook successfully Updated"
	resp.Data = eLogBook
	return c.Status(fiber.StatusOK).JSON(resp)
}

// ELogBook godoc
// @Security ApiKeyAuth
// @Summary Approved ELogBook [mahasiswa🔒, guest🔒].
// @Description Approved ELogBook by ID.
// @Tags ELogBook
// @Produce json
// @Success 200 {object} response.Response
// @Param id_elogbook path int64 true "ID ELogBook"
// @Router /api/elogbook/accepted/{id_elogbook} [put]
func AcceptedElogBook(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil
	userClaims := utils.DecodeJWT(c)

	if userClaims.Role == static.RoleMahasiswa {
		resp.Message = "Unauthorized user"
		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	id_elogbook, err := strconv.Atoi(c.AllParams()["id_elogbook"])
	if err != nil || id_elogbook < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	db := database.DB
	eLogBook := new(entity.ELogBook)

	if db.Where("id = ?", id_elogbook).First(&eLogBook).RowsAffected < 1 {
		resp.Message = "E-Log Book not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	if eLogBook.IsAccepted != static.AccOnReview {
		resp.Message = "ElogBook already Reviewed"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	// change to approve
	eLogBook.IsAccepted = static.AccApproved

	if err := db.Save(&eLogBook).Error; err != nil {
		resp.Message = "Duplicate Data Found"
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "ELogBook successfully Updated"
	resp.Data = eLogBook
	return c.Status(fiber.StatusOK).JSON(resp)
}

// ELogBook godoc
// @Security ApiKeyAuth
// @Summary Approved ELogBook [mahasiswa🔒, guest🔒].
// @Description Approved ELogBook by ID.
// @Tags ELogBook
// @Produce json
// @Success 200 {object} response.Response
// @Param id_elogbook path int64 true "ID ELogBook"
// @Router /api/elogbook/rejected/{id_elogbook} [put]
func RejectedElogBook(c *fiber.Ctx) error {
	resp := new(response.Response)
	resp.Status = static.StatusError
	resp.Data = nil
	userClaims := utils.DecodeJWT(c)

	if userClaims.Role == static.RoleMahasiswa {
		resp.Message = "Unauthorized user"
		return c.Status(fiber.StatusForbidden).JSON(resp)
	}

	id_elogbook, err := strconv.Atoi(c.AllParams()["id_elogbook"])
	if err != nil || id_elogbook < 1 {
		resp.Message = "ID is Not Valid"
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	db := database.DB
	eLogBook := new(entity.ELogBook)

	if db.Where("id = ?", id_elogbook).First(&eLogBook).RowsAffected < 1 {
		resp.Message = "E-Log Book not Found"
		return c.Status(fiber.StatusNotFound).JSON(resp)
	}

	if eLogBook.IsAccepted != static.AccOnReview {
		resp.Message = "ElogBook already Reviewed"
		return c.Status(fiber.StatusOK).JSON(resp)
	}

	// change to not approve
	eLogBook.IsAccepted = static.AccNotApproved

	if err := db.Save(&eLogBook).Error; err != nil {
		resp.Message = "Duplicate Data Found"
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}

	resp.Status = static.StatusSuccess
	resp.Message = "ELogBook successfully Updated"
	resp.Data = eLogBook
	return c.Status(fiber.StatusOK).JSON(resp)
}
