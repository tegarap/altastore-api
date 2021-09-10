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
