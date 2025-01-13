package handler

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"

	"trill/models"
	"trill/store"
	"trill/types"
)

func AddItem(c echo.Context) error {

	sess, _ := session.Get("session", c)

	uid := sess.Values["uid"]
	if uid == "" {
		return c.NoContent(http.StatusForbidden)
	}

	image, _ := c.FormFile("image")
	category := c.FormValue("category")
	title := c.FormValue("title")
	price := c.FormValue("price")

	converted_uid, _ := strconv.ParseUint(uid.(string), 10, 0)
	converted_category, _ := strconv.ParseUint(category, 10, 0)
	converted_price, _ := strconv.ParseUint(price, 10, 0)

	src, err := image.Open()
	if err != nil {
		slog.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	defer src.Close()

	var validation types.APIError
	buffer := make([]byte, 1)
	_, err = src.Read(buffer)
	if err != nil {
		validation.Messages = append(validation.Messages, "画像ファイルは必須項目です")
	}
	if category == "" {
		validation.Messages = append(validation.Messages, "カテゴリは必須項目です")
	}
	if title == "" {
		validation.Messages = append(validation.Messages, "タイトルは必須項目です")
	}
	if price == "" {
		validation.Messages = append(validation.Messages, "価格は必須項目です")
	}
	if converted_price < 100 {
		validation.Messages = append(validation.Messages, "価格は100以上の数値を入力してください")
	}
	if validation.Messages != nil {
		return c.JSON(http.StatusBadRequest, validation)
	}

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		slog.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	image_name := uuid.New().String()
	image_path := "/images/" + image_name + ".png"
	file, err := os.Create("../public" + image_path)
	if err != nil {
		slog.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	defer file.Close()

	_, err = io.Copy(file, src)
	if err != nil {
		slog.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	item := models.Item{
		Title:      title,
		Image:      image_path,
		CategoryID: converted_category,
		Price:      converted_price,
		CreatorID:  converted_uid,
	}
	iid, err := store.AddItem(item)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, iid)
}

func GetItem(c echo.Context) error {

	sess, _ := session.Get("session", c)

	uid := sess.Values["uid"]

	iid := c.Param("iid")

	item, err := store.GetItem(iid)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if strconv.FormatUint(uint64(item.Creator), 10) == uid {
		item.Seller = true
	}

	item.Purchased = false
	if uid != nil {
		purchasedItems, err := store.GetPurchasedItems(uid.(string))
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		for _, purchasedItem := range purchasedItems.Items {
			if strconv.FormatUint(uint64(purchasedItem.ID), 10) == iid {
				item.Purchased = true
			}
		}
	}

	if sess.Values["cart"] == nil {
		item.CartAdded = false
	} else {
		cartItems := sess.Values["cart"].([]string)
		for _, carItem := range cartItems {
			if carItem == iid {
				item.CartAdded = true
			}
		}
	}

	return c.JSON(http.StatusOK, item)
}

func GetItemList(c echo.Context) error {

	q := c.QueryParam("q")
	category := c.QueryParam("category")
	if category == "0" {
		category = ""
	}

	items, err := store.SearchItems(q, category)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, items)
}

func UpdateItem(c echo.Context) error {

	sess, _ := session.Get("session", c)

	uid := sess.Values["uid"]
	if uid == "" {
		return c.NoContent(http.StatusForbidden)
	}

	iid := c.Param("iid")

	price := c.FormValue("price")
	var validation types.APIError
	if price == "" {
		validation.Messages = append(validation.Messages, "価格は必須項目です")
	}
	converted_price, _ := strconv.Atoi(price)
	if converted_price < 100 {
		validation.Messages = append(validation.Messages, "価格は100以上の数値を入力してください")
	}
	if validation.Messages != nil {
		return c.JSON(http.StatusBadRequest, validation)
	}

	err := store.UpdateItemPrice(iid, price, uid.(string))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteItem(c echo.Context) error {

	sess, _ := session.Get("session", c)

	uid := sess.Values["uid"]
	if uid == "" {
		return c.NoContent(http.StatusForbidden)
	}

	iid := c.Param("iid")

	err := store.DeleteItem(iid, uid.(string))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}
