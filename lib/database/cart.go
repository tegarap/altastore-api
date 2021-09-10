package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateNewCart(cart *models.Carts) (interface{}, error) {
	result := config.Db.Create(&cart)
	if result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}

func GetAllCarts(carts *[]models.Carts) (interface{}, error) {
	result := config.Db.Find(&carts)
	if result.Error != nil {
		return nil, result.Error
	}
	return carts, nil
}

func GetSingleCart(cartId int, cart *models.Carts) (interface{}, error) {
	result := config.Db.Find(&cart, cartId)
	if result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}
