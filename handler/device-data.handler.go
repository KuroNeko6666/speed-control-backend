package handler

import (
	"net/http"

	"github.com/KuroNeko6666/speed-control-backend.git/database"
	"github.com/KuroNeko6666/speed-control-backend.git/database/model"
	"github.com/KuroNeko6666/speed-control-backend.git/model/form"
	"github.com/KuroNeko6666/speed-control-backend.git/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func ReadDeviceData(ctx *fiber.Ctx) error {
	var data []model.DeviceData
	var db *gorm.DB
	var totalPage int64 = 0
	var currentPage int64 = 0

	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)
	// search := ctx.Query("search", "")
	deviceID := ctx.Query("device_id", "")

	if len(deviceID) == 0 {
		var total int64
		// err = database.DB.Model(&model.Device{ID: deviceID}).Limit(limit).Offset(offset).Association("Data").Find(&data)
		db = database.DB.Model(&model.DeviceData{}).Find(&data)
		database.DB.Model(&model.DeviceData{}).Count(&total)
		totalPage = total / int64(limit)
		currentPage = (int64(offset) + int64(limit)) / int64(limit)

	} else {
		var total int64
		// err = database.DB.Model(&model.Device{ID: deviceID}).Limit(limit).Offset(offset).Where("device_id = ?", deviceID).Association("Data").Find(&data)
		db = database.DB.Model(&model.DeviceData{}).Preload("Device").Where("device_id = ?", deviceID).Find(&data)
		database.DB.Model(&model.DeviceData{}).Where("device_id = ?", deviceID).Count(&total)
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

func CreateDeviceData(ctx *fiber.Ctx) error {
	var form form.DeviceData
	var data model.DeviceData

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
