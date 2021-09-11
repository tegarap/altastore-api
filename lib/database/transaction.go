package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateTransaction(transaction *models.Transactions) (interface{}, error) {
	if err := config.Db.Create(&transaction).Error; err != nil {
		return nil, err
	}
	outTransaction := models.TransactionResponse{}
	return outTransaction, nil
}

func GetAllTransaction(transactions *[]models.Transactions) (interface{}, error) {
	if err := config.Db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	outTransaction := models.TransactionResponse{}
	return outTransaction, nil
}

func GetSingleTransaction(id int) (interface{}, error) {
	var transaction models.Transactions
	if err := config.Db.Find(&transaction, id).Error; err != nil {
		return transaction, err
	}
	outTransaction := models.TransactionResponse{}
	return outTransaction, nil
}