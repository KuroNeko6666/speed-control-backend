package handler

import (
	"net/http"

	"github.com/KuroNeko6666/speed-control-backend.git/config"
	"github.com/KuroNeko6666/speed-control-backend.git/database"
	"github.com/KuroNeko6666/speed-control-backend.git/database/model"
	"github.com/KuroNeko6666/speed-control-backend.git/helper"
	"github.com/KuroNeko6666/speed-control-backend.git/model/form"
	"github.com/KuroNeko6666/speed-control-backend.git/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func Login(ctx *fiber.Ctx) error {

	var form form.Login
	var data model.User

	if err := ctx.BodyParser(&form); err != nil {
		return BadRequest(ctx, err.Error())
	}

	db := database.DB.Model(&data).Where("username = ?", form.Username).Or("email = ?", form.Username).Find(&data)

	if db.Error != nil {
		return InternalServer(ctx, db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return UnAuthorized(ctx)
	}

	if err := helper.CompareHash(data.Password, form.Password); err != nil {
		return UnAuthorized(ctx)
	}

	token, err := helper.GenerateJWT(data, config.SecretKeyApp)

	if err != nil {
		return UnAuthorized(ctx)
	}

	return Success(ctx, response.Login{Token: token, User: data})

}

func Register(ctx *fiber.Ctx) error {
	var form form.Register
	var data model.User

	if err := ctx.BodyParser(&form); err != nil {
		return BadRequest(ctx, err.Error())
	}

	if err := copier.CopyWithOption(&data, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return InternalServer(ctx, err.Error())
	}

	if err := helper.GenerateHash(&data.Password); err != nil {
		return InternalServer(ctx, err.Error())
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

func Logout(ctx *fiber.Ctx) error {
	if err := helper.DeleteSession(ctx); err != nil {
		return InternalServer(ctx, err.Error())
	}
	return SuccessString(ctx)
}
