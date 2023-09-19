package model

import (
	"fmt"
	"restEcho1/configs"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitModel(config configs.ProgramConfig) *gorm.DB {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Model : cannot connect to database, ", err.Error())
		return nil

	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Users{})
}
