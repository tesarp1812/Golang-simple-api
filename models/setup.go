package models

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func ConnectDatabase() {

// 	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_animal"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	database.AutoMigrate(&Animal{})
// 	DB = database
// }

func ConnectDatabase() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	var err error
	for retries := 5; retries > 0; retries-- {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&Animal{})
}
