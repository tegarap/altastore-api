package cart

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func AddProductOnCart(c echo.Context) error {
	cartDetail := models.CartsDetail{}
	c.Bind(&cartDetail)

	newCartDetail, rowAffected, err := database.AddProductOnCart(&cartDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to add product on cart", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to add product on cart", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully add product on cart", newCartDetail))
}

func GetAllCartDetailController(c echo.Context) error {
	cartDetails := []models.CartsDetail{}
	findAllDetails, rowAffected, err := database.GetAllCartDetail(&cartDetails)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all cart details", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all cart details", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get all cart details", findAllDetails))
}

func GetSingleCartDetailController(c echo.Context) error {
	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid cart detail id", nil))
	}

	findDetails, rowAffected, err := database.GetSingleCartDetail(cartDetailId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get cart detail", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get cart detail", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get cart detail", findDetails))
}

func UpdatedProductOnCartController(c echo.Context) error {
	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid cart detail id", nil))
	}

	updatedDetail := models.CartsDetail{}
	c.Bind(&updatedDetail)

	updatedProduct, rowAffected, err := database.UpdatedProductOnCart(cartDetailId, &updatedDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update product on cart", nil))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update product on cart", nil))

	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully update product on cart", updatedProduct))
}

func DeleteProductOnCartController(c echo.Context) error {
	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product on cart", nil))
	}
	result, rowAffected, err := database.DeleteProductOnCart(cartDetailId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product on cart", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product on cart", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully delete product on cart", result))
}
