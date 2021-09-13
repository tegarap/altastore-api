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

func AddProductOnCart(c echo.Context) error {
	customerId, isAdmin := middleware.ExtractToken(c)
	if isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only customers can see their own cart", nil))
	}

	cartDetail := models.CartsDetail{}
	c.Bind(&cartDetail)

	newCartDetail, rowAffected, err := database.AddProductOnCart(&cartDetail, customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to add product on cart", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to add product on cart", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully add product on cart", newCartDetail))
}

func UpdatedProductOnCartController(c echo.Context) error {
	customerId, isAdmin := middleware.ExtractToken(c)
	if isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only customers can see their own cart", nil))
	}

	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid cart detail id", nil))
	}

	updatedDetail := models.CartsDetail{}
	c.Bind(&updatedDetail)

	updatedProduct, rowAffected, err := database.UpdatedProductOnCart(customerId, cartDetailId, &updatedDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update product on cart", nil))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update product on cart", nil))

	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully update product on cart", updatedProduct))
}

func DeleteProductOnCartController(c echo.Context) error {
	customerId, isAdmin := middleware.ExtractToken(c)
	if isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only customers can see their own cart", nil))
	}

	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product on cart", nil))
	}
	result, rowAffected, err := database.DeleteProductOnCart(cartDetailId, customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product on cart", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product on cart", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully delete product on cart", result))
}

// func GetAllCartDetailController(c echo.Context) error {
// 	cartDetails := []models.CartsDetail{}
// 	findAllDetails, rowAffected, err := database.GetAllCartDetail(&cartDetails)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all cart details", nil))
// 	}
// 	if rowAffected == 0 {
// 		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all cart details", nil))
// 	}
// 	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get all cart details", findAllDetails))
// }

// func GetSingleCartDetailController(c echo.Context) error {
// 	cartDetailId, errorId := strconv.Atoi(c.Param("id"))
// 	if errorId != nil {
// 		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid cart detail id", nil))
// 	}

// 	findDetails, rowAffected, err := database.GetSingleCartDetail(cartDetailId)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get cart detail", nil))
// 	}
// 	if rowAffected == 0 {
// 		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get cart detail", nil))
// 	}
// 	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get cart detail", findDetails))
// }
