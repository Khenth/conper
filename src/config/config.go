package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("conperdb")
	dbPass := os.Getenv("C0np3r*-2010")
	dbHost := os.Getenv("192.168.10.101")
	dbName := os.Getenv("conpercrmrodeo")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}

// DisconnectDB is used to disconnect database
func DissconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to disconnect database")
	}
	dbSQL.Close()
}
