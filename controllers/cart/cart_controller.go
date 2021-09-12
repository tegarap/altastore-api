package cart

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func CreateNewCartController(c echo.Context) error {
	cart := models.Carts{}
	c.Bind(&cart)

	newCart, rowAffected, err := database.CreateNewCart(&cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", newCart))
}

func GetAllCartsController(c echo.Context) error {
	carts := []models.Carts{}
	allCarts, rowAffected, err := database.GetAllCarts(&carts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", allCarts))
}

func GetSingleCartController(c echo.Context) error {
	cartId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	cart := models.Carts{}
	foundCart, rowAffected, err := database.GetSingleCart(cartId, &cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", foundCart))
}

func GetCustomersCartsController(c echo.Context) error {
	customerId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	listCarts, rowAffected, err := database.GetCustomersCarts(customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	type Output struct {
		CartId      uint                 `json:"cart_id"`
		CartStatus  string               `json:"cart_status"`
		CartsDetail []models.CartsDetail `json:"cart_detail"`
	}

	output := make([]Output, len(listCarts))
	for i := range listCarts {
		output[i].CartId = listCarts[i].ID
		output[i].CartStatus = listCarts[i].CartStatus
		output[i].CartsDetail = listCarts[i].CartsDetail
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("success", output))
}

func GetSingleCustomersCartController(c echo.Context) error {
	customerId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	cartId, errorId := strconv.Atoi(c.Param("cart_id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	listCarts, rowAffected, err := database.GetSingleCustomersCart(customerId, cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	type Output struct {
		CartId      uint                 `json:"cart_id"`
		CartStatus  string               `json:"cart_status"`
		CartsDetail []models.CartsDetail `json:"cart_detail"`
	}

	output := Output{
		CartId:      listCarts.ID,
		CartStatus:  listCarts.CartStatus,
		CartsDetail: listCarts.CartsDetail,
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("success", output))
}

func DeleteCartController(c echo.Context) error {
	cartId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	deletedCart, rowAffected, err := database.DeleteCart(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", deletedCart))
}

func UpdatedCartStatusController(c echo.Context) error {
	cartId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	newCart := models.Carts{
		CartStatus: "paid off",
	}

	updatedCart, rowAffected, err := database.UpdateStatusCart(cartId, &newCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", updatedCart))
}
