package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func AddProductOnCart(cartDetail *models.CartsDetail, customerId int) (interface{}, int, error) {
	cart := models.Carts{}
	findCart := config.Db.Find(&cart, cartDetail.CartsID)
	if findCart.Error != nil {
		return nil, 0, findCart.Error
	}

	if cart.CustomersID != uint(customerId) {
		return nil, 0, findCart.Error
	}

	product := models.Products{}
	findProduct := config.Db.Find(&product, cartDetail.ProductsID)
	if findProduct.Error != nil {
		return nil, 0, findProduct.Error
	}
	if product.Stock-cartDetail.Quantity < 0 {
		return nil, 0, findProduct.Error
	}

	product.Stock -= cartDetail.Quantity
	updatedProduct := config.Db.Save(&product)
	if updatedProduct.Error != nil {
		return nil, 0, updatedProduct.Error
	}

	result := config.Db.Create(&cartDetail)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return cartDetail, 1, nil
}

func UpdatedProductOnCart(customerId int, cartDetailId int, updatedDetail *models.CartsDetail) (interface{}, int, error) {
	cartDetails := models.CartsDetail{}
	findId := config.Db.Find(&cartDetails, cartDetailId)
	if findId.Error != nil {
		return nil, 0, findId.Error
	}

	cart := models.Carts{}
	findCart := config.Db.Find(&cart, cartDetails.CartsID)
	if findCart.Error != nil {
		return nil, 0, findCart.Error
	}

	if cart.CustomersID != uint(customerId) {
		return nil, 0, findCart.Error
	}

	if findId.RowsAffected > 0 {
		oldProduct := models.Products{}
		findOldProduct := config.Db.Find(&oldProduct, cartDetails.ProductsID)
		if findOldProduct.Error != nil {
			return nil, 0, findOldProduct.Error
		}
		oldProduct.Stock += cartDetails.Quantity
		saveOldProduct := config.Db.Save(&oldProduct)
		if saveOldProduct.Error != nil {
			return nil, 0, saveOldProduct.Error
		}

		updatedProduct := models.Products{}
		findUpdatedProduct := config.Db.Find(&updatedProduct, updatedDetail.ProductsID)
		if findUpdatedProduct.Error != nil {
			return nil, 0, findUpdatedProduct.Error
		}
		updatedProduct.Stock -= updatedDetail.Quantity
		saveUpdatedProduct := config.Db.Save(&updatedProduct)
		if saveUpdatedProduct.Error != nil {
			return nil, 0, saveUpdatedProduct.Error
		}

		saveUpdated := config.Db.Model(&cartDetails).Updates(updatedDetail)
		if saveUpdated.Error != nil {
			return nil, 0, saveUpdated.Error
		}
		return cartDetails, 1, nil
	}
	return "product id not found", 0, nil
}

func DeleteProductOnCart(cartDetailId, customerId int) (interface{}, int, error) {
	cartDetail := models.CartsDetail{}
	findCartDetail := config.Db.Find(&cartDetail, cartDetailId)
	if findCartDetail.Error != nil {
		return nil, 0, findCartDetail.Error
	}

	cart := models.Carts{}
	findCart := config.Db.Find(&cart, cartDetail.CartsID)
	if findCart.Error != nil {
		return nil, 0, findCart.Error
	}

	if cart.CustomersID != uint(customerId) {
		return nil, 0, findCart.Error
	}

	products := models.Products{}
	findProduct := config.Db.Find(&products, cartDetail.ProductsID)
	if findProduct.Error != nil {
		return nil, 0, findProduct.Error
	}

	products.Stock += cartDetail.Quantity
	updatedProduct := config.Db.Save(&products)
	if updatedProduct.Error != nil {
		return nil, 0, updatedProduct.Error
	}

	tx := config.Db.Delete(&cartDetail, cartDetailId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return "deleted", 1, nil
	}
	return nil, 0, nil
}

func GetAllCartDetail(cartDetail *[]models.CartsDetail) (interface{}, int, error) {
	result := config.Db.Find(&cartDetail)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return cartDetail, 1, nil
	}
	return nil, 0, nil
}

func GetSingleCartDetail(cartDetailId int) (interface{}, int, error) {
	cartDetail := models.CartsDetail{}
	result := config.Db.Find(&cartDetail, cartDetailId)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	if result.RowsAffected > 0 {
		return cartDetail, 1, nil
	}
	return nil, 0, nil
}
