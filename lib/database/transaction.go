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
