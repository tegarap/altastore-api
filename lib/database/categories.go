package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateNewCategories(category *models.Categories) (interface{}, error) {
	err := config.Db.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func GetSingleCategory(categoryId int) (interface{}, error) {
	category := models.Categories{}
	result := config.Db.Preload("Products").Find(&category, categoryId)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
