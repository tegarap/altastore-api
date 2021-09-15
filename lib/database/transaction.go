package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)
//backup
//func CreateTransaction(transaction *models.Transactions) (interface{}, error) {
//	if err := config.Db.Create(&transaction).Error; err != nil {
//		return nil, err
//	}
//	return transaction, nil
//}


func CreateTransaction(transaction *models.Transactions) (interface{}, error) {
	if err := config.Db.Select("CustomersID", "PaymentsID", "CartsID", "TotalPrice").Create(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func GetTransactionByAllCustomer(transactions *[]models.Transactions) (interface{}, error) {
	if err := config.Db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetTransaction(customerId int) (interface{}, error) {
	transaction := models.Transactions{}

	if err := config.Db.Where("customers_id = ?", customerId).Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func GetTotalPrice(cartId uint) (int, error) {
	type Result struct {
		Price int
		Qty   int
	}

	var result []Result

	if err := config.Db.Raw("select price, quantity from products inner join carts_details cd on products.id = cd.products_id where carts_id = ?", cartId).Scan(&result).Error; err != nil {
		return 0, err
	}

	var totalPrice int

	for _, r := range result {
		totalPrice += r.Price * r.Qty
	}

	return totalPrice, nil
}
