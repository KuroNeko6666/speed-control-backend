package database

import (
	"log"

	"github.com/KuroNeko6666/speed-control-backend.git/config"
	"github.com/KuroNeko6666/speed-control-backend.git/database/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DatabaseDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(
		&model.User{},
		&model.Device{},
	)

	if err != nil {
		log.Fatal(err)
	}

}
