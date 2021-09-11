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

func GetAllTransaction(transactions *[]models.Transactions) (interface{}, error) {
	if err := config.Db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetSingleTransaction(id int) (interface{}, error) {
	var transaction models.Transactions
	if err := config.Db.Find(&transaction, id).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}