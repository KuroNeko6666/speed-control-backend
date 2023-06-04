package handler

import (
	"log"
	"math"
	"net/http"

	"github.com/KuroNeko6666/speed-control-backend.git/database"
	"github.com/KuroNeko6666/speed-control-backend.git/database/model"
	"github.com/KuroNeko6666/speed-control-backend.git/model/form"
	"github.com/KuroNeko6666/speed-control-backend.git/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func ReadUsers(ctx *fiber.Ctx) error {
	var data []model.User
	var db *gorm.DB
	var totalPage int64 = 0
	var currentPage int64 = 0

	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)
	search := ctx.Query("search", "")
	role := ctx.Query("role", "")

	if len(role) == 0 || role == "all" {
		var total int64
		var page float64
		db = database.DB.Model(&model.User{}).Limit(limit).Offset(offset).Where("username LIKE ?", "%"+search+"%").Or("name LIKE ?", "%"+search+"%").Find(&data)
		database.DB.Model(&model.User{}).Where("username LIKE ?", "%"+search+"%").Or("name LIKE ?", "%"+search+"%").Count(&total)

		page = float64(total) / float64(limit)
		totalPage = int64(math.Ceil(page))
		currentPage = (int64(offset) + int64(limit)) / int64(limit)

	} else {

		log.Println("here")
		var total int64
		var page float64
		db = database.DB.Model(&model.User{}).Limit(limit).Offset(offset).Where(database.DB.Where("username LIKE ?", "%"+search+"%").Or("name LIKE ?", "%"+search+"%")).Where("role = ?", role).Find(&data)
		database.DB.Model(&model.User{}).Where(database.DB.Where("username LIKE ?", "%"+search+"%").Or("name LIKE ?", "%"+search+"%")).Where("role = ?", role).Count(&total)
		page = float64(total) / float64(limit)
		totalPage = int64(math.Ceil(page))
		currentPage = (int64(offset) + int64(limit)) / int64(limit)
	}

	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	return Success(ctx, response.Paginate{
		Page:      currentPage,
		TotalPage: totalPage,
		Data:      data,
	})

}

func FindUser(ctx *fiber.Ctx) error {
	var data model.User
	id := ctx.Query("id", "")

	data.ID = id
	db := database.DB.Model(&data).Find(&data)
	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return NotFound(ctx)
	}

	return Success(ctx, data)
}

func UpdateUser(ctx *fiber.Ctx) error {
	var form form.UpdateUser
	var data model.User
	id := ctx.Query("id", "")

	if err := ctx.BodyParser(&form); err != nil {
		return BadRequest(ctx, err.Error())
	}

	data.ID = id
	db := database.DB.Model(&data).Find(&data)

	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return NotFound(ctx)
	}

	if err := copier.CopyWithOption(&data, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return BadRequest(ctx, err.Error())
	}

	db = database.DB.Model(&data).Updates(&data)

	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return InternalServer(ctx, http.StatusText(http.StatusInternalServerError))
	}

	return SuccessString(ctx)
}

func DeleteUser(ctx *fiber.Ctx) error {
	var data model.User
	id := ctx.Query("id", "")

	data.ID = id
	db := database.DB.Model(&data).Preload("Devices").Find(&data)

	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return NotFound(ctx)
	}

	if err := database.DB.Model(&data).Association("Devices").Clear(); err != nil {
		return InternalServer(ctx, err.Error())
	}

	db = database.DB.Model(&data).Delete(&data)

	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return InternalServer(ctx, http.StatusText(http.StatusInternalServerError))
	}

	return SuccessString(ctx)

}
