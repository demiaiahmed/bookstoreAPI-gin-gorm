package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	database.AutoMigrate(&Book{})
	sqlDB, _ := database.DB()
	defer sqlDB.Close()
	DB = database
}
