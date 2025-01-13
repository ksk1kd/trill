package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"trill/store"
	"trill/types"
	"trill/utilities"
)

func AddCartItem(c echo.Context) error {

	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	}

	iid := c.FormValue("id")

	_, err := store.GetItem(iid)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if sess.Values["cart"] == nil {
		sess.Values["cart"] = []string{}
	}

	if !utilities.SliceContains(sess.Values["cart"].([]string), iid) {
		sess.Values["cart"] = append(sess.Values["cart"].([]string), iid)
	}
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusNoContent)
}

func GetCartItem(c echo.Context) error {

	var items types.APIItems

	sess, _ := session.Get("session", c)

	if sess.Values["cart"] == nil {
		items.Quantity = 0
		return c.JSON(http.StatusOK, items)
	}

	cart := sess.Values["cart"].([]string)
	ids := make([]string, len(cart))
	copy(ids, cart)

	if len(cart) == 0 {
		items.Quantity = 0
		return c.JSON(http.StatusOK, items)
	}

	items, err := store.GetItemsByIds(ids)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}

func DeleteCartItem(c echo.Context) error {

	sess, _ := session.Get("session", c)

	iid := c.Param("iid")

	sess.Values["cart"] = utilities.SliceRemove(sess.Values["cart"].([]string), iid)
	sess.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusNoContent)
}
