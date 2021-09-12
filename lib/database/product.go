package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func GetAllProduct(product *[]models.Products) (interface{}, int, error) {
	tx := config.Db.Preload("CartsDetail").Find(&product)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return product, 1, nil
	}
	return nil, 0, nil
}

func GetSingleProduct(productId int) (interface{}, int, error) {
	product := models.Products{}
	tx := config.Db.Preload("CartsDetail").Find(&product, productId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return product, 1, nil
	}
	return nil, 0, nil
}

func CreateNewProduct(products *models.Products) (interface{}, error) {
	err := config.Db.Create(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func DeleteProduct(productId int) (interface{}, int, error) {
	product := models.Products{}
	result := config.Db.Delete(&product, productId)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return "deleted", 0, nil
	}
	return "failed", 0, nil
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
