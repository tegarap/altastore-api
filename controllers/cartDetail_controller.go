package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func CreateNewCartDetailController(c echo.Context) error {
	cartDetail := models.CartsDetail{}
	c.Bind(&cartDetail)

	newCartDetail, err := database.CreateNewCartDetail(&cartDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", newCartDetail))
}

func GetAllCartDetailController(c echo.Context) error {
	cartDetails := []models.CartsDetail{}
	findAllDetails, err := database.GetAllCartDetail(&cartDetails)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", findAllDetails))
}

func GetSingleCartDetailController(c echo.Context) error {
	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	findDetails, err := database.GetSingleCartDetail(cartDetailId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", findDetails))
}

func UpdatedProductOnCartController(c echo.Context) error {
	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	updatedDetail := models.CartsDetail{}
	c.Bind(&updatedDetail)

	updatedProduct, rowAffected, err := database.UpdatedProductOnCart(cartDetailId, &updatedDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))

	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", updatedProduct))
}

func DeleteProductOnCartController(c echo.Context) error {
	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	result, err := database.DeleteProductOnCart(cartDetailId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", result))
}
