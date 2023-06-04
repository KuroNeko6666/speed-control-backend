package handler

import (
	"log"
	"net/http"

	"github.com/KuroNeko6666/speed-control-backend.git/database"
	"github.com/KuroNeko6666/speed-control-backend.git/database/model"
	"github.com/KuroNeko6666/speed-control-backend.git/model/form"
	"github.com/KuroNeko6666/speed-control-backend.git/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func ReadDevices(ctx *fiber.Ctx) error {
	var data []model.Device
	var db *gorm.DB
	var totalPage int64 = 0
	var currentPage int64 = 0

	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)
	search := ctx.Query("search", "")
	userID := ctx.Query("user_id", "")

	if len(userID) == 0 {
		var total int64
		db = database.DB.Model(&model.Device{}).Preload("User").Limit(limit).Offset(offset).Where("model LIKE ?", "%"+search+"%").Find(&data)
		database.DB.Model(&model.Device{}).Where("model LIKE ?", "%"+search+"%").Count(&total)
		totalPage = total / int64(limit)
		currentPage = (int64(offset) + int64(limit)) / int64(limit)

	} else {
		var total int64
		db = database.DB.Model(&model.Device{}).Preload("User").Limit(limit).Offset(offset).Where("model LIKE ?", "%"+search+"%").Where("user_id = ?", userID).Find(&data)
		database.DB.Model(&model.Device{}).Where("model LIKE ?", "%"+search+"%").Where("user_id = ?", userID).Count(&total)
		totalPage = total / int64(limit)
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

func FindDevice(ctx *fiber.Ctx) error {
	var data model.Device
	id := ctx.Query("id", "")

	data.ID = id
	db := database.DB.Model(&model.Device{}).Preload("User").Find(&data)
	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return NotFound(ctx)
	}

	return Success(ctx, data)
}

func CreateDevice(ctx *fiber.Ctx) error {
	var form form.Device
	var data model.Device

	if err := ctx.BodyParser(&form); err != nil {
		return BadRequest(ctx, err.Error())
	}

	if err := copier.CopyWithOption(&data, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return BadRequest(ctx, err.Error())
	}

	db := database.DB.Model(&data).Create(&data)

	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return InternalServer(ctx, http.StatusText(http.StatusInternalServerError))
	}

	return SuccessString(ctx)
}

func UpdateDevice(ctx *fiber.Ctx) error {
	var form form.DeviceCreate
	var data model.Device
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
	log.Println("here")

	return SuccessString(ctx)
}

func DeleteDevice(ctx *fiber.Ctx) error {
	var data model.Device
	id := ctx.Query("id", "")

	data.ID = id
	db := database.DB.Model(&data).Preload("Data").Find(&data)

	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return NotFound(ctx)
	}

	if err := database.DB.Model(&data).Association("Data").Clear(); err != nil {
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
