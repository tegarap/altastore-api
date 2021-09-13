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
	Db.AutoMigrate(&models.Transactions{})
}


//==============================================================
//	CONFIG FOR UNIT TESTING
//==============================================================

func InitDbTest(kind string) {
	connection := os.Getenv("CONNECTION_TEST")

	var err error
	Db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	switch kind {
	case "admin":
		InitMigrateAdminTest()
	case "cart":
		InitMigrateCartTest()
	case "cart detail":
		InitMigrateCartDetailTest()
	case "category":
		InitMigrateCategoryTest()
	case "customer":
		InitMigrateCustomerTest()
	case "payment":
		InitMigratePaymentTest()
	case "product":
		InitMigrateProductTest()
	case "transaction":
		InitMigrateTransactionTest()
	}
}
func InitMigrateAdminTest() {
	Db.Migrator().DropTable(&models.Admins{})
	Db.AutoMigrate(&models.Admins{})
}

func InitMigrateCartTest() {
	Db.Migrator().DropTable(&models.Carts{})
	Db.AutoMigrate(&models.Carts{})
}

func InitMigrateCartDetailTest() {
	Db.Migrator().DropTable(&models.CartsDetail{})
	Db.AutoMigrate(&models.CartsDetail{})
}

func InitMigrateCategoryTest() {
	Db.Migrator().DropTable(&models.Categories{})
	Db.AutoMigrate(&models.Categories{})
}

func InitMigrateCustomerTest() {
	Db.Migrator().DropTable(&models.Customers{})
	Db.AutoMigrate(&models.Customers{})
}

func InitMigratePaymentTest() {
	Db.Migrator().DropTable(&models.Payments{})
	Db.AutoMigrate(&models.Payments{})
}

func InitMigrateProductTest() {
	Db.Migrator().DropTable(&models.Products{})
	Db.AutoMigrate(&models.Products{})
}

func InitMigrateTransactionTest() {
	Db.Migrator().DropTable(&models.Transactions{})
	Db.AutoMigrate(&models.Transactions{})
}







