package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateDelivery(delivery *models.Deliveries) (interface{}, error) {
	if err := config.Db.Create(&delivery).Error; err != nil {
		return nil, err
	}
	return delivery, nil
}

func GetAllDelivery(deliveries *[]models.Deliveries) (interface{}, error) {
	if err := config.Db.Find(&deliveries).Error; err != nil {
		return nil, err
	}
	return deliveries, nil
}

func GetSingleDelivery(id int) (interface{}, error) {
	delivery := models.Deliveries{}
	if err := config.Db.Find(&delivery, id).Error; err != nil {
		return nil, err
	}
	return delivery, nil
}