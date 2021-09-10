package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
)

func CreateTransactionController(c echo.Context) error {
	var transaction models.Transactions
	c.Bind(&transaction)

	newTransaction, err := database.CreateTransaction(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Create Transaction", nil))
	}

	return c.JSON(http.StatusBadRequest, util.ResponseFail("Success Create Transaction", newTransaction))
}

