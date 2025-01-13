package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		sess, _ := session.Get("session", c)

		if sess.Values["authenticated"] != true {
			return c.NoContent(http.StatusUnauthorized)
		}

		return next(c)
	}
}

func AuthOnlySelf(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		if sess.Values["authenticated"] != true {
			return c.NoContent(http.StatusUnauthorized)
		}

		uid := c.Param("uid")

		if uid != sess.Values["uid"] {
			return c.NoContent(http.StatusForbidden)
		}

		return next(c)
	}
}
