package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"trill/store"
)

func AddPurchase(c echo.Context) error {

	uid := c.Param("uid")

	sess, _ := session.Get("session", c)

	if sess.Values["cart"] == nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	cart := sess.Values["cart"].([]string)
	ids := make([]string, len(cart))
	copy(ids, cart)

	if len(cart) == 0 {
		return c.NoContent(http.StatusInternalServerError)
	}

	err := store.AddPurchase(uid, ids)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	sess.Values["cart"] = nil
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusNoContent)
}

func GetPurchasedItemList(c echo.Context) error {

	uid := c.Param("uid")

	items, err := store.GetPurchasedItems(uid)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}

func GetPurchaseHistory(c echo.Context) error {

	uid := c.Param("uid")

	purchases, err := store.GetPurchaseHistory(uid)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, purchases)
}
