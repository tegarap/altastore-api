package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateNewCartDetail(cartDetail *models.CartsDetail) (interface{}, error) {
	result := config.Db.Create(&cartDetail)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartDetail, nil
}

func GetAllCartDetail(cartDetail *[]models.CartsDetail) (interface{}, error) {
	result := config.Db.Find(&cartDetail)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartDetail, nil
}

func GetSingleCartDetail(cartDetailId int) (interface{}, error) {
	cartDetail := models.CartsDetail{}
	result := config.Db.Find(&cartDetail, cartDetailId)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartDetail, nil
}

func UpdatedProductOnCart(cartDetailId int, updatedDetail *models.CartsDetail) (interface{}, int, error) {
	cartDetails := models.CartsDetail{}
	findId := config.Db.Find(&cartDetails, cartDetailId)
	if findId.Error != nil {
		return nil, 0, findId.Error
	}

	if findId.RowsAffected > 0 {
		saveUpdated := config.Db.Model(&cartDetails).Updates(updatedDetail)
		if saveUpdated.Error != nil {
			return nil, 0, saveUpdated.Error
		}
		return cartDetails, 1, nil
	}
	return "product id not found", 0, nil
}

func DeleteProductOnCart(cartDetailId int) (interface{}, error) {
	cartDetail := models.CartsDetail{}
	err := config.Db.Delete(&cartDetail, cartDetailId).Error
	if err != nil {
		return nil, err
	}
	return "deleted", nil
}
