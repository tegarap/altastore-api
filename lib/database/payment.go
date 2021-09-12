package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreatePaymentMethod(payment *models.Payments) (interface{}, error) {
	if err := config.Db.Create(&payment).Error; err != nil {
		return payment, err
	}

	return payment, nil
}

func GetAllPaymentMethod(payments *[]models.Payments) (interface{}, error) {
	if err := config.Db.Preload("Transactions").Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

func GetSinglePaymentMethod(id int) (interface{}, error) {
	payment := models.Payments{}

	if err := config.Db.Preload("Transactions").First(&payment, id).Error; err  != nil {
		return nil, err
	}

	return payment, nil
}
