package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"trill/store"
	"trill/types"
)

func Login(c echo.Context) error {

	var userInfo types.APIUser

	email := c.FormValue("email")
	password := c.FormValue("password")

	user := store.GetUserByEmail(email)
	userInfo.ID = uint64(user.ID)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	cookie := new(http.Cookie)
	cookie.Name = "uid"
	cookie.Value = strconv.FormatUint(uint64(user.ID), 10)
	cookie.Domain = os.Getenv("CORS_ORIGINS")
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 7)
	c.SetCookie(cookie)

	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	}
	sess.Values["authenticated"] = true
	sess.Values["uid"] = strconv.FormatUint(uint64(user.ID), 10)

	sess.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, userInfo)
}

func Logout(c echo.Context) error {

	cookie := new(http.Cookie)
	cookie.Name = "uid"
	cookie.Value = "0"
	cookie.Domain = os.Getenv("CORS_ORIGINS")
	cookie.Path = "/"
	cookie.Expires = time.Unix(0, 0)
	c.SetCookie(cookie)

	sess, _ := session.Get("session", c)

	sess.Options.MaxAge = -1

	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusNoContent)
}
