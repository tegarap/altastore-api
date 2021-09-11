package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func CreateTransaction(transaction *models.Transactions) (*models.Transactions, error) {
	if err := config.Db.Create(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func GetAllTransaction(transactions *[]models.Transactions) (*[]models.Transactions, error) {
	if err := config.Db.Preload("Deliveries").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetSingleTransaction(id int) (models.Transactions, error) {
	var transaction models.Transactions
	if err := config.Db.Preload("Deliveries").Find(&transaction, id).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}