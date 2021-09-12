package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateNewCategories(category *models.Categories) (interface{}, int, error) {
	tx := config.Db.Create(&category)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return category, 1, nil
	}
	return nil, 0, nil
}

func GetAllCategories() (interface{}, int, error) {
	categories := []models.Categories{}
	result := config.Db.Preload("Products").Find(&categories)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return categories, 1, nil
	}
	return nil, 0, nil
}

func GetSingleCategory(categoryId int) (interface{}, int, error) {
	outputProduct := []models.Products{}
	tx := config.Db.Raw("select products.id, products.name, products.description, products.price, products.stock from products where products.categories_id = ?", categoryId).Scan(&outputProduct)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return outputProduct, 1, nil
	}
	return nil, 0, nil
}

func DeleteCategory(categoryId int) (interface{}, int, error) {
	category := models.Categories{}
	tx := config.Db.Delete(&category, categoryId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return "deleted", 1, nil
	}
	return nil, 0, nil
}

func UpdateCategory(categoryId int, newCategory *models.Categories) (interface{}, int, error) {
	category := models.Categories{}
	findId := config.Db.Find(&category, categoryId)
	if findId.Error != nil {
		return nil, 0, findId.Error
	}

	if findId.RowsAffected > 0 {
		tx := config.Db.Model(&category).Updates(&newCategory)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}
		return category, 1, nil
	}
	return "category id not found", 0, nil
}
