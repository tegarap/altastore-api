package cart

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func CreateNewCartController(c echo.Context) error {
	customerId, isAdmin := middleware.ExtractToken(c)
	if isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only customer can create new cart", nil))
	}
	cart := models.Carts{}
	c.Bind(&cart)

	newCart, rowAffected, err := database.CreateNewCart(&cart, customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to create new cart", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to create new cart", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully create new cart", newCart))
}

func GetAllCartsController(c echo.Context) error {
	_, isAdmin := middleware.ExtractToken(c)
	if !isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only admin can get all carts", nil))
	}
	carts := []models.Carts{}
	allCarts, rowAffected, err := database.GetAllCarts(&carts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all carts", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all carts", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get all carts", allCarts))
}

func GetSingleCartController(c echo.Context) error {
	_, isAdmin := middleware.ExtractToken(c)
	if !isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only admin can get single carts", nil))
	}

	cartId, errorId := strconv.Atoi(c.QueryParam("cart_id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid cart id", nil))
	}

	cart := models.Carts{}
	foundCart, rowAffected, err := database.GetSingleCart(cartId, &cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get cart data", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get cart data", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get cart data", foundCart))
}

func GetCustomersCartsController(c echo.Context) error {
	customerId, isAdmin := middleware.ExtractToken(c)
	if isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only customers can see their own cart", nil))
	}

	listCarts, rowAffected, err := database.GetCustomersCarts(customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get customer carts", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get customer carts", nil))
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

	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get customer carts", output))
}

func GetSingleCustomersCartController(c echo.Context) error {
	customerId, isAdmin := middleware.ExtractToken(c)
	if isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only customers can see their own cart", nil))
	}

	cartId, errorId := strconv.Atoi(c.Param("cart_id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get customer cart", nil))
	}

	listCarts, rowAffected, err := database.GetSingleCustomersCart(customerId, cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get customer cart", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get customer carts", nil))
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

	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get customer cart", output))
}

func DeleteCartController(c echo.Context) error {
	customerId, isAdmin := middleware.ExtractToken(c)
	if isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only customers can see their own cart", nil))
	}

	cartId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid cart id", nil))
	}

	deletedCart, rowAffected, err := database.DeleteCart(cartId, customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete cart", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete cart", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully delete cart", deletedCart))
}

func UpdatedCartStatusController(c echo.Context) error {
	cartId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid cart id", nil))
	}

	newCart := models.Carts{
		CartStatus: "paid off",
	}

	updatedCart, rowAffected, err := database.UpdateStatusCart(cartId, &newCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update cart status", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update cart status", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully update cart status", updatedCart))
}
