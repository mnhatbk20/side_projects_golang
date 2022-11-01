package database

import (
	"log"
	"os"

	"build.cv/models"
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

	DBGorm.AutoMigrate(&models.ContactSocial{})
	DBGorm.AutoMigrate(&models.Education{})
	DBGorm.AutoMigrate(&models.WorkExperience{})
	DBGorm.AutoMigrate(&models.References{})
	DBGorm.AutoMigrate(&models.Hobby{})
	DBGorm.AutoMigrate(&models.Skill{})
	DBGorm.AutoMigrate(&models.User{})

}
