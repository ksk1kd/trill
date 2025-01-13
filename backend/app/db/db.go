package db

import (
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"trill/models"
	"trill/seeds"
)

var Connection *gorm.DB

func Init(seeding bool) {

	if err := godotenv.Load(".env.local"); err != nil {
		slog.Error("Error loading .env.local file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error

	for i := 0; i < 30; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			break
		}

		slog.Error("Failed to connect to database. Retrying...")

		time.Sleep(time.Second)
	}

	if seeding {
		db.Migrator().DropTable(
			&models.User{},
			&models.Category{},
			&models.Item{},
			&models.Purchase{},
			&models.PurchaseDetail{},
		)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Item{},
		&models.Purchase{},
		&models.PurchaseDetail{},
	)

	if seeding {
		seeds.RunAll(db)
	}

	Connection = db
}

func GetConnection() *gorm.DB {
	return Connection
}
