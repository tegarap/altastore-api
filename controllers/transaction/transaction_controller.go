package transaction

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
	"strconv"
)

func CreateTransactionController(c echo.Context) error {
	var transaction models.Transactions
	c.Bind(&transaction)

	newTransaction, err := database.CreateTransaction(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Create Transaction", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseFail("Success Create Transaction", newTransaction))
}

func GetAllTransactionController(c echo.Context) error {
	var transaction []models.Transactions
	transactions, err := database.GetAllTransaction(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get All Transaction", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get All Transaction", transactions))
}

func GetSingleTransactionController(c echo.Context) error {
	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid Parameter", nil))
	}

	transaction, err := database.GetSingleTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get Single Transaction", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Single Transaction", transaction))
}