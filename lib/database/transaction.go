package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateTransaction(transaction *models.Transactions) (interface{}, error) {
	if err := config.Db.Create(&transaction).Error; err != nil {
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