package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateNewCart(cart *models.Carts) (interface{}, int, error) {
	result := config.Db.Create(&cart)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return cart, 1, nil
	}
	return nil, 0, nil
}

func GetAllCarts(carts *[]models.Carts) (interface{}, int, error) {
	result := config.Db.Preload("CartsDetail").Find(&carts)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return carts, 1, nil
	}
	return nil, 0, nil
}

func GetSingleCart(cartId int, cart *models.Carts) (interface{}, int, error) {
	result := config.Db.Preload("CartsDetail").Find(&cart, cartId)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return cart, 1, nil
	}
	return nil, 0, nil
}

func GetCustomersCarts(customerId int) ([]models.Carts, int, error) {
	carts := []models.Carts{}
	result := config.Db.Preload("CartsDetail").Where("customers_id = ?", customerId).Find(&carts)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return carts, 1, nil
	}
	return nil, 0, nil
}

func GetSingleCustomersCart(customerId int, cartId int) (models.Carts, int, error) {
	cart := models.Carts{}
	result := config.Db.Preload("CartsDetail").Where("customers_id = ?", customerId).Find(&cart, cartId)
	if result.Error != nil {
		return models.Carts{}, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return cart, 1, nil
	}
	return models.Carts{}, 0, nil
}

func DeleteCart(cartId int) (interface{}, int, error) {
	cart := models.Carts{}
	result := config.Db.Where("cart_status = 'active'").Delete(&cart, cartId)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return "deleted", 1, nil
	}
	return nil, 0, nil
}

func UpdateStatusCart(cartId int, newCart *models.Carts) (interface{}, int, error) {
	cart := models.Carts{}
	result := config.Db.Find(&cart, cartId)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected > 0 {
		tx := config.Db.Model(&cart).Updates(newCart)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}
		return cart, 1, nil
	}
	return "cart id not found", 0, nil
}
