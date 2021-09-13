package category

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func CreateNewCategoriesController(c echo.Context) error {
	_, isAdmin := middleware.ExtractToken(c)
	if !isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only admin can create new product category", nil))
	}

	category := models.Categories{}
	c.Bind(&category)

	newCategory, rowAffected, err := database.CreateNewCategories(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to create new product category", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to create new product category", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully create new product category", newCategory))
}

func GetAllCategoriesController(c echo.Context) error {
	category, rowAffected, err := database.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all product categories", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all product categories", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get all product categories", category))
}

func GetSingleCategoryController(c echo.Context) error {
	categoryId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid category id", nil))
	}
	category, rowAffected, err := database.GetSingleCategory(categoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get product category", nil))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get product category", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get product category", category))
}

func DeleteCategoryController(c echo.Context) error {
	_, isAdmin := middleware.ExtractToken(c)
	if !isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only admin can delete product category", nil))
	}

	categoryId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid category id", nil))
	}

	deletedCategory, rowAffected, err := database.DeleteCategory(categoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product category", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product category", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully delete product category", deletedCategory))
}

func UpdateCategoryController(c echo.Context) error {
	_, isAdmin := middleware.ExtractToken(c)
	if !isAdmin {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only admin can update product category", nil))
	}
	categoryId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update product category", nil))
	}

	newCategory := models.Categories{}
	c.Bind(&newCategory)

	updatedCategory, rowAffected, err := database.UpdateCategory(categoryId, &newCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update product category", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update product category", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully update product category", updatedCategory))
}
