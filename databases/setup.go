package database

import (
	"fmt"
	"learn/api/databases/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db failed")
	}

	db.AutoMigrate(&model.User{})
	fmt.Println("Database Connected")

	DB = db
}
