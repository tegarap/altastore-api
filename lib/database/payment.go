package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreatePayment(payment *models.Payments) (*models.Payments, error) {
	if err := config.Db.Create(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}

func GetAllPayment(payments *[]models.Payments) (*[]models.Payments, error) {
	if err := config.Db.Find(&payments).Error; err != nil {
		return payments, err
	}
	return payments, nil
}

func GetSinglePayment(id int) (models.Payments, error) {
	payment := models.Payments{}
	if err := config.Db.First(&payment, id).Error; err != nil {
		return payment, err
	}
	return payment, nil
}