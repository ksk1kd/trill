package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"trill/store"
)

func GetCategoryList(c echo.Context) error {

	categories, err := store.GetCategories()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, categories)
}
