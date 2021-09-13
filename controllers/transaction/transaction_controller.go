package transaction

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
)

func CreateTransactionController(c echo.Context) error {
	id, isAdmin := middleware.ExtractToken(c)

	if isAdmin == true {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only Customer can Access", nil))
	}

	var transaction models.Transactions
	c.Bind(&transaction)

	transaction.CustomersID = uint(id)

	if transaction.PaymentsID == 0 || transaction.CartsID == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Field are required", nil))
	}

	newTransaction, err := database.CreateTransaction(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Create Transaction", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseFail("Success Create Transaction", newTransaction))
}

func GetTransactionController(c echo.Context) error {
	id, isAdmin := middleware.ExtractToken(c)

	if isAdmin == true {
		var transaction []models.Transactions
		transactions, err := database.GetTransactionByAllCustomer(&transaction)

		if err != nil {
			return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get All Transaction", nil))
		}

		return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get All Transaction", transactions))
	} else {
		transaction, err := database.GetTransaction(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get Transaction", nil))
		}

		return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Transaction", transaction))
	}
}