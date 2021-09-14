package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/models"
	"gorm.io/gorm/clause"
)

func LoginCustomer(customer *models.Customers) (interface{}, error){
	if er := config.Db.Where("email = ? AND password = ?", customer.Email, customer.Password).First(customer).Error; er != nil {
		return nil, er
	}

	var err error
	customer.Token, err = middleware.CreateToken(int(customer.ID), false)

	if err != nil {
		return nil, err
	}

	if err = config.Db.Save(customer).Error; err != nil{
		return nil, err
	}

	return customer, nil
}

func RegisterCustomer(customer *models.Customers) (interface{}, error) {
	if err := config.Db.Create(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func GetCustomerProfile(id int) (interface{}, error) {
	var customer models.Customers

	if err := config.Db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func GetAllCustomers(customers *[]models.Customers) (interface{}, error) {
	if err := config.Db.Preload("Carts.CartsDetail").Preload(clause.Associations).Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}
func GetAllCustomerSimple(customers *[]models.Customers) (interface{}, error) {
	if err := config.Db.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}
