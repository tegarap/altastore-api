package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreatePayment(payment *models.Payments) (interface{}, error) {
	if err := config.Db.Create(&payment).Error; err != nil {
		return payment, err
	}
	outPayment := models.PaymentResponse{}
	return outPayment, nil
}

func GetAllPayment(payments *[]models.Payments) (interface{}, error) {
	if err := config.Db.Find(&payments).Error; err != nil {
		return payments, err
	}
	outPayment := models.PaymentResponse{}
	return outPayment, nil
}

func GetSinglePayment(id int) (interface{}, error) {
	payment := models.Payments{}
	if err := config.Db.First(&payment, id).Error; err != nil {
		return payment, err
	}
	outPayment := models.PaymentResponse{}
	return outPayment, nil
}