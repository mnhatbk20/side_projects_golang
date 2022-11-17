package models

import (
	"blog-cms/database"
	"gorm.io/gorm"
)

type Tag struct {
	ID   uint `gorm:"primarykey"`
	Slug string
	Name string
}


func GetAllTags(tags *[]Tag) *gorm.DB{
	return database.DBGorm.Find(&tags)
}

func (tag *Tag) Migrate(){
	database.DBGorm.AutoMigrate(tag)
}
