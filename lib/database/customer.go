package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
	"gorm.io/gorm/clause"
)

func LoginCustomer(customer *models.Customers) (interface{}, error){
	if err := config.Db.Where("email = ? AND password = ?", customer.Email, customer.Password).First(customer).Error; err != nil {
		return nil, err
	}
	outCustomer := models.CustomerResponse{}
	return outCustomer, nil
}

func RegisterCustomer(customer *models.Customers) (interface{}, error) {
	if err := config.Db.Create(&customer).Error; err != nil {
		return nil, err
	}
	outCustomer := models.CustomerResponse{}
	return outCustomer, nil
}

func GetAllCustomers(customers *[]models.Customers) (interface{}, error) {
	if err := config.Db.Preload("Carts.CartsDetail").Preload(clause.Associations).Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
