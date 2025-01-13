package store

import (
	"log/slog"
	"trill/db"
	"trill/models"
)

func GetUserByEmail(email string) models.User {

	var user models.User

	if result := db.Connection.Where("email = ?", email).Find(&user); result.Error != nil {
		slog.Error(result.Error.Error())
	}

	return user
}
