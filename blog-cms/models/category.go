package models

import (
	"blog-cms/database"
	"gorm.io/gorm"

)


type Category struct {
	ID   uint `gorm:"primarykey"`
	Slug string
	Name string
}

func GetAllCategories(categories *[]Category) *gorm.DB{
	return database.DBGorm.Find(&categories)
}

func (category *Category) Migrate(){
	database.DBGorm.AutoMigrate(category)
}

