package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func GetAllProduct(product *[]models.Products) (interface{}, error) {
	err := config.Db.Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetSingleProduct(productId int) (interface{}, error) {
	product := models.Products{}
	err := config.Db.Find(&product, productId).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func CreateNewProduct(products *models.Products) (interface{}, error) {
	err := config.Db.Create(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
