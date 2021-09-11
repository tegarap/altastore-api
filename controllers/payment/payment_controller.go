package payment

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
	"strconv"
)

func CreatePaymentController(c echo.Context) error {
	var payment models.Payments
	c.Bind(&payment)

	newPayment, err := database.CreatePayment(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Create Payment Method", nil))
	}

	response := Response{
		ID:          newPayment.ID,
		PaymentName: newPayment.PaymentName,
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Create New Payment Method", response))
}

func GetAllPaymentController(c echo.Context) error {
	var payment []models.Payments

	payments, err := database.GetAllPayment(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get All Payment Method", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get All Payment Method", payments))
}

func GetSinglePaymentController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid Parameter", nil))
	}

	payment, errr := database.GetSinglePayment(id)
	if errr != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get Payment Method", nil))
	}

	response := Response{
		ID:          payment.ID,
		PaymentName: payment.PaymentName,
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Payment Method", response))
}
