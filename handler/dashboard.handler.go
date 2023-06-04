package handler

import (
	"log"
	"strconv"
	"time"

	"github.com/KuroNeko6666/speed-control-backend.git/database"
	"github.com/KuroNeko6666/speed-control-backend.git/database/model"
	"github.com/KuroNeko6666/speed-control-backend.git/model/response"
	"github.com/gofiber/fiber/v2"
)

func UserCount(ctx *fiber.Ctx) error {
	var count int64 = 0
	role := ctx.Query("role", "")
	if len(role) != 0 {
		log.Println(len(role))
		database.DB.Model(&model.User{}).Where("role =?", role).Count(&count)
	} else {
		database.DB.Model(&model.User{}).Count(&count)
	}
	// database.DB.Model(&model.User{}).Count(&count)
	return Success(ctx, count)
}

func DeviceCount(ctx *fiber.Ctx) error {
	var count int64 = 0
	userID := ctx.Query("user_id", "")

	if len(userID) != 0 {
		database.DB.Model(&model.Device{}).Where("user_id =?", userID).Count(&count)
	} else {
		database.DB.Model(&model.Device{}).Count(&count)
	}
	return Success(ctx, count)
}

func DataCount(ctx *fiber.Ctx) error {
	var count int64 = 0
	deviceID := ctx.Query("device_id", "")

	if len(deviceID) != 0 {
		database.DB.Model(&model.DeviceData{}).Where("device_id =?", deviceID).Count(&count)
	} else {
		database.DB.Model(&model.Device{}).Count(&count)
	}
	return Success(ctx, count)
}

func DataCountUserDevice(ctx *fiber.Ctx) error {
	var count int64 = 0
	var devices []model.Device
	userID := ctx.Query("user_id", "")

	err := database.DB.Model(&model.User{ID: userID}).Association("Devices").Find(&devices)

	if err != nil {
		return InternalServer(ctx, err.Error())
	}

	for _, v := range devices {
		var rawCount int64 = 0
		database.DB.Model(&model.DeviceData{}).Where("device_id =?", v.ID).Count(&rawCount)
		count = count + rawCount
	}

	return Success(ctx, count)
}

func UserDashboard(ctx *fiber.Ctx) error {
	var datetimes []time.Time
	var data []int64
	var label []string
	orderBy := ctx.Query("order_by", "week")
	role := ctx.Query("role", "")

	switch orderBy {
	case "second":
		datetimes = Second()
	case "minute":
		datetimes = Minute()
	case "week":
		datetimes = Week()
	case "month":
		datetimes = Month()
	default:
		datetimes = Week()
	}

	for _, v := range datetimes {
		var count int64
		var rangeTime time.Time
		switch orderBy {
		case "second":
			rangeTime = v.Add(time.Second)
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute())+":"+strconv.Itoa(v.Second()))
		case "minute":
			rangeTime = v.Add(time.Minute)
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute()))
		case "week":
			rangeTime = v.Add(time.Hour * 24)
			label = append(label, v.Weekday().String())
		case "month":
			rangeTime = v.AddDate(0, 1, -1)
			label = append(label, v.Month().String())
		default:
			rangeTime = v.Add(time.Hour * 24)
			label = append(label, v.Weekday().String())
		}

		if role != "" {
			database.DB.Model(&model.User{}).Where("created_at BETWEEN ? AND ?", v, rangeTime).Where("role = ?", role).Count(&count)
		} else {
			database.DB.Model(&model.User{}).Where("created_at BETWEEN ? AND ?", v, rangeTime).Count(&count)
		}
		data = append(data, count)
	}

	return Success(ctx, response.Dashboard{
		Label: label,
		Data:  data,
	})
}

