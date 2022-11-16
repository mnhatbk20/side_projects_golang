package database

import (
	"log"
	"os"

	// "blog-cms/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBGorm *gorm.DB

func Connect() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DBConnectionStr")

	DBGorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	// DBGorm.AutoMigrate(&models.BlogInfo{})
	// DBGorm.AutoMigrate(&models.Post{})
	// DBGorm.AutoMigrate(&models.User{})
	// DBGorm.AutoMigrate(&models.Comment{})
	// DBGorm.AutoMigrate(&models.Tag{})
	// DBGorm.AutoMigrate(&models.User{})

}
