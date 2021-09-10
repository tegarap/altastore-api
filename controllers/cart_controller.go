package controllers

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

	newCart, err := database.CreateNewCart(&cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", newCart))
}

func GetAllCartsController(c echo.Context) error {
	carts := []models.Carts{}
	allCarts, err := database.GetAllCarts(&carts)
	if err != nil {
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
	foundCart, err := database.GetSingleCart(cartId, &cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", foundCart))
}