func DeviceDashboard(ctx *fiber.Ctx) error {
	var datetimes []time.Time
	var data []int64
	var label []string
	orderBy := ctx.Query("order_by", "week")
	userID := ctx.Query("user_id", "")

	switch orderBy {
	case "second":
		datetimes = Second()
	case "minute":
		datetimes = Minute()
	case "week":
		datetimes = Week()
	case "month":
		datetimes = Month()
	default:
		datetimes = Week()
	}

	for _, v := range datetimes {
		var count int64
		var rangeTime time.Time
		switch orderBy {
		case "second":
			rangeTime = v.Add(time.Second)
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute())+":"+strconv.Itoa(v.Second()))
		case "minute":
			rangeTime = v.Add(time.Minute)
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute()))
		case "week":
			rangeTime = v.Add(time.Hour * 24)
			label = append(label, v.Weekday().String())
		case "month":
			rangeTime = v.AddDate(0, 1, -1)
			label = append(label, v.Month().String())
		default:
			rangeTime = v.Add(time.Hour * 24)
			label = append(label, v.Weekday().String())
		}

		if userID != "" {
			database.DB.Model(&model.Device{}).Where("created_at BETWEEN ? AND ?", v, rangeTime).Where("user_id = ?", userID).Count(&count)
		} else {
			database.DB.Model(&model.Device{}).Where("created_at BETWEEN ? AND ?", v, rangeTime).Count(&count)
		}
		data = append(data, count)
	}

	return Success(ctx, response.Dashboard{
		Label: label,
		Data:  data,
	})
}

func DeviceDataDashboard(ctx *fiber.Ctx) error {
	var datetimes []time.Time
	var data []int64
	var label []string
	orderBy := ctx.Query("order_by", "week")
	deviceID := ctx.Query("device_id", "")
	userID := ctx.Query("user_id", "")

	switch orderBy {
	case "minute":
		datetimes = Minute()
	case "second":
		datetimes = Second()
	case "week":
		datetimes = Week()
	case "month":
		datetimes = Month()
	default:
		datetimes = Week()
	}

	for _, v := range datetimes {
		var count int64
		var rangeTime time.Time
		switch orderBy {
		case "second":
			rangeTime = v.Add(time.Second)
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute())+":"+strconv.Itoa(v.Second()))
		case "minute":
			rangeTime = v.Add(time.Minute)
			label = append(label, strconv.Itoa(v.Hour())+":"+strconv.Itoa(v.Minute()))
		case "week":
			rangeTime = v.Add(time.Hour * 24)
			label = append(label, v.Weekday().String())
		case "month":
			rangeTime = v.AddDate(0, 1, -1)
			label = append(label, v.Month().String())
		default:
			rangeTime = v.Add(time.Hour * 24)
			label = append(label, v.Weekday().String())
		}

		if deviceID != "" {
			database.DB.Model(&model.DeviceData{}).Where("created_at BETWEEN ? AND ?", v, rangeTime).Where("device_id = ?", deviceID).Count(&count)
		} else if userID != "" {
			var devices []model.Device
			database.DB.Model(&model.Device{}).Where("user_id = ?", userID).Find(&devices)
			var rescount int64 = 0
			for _, device := range devices {
				var vcount int64
				database.DB.Model(&model.DeviceData{}).Where("created_at BETWEEN ? AND ?", v, rangeTime).Where("device_id = ?", device.ID).Count(&vcount)
				rescount += vcount
			}
			count = rescount
		} else {
			database.DB.Model(&model.DeviceData{}).Where("created_at BETWEEN ? AND ?", v, rangeTime).Count(&count)
		}
		data = append(data, count)
	}

	return Success(ctx, response.Dashboard{
		Label: label,
		Data:  data,
	})
}

func Minute() []time.Time {
	var data []time.Time
	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()
	currentHour, currentMinute, _ := now.Clock()
	currentLocation := now.Location()

	for i := 0; i < 6; i++ {
		currentDate := time.Date(currentYear, currentMonth, currentDay, currentHour, (currentMinute - i), 0, 0, currentLocation)
		data = append(data, currentDate)
	}

	return data
}

func Second() []time.Time {
	var data []time.Time
	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()
	currentHour, currentMinute, currentSecond := now.Clock()
	currentLocation := now.Location()

	for i := 0; i < 6; i++ {
		currentDate := time.Date(currentYear, currentMonth, currentDay, currentHour, currentMinute, (currentSecond - i), 0, currentLocation)
		data = append(data, currentDate)
	}

	return data
}

func Week() []time.Time {
	var data []time.Time
	now := time.Now()

	for i := 0; i < 6; i++ {
		data = append(data, now.AddDate(0, 0, (i*-1)).Truncate(24*time.Hour))
	}

	return data
}

func Month() []time.Time {
	var data []time.Time
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	currentDate := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	for i := 0; i < 6; i++ {
		data = append(data, currentDate.AddDate(0, (i*-1), 0))
	}

	return data
}
