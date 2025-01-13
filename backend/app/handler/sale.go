package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"trill/store"
)

func GetSalesItemList(c echo.Context) error {

	uid := c.Param("uid")

	items, err := store.GetItemsByCreatorId(uid)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}

func GetSalesHistory(c echo.Context) error {

	uid := c.Param("uid")

	sales, err := store.GetSalesHistory(uid)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, sales)
}
