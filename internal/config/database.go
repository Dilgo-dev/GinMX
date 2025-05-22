package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Dilgo-dev/GinMX/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Blog{})
	db.AutoMigrate(&models.User{})
	DB = db

	if os.Getenv("ENV") == "development" && DB.Find(&models.Blog{}).RowsAffected == 0 {
		DBSeed()
	}
}

func GetDB() *gorm.DB {
	return DB
}

func DBSeed() {
	DB.Create(&models.Blog{Title: "Blog 1", Content: "Content 1"})
	DB.Create(&models.Blog{Title: "Blog 2", Content: "Content 2"})
}