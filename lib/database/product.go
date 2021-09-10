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

func DeleteProduct(productId int) (interface{}, error) {
	product := models.Products{}
	result := config.Db.Delete(&product, productId)
	if result.Error != nil {
		return nil, result.Error
	}
	return "deleted", nil
}

func UpdateProduct(productId int, newProduct *models.Products) (interface{}, int, error) {
	product := models.Products{}
	findId := config.Db.Find(&product, productId)
	if findId.Error != nil {
		return nil, 0, findId.Error
	}

	if findId.RowsAffected > 0 {
		result := config.Db.Model(&product).Updates(newProduct)
		if result.Error != nil {
			return nil, 0, result.Error
		}
		return product, 1, nil
	}
	return "product not found", 0, nil
}
