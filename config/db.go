package config

import (
	entities "coverCraft/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	tz := os.Getenv("SERVER_TZ")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbName, port, tz)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	// Auto-migrate the model to create the table if it doesn't exist
	if autoMergeErr := database.AutoMigrate(&entities.Job{}); autoMergeErr != nil {
		panic(autoMergeErr)
	}
}

func DB() *gorm.DB {
	return database
}
