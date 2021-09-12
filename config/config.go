package config

import (
	"os"

	"github.com/tegarap/altastore-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	connection := os.Getenv("CONNECTION")

	var err error
	Db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	Db.AutoMigrate(&models.Admins{})
	Db.AutoMigrate(&models.Carts{})
	Db.AutoMigrate(&models.CartsDetail{})
	Db.AutoMigrate(&models.Categories{})
	Db.AutoMigrate(&models.Customers{})
	Db.AutoMigrate(&models.Payments{})
	Db.AutoMigrate(&models.Products{})
	//Db.AutoMigrate(&models.Transactions{})
}
