package category

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func CreateNewCategoriesController(c echo.Context) error {
	category := models.Categories{}
	c.Bind(&category)

	newCategory, rowAffected, err := database.CreateNewCategories(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", newCategory))
}

func GetAllCategoriesController(c echo.Context) error {
	category, rowAffected, err := database.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", category))
}

func GetSingleCategoryController(c echo.Context) error {
	categoryId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	category, rowAffected, err := database.GetSingleCategory(categoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("success", category))
}

func DeleteCategoryController(c echo.Context) error {
	categoryId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	deletedCategory, rowAffected, err := database.DeleteCategory(categoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", deletedCategory))
}

func UpdateCategoryController(c echo.Context) error {
	categoryId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	newCategory := models.Categories{}
	c.Bind(&newCategory)

	updatedCategory, rowAffected, err := database.UpdateCategory(categoryId, &newCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", updatedCategory))
}
