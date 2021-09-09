package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func initDb() {
	connection := os.Getenv("CONNECTION")

	var err error
	Db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	Db.AutoMigrate()
	Db.AutoMigrate()
}
