package database

import (
	"todoapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
	"log"
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

	log.Println("Connected to MySQL:", DBGorm)


	DBGorm.AutoMigrate(&models.ToDoItem{})


}
