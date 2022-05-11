package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	const DB_USERNAME = "root"
	const DB_PASSWORD = ""
	const DB_HOST = "localhost"
	const DB_PORT = "3306"
	const DB_NAME = "db_news"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print("Error can't connect to database, error: ", err)

		return nil
	}

	return db
}
